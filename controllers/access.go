package controllers

import (
	"encoding/json"
	"passapp-engine-api/models"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/validation"
	"log"
)

// Operations about Access Points
type AccessController struct {
	beego.Controller
}

// @Title LogAccess
// @Description log access points
// @Param	body		body 	models.User	true		"body for log report content"
// @Success 200 {int} models.JWT
// @Failure 403 body is empty
// @router / [post]
func (c *AccessController) Post() {
	var access models.Access
	fmt.Println(c.Ctx.Input.RequestBody);
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &access)
	
	if err != nil {
		panic(err)
	}

	valid := validation.Validation{}
	b, err := valid.Valid(access)
	
	if err != nil {
		panic(err)
	}

	if !b {
		// validation does not pass
		// print invalid message
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	res := models.AddAccess(access)
	c.Data["json"] = map[string]string{"access" :res}
	c.ServeJSON()
}

// @Title GetAll
// @Description get all Access
// @Success 200 {object} models.Access
// @router / [get]
func (c *AccessController) GetAll() {
	access := models.GetAllAccess()
	c.Data["json"] = access
	c.ServeJSON()
}