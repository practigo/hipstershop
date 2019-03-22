// Package infra provides infrastructure related utils
// to share among hipster shop services.
package infra

import "errors"

// environment variables
const (
	JaegerAddrEnv = "JAEGER_SERVICE_ADDR"
	DebugModeEnv  = "ENABLE_DEBUG_MODE"
	AppPortEnv    = "APP_PORT"
)

// exported errors
var (
	ErrNoEnv = errors.New("missing env variable")
)
