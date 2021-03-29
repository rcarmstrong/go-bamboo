package bamboo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Encryption service

type EncryptionResult struct {
	EncryptedText string `json:"encryptedText"`
}

type EncryptionRequest struct {
	Text string `text:"text"`
}

func (s *Encryption) Encrypt(requestParams EncryptionRequest) (encryptionResult EncryptionResult, r *http.Response, err error) {
	data, _ := json.Marshal(requestParams)
	request, err := s.client.NewRequest(http.MethodPost, "encrypt", data)
	if err != nil {
		return
	}

	response, err := s.client.Do(request, &encryptionResult)
	if err != nil {
		return
	}

	if !(response.StatusCode == 200) {
		return encryptionResult, response, &simpleError{fmt.Sprintf("Server pause returned %d", response.StatusCode)}
	}

	return
}
