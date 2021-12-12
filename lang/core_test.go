package lang_test

import (
	"testing"

	"github.com/boundedinfinity/abnfer/char"
	"github.com/boundedinfinity/abnfer/lang"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRegex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lang Suite")
}

var _ = Describe("Rule", func() {
	Context("CRLF", func() {
		It("should match \\r\\n", func() {
			expected := lang.MatchResult{2, true}
			actual := lang.CRLF.Matches("\r\n")
			Expect(actual).To(Equal(expected))
		})

		It("should match \\r\\n", func() {
			expected := lang.MatchResult{4, true}
			actual := lang.CRLF.New(1, 2).Matches("\r\n\r\n")
			Expect(actual).To(Equal(expected))
		})

		It("should not match 'BELL'", func() {
			expected := lang.MatchResult{}
			actual := lang.CRLF.Matches("\b")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("ALPHA", func() {
		It("should match A", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.ALPHA.Matches("A")
			Expect(actual).To(Equal(expected))
		})

		It("should match a", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.ALPHA.Matches("a")
			Expect(actual).To(Equal(expected))
		})

		It("should not match !", func() {
			expected := lang.MatchResult{}
			actual := lang.ALPHA.Matches("!")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := lang.MatchResult{}
			actual := lang.ALPHA.New(2, 2).Matches("A")
			Expect(actual).To(Equal(expected))
		})

		It("should match AA", func() {
			expected := lang.MatchResult{2, true}
			actual := lang.ALPHA.New(2, 2).Matches("AA")
			Expect(actual).To(Equal(expected))
		})

		It("should match AA1", func() {
			expected := lang.MatchResult{2, true}
			actual := lang.ALPHA.New(2, 2).Matches("AA1")
			Expect(actual).To(Equal(expected))
		})

		It("should not match AAA1", func() {
			expected := lang.MatchResult{}
			actual := lang.ALPHA.New(2, 2).Matches("AAA1")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("BIT", func() {
		It("should match 0", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.BIT.Matches("0")
			Expect(actual).To(Equal(expected))
		})

		It("should not match !", func() {
			expected := lang.MatchResult{}
			actual := lang.BIT.Matches("!")
			Expect(actual).To(Equal(expected))
		})

		It("should match 01", func() {
			expected := lang.MatchResult{2, true}
			actual := lang.BIT.New(1, 2).Matches("01")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("CHAR", func() {
		It("should match 0", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.CHAR.Matches("0")
			Expect(actual).To(Equal(expected))
		})

		It("should match 'BELL'", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.CHAR.Matches("\b")
			Expect(actual).To(Equal(expected))
		})

		It("should not match 'NULL'", func() {
			expected := lang.MatchResult{}
			actual := lang.CHAR.Matches(string(byte(0x00)))
			Expect(actual).To(Equal(expected))
		})
	})

	Context("CR", func() {
		It("should match \\r", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.CR.Matches("\r")
			Expect(actual).To(Equal(expected))
		})

		It("should not match 'BELL'", func() {
			expected := lang.MatchResult{}
			actual := lang.CR.Matches("\b")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("LF", func() {
		It("should match \\n", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.LF.Matches("\n")
			Expect(actual).To(Equal(expected))
		})

		It("should not match 'BELL'", func() {
			expected := lang.MatchResult{}
			actual := lang.LF.Matches("\b")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("CTL", func() {
		It("should match 'BELL'", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.CTL.Matches("\b")
			Expect(actual).To(Equal(expected))
		})

		It("should match 'DELETE'", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.CTL.Matches(char.DELETE.String())
			Expect(actual).To(Equal(expected))
		})

		It("should not match ' '", func() {
			expected := lang.MatchResult{}
			actual := lang.CTL.Matches(" ")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("DIGIT", func() {
		It("should match 1", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.DIGIT.Matches("1")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := lang.MatchResult{}
			actual := lang.CTL.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("DQUOTE", func() {
		It("should match \"", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.DQUOTE.Matches(`"`)
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := lang.MatchResult{}
			actual := lang.DQUOTE.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("HEXDIG", func() {
		It("should match 1", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.HEXDIG.Matches(`1`)
			Expect(actual).To(Equal(expected))
		})

		It("should match F", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.HEXDIG.Matches(`F`)
			Expect(actual).To(Equal(expected))
		})

		It("should not match G", func() {
			expected := lang.MatchResult{}
			actual := lang.HEXDIG.Matches("G")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("HTAB", func() {
		It("should match \\t", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.HTAB.Matches("\t")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := lang.MatchResult{}
			actual := lang.HTAB.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("SP", func() {
		It("should match 'SPACE'", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.SP.Matches(" ")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := lang.MatchResult{}
			actual := lang.SP.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("WSP", func() {
		It("should match 'SPACE'", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.WSP.Matches(" ")
			Expect(actual).To(Equal(expected))
		})

		It("should match \\t", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.HTAB.Matches("\t")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := lang.MatchResult{}
			actual := lang.WSP.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	Context("LWSP", func() {
		It("should match 'SPACE'", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.LWSP.Matches(" ")
			Expect(actual).To(Equal(expected))
		})

		It("should match \\t", func() {
			expected := lang.MatchResult{1, true}
			actual := lang.LWSP.Matches("\t")
			Expect(actual).To(Equal(expected))
		})

		It("should match \\r\\n\\t", func() {
			expected := lang.MatchResult{3, true}
			actual := lang.LWSP.Matches("\r\n\t")
			Expect(actual).To(Equal(expected))
		})

		It("should not match A", func() {
			expected := lang.MatchResult{}
			actual := lang.LWSP.Matches("A")
			Expect(actual).To(Equal(expected))
		})
	})

	// Context("OCTET", func() {
	// 	It("should match 'SPACE'", func() {
	// 		expected := lang.MatchResult{1, true}
	// 		actual := lang.OCTET.Matches(" ")
	// 		Expect(actual).To(Equal(expected))
	// 	})

	// 	It("should match \\t", func() {
	// 		expected := lang.MatchResult{1, true}
	// 		actual := lang.OCTET.Matches("\t")
	// 		Expect(actual).To(Equal(expected))
	// 	})
	// })
})
