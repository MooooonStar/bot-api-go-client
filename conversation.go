package bot

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type Participant struct {
	UserId    string    `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Conversation struct {
	ConversationId string    `json:"conversation_id"`
	CreatorId      string    `json:"creator_id"`
	Category       string    `json:"category"`
	Name           string    `json:"name"`
	IconURL        string    `json:"icon_url"`
	Announcement   string    `json:"announcement"`
	CreatedAt      time.Time `json:"created_at"`

	Participants []Participant `json:"participants"`
}

func CreateConversation(ctx context.Context, name, category, conversationId string, participants []Participant, uid, sid, key string) (*Conversation, error) {
	params, err := json.Marshal(map[string]interface{}{
		"name":            name,
		"category":        category,
		"conversation_id": conversationId,
		"participants":    participants,
	})
	if err != nil {
		return nil, err
	}
	accessToken, err := SignAuthenticationToken(uid, sid, key, "POST", "/conversations", string(params))
	if err != nil {
		return nil, err
	}
	body, err := Request(ctx, "POST", "/conversations", params, accessToken)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Data  Conversation `json:"data"`
		Error Error        `json:"error"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Error.Code > 0 {
		if resp.Error.Code == 401 {
			return nil, AuthorizationError(ctx)
		} else if resp.Error.Code == 403 {
			return nil, ForbiddenError(ctx)
		}
		return nil, ServerError(ctx, resp.Error)
	}
	return &resp.Data, nil
}

func ConversationShow(ctx context.Context, conversationId string, accessToken string) (*Conversation, error) {
	body, err := Request(ctx, "GET", "/conversations/"+conversationId, nil, accessToken)
	if err != nil {
		return nil, err
	}
	var resp struct {
		Data  Conversation `json:"data"`
		Error Error        `json:"error"`
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	if resp.Error.Code > 0 {
		if resp.Error.Code == 401 {
			return nil, AuthorizationError(ctx)
		} else if resp.Error.Code == 403 {
			return nil, ForbiddenError(ctx)
		}
		return nil, ServerError(ctx, resp.Error)
	}
	return &resp.Data, nil
}

type ParticipantAction string

const (
	ParticipantActionAdd    ParticipantAction = "ADD"
	ParticipantActionRemove                   = "REMOVE"
	ParticipantActionJoin                     = "JOIN"
	ParticipantActionExit                     = "EXIT"
)

type ConversationRequest struct {
	ConversationId string         `json:"conversation_id"`
	Category       string         `json:"category"`
	Name           string         `json:"name"`
	IconBase64     string         `json:"icon_base64"`
	Announcement   string         `json:"announcement"`
	Duration       string         `json:"duration"`
	Participants   []*Participant `json:"participants"`
}

type ParticipantRequest struct {
	UserId    string    `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func UpdateConversation(ctx context.Context, id string, req *ConversationRequest, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/conversations/"+id
	params, _ := json.Marshal(req)
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, string(params))
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, params, accessToken)
}

func ReadConversation(ctx context.Context, id string, uid, sid, key string) ([]byte, error) {
	method, uri := "GET", "/conversations/"+id
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}

func ManageParticipants(ctx context.Context, id, action string, req []*ParticipantRequest, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", fmt.Sprintf("/conversations/%s/participants/%s", id, action)
	params, _ := json.Marshal(req)
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, string(params))
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, params, accessToken)
}

func ExitConversation(ctx context.Context, id string, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/conversations/"+id+"/exit"
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}

func JoinConversation(ctx context.Context, codeId string, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/conversations/"+codeId+"/join"
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}

func RorateConversationQrcode(ctx context.Context, id string, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/conversations/"+id+"/rotate"
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}

func MuteConversation(ctx context.Context, id string, uid, sid, key string) ([]byte, error) {
	method, uri := "POST", "/conversations/"+id+"/mute"
	accessToken, err := SignAuthenticationToken(uid, sid, key, method, uri, "")
	if err != nil {
		return nil, err
	}
	return Request(ctx, method, uri, nil, accessToken)
}
