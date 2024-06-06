package avatar

import (
	"testing"
)

func TestMap(t *testing.T) {
	hex1 := colorHash("gmelum")
	hex2 := colorHash("gmelum")

	if hex1 != hex2 {
		t.Error("invalid color hash")
	}
}
