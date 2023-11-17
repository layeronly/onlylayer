package main

import (
	_ "embed"

	"github.com/onlylayer/onlylayer/command/root"
	"github.com/onlylayer/onlylayer/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
