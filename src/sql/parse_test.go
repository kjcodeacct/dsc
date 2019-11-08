package sql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const testOne = `
SELECT *
FROM test
WHERE test_a IS NOT NULL;
`

const testTwo = `
INSERT INTO test(a,b)VALUES(1,2);INSERT INTO test_other(b,c)VALUES(1,2);INSERT INTO test_other(b,c)
VALUES(1,2);

SELECT *
FROM test
WHERE text_d IS NOT NULL;
`

func TestParser(t *testing.T) {

	queryList, err := Parse(testOne)

	for _, query := range queryList {
		fmt.Println("Query:", query.Content)
		fmt.Println("LineStart:", query.LineStart)
		fmt.Println("LineEnd:", query.LineEnd)
		fmt.Println("")
	}

	require.Nil(t, err)

	queryList, err = Parse(testTwo)

	for _, query := range queryList {
		fmt.Println("Query:", query.Content)
		fmt.Println("----LineStart:", query.LineStart)
		fmt.Println("----LineEnd:", query.LineEnd)
		fmt.Println("")
	}

	require.Nil(t, err)
}
