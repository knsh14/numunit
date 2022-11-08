package main

import (
	"numunit"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(numunit.Analyzer) }
