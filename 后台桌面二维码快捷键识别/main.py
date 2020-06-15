import argparse
import keyboard
import pyzbar.pyzbar as pyzbar
import pyautogui as pag    #监听鼠标
from PIL import Image
from PIL import ImageGrab

def QrToStr(path):
    img = Image.open(path).convert('L')
    #img.show()
    data = ''
    for text in  pyzbar.decode(img):
        data += text.data.decode('utf-8')
    return data


while(True):
    if keyboard.wait(hotkey='ctrl+alt') == None:
        x1, y1 = pag.position()
        print("获取对角1")
    if keyboard.wait(hotkey='ctrl+alt') == None:
        x2, y2 = pag.position()
        if x1>x2: x1,x2 = x2,x1
        if y1>y2: y1,y2 = y2,y1
        print("获取对角2")
        image = ImageGrab.grab((x1, y1, x2, y2))
        image.save(".\\tmp.png")
        url = QrToStr('.\\tmp.png')
        print("解析结果："+url)