package mascot_test

import (
	"testing"

	"github.com/IvashkaKid/GoPractice/mascot"
)

func MascotTest(t *testing.T) {
	if mascot.BestMascot() != "Gp Gopher" {
		t.Fatal("It's not best mascot :(")
	}
}
