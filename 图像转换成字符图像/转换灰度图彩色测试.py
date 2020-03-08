
# -*- coding:utf8 -*-
 
import cv2
 
charSize = 2#字符尺寸
 
string = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
count = len(string)
img = cv2.imread('wm.jpg')
u, v, _= img.shape
c = img*0 + 255
gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
 
for i in range(0, u, 6):
    for j in range(0, v, 6):
        pix = gray[i, j]
        b, g, r = img[i, j]
        zifu = string[int(((count - 1) * pix) / 256)]
        cv2.putText(c, zifu, (j, i), cv2.FONT_HERSHEY_COMPLEX, charSize, (int(b), int(g), int(r)), 1)
 
cv2.imwrite('output.png', c)
