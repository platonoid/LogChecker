package parser

import (
	. "math"
	. "strconv"
)

type Parser struct {
	input_filename  string
	output_filename string
	stats           int //10
	window          int64
	begin           int64
	end             int64
}

func NewParser() Parser {
	return Parser{stats: 10, begin: MinInt64, end: MaxInt64}
}

func (parser Parser) GetInputFilename() string {
	return parser.input_filename
}

func (parser Parser) GetOutputFilename() string {
	return parser.output_filename
}

func (parser Parser) GetStatsNum() int {
	return parser.stats
}

func (parser Parser) GetFromNum() int64 {
	return parser.begin
}

func (parser Parser) GetToNum() int64 {
	return parser.end
}

func (parser Parser) GetWindowNum() int64 {
	return parser.window
}

var long_to_low = map[string]string{
	"--output": "-o",
	"--input":  "-i",
	"--stats":  "-s",
	"--window": "-w",
	"--from":   "-f",
	"--to":     "-t",
}

func (parser *Parser) ParseInput(arr []string) {
	for i := 0; i < len(arr); i++ {
		curr_key := long_to_low[arr[i]]
		if i != len(arr)-1 {
			switch curr_key {
			case ("-i"):
				parser.input_filename = arr[i+1]
			case ("-o"):
				parser.output_filename = arr[i+1]
			case ("-s"):
				value, err := Atoi(arr[i+1])
				if err != nil {
					panic("Bad num")
				}
				parser.stats = value
			case ("-w"):
				value, err := ParseInt(arr[i+1], 10, 64)
				if err != nil {
					panic("Bad num")
				}
				parser.window = value
			case ("-f"):
				value, err := ParseInt(arr[i+1], 10, 64)
				if err != nil {
					panic("Bad num")
				}
				parser.begin = value
			case ("-t"):
				value, err := ParseInt(arr[i+1], 10, 64)
				if err != nil {
					panic("Bad num")
				}
				parser.end = value
			default:
				panic("Bad t")
			}
			i++
		}
	}
}

func (parser Parser) IsEmpty() bool {
	return parser.input_filename == "" || parser.output_filename == ""
}
