package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type Result struct {
	Code int
}

func (c *MainController) Get() {
	type User struct {
		Id int
		Name string
		Age int
	}
	arr := []User{
		{ 1, "John", 24 },
		{ 2, "Bob", 32 },
		{ 3, "Lily", 18 },
	}

	for index, val := range arr {
		fmt.Println(index)
		fmt.Println(val.Name)
	}

	c.Data["json"] = arr
	c.ServeJSON()
}

func (c *MainController) Post() {

}
