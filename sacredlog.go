package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type LogLine struct {
	Time string        `json:"time,omitempty"`
	Corr string        `json:"corr,omitempty"`
	Level string       `json:"lvl,omitempty"`
	ShortFile string   `json:"shortfile,omitempty"`
	Msg string         `json:"msg,omitempty"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	l := LogLine{}
	line := 0
	for scanner.Scan() {
		line = line + 1
		err := json.Unmarshal(scanner.Bytes(), &l)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing line %d : %s", line, err)
			continue
		}
		fmt.Printf("%s %s %s %q\n", l.Time, l.Level, l.ShortFile, l.Msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}