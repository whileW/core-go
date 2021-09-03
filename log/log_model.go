package log

import (
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
	latency := over.Sub(start)
	l.WithKV("duration",latency)
	return l
}
func (l *Loger)WithError(err error) *Loger {
	l.WithKV("err",err)
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