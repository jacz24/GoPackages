package poker

import (
	"math/rand"
)

var UIDList map[int]bool

var UIDStruct = UniqueRand{}

var roomUIDList map[int]bool

var roomUID = UniqueRand{}

type UniqueRand struct {
	generated map[int]bool
	maxNumber int
}

func (u *UniqueRand) Int() int {
	for {
		i := rand.Intn(u.maxNumber)
		if !u.generated[i] {
			u.generated[i] = true
			//log.Println("UID ", i)
			return i
		}
	}
}


func calculateTotalPossibleUIDS(uniqueRand UniqueRand) int64{
	return int64(uniqueRand.maxNumber - len(uniqueRand.generated) + 1)}
