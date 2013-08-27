package pan

import (
	"testing"
)

type queryTest struct {
	ExpectedResult string
	Query *Query
}

var queryTests = []queryTest{
	queryTest{
		ExpectedResult: "This query expects $1 one arg;",
		Query: &Query{
			SQL: "This query expects ? one arg",
			Args: []interface{}{0},
		},
	},
	queryTest{
		ExpectedResult: "",
		Query: &Query{
			SQL: "This query expects ? one arg but won't get it;",
			Args: []interface{}{},
		},
	},
	queryTest{
		ExpectedResult: "",
		Query: &Query{
			SQL: "This query expects no arguments but will get one;",
			Args: []interface{}{0},
		},
	},
	queryTest{
		ExpectedResult: "",
		Query: &Query{
			SQL: "This query expects ? two args ? but will get one;",
			Args: []interface{}{0},
		},
	},
	queryTest{
		ExpectedResult: "",
		Query: &Query{
			SQL: "This query expects ? ? two args but will get three;",
			Args: []interface{}{0, 1, 2},
		},
	},
	queryTest{
		ExpectedResult: "Unicode test 世 $1;",
		Query: &Query{
			SQL: "Unicode test 世 ?",
			Args: []interface{}{0},
		},
	},
	queryTest{
		ExpectedResult: "Unicode boundary test $1 "+string(rune(0x80))+";",
		Query: &Query{
			SQL: "Unicode boundary test ? "+string(rune(0x80)),
			Args: []interface{}{0},
		},
	},
}

func TestQueriesFromTable(t *testing.T) {
	for pos, test := range queryTests {
		result := test.Query.String()
		if result != test.ExpectedResult {
			t.Logf("Expected\n%v\ngot\n%v\n.", []byte(test.ExpectedResult), []byte(result))
			t.Errorf("Query test %d failed. Expected \"%s\", got \"%s\".", pos+1, test.ExpectedResult, result)
		}
	}
}

func BenchmarkQueriesFromTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		test := queryTests[b.N % len(queryTests)]
		b.StartTimer()
		result := test.Query.String()
		b.StopTimer()
		if result != test.ExpectedResult {
			b.Errorf("Query test %d failed. Expected \"%s\", got \"%s\".", (b.N % len(queryTests)) + 1, test.ExpectedResult, result)
		}
		b.StartTimer()
	}
}
