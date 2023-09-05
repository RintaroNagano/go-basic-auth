# Gin API 使用ガイド

## 概要
/pingというapiと，basic認証の自己実装

## 動作確認手順
```
git clone git@github.com:RintaroNagano/go-basic-auth.git
cd go-bacic-auth
cp .env.sample .env
docker compose up
curl -X POST -H "Content-Type: application/json" --url http://localhost:8080/signup -d '{"user_id":"rintaro", "password":"password"}'
curl -X GET http://localhost:8080/users/rintaro -H "Authorization: Basic $(echo -n 'rintaro:password' | base64)"
curl -X PATCH http://localhost:8080/users/rintaro -H "Authorization: Basic $(echo -n 'rintaro:password' | base64)" -H "Content-Type: application/json" -d '{"nickname":"NewNickname","comment":"NewComment"}'
curl -X POST http://localhost:8080/close/ -H "Authorization: Basic $(echo -n 'rintaro:password' | base64)"
```

## 起動
サーバーは以下のコマンドで起動します。
```
go run main.go
```
起動すると、デフォルトで`0.0.0.0:8080`でlistenします。

## エンドポイントとその使用法

### 1. Ping
**Endpoint**: `/pong`  
**Method**: `GET`  
**Description**: サーバーの動作確認のためのシンプルなエンドポイント。  
**Response**:
```
{
    "message": "ping"
}
```

### 2. サインアップ (ユーザー作成)
**Endpoint**: `/signup`  
**Method**: `POST`  
**Description**: 新しいユーザーを作成します。  
**Request Body**:
```
{
    "user_id": "<USER_ID>",
    "password": "<PASSWORD>"
}
```
**Possible Errors**:  
- **"ユーザIDが存在しません"**: ユーザIDが存在しないまたは空の場合。
- **"パスワードが存在しません"**: パスワードが存在しないまたは空の場合。
- **"ユーザーIDは半角英数字のみである必要があります"**: ユーザIDが半角英数字以外の文字を含む場合。
- **"パスワードは半角英数字記号のみである必要があります"**: パスワードが半角英数字記号（空白や制御コードを除くASCII文字）以外の文字を含む場合。
- **"ユーザーIDは6文字以上20文字以内である必要があります"**: ユーザIDの長さが6文字未満または20文字を超える場合。
- **"パスワードは8文字以上20文字以内である必要があります"**: パスワードの長さが8文字未満または20文字を超える場合。

### 3. ユーザー情報取得
**Endpoint**: `/users/:id`  
**Method**: `GET`  
**Description**: 指定されたユーザIDのユーザー情報を取得します。  
**Headers**:  
`Authorization: Basic <BASE64_ENCODED_USERID_AND_PASSWORD>`  
**Possible Errors**:  
- "Authentication Failed": 認証情報がヘッダに存在しない、または不正な形式の場合。
- **"ニックネームが存在しません"**: ニックネームが存在しないまたは空の場合。
- **"コメントが存在しません"**: コメントが存在しないまたは空の場合。
- **"ニックネームとコメントが両方設定されていません"**: PATCHリクエストでニックネームとコメントの両方が空の場合。

### 4. ユーザー情報更新
**Endpoint**: `/users/:id`  
**Method**: `PATCH`  
**Description**: 指定されたユーザIDのユーザー情報を更新します。  
**Headers**:  
`Authorization: Basic <BASE64_ENCODED_USERID_AND_PASSWORD>`  
**Request Body**:
```
{
    "nickname": "<NICKNAME>",
    "comment": "<COMMENT>"
}
```
**Possible Errors**:  
- "Authentication Failed": 認証情報がヘッダに存在しない、または不正な形式の場合。
- "No Permission for Update": AuthorizationヘッダのユーザーIDがパスパラメータのIDと一致しない場合。

### 5. アカウント削除
**Endpoint**: `/close`  
**Method**: `POST`  
**Description**: 認証されたユーザーのアカウントを削除します。  
**Headers**:  
`Authorization: Basic <BASE64_ENCODED_USERID_AND_PASSWORD>`  
**Possible Errors**:  
- "Authentication Failed": 認証情報がヘッダに存在しない、または不正な形式の場合。

---

## 認証

一部のエンドポイント（ユーザー情報取得、ユーザー情報更新、アカウント削除）は、`Authorization`ヘッダーを必要とします。  
このヘッダーは`Basic <BASE64_ENCODED_USERID_AND_PASSWORD>`の形式であり、<BASE64_ENCODED_USERID_AND_PASSWORD>はユーザIDとパスワードをコロンで結合した文字列（例: `user123:password123`）をBase64でエンコードしたものです。

---

以上がAPIの使用方法となります。エンドポイントの正確な動作やエラーメッセージの詳細は、実際の実装を参照してください。