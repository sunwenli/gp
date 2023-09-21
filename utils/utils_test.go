package utils

import "testing"

func TestRobotSend(t *testing.T) {
	RobotSend("gp123")
}

func TestSendDingTextToSingleUser(t *testing.T) {
	SendDingTextToSingleUser("gp", "gp")
	/*
			curl 'https://oapi.dingtalk.com/robot/send?access_token=22506ca62efd3042464c0a7bc6e2386e6df96766430215f86215c1cbaec94553' \
		 -H 'Content-Type: application/json' \
		 -d '{"msgtype": "text","text": {"content":"我就是我, 是不一样的烟火"}}'
	*/
}
