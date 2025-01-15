# 使用__slots__
# 当我们定义了一个class，创建了一个class的实例后，我们可以给该实例绑定任何属性和方法，这就是动态语言的灵活性。
    # 尝试给实例绑定一个属性：
    # class Person:
    #     pass
    # p = Person()
    # p.name = 'Alice'
    # print(p.name)  # output: Alice
    # 尝试给实例绑定一个方法：
    # class Person:
    #     def say_hello(self):
    #         print('Hello, world!')
    # p = Person()
    # p.say_hello()  # output: Hello, world!
# 给一个实例绑定的方法，对另一个实例是不起作用的，为了给所有实例都绑定方法，可以给class绑定方法：
    # >>> def set_score(self, score):
    # ...     self.score = score
    # ...
    # >>> Person.set_score = set_score
# set_score方法可以直接定义在class中，但动态绑定允许我们在程序运行的过程中动态给class加上功能，这在静态语言中很难实现。
# 但是，Python允许在定义class的时候，定义一个特殊的__slots__变量，来限制该class实例能添加的属性：
    # class Person:
    #     __slots__ = ('name', 'age') # 用tuple定义允许绑定的属性名称
    # p = Person()
    # p.score = 99
    # AttributeError: 'Person' object has no attribute'score'
# 使用__slots__要注意，__slots__定义的属性仅对当前类实例起作用，对继承的子类是不起作用的：
    # class Student(Person):
    #     pass
    # s = Student()
    # s.score = 90
# 除非在子类中也定义__slots__，这样，子类实例允许定义的属性就是自身的__slots__加上父类的__slots__。

# 使用@property
# 既能检查参数，又可以用类似属性这样简单的方式来访问类的变量。
# 对于类的方法，装饰器一样起作用。Python内置的@property装饰器就是负责把一个方法变成属性调用的：
    # class Student(object):
    #     @property
    #     def score(self):
    #         return self._score
    #
    #     @score.setter
    #     def score(self, value):
    #         if not isinstance(value, int):
    #             raise ValueError('score must be an integer!')
    #         if value < 0 or value > 100:
    #             raise ValueError('score must between 0 ~ 100!')
    #         self._score = value
    # 把一个getter方法变成属性，只需要加上@property就可以了，此时，@property本身又创建了另一个装饰器@score.setter，负责把一个setter方法变成属性赋值，于是，我们就拥有一个可控的属性操作：
    # >>> s = Student()
    # >>> s.score = 60 # OK，实际转化为s.score(60)
    # >>> s.score # OK，实际转化为s.score()
    # 60
    # >>> s.score = 9999
    # Traceback (most recent call last):
    #   ...
    # ValueError: score must between 0 ~ 100!
    # @property，我们在对实例属性操作的时候，就知道该属性很可能不是直接暴露的，而是通过getter和setter方法来实现的。
    # 还可以定义只读属性，只定义getter方法，不定义setter方法就是一个只读属性：
    # class Student(object):
    #   @property
    #   def birth(self):
    #       return self._birth
    #
    #   @birth.setter
    #   def birth(self, value):
    #       self._birth = value
    #
    #   @property
    #   def age(self):
    #       return 2015 - self._birth
    # birth是可读写属性，而age就是一个只读属性，因为age可以根据birth和当前时间计算出来。
    # 特别注意：属性的方法名不要和实例变量重名。
    # class Student(object):
    #     # 方法名称和实例变量均为birth:
    #     @property
    #     def birth(self):
    #         return self.birth

# 多重继承
# Python支持多重继承，一个子类可以同时获得多个父类的所有功能。
# 多重继承的语法格式如下：
    # class SubClass(Parent1, Parent2, ...):
    #     # 子类定义
# MixIn
# 在设计类的继承关系时，通常，主线都是单一继承下来的，例如，Ostrich继承自Bird。但是，如果需要“混入”额外的功能，通过多重继承就可以实现，比如，让Ostrich除了继承自Bird外，再同时继承Runnable。这种设计通常称之为MixIn。
# 为了更好地看出继承关系，我们把Runnable和Flyable改为RunnableMixIn和FlyableMixIn。类似的，你还可以定义出肉食动物CarnivorousMixIn和植食动物HerbivoresMixIn，让某个动物同时拥有好几个MixIn：
# MixIn的语法格式如下：
    # class RunnableMixIn(object):
    #     def run(self):
    #         print('Running...')
    #
    # class Ostrich(Bird, RunnableMixIn):
    #     pass
# 这样，Ostrich就同时获得了Bird和RunnableMixIn的所有功能。
# 注意，MixIn的目的只是给类增加功能，并不是为了被继承，因此，不要把MixIn设计成可以被实例化的类。
# MixIn的目的就是给一个类增加多个功能，这样，在设计类的时候，我们优先考虑通过多重继承来组合多个MixIn的功能，而不是设计多层次的复杂的继承关系。
# Python自带的很多库也使用了MixIn。举个例子，Python自带了TCPServer和UDPServer这两类网络服务，而要同时服务多个用户就必须使用多进程或多线程模型，这两种模型由ForkingMixIn和ThreadingMixIn提供。
# 通过组合，我们就可以创造出合适的服务来。这样一来，我们不需要复杂而庞大的继承链，只要选择组合不同的类的功能，就可以快速构造出所需的子类。
    # 比如，编写一个多进程模式的TCP服务，定义如下：
    # class MyTCPServer(TCPServer, ForkingMixIn):
    #     pass
    # 编写一个多线程模式的UDP服务，定义如下：
    # class MyUDPServer(UDPServer, ThreadingMixIn):
    #     pass
    # 如果你打算搞一个更先进的协程模型，可以编写一个CoroutineMixIn：
    # class MyTCPServer(TCPServer, CoroutineMixIn):
    #     pass

# 定制类
# 看到类似__slots__这种形如__xxx__的变量或者函数名就要注意，这些在Python中是有特殊用途的。
# Python的class中还有许多这样有特殊用途的函数，可以帮助我们定制类。
# __str__    print()
    # >>> class Student(object):
    # ...     def __init__(self, name):
    # ...         self.name = name
    # ...     def __str__(self):
    # ...         return 'Student object (name: %s)' % self.name
    # ...
    # >>> print(Student('Michael'))
    # Student object (name: Michael)
# __repr__   直接输出变量名
    # >>> class Student(object):
    # ...     def __init__(self, name):
    # ...         self.name = name
    # ...     def __repr__(self):
    # ...         return 'Student object (name: %s)' % self.name
    # ...
    # >>> s = Student('Michael')
    # >>> print(s)
    # Student object (name: Michael)
    # >>> s
    # Student object (name: Michael)
# __iter__
# 如果一个类想被用于for ... in循环，类似list或tuple那样，就必须实现一个__iter__()方法，
#   该方法返回一个迭代对象，然后，Python的for循环就会不断调用该迭代对象的__next__()方法拿到循环的下一个值，直到遇到StopIteration错误时退出循环。
    # class Fib(object):
    #     def __init__(self):
    #         self.a, self.b = 0, 1 # 初始化两个计数器a，b
    #
    #     def __iter__(self):
    #         return self # 实例本身就是迭代对象，故返回自己
    #
    #     def __next__(self):
    #         self.a, self.b = self.b, self.a + self.b # 计算下一个值
    #         if self.a > 100000: # 退出循环的条件
    #             raise StopIteration()
    #         return self.a # 返回下一个值
# __getitem__
# 如果一个类想被用于索引操作，类似list那样，就必须实现一个__getitem__()方法，该方法根据索引值返回对应的值。
    # class Fib(object):
    #     def __getitem__(self, n):
    #         a, b = 1, 1
    #         for x in range(n):
    #             a, b = b, a + b
    #         return a
# 但是list有个神奇的切片方法，对于Fib却报错。原因是__getitem__()传入的参数可能是一个int，也可能是一个切片对象slice，所以要做判断：
    # class Fib(object):
    #     def __getitem__(self, n):
    #         if isinstance(n, int): # n是索引
    #             a, b = 1, 1
    #             for x in range(n):
    #                 a, b = b, a + b
    #             return a
    #         if isinstance(n, slice): # n是切片
    #             start = n.start
    #             stop = n.stop
    #             if start is None:
    #                 start = 0
    #             a, b = 1, 1
    #             L = []
    #             for x in range(stop):
    #                 if x >= start:
    #                     L.append(a)
    #                 a, b = b, a + b
    #             return L
# 但是没有对step参数作处理，也没有对负数作处理，所以，要正确实现一个__getitem__()还是有很多工作要做的。
# 与之对应的是__setitem__()方法，把对象视作list或dict来对集合赋值。最后，还有一个__delitem__()方法，用于删除某个元素。
# 总之，通过上面的方法，我们自己定义的类表现得和Python自带的list、tuple、dict没什么区别，这完全归功于动态语言的“鸭子类型”，不需要强制继承某个接口。

# __getattr__
# 正常情况下，当我们调用类的方法或属性时，如果不存在，就会报错。
# 要避免这个错误，除了可以加上一个score属性外，Python还有另一个机制，那就是写一个__getattr__()方法，动态返回一个属性。修改如下：
    # class Student(object):
    #     def __init__(self, name):
    #         self.name = name
    #
    #     def __getattr__(self, attr):
    #         if attr == 'score':
    #             return 99
# 这样，当调用不存在的属性时，Python解释器会试图调用__getattr__()方法，动态返回score属性。
# 返回函数也是完全可以的：
    # class Student(object):
    #     def __init__(self, name):
    #         self.name = name
    #
    #     def __getattr__(self, attr):
    #         if attr == 'age':
    #             return lambda: 2015 - 1999
    #         raise AttributeError('\'Student\' object has no attribute \'%s\'' % attr)
# 注意，只有在没有找到属性的情况下，才调用__getattr__()方法，已有的属性，比如name，不会在__getattr__()中查找。
# 任意调用如s.abc都会返回None，这是因为我们定义的__getattr__默认返回就是None。要让class只响应特定的几个属性，我们就要按照约定，抛出AttributeError的错误。
# 这实际上可以把一个类的所有属性和方法调用全部动态化处理了，不需要任何特殊手段。

# __call__
# 一个对象实例可以有自己的属性和方法，当我们调用实例方法时，我们用instance.method()来调用。能不能直接在实例本身上调用呢？在Python中，答案是肯定的。
# 任何类，只需要定义一个__call__()方法，就可以直接对实例进行调用。请看示例：
    # class Student(object):
    #     def __init__(self, name):
    #         self.name = name
    #
    #     def __call__(self):
    #         print('My name is %s.' % self.name)
# __call__()还可以定义参数。对实例进行直接调用就好比对一个函数进行调用一样，所以你完全可以把对象看成函数，把函数看成对象，因为这两者之间本来就没啥根本的区别。
# 如果你把对象看成函数，那么函数本身其实也可以在运行期动态创建出来，因为类的实例都是运行期创建出来的，这么一来，我们就模糊了对象和函数的界限。
# 那么，怎么判断一个变量是对象还是函数呢？其实，更多的时候，我们需要判断一个对象是否能被调用，能被调用的对象就是一个Callable对象，比如函数和我们上面定义的带有__call__()的类实例：
    # 通过callable()函数，我们就可以判断一个对象是否是“可调用”对象。
    # >>> callable(Student())
    # True
    # >>> callable(max)
    # True
    # >>> callable([1, 2, 3])
    # False
    # >>> callable(None)
    # False
    # >>> callable('str')
    # False

# 使用枚举类
# 为这样的枚举类型定义一个class类型，然后，每个常量都是class的一个唯一实例。Python提供了Enum类来实现这个功能：
    # from enum import Enum
    # Month = Enum('Month', ('Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'))
    # 这样我们就获得了Month类型的枚举类，可以直接使用Month.Jan来引用一个常量，或者枚举它的所有成员：
    # for name, member in Month.__members__.items():
    #     print(name, '=>', member, ',', member.value)
    # value属性则是自动赋给成员的int常量，默认从1开始计数。
# 如果需要更精确地控制枚举类型，可以从Enum派生出自定义类：
    # @unique装饰器可以帮助我们检查保证没有重复值。
    # from enum import Enum, unique
    #
    # @unique
    # class Weekday(Enum):
    #     Sun = 0 # Sun的value被设定为0
    #     Mon = 1
    #     Tue = 2
    #     Wed = 3
    #     Thu = 4
    #     Fri = 5
    #     Sat = 6
# 访问这些枚举类型可以有若干种方法：
    # >>> day1 = Weekday.Mon
    # >>> print(day1)
    # Weekday.Mon
    # >>> print(Weekday.Tue)
    # Weekday.Tue
    # >>> print(Weekday['Tue'])
    # Weekday.Tue
    # >>> print(Weekday.Tue.value)
    # 2
    # >>> print(day1 == Weekday.Mon)
    # True
    # >>> print(day1 == Weekday.Tue)
    # False
    # >>> print(Weekday(1))
    # Weekday.Mon
    # >>> print(day1 == Weekday(1))
    # True
    # >>> Weekday(7)
    # Traceback (most recent call last):
    #   ...
    # ValueError: 7 is not a valid Weekday
    # >>> for name, member in Weekday.__members__.items():
    # ...     print(name, '=>', member)
    # ...
    # Sun => Weekday.Sun
    # Mon => Weekday.Mon
    # Tue => Weekday.Tue
    # Wed => Weekday.Wed
    # Thu => Weekday.Thu
    # Fri => Weekday.Fri
    # Sat => Weekday.Sat

# 使用元类
# type()
# 动态语言和静态语言最大的不同，就是函数和类的定义，不是编译时定义的，而是运行时动态创建的。
    # class Hello(object):
    #     def hello(self, name='world'):
    #         print('Hello, %s.' % name)

    # >>> from hello import Hello
    # >>> h = Hello()
    # >>> h.hello()
    # Hello, world.
    # >>> print(type(Hello))
    # <class 'type'>
    # >>> print(type(h))
    # <class 'hello.Hello'>
# type()函数可以查看一个类型或变量的类型，Hello是一个class，它的类型就是type，而h是一个实例，它的类型就是class Hello。
# 我们说class的定义是运行时动态创建的，而创建class的方法就是使用type()函数。
# type()函数既可以返回一个对象的类型，又可以创建出新的类型，比如，我们可以通过type()函数创建出Hello类，而无需通过class Hello(object)...的定义：
    # def fn(self, name='world'):
    #     print('Hello, %s.' % name)
    #
    # Hello = type('Hello', (object,), dict(hello=fn))
    #
    # h = Hello()
    # h.hello()
    # Hello, world.
    # >>> print(type(Hello))
    # <class 'type'>
    # >>> print(type(h))
    # <class '__main__.Hello'>
# 要创建一个class对象，type()函数依次传入3个参数：
# 1. class的名称；
# 2. 继承的父类集合，注意Python支持多重继承，如果只有一个父类，别忘了tuple的单元素写法；
# 3. class的方法名称与函数绑定，这里我们把函数fn绑定到方法名hello上。
# 通过type()函数创建的类和直接写class是完全一样的，因为Python解释器遇到class定义时，仅仅是扫描一下class定义的语法，然后调用type()函数创建出class。
# 正常情况下，我们都用class Xxx...来定义类，但是，type()函数也允许我们动态创建出类来，也就是说，动态语言本身支持运行期动态创建类，
    # 这和静态语言有非常大的不同，要在静态语言运行期创建类，必须构造源代码字符串再调用编译器，或者借助一些工具生成字节码实现，本质上都是动态编译，会非常复杂。

# metaclass
# 除了使用type()动态创建类以外，要控制类的创建行为，还可以使用metaclass。
# metaclass，直译为元类，简单的解释就是：当我们定义了类以后，就可以根据这个类创建出实例，所以：先定义类，然后创建实例。
# 但是如果我们想创建出类呢？那就必须根据metaclass创建出类，所以：先定义metaclass，然后创建类。
# 连接起来就是：先定义metaclass，就可以创建类，最后创建实例。
# 所以，metaclass允许你创建类或者修改类。换句话说，你可以把类看成是metaclass创建出来的“实例”。
# 例如：
    # 定义ListMetaclass，按照默认习惯，metaclass的类名总是以Metaclass结尾，以便清楚地表示这是一个metaclass：
    # # metaclass是类的模板，所以必须从`type`类型派生：
    # class ListMetaclass(type):
    #     def __new__(cls, name, bases, attrs):
    #         attrs['add'] = lambda self, value: self.append(value)
    #         return type.__new__(cls, name, bases, attrs)
    # 有了ListMetaclass，我们在定义类的时候还要指示使用ListMetaclass来定制类，传入关键字参数metaclass：
    # class MyList(list, metaclass=ListMetaclass):
    #     pass
    # 当我们传入关键字参数metaclass时，魔术就生效了，它指示Python解释器在创建MyList时，要通过ListMetaclass.__new__()来创建，在此，我们可以修改类的定义，比如，加上新的方法，然后，返回修改后的定义。
    # __new__()方法接收到的参数依次是：
    # 1. 当前准备创建的类的对象；
    # 2. 类的名字；
    # 3. 类继承的父类集合；
    # 4. 类的方法集合
    # 测试一下MyList是否可以调用add()方法：
    # >>> L = MyList()
    # >>> L.add(1)
    # >> L
    # [1]
# 正常情况下，确实应该直接写，通过metaclass修改纯属变态。但是，总会遇到需要通过metaclass修改类定义的。ORM就是一个典型的例子。
# ORM全称“Object Relational Mapping”，即对象-关系映射，就是把关系数据库的一行映射为一个对象，也就是一个类对应一个表，这样，写代码更简单，不用直接操作SQL语句。
# 要编写一个ORM框架，所有的类都只能动态定义，因为只有使用者才能根据表的结构定义出对应的类来。
# ORM实现例子：
    # 1. 编写底层模块的第一步，就是先把调用接口写出来。比如，使用者如果使用这个ORM框架，想定义一个User类来操作对应的数据库表User，我们期待他写出这样的代码：
    # class User(Model):
    #     # 定义类的属性到列的映射：
    #     id = IntegerField('id')
    #     name = StringField('username')
    #     email = StringField('email')
    #     password = StringField('password')
    #
    # # 创建一个实例：
    # u = User(id=12345, name='Michael', email='test@orm.org', password='my-pwd')
    # # 保存到数据库：
    # u.save()
    # 其中，父类Model和属性类型StringField、IntegerField是由ORM框架提供的，剩下的魔术方法比如save()全部由父类Model自动完成。虽然metaclass的编写会比较复杂，但ORM的使用者用起来却异常简单。
    # 2. 按上面的接口来实现该ORM：
    # 2.1 首先来定义Field类，它负责保存数据库表的字段名和字段类型：
        # class Field(object):
        #     def __init__(self, name, column_type):
        #         self.name = name
        #         self.column_type = column_type
        #     def __str__(self):
        #         return '<%s:%s>' % (self.__class__.__name__, self.name)
    # 在Field的基础上，进一步定义各种类型的Field，比如StringField，IntegerField等等：
        # class StringField(Field):
        #     def __init__(self, name):
        #         super(StringField, self).__init__(name, 'varchar(100)')
        #
        # class IntegerField(Field):
        #     def __init__(self, name):
        #         super(IntegerField, self).__init__(name, 'bigint')
    # 2.2 下一步，就是编写最复杂的ModelMetaclass了：
        # class ModelMetaclass(type):
        #     def __new__(cls, name, bases, attrs):
        #         if name=='Model':
        #             return type.__new__(cls, name, bases, attrs)
        #         print('Found model: %s' % name)
        #         mappings = dict()
        #         for k, v in attrs.items():
        #             if isinstance(v, Field):
        #                 print('Found mapping: %s ==> %s' % (k, v))
        #                 mappings[k] = v
        #         for k in mappings.keys():
        #             attrs.pop(k)
        #         attrs['__mappings__'] = mappings # 保存属性和列的映射关系
        #         attrs['__table__'] = name # 假设表名和类名一致
        #         return type.__new__(cls, name, bases, attrs)
    # 以及基类Model：
        # class Model(dict, metaclass=ModelMetaclass):
        #     def __init__(self, **kw):
        #         super(Model, self).__init__(**kw)
        #
        #     def __getattr__(self, key):
        #         try:
        #             return self[key]
        #         except KeyError:
        #             raise AttributeError(r"'Model' object has no attribute '%s'" % key)
        #
        #     def __setattr__(self, key, value):
        #         self[key] = value
        #
        #     def save(self):
        #         fields = []
        #         params = []
        #         args = []
        #         for k, v in self.__mappings__.items():
        #             fields.append(v.name)
        #             params.append('?')
        #             args.append(getattr(self, k, None))
        #         sql = 'insert into %s (%s) values (%s)' % (self.__table__, ','.join(fields), ','.join(params))
        #         print('SQL: %s' % sql)
        #         print('ARGS: %s' % str(args))

