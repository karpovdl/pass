package main

import (
	afero "github.com/spf13/afero"
)

const (
	// Endl line break
	Endl = "\r\n"
)

var (
	appFs = afero.NewOsFs()

	isPprof bool
)
