package sql

import (
	"bufio"
	"dsc/fancy_errors"
	"strings"
	"time"
	"unicode"
)

/*
Query

The query object contains the metadata used for executing and logging a query ran via DSC.
*/
type Query struct {
	Content   string
	LineStart int
	LineEnd   int
	StartTime time.Time
	EndTime   time.Time
}

/*
Parse

This function parses text input and splits them into 'Query' objects containing the query's content,
and what line it starts and ends on. This is used for processing and logging by the sql
execution functions.

Potential cases for parsing query 'N':

1. Single line with 1 query:
'''
N;
'''

2. Single line with multiple queries. (N1; N2; N3; N4):
'''
N1; N2; N3; N4;
'''

3. Single line with multiple queries AND a multi line query(s):
N1; N2; N3; N4a
N4b
N4c;

4. Multi line query:
Na
Nb
Nc
Nd;

5. Multi line query AND single line query(s):
N1a
N1b
N1c
N1d;N2;N3;
*/
func Parse(input string) ([]Query, error) {

	scanner := bufio.NewScanner(strings.NewReader(input))

	lineCount := 0

	var queryList []Query

	query := &Query{}

	for scanner.Scan() {

		lineCount++

		lineString := scanner.Text()

		lineString = tabToSpace(lineString)

		if strings.Contains(lineString, ";") {

			// account for case 2, 3, & 5
			if strings.Count(lineString, ";") > 1 {
				lineQueryList := strings.SplitAfter(lineString, ";")

				for i, lineQuery := range lineQueryList {
					if i == 0 {
						query.Content += lineQuery + "\n"
						query.LineEnd = lineCount

						queryList = append(queryList, *query)

						query = &Query{}
					} else {

						if strings.Contains(lineQuery, ";") {
							query.Content += lineQuery + "\n"
							query.LineEnd = lineCount

							queryList = append(queryList, *query)

							query = &Query{}
						} else {
							query.Content += lineQuery + "\n"
						}
					}
				}

			} else {
				query.Content += lineString + "\n"
				query.LineEnd = lineCount

				queryList = append(queryList, *query)

				query = &Query{}
			}

		} else if strings.ReplaceAll(lineString, " ", "") != "\n" {

			// account for case 1, 3, 4, &5
			if query.LineStart == 0 {
				query.LineStart = lineCount
			}

			query.Content += lineString + "\n"

		}
	}

	if len(queryList) == 0 {
		return nil, fancy_errors.New("no valid queries present in input")
	}

	return queryList, nil
}

/*
tabToSpace

This function standardizes the input for easier parsing by stripping tabs and replacing them
with spaces.
*/
func tabToSpace(input string) string {
	var result []string

	for _, i := range input {
		switch {
		// all these considered as space, including tab \t
		// '\t', '\n', '\v', '\f', '\r',' ', 0x85, 0xA0
		case unicode.IsSpace(i):
			result = append(result, " ") // replace tab with space
		case !unicode.IsSpace(i):
			result = append(result, string(i))
		}
	}
	return strings.Join(result, "")
}
