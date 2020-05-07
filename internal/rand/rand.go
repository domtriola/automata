package rand

import (
	"fmt"
)

// Rand is a random number generator that supports different underlying rand
// implementations.
type Rand interface {
	Int(max int) (int64, error)
}

// NewRand initializes a new random number generator with the specified rand
// implementation.
func NewRand(randType string) (Rand, error) {
	switch randType {
	case "crypto":
		return cryptoRand{}, nil
	case "math":
		return mathRand{}, nil
	default:
		return nil, fmt.Errorf("%s randType not recognized", randType)
	}
}
