package liststream_manager

import (
	pb "gosmshub/proto"
)
func init() {
	Liststream = make(map[string]NameHomeNetwork)

	Viettel := NameHomeNetwork{}
	Vinaphone := NameHomeNetwork{}
	Mobiphone := NameHomeNetwork{}
	Unknown := NameHomeNetwork{}

	Viettel.NameHomeNetwork = make(map[pb.SMSBlockChainHub_BiDiMessageServer]*StreamInfo)
	Liststream["Viettel"] = Viettel
	Vinaphone.NameHomeNetwork = make(map[pb.SMSBlockChainHub_BiDiMessageServer]*StreamInfo)
	Liststream["Vinaphone"] = Vinaphone
	Mobiphone.NameHomeNetwork = make(map[pb.SMSBlockChainHub_BiDiMessageServer]*StreamInfo)
	Liststream["Mobifone"] = Mobiphone
	Unknown.NameHomeNetwork = make(map[pb.SMSBlockChainHub_BiDiMessageServer]*StreamInfo)
	Liststream["Unknown"] = Unknown
}
func TotalConnect() int {
	sum :=0
	for k, _ := range Liststream {
		sum = sum + len(Liststream[k].NameHomeNetwork)
	}
	return sum
}