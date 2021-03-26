package bamboo

import (
	"fmt"
	"net/http"
)

type Encryption service

type EncryptionResut struct {
	EncryptedText string `json:"encryptedText"`
}

func (s *Encryption) Encrypt(text string) (encryptionResult EncryptionResut, r *http.Response, err error) {
	request, err := s.client.NewRequest(http.MethodPost, "encrypt", `{text: "`+text+`"}`)
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
