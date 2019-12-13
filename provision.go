package bot

import (
	"context"
	"encoding/json"
	"time"
)

func GetProvisioningId(ctx context.Context, device, uid, sid, key string) (id string, err error) {
	method, uri := "POST", "/provisionings"
	params := map[string]interface{}{
		"device": device,
	}
	body, _ := json.Marshal(params)
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, string(body))
	if err != nil {
		return
	}
	data, err := Request(ctx, method, uri, body, accessToken)
	if err != nil {
		return
	}
	var resp struct {
		Data struct {
			DeviceID  string    `json:"device_id"`
			ExpiredAt time.Time `json:"expired_at"`
		} `json:"data"`
		Error Error
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		return
	} else if resp.Error.Code != 0 {
		return "", resp.Error
	}
	return resp.Data.DeviceID, nil
}

func GetProvisioning(ctx context.Context, id, uid, sid, key string) ([]byte, error) {
	method, uri := "GET", "/provisionings/"+id
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}

//
func VerifyProvisioning(ctx context.Context, id, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/provisionings/verify"
	params := map[string]interface{}{
		"app_version":      "0.5.1",
		"code":             "7539",
		"platform":         "Desktop",
		"platform_version": "OS X 10.14.6",
		"purpose":          "SESSION",
		"registration_id":  4224452628,
		"session_id":       "9b65b685-bd49-4256-9ec5-5b7a82990276",
		"session_secret":   "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAmv9Z4QjYU3H4WMuxfnUP\nOcokKGkDlYtlA8S72DHskdpr7qWIHEAR20eNEMwQW2qY2IbnRP2PTzqFCU+K7w2u\nZBsixHjapyTmoe7XS6URI1rQ9QrZtbc3VgX/rQhT5fve6XEHQiY8d0inbeXVmDwI\nHy0+XqhfjPWvqOC1gnUc+Y2V8bfjij+REEJYGz7n4IJeCKNQPqydg51OLlECl3Y/\nNXr663wi+IUMgwzlB9BAG863ZZbwr00LXyQuQX1SrXfrYlODs9R/7Jw2rn2NnOwy\ndSmTk/EVQIYS4eGP5EqnRHW1UgVlzlgswF2kWyXqCBTdCulP4a0d3GCe3RDAfSAs\nqwIDAQAB",
		"user_id":          "7b3f0a95-3ee9-4c1b-8ae9-170e3877d909",
	}
	body, _ := json.Marshal(params)
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, string(body))
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, body, accessToken)
}
