# GoSmsHub
step1: trong thu muc /cmd/server: $sh run.sh
step2: trong thu muc /cmd/clientThrift/: $sh run.sh
step3: trong thu muc /cmd/clientGo: $sh run.sh

Describe : trong file litstream_manager có 1 interface: citieria: tiêu chí
+ tiêu chí chọn theo nhà mạng: Viettel/Vinaphone/mobifone: mỗi nhà mạng chọn 1 connect theo cirlcle
mỗi 1 connect có nhiều sms được gửi đi, mỗi streamSmS được quản lí bởi ListStreamSMS với key là stream
