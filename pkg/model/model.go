package model

import (
	"GodFather-server/ent"
	"GodFather-server/ent/user"
	"context"
	"fmt"
	"log"

	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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

func (s User) IsValid(c *ent.Client) (ok bool) {
	_, err := c.User.Query().Select(user.FieldUserID).Where(user.UserID(s.UserID)).First(context.Background())
	if err != nil {
		log.Println(err)
		pwHash, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
		}

		s.Password = string(pwHash)
		s.InsertUser(c)
		return true
	}

	// if u.UserID != "" {
	// 	s.InsertUser(c)
	// 	return true
	// }
	return false
}

func (s User) InsertUser(c *ent.Client) {
	_, err := c.User.Create().
		SetUserID(s.UserID).
		SetPassword(s.Password).
		SetName(s.Name).
		SetRole(user.Role(s.Role)).
		Save(context.Background())

	if err != nil {
		panic(err)
	}
}

func (s User) Signin(c *ent.Client) bool {
	user, err := c.User.Query().Select(user.FieldUserID, user.FieldPassword).Where(user.UserID(s.UserID)).First(context.Background())
	if err != nil {
		log.Println(err)
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(s.Password))
	if err2 != nil {
		return false
	} else {
		return true
	}
}

func (s User) GetToken() (string, error) {
	expirationTime := time.Now().Add(expirationTime)

	claims := Claims{
		UserNo: int(s.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("token signed Error")
	} else {
		return tokenString, nil
	}
}
