package main

import (
	"context"
	"log"

	"github.com/MooooonStar/bot-api-go-client"
)

const (
	uid = ""
	sid = ""
	key = ``
)

func main() {
	ctx := context.Background()
	id := "8f873526-c131-4e84-985f-524f95319add"
	someone := "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"

	req := bot.ParticipantRequest{UserId: someone}
	_, err := bot.ManageParticipants(ctx, id, "ADD", []*bot.ParticipantRequest{&req}, uid, sid, key)
	if err != nil {
		log.Fatal(err)
	}
	_, err = bot.ManageParticipants(ctx, id, "REMOVE", []*bot.ParticipantRequest{&req}, uid, sid, key)
	if err != nil {
		log.Fatal(err)
	}
}
