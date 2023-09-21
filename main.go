package main

import (
	"fmt"
	"gp/utils"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var ttime time.Duration
var tc *time.Ticker

var (
	ttimeclose = time.Duration(60)
	ttimeopen  = time.Duration(6)
)

var percentMap = map[string]int{}

func main() {
	fmt.Println("hello")
	percentMap = make(map[string]int)
	ttime = ttimeopen
	tc = time.NewTicker(time.Second * ttime)
	getsocketdata()
	// return
	for {
		select {
		case <-tc.C:
			getsocketdata()
		default:
			// fmt.Println("hello")
		}
	}
}

var url = "https://hq.sinajs.cn/list=sh000001,sz399001,sz002585"

// var url = "https://hq.sinajs.cn/list=sh000001"

// var url2 = "https://stock.xueqiu.com/v5/stock/batch/quote.json?symbol=002585"

func getsocketdata() {
	if isclose() {
		return
	}
	client := http.DefaultClient
	rq, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		fmt.Println("rq err:", err1.Error())
		return
	}
	// rq.Header.Add("Referer", "https://stock.xueqiu.com/")
	rq.Header.Add("Referer", "http://finance.sina.com.cn/")
	for k, v := range utils.RandHeader() {
		// fmt.Println(k, v)
		rq.Header.Add(k, v)
	}
	resp, err := client.Do(rq)
	if err != nil {
		fmt.Println("get err:", err.Error())
		return
	}
	bd, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body err:", err.Error())
		return
	}
	defer resp.Body.Close()
	dec := simplifiedchinese.GB18030.NewDecoder()
	out, _ := dec.String(string(bd))
	// fmt.Println(out)
	var tm string
	for _, hqdata := range strings.Split(out, ";") {
		/* hqdata
		var hq_str_sz002585="双星新材,9.730,9.780,9.730,9.830,9.710,9.730,9.740,1371300,13365584.000,1500,9.730,30900,9.720,43200,9.710,79800,9.700,19600,9.690,14400,9.740,21600,9.750,6100,9.760,13900,9.770,13300,9.780,2023-09-12,09:38:36,00"
		*/
		// fmt.Println(hqdata)
		if !strings.Contains(hqdata, "hq_str_") {
			continue
		}
		hqdata = strings.ReplaceAll(hqdata, "\"", "")
		code := strings.Split(strings.Split(hqdata, "=")[0], "var hq_str_")[1]
		param := strings.Split(strings.Split(hqdata, "=")[1], ",")
		if len(param) < 1 {
			continue
		}
		// typ := code[:2] == "sh"
		symbol := code[2:]
		// fmt.Println(code, symbol)
		/* var hq_str_sz002585="双星新材,
		1	10.730,
		2	10.760,
		3	10.680,
		4	10.770,
		5	10.590,
		6	10.670,
		7	10.680,
		8	4438059,
		9	47411548.350,
		10	23200,
		11	10.670,
		12	55400,
		13	10.660,
		14	64700,
		15	10.650,
		16	23300,
		17	10.640,
		18	43400, ba5
		19	10.630, b5
		20	20300,
		21	10.680,
		22	39200,
		23	10.690,
		24	15100,
		25	10.700,
		26	18800,
		27	10.710,
		28	6000, sa5
		29	10.720, s5
			2023-08-07,
			10:12:33,
			00"

			集合竞价时返回
			双星新材,
			0.000,
			10.430,
			0.000,
			0.000,
			0.000,0.000,0.000,0,0.000,
			0,0.000,0,0.000,0,
			0.000,0,0.000,0,
			0.000,0,0.000,0,0.000,
			0,0.000,0,0.000,0,0.000,
			2023-08-09,09:12:09,00
		*/

		fmt.Println("\n" + param[0] + "[" + code + "]") //双星新材[sz002585]
		open := param[1]
		yestclose := param[2]
		price := param[3]
		high := param[4]
		low := param[5]
		volume := param[8]
		amount := param[9]

		s1 := param[21]
		s2 := param[23]
		s3 := param[25]
		s4 := param[27]
		s5 := param[29]

		b1 := param[11]
		b2 := param[13]
		b3 := param[15]
		b4 := param[17]
		b5 := param[19]
		var sa1, sa2, sa3, sa4, sa5, ba1, ba2, ba3, ba4, ba5 string
		if len(param[28]) > 2 {
			sa1 = param[20][:len(param[20])-2]
			sa2 = param[22][:len(param[22])-2]
			sa3 = param[24][:len(param[24])-2]
			sa4 = param[26][:len(param[26])-2]
			sa5 = param[28][:len(param[28])-2]

			ba1 = param[10][:len(param[10])-2]
			ba2 = param[12][:len(param[12])-2]
			ba3 = param[14][:len(param[14])-2]
			ba4 = param[16][:len(param[16])-2]
			ba5 = param[18][:len(param[18])-2]
		}
		if price == "0.000" {
			price = s1
		}
		tm = fmt.Sprintf("%s %s", param[30], param[31])
		pricef, _ := strconv.ParseFloat(price, 64)
		// openf, _ := strconv.ParseFloat(open, 64)
		yestclosef, _ := strconv.ParseFloat(yestclose, 64)
		prec := 0.0
		color := ""
		diff := 0.0
		flags5 := "  "
		flags4 := "  "
		flags3 := "  "
		flags2 := "  "
		flags1 := "  "
		flagb1 := "  "
		flagb2 := "  "
		flagb3 := "  "
		flagb4 := "  "
		flagb5 := "  "

		fb1 := 0.0
		fb2 := 0.0
		fb3 := 0.0
		fb4 := 0.0
		fb5 := 0.0

		fs1 := 0.0
		fs2 := 0.0
		fs3 := 0.0
		fs4 := 0.0
		fs5 := 0.0

		if yestclosef > 0 {
			diff = pricef - yestclosef
			if diff < 0 {
				// negative
				color = "green"
			} else if diff > 0 {
				// positive
				color = "red"
			}
			prec = diff * 100 / yestclosef
			flagb1, fb1 = getflag(b1, yestclosef)
			flagb2, fb2 = getflag(b2, yestclosef)
			flagb3, fb3 = getflag(b3, yestclosef)
			flagb4, fb4 = getflag(b4, yestclosef)
			flagb5, fb5 = getflag(b5, yestclosef)
			flags1, fs1 = getflag(s1, yestclosef)
			flags2, fs2 = getflag(s2, yestclosef)
			flags3, fs3 = getflag(s3, yestclosef)
			flags4, fs4 = getflag(s4, yestclosef)
			flags5, fs5 = getflag(s5, yestclosef)
		}
		precstr := fmt.Sprintf(" %.2f%%  %.2f ", prec, diff)
		if color == "green" {
			/*
				\033[30m 	\033[40m
				\033[31m 红	\033[41m
				\033[32m 绿	\033[42m
				\033[33m	\033[43m
				\033[34m	\033[44m
				\033[35m	\033[45m
				\033[36m	\033[46m
				\033[37m	\033[47m
			*/
			precstr = fmt.Sprintf(" \033[32m%.2f%%  %.2f\033[0m", prec, diff)
		} else if color == "red" {
			precstr = fmt.Sprintf(" \033[31m%.2f%%  %.2f\033[0m ", prec, diff)
		}
		// utils.SendDingTextToSingleUser("gp", precstr)
		tmpk := fmt.Sprintf("%.f", math.Floor(prec))

		precentkey := fmt.Sprintf("%s-%s", code, tmpk)
		fmt.Printf("map is :%+v\n", percentMap)
		if _, ok := percentMap[precentkey]; !ok {
			pstr := []string{"-10", "-9", "-8", "-7", "-6", "-5", "-4", "-3", "-2", "-1",
				"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
			//涨幅浮动1个点清空涨幅缓存
			for _, k := range pstr {
				delk := fmt.Sprintf("%s-%s", code, k)
				if _, ok := percentMap[delk]; ok && k != tmpk {
					delete(percentMap, delk)
				}
			}
			//将涨幅缓存,控制报警次数
			percentMap[precentkey] = 0
			if prec < -2.4 || prec > 5.3 {
				// https://oapi.dingtalk.com/robot/send?access_token=22506ca62efd3042464c0a7bc6e2386e6df96766430215f86215c1cbaec94553
				utils.SendDingTextToSingleUser("gp", precstr)
			}
		}
		// fmt.Println(pricef, openf)
		fmt.Println("开:", open, "昨收:", yestclose, "\n现:", price, "高:", high, "低:", low, precstr)
		if "000001" != symbol && "399001" != symbol {
			printf(flags5, "s5:", s5, sa5, fs5)
			printf(flags4, "s4:", s4, sa4, fs4)
			printf(flags3, "s3:", s3, sa3, fs3)
			printf(flags2, "s2:", s2, sa2, fs2)
			printf(flags1, "s1:", s1, sa1, fs1)
			fmt.Println()
			printf(flagb1, "b1:", b1, ba1, fb1)
			printf(flagb2, "b2:", b2, ba2, fb2)
			printf(flagb3, "b3:", b3, ba3, fb3)
			printf(flagb4, "b4:", b4, ba4, fb4)
			printf(flagb5, "b5:", b5, ba5, fb5)
		}
		printf("vol:", volume, "amount:", amount, 0.0)
	}
	fmt.Println("\ntm:", tm)
}

func getflag(sobstr string, cur float64) (string, float64) {
	sob, _ := strconv.ParseFloat(sobstr, 64)
	diff := sob - cur
	prec := diff * 100 / cur
	if diff < 0 {
		// negative
		return "-", prec
	} else if diff > 0 {
		// positive
		return "+", prec
	}
	return "±", prec
}

func printf(flag, lev, price, num string, precent float64) {
	fmt.Printf("\n%s %s %s  %s   %.2f%%", flag, lev, price, num, precent)
}

func isclose() bool {
	time1457, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 ")+"15:00:11")
	time1300, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 ")+"12:59:12")
	time1130, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 ")+"11:30:12")
	time0925, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 ")+"09:15:30")
	now, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	weekday := time.Now().Weekday()
	// fmt.Println(now, weekday)
	res1457 := now.Compare(time1457) // 大于15点 > 0
	res1300 := now.Compare(time1300) // 小于 13点 < 0
	res1130 := now.Compare(time1130) // 大于 11点半 > 0
	res0925 := now.Compare(time0925) //小于9点半 < 0

	if res0925 < 0 || res1457 > 0 || (res1130 > 0 && res1300 < 0) || weekday == time.Saturday || weekday == time.Sunday {
		fmt.Println("close")
		ttime = ttimeclose
		tc.Reset(time.Second * ttime)
		return true
	}
	ttime = ttimeopen
	tc.Reset(time.Second * ttime)
	fmt.Println("open")
	return false
}
