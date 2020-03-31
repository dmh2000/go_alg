package main

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// activate benchmarking if required
	testing.Init()

	// parse command line flags for test
	flag.Parse()

	// run and exit
	os.Exit(m.Run())
}
