package rand

import (
	"testing"
)

func TestNormIntRange(t *testing.T) {
	rnd := New()
	for i := 0; i < 100; i++ {
		v := rnd.NormIntRange(100, 50)
		println(v)
	}
}
