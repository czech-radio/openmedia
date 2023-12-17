package cmd

import (
	"fmt"
	"github/czech-radio/openmedia-archive/internal"
	"testing"
)

func TestArchive(t *testing.T) {
	// acfg := new(config_archive)
	acfg := Config_archive{}
	// internal.DeclareFlags(acfg)
	internal.ParseFlags(acfg)
	fmt.Printf("%+v", acfg)
}
