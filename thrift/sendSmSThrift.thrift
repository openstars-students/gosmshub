namespace go sendSms

service SendSmSThrift {
        string sendSMS(1:string toNumber, 2:string content),
}
