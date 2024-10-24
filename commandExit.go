package main

import (
	"os"
)

func commandExit(cfg *Config, area string) error {
	os.Exit(0)
	return nil
}
