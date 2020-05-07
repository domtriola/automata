package rand

import (
	"crypto/rand"
	"math/big"
)

var _ Rand = cryptoRand{}

type cryptoRand struct{}

func (r cryptoRand) Int(max int) (int64, error) {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}

	return num.Int64(), nil
}
