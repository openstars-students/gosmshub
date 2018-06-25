package common

import (
	"gosmshub/appconfig"
	"fmt"
)

//dau vao la sdt, tra ve nha mang tuong ung
func ClassifyNumber(number string ) string{
	var t string
	if len(number) ==10{
		t = number[0:3]
	}
	if len(number) ==11{
		t = number[0:4]
	}
	switch  {
	case appconfig.Viettel[t] != "":
		fmt.Println("Viettel")
		return "Viettel"
	case appconfig.Mobifone[t] != "":
		fmt.Println("Mobifone")
		return "Mobifone"
	case appconfig.Vinaphone[t] != "":
		fmt.Println("Vinaphone")
		return "Vinaphone"
	default:
		fmt.Println("Unknow")
		return "Unknown"
	}
}