package env

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	env := defaultEnv
	LoadEnv(env)
	fmt.Println(env)
	assert.Equal(t, "prod", env.Mode)
	assert.Equal(t, "/Users/letmehues/projects/gobase/gobase-core/", env.BasePath)
	assert.Equal(t, "/Users/letmehues/projects/gobase/gobase-core/conf", env.ConfPath)
}
