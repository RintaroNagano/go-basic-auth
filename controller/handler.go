package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"sample/db"
	"sample/model"
	"sample/model/user"
	"sample/myauth"
	"sample/myhash"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping",
	})
}

func SignupHandler(c *gin.Context) {
	req := struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}{
		UserID:   "",
		Password: "",
	}

	// JSONデータの受け取り
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Received data:", req.UserID, req.Password)
		fmt.Println(err)
		c.Abort()
	}

	u := user.User{
		UserID:   req.UserID,
		Password: req.Password,
		Nickname: req.UserID,
		Comment:  "",
	}

	// ユーザIDとパスワードの存在についてバリデーション
	if err := u.ValidateExistUserIDandPassword(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Account creation failed",
			"cause":   "requested user_id and password",
		})
		fmt.Println(err)
		c.Abort()
		return
	}

	// ユーザIDとパスワードの長さについてバリデーション
	if err := u.ValidateLength(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Account creation failed",
			"cause":   "length user_id and password",
		})
		fmt.Println(err)
		c.Abort()
		return
	}

	// パスワードのパターンについてバリデーション
	if err := u.ValidatePattern(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Account creation failed",
			"cause":   "pattern user_id and password",
		})
		fmt.Println(err)
		c.Abort()
		return
	}

	hashpass := myhash.PasswordToHash(u.Password)
	u.Password = hashpass

	// SELECT * FROM users WHERE user_id = '(valuable userId)' ORDER BY id LIMIT 1;
	err := db.GetDB().Where("user_id = ?", u.UserID).First(&u).Error

	if err != nil {
		// // INSERT INTO `users` (`userid`,`name`) VALUES ("user.UserID", "user.Name");
		db.GetDB().Create(u)

		// レスポンスの作成
		response := model.SignupResponse{
			Message: "Account successfully created",
			User: struct {
				UserID   string `json:"user_id"`
				Nickname string `json:"nickname"`
			}{
				UserID:   u.UserID,
				Nickname: u.UserID,
			},
		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Account creation failed",
			"cause":   "already same user_id is used",
		})
	}
}

func GetUserHandler(c *gin.Context) {
	id := c.Param("id")

	var u user.User

	// SELECT * FROM users WHERE user_id = '(valuable userId)' ORDER BY id LIMIT 1;
	if err := db.GetDB().Where("user_id = ?", id).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No User found",
		})
		c.Abort()
		return
	}

	response := struct {
		Message string
		User    struct {
			UserID   string `json:"user_id"`
			Nickname string `json:"nickname"`
			Comment  string `json:"comment"`
		}
	}{
		Message: "User details by user_id",
		User: struct {
			UserID   string `json:"user_id"`
			Nickname string `json:"nickname"`
			Comment  string `json:"comment"`
		}{
			UserID:   u.UserID,
			Nickname: u.Nickname,
			Comment:  u.Comment,
		},
	}

	c.JSON(http.StatusOK, response)
}

func PatchUserHandler(c *gin.Context) {
	// データベースに受け取ったidのレコードが存在するか調べる
	id := c.Param("id")

	var u user.User
	// SELECT * FROM users WHERE user_id = '(valuable userId)' ORDER BY id LIMIT 1;
	if err := db.GetDB().Where("user_id = ?", id).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No User found",
		})
		fmt.Println(err)
		c.Abort()
		return
	}

	// リクエストパラメータを受け取る
	var req user.User
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "JSON parse error",
			"cause":   "invalid request",
		})
		return
	}

	// ニックネームとコメントの存在確認
	if err := req.ValidateUsersPatch(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User updation failed",
			"cause":   "required nickname or comment",
		})
		return
	}

	// ユーザーネームとパスワードを変更しようとしていないか
	if err := req.ValidateExistUserIDandPassword(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User updation failed",
			"cause":   "not updatable user_id and password",
		})
		return
	}

	u.Comment = req.Comment
	u.Nickname = req.Nickname
	if req.Nickname == "" {
		u.Nickname = u.UserID
	}

	response := struct {
		Message string `json:"message"`
		Recipe  struct {
			Nickname string `json:"nickname"`
			Comment  string `json:"comment"`
		} `json:"recipe"`
	}{
		Message: "User Successfully updated",
		Recipe: struct {
			Nickname string `json:"nickname"`
			Comment  string `json:"comment"`
		}{
			Nickname: u.Nickname,
			Comment:  u.Comment,
		},
	}

	// UPDATE `users` SET `user_id` = 'u.UserID', `password` = 'u.Password', `nickname` = `u.Nickname`, `comment` = `u.Comment` WHERE `user_id` = 'id';
	err := db.GetDB().Model(&user.User{}).Where("user_id = ?", u.UserID).UpdateColumn(u).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "DB error update failure...",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func CloseHandler(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	userID, err := myauth.ExtractUserIDFromAuthHeader(auth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	var u user.User
	if err := db.GetDB().Where("user_id = ?", userID).First(&u).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No User found",
		})
		return
	}

	db.GetDB().Delete(&u)
	c.JSON(http.StatusOK, gin.H{
		"message": "Account and user successfully removed",
	})
}
