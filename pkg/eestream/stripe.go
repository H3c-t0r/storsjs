// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package eestream

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"

	"github.com/vivint/infectious"
)

// StripeReader can read and decodes stripes from a set of readers
type StripeReader struct {
	scheme ErasureScheme
	cond   *sync.Cond
	bufs   map[int]*PieceBuffer
	inbufs [][]byte
	inmap  map[int][]byte
	errmap map[int]error
}

// NewStripeReader creates a new StripeReader from the given readers, erasure
// scheme and max buffer memory.
func NewStripeReader(rs map[int]io.ReadCloser, es ErasureScheme, mbm int) *StripeReader {
	bufSize := mbm / es.TotalCount()
	bufSize -= bufSize % es.EncodedBlockSize()
	if bufSize < es.EncodedBlockSize() {
		bufSize = es.EncodedBlockSize()
	}

	r := &StripeReader{
		scheme: es,
		cond:   sync.NewCond(&sync.Mutex{}),
		bufs:   make(map[int]*PieceBuffer, es.TotalCount()),
		inbufs: make([][]byte, es.TotalCount()),
		inmap:  make(map[int][]byte, es.TotalCount()),
		errmap: make(map[int]error, es.TotalCount()),
	}

	for i := 0; i < es.TotalCount(); i++ {
		r.inbufs[i] = make([]byte, es.EncodedBlockSize())
		r.bufs[i] = NewPieceBuffer(make([]byte, bufSize), es.EncodedBlockSize(), r.cond)
	}

	// Kick off a goroutine each reader to be copied into a PieceBuffer.
	for i, buf := range r.bufs {
		go func(r io.Reader, buf *PieceBuffer) {
			_, err := io.Copy(buf, r)
			if err != nil {
				buf.SetError(err)
				return
			}
			buf.SetError(io.EOF)
		}(rs[i], buf)
	}

	return r
}

// Close closes the StripeReader and all PieceBuffers.
func (r *StripeReader) Close() error {
	errs := make(chan error, len(r.bufs))
	for _, buf := range r.bufs {
		go func(c io.Closer) {
			errs <- c.Close()
		}(buf)
	}
	var first error
	for range r.bufs {
		err := <-errs
		if err != nil && first == nil {
			first = Error.Wrap(err)
		}
	}
	return first
}

// ReadStripe reads and decodes the num-th stripe and concatenates it to p. The
// return value is the updated byte slice.
func (r *StripeReader) ReadStripe(num int64, p []byte) ([]byte, error) {
	for i := range r.inmap {
		delete(r.inmap, i)
	}

	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	for {
		for r.readAvailableShares(num) == 0 {
			r.cond.Wait()
		}
		if r.hasEnoughShares() {
			if !r.canRead() {
				return nil, r.combineErrs()
			}
			out, err := r.scheme.Decode(p, r.inmap)
			if err != nil {
				if r.shouldWaitForMore(err) {
					continue
				}
				return nil, err
			}
			return out, nil
		}
	}
}

// readAvailableShares reads the available num-th erasure shares from the piece
// buffers without blocking. The return value n is the number of erasure shares
// read.
func (r *StripeReader) readAvailableShares(num int64) (n int) {
	for i := 0; i < len(r.bufs); i++ {
		if r.inmap[i] != nil || r.errmap[i] != nil {
			continue
		}
		if r.bufs[i].HasShare(num) {
			err := r.bufs[i].ReadShare(num, r.inbufs[i])
			if err != nil {
				r.errmap[i] = err
			} else {
				r.inmap[i] = r.inbufs[i]
			}
			n++
		}
	}
	return n
}

// hasEnoughShares check if there are enough erasure shares read to attempt
// a decode.
func (r *StripeReader) hasEnoughShares() bool {
	return r.canRead() || len(r.inmap)+len(r.errmap) >= r.scheme.TotalCount()
}

// canRead returns if there's enough erasure shares for a successful decode
func (r *StripeReader) canRead() bool {
	return len(r.inmap) >= r.scheme.RequiredCount()+1
}

// shouldWaitForMore checks the returned decode error if it makes sense to wait
// for more erasure shares to attempt an error correction.
func (r *StripeReader) shouldWaitForMore(err error) bool {
	// check if the error is due to error detection
	if !infectious.NotEnoughShares.Contains(err) &&
		!infectious.TooManyErrors.Contains(err) {
		return false
	}
	// check if there are more input buffers to wait for
	return len(r.inmap)+len(r.errmap) < r.scheme.TotalCount()
}

// combineErrs makes a useful error message from the errors in errmap.
// combineErrs always returns an error.
func (r *StripeReader) combineErrs() error {
	if len(r.errmap) == 0 {
		return Error.New("programmer error: no errors to combine")
	}
	errstrings := make([]string, 0, len(r.errmap))
	for i, err := range r.errmap {
		errstrings = append(errstrings,
			fmt.Sprintf("\nerror retrieving piece %02d: %v", i, err))
	}
	sort.Strings(errstrings)
	return Error.New("failed to download stripe: %s",
		strings.Join(errstrings, ""))
}
