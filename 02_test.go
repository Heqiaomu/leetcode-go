package leetcode_go

import (
	"fmt"
	"testing"
)

func Test_getListNode(t *testing.T) {

	t.Run("", func(t *testing.T) {
		node := getListNode([]int{1, 2, 3, 4})
		fmt.Println(node)
	})
}

func Test_addTwoNumbers(t *testing.T) {

	t.Run("", func(t *testing.T) {
		numbers := addTwoNumbers(getListNode([]int{9, 9, 9, 9, 9, 9, 9}), getListNode([]int{9, 9, 9, 9}))
		fmt.Println(numbers.String())
	})

}
