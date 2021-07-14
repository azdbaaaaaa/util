package error

import (
	"errors"
	"fmt"
	"github.com/azdbaaaaaa/util/xutil/xerror"
	"testing"
)

func TestName(t *testing.T) {
	err2 := errors.New("xxx")
	//err2 := xerror.New(1, 1, "", "")
	switch err2.(type) {
	case xerror.Error:
		fmt.Println("Interface")
	case error:
		fmt.Println("Struct")
	default:
		//fmt.Println(err2.Kind())
	}
}
