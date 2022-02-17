package xcontext

import (
	"context"
	"unsafe"
)

type valueCtx struct {
	context.Context
	key, val interface{}
}

type iface struct {
	itab, data uintptr
}
func GetValueCtxKV(ctx context.Context) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	getValueCtxKV(ctx, m)
	return m
}

func getValueCtxKV(ctx context.Context, m map[interface{}]interface{}) {
	ictx := *(*iface)(unsafe.Pointer(&ctx))
	if ictx.data == 0 {
		return
	}

	valCtx := (*valueCtx)(unsafe.Pointer(ictx.data))
	if valCtx != nil && valCtx.key != nil && valCtx.val != nil {
		m[valCtx.key] = valCtx.val
	}
	getValueCtxKV(valCtx.Context, m)
}