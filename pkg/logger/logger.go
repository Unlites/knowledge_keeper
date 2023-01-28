package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(message string, args ...interface{})
	Error(message string, err error, args ...interface{})
	Fatal(message string, err error, args ...interface{})
	AddKey(key, value string)
}

var _ Logger = (*logger)(nil)

type logger struct {
	log *zerolog.Logger
}

func NewLogger() *logger {
	log := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "2006:01:02 15:04:05",
	}).With().Timestamp().Logger()
	return &logger{
		log: &log,
	}
}

func (l *logger) Info(message string, args ...interface{}) {
	if len(args) == 0 {
		l.log.Info().Msg(message)
	} else {
		l.log.Info().Msgf(message, args...)
	}
}

func (l *logger) Error(message string, err error, args ...interface{}) {
	if len(args) == 0 {
		l.log.Err(err).Msg("")
	} else {
		l.log.Err(err).Msgf(err.Error(), args...)
	}
}

func (l *logger) Fatal(message string, err error, args ...interface{}) {
	if len(args) == 0 {
		l.log.Fatal().Msg(message + ": " + err.Error())
	} else {
		l.log.Fatal().Msgf(message+": "+err.Error(), args...)
	}
}

func (l *logger) AddKey(key, value string) {
	l.log.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str(key, value)
	})
}
