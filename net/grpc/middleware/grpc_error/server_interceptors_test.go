package grpc_error

import (
	"errors"
	"github.com/azdbaaaaaa/util/xutil/xerror"
	"log"
	"testing"
)

func TestErrorWrapper(t *testing.T) {
	err := xerror.ErrParamInvalid
	if _, ok := err.(xerror.Error); ok {
		log.Println("is xerror")
	} else {
		log.Println("is not xerror")
	}
	err2 := errors.New(err.Error())
	if _, ok := err2.(xerror.Error); ok {
		log.Println("is xerror")
	} else {
		log.Println("is not xerror")
	}
}