// Package infra provides infrastructure related utils
// to share among hipster shop services.
package infra

import "errors"

// environment variables
const (
	JaegerAddrEnv = "JAEGER_SERVICE_ADDR"
)

// exported errors
var (
	ErrNoEnv = errors.New("missing env variable")
)
