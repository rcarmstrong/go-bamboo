package bamboo

import "encoding/base64"

// Authorizer is the interface that wraps the Authorization method
// Authorization returns the string to be used in the Authorization value of header
type Authorizer interface {
	Authorization() string
}

// SimpleCredentials are the username and password used to communicate with the API
type SimpleCredentials struct {
	Username string
	Password string
}

func (sc *SimpleCredentials) Authorization() string {
	return "Basic " + basicAuth(sc.Username, sc.Password)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// TokenCredentials are the token used to communicate with the API
// https://developer.atlassian.com/server/bamboo/using-the-bamboo-rest-apis/
type TokenCredentials struct {
	Token string
}

func (tc *TokenCredentials) Authorization() string {
	return "Bearer " + tc.Token
}
