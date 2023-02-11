package bee

import (
	"github.com/beego/beego/v2/client/orm"
)

type Service struct{}

func (s Service) Query(table interface{}) (query orm.QuerySeter) {
	query = orm.NewOrm().QueryTable(table)
	return query
}

func (s Service) Order(query orm.QuerySeter, field string, mode string) orm.QuerySeter {
	switch mode {
	case "-":
		query = query.OrderBy(mode + field)
	default:
		query = query.OrderBy(field)
	}
	return query
}
