package infra

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// InitLogrus returns logger that prints to stderr.
// If ENABLE_DEBUG_MODE is set to "1", it outputs
// colored plain text for human read; otherwise it
// outputs JSON-formatted lines for machines parsing.
func InitLogrus() *logrus.Logger {
	log := logrus.New()
	// Out:          os.Stderr,
	// Formatter:    new(TextFormatter),
	// Level:        InfoLevel,

	if os.Getenv(DebugModeEnv) == "1" {
		log.Level = logrus.DebugLevel
		log.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	} else {
		// production setting
		log.Formatter = &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "ts",
				logrus.FieldKeyLevel: "le",
				logrus.FieldKeyMsg:   "msg",
			},
			TimestampFormat: time.RFC3339Nano,
		}
	}

	return log
}
