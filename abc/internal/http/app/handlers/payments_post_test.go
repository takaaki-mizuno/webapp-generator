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

func TestHandler_PaymentsPost(t *testing.T) {
	t.Run("API to make a payment for a specified amount. if the payment is successful, the API returns the receipt information and the Payment ID, along with the status code 200. If the payment is unsuccessful, it returns the reason for the failure, with status code 400.", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodPost, "/payments", strings.NewReader(""))
		// Set the request header for testing
		

		handler := createHandler()
		handler.PaymentsPost(c)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		response, err := ioutil.ReadAll(w.Result().Body)
		var result map[string]interface{}
		err = json.Unmarshal(response, &result)
		assert.Nil(t, err)
	})
}
