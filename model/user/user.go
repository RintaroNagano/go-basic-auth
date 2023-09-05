package user

import (
	"errors"
	"regexp"
)

type User struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Comment  string `json:"comment"`
}

// /ユーザIDとパスワードについて存在確認を実行する関数
func (u User) ValidateExistUserIDandPassword() error {
	if err := u.ValidateExistUserID(); err == nil {
		return err
	}

	if err := u.ValidateExistPassword(); err == nil {
		return err
	}

	return nil
}

// PATCH users/{id}についてバリデーションする関数
func (u User) ValidateUsersPatch() error {
	if len(u.Nickname) == 0 && len(u.Comment) == 0 {
		return errors.New("ニックネームとコメントが両方設定されていません")
	}

	return nil
}

// ユーザーIDのパターンについてバリデーションを実行する関数
func (u *User) ValidatePattern() error {
	// 半角英数字の正規表現パターン
	pattern := "^[a-zA-Z0-9]+$"
	match, err := regexp.MatchString(pattern, u.UserID)
	if err != nil {
		return errors.New("ユーザーIDのバリデーション中にエラーが発生しました")
	}

	if !match {
		return errors.New("ユーザーIDは半角英数字のみである必要があります")
	}

	// 半角英数字記号（空白と制御コードを除くASCII文字）の正規表現パターン
	pattern = "^[[:print:]]+$"
	match, err = regexp.MatchString(pattern, u.Password)
	if err != nil {
		return errors.New("パスワードのバリデーション中にエラーが発生しました")
	}
	if err != nil {
		return errors.New("パスワードのバリデーション中にエラーが発生しました")
	}

	if !match {
		return errors.New("パスワードは半角英数字記号のみである必要があります")
	}

	return nil
}

// ユーザーID, パスワードの長さについてバリデーションを実行する関数
func (u *User) ValidateLength() error {
	if len(u.UserID) < 6 || 20 < len(u.UserID) {
		return errors.New("ユーザーIDは6文字以上20文字以内である必要があります")
	}

	if len(u.Password) < 8 || 20 < len(u.Password) {
		return errors.New("パスワードは8文字以上20文字以内である必要があります")
	}

	return nil
}

// ニックネームの存在についてバリデーションを実行する関数
func (u *User) ValidateExistUserID() error {
	if u.UserID == "" {
		return errors.New("ユーザIDが存在しません")
	}

	return nil
}

// コメントの存在についてバリデーションを実行する関数
func (u *User) ValidateExistPassword() error {
	if u.Password == "" {
		return errors.New("パスワードが存在しません")
	}

	return nil
}

// ニックネームの存在についてバリデーションを実行する関数
func (u *User) ValidateExistNickname() error {
	if len(u.Nickname) == 0 {
		return errors.New("ニックネームが存在しません")
	}

	return nil
}

// コメントの存在についてバリデーションを実行する関数
func (u *User) ValidateExistComment() error {
	if len(u.Comment) == 0 {
		return errors.New("コメントが存在しません")
	}

	return nil
}
