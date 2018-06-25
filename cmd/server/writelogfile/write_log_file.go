package writelogfile

import (
	"fmt"
	"os"
	"gosmshub/appconfig"
	pb "gosmshub/proto"
	"io/ioutil"
	"log"
	"time"
)
func CreateFile(path string) {
	appconfig.InitConfig()
	var _, err = os.Stat(path)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) { return }
		defer file.Close()
	}
}

var file, err = os.OpenFile(appconfig.PathLogInCommingMsg, os.O_RDWR, 0777)

func LogInCommingMsg(inCommingMsg *pb.C2SMessage){
	//fmt.Print("time: ",now)
	fmt.Print("Phan hoi tu client: ")
	fmt.Print("  fromNumber: ", inCommingMsg.IncomingMessage.FromNumber)
	fmt.Print("  content: ",inCommingMsg.IncomingMessage.Content )
	fmt.Print("  Time: ",inCommingMsg.IncomingMessage.TimeStamp )

	str := "time: "+ string(inCommingMsg.IncomingMessage.TimeStamp)  +
		"  fromNumber: " +inCommingMsg.IncomingMessage.FromNumber +
		"  content: " +  inCommingMsg.IncomingMessage.Content
	data, err := ioutil.ReadFile(appconfig.PathLogInCommingMsg)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	file, err := os.Create(appconfig.PathLogInCommingMsg) // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done
	x := string(data) + str + "\n"
	file.WriteString(x)
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func LogAppSendMsg(fromNumber string ,commandSendMessage *pb.S2CMesssage){
	//fmt.Print("time: ",now)
	now := time.Now()
	time := now.Format("2006-01-02 15:04:05")
	fmt.Print("[LogAppSendMsg")
	fmt.Print("  fromNumber: ", fromNumber)
	fmt.Print("  content: ",commandSendMessage.GetMessage().GetContent() )
	fmt.Print("  Time: ",commandSendMessage.GetMessage() )

	str := "time: "+ time  +
		"  fromNumber: " +fromNumber +
		"  content: " +  commandSendMessage.GetMessage().GetContent() +
			" toNumber" +  commandSendMessage.GetMessage().GetToNumber()
	data, err := ioutil.ReadFile(appconfig.PathLogAppSendMsg)

	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	file, err := os.Create(appconfig.PathLogAppSendMsg) // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done
	x := string(data) + str + "\n"
	file.WriteString(x)
}