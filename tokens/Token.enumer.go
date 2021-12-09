//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package tokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Token string
type Tokens []Token

func Slice(es ...Token) Tokens {
	var s Tokens

	for _, e := range es {
		s = append(s, e)
	}

	return s
}

const (
	Alternative            Token = "alternative"
	Comment                Token = "comment"
	FileEnd                Token = "file-end"
	FileStart              Token = "file-start"
	IncrementalAlternative Token = "incremental-alternative"
	RuleEnd                Token = "rule-end"
	RuleName               Token = "rule-name"
	RuleStart              Token = "rule-start"
	SequenceGroupBegin     Token = "sequence-group-begin"
	SequenceGroupEnd       Token = "sequence-group-end"
	String                 Token = "string"
	SyntaxError            Token = "syntax-error"
	Terminal               Token = "terminal"
	ValueRangeAlternative  Token = "value-range-alternative"
)

var (
	All = Tokens{
		Alternative,
		Comment,
		FileEnd,
		FileStart,
		IncrementalAlternative,
		RuleEnd,
		RuleName,
		RuleStart,
		SequenceGroupBegin,
		SequenceGroupEnd,
		String,
		SyntaxError,
		Terminal,
		ValueRangeAlternative,
	}
)

func Is(v string) bool {
	return All.Is(v)
}

func Parse(v string) (Token, error) {
	return All.Parse(v)
}

func Strings() []string {
	return All.Strings()
}

func (t Token) String() string {
	return string(t)
}

var ErrTokenInvalid = errors.New("invalid enumeration type")

func Error(vs Tokens, v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrTokenInvalid, v, strings.Join(vs.Strings(), ","),
	)
}

func (t Tokens) Strings() []string {
	var ss []string

	for _, v := range t {
		ss = append(ss, v.String())
	}

	return ss
}

func (t Tokens) Parse(v string) (Token, error) {
	var o Token
	var f bool
	n := strings.ToLower(v)

	for _, e := range t {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, Error(t, v)
	}

	return o, nil
}

func (t Tokens) Is(v string) bool {
	var f bool

	for _, e := range t {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func (t Tokens) Contains(v Token) bool {
	for _, e := range t {
		if e == v {
			return true
		}
	}

	return false
}

func (t Token) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *Token) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := Parse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}