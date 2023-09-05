package user

import (
	"testing"
)

func TestValidateExistUserID(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want bool
	}{
		{
			name: "user id exists",
			u:    User{UserID: "test"},
			want: true,
		},
		{
			name: "user id does not exist",
			u:    User{UserID: ""},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.u.ValidateExistUserID()
			if (err == nil) != tt.want {
				t.Errorf("ValidateExistUserID() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestValidateExistPassword(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "password exists",
			user: User{Password: "password"},
			want: true,
		},
		{
			name: "password does not exist",
			user: User{Password: ""},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.ValidateExistPassword()
			if (err == nil) != tt.want {
				t.Errorf("ValidateExistPassword() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestValidatePattern(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "valid user id and password",
			user: User{UserID: "test123", Password: "password!@#"},
			want: true,
		},
		{
			name: "invalid user id",
			user: User{UserID: "test@", Password: "password!@#"},
			want: false,
		},
		{
			name: "invalid password",
			user: User{UserID: "test123", Password: "passwordあ"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.ValidatePattern()
			if (err == nil) != tt.want {
				t.Errorf("ValidatePattern() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestValidateLength(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "valid length of user id and password",
			user: User{UserID: "test123", Password: "password!@#"},
			want: true,
		},
		{
			name: "invalid length of user id",
			user: User{UserID: "test", Password: "password!@#"},
			want: false,
		},
		{
			name: "invalid length of password",
			user: User{UserID: "test123", Password: "pass"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.ValidateLength()
			if (err == nil) != tt.want {
				t.Errorf("ValidateLength() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

// ValidateExistUserID and ValidateExistPassword are already tested above.

func TestValidateExistNickname(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{
			name: "nickname exists",
			user: User{Nickname: "nickname"},
			want: true,
		},
		{
			name: "nickname does not exist",
			user: User{Nickname: ""},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.ValidateExistNickname()
			if (err == nil) != tt.want {
				t.Errorf("ValidateExistNickname() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestValidateExistComment(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want bool
	}{
		{
			name: "comment exists",
			u:    User{Comment: "comment"},
			want: true,
		},
		{
			name: "comment does not exist",
			u:    User{Comment: ""},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.u.ValidateExistComment()
			if (err == nil) != tt.want {
				t.Errorf("ValidateExistComment() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
