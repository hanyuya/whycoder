# 高阶函数
# 把函数作为参数传入，这样的函数称为高阶函数，函数式编程就是指这种高度抽象的编程范式。
    # 变量可以指向函数: 函数本身也可以赋值给变量，即：变量可以指向函数。
    # f = abs
    # print(f(-10))  # 10

    # 函数名也是变量，函数名其实就是指向函数的变量。
    # print(abs(-10))  # 10

    # 传入函数: 我们可以把函数作为参数传入另一个函数，这种函数称为高阶函数。
    # def add(x, y, f):
    #     return f(x) + f(y)

# Python内建了map()和reduce()函数。
    # map()函数接收两个参数，一个是函数，一个是可迭代对象，map将传入的函数依次作用到序列的每个元素，并把结果作为新的Iterator返回。
    # def f(x):
    #     return x * x
    # r = map(f, [1, 2, 3, 4, 5, 6, 7, 8, 9])
    # print(list(r))  # [1, 4, 9, 16, 25, 36, 49, 64, 81]
    # list(map(str, [1, 2, 3, 4, 5, 6, 7, 8, 9]))  # ['1', '2', '3', '4', '5', '6', '7', '8', '9']

    # reduce()函数接收两个参数，一个是函数，一个是序列，reduce把函数作用在序列的元素上，这个函数必须接收两个参数，reduce把结果继续和序列的下一个元素做累积计算。
    # from functools import reduce
    # def add(x, y):
    #     return x + y
    # r = reduce(add, [1, 3, 5, 7, 9])
    # print(r)  # 25

    # 利用map()和reduce()函数，可以写出更复杂的函数式程序。
    # 把str转换为int的函数
    # from functools import reduce
    # def fn(x, y):
    #     return x * 10 + y
    #
    # def char2num(s):
    # digits = {'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
    #     return digits[s]
    #
    # reduce(fn, map(char2num, '13579'))  # 13579

    # 还可以用lambda函数进一步简化成一行代码
    # from functools import reduce
    # reduce(lambda x, y: x * 10 + y, map(lambda s: {'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}[s], '13579'))  # 13579

# Python内建的filter()函数用于过滤序列。
    # filter()也接收一个函数和一个序列。和map()不同的是，filter()把传入的函数依次作用于每个元素，然后根据返回值是True还是False决定保留还是丢弃该元素。
    # def is_odd(n):
    #     return n % 2 == 1
    # r = filter(is_odd, [1, 2, 3, 4, 5, 6, 7, 8, 9])
    # print(list(r))  # [1, 3, 5, 7, 9]

# Python内置的sorted()函数就可以对list进行排序。
    # sorted()也是一个高阶函数，它可以接收一个可迭代对象，并返回一个新的排序后的列表。
    # sorted([36, 5, -12, 9, -21])  # [-21, -12, 5, 9, 36]
    # sorted(['bob', 'about', 'Zoo', 'Credit'])  # ['Credit', 'Zoo', 'about', 'bob']
    # sorted([36, 5, -12, 9, -21], key=abs)  # [5, 9, -12, -21, 36]
    # sorted(['bob', 'about', 'Zoo', 'Credit'], key=str.lower)  # ['about', 'bob', 'Credit', 'Zoo']
    # 要进行反向排序，不必改动key函数，可以传入第三个参数reverse=True：
    # sorted([36, 5, -12, 9, -21], key=abs, reverse=True)  # [36, -21, -12, 9, 5]

# 返回函数: 高阶函数除了可以接受函数作为参数外，还可以返回一个函数。
    # 我们可以定义一个函数，它接受一个函数作为参数，并返回一个新的函数。
    # def lazy_sum(*args):
    #     def sum():
    #         return sum(args)
    #     return sum
    # f = lazy_sum(1, 3, 5, 7, 9)
    # print(f())  # 25
    # 在函数lazy_sum中又定义了函数sum，并且，内部函数sum可以引用外部函数lazy_sum的参数和局部变量，当lazy_sum返回函数sum时，相关参数和变量都保存在返回的函数中，这种称为“闭包（Closure）”的程序结构拥有极大的威力。
    # 当我们调用lazy_sum()时，每次调用都会返回一个新的函数，即使传入相同的参数。

# 闭包
    # 返回的函数在其定义内部引用了局部变量args，所以，当一个函数返回了一个函数后，其内部的局部变量还被新函数引用，所以，闭包用起来简单，实现起来可不容易。
    # 返回的函数并没有立刻执行，而是直到调用了f()才执行。
    # 返回函数不要引用任何循环变量，或者后续会发生变化的变量。
    # def count():
    #     fs = []
    #     for i in range(1, 4):
    #         def f():
    #              return i*i
    #         fs.append(f)
    #     return fs
    #
    # f1, f2, f3 = count()
    # print(f1(), f2(), f3())  # 9 9 9
    # 如果一定要引用循环变量怎么办？方法是再创建一个函数，用该函数的参数绑定循环变量当前的值，无论该循环变量后续如何更改，已绑定到函数参数的值不变
    # def count():
    #     def f(j):
    #         def g():
    #             return j*j
    #         return g
    #     fs = []
    #     for i in range(1, 4):
    #         fs.append(f(i)) # f(i)立刻被执行，因此i的当前值被传入f()
    #     return fs
    #
    # f1, f2, f3 = count()
    # print(f1(), f2(), f3())  # 1 4 9

# nonlocal关键字
    # 使用闭包，就是内层函数引用了外层函数的局部变量。如果只是读外层变量的值，我们会发现返回的闭包函数调用一切正常，但是，如果想修改外层变量的值，就会发现修改失败。
    # 使用闭包时，对外层变量赋值前，需要先使用nonlocal声明该变量不是当前函数的局部变量。
    # def inc():
    #     x = 0
    #     def fn():
    #         # nonlocal x
    #         x = x + 1
    #         return x
    #     return fn
    # 原因是x作为局部变量并没有初始化，直接计算x+1是不行的。
    # 正确的做法是用nonlocal关键字声明要修改的外层变量。

# 匿名函数
    # 除了使用def关键字定义函数外，Python还允许使用lambda关键字创建匿名函数。
    # 关键字lambda表示匿名函数，冒号前面的x表示函数参数。
    # 匿名函数也是一个函数对象，也可以把匿名函数赋值给一个变量，再利用变量来调用该函数
    # f = lambda x: x * x

# 装饰器
    # 装饰器是Python的另一种高阶函数。
    # 函数对象有一个__name__属性（注意：是前后各两个下划线），可以拿到函数的名字：
    # def foo():
    #     pass
    # print(foo.__name__)  # foo
    # 我们要增强foo()函数的功能，比如，在函数调用前后自动打印日志，但又不希望修改foo()函数的定义，这种在代码运行期间动态增加功能的方式，称之为“装饰器”（Decorator）。
    # 装饰器是通过@语法实现的，其基本语法如下：
    # @decorator
    # def func():
    #     pass
    # 这里，decorator是装饰器函数，func是被装饰的函数。
    # 装饰器函数的定义如下：
    # def decorator(func):
    #     def wrapper(*args, **kw):
    #         print('call %s():' % func.__name__)
    #         return func(*args, **kw)
    #     return wrapper
    # 装饰器函数的作用就是在调用func()函数前后自动打印日志。
    # 我们可以定义一个@decorator的装饰器，把它放在foo()函数的定义处：
    # @decorator
    # def foo():
    #     pass
    # 这样，调用foo()函数时，会自动打印日志：
    # foo()
    # call foo():

    # 我们还可以把多个装饰器叠加使用：
    # @decorator1
    # @decorator2
    # def foo():
    #     pass
    # 这样，先执行decorator2，再执行decorator1。

    # 装饰器不但可以打印日志，还可以做很多其他有用的事情，比如，检查参数，限定函数的执行时间，统计函数的执行次数，甚至修改函数的返回值。
    # 如果decorator本身需要传入参数，那就需要编写一个返回decorator的高阶函数，写出来会更复杂。比如，要自定义log的文本：
    # def my_decorator(text):
    #     def decorator(func):
    #         def wrapper(*args, **kw):
    #             print('%s %s():' % (text, func.__name__))
    #             return func(*args, **kw)
    #         return wrapper
    #     return decorator
    # 现在，我们可以自定义log的文本：
    # @my_decorator('execute')
    # def foo():
    #     pass
    # 这样，调用foo()函数时，会自动打印日志：
    # foo()
    # execute foo():

    # 因为函数也是对象，它有__name__等属性，但你去看经过decorator装饰之后的函数，它们的__name__已经从原来的'now'变成了'wrapper'：
    # print(foo.__name__)  # wrapper
    # 因为返回的那个wrapper()函数名字就是'wrapper'，所以，需要把原始函数的__name__等属性复制到wrapper()函数中，否则，有些依赖函数签名的代码执行就会出错。
    # 不需要编写wrapper.__name__ = func.__name__这样的代码，Python内置的functools.wraps就是干这个事的，所以，一个完整的decorator的写法如下：
    # import functools
    # def log(func):
    #     @functools.wraps(func)
    #     def wrapper(*args, **kw):
    #         print('call %s():' % func.__name__)
    #         return func(*args, **kw)
    #     return wrapper
    # 这样，无论调用原始函数还是装饰后的函数，它们的__name__都不会丢失。
    # 带参数的decorator:
    # import functools
    # def log(text):
    #     def decorator(func):
    #         @functools.wraps(func)
    #         def wrapper(*args, **kw):
    #             print('%s %s():' % (text, func.__name__))
    #             return func(*args, **kw)
    #         return wrapper
    #     return decorator

# 偏函数
    # 偏函数（Partial function）是指固定一个函数的某些参数，返回一个新的函数，这样返回的新函数可以固定参数的部分，从而简化调用。
    # import functools
    # int2 = functools.partial(int, base=2)
    # int2('1000000')  # 64