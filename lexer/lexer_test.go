package lexer_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRegex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lexer Suite")
}

// var (
// 	spec1 = `
// 	                            ; quoted string of SP and VCHAR
//                                 ;  without DQUOTE

// 	`
// )

var _ = Describe("Lex", func() {
	Context("should match full year", func() {
		for i := 0; i <= 255; i++ {
			fmt.Printf("%d - %X - %s\n", i, i, string(byte(i)))
		}
	})
	// Context("should match full year", func() {
	// 	l := lexer.New("test", spec1)
	// 	wg := sync.WaitGroup{}

	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for token := range l.Ch {
	// 			fmt.Printf("token type: %v, Value: %v\n", token.Type, token.Value)
	// 		}
	// 	}()

	// 	l.Run()

	// 	wg.Wait()
	// })
})
