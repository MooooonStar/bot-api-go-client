package bot

import (
	"context"
	"encoding/json"
)

func SearchUser(ctx context.Context, q string, uid, sid, key string) ([]byte, error) {
	method, uri := "GET", "/search/"+q
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}

func ReadUser(ctx context.Context, id string, uid, sid, key string) ([]byte, error) {
	method, uri := "GET", "/users/"+id
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}

func FetchUsers(ctx context.Context, ids []string, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/users/fetch"
	body, _ := json.Marshal(ids)
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, string(body))
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, body, accessToken)
}
