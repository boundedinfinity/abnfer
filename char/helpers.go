package char

import "bytes"

type Pattern []byte

func HasPrefix(s string, patterns ...Pattern) bool {
	x := []byte(s)

	for _, pattern := range patterns {
		if bytes.HasPrefix(x, pattern) {
			return true
		}
	}

	return false
}

func OneOf(s string) bool {
	_, ok := MAP[s]
	return ok
}

func Concat(ps ...Pattern) Pattern {
	var n Pattern

	for _, p := range ps {
		n = append(n, p...)
	}

	return n
}

func Range(s, e int) Pattern {
	var p Pattern

	for i := s; i <= e; i++ {
		p = append(p, byte(i))
	}

	return p
}

func List(is ...int) Pattern {
	var p Pattern

	for _, i := range is {
		p = append(p, byte(i))
	}

	return p
}
