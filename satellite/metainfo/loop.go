// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package metainfo

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storage"
)

var (
	// LoopError is a standard error class for this component.
	LoopError = errs.Class("metainfo loop error")
	// LoopClosedError is a loop closed error
	LoopClosedError = LoopError.New("loop closed")
)

// Observer is an interface defining an observer that can subscribe to the metainfo loop.
//
// architecture: Observer
type Observer interface {
	RemoteSegment(context.Context, ScopedPath, *pb.Pointer) error
	RemoteObject(context.Context, ScopedPath, *pb.Pointer) error
	InlineSegment(context.Context, ScopedPath, *pb.Pointer) error
}

// ScopedPath contains full expanded information about the path
type ScopedPath struct {
	ProjectPath string
	ProjectID   uuid.UUID
	BucketName  string

	// TODO: should these be a []byte?

	// Raw is the same path as pointerDB is using.
	Raw string
}

type observerContext struct {
	Observer
	ctx  context.Context
	done chan error
}

func (observer *observerContext) HandleError(err error) bool {
	if err != nil {
		observer.done <- err
		observer.Finish()
		return true
	}
	return false
}

func (observer *observerContext) Finish() {
	close(observer.done)
}

func (observer *observerContext) Wait() error {
	return <-observer.done
}

// LoopConfig contains configurable values for the metainfo loop.
type LoopConfig struct {
	CoalesceDuration time.Duration `help:"how long to wait for new observers before starting iteration" releaseDefault:"5s" devDefault:"5s"`
}

// Loop is a metainfo loop service.
//
// architecture: Service
type Loop struct {
	config   LoopConfig
	metainfo *Service
	join     chan *observerContext
	done     chan struct{}
}

// NewLoop creates a new metainfo loop service.
func NewLoop(config LoopConfig, metainfo *Service) *Loop {
	return &Loop{
		metainfo: metainfo,
		config:   config,
		join:     make(chan *observerContext),
		done:     make(chan struct{}),
	}
}

// Join will join the looper for one full cycle until completion and then returns.
// On ctx cancel the observer will return without completely finishing.
// Only on full complete iteration it will return nil.
// Safe to be called concurrently.
func (loop *Loop) Join(ctx context.Context, observer Observer) (err error) {
	defer mon.Task()(&ctx)(&err)

	obsContext := &observerContext{
		Observer: observer,
		ctx:      ctx,
		done:     make(chan error),
	}

	select {
	case loop.join <- obsContext:
	case <-ctx.Done():
		return ctx.Err()
	case <-loop.done:
		return LoopClosedError
	}

	return obsContext.Wait()
}

// Run starts the looping service.
// It can only be called once, otherwise a panic will occur.
func (loop *Loop) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	defer close(loop.done)

	for {
		err := loop.runOnce(ctx)
		if err != nil {
			return err
		}
	}
}

// runOnce goes through metainfo one time and sends information to observers.
func (loop *Loop) runOnce(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	var observers []*observerContext
	defer func() {
		if err != nil {
			for _, observer := range observers {
				observer.HandleError(err)
			}
			return
		}
		for _, observer := range observers {
			observer.Finish()
		}
	}()

	// wait for the first observer, or exit because context is canceled
	select {
	case observer := <-loop.join:
		observers = append(observers, observer)
	case <-ctx.Done():
		return ctx.Err()
	}

	// after the first observer is found, set timer for CoalesceDuration and add any observers that try to join before the timer is up
	timer := time.NewTimer(loop.config.CoalesceDuration)
waitformore:
	for {
		select {
		case observer := <-loop.join:
			observers = append(observers, observer)
		case <-timer.C:
			break waitformore
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	err = loop.metainfo.Iterate(ctx, "", "", true, false,
		func(ctx context.Context, it storage.Iterator) error {
			var item storage.ListItem

			// iterate over every segment in metainfo
			for it.Next(ctx, &item) {
				rawPath := item.Key.String()
				pointer := &pb.Pointer{}

				err = proto.Unmarshal(item.Value, pointer)
				if err != nil {
					return LoopError.New("unexpected error unmarshalling pointer %s", err)
				}

				pathElements := storj.SplitPath(rawPath)
				if len(pathElements) < 3 {
					return LoopError.New("invalid path %q", rawPath)
				}

				isLastSegment := pathElements[1] == "l"

				path := ScopedPath{
					Raw:         rawPath,
					ProjectPath: pathElements[0],
					BucketName:  pathElements[2],
				}

				projectID, err := uuid.Parse(path.ProjectPath)
				if err != nil {
					return LoopError.Wrap(err)
				}
				path.ProjectID = *projectID

				nextObservers := observers[:0]
				for _, observer := range observers {
					keepObserver := handlePointer(ctx, observer, path, isLastSegment, pointer)
					if keepObserver {
						nextObservers = append(nextObservers, observer)
					}
				}

				observers = nextObservers
				if len(observers) == 0 {
					return nil
				}

				// if context has been canceled exit. Otherwise, continue
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
				}
			}
			return nil
		})

	return err
}

// handlePointer deals with a pointer for a single observer
// if there is some error on the observer, handle the error and return false. Otherwise, return true
func handlePointer(ctx context.Context, observer *observerContext, path ScopedPath, isLastSegment bool, pointer *pb.Pointer) bool {
	remote := pointer.GetRemote()
	if remote != nil {
		if observer.HandleError(observer.RemoteSegment(ctx, path, pointer)) {
			return false
		}
		if isLastSegment {
			if observer.HandleError(observer.RemoteObject(ctx, path, pointer)) {
				return false
			}
		}
	} else if observer.HandleError(observer.InlineSegment(ctx, path, pointer)) {
		return false
	}

	select {
	case <-observer.ctx.Done():
		observer.HandleError(observer.ctx.Err())
		return false
	default:
	}

	return true
}

// Wait waits for run to be finished.
// Safe to be called concurrently.
func (loop *Loop) Wait() {
	<-loop.done
}
