package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        uint      `json:"id"`
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID        uint
	Name      string
	Price     int
	Count     int
	CreatedAt time.Time
}

type ApprovalInformation struct {
	ID           uint
	Status       string
	ApprovalDate time.Time
	Count        int
	UserID       string
	ProductID    string
	CreatedAt    time.Time
}

type signupReturn struct {
	status  int
	message string
}

type Claims struct {
	UserNo int
	jwt.StandardClaims
}

var (
	message = map[string]interface{}{
		"success": signupReturn{200, "회원가입 성공!"},
		"fail":    signupReturn{500, "이미 존재하는 아이디입니다."},
	}

	expirationTime = 5 * time.Minute
	JwtKey         = []byte("This application was made by K-brother")
)
