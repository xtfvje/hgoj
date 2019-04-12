package controllers

import (
	//"errors"
	"github.com/astaxie/beego"
	//"github.com/yinrenxin/joeblog/syserror"

	//uuid "github.com/satori/go.uuid"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/syserror"
)

//const SESSION_USER_KEY = "SESSION_USER_KEY"

//type MAP_H = map[string]interface{}

type BaseController struct {
	beego.Controller
	User    models.Users
	IsLogin bool
	IsAdmin bool
}

const OJ_DATA = "./judge/data"

type MAP_H = map[string]interface{}

const SESSION_USER_KEY = "JOE"

func (this *BaseController) Prepare() {
	path := this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SESSION_USER_KEY).(models.Users)
	//logs.Info("登录id：",u.UserId, "登录邮箱：",u.Email)
	this.IsLogin = false
	this.IsAdmin = false
	if ok {
		this.User = u
		this.IsLogin = true
		if u.Role == 1 {
			this.IsAdmin = true
		}
		this.Data["User"] = this.User
	}
	//logs.Info("是否登录", this.IsLogin)
	this.Data["Path"] = path
	this.Data["islogin"] = this.IsLogin
	this.Data["isadmin"] = this.IsAdmin
	this.Data["menu"] = path
}

func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}

func (this *BaseController) Abort401(err error) {
	this.Data["error"] = err
	this.Abort("401")
}

func (this *BaseController) GetMushString(key, msg string) string {
	k := this.GetString(key)
	if len(k) == 0 {
		//this.Abort500(errors.New(msg))
		this.JsonErr(msg, syserror.KEY_NOT_NULL, "")
	}
	return k
}



/**
	过滤代码
 */
func (this *BaseController) FilterSource(key, msg string) (string, int) {
	k := this.GetString(key)
	if len(k) == 0 {
		this.JsonErr("代码不能为空", syserror.SOURCE_NOT_NULL, "/problem")
	}
	return k, len(k)
}

func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
}

func (this *BaseController) JsonErr(msg string, code int32, action string) {
	this.Data["json"] = MAP_H{
		"code":   code,
		"msg":    msg,
		"action": action,
	}
	this.ServeJSON()
	this.Abort("500")
}

func (this *BaseController) JsonOK(msg, action string) {
	this.Data["json"] = MAP_H{
		"code":   0,
		"msg":    msg,
		"action": action,
	}
	this.ServeJSON()
}

func (this *BaseController) JsonOKH(msg string, data MAP_H) {
	data["code"] = 0
	data["msg"] = msg
	this.Data["json"] = data
	this.ServeJSON()
}

//func (this *BaseController) UUID() string {
//	u, err := uuid.NewV4()
//	if err != nil {
//		this.Abort500(syserror.New("系统错误", nil))
//	}
//	return u.String()
//}

