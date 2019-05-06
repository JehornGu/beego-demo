package controllers

import (
	"civ/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// User APIs
type UserController struct {
	beego.Controller
}

// @Title GetUserList
// @Description Get user list
// @Param  name query string "The filter key of Name"
// @router /list [get]
func (c *UserController) List() {
	r := models.Result{}
	usr := models.User{}
	name := c.GetString("name")

	userList, err := usr.List(name)

	r.Data = userList
	r.Code = models.SUCCESS
	r.Msg = "查询成功"

	if err != nil {
		r.Data = nil
		r.Code = models.ERR_DATA
		r.Msg = "查询用户不存在"
	}

	c.Data["json"] = r
	c.ServeJSON()
}

// @Title AddUser
// @Description Add users
// @Param   user     formData    models.UserAddObj        "The user list for add"
// @router /add [post]
func (c *UserController) Add() {
	r := models.Result{}

	usr := models.User{}

	var err error
	var obj models.UserAddObj
	var errs []error
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &obj); err == nil {

		err := usr.Add(obj.User)

		if err != nil {
			errs = append(errs, err)
		}

		r.Data = nil
		r.Msg = "新增用户成功"
		r.Code = models.SUCCESS

		if len(errs) > 0 {
			r.Data = nil
			r.Msg = errs[0].Error()
			r.Code = models.ERR_DATA
		}

	} else {
		r.Data = nil
		r.Msg = err.Error()
		r.Code = models.ERR_BIZ
	}

	c.Data["json"] = r
	c.ServeJSON()
}
