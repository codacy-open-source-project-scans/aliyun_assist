package lock

import (
	"context"
	"time"
)

// ChanMutex provides interfaces of spinlock and trylock implemented by channel.
type ChanMutex interface {
	// Lock acquires the lock.
	// If it is currently held by others, Lock will wait until it has a chance to acquire it.
	Lock()
	// TryLock attempts to acquire the lock without blocking.
	// Return false if someone is holding it now.
	TryLock() bool
	// TryLockWithTimeout attempts to acquire the lock within a period of time.
	// Return false if spending time is more than duration and no chance to acquire it.
	TryLockWithTimeout(time.Duration) bool
	// TryLockWithContext attempts to acquire the lock, blocking until resources
	// are available or ctx is done (timeout or cancellation)
	TryLockWithContext(ctx context.Context) bool
	// Unlock releases the lock
	Unlock()
}

type chanMutex struct {
	lockChan chan struct{}
}

func (m *chanMutex) Lock() {
	m.lockChan <- struct{}{}
}

func (m *chanMutex) Unlock() {
	<-m.lockChan
}

func (m *chanMutex) TryLock() bool {
	select {
	case m.lockChan <- struct{}{}:
		return true
	default:
		return false
	}
}

func (m *chanMutex) TryLockWithContext(ctx context.Context) bool {
	select {
	case m.lockChan <- struct{}{}:
		return true
	case <-ctx.Done():
		// timeout or cancellation
		return false
	}
}

func (m *chanMutex) TryLockWithTimeout(duration time.Duration) bool {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	return m.TryLockWithContext(ctx)
}

// NewChanMutex returns ChanMutex lock
func NewChanMutex() ChanMutex {
	return &chanMutex{
		lockChan: make(chan struct{}, 1),
	}
}
