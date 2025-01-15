# 函数
    # 内置函数 len() 返回对象（字符串、列表、元组等）的长度。
    # 内置函数 type() 返回对象类型。
    # 内置函数 print() 输出字符串。
    # 内置函数 input() 输入字符串，并返回输入值，输入值类型为字符串。
    # 内置函数 range() 创建一个整数序列。
    # 内置函数 sorted() 对列表进行排序。
    # 内置函数 max() 和 min() 返回列表中的最大值和最小值。
    # 内置函数 abs() 返回数字的绝对值。
    # 内置函数 round() 四舍五入。
    # 内置函数 isinstance() 判断对象是否为指定类型。isinstance(x, (int, float)): 判断x是否为整数或浮点数。

# 数据类型转换
    # 内置函数 str() 将对象转换为字符串。
    # 内置函数 int() 将字符串、浮点数转换为整数。
    # 内置函数 float() 将字符串、整数转换为浮点数。
    # 内置函数 bool() 将对象转换为布尔值。

# 自定义函数
    # 定义函数的语法如下：
    # def 函数名(参数列表):
    #     函数体
    # 函数名：自定义函数的名称，可以取任何有效的标识符。
    # 参数列表：函数参数的列表，可以为空。
    # 函数体：函数执行的语句块，可以包含多个语句。
    # 函数的返回值用return语句返回。
    # 如果没有return语句，函数执行完毕后也会返回结果，只是结果为None。return None可以简写为return

# 空函数可以用pass语句占位。(pass可以用来作为占位符，比如现在还没想好怎么写函数的代码，就可以先放一个pass，让代码能运行起来。)
    # 举例：
    # def nop():
    #     pass
    # 函数可以返回多个值, Python的函数返回多值其实就是返回一个tuple。

# 默认参数：定义函数时，可以给参数指定默认值，这样，当调用函数时，可以省略参数，使用默认值。
    # 必选参数在前，默认参数在后，否则Python的解释器会报错
    # 默认参数必须指向不变对象，可以用None这个不变对象作为默认值。
    # 举例：def greet(name, msg="Hello,"):

# 可变参数：传入的参数个数是可变的，可以是1个、2个到任意个，还可以是0个。
    # 定义可变参数和定义一个list或tuple参数相比，仅仅在参数前面加了一个*号
    # Python允许你在list或tuple前面加一个*号，把list或tuple的元素变成可变参数
    # 定义可变参数的语法如下：
    # def func(arg, *args):
    #     函数体
    # 参数arg：必选参数，指定参数个数。
    # 参数*args：可变参数，收集剩余的参数，参数名可以自定义。


# 关键字参数：关键字参数允许你传入0个或任意个含参数名的参数，这些关键字参数在函数内部自动组装为一个dict(tuple)。
    # 举例：
    # def person(name, age, **kw):
    #     print('name:', name, 'age:', age, 'other:', kw)
    # 调用:
    # person('Alice', 25, city='Shanghai', job='Doctor')
    # extra = {'city': 'Beijing', 'job': 'Engineer'}
    # person('Jack', 24, city=extra['city'], job=extra['job'])
    # person('Jack', 24, **extra)

# 命名关键字参数：函数的调用者可以传入任意不受限制的关键字参数。至于到底传入了哪些，就需要在函数内部通过kw检查。
    # 和关键字参数**kw不同，命名关键字参数需要一个特殊分隔符*，*后面的参数被视为命名关键字参数。
    # 举例:
    # def person(name, age, *, city, job):
    #     print('name:', name, 'age:', age, 'city:', city, 'job:', job)
    # 调用:
    # person('Alice', 25, city='Shanghai', job='Doctor')
    # 如果函数定义中已经有了一个可变参数，后面跟着的命名关键字参数就不再需要一个特殊分隔符*了：
    # def person(name, age, *args, city, job):
    #     print('name:', name, 'age:', age, 'args:', args, 'city:', city, 'job:', job)
    # 调用:
    # person('Alice', 25, 'Engineer', city='Shanghai', job='Doctor')
    # 命名关键字参数可以有缺省值，从而简化调用：
    # def person(name, age, *, city='Beijing', job):
    #     print(name, age, city, job)
    # 使用命名关键字参数时，要特别注意，如果没有可变参数，就必须加一个*作为特殊分隔符。如果缺少*，Python解释器将无法识别位置参数和命名关键字参数：
    # def person(name, age, city, job):
    #     # 缺少 *，city和job被视为位置参数
    #     pass

# 参数组合
    # 参数定义的顺序：必选参数、默认参数、可变参数、命名关键字参数和关键字参数。
    # 举例：
    # def person(name, age, city='Beijing', *args, job, **kw):
    #     print(name, age, city, args, job, kw)
    # 对于任意函数，都可以通过类似func(*args, **kw)的形式调用它，无论它的参数是如何定义的。

# 递归函数：函数可以调用自己，这种函数称为递归函数。
    # 举例：
    # def fact(n):
    #     if n == 1:
    #         return 1
    #     else:
    #         return n * fact(n-1)
    # 调用：fact(5)
    # 递归函数的优点是逻辑简单，缺点是效率低，因为函数调用会带来栈的压入和弹出，内存开销大。
    # 要改进递归函数的效率，通常使用尾递归优化。
    # 尾递归是指，在函数返回的时候，调用自身本身，并且，return语句不能包含表达式。这样，编译器或者解释器就可以把尾递归做优化，使递归本身无论调用多少次，都只占用一个栈帧，不会出现栈溢出的情况。
    # 大多数编程语言没有针对尾递归做优化，Python解释器也没有做优化，所以，即使把上面的fact(n)函数改成尾递归方式，也会导致栈溢出。