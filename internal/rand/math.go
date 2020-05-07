package rand

import "math/rand"

var _ Rand = mathRand{}

type mathRand struct{}

func (r mathRand) Int(max int) (int64, error) {
	return int64(rand.Intn(max)), nil
}
