import json
import time
import requests
import smtplib
from email.mime.text import MIMEText
from email.header import Header
from_addr = '' #发件人
password='' #发件人SMTP授权码
to_addr='' #收件人
zfytime = time.strftime('%Y-%m-%d')
data ={ 
    'xingming' : "",#姓名
    "xuehao": "",#学号
    "shoujihao": "",#手机
    "danweiyuanxi": "",#学院
    "dangqiansuozaishengfen": "",
    "dangqiansuozaichengshi": "",
    "shifouyuhubeiwuhanrenyuanmiqie": "",
    "shifoujiankangqingkuang": "",
    "shifoujiechuguohubeihuoqitayou": "",
    "fanhuididian": "",
    "shifouweigelirenyuan": "",
    "shentishifouyoubushizhengzhuan": "",
    "shifouyoufare": "",
    "qitaxinxi": "",
    "tiwen": "36.5",#体温
    "tiwen1": "36.5",#体温
    "tiwen2": "36.5",#体温
    "riqi": zfytime,#每天的时间
    "id": "431243206"#抓包后的id
    }
url = "https://app.sau.edu.cn/form/wap/default/save?formid=10"
header1 ={
    "Host": "app.sau.edu.cn",
    "Connection": "keep-alive",
    "Cache-Control": "max-age=0",
    "sec-ch-ua": '"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"',
    "sec-ch-ua-mobile": "?0",
    "Upgrade-Insecure-Requests": "1",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
    "Sec-Fetch-Site": "same-origin",
    "Sec-Fetch-Mode": "navigate",
    "Sec-Fetch-User": "?1",
    "Sec-Fetch-Dest": "document",
    "Referer": "https://app.sau.edu.cn/form/wap/default/index?formid=10",
    "Accept-Language": "zh,zh-CN;q=0.9",
    "Cookie": ""}#浏览器的Cooke
zfy=requests.post(url=url,headers=header2,data=data)
a=json.loads(zfy.text)
if a['m']=='操作成功'and a['e']==0 and a['d']==[]:
    msg=MIMEText('打卡成功','plain','utf-8')
    msg['From']=Header(from_addr)
    msg['To'] = Header(to_addr)
    msg['Subject'] = Header('打卡信息')
    server =smtplib.SMTP_SSL('smtp.qq.com',465)
    server.login(from_addr,password)
    server.sendmail(from_addr,to_addr,msg.as_string())
    server.quit()
else :
    msg=MIMEText('打卡失败','plain','utf-8')
    msg['From']=Header(from_addr)
    msg['To'] = Header(to_addr)
    msg['Subject'] = Header('打卡信息')
    server =smtplib.SMTP_SSL('smtp.qq.com',465)
    server.login(from_addr,password)
    server.sendmail(from_addr,to_addr,msg.as_string())
    server.quit()
