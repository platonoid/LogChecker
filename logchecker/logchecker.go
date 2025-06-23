package logchecker

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type LogChecker struct {
	logTimeFrequency map[int64]int
	logFrequency     map[string]int
}

type KeyValue struct {
	key   string
	value int
}

func NewLogchecker() LogChecker {
	return LogChecker{logTimeFrequency: make(map[int64]int), logFrequency: make(map[string]int)}
}

func GetLogTime(line string) int64 {
	var first_boarder int = strings.Index(line, "[")
	var second_boarder int = strings.Index(line, "]")
	if first_boarder < 0 || second_boarder < 0 {
		panic("Bad log")
	}
	var logtime string = line[first_boarder+1 : second_boarder]
	time_zone_index := strings.Index(logtime, "-")
	if time_zone_index < 0 {
		panic("Bad log")
	}
	example := "02/Jan/2006:15:04:05"
	logtime = logtime[0 : time_zone_index-1]
	t, err := time.Parse(example, logtime)
	if err != nil {
		panic(err)
	}
	return t.Unix()
}

func GetLog(line string) string {
	var first_boarder int = strings.Index(line, "\"")
	if first_boarder < 0 {
		panic("Bad log")
	}
	line = line[first_boarder+1:]
	var second_boarder int = strings.Index(line, "\"")
	if second_boarder < 0 {
		panic("Bad log")
	}
	return line[0:second_boarder]
}

func (logger LogChecker) MaxLogFrequency(input_path string, from int64, to int64) int {
	file, err := os.Open(input_path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		logtime := GetLogTime(scanner.Text())
		if logtime < from || logtime > to {
			continue
		}
		logger.logTimeFrequency[logtime]++
	}
	var maximum int
	for _, element := range logger.logTimeFrequency {
		if maximum < element {
			maximum = element
		}
	}
	return maximum
}

func (logger LogChecker) StatsOfLogs(input_path string, outputFile *os.File, stats int, from int64, to int64) {
	file, err := os.Open(input_path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currstr := scanner.Text()
		logtime := GetLogTime(currstr)
		if logtime < from || logtime > to {
			continue
		}
		logger.logFrequency[GetLog(currstr)]++
	}
	var logsKV []KeyValue

	for key, value := range logger.logFrequency {
		logsKV = append(logsKV, KeyValue{key, value})
	}

	sort.Slice(logsKV, func(i, j int) bool {
		return logsKV[i].value > logsKV[j].value
	})

	fmt.Println(min(len(logsKV), stats), "most recent logs")
	fmt.Fprintln(outputFile, min(len(logsKV), stats), "most recent logs")
	for i := 0; i < min(len(logsKV), stats); i++ {
		fmt.Fprintln(outputFile, logsKV[i].key, logsKV[i].value)
		fmt.Println(logsKV[i].key, logsKV[i].value)
	}
	fmt.Fprintln(outputFile, "------------------")
	fmt.Println("------------------")
}

func (logger LogChecker) LogsFromTo(input_path string, outputFile *os.File, from int64, to int64) {
	file, err := os.Open(input_path)

	if err != nil {
		panic("os.File does not exist")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Fprintln(outputFile, "Log between", max(0, from), "-", to)
	fmt.Fprintln(outputFile, "Format logs : time")
	for scanner.Scan() {
		line := scanner.Text()
		if GetLogTime(line) < from {
			continue
		}

		if GetLogTime(line) > to {
			continue
		}

		fmt.Fprintln(outputFile, GetLog(line), ":", GetLogTime(line))
	}
	fmt.Fprintln(outputFile, "------------------")
}

func (logger LogChecker) LogsWindow(input_path string, outputFile *os.File, window int64, from int64, to int64) {
	file, err := os.Open(input_path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxLogs := 0

	var logsTime []int64
	for scanner.Scan() {
		time := GetLogTime(scanner.Text())
		if from > time || time > to {
		}
		logsTime = append(logsTime, time)

		for len(logsTime) > 0 && (time-logsTime[0]) > window {
			logsTime = logsTime[1:]
		}

		if len(logsTime) > maxLogs {
			maxLogs = len(logsTime)
		}
	}

	fmt.Println("Maximum number of requests in a window of length", window)
	fmt.Fprintln(outputFile, "Maximum number of requests in a window of length", window)
	fmt.Println(logsTime[0], "-", logsTime[len(logsTime)-1])
	fmt.Fprintln(outputFile, logsTime[0], "-", logsTime[len(logsTime)-1])
	fmt.Println("Count of this logs", maxLogs)
	fmt.Fprintln(outputFile, "Count of this logs", maxLogs)
	fmt.Fprintln(outputFile, "------------------")
	fmt.Println("------------------")
}
