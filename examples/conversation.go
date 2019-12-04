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
	uid = "fce4a2bf-f239-44d0-a4ce-b7fd3ff02327"
	sid = "db50b3d0-2303-4a3c-9ab3-b568556f9b63"
	key = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAqkQUTXCkCOEpUXB2H9wqL+D+S7eLxbglYW3NyAm1317PrNpZ
njni7tP1BvrfSbQYJH9MJlVsFtF6zwiWLR4dgPHeZkcu+Tj5b1qtG7Wao0Z9NeQg
/QnCxA/n0rVaHnO3XKNeKb2a3X0TAX3XkPwPDGGpSnZfeliD/Z3Dtl3zgnYbCD07
QVl3efRLGtwPYG3I1pc/FDOMcrbAPnE4Dg7fe5UE+O/Ks1mwpyodel+Tp+Nj/dZP
ENVdYvjXTGK7XUaD4GyyldpobR6WJ1mAbi2FRe1N+BSQq8bH8MVqFBxvYjiE1AvI
n6D29M3fH6sWi+IkPn/BigPDiSEab9qsDpOqCwIDAQABAoIBAANQb0NmoDrFWeQ0
7TJc5bv9hCr9zUxcSOtRk4IuS4pSh8up4GQHumwYV2+lrcz7IsaAbSoIhaHCvWzk
POVvQvcXPiioKrjn1qsdN7tkreLkZLGKLujd/3AzjblN6w18BniuJroR777EnXxs
8lI8ex/W6CWuki3eeos9k+lWgNPsPRcJvS74vfANzj2NMQpuZqLozBBdW10obMG/
Vs5+XOmxW7kFDkmrCOTrV2VUxwB2Jdx+isJoCKdO9XqUKqpwJB7FKwvtBpMvNDeE
5VdqLTGkBxzTk5ebgMOxz0CaouKGiJGC37psvwtWwbiWSZ0hgnUFPRs/cxYhqGG9
5eDGqtECgYEA6n4d46hKyogqmKA6YyaFXeOZZfWOxGmMs94AsvWAUItmHqIeRqWZ
Q4+Zf/sK4tQH7m0ccFeoihDwU7LotrB3dQbLMx7SHK3yUMCNZkpVfcgF/EOI1VKM
1VcRS2gh/5znH/7knRR1YQ3Ic9LOZK3y7JldmKshkoCHPwkpD68kfakCgYEAueHr
p/6dthU50WD93c2E7yJ3MpcZw9odJ7ORhVSGpnkxW7pEG+O+tTZJ2JAWFC7NR656
vSdBe8Fr037sn77ooghl7tK6dy9lptp/8FYcRZr1Jg2kmVR1NQ1ygiAxqLszNMtd
X1NrghXApBclKqipqg9Jd6PW0cuDE8ddd/tAspMCgYBHWWp0y88+AuKbIB3D74Oj
Bs61FFuFbexPClzCFysc1t2SspdtvZN92CLHqzYrwIJNTMvbbwCsSe0mOojOlbEz
XYbUnismaNMTORONLU8H/6W6lT38UqrheyWbAStfIHm+5YrgNMpmZpA4lcsRHVCO
JzgwbNnBPvRn/LyOtibvSQKBgBFGoQtVC71fjacbF50ku8Up8+Ac0QU14V3OP6kf
qQAR5Xd8AVaCPThEWmqpQZQQoX2XBstLubdRy761XjKeoKpmk2q15SSsoAr45ckl
h60kcMJfpDun8NjEbi8P5O2ksGaou/3fF+/YK7ZUfFzRniszDoXzEUUgYAdXkSVU
kbfHAoGAMsVPZhb9JPpM54QVw7MaB2NC1hGULz22IjUf1MO+c4Fhg1UC7wSmid6B
jJ495odSDwF0j64Ks3e/+rQR8FxRBp0h0Hd2WYvslLgahcNq182swRntsuEwoBcp
eRPtBuVbwq6VphDHvVHuD85IOanZhiHH0qXlWrIUXBpGZGkZjow=
-----END RSA PRIVATE KEY-----`
)

// const (
// 	cat  = "14521f6b-2619-41ba-89ff-d440330cbde0"
// 	snow = "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"
// 	uid  = "fce4a2bf-f239-44d0-a4ce-b7fd3ff02327"
// 	sid  = "36250dc4-f987-4eae-a986-760f7d8e57dc"
// 	key  = `-----BEGIN RSA PRIVATE KEY-----
// MIICXAIBAAKBgQDAgWryZroTizbqJcZ3P07/ho0pW5FbnsGn9tFDdFnZZFbs0VLy
// hHAkfPdrYYHXEGDYEnNg4sZ8bd7weNwFlLhsGkV5Dn6b0XkW11X8Unp240uVf3fm
// TpX899yFoGcfaXB0LoxdMSn8jABFDtDkFaTdQ4bmilrem7jVR57VRiFaxQIDAQAB
// AoGABHVK0oPOBeNYuXn3+22AKRtG3CEWf/JwhjEJXiBM72OrPq9HVj36ceHiAf7f
// WKl0sLRWrzT5R0rtGZEW+VsdHTt/sNIq/vIdsM1qyJfKaHIk+P/7rB6sxUHhZmjm
// Rr8RlGWD7fg/yPerzvhF+c5kT3ihJUbETg9V/tIjINYjRgECQQDp1TwhljhShZzh
// RYfSONKTwKTLmGwd10Q4OFtfLdkKU2C2i064vDgKUhsy35NdfTZEU4FNLtqh+aFR
// hDiNPajFAkEA0sE696Y1WHidEZXUa4J/fU70FWC0Ml2pwOwDEpjYK9309ytE1iqv
// OCKw81rY/Ec3hAJezVRbaFfoaqn8m9wKAQJAYtqM0z4ojWqDChFU/CCdlW0pvhCw
// FGV/J61bo39EdEMPbdyy6RJZNrPDX9x49lsK6BPNZA3czAE2m7mkFsHsaQJBAKWo
// 0Wr8MI5cMEo0VYY0S3X5644qNzonuVk1qqyhi7nfe8AVOMSHRLeBPUnsLQcQC6ku
// G8qV9kusXboJ58+zRAECQFJrp9GDK1JtjxnpbtjwGEThygbVWfd8guuuKz2rDDdP
// XPFRmZlYqdGSGvi87iEKqaite3rHZKyMZWKFxZZLaNc=
// -----END RSA PRIVATE KEY-----`
// )

func main() {
	ctx := context.Background()

	// auth, err := bot.ReadAuthorization(ctx, "2c2753af-4c2e-41d9-8c5f-9f1915c951c9", uid, sid, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(auth))
	// ctx := context.Background()
	// user, err := bot.SearchUser(ctx, "7000102353", uid, sid, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(user))
	// id := "8f873526-c131-4e84-985f-524f95319add"
	// someone := "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"

	// req := bot.ParticipantRequest{UserId: someone}
	// _, err := bot.ManageParticipants(ctx, id, "ADD", []*bot.ParticipantRequest{&req}, uid, sid, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = bot.ManageParticipants(ctx, id, "REMOVE", []*bot.ParticipantRequest{&req}, uid, sid, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bot.ConversationShow(ctx)
	// accessToken, err := bot.SignAuthenticationToken(uid, sid, key, "GET", "/assets", "")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// assets, err := bot.AssetList(ctx, accessToken)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(assets)

	// authorizations, err := bot.ReadAuthorizations(ctx, uid, sid, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(string(authorizations))
	// attachment, err := bot.AttachemntShow(ctx, uid, sid, key, "0839a7c6-2b7a-47ec-85d0-cc80c0164421")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bt, _ := json.Marshal(attachment)
	// log.Println(string(bt))

	// ids := []string{"e89acc53-6a5b-407f-8558-d574e82e5922", "2b9a8b76-80a6-4cd3-a586-1c33ff2cbd11", "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909"}
	// data, err := bot.FetchUsers(ctx, ids, uid, sid, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(":", string(data))

	// id := "29890426-64b6-4588-aac6-432ce6b1e52c"
	// accessToken, err := bot.SignAuthenticationToken(uid, sid, key, "GET", "/conversations/"+id, "")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// conversation, err := bot.ConversationShow(ctx, id, accessToken)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// bt, _ := json.Marshal(conversation)
	// log.Println(string(bt))

	// ctx := context.Background()
	accessToken, err := bot.SignAuthenticationToken(uid, sid, key, "GET", "/me", "")
	if err != nil {
		log.Fatal(err)
	}
	assets, err := bot.Request(ctx, "GET", "/me", nil, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(assets))
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

type Listener struct{}

func (l *Listener) OnMessage(ctx context.Context, msg bot.MessageView, userId string) error {
	bt, _ := json.Marshal(msg)
	log.Println("message: ", string(bt))
	return nil
}
