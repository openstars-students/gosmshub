package serverthrift

import (
	"fmt"
	"os"
	"gosmshub/appconfig"
	"git.apache.org/thrift.git/lib/go/thrift"
	"gosmshub/thrift/gen-go/sendSms"

	pb "gosmshub/proto"
	"gosmshub/cmd/server/liststream_manager"
	"strconv"
)

type mythriftThrift struct{}

func (this *mythriftThrift) SendSMS(toNumber string, content string)(r string, err error)  {

	//chuyen noi dung tin nhan cho ListenToThriftServer() package manaliststream
	fmt.Println("thriftserver lenght:= ",liststream_manager.TotalConnect())
	if liststream_manager.TotalConnect() != 0 {
		fmt.Println("nhan tin nhan ....................... ben thrift server")
		liststream_manager.Dem ++
		r = toNumber + " " + content
		fmt.Println("[SendSMSofServerThrift]serverThrift: ", r)
		var messageToSend pb.CommandSendSmS
		messageToSend.ToNumber = toNumber
		messageToSend.Content = content + " "+ strconv.Itoa(liststream_manager.Dem)

		CommandSendSmS := &pb.S2CMesssage{
			ChoseCommand: &pb.S2CMesssage_Message{&messageToSend},
		}
		liststream_manager.ListenToThriftServer(CommandSendSmS)

		return "true", nil
		}

	return "false", nil
}

func RunServerThrift(){
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocket(appconfig.AddressServerThrift)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	handler := &mythriftThrift{}
	processor := sendSms.NewSendSmSThriftProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", appconfig.AddressServerThrift)
	server.Serve()
}