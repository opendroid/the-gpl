// Package userip stores and retrieves a client's IP address in a context.Context.
// This is the canonical pattern for threading request-scoped values through a call stack
// without extending every function signature. Modelled after golang.org/x/blog/content/context/userip.
package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// key Type to store request scope values
type key int

// userIPKey where userIP is stored
const userIPKey key = 0

// FromRequest extracts the client IP address from req.RemoteAddr.
func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not a IP:port", req.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not a IP:port", req.RemoteAddr)
	}
	return userIP, nil
}

// NewContext returns a copy of ctx with userIP stored under the package-private key.
func NewContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

// FromContext retrieves the IP address stored by NewContext. ok is false if none was set.
func FromContext(ctx context.Context) (net.IP, bool) {
	// ctx.Value returns nil if ctx has no value for "userIPKey"
	// net.IP assertion return ok=false for nil
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}
