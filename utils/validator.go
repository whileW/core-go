package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 校验方法 接收两个参数  入参实例，规则map
func Verify(st interface{}) (err error) {
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)
	kd := val.Kind()
	if kd != reflect.Struct {
		return errors.New("只接受struct类型的参数")
	}
	return verifyStruct(typ,&val)
}

func verify(val *reflect.Value,sfield *reflect.StructField) error {
	kd := val.Kind() // 获取到st对应的类别
	switch kd {
	case reflect.Struct:
		return verifyStruct(sfield.Type,val)
	case reflect.Array,reflect.Slice:
		return verifyArraySlice(sfield,val)
	default:
		return verifyFieldV2(val,sfield)
	}
	return nil
}
//todo 报错信息优化 field下的第i个值不可为空
func verifyArraySlice(typ *reflect.StructField,val *reflect.Value) error {
	for i := 0; i < val.Len(); i++ {
		val_t := val.Index(i)
		if err := verify(&val_t,typ);err != nil {
			return err
		}
	}
	return nil
}
func verifyStruct(typ reflect.Type,val *reflect.Value) error {
	num := val.NumField()
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		val_t := val.Field(i)
		typ_t := typ.Field(i)
		if err := verify(&val_t,&typ_t);err != nil{
			return err
		}
	}
	return nil
}
func verifyFieldV2(val *reflect.Value,sfield *reflect.StructField) error {
	kd := val.Kind()
	name := getFieldName(sfield)
	verRol := strings.Split(sfield.Tag.Get("v"),";")
	if len(verRol) <= 0 || verRol[0] == "" {
		return nil
	}
	switch kd {
	case reflect.String:
		return verifyFieldString(val,name,verRol)
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return verifyFieldInt(val,name,verRol)
	default:
		return errors.New("错误的结构类型")
	}
	return nil
}
//notEmpty	不为空
//eq		值等于
//ne		值不等于
//lt  		长度小于
//le		长度小于等于
//ge		长度大于等于
//gt		长度大于
//phone		值为手机号
//email 	值为邮箱
//regexp	正则校验
func verifyFieldString(val *reflect.Value,name string,verRol []string) error {
	for _,vr := range verRol {
		vrs := strings.Split(vr,"=")
		switch vrs[0] {
		case "notEmpty":	//非空
			if val.String() == "" {
				return errors.New(fmt.Sprintf("[%s]不可为空",name))
			}
		case "eq":
			if val.String() != vrs[1] {
				return errors.New(fmt.Sprintf("[%s]的值必须为：%s",name,vrs[1]))
			}
		case "ne":
			if val.String() == vrs[1] {
				return errors.New(fmt.Sprintf("[%s]的值不可以等于：%s",name,vrs[1]))
			}
		case "lt":	//长度小于
			c,_ := strconv.Atoi(vrs[1])
			if len(val.String()) >= c {
				return errors.New(fmt.Sprintf("[%s]的长度必须小于：%d",name,c))
			}
		case "le":	//小于等于
			c,_ := strconv.Atoi(vrs[1])
			if len(val.String()) > c {
				return errors.New(fmt.Sprintf("[%s]的长度必须小于等于：%d",name,c))
			}
		case "ge":	//大于等于
			c,_ := strconv.Atoi(vrs[1])
			if len(val.String()) < c {
				return errors.New(fmt.Sprintf("[%s]的长度必须大于等于：%d",name,c))
			}
		case "gt":	//大于
			c,_ := strconv.Atoi(vrs[1])
			if len(val.String()) <= c {
				return errors.New(fmt.Sprintf("[%s]的长度必须大于：%d",name,c))
			}
		case "phone":
			if !MatchPhone(val.String()) {
				return errors.New(fmt.Sprintf("[%s]的格式错误，手机号格式校验失败",name))
			}
		case "email":
			if !MatchEmail(val.String()) {
				return errors.New(fmt.Sprintf("[%s]的格式错误，邮箱格式校验失败",name))
			}
		case "regexp":
			if !MatchString(vrs[1],val.String()) {
				return errors.New(fmt.Sprintf("[%s]的格式错误，正则校验失败",name))
			}
		}
	}
	return nil
}
func verifyFieldInt(val *reflect.Value,name string,verRol []string) error {
	for _,vr := range verRol {
		vrs := strings.Split(vr,"=")
		switch vrs[0] {
		case "notEmpty":	//非空
			if val.Int() == 0 {
				return errors.New(fmt.Sprintf("[%s]不可为0",name))
			}
		case "eq":
			if val.String() != vrs[1] {
				return errors.New(fmt.Sprintf("[%s]的值必须为：%s",name,vrs[1]))
			}
		case "ne":
			if val.String() == vrs[1] {
				return errors.New(fmt.Sprintf("[%s]的值不可以等于：%s",name,vrs[1]))
			}
		case "lt":	//小于
			c,_ := strconv.ParseInt(vrs[1],10,64)
			if val.Int() >= c {
				return errors.New(fmt.Sprintf("[%s]必须小于：%d",name,c))
			}
		case "le":	//小于等于
			c,_ := strconv.ParseInt(vrs[1],10,64)
			if val.Int() > c {
				return errors.New(fmt.Sprintf("[%s]必须小于等于：%d",name,c))
			}
		case "ge":	//大于等于
			c,_ := strconv.ParseInt(vrs[1],10,64)
			if val.Int() < c {
				return errors.New(fmt.Sprintf("[%s]必须大于等于：%d",name,c))
			}
		case "gt":	//大于
			c,_ := strconv.Atoi(vrs[1])
			if len(val.String()) <= c {
				return errors.New(fmt.Sprintf("[%s]的长度必须大于：%d",name,c))
			}
		case "phone":
			if !MatchPhone(val.String()) {
				return errors.New(fmt.Sprintf("[%s]的格式错误，手机号格式校验失败",name))
			}
		case "email":
			if !MatchEmail(val.String()) {
				return errors.New(fmt.Sprintf("[%s]的格式错误，邮箱格式校验失败",name))
			}
		case "regexp":
			if !MatchString(vrs[1],val.String()) {
				return errors.New(fmt.Sprintf("[%s]的格式错误，正则校验失败",name))
			}
		}
	}
	return nil
}
func getFieldName(sfield *reflect.StructField) string {
	name := sfield.Tag.Get("name")
	if name == "" {
		name = sfield.Tag.Get("json")
		if name == "" {
			name = sfield.Name
		}
	}
	return name
}