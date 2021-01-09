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

func TestHandler_PaymentVerifyGet(t *testing.T) {
	t.Run("This API is used to verify the receipt obtained at the time of payment, and MiniApp Backend can confirm that the receipt is correct by verifying it with this API. Also, since the values sent by MiniApp Frontend, such as the payment amount, are not reliable, MiniApp backend must use the values obtained through this API.", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = httptest.NewRequest(http.MethodGet, "/payments/{paymentId}/verify", strings.NewReader(""))
		// Set the request header for testing
		

		handler := createHandler()
		handler.PaymentVerifyGet(c)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)

		response, err := ioutil.ReadAll(w.Result().Body)
		var result map[string]interface{}
		err = json.Unmarshal(response, &result)
		assert.Nil(t, err)
	})
}
