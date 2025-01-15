# 错误、调试和测试

# 错误处理
# try-except-finally
    # try:
    #     print('try...')
    #     r = 10 / 0
    #     print('result:', r)
    # except ZeroDivisionError as e:
    #     print('except:', e)
    # finally:
    #     print('finally...')
    # print('END')
# 当错误发生时，后续语句print('result:', r)不会被执行，except由于捕获到ZeroDivisionError，因此被执行。最后，finally语句被执行。然后，程序继续按照流程往下走。
# 有错误发生，所以except语句块不会被执行，但是finally如果有，则一定会被执行（可以没有finally语句）。
# 错误应该有很多种类，如果发生了不同类型的错误，应该由不同的except语句块处理。没错，可以有多个except来捕获不同类型的错误：
    # try:
    #     print('try...')
    #     r = 10 / int('a')
    #     print('result:', r)
    # except ValueError as e:
    #     print('ValueError:', e)
    # except ZeroDivisionError as e:
    #     print('ZeroDivisionError:', e)
    # finally:
    #     print('finally...')
    # print('END')
# 此外，如果没有错误发生，可以在except语句块后面加一个else，当没有错误发生时，会自动执行else语句：
    # try:
    #     print('try...')
    #     r = 10 / int('2')
    #     print('result:', r)
    # except ValueError as e:
    #     print('ValueError:', e)
    # except ZeroDivisionError as e:
    #     print('ZeroDivisionError:', e)
    # else:
    #     print('no error!')
    # finally:
    #     print('finally...')
    # print('END')
# Python的错误其实也是class，所有的错误类型都继承自BaseException，所以在使用except时需要注意的是，它不但捕获该类型的错误，还把其子类也“一网打尽”。比如：
    # try:
    #     foo()
    # except ValueError as e:
    #     print('ValueError')
    # except UnicodeError as e:
    #     print('UnicodeError')
    # 以上代码只会捕获ValueError，因为UnicodeError是ValueError的子类，如果没有UnicodeError，也不会捕获。
# 使用try...except捕获错误还有一个巨大的好处，就是可以跨越多层调用，比如函数main()调用bar()，bar()调用foo()，结果foo()出错了，这时，只要main()捕获到了，就可以处理：
    # def foo(s):
    #     return 10 / int(s)
    #
    # def bar(s):
    #     return foo(s) * 2
    #
    # def main():
    #     try:
    #         bar('0')
    #     except Exception as e:
    #         print('Error:', e)
    #     finally:
    #         print('finally...')

# 调用栈
# 如果错误没有被捕获，它就会一直往上抛，最后被Python解释器捕获，打印一个错误信息，然后程序退出。来看看err.py：
    # def foo(s):
    #     return 10 / int(s)
    #
    # def bar(s):
    #     return foo(s) * 2
    #
    # def main():
    #     bar('0')
    #
    # main()
    # 运行后，会得到如下错误信息：
    # $ python3 err.py
    # Traceback (most recent call last):
    #   File "err.py", line 11, in <module>
    #     main()
    #   File "err.py", line 9, in main
    #     bar('0')
    #   File "err.py", line 6, in bar
    #     return foo(s) * 2
    #   File "err.py", line 3, in foo
    #     return 10 / int(s)
    # ZeroDivisionError: division by zero

# 记录错误
# 如果不捕获错误，自然可以让Python解释器来打印出错误堆栈，但程序也被结束了。既然我们能捕获错误，就可以把错误堆栈打印出来，然后分析错误原因，同时，让程序继续执行下去。
# 我们可以用logging模块来记录错误信息，通过配置，logging还可以把错误记录到日志文件里，方便事后排查。比如：
    # import logging
    #
    # def foo(s):
    #     return 10 / int(s)
    #
    # def bar(s):
    #     return foo(s) * 2
    #
    # def main():
    #     try:
    #         bar('0')
    #     except Exception as e:
    #         logging.exception(e)
    #
    # main()
    # print('END')
    # 同样是出错，但程序打印完错误信息后会继续执行，并正常退出。

# 抛出错误
# 因为错误是class，捕获一个错误就是捕获到该class的一个实例。因此，错误并不是凭空产生的，而是有意创建并抛出的。Python的内置函数会抛出很多类型的错误，我们自己编写的函数也可以抛出错误。
# 只有在必要的时候才定义我们自己的错误类型。如果可以选择Python已有的内置的错误类型（比如ValueError，TypeError），尽量使用Python内置的错误类型。
# 如果要抛出错误，首先根据需要，可以定义一个错误的class，选择好继承关系，然后，用raise语句抛出一个错误的实例：
    # class FooError(ValueError):
    #     pass
    #
    # def foo(s):
    #     n = int(s)
    #     if n==0:
    #         raise FooError('invalid value: %s' % s)
    #     return 10 / n
    #
    # foo('0')
    # 执行，可以最后跟踪到我们自己定义的错误。
# raise语句如果不带参数，就会把当前错误原样抛出。
    # def foo(s):
    #     n = int(s)
    #     if n==0:
    #         raise ValueError('invalid value: %s' % s)
    #     return 10 / n
    #
    # def bar():
    #     try:
    #         foo('0')
    #     except ValueError as e:
    #         print('ValueError!')
    #         raise
    #
    # bar()
# 此外，在except中raise一个Error，还可以把一种类型的错误转化成另一种类型，只要是合理的转换逻辑就可以，但是，决不应该把一个IOError转换成毫不相干的ValueError。
    # try:
    #     10 / 0
    # except ZeroDivisionError:
    #     raise ValueError('input error!')

# 调试

# 第一种方法简单直接粗暴有效，就是用print()把可能有问题的变量打印出来看看。

# 断言
# 凡是用print()来辅助查看的地方，都可以用断言（assert）来替代：
    # def foo(s):
    #     n = int(s)
    #     assert n != 0, 'n is zero!'
    #     return 10 / n
    #
    # def main():
    #     foo('0')
    # assert的意思是，表达式n != 0应该是True，否则，根据程序运行的逻辑，后面的代码肯定会出错。
    # 如果断言失败，assert语句本身就会抛出AssertionError：
        # $ python err.py
        # Traceback (most recent call last):
        #   ...
        # AssertionError: n is zero!
    # 程序中如果到处充斥着assert，和print()相比也好不到哪去。不过，启动Python解释器时可以用-O参数来关闭assert：
    # $ python -O err.py
    # 关闭后，你可以把所有的assert语句当成pass来看。

# logging
# 把print()替换为logging是第3种方式，和assert比，logging不会抛出错误，而且可以输出到文件：
    # import logging
    # logging.basicConfig(level=logging.INFO)
    # s = '0'
    # n = int(s)
    # logging.info('n = %d' % n)
    # print(10 / n)

    # $ python err.py
    # INFO:root:n = 0
    # Traceback (most recent call last):
    #   File "err.py", line 8, in <module>
    #     print(10 / n)
    # ZeroDivisionError: division by zero
# 这就是logging的好处，它允许你指定记录信息的级别，有debug，info，warning，error等几个级别，当我们指定level=INFO时，logging.debug就不起作用了。
    # 同理，指定level=WARNING后，debug和info就不起作用了。这样一来，你可以放心地输出不同级别的信息，也不用删除，最后统一控制输出哪个级别的信息。
# logging的另一个好处是通过简单的配置，一条语句可以同时输出到不同的地方，比如console和文件。

# pdb
# 第4种方式是启动Python的调试器pdb，让程序以单步方式运行，可以随时查看运行状态。我们先准备好程序：
    # # err.py
    # s = '0'
    # n = int(s)
    # print(10 / n)

    # $ python -m pdb err.py
    # > /Users/michael/Github/learn-python3/samples/debug/err.py(2)<module>()
    # -> s = '0'

    # 以参数-m pdb启动后，pdb定位到下一步要执行的代码-> s = '0'。输入命令l来查看代码：
    # (Pdb) l
    #   1     # err.py
    #   2  -> s = '0'
    #   3     n = int(s)
    #   4     print(10 / n)
    # 输入命令n可以单步执行代码：
    # (Pdb) n
    # > /Users/michael/Github/learn-python3/samples/debug/err.py(3)<module>()
    # -> n = int(s)
    # (Pdb) n
    # > /Users/michael/Github/learn-python3/samples/debug/err.py(4)<module>()
    # -> print(10 / n)
    # 任何时候都可以输入命令p 变量名来查看变量：
    # (Pdb) p s
    # '0'
    # (Pdb) p n
    # 0
    # 输入命令q结束调试，退出程序：
    # (Pdb) q
# 这种通过pdb在命令行调试的方法理论上是万能的，但实在是太麻烦了，如果有一千行代码，要运行到第999行得敲多少命令啊。还好，我们还有另一种调试方法。

# pdb.set_trace()
# 这个方法也是用pdb，但是不需要单步执行，我们只需要import pdb，然后，在可能出错的地方放一个pdb.set_trace()，就可以设置一个断点：
    # # err.py
    # import pdb
    #
    # s = '0'
    # n = int(s)
    # pdb.set_trace() # 运行到这里会自动暂停
    # print(10 / n)
# 运行代码，程序会自动在pdb.set_trace()暂停并进入pdb调试环境，可以用命令p查看变量，或者用命令c继续运行：
    # $ python err.py
    # > /Users/michael/Github/learn-python3/samples/debug/err.py(7)<module>()
    # -> print(10 / n)
    # (Pdb) p n
    # 0
    # (Pdb) c
    # Traceback (most recent call last):
    #   File "err.py", line 7, in <module>
    #     print(10 / n)
    # ZeroDivisionError: division by zero
# 这个方式比直接启动pdb单步调试效率要高很多，但也高不到哪去。

# 单元测试
# mydict.py代码如下：
# class Dict(dict):
#     def __init__(self, **kw):
#         super().__init__(**kw)
#
#     def __getattr__(self, key):
#         try:
#             return self[key]
#         except KeyError:
#             raise AttributeError(r"'Dict' object has no attribute '%s'" % key)
#
#     def __setattr__(self, key, value):
#         self[key] = value
# 为了编写单元测试，我们需要引入Python自带的unittest模块。
# 编写单元测试时，我们需要编写一个测试类，从unittest.TestCase继承。
# 以test开头的方法就是测试方法，不以test开头的方法不被认为是测试方法，测试的时候不会被执行。
# 对每一类测试都需要编写一个test_xxx()方法。由于unittest.TestCase提供了很多内置的条件判断，我们只需要调用这些方法就可以断言输出是否是我们所期望的。
    # 最常用的断言就是assertEqual()：self.assertEqual(abs(-1), 1) # 断言函数返回的结果与1相等
# 另一种重要的断言就是期待抛出指定类型的Error，比如通过d['empty']访问不存在的key时，断言会抛出KeyError：
    # with self.assertRaises(KeyError):
    #     value = d['empty']
# 而通过d.empty访问不存在的key时，我们期待抛出AttributeError：
    # with self.assertRaises(AttributeError):
    #     value = d.empty
# 编写mydict_test.py如下：
# import unittest
# from mydict import Dict
# class TestDict(unittest.TestCase):
#     def test_init(self):
#         d = Dict(a=1, b='test')
#         self.assertEqual(d.a, 1)
#         self.assertEqual(d.b, 'test')
#         self.assertTrue(isinstance(d, dict))
#
#     def test_key(self):
#         d = Dict()
#         d['key'] = 'value'
#         self.assertEqual(d.key, 'value')
#
#     def test_attr(self):
#         d = Dict()
#         d.key = 'value'
#         self.assertTrue('key' in d)
#         self.assertEqual(d['key'], 'value')
#
#     def test_keyerror(self):
#         d = Dict()
#         with self.assertRaises(KeyError):
#             value = d['empty']
#
#     def test_attrerror(self):
#         d = Dict()
#         with self.assertRaises(AttributeError):
#             value = d.empty

# 运行单元测试
# 一旦编写好单元测试，我们就可以运行单元测试。最简单的运行方式是在mydict_test.py的最后加上两行代码：
    # if __name__ == '__main__':
    #     unittest.main()
# 样就可以把mydict_test.py当做正常的python脚本运行：
    # $ python mydict_test.py
# 另一种方法是在命令行通过参数-m unittest直接运行单元测试：
    # $ python -m unittest mydict_test

# setUp与tearDown
# 可以在单元测试中编写两个特殊的setUp()和tearDown()方法。这两个方法会分别在每调用一个测试方法的前后分别被执行。
# setUp()和tearDown()方法有什么用呢？设想你的测试需要启动一个数据库，这时，就可以在setUp()方法中连接数据库，在tearDown()方法中关闭数据库，这样，不必在每个测试方法中重复相同的代码：
# class TestDict(unittest.TestCase):
#     def setUp(self):
#         print('setUp...')
#
#     def tearDown(self):
#         print('tearDown...')

# 文档测试
# 这些代码与其他说明可以写在注释中，然后，由一些工具来自动生成文档。既然这些代码本身就可以粘贴出来直接运行，那么，可不可以自动执行写在注释中的这些代码呢？答案是肯定的。
# Python内置的“文档测试”（doctest）模块可以直接提取注释中的代码并执行测试。
# doctest严格按照Python交互式命令行的输入和输出来判断测试结果是否正确。只有测试异常的时候，可以用...表示中间一大段烦人的输出。
# 让我们用doctest来测试上次编写的Dict类：
    # # mydict2.py
    # class Dict(dict):
    #     '''
    #     Simple dict but also support access as x.y style.
    #
    #     >>> d1 = Dict()
    #     >>> d1['x'] = 100
    #     >>> d1.x
    #     100
    #     >>> d1.y = 200
    #     >>> d1['y']
    #     200
    #     >>> d2 = Dict(a=1, b=2, c='3')
    #     >>> d2.c
    #     '3'
    #     >>> d2['empty']
    #     Traceback (most recent call last):
    #         ...
    #     KeyError: 'empty'
    #     >>> d2.empty
    #     Traceback (most recent call last):
    #         ...
    #     AttributeError: 'Dict' object has no attribute 'empty'
    #     '''
    #     def __init__(self, **kw):
    #         super(Dict, self).__init__(**kw)
    #
    #     def __getattr__(self, key):
    #         try:
    #             return self[key]
    #         except KeyError:
    #             raise AttributeError(r"'Dict' object has no attribute '%s'" % key)
    #
    #     def __setattr__(self, key, value):
    #         self[key] = value
    #
    # if __name__=='__main__':
    #     import doctest
    #     doctest.testmod()
# 运行python mydict2.py：
# $ python mydict2.py

