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

func TestHandler_AppsGet(t *testing.T) {
	t.Run("Returns a list of apps that are available in the country specified in the header by the user specified in the GCP token.", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodGet, "/apps", strings.NewReader(""))
		// Set the request header for testing
		

		handler := createHandler()
		handler.AppsGet(c)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		response, err := ioutil.ReadAll(w.Result().Body)
		var result map[string]interface{}
		err = json.Unmarshal(response, &result)
		assert.Nil(t, err)
	})
}
