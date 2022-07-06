package logger

import (
	"os"
	"time"
)

type transaction struct {
	time   time.Time
	route  string
	params map[string]interface{}
}

type Logger struct {
	file             *os.File
	transactionsChan chan<- transaction
	errorsChan       <-chan error
}

func NewLogger(filepath string) (*Logger, error) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}
	return &Logger{file: file, transactionsChan: make(chan transaction), errorsChan: make(chan error)}, nil
}

func (l *Logger) Close() error {
	if l.transactionsChan != nil {
		close(l.transactionsChan)
	}

	return l.file.Close()
}

func (l *Logger) WriteTransaction(transactTime time.Time, route string, params map[string]interface{}) {
}
