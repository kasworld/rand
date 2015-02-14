// extend stardard rand package
// can use instead of standard rand package
package rand

import (
	r "math/rand"
	"time"
)

type Rand struct {
	*r.Rand
}

func (rnd *Rand) IntRange(s, e int) int {
	return rnd.Intn(e-s) + s
}

func (rnd *Rand) NormIntRange(desiredMean, desiredStdDev int) int {
	return int(rnd.NormFloat64()*float64(desiredStdDev)) + desiredMean
}

func (rnd *Rand) NormFloat64Range(desiredMean, desiredStdDev float64) float64 {
	return rnd.NormFloat64()*desiredStdDev + desiredMean
}

func New() *Rand {
	return &Rand{
		r.New(r.NewSource(time.Now().UnixNano())),
	}
}
