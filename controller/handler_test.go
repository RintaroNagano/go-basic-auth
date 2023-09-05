package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sample/db"
	"sample/model"
	"sample/model/user"
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

func TestSignupHandler(t *testing.T) {
	db.GormConnect()
	defer db.GetDB().Close()

	// Migrate the schema
	db.GetDB().AutoMigrate(&user.User{})

	gin.SetMode(gin.TestMode)

	// Handlerをセットアップ
	r := gin.Default()
	r.POST("/signup", SignupHandler)

	// テストデータ
	reqBody := map[string]string{
		"user_id":  "testuser",
		"password": "testpass",
	}
	jsonValue, _ := json.Marshal(reqBody)

	// リクエストを作成
	req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("リクエストの作成に失敗: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// レスポンスを記録
	w := httptest.NewRecorder()

	// リクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスを検証（ここではステータスコードとレスポンスボディの一部だけ検証していますが、
	// 本当はモックを使用してデータベースの変更も検証するべきです。）
	assert.Equal(t, http.StatusOK, w.Code)

	resp := model.SignupResponse{}
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Fatalf("レスポンスのパースに失敗: %v\n", err)
	}
	assert.Equal(t, "Account successfully created", resp.Message)
	assert.Equal(t, "testuser", resp.User.UserID)
	assert.Equal(t, "testuser", resp.User.Nickname)
}
