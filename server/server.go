package main

import (
	"fmt"
	"log"
	"net"
	pb "example/SendSMS"
	"google.golang.org/grpc"
	"git.apache.org/thrift.git/lib/go/thrift"
	"example/SendSMS/serverThrift/gen-go/demo"
	"strings"
	"os"
	"github.com/claudiu/gocron"

)
const (
	address = "0.0.0.0:10023"
)
const (
	HOST = "127.0.0.1"
	PORT = "9090"
)
type SendSMSService struct {
}

var clients = make(map[string]*Client)

type Client struct {
	sdt string
	index int
	stream pb.SendSMS_SendServer
	status bool
}

func listenToClient(stream pb.SendSMS_SendServer,sdt string ){
	for {
		a, _ := stream.Recv()
		if a != nil {
			log.Print("[listenToClient]: Phan hoi tu client: ", a.GetContent())
			if clients[sdt].status {
				clients[sdt].index ++
				log.Println("[listenToClient]: index: ", clients[sdt].index)
			}
		} else {
			log.Println("[listenToClient]: MAT KET NOI vs Sdt: ", clients[sdt].sdt)
			clients[sdt].status = false
		}
	}
}
func task() {
	fmt.Println("I am runnning task.")
}
func abc(){
	s := gocron.NewScheduler()
	s.Every(1).Day().Do(task)

	sc := s.Start() // keep the channel
	//go test(s, sc)  // wait
	<-sc            // it will happens if the channel is closed
}

var dem = 0
func (svr *SendSMSService) Send(stream pb.SendSMS_SendServer) error {
//luu cau truc thong tin cua 1 client

	dangki, _ := stream.Recv()
	client := &Client{
		sdt:  dangki.GetToNumber(),
		stream: stream,
		status:true,
		index: 0,
	}
	clients[dangki.GetToNumber()] = client
//lang nghe phan hoi tu client
	go listenToClient(stream, dangki.GetToNumber())

	//abc()
	log.Println("[Send]: ket thuc ham Send")
	return nil
}
//nhan lenh tu thrift
func listenFromThrift() {
	for {
		log.Println("listenFromThrift")
		transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

		transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}

		useTransport := transportFactory.GetTransport(transport)
		client := demo.NewMyThriftClientFactory(useTransport, protocolFactory)

		if err := transport.Open(); err != nil {
			fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
			os.Exit(1)
		}

		defer transport.Close()
//gui tin hieu cho serverthrift va doi lenh
		sendmess, errr := client.SendSMS()

		if errr != nil{
			log.Println("ko ket noi duoc voi server thrift")
			continue
		}
		if len(clients) > 0 && sendmess != "" && errr == nil {
			s := strings.Split(string(sendmess), " ")

			var mess pb.SendMessage
			mess.Content = s[1]
			mess.ToNumber = s[0]
			sdt := choseConnect(clients)
			sendSmS(sdt,&mess)
			//broascast(&mess)
		}
	}
}

func choseConnect(clients map[string]*Client) string {
	var token string
	var min = 1000000000

	for _, client := range clients{
		log.Println("[choseConnect]:Chay den sdt: ", client.sdt)
		if client.index <= min && client.status == true {
			log.Println("[choseConnect]:client: ", client.sdt)
			log.Println("[choseConnect]:min: ", min)
			min = client.index
			token = client.sdt
		}
	}
	return token
}

func sendSmS(sdt string, mess *pb.SendMessage){
	if clients[sdt] != nil{
		clients[sdt].stream.Send(mess)
	}
}
func broascast(mess *pb.SendMessage){

	for _, client := range clients{
		client.stream.Send(mess)
	}
}

func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterSendSMSServer(s, &SendSMSService{})

	fmt.Println("Listening on the 0.0.0.0:10023")

	go listenFromThrift()

	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}