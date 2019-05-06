package models

type Code int32

const (
	SUCCESS Code = iota // value --> 0
	ERR_BIZ             // value --> 1
	ERR_DATA            // value --> 2
	ERR_AUTH            // value --> 3
	ERR_LOST            // value --> 4
	ERR_UNKNOW          // value --> 5
)

func (c Code) String() string  {
	switch c {
	case SUCCESS:
		return "操作成功"
	case ERR_BIZ:
		return "业务操作失败"
	case ERR_DATA:
		return "数据操作失败"
	case ERR_AUTH:
		return "权限操作失败"
	case ERR_LOST:
		return "操作不存在"
	case ERR_UNKNOW:
		return "未知错误200"
	default:
		return "未知错误"
	}
}

type Result struct {
	Code Code
	Msg string
	Data interface{}
}
