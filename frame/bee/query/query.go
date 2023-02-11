package query

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"strings"
)

func List[S comparable](insModel S) []S {
	q := orm.NewOrm().QueryTable(insModel)
	list := make([]S, 0)
	_, _ = q.All(&list)
	return list
}

func UniqueFilter[P comparable, T comparable](params P, insModel T) error {
	o := orm.NewOrm()
	tag := "haloValid"
	t := reflect.TypeOf(params)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if strings.Contains(field.Tag.Get(tag), "unique") {
			qs := o.QueryTable(insModel)
			qs = qs.Filter(strings.ToLower(field.Name), reflect.ValueOf(&params).
				Elem().FieldByName(strings.Title(field.Name)).String())
			err := qs.One(&insModel)
			if err != nil {
				println(err.Error())
			}
			if reflect.ValueOf(&insModel).Elem().FieldByName(field.Name).String() == reflect.ValueOf(&params).Elem().
				FieldByName(field.Name).String() {
				return errors.New(field.Name + "已存在")
			}
		}
	}
	return nil
}

// Where whereField: key = condition, value = [ param, value ]
func Where[T string | int](query orm.QuerySeter, whereField map[string]map[string]T) orm.QuerySeter {
	for kc, vc := range whereField {
		switch kc {
		case "=":
			for kvc, vvc := range vc {
				query = query.Filter(kvc, vvc)
			}
		case "like":
			for kvc, vvc := range vc {
				query = query.Filter(kvc+"__contains", vvc)
			}
		}
	}

	return query
}
