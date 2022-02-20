// String parser
// @fixme:
// (1) limitation for segment repetition is [1..9]
// (2) error handling - bad string format

package parser

import (
	"strconv"
	"strings"
)

func Parse(in string) string {
	out := ""
	index := 0

	for index < len(in) {
		segment, i := ParseSegment(in[index:], index)

		out = out + segment
		index = i
	}

	return out
}

func ParseSegment(str string, index int) (string, int) {
	segment := ""
	next := index

	for i, r := range str {
		if r == '}' {
			next := next + 2
			return segment, next
		}

		if '0' < r && r <= '9' {
			count, _ := strconv.Atoi(string(r))
			s, next := ParseSegment(str[i+2:], next+1)
			ss := strings.Repeat(s, count)
			return segment + ss, next
		}

		if r == '{' {
			s, next := ParseSegment(str[i+1:], next)
			return segment + s, next
		}

		segment = segment + string(r)
		next++
	}

	return segment, next
}
