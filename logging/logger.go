package logging

import (
	"github.com/bloom42/rz-go"
)

var (
	Log rz.Logger
)

func SetupLogger(requestId string) {
	Log = rz.New(
		rz.Fields(
			rz.String("X-Request-ID", requestId),
			rz.String("hostname", ""),
			rz.String("environment", ""),
		),
	)
}
