package creativeword

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// DeleteRequest 删除动态创意词包 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativeWordID 创意词包id
	CreativeWordID uint64 `json:"creative_word_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// DeleteResponse 删除动态创意词包 API Response
type DeleteResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// CreativeWordID 创意词包ID
		CreativeWordID uint64 `json:"creative_word_id,omitempty"`
	} `json:"data,omitempty"`
}
