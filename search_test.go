package bot

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	uid = ""
	sid = ""
	key = ""
)

func TestSearch(t *testing.T) {
	ctx := context.Background()
	data, err := SearchUser(ctx, "7000", uid, sid, key)
	assert.Nil(t, err)
	log.Println(string(data))

	data, err = ReadUser(ctx, "773e5e77-4107-45c2-b648-8fc722ed77f5", uid, sid, key)
	assert.Nil(t, err)
	log.Println(string(data))
}
