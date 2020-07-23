package gameUtils

import (
	"math/rand"
)

type UniqueRand struct {
	Generated map[int]bool
	MaxNumber int
}

func (u *UniqueRand) Int() int {
	for {
		i := rand.Intn(u.MaxNumber)
		if !u.Generated[i] {
			u.Generated[i] = true
			//log.Println("UID ", i)
			return i
		}
	}
}


func CalculateTotalPossibleUIDS(uniqueRand UniqueRand) int64{
	return int64(uniqueRand.MaxNumber - len(uniqueRand.Generated) + 1)}
