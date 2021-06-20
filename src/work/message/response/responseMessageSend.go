package response

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseMessageSend struct {
	*response.ResponseWX

	InvalidUser  string `json:"invaliduser"`  // "userid1|userid2",
	InvalidParty string `json:"invalidparty"` // "partyid1|partyid2",
	InvalidTag   string `json:"invalidtag"`   // "tagid1|tagid2"

}