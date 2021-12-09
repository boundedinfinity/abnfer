package char

import "bytes"

func HasPrefix(s string, cs []Char) bool {
	var bs []byte
	for _, c := range cs {
		bs = append(bs, byte(c))
	}
	return bytes.HasPrefix([]byte(s), bs)
}

func Range(s, e Char) []Char {
	var l []Char

	for c := s; c <= e; c++ {
		l = append(l, c)
	}

	return l
}

func Concat(cs ...[]Char) []Char {
	var o []Char

	for _, c := range cs {
		o = append(o, c...)
	}

	return o
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

func String(s string) []Char {
	var l []Char

	for _, c := range s {
		l = append(l, Char(c))
	}

	return l
}

func Single(c Char) []Char {
	return []Char{c}
}
