package HaloRequest

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func RequestBindJson(ctx *gin.Context, params interface{}) (err error) {
	err = ctx.ShouldBind(params)
	return err
}

func RequestBindForm(ctx *gin.Context, params interface{}) (err error) {
	err = ctx.Bind(params)
	return err
}

func RequestValidate(params interface{}) (err error) {
	v := validate.Struct(params)
	if !v.Validate() {
		return errors.New("验证bbbb错误，这里需要改正")
	}
	return
}

// Handle 注意GET请求取参方式
func Handle(ctx *gin.Context, params interface{}, contentType string) (err error) {
	switch contentType {
	case "form":
		err = ctx.ShouldBind(params)
		if err != nil {
			return
		}
	default:
		// form 表单
		err = ctx.Bind(params)
		if err != nil {
			return
		}
	}

	v := validate.Struct(params)
	if !v.Validate() {
		return errors.New("验证错误，这里需要改正")
	}
	return
}
