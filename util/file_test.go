package util

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCurrentAbsPath(t *testing.T) {
	fmt.Println(CurrentAbsPath())
	assert.Equal(t, "/Users/letmehues/projects/gobase/gobase-core", CurrentAbsPath())
}
