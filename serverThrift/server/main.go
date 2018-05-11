package main

import (
    "fmt"
    "os"
    "git.apache.org/thrift.git/lib/go/thrift"
    "example/gosmshub/serverThrift/gen-go/demo"
    "bufio"
    "strings"
)

const (
    NetworkAddr = "127.0.0.1:9090"
)

type mythriftThrift struct{}

func (this *mythriftThrift) SendSMS()(r string, err error)  {

    var sdt string
    fmt.Print("Nhap sdt: ")
    user := bufio.NewReader(os.Stdin)
    sdt,_ = user.ReadString('\n')
    sdt = strings.TrimSpace(sdt)

    var content string
    fmt.Print("Nhap content: ")
    user = bufio.NewReader(os.Stdin)
    content,_ = user.ReadString('\n')
    content = strings.TrimSpace(content)

    r = sdt + " " + content
    //r = "1" + " " +"2"
    //time.Sleep(time.Second*2)
    return
}

func main() {

    transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
    if err != nil {
        fmt.Println("Error!", err)
        os.Exit(1)
    }

    handler := &mythriftThrift{}
    processor := demo.NewMyThriftProcessor(handler)

    server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
    fmt.Println("thrift server in", NetworkAddr)
    server.Serve()
}
