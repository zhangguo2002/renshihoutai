package routers

import (
	"renshihoutai/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.RegisterController{}, "get:ShowRegister;post:HandleRegister")
	beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/adminhoutai", &controllers.HoutaiController{}, "get:ShowHoutai;post:HandleHoutai")
	beego.Router("/yuangonggongzi", &controllers.YuanGongGongZiController{}, "get:ShowYuanGongGongZi;post:HandleYuanGongGongZi")
	beego.Router("/bumenxinxi", &controllers.BuMenXinXiController{}, "get:ShowBuMenXinXi;post:HandleBuMenXinXi")
	beego.Router("/ruzhixinxi", &controllers.RuZhiXinXiController{}, "get:ShowRuZhiXinXi;post:HandleRuZhiXinXi")
	beego.Router("/userhoutai", &controllers.UserHoutaiController{}, "get:ShowUserHoutai")
	beego.Router("/usergongzi", &controllers.UserGongZiController{}, "get:ShowUserGongZi")
	beego.Router("/userupdate", &controllers.UserUpdateController{}, "get:ShowUserUpdate;post:HandleUserUpdate")
	beego.Router("/adminupdate", &controllers.AdminUpdateController{}, "get:ShowAdminUpdate;post:HandleAdminUpdate")
	beego.Router("/admindelete", &controllers.AdminDeleteController{}, "get:ShowAdminDelete;post:HandleAdminDelete")
	beego.Router("/adyuangongziliaosousuo", &controllers.AdminSouSuoYuanGongZiLiaoController{}, "get:ShowSouSuoYuanGongZiLiao")
	beego.Router("/adbumenxinxisousuo", &controllers.AdminSouSuoYuanGongZiLiaoController{}, "get:ShowSouSuoYuanGongZiLiao")
	beego.Router("/adyuangongziliaosousuo", &controllers.AdminSouSuoYuanGongZiLiaoController{}, "get:ShowSouSuoYuanGongZiLiao")
}
