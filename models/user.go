package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
)

func init() {
	orm.RegisterModel(new(User))
	//orm.RegisterModel(new(UserProfile))
}
// 进行用户的增删改查
type User struct {
	Id          int
	Phone       string
	//UserProfile *UserProfile `orm:"rel(one)"`
	Password    string
	Status      int
	Created     int64
	Changed     int64
}
func (this *User) TableName() string {
	return "user"
}
func LoginUser(phone string, password string) (err error, users []User) {
	// 1.创建查询条件
	cond := orm.NewCondition()
	cond = cond.And("phone", phone)
	cond = cond.And("password", com.Md5(password))
	cond = cond.And("status", 1)

	// 2.查询数据库
	qs := orm.NewOrm().QueryTable("user")
	qs = qs.SetCond(cond)
	err = qs.Limit(1).One(&users)
	return err, users
}

func GetUser(id int) (user User, err error) {
	// 1.构建查询条件
	user = User{Id: id}

	// 2.进行条件查询
	err = orm.NewOrm().Read(&user)
	if err == orm.ErrNoRows {
		return user, nil
	}
	return user, err
}

func UpdatePassword(id int, oldPwd string, newPwd string) error {
	// 1.构建查询条件
	user := User{Id: id}

	// 2.进行条件查询
	o := orm.NewOrm()
	err := o.Read(&user)
	if nil != err {
		return err
	} else {
		// 3.更新数据库
		if user.Password == com.Md5(oldPwd) {
			user.Password = com.Md5(newPwd)
			_, err := o.Update(&user)
			return err
		} else {
			return fmt.Errorf("验证出错")
		}
	}
}
