## 模块（Module）
- 在Python中，一个.py文件就称之为一个模块（Module）。
- 为了避免模块名冲突，Python又引入了按目录来组织模块的方法，称为包（Package）。
- 包可以看成是模块的容器，可以包含多个模块，每个模块可以有自己的变量、函数和类。
- 每一个包目录下面都会有一个`__init__.py`的文件，这个文件是必须存在的，否则，Python就把这个目录当成普通目录，而不是一个包。
- `__init__.py`可以是空文件，也可以有Python代码，因为`__init__.py`本身就是一个模块，而它的模块名就是`mycompany`。
```python
mycompany
├─ __init__.py
├─ abc.py
└─ xyz.py
```
- 在上面的例子中，mycompany是一个包，里面包含了两个模块：abc.py和xyz.py。
- 可以有多级目录，组成多级层次的包结构。比如如下的目录结构：
```python
mycompany
 ├─ web
 │  ├─ __init__.py
 │  ├─ utils.py
 │  └─ www.py
 ├─ __init__.py
 ├─ abc.py
 └─ utils.py
```
- 文件www.py的模块名就是mycompany.web.www，两个文件utils.py的模块名分别是mycompany.utils和mycompany.web.utils。
```python
#!/usr/bin/env python3
# -*- coding: utf-8 -*-

' a test module '

__author__ = 'Michael Liao'

import sys

def test():
    args = sys.argv
    if len(args)==1:
        print('Hello, world!')
    elif len(args)==2:
        print('Hello, %s!' % args[1])
    else:
        print('Too many arguments!')

if __name__=='__main__':
    test()
```
- 第1行和第2行是标准注释，第1行注释可以让这个hello.py文件直接在Unix/Linux/Mac上运行，第2行注释表示.py文件本身使用标准UTF-8编码；
- 第4行是模块的文档注释，任何模块代码的第一个字符串都被视为模块的文档注释；
- 第6行是作者注释，表示模块的作者是谁；
- 第8行是导入sys模块，后面会用到；
- 第10行定义了一个test()函数，用于输出一句话；
- 第12行是主程序入口，只有在直接运行hello.py文件时才会执行这里的代码，否则，当hello.py被导入到其他模块中时，这里的代码就不会被执行。
- 当我们在命令行运行hello模块文件时，Python解释器把一个特殊变量__name__置为__main__，而如果在其他地方导入该hello模块时，if判断将失败，因此，这种if测试可以让一个模块通过命令行运行时执行一些额外的代码，最常见的就是运行测试。

## 作用域
- 正常的函数和变量名是公开的（public），可以被直接引用，比如：abc，x123，PI等；
- 类似__xxx__这样的变量是特殊变量，可以被直接引用，但是有特殊用途，比如上面的__author__，__name__就是特殊变量，hello模块定义的文档注释也可以用特殊变量__doc__访问，我们自己的变量一般不要用这种变量名；
- 类似_xxx和__xxx这样的函数或变量就是非公开的（private），不应该被直接引用，比如_abc，__abc等；
- 之所以我们说，正常的函数和变量名是公开的，是因为Python本身没有任何机制阻止我们直接访问private函数或变量，但是，按照约定俗成的规定，当函数或变量名以单下划线开头时，表示它是一个非公开的函数或变量，不应该被直接引用。
- 外部不需要引用的函数全部定义成private，只有外部需要引用的函数才定义为public。

## 安装模块

在Python中，安装第三方模块，有两种方式：
1. 通过包管理工具pip安装，pip是一个第三方包管理工具，可以自动安装和管理Python第三方模块。
2. 通过源码安装，这种方式需要先下载模块的源码包，然后手动安装。

推荐使用pip安装，因为pip可以自动处理依赖关系，自动下载安装依赖的模块。

举例：
1. 安装requests模块，可以直接在命令行通过pip安装：`pip install requests`
2. 安装lxml模块，可以先下载源码包，然后手动安装：
   - 下载lxml源码包：`wget https://pypi.python.org/packages/source/l/lxml/lxml-3.4.2.tar.gz`
   - 解压源码包：`tar -zxvf lxml-3.4.2.tar.gz`
   - 进入源码包目录：`cd lxml-3.4.2`
   - 安装lxml：`python setup.py install`
   - 这样，lxml模块就安装成功了。

## 模块搜索路径
当我们导入一个模块时，Python解释器会按照以下路径搜索模块：
1. 当前目录、所有已安装的内置模块和第三方模块，**搜索路径存放在sys模块的path变量中**。
2. 如果当前目录下没有找到模块，Python解释器还会到环境变量**PYTHONPATH**指定的目录搜索。
```python
>>> import sys
>>> sys.path
['', '/usr/lib64/python38.zip', '/usr/lib64/python3.8', '/usr/lib64/python3.8/lib-dynload', '/home/.local/lib/python3.8/site-packages', '/usr/lib64/python3.8/site-packages', '/usr/lib/python3.8/site-packages']
```









