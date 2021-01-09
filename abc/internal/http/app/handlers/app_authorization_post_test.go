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

func TestHandler_AppAuthorizationPost(t *testing.T) {
	t.Run("Authorize the mini-app. If the authorization is successful, API returns status code 200. If the mini app can be used in multiple countries, all countries can be approved at once, but if it is not available in the country specified in the header, an error occurs and API returns status code 403.", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodPost, "/apps/{bundle_id}/authorization", strings.NewReader(""))
		// Set the request header for testing
		

		handler := createHandler()
		handler.AppAuthorizationPost(c)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		response, err := ioutil.ReadAll(w.Result().Body)
		var result map[string]interface{}
		err = json.Unmarshal(response, &result)
		assert.Nil(t, err)
	})
}
