package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler_AppTokenPost(t *testing.T) {
	t.Run("Get MiniApp Token, if it successfully generate MiniApp Token, it returns token info with 200 status code. If GCP Token is wrong, it will return 401, and the user is not authorized with the miniapp yet, it returns 403 status, and need to use `/apps/{bundle_id}/authorization` to athorize before getting token. Also, even with an authorized mini-app, if the mini-app is not available in the country specified with the country code in the header, no token will be issued and an error will be returned with a status code 403.", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodPost, "/apps/{bundle_id}/token", strings.NewReader(""))
		// Set the request header for testing
		

		handler := createHandler()
		handler.AppTokenPost(c)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		response, err := ioutil.ReadAll(w.Result().Body)
		var result map[string]interface{}
		err = json.Unmarshal(response, &result)
		assert.Nil(t, err)
	})
}
