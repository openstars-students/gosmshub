package appconfig
import "fmt"
var (
	PathLogAppSendMsg string
	PathLogInCommingMsg string
	Viettel = make(map[string] string)
	Vinaphone = make(map[string] string)
	Mobifone = make(map[string] string)

	Default = make(map[string] string)

)

func InitConfig(){
	fmt.Println("[InitConfig]")
	PathLogAppSendMsg = "/home/AppSendMsg.txt"
	PathLogInCommingMsg = "/home/InCommingMsg.txt"

	Viettel["086"] = "086"
	Viettel["096"] = "096"
	Viettel["097"] = "097"
	Viettel["098"] = "098"
	Viettel["0163"] = "0163"
	Viettel["0164"] = "0164"
	Viettel["0165"] = "0165"
	Viettel["0166"] = "0166"
	Viettel["0167"] = "0167"
	Viettel["0168"] = "0168"
	Viettel["0169"] = "0169"

	Mobifone["090"] = "090"
	Mobifone["093"] = "093"
	Mobifone["0120"] = "0120"
	Mobifone["0121"] = "0121"
	Mobifone["0122"] = "0122"
	Mobifone["0126"] = "0126"
	Mobifone["0128"] = "0128"
	Mobifone["08966"] = "08966"

	Vinaphone["091"] = "091"
	Vinaphone["094"] = "094"
	Vinaphone["0123"] = "0123"
	Vinaphone["0124"] = "0124"
	Vinaphone["0125"] = "0125"
	Vinaphone["0127"] = "0127"
	Vinaphone["0129"] = "0129"
}