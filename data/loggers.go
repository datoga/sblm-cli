package data

import (
	"sort"
	"strings"

	tm "github.com/buger/goterm"
)

//GetSortedLoggerNames gets a list of the loggers sorted alphabetically
func (loggers Loggers) GetSortedLoggerNames() []string {
	var keys []string

	for logger := range loggers {
		keys = append(keys, logger)
	}

	sort.Strings(keys)

	return keys
}

func (loggers Loggers) String() string {
	return loggers.Raw()
}

//Raw returns the raw representation of a list of Loggers
func (loggers Loggers) Raw() string {
	sortedLoggers := loggers.GetSortedLoggerNames()

	var str strings.Builder

	for i, logger := range sortedLoggers {
		str.WriteString(logger + ": " + loggers[logger].EffectiveLevel)

		if i < len(sortedLoggers)-1 {
			str.WriteString("\n")
		}
	}

	return str.String()
}

func getLongerValue(arr []string) int {
	max := -1

	for _, v := range arr {
		if len(v) > max {
			max = len(v)
		}
	}

	return max
}

//PrettyPrint returns a pretty representation of a list of Loggers
func (loggers Loggers) PrettyPrint(maxWidth int) {
	longerValue := getLongerValue(loggers.GetSortedLoggerNames())

	if longerValue > maxWidth {
		longerValue = maxWidth
	}

	for _, logger := range loggers.GetSortedLoggerNames() {
		level := loggers[logger].EffectiveLevel

		if level == "" {
			continue
		}

		if len(logger) > maxWidth {
			logger = logger[:maxWidth-2] + ".."
		}

		differenceWithMax := longerValue - len(logger)

		var str strings.Builder

		str.WriteString(logger)

		str.WriteString(strings.Repeat(" ", differenceWithMax))

		str.WriteString("\t")

		for _, predefinedLevel := range PredefinedLevels {
			if level == predefinedLevel {
				str.WriteString(tm.Color(tm.Bold(tm.Background(predefinedLevel, tm.BLUE)), tm.WHITE))
			} else {
				str.WriteString(predefinedLevel)
			}

			str.WriteString("\t")
		}

		tm.Println(str.String())

		tm.Flush()
	}
}
