package chapter29

import (
	"encoding/json"
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	assert := assert2.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Holymoly", string(data))

}

func TestMarshal(t *testing.T) {
	assert := assert2.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	std := new(Student)
	err := json.NewDecoder(res.Body).Decode(std)
	assert.Nil(err)
	fmt.Println(std)
}
