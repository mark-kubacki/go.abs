// Provides functions to compute ›absolute differences‹.
//
// Golang Go's missing »left-pad«.
//
// No copyright: triviality.
package abs // import "plugin.hosting/go/abs"

// Returns the absolute value of n.
//
// Branchless, constant time.
func Abs64(n int64) uint64 {
	m := n >> (64 - 1)
	return uint64((n ^ m) - m)
}

// Returns the absolute value of n.
//
// Branchless, constant time, boring.
func Abs32(n int32) uint32 {
	m := n >> (32 - 1)
	return uint32((n ^ m) - m)
}
