package bot

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	var (
		uid = "e89acc53-6a5b-407f-8558-d574e82e5922"
		sid = "de3432cb-38ef-4418-80ad-512693aebd55"
		key = "-----BEGIN RSA PRIVATE KEY-----\r\nMIICXQIBAAKBgQCF4Ks6Vv7J99J3SOsAKmL3A+zwdFkQbZmNTiz3ucXhcyooMY+l\r\nmo7ks9ke421nXdPI6fJIykXCX8HffpsahhjWgBrD+SmfV8rjJpOlX7xydhr4h+7C\r\nJPxVnYGGW/elanzBjZba7p6zO4b3wlU/BEo60xiMrR9KGRvoKMu4ZawKewIDAQAB\r\nAoGACXEU4mJsSbYuo7cCy9oxbJZcSYELHvt7ztlqlnYFnKbQnFESfm1uEncUovcL\r\nKjUV0J+fNL7/OHwkYiy9p5s5eRvSndp1vOTdEenATXFs9wUg68mTlDYmAnzrXHWU\r\nO2BBf4NE4Zk6jIUFTg70KU0ReskgRcZAQAGjzS6sArvVd5ECQQD+Xj5AO1vPvFFg\r\nl4/FujnzAO1FTMKuxpglE4FxNnRR/cUbJHJeS8Whddo7FnWjIjwkTjM7wU7yt30T\r\nsy2j44l/AkEAhryKULbpjJnFIJYbC2tSNOXt4KHmHAb1P9US7u0Bb4Q0kThs4Lz0\r\nDPMlfV/IiTagDDm9HGbUKdY/vksWkCUlBQJBAJ6zvBElN8YnT3p9dVU/OFKV3HOl\r\nTb3v6BMR1WwGNpLFSwoBNl4A1oBDYHIHNEQg4vcf1zVMCW7D7oSxASPEng8CQGXz\r\nAaGjVcLKk4tdqScR1Mkr9buUJV6tsSLUohfhg8WSvofnrxK6RtwJmyNrt3yDtcSq\r\nZ7q23/CsKd1eSbtWL9UCQQDbF3h7A7lsMhvSOq5e0gGmJiKxLDeLbWnI6Yzac3oc\r\nhkaW2wzxgMRgSWH5yzvCbeiagOE/rm/mrPm019CcFUgC\r\n-----END RSA PRIVATE KEY-----\r\n"
	)

	ctx := context.Background()
	data, err := SearchUser(ctx, "7000000016", uid, sid, key)
	assert.Nil(t, err)
	log.Println(string(data))

	// data, err = ReadUser(ctx, "773e5e77-4107-45c2-b648-8fc722ed77f5", uid, sid, key)
	// assert.Nil(t, err)
	// log.Println(string(data))
}
