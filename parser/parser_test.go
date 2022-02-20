package parser

import (
	"testing"
)

const success = "\u2713"
const failure = "\u2717"

func TestParse(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{
			// nested segments
			in:  "3{a2{bb}}",
			out: "abbbbabbbbabbbb",
		},
		{
			// serial segments
			in:  "2{a}3{b}4{c}",
			out: "aabbbcccc",
		},
		{
			// nested and serial segments
			in:  "2{a3{bc}}4{de}",
			out: "abcbcbcabcbcbcdededede",
		},
		{
			// empty segment
			in:  "2{3{ab}}",
			out: "abababababab",
		},
		{
			// empty deepest segment
			in:  "2{a3{}}",
			out: "aa",
		},
		{
			// no nested segments
			in:  "abc",
			out: "abc",
		},
		{
			// empty string
			in:  "",
			out: "",
		},
	}

	t.Log("Parse test")
	{
		for _, tc := range testCases {
			t.Logf("\tTest 0:\tWhen string '%s' is parced", tc.in)
			{
				out := Parse(tc.in)
				if out != tc.out {
					t.Fatalf("\t%s\tShould return '%s', got '%s'", failure, tc.out, out)
				}
				t.Logf("\t%s\tShould return '%s'.", success, tc.out)
			}
		}
	}
}
