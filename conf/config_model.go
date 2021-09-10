package conf

import (
	"fmt"
	"reflect"
	"strconv"
)

//todo 手动增加配置   conf.GetConf().Add()|.AddChild()
//todo 修改配置反写到配置文件
//todo 配置来源，from env\file\system
var conf Config

func GetConf() *Config {
	return &conf
}

//总配置
type Config struct {
	SysSetting			sysSetting
	Setting				Settings
}

//系统设置
type sysSetting struct {
	//环境 - 默认debug
	//dev	-	开发
	//debug - 测试
	//release - 正式
	Env 			string
	//http 监听端口地址 - 默认8080
	HttpAddr		string
	//rpc 监听端口地址 - 默认30010
	RpcAddr 		string
	//本机ip
	Host 			string
	//系统名称
	SystemName 		string
	//从哪里读取配置		- 默认file,可选值file、acm(阿里配置中心)
	ConfFrom		string
}

//设置默认环境
func (s *sysSetting)SetDefaultEnv() {
	if s.Env == "" {
		s.Env = "debug"
	}
}
//设置默认系统名称
func (s *sysSetting)SetDefaultSystemName() {
	if s.SystemName == "" {
		s.SystemName = "core"
	}
}
//设置默认http监听地址
func (s *sysSetting)SetDefaultHttpAddr() {
	if s.HttpAddr == "" {
		s.HttpAddr = "8080"
	}
}
//设置默认rpc监听地址
func (s *sysSetting)SetDefaultRpcAddr() {
	if s.HttpAddr == "" {
		s.HttpAddr = "30010"
	}
}
//设置默认host
func (s *sysSetting)SetDefaultHost()  {
	if s.Host == "" {
		s.Host = "127.0.0.1"
	}
}
//设置默认值
func (s *sysSetting)SetDefault()  {
	s.SetDefaultSystemName()
	s.SetDefaultEnv()
	s.SetDefaultHttpAddr()
	s.SetDefaultRpcAddr()
	s.SetDefaultHost()
}

//其他设置
type Settings map[string]Setting
type Setting struct {
	Key 			string
	Value 			interface{}
	change_chan 	[]chan int
}

func (s *Settings)Get(key string) interface{} {
	if v, ok := (*s)[key]; ok {
		return v.Value
	} else {
		panic("key not find")
	}
}
func (s *Settings)GetInt(key string) int {
	tv := s.Get(key)
	if v,ok := tv.(int);ok{
		return v
	}
	switch reflect.TypeOf(tv).String() {
	case "string":
		if v,err := strconv.Atoi(tv.(string));err != nil{
			panic(err)
		}else {
			return v
		}
	case "float32":
		return int(tv.(float32))
	case "float64":
		return int(tv.(float64))
	default:
		panic("conf.GetInt():err type to int:"+reflect.TypeOf(tv).String())
	}
}
func (s *Settings)GetString(key string) string {
	return s.Get(key).(string)
}
func (s *Settings)GetBool(key string) bool {
	return s.Get(key).(bool)
}
func (s *Settings)GetChild(key string) *Settings {
	return s.Get(key).(*Settings)
}

func (s *Settings)Getd(key string,d interface{})interface{} {
	if s == nil {
		return d
	}
	if v, ok := (*s)[key]; ok {
		return v.Value
	} else {
		return d
	}
}
func (s *Settings)GetIntd(key string,d int) int {
	tv := s.Get(key)
	var v int
	var ok bool
	if v,ok = tv.(int);!ok{
		switch reflect.TypeOf(tv).String() {
		case "string":
			if sv,err := strconv.Atoi(tv.(string));err != nil{
				return d
			}else {
				v = sv
			}
		case "float32":
			v = int(tv.(float32))
		case "float64":
			v =  int(tv.(float64))
		default:
			return d
		}
	}
	if v == 0 {
		 return d
	}
	return v
}
func (s *Settings)GetStringd(key string,d string) string {
	tv := s.Getd(key,d)
	var v string
	var ok bool
	if v,ok = tv.(string);!ok {
		switch reflect.TypeOf(tv).String() {
		case "int":
			v = strconv.Itoa(tv.(int))
		}
	}
	if v == "" {
		return d
	}
	return v
}
func (s *Settings)GetBoold(key string,d bool) bool {
	return s.Getd(key,d).(bool)
}
func (s *Settings)GetChildd(key string) *Settings {
	v := s.Getd(key,nil)
	if v == nil {
		return &Settings{}
	}
	return v.(*Settings)
}

func (s *Settings)Getd_c(key string,d interface{})(interface{},chan int) {
	var value interface{}
	change_chan := make(chan int)
	if v, ok := (*s)[key]; ok {
		value = v.Value
		if change_chan == nil || len(v.change_chan) <= 0 {
			v.change_chan = []chan int{}
		}
		v.change_chan = append(v.change_chan, change_chan)
		(*s)[key] = v
	} else {
		value = d
	}
	return value,change_chan
}
func (s *Settings)GetIntd_c(key string,d int) (int,chan int) {
	v,ch := s.Getd_c(key,d)
	return v.(int),ch
}
func (s *Settings)GetStringd_c(key string,d string) (string,chan int) {
	v,ch := s.Getd_c(key,d)
	return v.(string),ch
}
func (s *Settings)GetBoold_c(key string,d bool) (bool,chan int) {
	v,ch := s.Getd_c(key,d)
	return v.(bool),ch
}
func (s *Settings)GetChildd_c(key string) (*Settings,chan int) {
	v,ch := s.Getd_c(key,nil)
	if v == nil {
		return nil,ch
	}
	return v.(*Settings),ch
}

func (s *Setting)send_change()  {
	if s.change_chan != nil {
		for _,t := range s.change_chan {
			fmt.Println("配置修改:",s.Key)
			t<-1
		}
	}
}