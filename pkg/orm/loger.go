package orm

import (
	"context"
	"fmt"
	"github.com/whileW/core-go/pkg/log"
	"github.com/whileW/core-go/pkg/util/xcontext"
	"gorm.io/gorm/logger"
	"time"
)

type ormLoger struct {
	*log.Loger
}

func (l *ormLoger)LogMode(logger.LogLevel) logger.Interface {
	return l
}
func (l *ormLoger)Info(ctx context.Context,msg string,arg ...interface{}){
	l.Clone().Info(fmt.Sprintf(msg,arg...))
}
func (l *ormLoger)Warn(ctx context.Context,msg string,arg ...interface{}){
	l.Clone().Warn(fmt.Sprintf(msg,arg...))
}
func (l *ormLoger)Error(ctx context.Context,msg string,arg ...interface{}){
	l.Clone().Error(fmt.Sprintf(msg,arg...))
}
func (l *ormLoger)Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error){
	sql,rows := fc()
	arg := xcontext.GetValueCtxKV(ctx)
	l.Clone().WithDuration(begin).Info(fmt.Sprintf("[trace] %s %d",sql,rows),arg)
}
