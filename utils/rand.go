package utils

import (
	"math/rand"
	"time"
	"errors"
)

type Rand struct {
	rand.Rand
}

func GenerateRand(seed *int) (*Rand, error) {
	_seed := time.Now().UnixMilli()
	if seed != nil {
		_seed = int64(*seed)
	}
	source := rand.NewSource(int64(_seed))
	newRand := rand.New(source)
	if newRand == nil {
		return nil, errors.New("error creating new Rand")
	}
	r := Rand{*newRand}
	return &r, nil
}

func (r *Rand) GetInt(i int) int {
	return r.Intn(i)
}
