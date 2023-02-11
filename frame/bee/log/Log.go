package log

import (
	"github.com/beego/beego/v2/core/logs"
)

type Msg string

type FormatterOperate interface {
	Format(lm *Msg) string
}

var LOG = logs.NewLogger()

func init() {
	err := LOG.SetLogger(
		logs.AdapterFile,
		`{"filename":"runtime/logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	if err != nil {
		return
	}
	// 异步输出 提升性能
	LOG.Async()
}

// RecordRequest 请求日志记录
func RecordRequest(user string, role string, url string, method string, ip string, params interface{}) {
	f := &logs.PatternLogFormatter{
		Pattern:    "%F:%n|%w%t>> %m",
		WhenFormat: "2006-01-02 15:04:05",
	}
	logs.RegisterFormatter("pattern", f)
	_ = logs.SetGlobalFormatter("pattern")

	LOG.Info("user: %s | role: %s | ip: %s | method: %s | path: %s | params: %s ", user, role, ip, method, url, params)
}

func RecordError(file string, method string) {

}
