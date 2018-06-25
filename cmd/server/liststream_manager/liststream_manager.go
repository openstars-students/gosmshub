package liststream_manager
import (
	pb "gosmshub/proto"
	"fmt"
	"log"
	"time"
	"gosmshub/cmd/server/writelogfile"
	"gosmshub/cmd/server/common"
)

var  Liststream map[string] NameHomeNetwork
type NameHomeNetwork struct {
	NameHomeNetwork map[pb.SMSBlockChainHub_BiDiMessageServer]*StreamInfo
}

type StreamInfo struct {
	FullNumber string
	Stream pb.SMSBlockChainHub_BiDiMessageServer
	//sdt vua gui tin nhan thi token = tru
	StatusSending bool
	NameHomeNetwork string
	DailyLimit int32
	Index int32
	TotalLimit int32
	RewardAddress string
}

//interface theo tieu chi
type SelectOn interface {
	ChoseConnect(clients map[string]*StreamInfo) string
}
//chon theo nha mang
type SelectOnNameHomeNetwork struct {
}
//chon theo so dien thoai mac dinh biet truoc
type SelectOnDefaultNumber struct {
	defaultNumber string
}
var Dem = 0
func ListenToThriftServer(CommandSendMessage *pb.S2CMesssage) (bool, error){
	fmt.Println("nhan message tu server Thrift")
	selectOnLabelName := SelectOnNameHomeNetwork{}
	selectOnLabelName.ChoseConnect(CommandSendMessage)
	return false, nil
}
//gui mess theo nha mang
func(Select SelectOnNameHomeNetwork)ChoseConnect(CommandSendMessage *pb.S2CMesssage) (bool, error){
	fmt.Println("mess from server thrift: ", CommandSendMessage.GetMessage())
	nameHomeNetwork := common.ClassifyNumber(CommandSendMessage.GetMessage().GetToNumber())
	fmt.Println("labelName: ",nameHomeNetwork )

	stream,nameHomeNetworkSelected := ChoseConnectOnCircle(nameHomeNetwork)
	fmt.Println("chi co 1 connect: ",nameHomeNetworkSelected)
	if stream  != nil {
		fmt.Println("sending ...")
		stream.Send(CommandSendMessage)
		ticker := time.NewTicker(time.Second)
		started := time.Now()
		indexBefore := Liststream[nameHomeNetworkSelected].NameHomeNetwork[stream].Index
		go func() {
			for now := range ticker.C {
				//	dat timeout la 3s
				if now.Sub(started) > time.Second*4 {
					if Liststream[nameHomeNetwork].NameHomeNetwork[stream] != nil {
						{
							indexAfter := Liststream[nameHomeNetwork].NameHomeNetwork[stream].Index
							fmt.Println("indexBefore: ", indexBefore)
							if indexBefore == indexAfter {
								fmt.Println("time out")
								fmt.Println("Loi duong truyen voi client")
								fmt.Println("delete ", )
								fmt.Println("SUM before delete", TotalConnect())
								delete(Liststream[nameHomeNetwork].NameHomeNetwork, stream)
								fmt.Println("SUM after delete", TotalConnect())

								stream,nameHomeNetwork := ChoseConnectOnCircle(nameHomeNetwork)
								if stream  != nil {
									fmt.Println("gui lai ...")
									stream.Send(CommandSendMessage)
									writelogfile.LogAppSendMsg(Liststream[nameHomeNetwork].NameHomeNetwork[stream].FullNumber,CommandSendMessage)
									}
							}else {
								writelogfile.LogAppSendMsg(Liststream[nameHomeNetwork].NameHomeNetwork[stream].FullNumber,CommandSendMessage)
							}
						}
					}
					ticker.Stop()
					return
				}
			}
		}()

	}
	return false, nil
}

func ChoseConnectOnCircle(nameHomeNetwork string) (pb.SMSBlockChainHub_BiDiMessageServer, string){
//sum = ????
//moi 1 nha mang chi co 1 connect,
	for k, v := range Liststream {
		if len(Liststream[k].NameHomeNetwork)==1 {
			for _, eachStream  := range v. NameHomeNetwork {
					if eachStream.Stream != nil && eachStream.Index <= eachStream.DailyLimit{
						eachStream.StatusSending = false
						}
			}
		}
	}
//neu tat ca chi co 1 connect ==>> luon chon connect do
	sum := TotalConnect()
	if sum ==1 {
		for _, v := range Liststream {
			for _, eachStream  := range v. NameHomeNetwork {
				if eachStream.Stream != nil && eachStream.Index <= eachStream.DailyLimit {
					fmt.Println("sum = 1, chi co 1 connect")
					eachStream.StatusSending = true
					return eachStream.Stream, common.ClassifyNumber(eachStream.FullNumber)
					}
		}
		}
	}
//check cac connect thuoc nha mang NameHomeNetwork do deu da gui tin nhan trong 1 chu ky??
	check := false
	for k, v := range Liststream {
		if k == nameHomeNetwork {
			for _, eachStream  := range v. NameHomeNetwork {
				if eachStream.StatusSending == false && eachStream.Index <= eachStream.DailyLimit{
					eachStream.StatusSending = true
					return eachStream.Stream, common.ClassifyNumber(eachStream.FullNumber)
				}
				if eachStream.Stream != nil{check = true}
				}
			}
	}
//set StatusSending = false cho tat cac cac connect cua NameHomeNetwork
	if check {
		fmt.Println("co nhieu nha mang do")
		for k, v := range Liststream {
			if k == nameHomeNetwork {
				for _, eachStream  := range v. NameHomeNetwork{
				fmt.Println("chuyen tat ca trang thai thanh false")
					eachStream.StatusSending = false
				}
			}
		}
		for k, v := range Liststream {
			if k == nameHomeNetwork {
				for _, eachStream  := range v. NameHomeNetwork {
					if eachStream.StatusSending == false && eachStream.Index <= eachStream.DailyLimit{
						fmt.Println("chon duoc 1 connect trong n connect nha mang: ", nameHomeNetwork)
						eachStream.StatusSending = true
						return eachStream.Stream, common.ClassifyNumber(eachStream.FullNumber)
					}
				}
			}
		}
	}else{
		fmt.Println("khong co nha mang nao, chon 1 connect cua nha mang khac")
		for _, v := range Liststream {
			for _, eachStream  := range v. NameHomeNetwork {
				if eachStream.StatusSending == false && eachStream.Index <= eachStream.DailyLimit{
					eachStream.StatusSending = true
					return eachStream.Stream, common.ClassifyNumber(eachStream.FullNumber)
				}
			}
		}
	}
	return nil,""
}

func ListenToClient(nameHomeNetwork string, stream pb.SMSBlockChainHub_BiDiMessageServer){
	for {
		a, _ := stream.Recv()
		if a != nil {
		//app nhan tin nhan tu client
			if a.GetIncomingMessage().GetContent() != ""{
				log.Println("[listenToClient]:APP NHAN TIN NHAN TU CLIENT: ")
				if Liststream[nameHomeNetwork].NameHomeNetwork != nil && Liststream[nameHomeNetwork].NameHomeNetwork[stream] != nil{
					fmt.Println("LogInCommingMsg: ")
					writelogfile.LogInCommingMsg(a)
				}
			} else if a.GetSuccess() == true {
				log.Println("[listenToClient]: Phan hoi tu client: ")
				if Liststream[nameHomeNetwork].NameHomeNetwork != nil && Liststream[nameHomeNetwork].NameHomeNetwork[stream] != nil{
					Liststream[nameHomeNetwork].NameHomeNetwork[stream].Index ++
					fmt.Println("indexAfter: ", Liststream[nameHomeNetwork].NameHomeNetwork[stream].Index)
				}
			//app response to server
			}
		} else {
			if Liststream[nameHomeNetwork].NameHomeNetwork != nil && Liststream[nameHomeNetwork].NameHomeNetwork[stream] != nil{
				log.Println("[listenToClient]: MAT KET NOI vs Number: ", Liststream[nameHomeNetwork].NameHomeNetwork[stream].FullNumber)
				fmt.Println("delete connect vs number: ")
				delete(Liststream[nameHomeNetwork].NameHomeNetwork, stream)
				return
			}
		}
	}
}

//neu khong tach thanh 2 ham rpc, phai co unitStream