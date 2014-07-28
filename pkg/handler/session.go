package handler

//
// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/mholt/binding"
//
// 	db "github.com/lmdtfy/lmdtfy/pkg/database"
// 	m "github.com/lmdtfy/lmdtfy/pkg/model"
// )
//
// var ErrLoginFailed = "Login Failed! Email and/or password incorrect."
//
// // LoginForm ...
// type LoginForm struct {
// 	Email    string
// 	Password string
// }
//
// // FieldMap ...
// func (lf *LoginForm) FieldMap() binding.FieldMap {
// 	return binding.FieldMap{
// 		&lf.Email:    binding.Field{Form: "email", Required: true},
// 		&lf.Password: binding.Field{Form: "password", Required: true},
// 	}
// }
//
// // Login ...
// func Login(c *gin.Context) {
// 	// Get username and password
// 	lf := new(LoginForm)
//
// 	errs := binding.Bind(c.Request, lf)
// 	if errs != nil {
// 		c.JSON(400, errs)
// 		return
// 	}
//
// 	user, err := db.Users.FindByEmail(lf.Email)
// 	if err != nil {
// 		c.JSON(401, ErrLoginFailed)
// 		return
// 	}
//
// 	if err := user.ComparePassword(lf.Password); err != nil {
// 		c.JSON(401, ErrLoginFailed)
// 		return
// 	}
//
// 	// Create a session for the user.
// 	session, err := m.NewSession(user)
// 	if err != nil {
// 		c.JSON(500, err)
// 		return
// 	}
//
// 	// Send token to the user so they can use it to to authenticate all further requests.
// 	c.JSON(200, &gin.H{"session": []m.Session{session}})
// 	return
// }
