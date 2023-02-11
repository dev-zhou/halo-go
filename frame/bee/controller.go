package bee

import (
	"HaloAdmin/base/status"
	"HaloAdmin/halo/_jwt"
	"HaloAdmin/halo/types"
	"HaloAdmin/halo/utils/_str"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/gookit/validate"
	"time"
)

// Controller 基准控制器
type Controller struct {
	web.Controller
}

var Valid = validation.Validation{}

// #### Response 响应处理

// DATA SUCCESS 操作成功
func (c *Controller) DATA(data interface{}, code int, message string) {
	c.Data["json"] = &types.Data{
		Data:    data,
		Code:    code,
		Message: message,
	}
	e := c.ServeJSON()
	if e != nil {
		println(e)
	}
	c.StopRun()
}

// INFO ERROR 操作失败
func (c *Controller) INFO(code int, message string) {
	c.Data["json"] = &types.Info{
		Code:    code,
		Message: message,
	}
	e := c.ServeJSON()
	if e != nil {
		println(e)
	}
	c.StopRun()
}

// PAGE 分页处理
func (c *Controller) PAGE(data interface{}, page int64, count int64, limit int64, code int, message string) {
	c.Data["json"] = &types.Page{
		Page:      page,
		Limit:     limit,
		Count:     count,
		PageCount: (count + limit - 1) / limit,
		Data:      data,
		Code:      code,
		Message:   message,
	}
	e := c.ServeJSON()
	if e != nil {
		println(e)
	}
	c.StopRun()
}

func (c *Controller) CurrentUserInfo() (userInfo _jwt.UserInfo) {
	println("全局登录用户")
	ctx := c.Ctx
	token := ctx.Input.Header("Authorization")
	token = _str.Replace(token, "Bearer ", "")
	userInfo = _jwt.UserInfo{}
	println("全局登录用户 token: ", token)
	claim, err := _jwt.ParseToken(token)

	// claim 存在nil情况
	if claim != nil {
		fmt.Printf("%#v\n", claim)
		fmt.Printf("UserId: %#v\n", claim.Id)
		fmt.Printf("解析token err: %#v\n", err)
	}

	if err != nil {
		// errors.New("Token Parse Fail")
		fmt.Printf("Token Parse Fail")
		return userInfo // 未知用户
	} else if time.Now().Unix() > claim.ExpiresAt {
		// errors.New("fmt.Printf("解析token err: %#v\n", err)")
		fmt.Printf("Token Expired\n")
		fmt.Printf("time.Now.Unix: %v\n", time.Now().Unix())
		fmt.Printf("设置过期时间: %v\n", claim.ExpiresAt)
		return userInfo // 未知用户
	}
	userInfo.Id = claim.Id
	userInfo.Uuid = claim.Uuid
	userInfo.Username = claim.Username
	fmt.Printf("UserInfo %#v\n", userInfo)
	return userInfo
}

// ReqFormBind 请求参数绑定 form json params为指针结构体
func (c *Controller) ReqFormBind(params interface{}) {
	if err := c.BindForm(params); err != nil {
		c.INFO(status.ParamBindFail, err.Error())
	}
	// TODO Auto Creator
	//if reflect.TypeOf(params).Kind() == reflect.Struct {
	//	fmt.Println("u kind is struct")
	//}
}

// ReqJsonBind 请求参数绑定 form json params为指针结构体
func (c *Controller) ReqJsonBind(params interface{}) {
	if err := c.BindJSON(params); err != nil {
		c.INFO(status.ParamBindFail, err.Error())
	}
	// TODO Auto Creator
	//if reflect.TypeOf(params).Kind() == reflect.Struct {
	//	fmt.Println("u kind is struct")
	//}
}

// ParamsValid 请求参数验证
func (c *Controller) ParamsValid(params interface{}) {
	v := validate.Struct(params)
	if !v.Validate() {
		c.INFO(status.ParamValidFail, v.Errors.One())
	}
}

func (c *Controller) CurrentUrl() string {
	var ctx context.Context
	return ctx.Request.RequestURI
}
