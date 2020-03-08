import os

print('开始扫描目录！')
path = []
inform = {}
for root,dirs,files in os.walk('.\\'):
    print('\n目录：'+str(root))
    path.append(root)
    inform[root] = set()
    for f in files:
        inform[root].add(f)
        print(f,end='  ')
    for d in dirs:
        inform[root].add(d)
        print(d,end='  ')

     
num = 0
path.sort(reverse=True)
for p in path:
    if len(inform[p])==0:#文件夹为空
        os.rmdir(p)
        father,child = reversed(p[::-1].split('\\',1))
        father = father[::-1]
        child = child[::-1]
        if father=='.':
            father='.\\'
        
        print(father,child)
        inform[father].remove(child)
        
        print('rm '+str(p))
        num+=1

print('---------------------')
print('删除完成！,共删除%d个空文件夹'%(num))
os.system('pause>nul')
    
