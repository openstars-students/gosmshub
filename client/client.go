package main

import (
	"fmt"
	"log"
	pb "example/gosmshub"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
	"time"
)
const (
	address = "0.0.0.0:10023"
)
func send(client pb.SendSMSClient) {
	stream, err := client.Send(context.Background())
	if err != nil {
		log.Fatal("client Chat get stream error: ", err)
	}
	t := strconv.Itoa(int(time.Now().Unix()))
	var dangki pb.SendMessage
	dangki.ToNumber = t
	stream.Send(&dangki)

	for {
		recv,_ := stream.Recv()
		fmt.Println("toNumber: ",recv.ToNumber,"  content: ",recv.Content)

		var phanhoi pb.SendMessage
		phanhoi.Content ="sdt: " +t + " da nhan duoc lenh"
		stream.Send(&phanhoi)
	}
}
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("dail error:", err)
	}
	defer conn.Close()
	c := pb.NewSendSMSClient(conn)

	send(c)
}