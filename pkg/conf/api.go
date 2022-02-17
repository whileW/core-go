package conf

import (
	"fmt"
	"github.com/whileW/core-go/pkg/util/xcast"
	"strings"
)

const (
	defaultSearchKeyDelim = "."
)

// todo 检查key是否存在

func Get(key string) interface{} {
	return defaultConfiguration.settings.Get(key)
}
func GetString(key string) string {
	return defaultConfiguration.settings.GetString(key)
}
func GetStringd(key,dv string) string {
	return defaultConfiguration.settings.GetStringd(key,dv)
}
func GetBool(key string) bool {
	return defaultConfiguration.settings.GetBool(key)
}
func GetBoold(key string,dv bool) bool {
	return defaultConfiguration.settings.GetBoold(key,dv)
}
func GetInt(key string) int {
	return defaultConfiguration.settings.GetInt(key)
}
func GetIntd(key string,dv int) int {
	return defaultConfiguration.settings.GetIntd(key,dv)
}
func GetInt64(key string) int64 {
	return defaultConfiguration.settings.GetInt64(key)
}
func GetInt64d(key string,dv int64) int64 {
	return defaultConfiguration.settings.GetInt64d(key,dv)
}
func GetFloat64(key string) float64 {
	return defaultConfiguration.settings.GetFloat64(key)
}
func GetFloat64d(key string,dv float64) float64 {
	return defaultConfiguration.settings.GetFloat64d(key,dv)
}
func GetStringSlice(key string) []string {
	return defaultConfiguration.settings.GetStringSlice(key)
}
func GetChildd(key string) *Settings {
	return defaultConfiguration.settings.GetChildd(key)
}
func PrintSourceData()  {
	defaultConfiguration.PrintSourceData()
}

func (c *Configuration)find(key string) interface{} {
	keys := strings.Split(key,defaultSearchKeyDelim)
	return c.settings.find(keys)
}
func (c *Configuration)Get(key string) interface{} {
	return c.settings.Get(key)
}
func (c *Configuration)GetString(key string) string {
	return c.settings.GetString(key)
}
func (c *Configuration)GetStringd(key,dv string) string {
	return c.settings.GetStringd(key,dv)
}
func (c *Configuration)GetBool(key string) bool {
	return c.settings.GetBool(key)
}
func (c *Configuration)GetBoold(key string,dv bool) bool {
	return c.settings.GetBoold(key,dv)
}
func (c *Configuration)GetInt(key string) int {
	return c.settings.GetInt(key)
}
func (c *Configuration)GetIntd(key string,dv int) int {
	return c.settings.GetIntd(key,dv)
}
func (c *Configuration)GetInt64(key string) int64 {
	return c.settings.GetInt64(key)
}
func (c *Configuration)GetInt64d(key string,dv int64) int64 {
	return c.settings.GetInt64d(key,dv)
}
func (c *Configuration)GetFloat64(key string) float64 {
	return c.settings.GetFloat64(key)
}
func (c *Configuration)GetFloat64d(key string,dv float64) float64 {
	return c.settings.GetFloat64d(key,dv)
}
func (c *Configuration)GetStringSlice(key string) []string {
	return c.settings.GetStringSlice(key)
}
func (c *Configuration)GetChildd(key string) *Settings {
	return c.settings.GetChildd(key)
}
func (c *Configuration)PrintSourceData()  {
	fmt.Printf("%+v\n", c.source_data)
	fmt.Println()
}

func (s *Settings)find(keys []string) interface{} {
	if v, ok := (*s)[keys[0]]; ok {
		if len(keys) > 1 {
			if v,ok := v.Value.(*Settings);ok {
				return v.find(keys[1:])
			}else {
				return nil
			}
		}
		return v.Value
	} else {
		return nil
	}
}
func (s *Settings)Get(key string) interface{} {
	keys := strings.Split(key,defaultSearchKeyDelim)
	return s.find(keys)
}
func (s *Settings)GetString(key string) string {
	return xcast.ToString(s.Get(key))
}
func (s *Settings)GetStringd(key,dv string) string {
	v := s.GetString(key)
	if v == "" {
		return dv
	}
	return v
}
func (s *Settings)GetBool(key string) bool {
	return xcast.ToBool(s.Get(key))
}
func (s *Settings)GetBoold(key string,dv bool) bool {
	if v := s.Get(key);v == nil {
		return dv
	}else {
		return xcast.ToBool(dv)
	}
}
func (s *Settings)GetInt(key string) int {
	return xcast.ToInt(s.Get(key))
}
func (s *Settings)GetIntd(key string,dv int) int {
	if v := s.Get(key);v == nil {
		return dv
	}else {
		return xcast.ToInt(v)
	}
}
func (s *Settings)GetInt64(key string) int64 {
	return xcast.ToInt64(s.Get(key))
}
func (s *Settings)GetInt64d(key string,dv int64) int64 {
	if v := s.Get(key);v == nil {
		return dv
	}else {
		return xcast.ToInt64(v)
	}
}
func (s *Settings)GetFloat64(key string) float64 {
	return xcast.ToFloat64(s.Get(key))
}
func (s *Settings)GetFloat64d(key string,dv float64) float64 {
	if v := s.Get(key);v == nil {
		return dv
	}else {
		return xcast.ToFloat64(v)
	}
}
func (s *Settings)GetStringSlice(key string) []string {
	return xcast.ToStringSlice(key)
}
func (s *Settings)GetChildd(key string) *Settings {
	if v := s.Get(key);v == nil {
		return &Settings{}
	}else {
		if v,ok := v.(*Settings);ok {
			return v
		}else {
			return &Settings{}
		}
	}
}