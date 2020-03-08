import os
import time
import subprocess 

infile = open('test.txt','r')

i = 0
screen = ''


for line in infile:
    if i == 31:
        #print("\n"*80)
        #subprocess.call("cls", shell=True)
        time.sleep(1/35)
        os.system('clear')
        print(screen)
        screen = ''
        i = 0

    screen += line
    i+=1

infile.close()

#220s 6600
