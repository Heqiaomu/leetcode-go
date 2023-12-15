package _8

import (
	"fmt"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	justify := fullJustify([]string{"Here", "is", "an", "example", "of", "text", "justification."}, 16)

	for _, value := range justify {

		fmt.Println("value -> ", value)
	}
}
