package model

import (
	"GodFather-server/ent"
	"GodFather-server/ent/user"
	"context"
	"log"

	"time"
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

var (
	message = map[string]interface{}{
		"success": signupReturn{200, "회원가입 성공!"},
		"fail":    signupReturn{500, "이미 존재하는 아이디입니다."},
	}
)

func (s User) IsValid(c *ent.Client) (ok bool) {
	_, err := c.User.Query().Select(user.FieldUserID).Where(user.UserID(s.UserID)).First(context.Background())
	if err != nil {
		log.Println(err)
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
