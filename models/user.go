package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type User struct {
	Id int
	Name string
	Pwd string
}

type UserAddObj struct {
	User []User
}

func (u User) Add(list []User) error {
	for index, item := range list {
		u.Name = strings.TrimSpace(item.Name)
		u.Pwd = strings.TrimSpace(item.Pwd)
		if u.Name == "" {
			var sb strings.Builder
			sb.WriteString("第")
			sb.WriteString(strconv.Itoa(index + 1))
			sb.WriteString("行姓名不能为空")

			return errors.New(sb.String())
		}
		max := 3
		if u.Pwd == "" || len(u.Pwd) < max {
			var sb strings.Builder
			sb.WriteString("第")
			sb.WriteString(strconv.Itoa(index + 1))
			sb.WriteString("行密码应大于等于")
			sb.WriteString(strconv.Itoa(max))
			sb.WriteString("个字符")

			return errors.New(sb.String())
		}
	}

	o := orm.NewOrm()
	_, err := o.InsertMulti(1, &list)

	return err
}

func (u User) List(name string) ([]User, error) {
	var list []User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	_, err := qs.Filter("name__icontains", strings.Trim(name, " ")).All(&list)

	return list, err
}
