package frame

import (
	"log"
	"testing"
)

func TestNewFrame(t *testing.T) {
	f := NewFrame(getDef())
	log.Printf("%#v", f)
}
