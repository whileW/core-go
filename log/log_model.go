package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Loger struct {
	args 		[]interface{}
}

func (l *Loger)New() *Loger {
	tl := *l
	return &tl
}
func (l *Loger)WithKV(args ...interface{}) *Loger {
	l.args = append(l.args, args...)
	return l
}
func (l *Loger)WithDuration(start time.Time) *Loger {
	over := time.Now()
	l.WithKV("duration",fmt.Sprint(over.Sub(start)))
	return l
}
func (l *Loger)WithError(err error) *Loger {
	l.WithKV("err",err)
	return l
}
func (l *Loger)WithModule(name string) *Loger {
	l.WithKV("module",name)
	return l
}
func (l *Loger)WithModule_ApiReq() *Loger {
	l.WithModule("api_req_log")
	return l
}
func (l *Loger)WithModule_Db() *Loger {
	l.WithModule("db_log")
	return l
}

func (l *Loger)WithReqId(c *gin.Context) *Loger {
	l.WithKV("req_id",c.MustGet("req_id").(string))
	return l
}
func (l *Loger)handWith(args ...interface{}) *logrus.Entry {
	if len(l.args) > 0 {
		args = append(args, l.args...)
	}
	var e *logrus.Entry
	for i:=0;i< len(args);  {
		var (
			k string
			v interface{}
		)
		if i+1 >= len(args) {
			k = "unkown"
			v = args[i]
			i++
		}else {
			if tk,ok := args[i].(string);ok {
				k = tk
				v = args[i+1]
				i = i+2
			}else {
				k = "unkown"+strconv.Itoa(i+1)
				v = args[i]
				i++
			}
		}
		if e == nil {
			e = logrus.WithField(k,v)
		}else {
			e = e.WithField(k,v)
		}
	}
	return e
}

func (l *Loger)Infow(msg string, keysAndValues ...interface{})  {
	e := l.handWith(keysAndValues...)
	if e != nil {
		e.Info(msg)
	}else {
		logrus.Info(msg)
	}
}
func (l *Loger)Info(args ...interface{})  {
	logrus.Info(args...)
}
//todo 增加堆栈信息
func (l *Loger)Errorw(msg string, keysAndValues ...interface{})  {
	e := l.handWith(keysAndValues...)
	if e != nil {
		e.Error(msg)
	}else {
		logrus.Error(msg)
	}
}
func (l *Loger)Error(args ...interface{})  {
	logrus.Error(args...)
}

func (l *Loger)Warnw(msg string, keysAndValues ...interface{})  {
	e := l.handWith(keysAndValues...)
	if e != nil {
		e.Warn(msg)
	}else {
		logrus.Warn(msg)
	}
}
func (l *Loger)Warn(args ...interface{})  {
	logrus.Warn(args...)
}