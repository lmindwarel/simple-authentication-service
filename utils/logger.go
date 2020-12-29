package utils

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

//Format for console stdout
var loggerFormatStdout = logging.MustStringFormatter(
	`%{color}%{time:2006-01-02 15:04:05.000} [%{module} / %{shortfunc}] ▶ %{level:.1s} %{color:reset}%{message}`,
)

//Format for file
var loggerFormatFile = logging.MustStringFormatter(
	`%{time:2006-01-02 15:04:05.000} [%{module} / %{shortfunc}] ▶ %{level:.1s} %{message}`,
)

// InitLogger intialize the logger and create log file at the given path
func InitLogger(filePath string) {

	var backends []logging.Backend

	// Create console output
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, loggerFormatStdout)
	backends = append(backends, backendFormatter)

	// Create a file backend
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0640)

	if err != nil {
		panic(fmt.Sprintf("Unable to create log file: %s", err))
	}

	backend = logging.NewLogBackend(file, "", 0)
	backendFormatter = logging.NewBackendFormatter(backend, loggerFormatFile)
	backends = append(backends, backendFormatter)

	logging.SetBackend(backends...)
}

// GetLogger return a logger
func GetLogger(identifier string) *logging.Logger {
	return logging.MustGetLogger(identifier)
}
