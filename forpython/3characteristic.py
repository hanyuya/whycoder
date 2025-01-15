# 切片（Slice）操作符
# 切片操作符的语法如下：
# [start:stop:step]
# 其中，start表示切片开始的索引（包含），stop表示切片结束的索引（不包含），step表示切片的步长。
# 如果省略start，则默认为0；如果省略stop，则默认为列表长度；如果省略step，则默认为1。
# 切片操作符的结果是一个新的列表，包含原列表中指定范围内的元素。
# 字符串'xxx'也可以看成是一种list，每个元素就是一个字符。因此，字符串也可以用切片操作，只是操作结果仍是字符串：
# s = 'hello world'
# print(s[0:5])  # 输出 'hello'

# 迭代（Iteration）
# Python的for循环抽象程度要高于C的for循环，因为Python的for循环不仅可以用在list或tuple上，还可以作用在其他可迭代对象上。
# 迭代dict时，只迭代key。如果要迭代value，可以用for value in d.values()，如果要同时迭代key和value，可以用for k, v in d.items()。
# 迭代字符串时，可以用for c in s，其中c表示字符串中的每一个字符。
# 判断对象是否是可迭代对象（Iterable），通过collections.abc模块的Iterable类型判断：
    # from collections.abc import Iterable
    # if isinstance(obj, Iterable):
    #     # obj is iterable
    # else:
    #     # obj is not iterable
    # isinstance('abc', Iterable) # str是否可迭代
    # isinstance([1, 2, 3], Iterable) # list是否可迭代
    # isinstance(123, Iterable) # 数字是否可迭代
# Python内置的enumerate函数可以把一个list变成索引-元素对，这样就可以在for循环中同时迭代索引和元素本身：
    # for i, value in enumerate(['A', 'B', 'C']):
    #     print(i, value)  # 输出 0 A 1 B 2 C

# 列表生成式（List Comprehension）
# 列表生成式是Python内置的非常简单却强大的可以用来创建list的生成式。
# 语法格式如下：
# [expression for item in iterable if condition]
# 表达式可以是任意合法的Python表达式，而item和iterable分别表示待生成元素和待迭代对象。
# 条件（condition）是可选的，表示对每个元素进行过滤，只有满足条件的元素才会被包含在生成的列表中。
# 举例如下：
# 1. 计算1到10的平方
# squares = [x**2 for x in range(1, 11)]
# print(squares)  # 输出 [1, 4, 9, 16, 25, 36, 49, 64, 81, 100]
# 2. 筛选出列表中大于5的元素
# numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
# 还可以使用两层循环，可以生成全排列：
    # permutations = [m + n for m in 'ABC' for n in 'XYZ']
# for循环其实可以同时使用两个甚至多个变量，比如dict的items()可以同时迭代key和value：
    # d = {'x': 'A', 'y': 'B', 'z': 'C' }
    # [k + '=' + v for k, v in d.items()]
    # ['y=B', 'x=A', 'z=C']
# 在一个列表生成式中，for前面的if ... else是表达式，而for后面的if是过滤条件，不能带else。
    # [x if x % 2 == 0 else -x for x in range(1, 11)]
    # [-1, 2, -3, 4, -5, 6, -7, 8, -9, 10]

# 生成器表达式（Generator Expression）
# 生成器表达式是列表生成式的一种简化写法，只用一对括号。
# 语法格式如下：
# (expression for item in iterable if condition)
# 与列表生成式的区别在于，生成器表达式不创建完整的列表，而是返回一个生成器对象。
# 调用生成器对象的next()方法，可以获取生成器的下一个元素，直到遇到StopIteration错误时结束。
# 举例如下：
# 1. 计算1到10的平方
# squares = (x**2 for x in range(1, 11))
# print(next(squares))  # 输出 1
# print(next(squares))  # 输出 4
# print(next(squares))  # 输出 9
# 2. 筛选出列表中大于5的元素
# numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
# 正确的方法是使用for循环，因为generator也是可迭代对象：
    # filtered = (x for x in numbers if x > 5)
    # for x in filtered:
    #     print(x)  # 输出 6 7 8 9 10
# 如果一个函数定义中包含yield关键字，那么这个函数就不再是一个普通函数，而是一个generator函数，调用一个generator函数将返回一个generator。
# generator函数和普通函数的执行流程不一样。
    # 普通函数是顺序执行，遇到return语句或者最后一行函数语句就返回。
    # generator的函数，在每次调用next()的时候执行，遇到yield语句返回，再次执行时从上次返回的yield语句处继续执行。
    # def odd():
    #   print('step 1')
    #   yield 1
    #   print('step 2')
    #   yield(3)
    #   print('step 3')
    #   yield(5)
    # o = odd()
    # print(next(o))  # 输出 step 1 1
    # print(next(o))  # 输出 step 2 3
    # print(next(o))  # 输出 step 3 5
    # print(next(o))  # 输出 StopIteration
    # for n in odd():
    #     print(n)  # 输出 1 3 5

# 迭代器（Iterator）
# 可以直接作用于for循环的数据类型有以下几种：
    # 一类是集合数据类型，如list、tuple、dict、set、str等；
    # 一类是generator，包括生成器和带yield的generator function。
# 这些可以直接作用于for循环的对象统称为可迭代对象：Iterable。
# 可以使用isinstance()判断一个对象是否是Iterable对象：
    # from collections.abc import Iterable
    # isinstance([1, 2, 3], Iterable) # True
    # isinstance('abc', Iterable) # True
    # isinstance({}, Iterable) # True
    # isinstance((x for x in range(10)), Iterable)
    # isinstance(100, Iterable) # False
# 生成器不但可以作用于for循环，还可以被next()函数不断调用并返回下一个值，直到最后抛出StopIteration错误表示无法继续返回下一个值了。
# 可以被next()函数调用并不断返回下一个值的对象称为迭代器：Iterator。
# 可以使用isinstance()判断一个对象是否是Iterator对象：
    # from collections.abc import Iterator
    # isinstance((x for x in range(10)), Iterator) # True
    # isinstance([], Iterator) # False
    # isinstance({}, Iterator) # False
    # isinstance('abc', Iterator) # False
# 生成器都是Iterator对象，但list、dict、str虽然是Iterable，却不是Iterator。
# 把list、dict、str等Iterable变成Iterator可以使用iter()函数：
    # iter_list = iter([1, 2, 3])
    # iter_dict = iter({'x': 1, 'y': 2, 'z': 3})
    # iter_str = iter('hello')
    # isinstance(iter_list, Iterator) # True
    # isinstance(iter_dict, Iterator) # True
    # isinstance(iter_str, Iterator) # True
# 也可以使用next()函数获取一个Iterator的下一个值，直到遇到StopIteration错误表示无法继续返回下一个值了：
    # while True:
    #     try:
    #         n = next(iter_list)
    #         print(n)
    #     except StopIteration:
    #         break
    # 输出 1 2 3
# Python的Iterator对象表示的是一个数据流，Iterator对象可以被next()函数调用并不断返回下一个数据，直到没有数据时抛出StopIteration错误。
# 数据流看做是一个有序序列，但我们却不能提前知道序列的长度，只能不断通过next()函数实现按需计算下一个数据，所以Iterator的计算是惰性的，只有在需要返回下一个数据时它才会计算。
# Iterator甚至可以表示一个无限大的数据流，例如全体自然数。而使用list是永远不可能存储全体自然数的。
# Python的for循环本质上就是通过不断调用next()函数实现的。