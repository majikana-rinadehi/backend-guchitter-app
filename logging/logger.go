package logging

import (
	"os"

	"github.com/bloom42/rz-go"
)

var (
	Log rz.Logger
)

// SetupLoggerはLoggerのフィールドをセットアップします。
func SetupLogger(requestId string) {
	hostName, _ := os.Hostname()
	env := os.Getenv("GO_ENV")
	Log = rz.New(
		rz.Fields(
			rz.String("X-Request-ID", requestId),
			rz.String("hostname", hostName),
			rz.String("environment", env),
		),
	)
}
