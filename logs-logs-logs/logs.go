package logs

import "strings"

// Message extracts the message from the provided log line.
func Message(line string) string {
	return strings.TrimSpace(strings.Split(line, ":")[1])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return len(strings.Split(Message(line), ""))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	return strings.ToLower(strings.Split(strings.Split(line, "[")[1], "]")[0])
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return Message(line) + " (" + LogLevel(line) + ")"
}
