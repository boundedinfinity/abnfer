package char

import "bytes"

func HasPrefix(s string, cs []Char) bool {
	return bytes.HasPrefix([]byte(s), []byte(string(cs)))
}

func Range(s, e Char) []Char {
	var l []Char

	for c := s; c <= e; c++ {
		l = append(l, c)
	}

	return l
}

func FilteredRange(s, e Char, fs []Char) []Char {
	var l []Char

	for c := s; c <= e; c++ {
		for _, f := range fs {
			if c == f {
				continue
			}
		}

		l = append(l, c)
	}

	return l
}

func FromString(s string) []Char {
	var l []Char

	for _, c := range s {
		l = append(l, Char(c))
	}

	return l
}
