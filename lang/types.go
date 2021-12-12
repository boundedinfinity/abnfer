package lang

import (
	"strings"

	"github.com/boundedinfinity/abnfer/char"
)

type Rule interface {
	Matches(string) MatchResult
	New(int, int) Rule
}

type Repetition struct {
	Name string
	Min  int
	Max  int
}

type MatchResult struct {
	Count   int
	Matched bool
}

type TerminalRule struct {
	Repetition
	Terminals []char.Char
}

func NewTerminal(name string, min, max int, cs []char.Char) Rule {
	return TerminalRule{
		Repetition: Repetition{
			Name: name,
			Min:  min,
			Max:  max,
		},
		Terminals: cs,
	}
}

func (r TerminalRule) New(min, max int) Rule {
	return NewTerminal(r.Name, min, max, r.Terminals)
}

func (r TerminalRule) Matches(s string) MatchResult {
	var result MatchResult

	if len(s) < r.Min {
		return result
	}

	b := make([]bool, 0)

	for i := 0; i < len(s); i++ {
		var found bool

		for _, v := range r.Terminals {
			if strings.HasPrefix(s[i:], v.String()) {
				b = append(b, true)
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	bl := len(b)

	if bl >= r.Min && bl <= r.Max {
		result.Matched = true
		result.Count = bl
	}

	return result
}

type ConcatenationRule struct {
	Repetition
	Rules []Rule
}

func NewConcatenation(name string, min, max int, rs []Rule) Rule {
	return ConcatenationRule{
		Repetition: Repetition{
			Name: name,
			Min:  min,
			Max:  max,
		},
		Rules: rs,
	}
}

func (r ConcatenationRule) New(min, max int) Rule {
	return NewConcatenation(r.Name, min, max, r.Rules)
}

func (r ConcatenationRule) Matches(s string) MatchResult {
	var result MatchResult
	var mr1 []MatchResult
	s2 := s

	for {
		var mr2 []MatchResult

		for _, v := range r.Rules {
			mr := v.Matches(s2)

			if !mr.Matched {
				break
			} else {
				mr2 = append(mr2, mr)
				s2 = s2[mr.Count:]
			}
		}

		if len(mr2) != len(r.Rules) {
			break
		} else {
			mr1 = append(mr1, mr2...)
		}

		if s2 == "" {
			break
		}
	}

	dividend := len(mr1) / len(r.Rules)
	remainder := len(mr1) % len(r.Rules)

	if remainder == 0 && dividend >= r.Min && dividend <= r.Max {
		result.Matched = true

		for _, v := range mr1 {
			result.Count += v.Count
		}
	}

	return result
}

type AlternativesRule struct {
	Repetition
	Rules []Rule
}

func NewAlternatives(name string, min, max int, rs []Rule) Rule {
	return AlternativesRule{
		Repetition: Repetition{
			Name: name,
			Min:  min,
			Max:  max,
		},
		Rules: rs,
	}
}

func (r AlternativesRule) New(min, max int) Rule {
	return NewAlternatives(r.Name, min, max, r.Rules)
}

func (r AlternativesRule) Matches(s string) MatchResult {
	for _, v := range r.Rules {
		if result := v.Matches(s); result.Matched {
			return result
		}
	}

	return MatchResult{}
}
