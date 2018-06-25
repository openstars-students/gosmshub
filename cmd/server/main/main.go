package main

import (
	"fmt"
	"log"
	"net"
	pb "gosmshub/proto"
	"google.golang.org/grpc"
	"gosmshub/appconfig"
	"gosmshub/cmd/server/writelogfile"
	"gosmshub/cmd/server/liststream_manager"
	"gosmshub/cmd/server/thriftserver"
	"gosmshub/cmd/server/common"
	"github.com/claudiu/gocron"
	"strconv"
)

type SendSmsService struct {
}

func KeppConnect(){
	s := gocron.NewScheduler()
	s.Every(1).Day().Do(task)
	sc := s.Start() // keep the channel
	<-sc            // it will happens if the channel is closed
}
func task() {
	fmt.Println("keep connect client....")
}

func (svr *SendSmsService) BiDiMessage(stream pb.SMSBlockChainHub_BiDiMessageServer) error {
	//authentication client
	auth, _ := stream.Recv()
	appID := auth.GetAuth().GetAppID()
	timeStamp := auth.GetAuth().GetTimeStamp()
	authSignature := auth.GetAuth().GetAuthSignature()
	fmt.Println("authID: " + appID + " timeStamp: "  + timeStamp + " signature: " + authSignature)
	//	CommandSendSmS kieu du lieu con trong S2CMessage; choseCommand kieu du lieu oneof; message la bien
	var message pb.CommandSendSmS
	message.Content = "the first stream success"
	CommandSendSmS := &pb.S2CMesssage{
		ChoseCommand: &pb.S2CMesssage_Message{&message},
	}
	var mess *pb.S2CMesssage
	mess = CommandSendSmS
	stream.Send(mess)

	//register information client
	agentInfo, _ := stream.Recv()
	fullNumber := agentInfo.GetSmsAgent().GetFullNumber()
	dailyLimit_Int := agentInfo.GetSmsAgent().GetDailyLimit(); dailyLimit_String := strconv.Itoa(int(dailyLimit_Int))
	totalLimit_Int := agentInfo.GetSmsAgent().GetTotalLimit(); totalLimit_String := strconv.Itoa(int(totalLimit_Int))
	rewardAddress := agentInfo.GetSmsAgent().GetRewardAddress()
	fmt.Println("fullNumber: " + fullNumber + " dailyLimit: "  + dailyLimit_String+" totalLimit: " + totalLimit_String+ " rewardAddress:" + rewardAddress)
	message.Content = "the second stream success"
	CommandSendSmS = &pb.S2CMesssage{
		ChoseCommand: &pb.S2CMesssage_Message{&message},
	}
	mess = CommandSendSmS
	stream.Send(mess)

	nameHomeNetwork := common.ClassifyNumber(fullNumber)
	//khoi tao cau truc cua streamInfo
	streamInfo := &liststream_manager.StreamInfo{
		FullNumber:  fullNumber,
		Stream: stream,
		StatusSending :false,
		NameHomeNetwork:nameHomeNetwork,
		Index: 0,
		DailyLimit : dailyLimit_Int,
		TotalLimit : totalLimit_Int,
		RewardAddress : rewardAddress,

	}

//khoi tao
	liststream_manager.Liststream[nameHomeNetwork].NameHomeNetwork[stream] = &liststream_manager.StreamInfo{}
//gan
	liststream_manager.Liststream[nameHomeNetwork].NameHomeNetwork[stream] = streamInfo

	go liststream_manager.ListenToClient(nameHomeNetwork, stream)
	KeppConnect()
	log.Println("[Send]: ket thuc ham Send")
	return nil
}

func main() {
	appconfig.InitConfig()
	listen, err := net.Listen("tcp", appconfig.AddressServerGo)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterSMSBlockChainHubServer(s, &SendSmsService{})
	fmt.Println("Listening on the 127.0.0.1:10023")
	//create file to write log
	writelogfile.CreateFile(appconfig.PathLogInCommingMsg)
	writelogfile.CreateFile(appconfig.PathLogAppSendMsg)
	//run server
	go serverthrift.RunServerThrift()
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
