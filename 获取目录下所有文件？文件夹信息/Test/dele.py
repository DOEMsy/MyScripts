import os

print('即将读取文件清单，并删除清单内的文件')
print('请确保文件清单在同一目录内')
print('输入文件清单的8位编号以开始删除：')
sel = input()
mapc = {}
try:
    infile = open(sel+'.txt')
    for linestr in infile:
        linelis = linestr.split('?')
        path = linelis[0]
        mapc[path] = set()
        for f in linelis[1:]:
            mapc[path].add(f)

    lis = []
    print('重复文件如下：')
    for root,dirs,files in os.walk('.\\'):
        for f in files:
            if f in mapc[root]:
                #发现重复文件
                #os.remove(str(root)+str(f))
                if root!='.\\':
                    lis.append(str(root)+'\\'+str(f))    
                    print('$ '+str(root)+'\\'+str(f))
                else:
                    lis.append(str(root)+str(f))    
                    print('$ '+str(root)+str(f))
    print('dele[y/n]?:')
    pre = input()
    if pre != 'y':
        print('没有删除，任意键退出')
        os.system('pause>nul')
        exit()


    for f in lis:
        os.remove(f)
        print('rm '+str(f))

    print('---------------------')
    print('删除完成！,共删除%d个文件'%(len(lis)))
    os.system('pause>nul')
    

except FileNotFoundError:
    print('未找到文件清单！')
    os.system('pause>nul')
