package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var Version string
var Revision = "HEAD"

var embedVersion = "0.4.0"

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	if Version == "" {
		Version = embedVersion
	}
	fmt.Printf("p2 version %s (rev:%s)\n", Version, Revision)
	app.Exit(0)

	return nil
}
