package exception

import (
	"HaloAdmin/halo/frame/bee"
)

// ErrorController 异常控制器
type ErrorController struct {
	bee.Controller
}

// #### Exception 异常处理 beego默认处理 401、403、404、500、503

// Error401 No Authorization
func (c *ErrorController) Error401() {
	c.INFO(401, "no authorization")
}

// Error403 Forbidden
func (c *ErrorController) Error403() {
	c.INFO(403, "forbidden")
}

// Error404 Not Found
func (c *ErrorController) Error404() {
	c.INFO(404, "not found")
}

// Error500 Server Error
func (c *ErrorController) Error500() {
	c.INFO(500, "system error")
}

// Error501 Not Implemented
func (c *ErrorController) Error501() {
	c.INFO(501, "not implemented")
}

// ErrorDB DB异常
func (c *ErrorController) ErrorDB() {
	c.INFO(600, "database error")
}
