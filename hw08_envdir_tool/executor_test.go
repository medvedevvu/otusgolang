package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("empty command", func(t *testing.T) {
		code := RunCmd([]string{}, Environment{})
		require.Equal(t, 1, code)
	})

	t.Run("invalid command", func(t *testing.T) {
		code := RunCmd([]string{"zzz"}, Environment{})
		require.NotEqual(t, 0, code)
	})

	t.Run("success simple", func(t *testing.T) {
		code := RunCmd([]string{"ls", "-g"}, Environment{})
		require.Equal(t, 0, code)
	})

	t.Run("success simple with env", func(t *testing.T) {
		code := RunCmd([]string{"ls", "-g"}, Environment{"TEST_ENV": EnvValue{Value: "TEST_ENV_VALUE"}})
		require.Equal(t, 0, code)
		require.Contains(t, os.Environ(), "TEST_ENV=TEST_ENV_VALUE")
	})
}
