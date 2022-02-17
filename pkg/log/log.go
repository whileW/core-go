package log

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/whileW/core-go/pkg/log/loki"
	"github.com/whileW/core-go/pkg/system_variable"
	"github.com/whileW/core-go/pkg/util/xgo"
	"strconv"
	"time"
)

type Loger struct {
	args 		[]interface{}
}

func Initialize() error {
	fmt.Println("开始初始化 log ...")
	WithPubKV("app",system_variable.SystemName)
	WithPubKV("env",system_variable.Env)
	return xgo.SerialUntilError(
			loki.Initialize,
		)()
}

var default_loger = &Loger{}

func GetLoger() *Loger {
	return default_loger.Clone()
}
func WithPubKV(args ...interface{}) *Loger {
	return default_loger.WithKV(args...).Clone()
}
func WithKV(args ...interface{}) *Loger {
	return default_loger.Clone().WithKV(args...)
}
func WithDuration(start time.Time) *Loger {
	return default_loger.Clone().WithDuration(start)
}
func WithError(err error) *Loger {
	return default_loger.Clone().WithError(err)
}
func WithModule(name string) *Loger {
	return default_loger.Clone().WithModule(name)
}

func (l *Loger)Clone() *Loger {
	copy := &Loger{args:l.args}
	return copy
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

func (l *Loger)Info(msg string, keysAndValues ...interface{})  {
	e := l.handWith(keysAndValues...)
	if e != nil {
		e.Info(msg)
	}else {
		logrus.Info(msg)
	}
}
//todo 增加堆栈信息
func (l *Loger)Error(msg string, keysAndValues ...interface{})  {
	e := l.handWith(keysAndValues...)
	if e != nil {
		e.Error(msg)
	}else {
		logrus.Error(msg)
	}
}
func (l *Loger)Warn(msg string, keysAndValues ...interface{})  {
	e := l.handWith(keysAndValues...)
	if e != nil {
		e.Warn(msg)
	}else {
		logrus.Warn(msg)
	}
}


func (l *Loger)ValueCtx() context.Context {
	ctx := context.Background()
	return l.ValueCtxb(ctx)
}
func (l *Loger)ValueCtxb(ctx context.Context) context.Context {
	for k,v := range l.handWith().Data {
		ctx = context.WithValue(ctx,k,v)
	}
	return ctx
}