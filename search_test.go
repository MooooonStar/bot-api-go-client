package bot

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	uid = "e89acc53-6a5b-407f-8558-d574e82e5922"
	sid = "1cff82d1-5873-488b-b210-b25421db7ccb"
	key = "-----BEGIN RSA PRIVATE KEY-----\r\nMIICXQIBAAKBgQC529vlrNBDpxAvBFRp/PyydQaA67P6jfK2CmaRcj2jAjIM34sT\r\nee0ufQFE6AjfuYGE1rzH82FzygUBFs9Ag2ZJrYjQyITi89aVDI2k9O29z+Xc3Ss7\r\nOK46Y/O7QRuJNFL9GL3nOE9sVEZMfZ9GojgmUEVvge6GTi8391iEclTf4wIDAQAB\r\nAoGBAJRYVOK0FaL59GPiq9HQ+I0j9OXvu76vj9sW4BkVOvch6Hr37hTEi7hAVEQA\r\nhBy1/xfdId7idpUp2OauowuWFYw6LtfG/CaFgVI0u0fip9z3LZ2gMpAAtJG29Hs2\r\nv7LFfHwRFjHy8guIWVrqev1i5YSaRe4Nn64IRzdBwhzfLMIpAkEA6YbDR1iQlowi\r\nhRFdBOSJJOr0aS55Y1Ja52wpofupQOP/lXYB+2EdUtHPKaju0GutXcLqcmhdQLFR\r\nImFoWPg43wJBAMu+va3xmeHOaRVk7wJYewOUAgFrANKezl1b1xS3xaeSCxOALl5I\r\nk3YBjsQeUV6q/r2I3Z1h+dDpp7HQlgbTRX0CQAodu2Hg1UqgQXKQQP5QRBykTPgt\r\nejayP5vc1KTeuQNFnTHFTeqs109aJWZBmi7PK1BvfCZ+nOOLit85ZwLLpfMCQCg0\r\nJzGfrwHDvTCfjhokY4G9PvgQqUTmWUpKyRIpgTMjulPKw1uPbqmmfQluAhhwHXxA\r\n+ZoIr0buFE0eH7yRC/ECQQDXKRJodBt/VIlB33N6w/qXtusb8zMKdWlNLdi1tiyd\r\nhZ63Dg/WtrbHKb3l41fno4k7OYhRJONn3i/ThiHnzQrk\r\n-----END RSA PRIVATE KEY-----\r\n"
)

func TestSearch(t *testing.T) {
	ctx := context.Background()
	data, err := SearchUser(ctx, "38451826", uid, sid, key)
	assert.Nil(t, err)
	log.Println(string(data))

	// data, err = ReadUser(ctx, "773e5e77-4107-45c2-b648-8fc722ed77f5", uid, sid, key)
	// assert.Nil(t, err)
	// log.Println(string(data))
}
