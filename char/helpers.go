package char

import "bytes"

func HasPrefix(s string, cs []Char) bool {
	var bs []byte
	for _, c := range cs {
		bs = append(bs, byte(c))
	}
	return bytes.HasPrefix([]byte(s), bs)
}

func Bytes(cs ...Char) []byte {
	var bs []byte

	for _, c := range cs {
		bs = append(bs, byte(c))
	}

	return bs
}

func String(cs ...Char) string {
	bs := Bytes(cs...)
	return string(bs)
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

func AnyOf(cs ...Char) []Char {
	return append([]Char{}, cs...)
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
