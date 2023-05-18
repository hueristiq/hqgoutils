package log

import (
	"os"
	"strings"

	"github.com/hueristiq/hqgoutils/log/formatter"
	"github.com/hueristiq/hqgoutils/log/levels"
	"github.com/hueristiq/hqgoutils/log/writer"
)

type Logger struct {
	formatter formatter.Formatter
	maxLevel  levels.LevelInt
	writer    writer.Writer
}

func (logger *Logger) SetFormatter(formatter formatter.Formatter) {
	logger.formatter = formatter
}

func (logger *Logger) SetMaxLevel(level levels.LevelStr) {
	logger.maxLevel = levels.Levels[level]
}

func (logger *Logger) SetWriter(writer writer.Writer) {
	logger.writer = writer
}

func (logger *Logger) Log(event *Event) {
	if event.level > logger.maxLevel {
		return
	}

	var (
		ok    bool
		label string
	)

	if _, ok = event.metadata["label"]; !ok {
		labels := map[levels.LevelInt]string{
			levels.Levels[levels.LevelFatal]:   "FTL",
			levels.Levels[levels.LevelError]:   "ERR",
			levels.Levels[levels.LevelWarning]: "WRN",
			levels.Levels[levels.LevelInfo]:    "INF",
			levels.Levels[levels.LevelDebug]:   "DBG",
		}

		if label, ok = labels[event.level]; ok {
			event.metadata["label"] = label
		}
	}

	event.message = strings.TrimSuffix(event.message, "\n")

	data, err := logger.formatter.Format(&formatter.Log{
		Message:  event.message,
		Level:    event.level,
		Metadata: event.metadata,
	})
	if err != nil {
		return
	}

	if character, ok := event.metadata["rest"]; ok {
		data = appendRest(data, character)
	}

	logger.writer.Write(data, event.level)

	if event.level == levels.Levels[levels.LevelFatal] {
		os.Exit(1)
	}
}

func (logger *Logger) Print() *Event {
	event := &Event{
		logger:   logger,
		level:    levels.LevelInt(-1),
		metadata: make(map[string]string),
	}

	return event
}

func (logger *Logger) Debug() *Event {
	level := levels.Levels[levels.LevelDebug]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

func (logger *Logger) Info() *Event {
	level := levels.Levels[levels.LevelInfo]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

func (logger *Logger) Warning() *Event {
	level := levels.Levels[levels.LevelWarning]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

func (logger *Logger) Error() *Event {
	level := levels.Levels[levels.LevelError]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}

func (logger *Logger) Fatal() *Event {
	level := levels.Levels[levels.LevelFatal]

	event := &Event{
		logger:   logger,
		level:    level,
		metadata: make(map[string]string),
	}

	return event
}
