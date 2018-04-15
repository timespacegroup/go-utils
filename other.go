package tsgutils

import (
	"fmt"
	"go/types"
)

/*
 other utils
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

/*
  ternary operator, replace other language: a == b ? c : d
*/
func IIIInterfaceOperator(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func IIITypeOperator(condition bool, trueValue, falseValue types.Type) types.Type {
	if condition {
		return trueValue
	}
	return falseValue
}

func FmtPrintln(info interface{}) {
	if info != nil {
		switch t := info.(type) {
		case struct{}:
			fmt.Println(StructToJson(t))
		default:
			fmt.Println(t)
		}
	} else {
		fmt.Println(nil)
	}
}
