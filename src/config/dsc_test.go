package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWorkingDir(t *testing.T) {

	location, err := findWorkingDir()
	require.Nil(t, err)
}
