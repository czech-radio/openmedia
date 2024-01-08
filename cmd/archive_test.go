package cmd

import (
	"fmt"
	"github/czech-radio/openmedia-archive/internal"
	"testing"
)

func TestArchive(t *testing.T) {
	acfg := &Config_archivate{}
	internal.SetupRootFlags(acfg)
	fmt.Printf("%+v", acfg)
}
