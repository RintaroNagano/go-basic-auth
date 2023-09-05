package myauth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"sample/db"
	"sample/model/user"
	"sample/myhash"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	// この関数はミドルウェアとして使用できるHandlerFuncを返す
	return func(c *gin.Context) {
		// リクエストからAuthorizationヘッダーを取得する
		auth := c.GetHeader("Authorization")

		// ヘッダーがない場合、401ステータスを返してリクエストを中止する
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication Failed"})
			return
		}

		// Authorizationヘッダーは"Basic "で始まることを期待する
		const basicSchema = "Basic "
		if !strings.HasPrefix(auth, basicSchema) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication Failed"})
			return
		}

		// "Basic "をトリムした後、Base64エンコードされた文字列をデコードする
		str, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, basicSchema))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication Failed"})
			return
		}

		// デコードされた文字列を":"で分割して、ユーザーIDとパスワードを取得する
		creds := strings.SplitN(string(str), ":", 2)
		if len(creds) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication Failed"})
			return
		}

		// 取得したユーザーIDとパスワードを使って認証を試みる
		userID, password := creds[0], creds[1]
		if !CheckUserCredentials(userID, password) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authentication Failed"})
			return
		}

		// Authorizationヘッダに他人のIDを入れて認証を突破されないように，
		// パスパラメータがあるときは，パスパラメータと比較
		if id := c.Param("id"); id != "" {
			if id != userID {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "No Permission for Update"})
				return
			}
		}

		// 認証が成功したら、リクエスト処理を次に進める
		c.Next()
	}
}

// 与えられたユーザー ID とパスワードが正しいかどうかをチェックする
func CheckUserCredentials(userID, password string) bool {
	var u user.User

	// SELECT * FROM users WHERE user_id = '(valuable userId)' ORDER BY id LIMIT 1;
	err := db.GetDB().Where("user_id = ?", userID).First(&u).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("Error: myauth.CheckUserCredential(), User not found")
		}
	}

	return u.Password == myhash.PasswordToHash(password)
}

// AuthorizationヘッダからユーザIDを取り出す
func ExtractUserIDFromAuthHeader(authString string) (string, error) {
	const basicSchema = "Basic "

	if !strings.HasPrefix(authString, basicSchema) {
		return "", errors.New("Authentication Failed")
	}

	str, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authString, basicSchema))
	if err != nil {
		return "", errors.New("Authentication Failed")
	}

	creds := strings.SplitN(string(str), ":", 2)
	if len(creds) != 2 {
		return "", errors.New("Authentication Failed")
	}

	return creds[0], nil
}
