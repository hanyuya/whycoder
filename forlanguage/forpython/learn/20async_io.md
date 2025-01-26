## 协程

在学习异步IO模型前，我们先来了解协程。

协程，又称微线程，纤程。英文名Coroutine。

协程的概念很早就提出来了，但直到最近几年才在某些语言（如Lua）中得到广泛应用。

子程序，或者称为函数，在所有语言中都是层级调用，比如A调用B，B在执行过程中又调用了C，C执行完毕返回，B执行完毕返回，最后是A执行完毕。

所以子程序调用是通过栈实现的，一个线程就是执行一个子程序。

子程序调用总是一个入口，一次返回，调用顺序是明确的。而协程的调用和子程序不同。

协程看上去也是子程序，但执行过程中，在子程序内部可中断，然后转而执行别的子程序，在适当的时候再返回来接着执行。

注意，在一个子程序中中断，去执行其他子程序，不是函数调用，有点类似CPU的中断。比如子程序A、B：
```python
def A():
    print('1')
    print('2')
    print('3')

def B():
    print('x')
    print('y')
    print('z')
```
假设由协程执行，在执行A的过程中，可以随时中断，去执行B，B也可能在执行过程中中断再去执行A，结果可能是：
```python
1
2
x
y
3
z
```
但是在A中是没有调用B的，所以协程的调用比函数调用理解起来要难一些。

看起来A、B的执行有点像多线程，但协程的特点在于是一个线程执行，那和多线程比，协程有何优势？

**最大的优势就是协程极高的执行效率。因为子程序切换不是线程切换，而是由程序自身控制，因此，没有线程切换的开销，和多线程比，线程数量越多，协程的性能优势就越明显。**

**第二大优势就是不需要多线程的锁机制，因为只有一个线程，也不存在同时写变量冲突，在协程中控制共享资源不加锁，只需要判断状态就好了，所以执行效率比多线程高很多。**

因为协程是一个线程执行，那怎么利用多核CPU呢？最简单的方法是多进程+协程，既充分利用多核，又充分发挥协程的高效率，可获得极高的性能。

Python对协程的支持是通过`generator`实现的。

在`generator`中，我们不但可以通过`for`循环来迭代，还可以不断调用`next()`函数获取由`yield`语句返回的下一个值。

但是Python的`yield`不但可以返回一个值，它还可以接收调用者发出的参数。

来看例子：

- 传统的生产者-消费者模型是一个线程写消息，一个线程取消息，通过锁机制控制队列和等待，但一不小心就可能死锁。

如果改用协程，生产者生产消息后，直接通过`yield`跳转到消费者开始执行，待消费者执行完毕后，切换回生产者继续生产，效率极高：
```python
def consumer():
    r = ''
    while True:
        n = yield r
        if not n:
            return
        print('[CONSUMER] Consuming %s...' % n)
        r = '200 OK'

def produce(c):
    c.send(None)
    n = 0
    while n < 5:
        n = n + 1
        print('[PRODUCER] Producing %s...' % n)
        r = c.send(n)
        print('[PRODUCER] Consumer return: %s' % r)
    c.close()

c = consumer()
produce(c)
```
执行结果：
```python
[PRODUCER] Producing 1...
[CONSUMER] Consuming 1...
[PRODUCER] Consumer return: 200 OK
[PRODUCER] Producing 2...
[CONSUMER] Consuming 2...
[PRODUCER] Consumer return: 200 OK
[PRODUCER] Producing 3...
[CONSUMER] Consuming 3...
[PRODUCER] Consumer return: 200 OK
[PRODUCER] Producing 4...
[CONSUMER] Consuming 4...
[PRODUCER] Consumer return: 200 OK
[PRODUCER] Producing 5...
[CONSUMER] Consuming 5...
[PRODUCER] Consumer return: 200 OK
```
注意到consumer函数是一个generator，把一个consumer传入produce后：
1. 首先调用`c.send(None)`启动生成器；
2. 然后，一旦生产了东西，通过`c.send(n)`切换到consumer执行；
3. consumer通过yield拿到消息，处理，又通过yield把结果传回；
4. produce拿到consumer处理的结果，继续生产下一条消息；
5. produce决定不生产了，通过`c.close()`关闭consumer，整个过程结束。
整个流程无锁，由一个线程执行，produce和consumer协作完成任务，所以称为“协程”，而非线程的抢占式多任务。

最后套用Donald Knuth的一句话总结协程的特点：“子程序就是协程的一种特例。”


## 使用asyncio

`asyncio`是Python 3.4版本引入的标准库，直接内置了对异步IO的支持。

`asyncio`的编程模型就是一个消息循环。`asyncio`模块内部实现了`EventLoop`，把需要执行的协程扔到`EventLoop`中执行，就实现了异步IO。

用`asyncio`提供的`@asyncio.coroutine`可以把一个`generator`标记为coroutine``类型，然后`在coroutine`内部用`yield from`调用另一个`coroutine`实现异步操作。

**为了简化并更好地标识异步IO，从Python 3.5开始引入了新的语法`async`和`await`，可以让coroutine的代码更简洁易读**。

用asyncio实现Hello world代码如下：
```python
import asyncio

async def hello():
    print("Hello world!")
    # 异步调用asyncio.sleep(1):
    await asyncio.sleep(1)
    print("Hello again!")

asyncio.run(hello())
```
async把一个函数变成coroutine类型，然后，我们就把这个async函数扔到asyncio.run()中执行。执行结果如下：
```python
Hello!
(等待约1秒)
Hello again!
```
`hello()`会首先打印出Hello world!，然后，`await`语法可以让我们方便地调用另一个`async`函数。由于`asyncio.sleep()`也是一个`async`函数，所以线程不会等待`asyncio.sleep()`，而是直接中断并执行下一个消息循环。当`asyncio.sleep()`返回时，就接着执行下一行语句。

把`asyncio.sleep(1)`看成是一个耗时1秒的IO操作，在此期间，主线程并未等待，而是去执行EventLoop中其他可以执行的async函数了，因此可以实现并发执行。

上述`hello()`还没有看出并发执行的特点，我们改写一下，让两个`hello()`同时并发执行：
```python
# 传入name参数:
async def hello(name):
    # 打印name和当前线程:
    print("Hello %s! (%s)" % (name, threading.current_thread))
    # 异步调用asyncio.sleep(1):
    await asyncio.sleep(1)
    print("Hello %s again! (%s)" % (name, threading.current_thread))
    return name
```
用`asyncio.gather()`同时调度多个async函数：
```python
async def main():
    L = await asyncio.gather(hello("Bob"), hello("Alice"))
    print(L)

asyncio.run(main())
```
执行结果如下：
```python
Hello Bob! (<function current_thread at 0x10387d260>)
Hello Alice! (<function current_thread at 0x10387d260>)
(等待约1秒)
Hello Bob again! (<function current_thread at 0x10387d260>)
Hello Alice again! (<function current_thread at 0x10387d260>)
['Bob', 'Alice']
```
从结果可知，用`asyncio.run()`执行async函数，所有函数均由同一个线程执行。两个`hello()`是并发执行的，并且可以拿到async函数执行的结果（即return的返回值）。

如果把`asyncio.sleep()`换成真正的IO操作，则多个并发的IO操作实际上可以由一个线程并发执行。

我们用`asyncio`的异步网络连接来获取sina、sohu和163的网站首页：
```python
import asyncio

async def wget(host):
    print(f"wget {host}...")
    # 连接80端口:
    reader, writer = await asyncio.open_connection(host, 80)
    # 发送HTTP请求:
    header = f"GET / HTTP/1.0\r\nHost: {host}\r\n\r\n"
    writer.write(header.encode("utf-8"))
    await writer.drain()

    # 读取HTTP响应:
    while True:
        line = await reader.readline()
        if line == b"\r\n":
            break
        print("%s header > %s" % (host, line.decode("utf-8").rstrip()))
    # Ignore the body, close the socket
    writer.close()
    await writer.wait_closed()
    print(f"Done {host}.")

async def main():
    await asyncio.gather(wget("www.sina.com.cn"), wget("www.sohu.com"), wget("www.163.com"))

asyncio.run(main())
```
执行结果如下：
```python
wget www.sohu.com...
wget www.sina.com.cn...
wget www.163.com...
(等待一段时间)
(打印出sohu的header)
www.sohu.com header > HTTP/1.1 200 OK
www.sohu.com header > Content-Type: text/html
...
(打印出sina的header)
www.sina.com.cn header > HTTP/1.1 200 OK
www.sina.com.cn header > Date: Wed, 20 May 2015 04:56:33 GMT
...
(打印出163的header)
www.163.com header > HTTP/1.0 302 Moved Temporarily
www.163.com header > Server: Cdn Cache Server V2.0
...
```
可见3个连接由一个线程并发执行3个async函数完成。

### 小结
asyncio提供了完善的异步IO支持，用`asyncio.run()`调度一个coroutine；

在一个async函数内部，通过`await`可以调用另一个async函数，这个调用看起来是串行执行的，但实际上是由asyncio内部的消息循环控制；

在一个async函数内部，通过`await asyncio.gather()`可以并发执行若干个async函数。

## 使用aiohttp

`asyncio`可以实现单线程并发IO操作。如果仅用在客户端，发挥的威力不大。如果把`asyncio`用在服务器端，例如Web服务器，由于HTTP连接就是IO操作，因此可以用`单线程+async`函数实现多用户的高并发支持。

`asyncio`实现了TCP、UDP、SSL等协议，`aiohttp`则是基于`asyncio`实现的HTTP框架。

我们先安装aiohttp：
```python
$ pip install aiohttp
```
然后编写一个HTTP服务器，分别处理以下URL：
- `/` - 首页返回Index Page；
- `/{name}` - 根据URL参数返回文本Hello, {name}!。

代码如下：
```python
# app.py
from aiohttp import web

async def index(request):
    text = "<h1>Index Page</h1>"
    return web.Response(text=text, content_type="text/html")

async def hello(request):
    name = request.match_info.get("name", "World")
    text = f"<h1>Hello, {name}</h1>"
    return web.Response(text=text, content_type="text/html")

app = web.Application()

# 添加路由:
app.add_routes([web.get("/", index), web.get("/{name}", hello)])

if __name__ == "__main__":
    web.run_app(app)
```
直接运行app.py，访问首页：
```python
Index
```
访问http://localhost:8080/Bob：
```python
Hello
```
使用aiohttp时，定义处理不同URL的async函数，然后通过`app.add_routes()`添加映射，最后通过`run_app()`以`asyncio`的机制启动整个处理流程。

