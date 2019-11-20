package bot

import "context"

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
