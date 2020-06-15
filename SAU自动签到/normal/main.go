package main

import (
	//"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/axgle/mahonia"

	//"reflect"
	"io/ioutil"
)

/*
	直接访问尝试获取 cookies
*/
func httpGetFormCookies_hasLog() []*http.Cookie {
	response, err := http.Get("https://app.sau.edu.cn/form/wap/default?formid=10")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	cookies := response.Cookies()
	return cookies
}

/*
	模拟登录智慧沈航并获取一个临时的 cookies
*/
func httpPostLoginCookies() []*http.Cookie {
	response, err := http.PostForm(
		"https://ucapp.sau.edu.cn/wap/login/invalid",
		url.Values{
			"username": {"YourID"},
			"password": {"YourPassword"},
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("模拟登陆结果:\n" + string(body))
	return response.Cookies()
}

/*
	使用 cookies 登录并发送表单信息
*/
func httpPostFormDate(cookies []*http.Cookie) string {
	client := &http.Client{}

	//获取当前日期和随机温度
	date := time.Now().Format("2006-01-02 15:04:05")
	date = date[:10]
	temperature := fmt.Sprintln(36.5 + float64(rand.Intn(3))/10)
	temperature = temperature[:4]

	//表单内容
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

		//这里的ID好像是每个人都特有一个ID，在提交表单的时候会一并发送，需自行抓包查看
		"id":                             {"你的ID"},	
		"fanhuiliaoningshijian":          {""},
		"fanhuididian":                   {""},
		"daofangshijian":                 {""},
		"danweiyuanxi":                   {"计算机学院"},
		"dangqiansuozaishengfen":         {"辽宁省"},
		"dangqiansuozaichengshi":         {"沈阳市"},
	}

	req, err := http.NewRequest(
		"POST",
		"https://app.sau.edu.cn/form/wap/default/save?formid=10",
		strings.NewReader(postData.Encode()),
	)

	if err != nil {
		fmt.Println(err)
	}

	//照着抓来的头原封不动抄的，伪造浏览器访问
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
	req.Header.Add("Content-Length", strconv.Itoa(len(postData.Encode())))

	//套入临时 cookies
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	fmt.Println("post headers:")
	fmt.Println(req)
	fmt.Println("提交表单信息:")
	fmt.Println(postData)

	//发送数据并解压回应信息
	//如果在返回的表单中存在 “操作成功” 字串则提交成功
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	decoder := mahonia.NewDecoder("utf-8")
	body, err := ioutil.ReadAll(decoder.NewReader(resp.Body))
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func main() {

	cookies := httpPostLoginCookies()
	//cookies = httpGetFormCookies_hasLog()
	fmt.Println("获取cookies:")
	fmt.Println(cookies)

	responseBody := httpPostFormDate(cookies)
	fmt.Println("表单提交结果:\n"+responseBody)
}
