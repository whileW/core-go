package xcodec

import "errors"

// 编解码库

var(
	NoFindXCoderERR 	=	errors.New("没有找到对应得编解码器")
)

var(
	xcodecBuilders		=	make(map[string]Xcodec)
)

type Xcodec interface {
	Decode(data []byte, i interface{}) error
	Encode(i interface{}) ([]byte, error)
}

func Decode(typ string,data []byte,i interface{}) error {
	if v,exist := xcodecBuilders[typ];exist {
		return v.Decode(data,i)
	}
	return NoFindXCoderERR
}
func Encode(typ string,i interface{}) ([]byte, error) {
	if v,exist := xcodecBuilders[typ];exist {
		return v.Encode(i)
	}
	return nil,NoFindXCoderERR
}
func Register(name string, creator Xcodec) {
	xcodecBuilders[name] = creator
}