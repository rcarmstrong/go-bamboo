package bamboo_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	bamboo "github.com/rcarmstrong/go-bamboo"
)

var (
	testComment = "hello world"
)

func TestAddComment(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(addCommentStub))
	defer ts.Close()

	client := bamboo.NewSimpleClient(nil, "", "")
	client.SetURL(ts.URL)

	success, resp, err := client.Comments.AddComment("TEST-TEST-1", testComment)
	if err != nil {
		t.Error(err)
	}

	if success == false || resp.StatusCode != 200 {
		t.Error(fmt.Sprintf("Adding comment \"%s\" was unsuccessful. Returned %s", testComment, resp.Status))
	}
}

func addCommentStub(w http.ResponseWriter, r *http.Request) {
	comment := &bamboo.Comment{}

	json.NewDecoder(r.Body).Decode(comment)

	if comment.Content != testComment {
		http.Error(w, "comments do not match", http.StatusBadRequest)
		return
	}
}
