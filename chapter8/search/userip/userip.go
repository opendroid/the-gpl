// userip package is similar to golang.org/x/blog/content/context/userip
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
const userIPKey key  = 0

// FromRequest gets the userIP from a request
func FromRequest(req * http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not a IP:port", req.RemoteAddr)
	}
	userIP  := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not a IP:port", req.RemoteAddr)
	}
	return userIP, nil
}

// NewContext creates a new context with request scoped value ["0"]: "IP"
func NewContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

// FromContext gets  the value from ctx
func FromContext(ctx context.Context) (net.IP, bool) {
	// ctx.Value returns nil if ctx has no value for "userIPKey"
	// net.IP assertion return ok=false for nil
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}