package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	// Ginをテストモードに設定
	gin.SetMode(gin.TestMode)

	// ルートの設定
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// ハンドラで使用するリクエストを作成。
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("リクエストの作成に失敗: %v\n", err)
	}

	// http.ResponseWriterを満たすResponseRecorderを作成して、レスポンスを記録。
	w := httptest.NewRecorder()

	// ルーターでリクエストを実行
	r.ServeHTTP(w, req)

	// ステータスコードが期待するものであるかチェック。
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスボディが期待するものであるかチェック。
	expected := "pong"
	assert.Equal(t, expected, w.Body.String())
}
