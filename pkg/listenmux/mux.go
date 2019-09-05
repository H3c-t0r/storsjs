// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package listenmux

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/zeebo/errs"
)

// Closed is returned by routed listeners when the mux is closed.
var Closed = errs.New("listener closed")

// Mux lets one multiplex a listener into different listeners based on the first
// bytes sent on the connection.
type Mux struct {
	base      net.Listener
	prefixLen int
	addr      net.Addr
	def       *listener

	mu     sync.Mutex
	routes map[string]*listener

	once sync.Once
	done chan struct{}
	err  error
}

// New creates a mux that reads the prefixLen bytes from any connections Accepted by the
// passed in listener and dispatches to the appropriate route.
func New(base net.Listener, prefixLen int) *Mux {
	addr := base.Addr()
	return &Mux{
		base:      base,
		prefixLen: prefixLen,
		addr:      addr,
		def:       newListener(addr),

		routes: make(map[string]*listener),

		done: make(chan struct{}),
	}
}

//
// set up the routes
//

// Default returns the net.Listener that is used if no route matches.
func (m *Mux) Default() net.Listener { return m.def }

// Route returns a listener that will be used if the first bytes are the given prefix. The
// length of the prefix must match the original passed in prefixLen.
func (m *Mux) Route(prefix string) net.Listener {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(prefix) != m.prefixLen {
		panic(fmt.Sprintf("invalid prefix: has %d but needs %d bytes", len(prefix), m.prefixLen))
	}

	lis, ok := m.routes[prefix]
	if !ok {
		lis = newListener(m.addr)
		m.routes[prefix] = lis
		go m.monitorListener(prefix, lis)
	}
	return lis
}

//
// run the muxer
//

// Run calls listen on the provided listener and passes connections to the routed
// listeners.
func (m *Mux) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go m.monitorBase()
	go m.monitorSignal()
	go m.monitorContext(ctx)

	<-m.done

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, lis := range m.routes {
		<-lis.done
	}

	return m.err
}

func (m *Mux) monitorSignal() {
	<-m.done
	// TODO(jeff): do we care about this error?
	_ = m.base.Close()
}

func (m *Mux) monitorContext(ctx context.Context) {
	<-ctx.Done()
	m.once.Do(func() { close(m.done) })
}

func (m *Mux) monitorListener(prefix string, lis *listener) {
	select {
	case <-m.done:
		lis.once.Do(func() {
			if m.err != nil {
				lis.err = m.err
			} else {
				lis.err = Closed
			}
			close(lis.done)
		})
	case <-lis.done:
	}
	m.mu.Lock()
	delete(m.routes, prefix)
	m.mu.Unlock()
}

func (m *Mux) monitorBase() {
	for {
		conn, err := m.base.Accept()
		switch {
		case err != nil:
			// TODO(jeff): temporary errors?
			m.once.Do(func() {
				m.err = err
				close(m.done)
			})
			return
		case conn == nil:
			<-m.done
			return
		}
		go m.routeConn(conn)
	}
}

func (m *Mux) routeConn(conn net.Conn) {
	buf := make([]byte, m.prefixLen)
	if _, err := io.ReadFull(conn, buf); err != nil {
		// TODO(jeff): how to handle this error?
		return
	}

	m.mu.Lock()
	lis, ok := m.routes[string(buf)]
	if !ok {
		lis = m.def
		conn = newPrefixConn(buf, conn)
	}
	m.mu.Unlock()

	// TODO(jeff): a timeout for the listener to get to the conn?

	select {
	case <-lis.done:
		// TODO(jeff): better way to signal to the caller the listener is closed?
		_ = conn.Close()
	case lis.Conns() <- conn:
	}
}
