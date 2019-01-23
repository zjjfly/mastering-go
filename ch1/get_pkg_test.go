package ch1

import (
	"github.com/mactsouk/go/simpleGitHub"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(simpleGitHub.AddTwo(1, 2))
}
