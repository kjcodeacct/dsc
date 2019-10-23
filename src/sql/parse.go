package sql

import (
	"bufio"
	"fmt"
	"strings"
	"time"
)

type QueryGroup struct {
	Priority   int
	CommitHash string
	List       []Query
}

type Query struct {
	Priority  int
	Filename  string
	LineStart int
	LineEnd   int
	StartTime time.Time
	EndTime   time.Time
	Content   string
}

/*
Cases:

insert into....; select from ...;



*/

func Parse(input string) (QueryGroup, error) {
	var newQueryGroup QueryGroup

	scanner := bufio.NewScanner(strings.NewReader(input))

	lineCount := 0
	var newQuery Query
	for scanner.Scan() {
		lineCount++
		lineStr := scanner.Text()

		if strings.Contains(lineStr, ";") {
			// detect multiple splits
			newQuery.LineEnd = lineCount
		}
	}

	if lineCount == 1 && len(newQueryGroup.List) > 1 {
		fmt.Println("Why did you single line this query?\nWe're compressing it before storage man, it'll be ok...")
	}
}
