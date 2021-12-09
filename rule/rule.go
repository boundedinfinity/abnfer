package rule

import (
	"fmt"

	"github.com/boundedinfinity/abnfer/char"
)

type Rule struct {
	Name    string
	Min     int
	Max     int
	Pattern Pattern
}

type Pattern [][]char.Char

type Rules map[string]Rule

func (t Rules) Append(r Rule) error {
	if _, ok := t[r.Name]; ok {
		return fmt.Errorf("alread exists")
	}

	t[r.Name] = r
	return nil
}
