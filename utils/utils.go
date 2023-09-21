package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strings"
)

func RandHeader() map[string]string {

	head_connection := []string{"Keep-Alive", "close"}
	head_accept := []string{"text/html, application/xhtml+xml, */*"}
	head_accept_language := []string{
		"zh-CN,fr-FR;q=0.5",
		"zh-CN,zh-Hans;q=0.9",
		"en-US,en;q=0.8,zh-Hans-CN;q=0.5,zh-Hans;q=0.3",
	}
	head_user_agent := []string{
		"Opera/8.0 (Macintosh; PPC Mac OS X; U; en)",
		"Opera/9.27 (Windows NT 5.2; U; zh-cn)",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Win64; x64; Trident/4.0)",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.2; .NET4.0C; .NET4.0E)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.2; .NET4.0C; .NET4.0E; QQBrowser/7.3.9825.400)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0; BIDUBrowser 2.x)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1) Gecko/20070309 Firefox/2.0.0.3",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1) Gecko/20070803 Firefox/1.5.0.12",
		"Mozilla/5.0 (Windows; U; Windows NT 5.2) Gecko/2008070208 Firefox/3.0.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.12) Gecko/20080219 Firefox/2.0.0.12 Navigator/9.0.0.6",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.95 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; rv:11.0) like Gecko)",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:21.0) Gecko/20100101 Firefox/21.0 ",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Maxthon/4.0.6.2000 Chrome/26.0.1410.43 Safari/537.1 ",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.92 Safari/537.1 LBBROWSER",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.75 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/3.0 Safari/536.11",
		"Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Macintosh; PPC Mac OS X; U; en) Opera 8.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.5.1 Safari/605.1.15",
	}
	result := map[string]string{
		"Connection":      head_connection[0],
		"Accept":          head_accept[0],
		"Accept-Language": head_accept_language[1],
		"User-Agent":      head_user_agent[int(math.Floor(rand.Float64()))],
	}
	return result
}

// SendDingDingMarkdownToSingleUser 发送个人钉钉消息
func SendDingDingMarkdownToSingleUser(title, message string, template ...string) {
	// defer mkbLog.CrashLogs("", StartProgProduct)
	// go feishu.SendTextToSingleUser(receiverName, title, message, template...)
	// machineID, _ := diffconfig.GetMyMachineId()
	// receiverUserID := "206203482436277267"
	// if machineID != "start" || receiverName == "" {
	// 	if machineID != "start" {
	// 		title += "- 本地测试"
	// 	}
	// } else {
	// 	dingUser, err := casadmin.NewModelDingdingUsers().GetUsersByName(receiverName)
	// 	if err == nil {
	// 		receiverUserID = dingUser.UserID
	// 	}
	// }
	msg := map[string]interface{}{}
	msg["msgtype"] = "markdown"
	msg["markdown"] = map[string]string{"title": title, "text": "### gp123" + title + "gp 123\n" + message}
	tokenUser := "22506ca62efd3042464c0a7bc6e2386e6df96766430215f86215c1cbaec94553"
	jb, _ := json.Marshal(&msg)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s", tokenUser), strings.NewReader(string(jb)))
	// req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// req.Header("Content-Type", "application/json")

	req.Header.Add("content-type", "application/json;charset=utf-8")
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		_ = fmt.Errorf("dingtalk do req err: %v", err)
		return
	}
	defer r.Body.Close()
	// req.JSONBody(map[string]interface{}{
	// 	// "agent_id":    startAgentID,
	// 	// "userid_list": receiverUserID,
	// 	"to_all_user": false,
	// 	"msg":         msg,
	// })
	// ee, _ := req.String()
	// fmt.Println("ddErr:", ee)
}

// SendDingTextToSingleUser 发送个人钉钉消息
func SendDingTextToSingleUser(title, message string, template ...string) {
	// defer mkbLog.CrashLogs("", StartProgProduct)
	// go feishu.SendTextToSingleUser(receiverName, title, message, template...)
	// machineID, _ := diffconfig.GetMyMachineId()
	// receiverUserID := "206203482436277267"
	// if machineID != "start" || receiverName == "" {
	// 	if machineID != "start" {
	// 		title += "- 本地测试"
	// 	}
	// } else {
	// 	dingUser, err := casadmin.NewModelDingdingUsers().GetUsersByName(receiverName)
	// 	if err == nil {
	// 		receiverUserID = dingUser.UserID
	// 	}
	// }
	msg := map[string]interface{}{}
	msg["msgtype"] = "text"
	msg["text"] = map[string]string{"title": title, "content": "### gp123" + title + "gp" + message}
	jb, _ := json.Marshal(&msg)
	tokenUser := "22506ca62efd3042464c0a7bc6e2386e6df96766430215f86215c1cbaec94553"
	// quoted := strconv.QuoteToASCII("gp")
	// subject := "{\"msgtype\":\"text\",\"text\":{\"content\":" + quoted + "}}"
	// url := "http://198.218.6.203:8180/mon/DingDing2?tocken=" + tokenUser + "&subject=" + url.QueryEscape(subject)
	// https://oapi.dingtalk.com/robot/send?access_token
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + tokenUser
	req, _ := http.NewRequest("POST", url, strings.NewReader(string(jb)))
	// req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// req.Header("Content-Type", "application/json")

	req.Header.Add("content-type", "application/json;charset=utf-8")
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		_ = fmt.Errorf("dingtalk do req err: %v", err)
		return
	}
	defer r.Body.Close()
	// req.JSONBody(map[string]interface{}{
	// 	// "agent_id":    startAgentID,
	// 	// "userid_list": receiverUserID,
	// 	"to_all_user": false,
	// 	"msg":         msg,
	// })
	// ee, _ := req.String()
	// fmt.Println("ddErr:", ee)
}
