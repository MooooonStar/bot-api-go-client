package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"time"

	"github.com/MooooonStar/bot-api-go-client"
	"github.com/gofrs/uuid"
)

const (
	snow = "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"
	uid  = "e89acc53-6a5b-407f-8558-d574e82e5922"
	sid  = "981b9559-dbcd-4606-beb8-74e70168daf3"
	key  = "-----BEGIN RSA PRIVATE KEY-----\r\nMIICXAIBAAKBgQCLNAuQpGpyAWmZhh4Rj2WIrkdUhSFUzXLb9laBQO5lN1bWA7de\r\nib8h9NeoNIN09CkOjuoX37qHKJ4eu74gAg2qPxWBBfZqfbV+F1fWfNZhjK0KZs+4\r\nwod7hiCZBHagSeb4xk2IKuoCq6T9gNUnSDcyZRlryXh5wpmKW0+tLue5YwIDAQAB\r\nAoGAQTOiiyCH/1hsBw4L/XumOkwvewZUo20L9L3ArRAl2eD+2oWpGS36CcSuUsou\r\nHNTUek40lfQkYXSEA686kbGEhlxiToQImBmXURweocRj/avL4U9T+914vLxyD24H\r\nhCmoy8C8WZThUVvxdpdhittgI4frOaHoSdxtBxtMfTByhHkCQQDXOFlc9q/F+gEc\r\nJtyH5r4fw+Dj82CfTnC9g+DOIjTjXgKqbM3OuNMWA/Y2fv0etJpL/CcOl/yONZPy\r\nVXD7rFiXAkEApZRcDgx2eTC+wpnhgUSmyb4Lr+Z7zVp/r7H6a1yQ0ZlNse5V5C2Z\r\nDY256sWyOztVTlKnYtYu7GVViXJjQUDTFQJAEYebb8EADCYS1RItQDyBG+E+fhKi\r\nuhcPg8AsHV4X5LURjoQSN1ZCizaDZWjugQPA1ZOSMtQaVBL1G/Xtn/InvQJBAJiR\r\nan0baLLkgnFHXy1Ad7qenaA37M1cXRkkEo2InnwzGQ8hgA0fIs8FT1EPVtoTIthn\r\neoqstlHmUXmzZ9eBhr0CQC/qADQI/tPJrMh9q5GZorEKO69kGrljgdBOPPJbg8ye\r\ny6bq+WKGaH0sRmVzFzYTSe0O+M3hXD3qYZIAe9/V+sI=\r\n-----END RSA PRIVATE KEY-----\r\n"
)

type Listener struct{}

func (l *Listener) OnMessage(ctx context.Context, msgView bot.MessageView, userID string) error {
	bt, _ := json.Marshal(msgView)
	log.Println("On message ", string(bt))
	return nil
}

func main() {
	ctx := context.Background()
	l := &Listener{}
	id := bot.UniqueConversationId(snow, uid)

	var client *bot.BlazeClient
	go func() {
		for {
			client = bot.NewBlazeClient(uid, sid, key)
			if err := client.Loop(ctx, l); err != nil {
				log.Println("LOOP ERROR: ", err)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	log.Println("I am here.")
	btns, _ := json.Marshal(map[string]interface{}{
		"title":       "hello",
		"action":      "www.google.cn",
		"description": "",
		"icon_url":    "https://gss0.baidu.com/-4o3dSag_xI4khGko9WTAnF6hhy/zhidao/pic/item/5243fbf2b2119313ac02f75c69380cd790238d9c.jpg",
	})
	params := map[string]interface{}{
		"conversation_id": id,
		"recipient_id":    snow,
		"message_id":      bot.UuidNewV4().String(),
		"category":        "APP_CARD",
		"data":            base64.StdEncoding.EncodeToString(btns),
	}
	for {
		err := client.SendRawMessage(ctx, params)
		log.Println("Send raw message error: ", err)
		time.Sleep(2 * time.Second)
	}
}

func talk() {
	ctx := context.Background()

	l := &Listener{}
	var client *bot.BlazeClient
	go func() {
		for {
			client = bot.NewBlazeClient(uid, sid, key)
			if err := client.Loop(ctx, l); err != nil {
				log.Println("ERROR: ", err)
			}
		}
	}()

	time.Sleep(5 * time.Second)

	cat := "14521f6b-2619-41ba-89ff-d440330cbde0"
	for {
		id := bot.UniqueConversationId(uid, cat)
		data := base64.StdEncoding.EncodeToString([]byte("hello"))
		// err := bot.PostMessage(ctx, id, cat, uuid.Must(uuid.NewV4()).String(), "PLAIN_TEXT", data, uid, sid, key)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		err := client.SendMessage(ctx, id, cat, uuid.Must(uuid.NewV4()).String(), "PLAIN_TEXT", data, "")
		time.Sleep(1 * time.Second)
		log.Println("I am here.", err)
	}
}
