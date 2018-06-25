package main

import (
	"fmt"
	"net"
	"os"
	"git.apache.org/thrift.git/lib/go/thrift"
	"gosmshub/thrift/gen-go/sendSms"
	"github.com/claudiu/gocron"
	"log"
	"gosmshub/appconfig"
)

func run(){
	log.Println("[run:] client gui lenh len cho serverGo ")
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocket(net.JoinHostPort(appconfig.HostClientThrift, appconfig.PortClientThrift))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	useTransport := transportFactory.GetTransport(transport)
	client := sendSms.NewSendSmSThriftClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+appconfig.HostClientThrift+":"+appconfig.PortClientThrift, " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	//client.SendSMS("01248141095",strconv.Itoa(int(time.Now().Unix())))
	r,_:=  client.SendSMS("01248111113","Hello World")
	fmt.Println(r)
}

func main() {
	s := gocron.NewScheduler()
	s.Every(6).Seconds().Do(run)
	sc := s.Start() // keep the channel:=
	<-sc            // it will happens if the channel is closed
}

