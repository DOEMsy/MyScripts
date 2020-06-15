package main

import (
	//"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	//"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/go-gomail/gomail"

	//"reflect"
	"io/ioutil"
)

func httpGetFormCookies_hasLog() []*http.Cookie {
	response, err := http.Get("https://app.sau.edu.cn/form/wap/default?formid=10")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	//body, _ := ioutil.ReadAll(reponse.Body)
	//fmt.Println(string(body))
	//fmt.Println(response)
	cookies := response.Cookies()
	//fmt.Println(cookies)
	//fmt.Println(reflect.TypeOf(cookies))
	return cookies
}

func httpPostLoginCookies() ([]*http.Cookie,string){
	logMessage := "开始模拟登陆：\n"
	response, err := http.PostForm(
		"https://ucapp.sau.edu.cn/wap/login/invalid",
		url.Values{
			"username": {"YourID"},
			"password": {"YourPassword"},
		},
	)
	if err != nil {
		fmt.Println(err)
		logMessage+=err.Error()+"\n"
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		logMessage+=err.Error()+"\n"
	}

	fmt.Println("模拟登陆结果:\n" + string(body))
	logMessage+="模拟登陆结果:\n" + string(body) + "\n"

	return response.Cookies(),logMessage
}

func httpPostFormDate(cookies []*http.Cookie) (string,string){
	logMessage := "开始提交表单：\n" 

	client := &http.Client{}
	date := time.Now().Format("2006-01-02 15:04:05")
	date = date[:10]
	temperature := fmt.Sprintln(36.5 + float64(rand.Intn(3))/10)
	temperature = temperature[:4]

	postData := url.Values{
		"xuehao":                         {"你的学号"},
		"xingming":                       {"你的姓名"},
		"tiwen":                          {temperature},
		"shoujihao":                      {"你的电话号"},
		"shifouyuhubeiwuhanrenyuanmiqie": {"否"},
		"shifouyouxiaohuadabushizhengzh": {"否"},
		"shifouyouhuxidaobushizhengzhua": {"否"},
		"shifouyoufare":                  {"否"},
		"shifouweigelirenyuan":           {"否"},
		"shifouquguomouyiliaojigoufaren": {"否"},
		"shifouquguohuoluguowuhanhuohub": {"否"},
		"shifoujiechuguoyisizhengzhuang": {"否"},
		"shifoujiechuguohubeihuoqitayou": {"否"},
		"shifoujiankangqingkuang":        {"是"},
		"shifouhuanhuiliaoning":          {""},
		"shentishifouyoubushizhengzhuan": {"否"},
		"shenbianjiarenyouwuyisizhengzh": {"否"},
		"riqi":                           {date},
		"qitaxinxi":                      {""},
		"jiechushijian":                  {""},
		"jiechudidian":                   {""},
		"id":                             {"你的ID"},
		"fanhuiliaoningshijian":          {""},
		"fanhuididian":                   {""},
		"daofangshijian":                 {""},
		"danweiyuanxi":                   {"学院"},
		"dangqiansuozaishengfen":         {"省份"},
		"dangqiansuozaichengshi":         {"城市"},
	}

	req, err := http.NewRequest(
		"POST",
		"https://app.sau.edu.cn/form/wap/default/save?formid=10",
		strings.NewReader(postData.Encode()),
	)

	if err != nil {
		fmt.Println(err)
		logMessage += fmt.Sprintln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	req.Header.Add("Host", "app.sau.edu.cn")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Origin", "https://app.sau.edu.cn")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	req.Header.Add("Referer", "https://app.sau.edu.cn/form/wap/default/index?formid=10&nn=4669.797748311082")
	req.Header.Add("Accept-Encoding", "")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	//req.Header.Add("Cookie",fmt.Sprint(cookies))
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	fmt.Println("post headers:")
	fmt.Println(req)
	fmt.Println("提交表单信息:")
	fmt.Println(postData)
	logMessage += "post headers：\n"
	for key,value := range req.Header{
		logMessage += key + " :  " + fmt.Sprintln(value)
	}

	logMessage += "提交表单信息：\n"
	for key,value := range postData{
		logMessage += key + " :  " + fmt.Sprintln(value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		logMessage += fmt.Sprintln(err)
	}

	defer resp.Body.Close()
	decoder := mahonia.NewDecoder("utf-8")
	body, err := ioutil.ReadAll(decoder.NewReader(resp.Body))
	if err != nil {
		fmt.Println(err)
		logMessage += fmt.Sprintln(err)
	}
	//fmt.Println(string(body))
	return string(body),logMessage
}

func main() {
	
	m := gomail.NewMessage()

	m.SetAddressHeader("From", "你的IMAP/SMTP邮箱" /*"发件人地址"*/, "发件人") // 发件人

	m.SetHeader("To",                                                            
		m.FormatAddress("收件邮箱", "收件人")) // 收件人
//	m.SetHeader("Cc",
//		m.FormatAddress("xxxx@foxmail.com", "收件人")) //抄送
//	m.SetHeader("Bcc",  -
//		m.FormatAddress("xxxx@gmail.com", "收件人")) /暗送

	m.SetHeader("Subject", "签到成功")     // 主题

	//m.SetBody("text/html",xxxxx ") // 可以放html..还有其他的	
	var SendMessage string

	cookies, logMessage := httpPostLoginCookies()
	//cookies = httpGetFormCookies_hasLog()
	fmt.Println("获取cookies:")
	fmt.Println(cookies)
	for _,value := range cookies{
		logMessage += fmt.Sprintln(value)
	}

	SendMessage += logMessage + "\n"

	responseBody, logMessage := httpPostFormDate(cookies)
	fmt.Println("表单提交结果:\n"+responseBody)
	logMessage += "表单提交结果:\n"+responseBody+"\n"

	SendMessage += logMessage + "\n"
	SendMessageHtml := ""
	//bot.Send(msg)
	for i:=0;i<len(SendMessage);i++{
		if SendMessage[i] != '\n'{
			SendMessageHtml += string(SendMessage[i])
		}else{
			SendMessageHtml += "</br>"
		} 
	}

	m.SetBody("text/html",SendMessage) // 正文

	//	m.Attach("我是附件")  //添加附件
	
	d := gomail.NewPlainDialer("smtp.example.com", 587, "你的IMAP/SMTP邮箱账号", "密码") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err)
		return
	}
	log.Println("done.发送成功")
}
