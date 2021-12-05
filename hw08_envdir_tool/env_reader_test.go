package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	// Place your code here
	t.Run("invalid dir path", func(t *testing.T) {
		_, err := ReadDir("invalid/path")
		require.Error(t, err)
	})

	t.Run("empty folder", func(t *testing.T) {
		dir, err := ioutil.TempDir("testdata", "test")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir)

		env, err := ReadDir(dir)
		require.NoError(t, err)
		require.Equal(t, 0, len(env))
	})
}
