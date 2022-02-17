package conf

import (
	"errors"
	"fmt"
	"github.com/whileW/core-go/pkg/conf/datasource"
	"sync"
)

// env < flag < datasource
type Configuration struct {
	settings		Settings
	source_data 	map[string]interface{}

	initOnce 		sync.Once
	//listen			map[string]chan Setting
}

type Settings map[string]Setting
type Setting struct {
	Key 			string
	Value 			interface{}
	//change_chan 	[]chan int
}

var defaultConfiguration = &Configuration{}

func Initialize() (err error) {
	defaultConfiguration.initOnce.Do(func() {
		fmt.Println("开始初始化 conf ...")
		d,e := datasource.NewDataSource()
		if e != nil {
			err = errors.New(fmt.Sprintf("【conf】initialize config error: %v", e))
			return
		}
		if e := defaultConfiguration.loadFromDataSource(d);e != nil{
			err = errors.New(fmt.Sprintf("【conf】%v",e))
			return
		}
	})
	return err
}

func (c *Configuration)loadFromDataSource(ds datasource.DataSource) error {
	conf,err := ds.ReadConfig()
	if err != nil {
		return errors.New(fmt.Sprintf("datasource read config failed: %v", err))
	}
	return c.apply(conf)
}
func (c *Configuration)apply(conf map[string]interface{}) error {
	temp := &Settings{}
	c.analysisSetting(&c.settings,temp,conf,0)
	c.settings = *temp
	c.source_data = conf
	fmt.Println("读取到配置：",conf)
	return nil
}
func (c *Configuration)analysisSetting(source *Settings,temp *Settings,s map[string]interface{},h int) bool {
	is_change := false
	for k,v := range s {
		//变量初始化
		var source_s *Setting
		if t,ok := (*source)[k];ok {
			source_s = &t
		}else {
			source_s = &Setting{
				Key:k,
				Value:&Settings{},
			}
		}
		temp_s := &Setting{
			Key:k,
		}

		//配置文件赋值
		if v != nil {
			if d,ok:=v.(map[string]interface{});ok {
				ts := &Settings{}
				temp_s.Value = ts
				if ic := c.analysisSetting(source_s.Value.(*Settings),ts,d,h+1);ic{
					is_change = ic
				}
			}else {
				if v != source_s.Value {
					is_change = true
				}
				temp_s.Value = v
			}
		}

		(*temp)[k] = *temp_s
	}
	return is_change
}