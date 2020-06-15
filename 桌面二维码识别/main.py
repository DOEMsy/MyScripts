import argparse
import pyzbar.pyzbar as pyzbar
from PIL import Image

def QrToStr(path):
    img = Image.open(path).convert('L')
    #img.show()
    data = ''
    for text in  pyzbar.decode(img):
        data += text.data.decode('utf-8')
    return data

parser = argparse.ArgumentParser(description='二维码识别')
parser.add_argument('-i','--img', help='the path of qrcode img.',nargs='*')
args = parser.parse_args()
for path in args.img:
    url = QrToStr(path)
    print(url)
