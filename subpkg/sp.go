package subpkg

import (
	"fmt"
)

type ReleaseType string

const (
	nightly ReleaseType = "nightly"
	EAP     ReleaseType = "EAP"
	Beta    ReleaseType = "Beta"
	Stable  ReleaseType = "Stable"
)

func MoveMe() {
	fmt.Printf("Download GoLand 2018.2 %s today!", Beta)
}