package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// gin用のRecorder, Contextを作成
func SetupGinContext() (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return w, c
}

// gin用のRecorder, Contextを作成(リクエストパラメータ)
func SetupGinContextWithParam(params []gin.Param) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// リクエストパラメータを追加する
	for _, p := range params {
		c.Params = append(
			c.Params, p,
		)
	}
	return w, c
}

// gin用のRecorder, Contextを作成(リクエストボディ)
func SetupGinContextWithBody(model any) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// リクエストボディを追加する
	jsonValue, _ := json.Marshal(model)
	reqBody := bytes.NewBuffer(jsonValue)
	req, _ := http.NewRequest(
		"POST",
		"",
		reqBody,
	)
	c.Request = req
	return w, c
}

// gin用のRecorder, Contextを作成(リクエストクエリ)
func SetupGinContextWithQuery(queryList []gin.Param) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// リクエストクエリを設定する
	req, _ := http.NewRequest(
		http.MethodPost,
		"",
		nil,
	)
	q := req.URL.Query()
	for _, query := range queryList {
		q.Add(query.Key, query.Value)
	}
	req.URL.RawQuery = q.Encode()
	c.Request = req
	return w, c
}
