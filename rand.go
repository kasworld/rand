// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
