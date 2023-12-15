package _1

import "testing"

func Test_simplifyPath(t *testing.T) {
	simplifyPath("/a/./b/../../c/")
}
