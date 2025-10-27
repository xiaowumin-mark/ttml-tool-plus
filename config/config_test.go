package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	u, err := os.UserHomeDir()
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}

func TestInit(t *testing.T) {
	ConfigInit()
}
