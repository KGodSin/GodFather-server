package v1handler

import (
	"GodFather-server/pkg/model"
	"GodFather-server/pkg/util"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	url     = "http://localhost:3333"
	urlList = map[string]string{
		"signup": url + "/auth/signup",
		"signin": url + "/auth/signin",
	}
)

func Signup(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	pbytes, _ := json.Marshal(user)
	buff := bytes.NewBuffer(pbytes)

	res := util.PostJSON(urlList["signup"], "application/json", buff)
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body) // 나중에 응답 읽은걸로 보내줘야됨
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"user_id": user.UserID,
		"role":    user.Role,
		"name":    user.Name,
	})
}

func Signin(c *gin.Context) {

}
