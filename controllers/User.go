package controllers

import (
	"renshihoutai/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

// 注册展示页面
func (c *RegisterController) ShowRegister() {
	c.TplName = "register.html"
}

// 注册获取数据页面
func (c *RegisterController) HandleRegister() {
	// 获取浏览器传递的值，并去除两边的空格

	Name := strings.TrimSpace(c.GetString("userName"))
	Pwd := strings.TrimSpace(c.GetString("passWord"))
	// beego.Info("账号:", Name, "密码:", Pwd)
	// 数据处理
	if Name == "" || Pwd == "" {
		beego.Info("用户名或密码不能为空")
		c.TplName = "register.html"
		c.Data["errmsg"] = "用户名或密码不能为空 ！"
		return
	}
	// 插入数据库（数据库表，Users）
	//获取orm对象
	o := orm.NewOrm()
	//   获取插入对象
	user := models.Users{}
	//   插入数值
	user.Name = Name
	user.Pwd = Pwd

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("插入数据失败，用户相同或者其他错误！！！")
		c.TplName = "register.html"
		c.Data["errmsg"] = "插入数据失败，用户相同或者其他错误！！！！"
		return
	}
	// 测试返回视图
	// c.Ctx.WriteString("注册成功！！！")
	// 实际情况注册成功返回到登录页面
	c.Redirect("login", 302)
}

type LoginController struct {
	beego.Controller
}

// 登录页面 get
func (c *LoginController) ShowLogin() {
	c.TplName = "login.html"
}

// 登录页面 post
func (c *LoginController) HandleLogin() {
	// 拿到浏览器数据，并去除两边的空格
	Name := strings.TrimSpace(c.GetString("userName"))
	Pwd := strings.TrimSpace(c.GetString("passWord"))
	beego.Info("账号:", Name, "密码:", Pwd)

	//数据处理
	if Name == "" || Pwd == "" {
		beego.Info("登录失败！！")
		c.TplName = "login.html"
		c.Data["errmsg"] = "登录失败！！！！！"
		return
	}
	// 查找数据库
	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var user models.Users
	// 查询
	user.Name = Name
	err := o.Read(&user, "Name")
	if err != nil {
		beego.Info("用户名登录失败！！！")
		c.TplName = "login.html"
		c.Data["errmsg"] = "用户名登录失败！！！！！"
		return
	}

	// 判断密码是否一致
	if user.Pwd != Pwd {
		beego.Info("密码登录失败！！！")
		c.TplName = "login.html"
		c.Data["errmsg"] = "密码登录失败！！"
		return
	}
	// 测试返回视图
	// c.Ctx.WriteString("登录成功")
	// 实际情况登录成功跳转到后台页面
	if Name == "admin" {

		//设置cookie
		c.Ctx.SetCookie("username", Name, time.Second*60*60)
		c.Redirect("adminhoutai.html", 302)

	} else {
		//设置cookie
		c.Ctx.SetCookie("username", Name, time.Second*60*60)
		c.Redirect("userhoutai", 302)
	}

}

type HoutaiController struct {
	beego.Controller
}

// 管理员后台页面(员工资料页面) get
func (c *HoutaiController) ShowHoutai() {

	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var yuangongziliao []models.YuanGongZiLiao
	_, err := o.QueryTable("yuan_gong_zi_liao").All(&yuangongziliao)
	if err != nil {
		beego.Info("查询所有用户信 息出错")
		return
	}
	c.Data["yuangongziliao"] = yuangongziliao

	c.TplName = "adminhoutai.html"

}

// 管理员后台页面(员工资料页面) post
func (c *HoutaiController) HandleHoutai() {
	// 获取浏览器传递的值
	Id := c.Input().Get("BianHao")
	c.Ctx.SetCookie("bianhao", Id, time.Second*60*60)
	c.Redirect("adyuangongziliaosousuo", 302)
}

type YuanGongGongZiController struct {
	beego.Controller
}

// 管理员后台页面(员工工资页面) get
func (c *YuanGongGongZiController) ShowYuanGongGongZi() {
	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var yuangonggongzi []models.YuanGongZiLiao
	_, err := o.QueryTable("yuan_gong_zi_liao").All(&yuangonggongzi)
	if err != nil {
		beego.Info("查询所有用户信 息出错")
		return
	}
	c.Data["yuangonggongzi"] = yuangonggongzi

	c.TplName = "yuangonggongzi.html"

}

// 管理员后台页面(员工工资页面) post
func (c *YuanGongGongZiController) HandleYuanGongGongZi() {
	// 获取浏览器传递的值
	Id := c.Input().Get("BianHao")
	c.Ctx.SetCookie("bianhao", Id, time.Second*60*60)
	c.Redirect("adyuangongziliaosousuo", 302)
}

type BuMenXinXiController struct {
	beego.Controller
}

// 管理员后台页面(部门信息页面) get
func (c *BuMenXinXiController) ShowBuMenXinXi() {
	//获取orm对象
	o := orm.NewOrm()
	//获取查询对象
	var bumenxinxi []models.BuMenXinXi
	_, err := o.QueryTable("bu_men_xin_xi").All(&bumenxinxi)
	if err != nil {
		beego.Info("查询所有用户信息出错")
		return
	}
	c.Data["bumenxinxi"] = bumenxinxi

	c.TplName = "bumenxinxi.html"

}

// 管理员后台页面(部门信息页面) post
func (c *BuMenXinXiController) HandleBuMenXinXi() {
	// 获取浏览器传递的值
	Id := c.Input().Get("BianHao")
	c.Ctx.SetCookie("bianhao", Id, time.Second*60*60)
	c.Redirect("adbumenxinxisousuo", 302)
}

type RuZhiXinXiController struct {
	beego.Controller
}

// 管理员后台页面(入职信息页面) get
func (c *RuZhiXinXiController) ShowRuZhiXinXi() {
	c.TplName = "ruzhixinxi.html"
}

// 管理员后台页面(入职信息页面) post
func (c *RuZhiXinXiController) HandleRuZhiXinXi() {
	// 获取浏览器传递的值，并去除两边的空格
	Id := c.Input().Get("id")
	intid, _ := strconv.Atoi(Id)
	Name := c.GetString("name")
	Department := c.GetString("department")
	Entrytime := c.GetString("entrytime")
	Salary := c.Input().Get("salary")
	intsalary, _ := strconv.Atoi(Salary)
	// beego.Info("账号:", Name, "密码:", Pwd)
	// 数据处理
	if intid == 0 || Name == "" || Department == "" || Entrytime == "" || intsalary == 0 {
		beego.Info("不能有空白项")
		c.TplName = "ruzhixinxi.html"
		return
	}
	// 插入数据库（数据库表，Users）
	//获取ocrm对象
	o := orm.NewOrm()
	//   获取插入对象
	yuangongziliao := models.YuanGongZiLiao{}
	//   插入数值
	yuangongziliao.Id = intid
	yuangongziliao.Name = Name
	yuangongziliao.Department = Department
	yuangongziliao.Entrytime = Entrytime
	yuangongziliao.Salary = intsalary

	_, err := o.Insert(&yuangongziliao)
	if err != nil {
		beego.Info("插入数据失败，用户相同或者其他错误！！！")
		c.TplName = "ruzhixinxi.html"
		c.Data["errmsg"] = "插入数据失败，用户相同或者其他错误！！！！"
		return
	}

	// 实际情况添加成功返回到登录页面
	c.Redirect("adminhoutai", 302)

}

type AdminUpdateController struct {
	beego.Controller
}

// 管理员修改页面 get
func (c *AdminUpdateController) ShowAdminUpdate() {
	c.TplName = "adminupdate.html"

}

// 管理员修改页面 post
func (c *AdminUpdateController) HandleAdminUpdate() {

	// 获取浏览器传递的值，并去除两边的空格
	Id := c.Input().Get("id")
	intid, _ := strconv.Atoi(Id)
	Name := c.GetString("name")
	Department := c.GetString("department")
	Entrytime := c.GetString("entrytime")
	Salary := c.Input().Get("salary")
	intsalary, _ := strconv.Atoi(Salary)
	// 数据处理
	if intid == 0 || Name == "" || Department == "" || Entrytime == "" || intsalary == 0 {
		beego.Info("不能有空白项")
		c.TplName = "adminupdate.html"
		return
	}
	// 插入数据库（数据库表，Users）
	//获取orm对象
	o := orm.NewOrm()
	//   获取更新对象
	yuangongziliao := models.YuanGongZiLiao{intid, Name, Department, Entrytime, intsalary}
	//   插入数值
	_, err := o.Update(&yuangongziliao)
	if err != nil {
		beego.Info("插入数据失败，用户相同或者其他错误！！！")
		c.TplName = "adminupdate.html"
		return
	}

	// 实际情况添加成功返回到登录页面
	c.Redirect("adminupdate", 302)
}

type AdminDeleteController struct {
	beego.Controller
}

// 管理员删除页面 get
func (c *AdminDeleteController) ShowAdminDelete() {
	c.TplName = "admindelete.html"

}

// 管理员删除页面 post
func (c *AdminDeleteController) HandleAdminDelete() {

	// 获取浏览器传递的值，并去除两边的空格
	Id := c.Input().Get("id")
	intid, _ := strconv.Atoi(Id)
	// 数据处理
	if intid == 0 {
		beego.Info("不能有空白项")
		c.TplName = "admindelete.html"
		return
	}
	// 插入数据库（数据库表，yuangonggongzi）
	//获取orm对象
	o := orm.NewOrm()
	user := models.YuanGongZiLiao{Id: intid}

	_, err := o.Delete(&user)
	if err == nil {
		beego.Info("删除成功")
	}

	// 实际情况添加成功返回到登录页面
	c.Redirect("admindelete", 302)
}

type AdminSouSuoYuanGongZiLiaoController struct {
	beego.Controller
}

// 管理员搜索员工资料 get
func (c *AdminSouSuoYuanGongZiLiaoController) ShowSouSuoYuanGongZiLiao() {
	//获取cookie
	Id := c.Ctx.GetCookie("bianhao")
	bianhao, _ := strconv.Atoi(Id)
	//获取orm对象
	orm := orm.NewOrm()
	//获取查询对象
	var yuangongziliao []models.YuanGongZiLiao
	qs := orm.QueryTable("yuan_gong_zi_liao")
	// select * from student where Id =intid;
	err := qs.Filter("Id__exact", bianhao).One(&yuangongziliao) // 过滤器
	if err != nil {
		beego.Info(yuangongziliao)
	} else {
		beego.Info(yuangongziliao)
		beego.Info("查询成功")
	}
	c.Data["yuangongziliao"] = yuangongziliao
	c.TplName = "adyuangongziliaosousuo.html"

}

// 管理员搜索员工资料 post
func (c *AdminSouSuoYuanGongZiLiaoController) HandleSouSuoYuanGongZiLiao() {

}

type AdminBuMenXinXiSouSuoController struct {
	beego.Controller
}

// 部门信息搜索页面 get
func (c *AdminBuMenXinXiSouSuoController) ShowBuMenXinXi() {
	//获取cookie
	Id := c.Ctx.GetCookie("bianhao")
	bianhao, _ := strconv.Atoi(Id)
	//获取orm对象
	orm := orm.NewOrm()
	//获取查询对象
	var bumenxinxi []models.BuMenXinXi
	qs := orm.QueryTable("bu_men_xin_xi")
	// select * from student where Id =intid;
	err := qs.Filter("Id__exact", bianhao).One(&bumenxinxi) // 过滤器
	if err != nil {
		beego.Info("查询所有用户信息出错")
		return
	}
	c.Data["bumenxinxi"] = bumenxinxi

	c.TplName = "bumenxinxi.html"

}

// 部门信息搜索页面 post
func (c *AdminBuMenXinXiSouSuoController) HandleBuMenXinXi() {

}

type AdminYuanGongGongZiSouSuoController struct {
	beego.Controller
}

// 员工工资搜索页面 get
func (c *AdminYuanGongGongZiSouSuoController) ShowAdminYuanGongGongZiSouSuo() {
	//获取cookie
	Id := c.Ctx.GetCookie("bianhao")
	bianhao, _ := strconv.Atoi(Id)
	//获取orm对象
	orm := orm.NewOrm()
	//获取查询对象
	var yuangongziliao []models.YuanGongZiLiao
	qs := orm.QueryTable("yuan_gong_zi_liao")
	// select * from student where Id =intid;
	err := qs.Filter("Id__exact", bianhao).One(&yuangongziliao) // 过滤器
	if err != nil {
		beego.Info(yuangongziliao)
	} else {
		beego.Info(yuangongziliao)
		beego.Info("查询成功")
	}
	c.Data["yuangongziliao"] = yuangongziliao
	c.TplName = "adyuangongziliaosousuo.html"

}

// 员工工资搜索页面 post
func (c *AdminYuanGongGongZiSouSuoController) HandleAdminYuanGongGongZiSouSuo() {

}

type UserHoutaiController struct {
	beego.Controller
}

// 员工后台页面(员工资料页面) get
func (c *UserHoutaiController) ShowUserHoutai() {

	//获取cookie
	username := c.Ctx.GetCookie("username")
	id, _ := strconv.Atoi(username)
	//获取orm对象
	orm := orm.NewOrm()
	//获取查询对象
	var userziliao []models.YuanGongZiLiao
	qs := orm.QueryTable("yuan_gong_zi_liao")
	// select * from student where Name =username;
	err := qs.Filter("Id__exact", id).One(&userziliao) // 过滤器
	if err != nil {
		beego.Info(userziliao)
	} else {
		beego.Info(userziliao)
		beego.Info("查询失败")
	}
	c.Data["userziliao"] = userziliao
	c.TplName = "userhoutai.html"

}

// 员工后台页面(员工资料页面)post
func (c *UserHoutaiController) HandleUserHoutai() {

}

type UserGongZiController struct {
	beego.Controller
}

// 员工后台页面(员工工资页面) get
func (c *UserGongZiController) ShowUserGongZi() {
	//获取cookie
	username := c.Ctx.GetCookie("username")
	id, _ := strconv.Atoi(username)
	//获取orm对象
	orm := orm.NewOrm()
	//获取查询对象
	var usergongzi []models.YuanGongZiLiao
	qs := orm.QueryTable("yuan_gong_zi_liao")
	// select * from student where Name =username;
	err := qs.Filter("Id__exact", id).One(&usergongzi) // 过滤器
	if err != nil {
		beego.Info(usergongzi)
	} else {
		beego.Info(usergongzi)
		beego.Info("查询失败")
	}
	c.Data["usergongzi"] = usergongzi
	c.TplName = "usergongzi.html"

}

// 员工后台页面(员工工资页面) post
func (c *UserGongZiController) HandleUserGongZi() {

}

type UserUpdateController struct {
	beego.Controller
}

// 员工后台页面(员工修改页面) get
func (c *UserUpdateController) ShowUserUpdate() {
	c.TplName = "userupdate.html"

}

// 员工后台页面(员工修改页面) post
func (c *UserUpdateController) HandleUserUpdate() {

	// 获取浏览器传递的值，并去除两边的空格
	Id := c.Input().Get("id")
	intid, _ := strconv.Atoi(Id)
	Name := c.GetString("name")
	Department := c.GetString("department")
	Entrytime := c.GetString("entrytime")
	Salary := c.Input().Get("salary")
	intsalary, _ := strconv.Atoi(Salary)
	// beego.Info("账号:", Name, "密码:", Pwd)
	// 数据处理
	if intid == 0 || Name == "" || Department == "" || Entrytime == "" || intsalary == 0 {
		beego.Info("不能有空白项")
		c.TplName = "userupdate.html"
		return
	}
	// 插入数据库（数据库表，Users）
	//获取orm对象
	o := orm.NewOrm()
	//   获取插入对象
	yuangongziliao := models.YuanGongZiLiao{}
	//   插入数值
	yuangongziliao.Id = intid
	yuangongziliao.Name = Name
	yuangongziliao.Department = Department
	yuangongziliao.Entrytime = Entrytime
	yuangongziliao.Salary = intsalary

	_, err := o.Update(&yuangongziliao)
	if err != nil {
		beego.Info("插入数据失败，用户相同或者其他错误！！！")
		c.TplName = "userupdate.html"
		return
	}

	// 实际情况添加成功返回到登录页面
	c.Redirect("userhoutai", 302)
}
