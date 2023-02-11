package bee

import (
	"HaloAdmin/base/alias"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// Model for beego 模型基础
type Model struct {
	Creator   alias.Id  `orm:"description(数据权限)"`
	CreatedAt time.Time `orm:"description(创建时间);type(datetime);auto_now_add"`
	UpdatedAt time.Time `orm:"description(修改时间);type(datetime);auto_now"`
}

// CRUDer CRUD 接口
type CRUDer interface {
	Create(interface{}) (int64, error)
	Delete(interface{}, ...string) (int64, error)
	Update(interface{}, ...string) (int64, error)
	Reader(interface{}, ...string) (interface{}, error)
	ReaderCreate(interface{}, string, ...string) (bool, int64, error)
}

// var CRUD CRUDer

// Create 新增 obj 指针对象
func (m *Model) Create(obj interface{}) (row int64, err error) {
	o := orm.NewOrm()
	row, err = o.Insert(obj)
	println()
	println(obj, "SQL: ---- Create")
	if err != nil {
		return row, errors.New(err.Error())
	}
	return row, err
}

// Update 更新 obj 指针对象
func (m *Model) Update(obj interface{}, cols ...string) (row int64, err error) {
	o := orm.NewOrm()
	row, err = o.Update(obj, cols...)
	println()
	println(obj, "SQL: ---- Update")
	if err != nil {
		return row, errors.New(err.Error())
	}
	return row, err
}

// Delete 删除 obj 指针对象
func (m *Model) Delete(obj interface{}, cols ...string) (row int64, err error) {
	o := orm.NewOrm()
	row, err = o.Delete(obj, cols...)
	if err != nil {
		return row, errors.New(err.Error())
	}
	return row, err
}

// Reader 查找
func (m *Model) Reader(obj interface{}, cols ...string) (info interface{}, err error) {
	o := orm.NewOrm()
	err = o.Read(obj, cols...)
	println()
	println(&obj, "SQL: ---- Reader")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &obj, err
}

// ReaderCreate return 3 values if ok first value created id else get id
func (m *Model) ReaderCreate(obj interface{}, cols1 string, cols ...string) (created bool, row int64, err error) {
	o := orm.NewOrm()
	created, row, err = o.ReadOrCreate(obj, cols1, cols...)
	println()
	println(obj, "SQL: ---- ReadOrCreate")
	if err != nil {
		return false, row, errors.New(err.Error())
	}
	return created, row, err
}

// Where QuerySeter
func (m *Model) Where(query orm.QuerySeter, where string, like string, operate uint8) orm.QuerySeter {
	// 这里可以采用 ...where方式循环 但考虑到reflect包，暂时不进行封装
	return query
}

func (m *Model) QuerySeterOrderBy() {

}

type Id11 struct {
	Id int `orm:"description(自增主键);size(11);auto;pk;unique"`
}

type Name20 struct {
	Name string `orm:"description(名称字段);size(20);unique"`
}

type Code20 struct {
	Code string `orm:"description(编码字段);size(20);unique"`
}

type Note20 struct {
	Note string `orm:"description(备注字段);size(20)"`
}

type Key20 struct {
	Key string `orm:"description(键字段);size(20)"`
}

type Value20 struct {
	Value string `orm:"description(值字段);size(20)"`
}

type Order20 struct {
	Order string `orm:"description(订单字段);size(20)"`
}

type Type255 struct {
	Type int `orm:"description(类型字段)"`
}

type TinyType struct {
	Type uint8 `orm:"description(类型字段)"`
}

type TinyStatus struct {
	Status uint8 `orm:"description(状态字段);default(1)"`
}

type Status255 struct {
	Status alias.Tiny `orm:"description(状态字段);default(1)"`
}

// Deleted 软删除字段
type Deleted struct {
	DeletedAt time.Time `orm:"description(删除时间);type(datetime);default(0)"`
}
