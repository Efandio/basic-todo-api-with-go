package helper

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-api/model"

	"github.com/stretchr/testify/assert"
)


func TestDecoderJson(t *testing.T) {

	// Sub test using string literal
	t.Run("Valid JSON input from string literal", func(t *testing.T) {
		mockJsonString := `{"id":"129jans","title":"BelajarTesting","complete":false}`
		reqBody := strings.NewReader(mockJsonString)

		req := httptest.NewRequest(http.MethodPost, "/", reqBody)
		req.Header.Set("Content-Type", "application/json")

		var todo model.Todo
		err := DecoderJson(req, &todo)

		assert.NotNil(t, todo, "todo should not nil")
		assert.Nil(t, err, "result err should not nil")
		assert.NoError(t, err, "Should not error")
	})
	
}

func TestEncoderJson(t *testing.T) {
	
	t.Run("Success with struct data", func(t *testing.T) {

		type MockData struct {
			ID	string	`json:"id"`
			Title	string	`json:"title"`
		}

		mockResData := MockData{ID: "kaoa22223qqqwas", Title: "Test Encode"}
		recorder := httptest.NewRecorder()

		err := EncoderJson(recorder, http.StatusOK, mockResData)
		assert.NoError(t, err, "expected to success")
	})

}