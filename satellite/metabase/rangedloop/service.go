// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package rangedloop

import (
	"context"
	"fmt"
	"time"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/common/errs2"
	"storj.io/common/sync2"
	"storj.io/storj/satellite/metabase/segmentloop"
)

var (
	mon = monkit.Package()
)

// Config contains configurable values for the shared loop.
type Config struct {
	Parallelism        int           `help:"how many chunks of segments to process in parallel" default:"2"`
	BatchSize          int           `help:"how many items to query in a batch" default:"2500"`
	AsOfSystemInterval time.Duration `help:"as of system interval" releaseDefault:"-5m" devDefault:"-1us" testDefault:"-1us"`
	Interval           time.Duration `help:"how often to run the loop" releaseDefault:"2h" devDefault:"10s" testDefault:"10s"`
}

// Service iterates through all segments and calls the attached observers for every segment
//
// architecture: Service
type Service struct {
	log       *zap.Logger
	config    Config
	provider  RangeSplitter
	observers []Observer

	Loop *sync2.Cycle
}

// NewService creates a new instance of the ranged loop service.
func NewService(log *zap.Logger, config Config, provider RangeSplitter, observers []Observer) *Service {
	return &Service{
		log:       log,
		config:    config,
		provider:  provider,
		observers: observers,
		Loop:      sync2.NewCycle(config.Interval),
	}
}

// observerState contains information to manage an observer during a loop iteration.
type observerState struct {
	observer       Observer
	rangeObservers []*rangeObserverState
	// err is the error that occurred during the observer's Start method.
	// If err is set, the observer will be skipped during the loop iteration.
	err error
}

type rangeObserverState struct {
	rangeObserver Partial
	duration      time.Duration
	// err is the error that is returned by the observer's Fork or Process method.
	// If err is set, the range observer will be skipped during the loop iteration.
	err error
}

// ObserverDuration reports back on how long it took the observer to process all the segments.
type ObserverDuration struct {
	Observer Observer
	// Duration is set to -1 when the observer has errored out
	// so someone watching metrics can tell that something went wrong.
	Duration time.Duration
}

// Run starts the looping service.
func (service *Service) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	service.log.Info("ranged loop initialized")

	return service.Loop.Run(ctx, func(ctx context.Context) error {
		service.log.Info("ranged loop started")
		_, err := service.RunOnce(ctx)
		if err != nil {
			service.log.Error("ranged loop failure", zap.Error(err))

			if errs2.IsCanceled(err) {
				return err
			}

			if ctx.Err() != nil {
				return errs.Combine(err, ctx.Err())
			}

			mon.Event("rangedloop_error") //mon:locked
		}

		service.log.Info("ranged loop finished")
		return nil
	})
}

// RunOnce goes through one time and sends information to observers.
func (service *Service) RunOnce(ctx context.Context) (observerDurations []ObserverDuration, err error) {
	defer mon.Task()(&ctx)(&err)

	observerStates, err := startObservers(ctx, service.log, service.observers)
	if err != nil {
		return nil, err
	}

	rangeProviders, err := service.provider.CreateRanges(service.config.Parallelism, service.config.BatchSize)
	if err != nil {
		return nil, err
	}

	group := errs2.Group{}
	for _, rangeProvider := range rangeProviders {
		rangeObservers := []*rangeObserverState{}
		for i, observerState := range observerStates {
			if observerState.err != nil {
				continue
			}
			rangeObserver, err := observerState.observer.Fork(ctx)
			rangeState := &rangeObserverState{
				rangeObserver: rangeObserver,
				err:           err,
			}
			rangeObservers = append(rangeObservers, rangeState)
			observerStates[i].rangeObservers = append(observerStates[i].rangeObservers, rangeState)
		}

		// Create closure to capture loop variables.
		group.Go(createGoroutineClosure(ctx, rangeProvider, rangeObservers))
	}

	// Improvement: stop all ranges when one has an error.
	errList := group.Wait()
	if errList != nil {
		return nil, errs.Combine(errList...)
	}

	return finishObservers(ctx, service.log, observerStates)
}

func createGoroutineClosure(ctx context.Context, rangeProvider SegmentProvider, states []*rangeObserverState) func() error {
	return func() (err error) {
		defer mon.Task()(&ctx)(&err)

		return rangeProvider.Iterate(ctx, func(segments []segmentloop.Segment) error {
			// check for cancellation every segment batch
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return processBatch(ctx, states, segments)
			}
		})
	}
}

func startObservers(ctx context.Context, log *zap.Logger, observers []Observer) (observerStates []observerState, err error) {
	startTime := time.Now()

	for _, obs := range observers {
		observerStates = append(observerStates, startObserver(ctx, log, startTime, obs))
	}

	return observerStates, nil
}

func startObserver(ctx context.Context, log *zap.Logger, startTime time.Time, observer Observer) observerState {
	err := observer.Start(ctx, startTime)

	if err != nil {
		log.Error(
			"Starting observer failed. This observer will be excluded from this run of the ranged segment loop.",
			zap.String("observer", fmt.Sprintf("%T", observer)),
			zap.Error(err),
		)
	}

	return observerState{
		observer: observer,
		err:      err,
	}
}

func finishObservers(ctx context.Context, log *zap.Logger, observerStates []observerState) (observerDurations []ObserverDuration, err error) {
	for _, state := range observerStates {
		observerDurations = append(observerDurations, finishObserver(ctx, log, state))
	}

	sendObserverDurations(observerDurations)

	return observerDurations, nil
}

// Iterating over the segments is done.
// This is the reduce step.
func finishObserver(ctx context.Context, log *zap.Logger, state observerState) ObserverDuration {
	if state.err != nil {
		return ObserverDuration{
			Observer: state.observer,
			Duration: -1 * time.Second,
		}
	}
	for _, rangeObserver := range state.rangeObservers {
		if rangeObserver.err != nil {
			log.Error(
				"Observer failed during Process(), it will not be finalized in this run of the ranged segment loop",
				zap.String("observer", fmt.Sprintf("%T", state.observer)),
				zap.Error(rangeObserver.err),
			)
			return ObserverDuration{
				Observer: state.observer,
				Duration: -1 * time.Second,
			}
		}
	}

	var duration time.Duration
	for _, rangeObserver := range state.rangeObservers {
		err := state.observer.Join(ctx, rangeObserver.rangeObserver)
		if err != nil {
			log.Error(
				"Observer failed during Join(), it will not be finalized in this run of the ranged segment loop",
				zap.String("observer", fmt.Sprintf("%T", state.observer)),
				zap.Error(rangeObserver.err),
			)
			return ObserverDuration{
				Observer: state.observer,
				Duration: -1 * time.Second,
			}
		}
		duration += rangeObserver.duration
	}

	err := state.observer.Finish(ctx)
	if err != nil {
		log.Error(
			"Observer failed during Finish()",
			zap.String("observer", fmt.Sprintf("%T", state.observer)),
			zap.Error(err),
		)
		return ObserverDuration{
			Observer: state.observer,
			Duration: -1 * time.Second,
		}
	}

	return ObserverDuration{
		Duration: duration,
		Observer: state.observer,
	}
}

func processBatch(ctx context.Context, states []*rangeObserverState, segments []segmentloop.Segment) (err error) {
	for _, state := range states {
		if state.err != nil {
			// this observer has errored in a previous batch
			continue
		}
		start := time.Now()
		err := state.rangeObserver.Process(ctx, segments)
		state.duration += time.Since(start)
		if err != nil {
			// unsure if this is necessary here
			if errs2.IsCanceled(err) {
				return err
			}
			state.err = err
		}
	}
	return nil
}
