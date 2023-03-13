package binflags

import (
	"sync"
)

// NewSync returns new instance of binflags.Sync.
// It is safe to use it concurrently.
// As a parameter it takes a pointer to binflags.Flags implementation.
// It implements binflags.Flags interface.
func NewSync(flags Flags) *Sync {
	return &Sync{
		flags: flags,
		mu:    sync.RWMutex{},
	}
}

// Sync is a wrapper around binflags.Flags implementation.
// It is safe to use it concurrently.
type Sync struct {
	mu    sync.RWMutex
	flags Flags
}

// IsSet returns true if bitNumber is set.
// It is safe to use it concurrently.
// If the flag is not set, then it returns false.
// This function has O(1) complexity.
func (s *Sync) IsSet(bitNumber uint) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.flags.IsSet(bitNumber)
}

// Set sets bitNumber.
// It is safe to use it concurrently.
// This function is idempotent.
// This function has O(1) complexity.
// It always returns true.
func (s *Sync) Set(bitNumber uint) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.flags.Set(bitNumber)
}

// Unset unsets bitNumber.
// It is safe to use it concurrently.
// This function is idempotent.
// This function has O(1) complexity.
// It always returns true.
func (s *Sync) Unset(bitNumber uint) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.flags.Unset(bitNumber)
}
