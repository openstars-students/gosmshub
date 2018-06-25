package main

import (
	"fmt"
	"log"
	pb "gosmshub/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gosmshub/appconfig"
	"github.com/claudiu/gocron"

)
func run(){

	var client_send pb.C2SMessage
	var incomingMess pb.IncomingMsg
	incomingMess.Content = "XXXXXXXXXXXXXXXXXXXXXX"
	incomingMess.FromNumber = "1111111111111111"

	client_send.IncomingMessage = &incomingMess
	Stream.Send(&client_send)

}
var Stream pb.SMSBlockChainHub_BiDiMessageClient

func send(client pb.SMSBlockChainHubClient) {

	stream,err := client.BiDiMessage(context.Background())
	if err != nil {
		log.Fatal("client  get stream error: ", err)
	}
	Stream = stream
	var authClient pb.AuthMessage
	authClient.AppID = "AppID_111 "
	authClient.AuthSignature = "AuthSignature_111"
	authClient.TimeStamp = "TimeStamp_111"
	theFirstStream_Send := &pb.C2SMessage{
		AuthClient: &pb.C2SMessage_Auth{&authClient},
	}

	stream.Send(theFirstStream_Send)
	theFirstStream_Recv,_ := stream.Recv()
	fmt.Println(theFirstStream_Recv)

	var agentInfo pb.SmsAgentInfo
//viettel
	//agentInfo.FullNumber = "0989253496"
//vina
	agentInfo.FullNumber = "01248141091"
//mobi
	//agentInfo.FullNumber = "0938141096"

	agentInfo.DailyLimit = 100
	agentInfo.TotalLimit = 1000
	agentInfo.RewardAddress = "0xf40198036e57B91Ed0CB11872F9259aEF59FE39C"

	theSecondStream_Send := &pb.C2SMessage{
		AuthClient: &pb.C2SMessage_SmsAgent{&agentInfo},
	}

	stream.Send(theSecondStream_Send)
	theSecondStream_Recv,_ := stream.Recv()
	fmt.Println(theSecondStream_Recv)

	go func() {
		s := gocron.NewScheduler()
		s.Every(10).Seconds().Do(run)
		sc := s.Start() // keep the channel:=
		<-sc            // it will happens if the channel is closed
	}()

	for {
		recv,_ := stream.Recv()
		fmt.Println("server said: ",recv.GetMessage())

		var client_send pb.C2SMessage
		client_send.Success = true
		//fmt.Println(" client_send.IncomingMessage.Content ",client_send.IncomingMessage.Content )
		//time.Sleep(time.Second*10)
		stream.Send(&client_send)
	}
}

func main() {
	conn, err := grpc.Dial(appconfig.AddressServerGo, grpc.WithInsecure())
	if err != nil {
		log.Fatal("dail error:", err)
	}
	defer conn.Close()
	c := pb.NewSMSBlockChainHubClient(conn)
	send(c)
}