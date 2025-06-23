package main

import (
	"fmt"
	parser "logger/Parser"
	logchecker "logger/logchecker"
	"os"
)

func main() {
	args := os.Args[1:]
	parser := parser.NewParser()
	parser.ParseInput(args)
	if parser.IsEmpty() {
		panic("No files for work")
	}

	outputFile, err := os.OpenFile(parser.GetOutputFilename(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println("Failed for open output file")
		panic(err)
	}

	defer outputFile.Close()

	logs := logchecker.NewLogchecker()
	fmt.Println("Maximum number of logs per second:",
		logs.MaxLogFrequency(parser.GetInputFilename(), parser.GetFromNum(), parser.GetToNum()))
	fmt.Println("------------------")
	logs.StatsOfLogs(parser.GetInputFilename(), outputFile, parser.GetStatsNum(), parser.GetFromNum(), parser.GetToNum())
	if parser.GetWindowNum() > 0 {
		logs.LogsWindow(parser.GetInputFilename(), outputFile, parser.GetWindowNum(), parser.GetFromNum(), parser.GetToNum())
	}

	logs.LogsFromTo(parser.GetInputFilename(), outputFile, parser.GetFromNum(), parser.GetToNum())
}
