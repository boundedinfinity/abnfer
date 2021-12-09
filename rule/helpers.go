package rule

import "github.com/boundedinfinity/abnfer/char"

func AnyOf(cs ...[]char.Char) Pattern {
	var p Pattern

	p = append(p, cs...)

	return p
}

func Concat(rs ...Rule) Pattern {
	var p Pattern

	for _, r := range rs {
		for _, rp := range r.Pattern {
			p = append(p, rp)
		}
	}

	return p
}
