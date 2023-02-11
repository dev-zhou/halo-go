package _interface

type INode interface {
	GetId() int           // GetId 获取id
	GetTitle() string     // GetTitle 获取显示名字
	GetParentId() int     // GetParentId 获取父id
	GetData() interface{} // GetData 获取附加数据
	IsRoot() bool         // IsRoot 判断当前节点是否是顶层根节点
}
