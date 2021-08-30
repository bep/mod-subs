package p

import (
	"fmt"
	"runtime/debug"
	"strings"

	mb "github.com/bep/mod-b/p"
)

func Version() string {
	return "self: " + getVersion() + "\nmod-b: " + mb.Version() + "\n"
}

func getVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if ok {
		version := bi.Main.Version
		version = strings.Trim(version, "()")
		return fmt.Sprintf("%s@%s\n", bi.Main.Path, version)
	}
	return "unknown"
}
