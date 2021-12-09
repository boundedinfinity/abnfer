package rule_test

import (
	"testing"

	"github.com/boundedinfinity/abnfer/char"
	"github.com/boundedinfinity/abnfer/rule"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRegex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rule Helpers Suite")
}

var _ = Describe("AnyOf", func() {
	Context("correctly assign 2 singles", func() {
		expected := rule.Pattern{
			[]char.Char{char.ZERO},
			[]char.Char{char.ONE},
		}
		actual := rule.AnyOf(char.Single(char.ZERO), char.Single(char.ONE))

		Expect(len(actual)).To(Equal(len(expected)))
	})

	Context("correctly assign a single and string", func() {
		expected := rule.Pattern{
			[]char.Char{char.CARRIAGE_RETURN, char.LINE_FEED},
			[]char.Char{char.LINE_FEED},
		}
		actual := rule.AnyOf(
			char.Concat(char.Single(char.CARRIAGE_RETURN), char.Single(char.LINE_FEED)),
			char.Single(char.LINE_FEED),
		)

		Expect(len(actual)).To(Equal(len(expected)))
	})
})
