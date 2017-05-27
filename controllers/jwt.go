package controllers

import (
	"encoding/json"
	"passapp-engine-api/models"

	"github.com/astaxie/beego"
)

// Operations about JWT
type JWTController struct {
	beego.Controller
}

// @Title CreateToken
// @Description create tokens
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.JWT
// @Failure 403 body is empty
// @router / [post]
func (j *JWTController) Post() {
	var user models.User
	json.Unmarshal(j.Ctx.Input.RequestBody, &user)
	token := models.AddToken(user)
	j.Data["json"] = map[string]string{"token" :token}
	j.ServeJSON()
}
