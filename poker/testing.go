package poker

import (
	"github.com/hewiefreeman/GopherGameServer/core"
	"log"
)

func openTestTable(){
	log.Println("Opening Test Table")
	room, roomErr := core.NewRoom("super", "PokerTable", false, 0, "")
	if roomErr != nil {
		log.Println("Error while opening Room:", roomErr)
	} else {
		log.Println(room)
	}
}
func openSecondTestTable(){
	log.Println("Opening Test Table")
	room, roomErr := core.NewRoom("cool", "PokerTable", false, 0, "")
	if roomErr != nil {
		log.Println("Error while opening Room:", roomErr)
	} else {
		log.Println(room)
	}
}

