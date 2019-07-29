package logger

import (
	"github.com/sirupsen/logrus"
	"os"
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
	systemErrorMessage = Event{100, "Error opening the app config file - %s"}
	spotifyErrorMessage = Event{200, "There was an Error in the Request to Spotify - %s"}
	logFilename = "logfile.json"
)

func (l *StandardLogger) OpenAppLogFile(){
	f, _ := os.OpenFile(logFilename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	l.SetOutput(f)
}

func (l *StandardLogger) SpotifyError(errMsg string) {
	l.OpenAppLogFile()
	l.Errorf(spotifyErrorMessage.message, errMsg)
}

func (l *StandardLogger) AppConfigFileError(errMsg string) {
	l.OpenAppLogFile()
	l.Errorf(systemErrorMessage.message, errMsg)
}
