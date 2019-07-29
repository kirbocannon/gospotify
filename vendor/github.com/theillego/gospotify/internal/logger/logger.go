package logger

import (
	"github.com/sirupsen/logrus"
	"strings"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

func FormatSpotifyErrorMessage(err error) (parsedError []string) {
	parsedError = strings.Split(err.Error(), "Response:")

	// strip any new lines found in messages
	// This is currently NOT working
	for i := 0; i < 2; i++ {
		parsedError = append(parsedError, strings.TrimSuffix(parsedError[i], "\n"))
	}

	return
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	baseLogger := logrus.New()

	standardLogger := &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	spotifyErrorMessage = Event{1, "There was an Error in the Request to Spotify - %s", }
)

func (l *StandardLogger) SpotifyError(err string) {
	l.Errorf(spotifyErrorMessage.message, err)
}

