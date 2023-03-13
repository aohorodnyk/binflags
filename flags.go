package binflags

// Flags is an interface to standardize the interface for all binflags implementations.
type Flags interface {
	// IsSet returns true if the flag is set.
	// If the flag is not set, then it returns false.
	IsSet(bitNumber uint) bool

	// Set sets the flag.
	// This function is idempotent.
	Set(bitNumber uint) bool

	// Unset unsets the flag.
	// This function is idempotent.
	Unset(bitNumber uint) bool
}
