package main

type Logger interface {
	Log(debugType string, message string)
}

type ConsoleLogger struct {
}

func (ConsoleLogger) Log(debugType, message string) {
	println(debugType, message)
}

func main() {
	logger := ConsoleLogger{}
	printLogger(logger)

}

func printLogger(logger Logger) {
	logger.Log("debug", "I'm a message!")
}