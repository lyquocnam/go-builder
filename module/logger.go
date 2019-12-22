package module

import (
	"fmt"
	"github.com/labstack/gommon/color"
	"log"
)
type Logger interface {
	Infof(format string, params ...interface{})
	Warnf(format string, params ...interface{})
}

type logger struct {
	logger *color.Color
}

func NewLogger() *logger {
	l := color.New()
	l.SetOutput(log.Writer())

	return &logger{
		logger: l,
	}
}

func (s *logger) Infof(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	s.logger.Println(s.logger.Green(msg))
}

func (s *logger) Warnf(format string, params ...interface{}) {
	msg := fmt.Sprintf(format, params...)
	s.logger.Println(s.logger.Yellow(msg))
}