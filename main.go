//go:build !go1.12
// +build !go1.12

package main

import (
	"github.com/bfirestone/bodyclose/passes/bodyclose"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(bodyclose.Analyzer) }
