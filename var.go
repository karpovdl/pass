package main

import (
	afero "github.com/spf13/afero"
)

const (
	// Endl line break
	Endl = "\r\n"
)

// Resp ...
type Resp struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

var (
	appFs = afero.NewOsFs()

	port int // Server port

	isPprof   bool // Pprof flag
	pprofPort int  // Pprof server port
)
