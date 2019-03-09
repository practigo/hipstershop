package infra_test

import (
	"os"
	"testing"

	"github.com/practigo/hipstershop/infra"
)

func TestInitJaegerTracing(t *testing.T) {
	var err error
	if os.Getenv(infra.JaegerAddrEnv) == "" {
		err = infra.InitJaegerTracing("test-service")
		if err != infra.ErrNoEnv {
			t.Error("should return no env error")
		}
		os.Setenv(infra.JaegerAddrEnv, "localhost:6831")
		defer os.Setenv(infra.JaegerAddrEnv, "")
	}
	err = infra.InitJaegerTracing("test-service")
	// should have no error, if you have run
	// docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:1.10
	t.Log(err)
}
