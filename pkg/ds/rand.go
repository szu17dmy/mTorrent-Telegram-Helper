package ds

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandInt(i int) int {
	return r.Intn(i)
}
