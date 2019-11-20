package bot

import (
	"context"
	"encoding/json"
)

func ReadAuthorization(ctx context.Context, id string, uid, sid, key string) ([]byte, error) {
	method, uri := "GET", "/authorizations/"+id
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}

func CancelAuthorization(ctx context.Context, clientId string, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/oauth/cancel"
	params, _ := json.Marshal(map[string]interface{}{
		"client_id": clientId,
	})
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, string(params))
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, params, accessToken)
}

func Authorize(ctx context.Context, id string, scopes []string, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/oauth/authorize"
	params, _ := json.Marshal(map[string]interface{}{
		"authorization_id": id,
		"scopes":           scopes,
	})
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, string(params))
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, params, accessToken)
}
