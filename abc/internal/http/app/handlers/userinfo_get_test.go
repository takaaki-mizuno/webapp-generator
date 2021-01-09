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

func TestHandler_UserinfoGet(t *testing.T) {
	t.Run("Get user information by MiniApp Token. MiniApp Backend/Fronend will access to this API to get user information to use for the service.", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodGet, "/userinfo", strings.NewReader(""))
		// Set the request header for testing
		

		handler := createHandler()
		handler.UserinfoGet(c)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		response, err := ioutil.ReadAll(w.Result().Body)
		var result map[string]interface{}
		err = json.Unmarshal(response, &result)
		assert.Nil(t, err)
	})
}
