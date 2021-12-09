package lexer_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/boundedinfinity/abnfer/lexer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRegex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lexer Suite")
}

var (
	spec1 = `; quoted string of SP and VCHAR
;  without DQUOTE
	`
)

var _ = Describe("Lex", func() {
	Context("should match full year", func() {
		l := lexer.New("test", spec1)
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for token := range l.Events {
				fmt.Printf("token type: %v, Value: %v\n", token.Type, token.Value)
			}
		}()

		l.Run()

		wg.Wait()
	})
})
