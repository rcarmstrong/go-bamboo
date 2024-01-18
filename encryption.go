package bamboo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Encryption service

type IEncryption interface {
	Encrypt(requestParams EncryptionRequest) (encryptionResult EncryptionResult, err error)
}

type EncryptionResult struct {
	EncryptedText string `json:"encryptedText"`
}

type EncryptionRequest struct {
	Text string `json:"text"`
}

func (s *Encryption) Encrypt(requestParams EncryptionRequest) (encryptionResult EncryptionResult, err error) {
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
		return encryptionResult, &simpleError{fmt.Sprintf("Server pause returned %d", response.StatusCode)}
	}

	return
}
