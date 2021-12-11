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

var _ = Describe("Pattern", func() {
	Context("char range", func() {
		var p lang.Pattern
		p.CharRange(char.UPPERCASE_A, char.UPPERCASE_C)

		Expect(len(p)).To(Equal(3))
		Expect(len(p[0])).To(Equal(1))
		Expect(p[0][0]).To(Equal(char.UPPERCASE_A))
		Expect(len(p[1])).To(Equal(1))
		Expect(p[1][0]).To(Equal(char.UPPERCASE_B))
		Expect(len(p[2])).To(Equal(1))
		Expect(p[2][0]).To(Equal(char.UPPERCASE_C))
	})

	Context("append", func() {
		var p lang.Pattern
		p.CharAnyOf(char.CARRIAGE_RETURN, char.LINE_FEED)
		p.CharAnyOf(char.LINE_FEED)

		Expect(len(p)).To(Equal(2))
		Expect(len(p[0])).To(Equal(2))
		Expect(p[0][0]).To(Equal(char.CARRIAGE_RETURN))
		Expect(p[0][1]).To(Equal(char.LINE_FEED))
		Expect(len(p[1])).To(Equal(1))
		Expect(p[1][0]).To(Equal(char.LINE_FEED))
	})

	Context("combine", func() {
		var p1 lang.Pattern
		var p2 lang.Pattern
		var p3 lang.Pattern

		p1.CharAnyOf(char.CARRIAGE_RETURN, char.LINE_FEED)
		p2.CharAnyOf(char.LINE_FEED)
		p3.AppendPattern(p1, p2)

		Expect(len(p3)).To(Equal(2))
		Expect(len(p3[0])).To(Equal(2))
		Expect(p3[0][0]).To(Equal(char.CARRIAGE_RETURN))
		Expect(p3[0][1]).To(Equal(char.LINE_FEED))
		Expect(len(p3[1])).To(Equal(1))
		Expect(p3[1][0]).To(Equal(char.LINE_FEED))
	})

	Context("WSP", func() {
		Expect(len(lang.WSP)).To(Equal(2))
		Expect(len(lang.WSP[0])).To(Equal(1))
		Expect(len(lang.WSP[1])).To(Equal(1))
	})

	Context("LWSP", func() {
		Expect(len(lang.LWSP)).To(Equal(2))
		Expect(len(lang.LWSP[0])).To(Equal(2))
		Expect(len(lang.LWSP[1])).To(Equal(4))
	})
})
