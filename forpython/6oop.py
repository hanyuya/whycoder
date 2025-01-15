# 面相对象编程
# 1、面向过程的程序设计把计算机程序视为一系列的命令集合，即一组函数的顺序执行。为了简化程序设计，面向过程把函数继续切分为子函数，即把大块函数通过切割成小块函数来降低系统的复杂度。
# 2、面向对象的程序设计把计算机程序视为一组对象的集合，而每个对象都可以接收其他对象发过来的消息，并处理这些消息，计算机程序的执行就是一系列消息在各个对象之间传递。
# 3、在Python中，所有数据类型都可以视为对象，当然也可以自定义对象。自定义的对象数据类型就是面向对象中的类（Class）的概念。

# 类（Class）和 实例（Instance）
# 1、类是抽象的模板，它定义了对象的属性和方法。实例是根据类创建出来的具体对象，每个对象都拥有相同的方法，但各自的数据可能不同。
# 定义类：
# class 类名(object):
#     pass
# 2、可以自由地给一个实例变量绑定属性，比如，给实例bart绑定一个name属性：
# >>> bart.name = 'Bart Simpson'
# >>> bart.name
# 'Bart Simpson'
# 3、__init__方法：
# __init__方法的第一个参数永远是self，表示创建的实例本身，因此，在__init__方法内部，就可以把各种属性绑定到self，因为self就指向创建的实例本身。
# 可以在创建实例的时候，把一些我们认为必须绑定的属性强制填写进去。通过定义一个特殊的__init__方法，在创建实例的时候，就把name，score等属性绑上去。
# class Student(object):
#     def __init__(self, name, score):
#         self.name = name
#         self.score = score
# 在创建实例的时候，就不能传入空的参数了，必须传入与__init__方法匹配的参数，但self不需要传，Python解释器自己会把实例变量传进去：
# bart = Student('Bart Simpson', 59)
# 4、要定义一个方法，除了第一个参数是self外，其他和普通函数一样。要调用一个方法，只需要在实例变量上直接调用，除了self不用传递，其他参数正常传入：
# class Student(object):
#     def __init__(self, name, score):
#         self.name = name
#         self.score = score
#
#     def print_score(self):
#         print('%s: %s' % (self.name, self.score))

# 访问限制
# 在Python中，可以用双下划线__前缀来表示实例变量或方法不应该被直接访问，只能通过对象的方法来访问。（私有变量（private））
# 在Python中，变量名类似__xxx__的，也就是以双下划线开头，并且以双下划线结尾的，是特殊变量，特殊变量是可以直接访问的，不是private变量，所以，不能用__name__、__score__这样的变量名。
# 比如_name，这样的实例变量外部是可以访问的，但是，按照约定俗成的规定，当你看到这样的变量时，意思就是，“虽然我可以被访问，但是，请把我视为私有变量，不要随意访问”。
    #  不能直接访问__name是因为Python解释器对外把__name变量改成了_Student__name，所以，仍然可以通过_Student__name来访问__name变量：
    # >>> bart._Student__name
    # 'Bart Simpson'
    # 例如，我们定义一个Student类：
    # class Student(object):
    #     def __init__(self, name, score):
    #         self.__name = name
    #         self.__score = score

# 继承和多态
# 在OOP程序设计中，当我们定义一个class的时候，可以从某个现有的class继承，新的class称为子类（Subclass），而被继承的class称为基类、父类或超类（Base class、Super class）。
# 例如：
# class Dog(Animal):
#     pass
# class Cat(Animal):
#     pass
# 子类获得了父类的全部功能，同时还可以新增自己特有的功能。
# 当子类和父类都存在相同的run()方法时，我们说，子类的run()覆盖了父类的run()，在代码运行的时候，总是会调用子类的run()。
# 判断一个变量是否是某个类型可以用isinstance()判断：
# if isinstance(obj, Animal):
#     obj.run()

# 多态，即一个变量可以指向不同类型的对象，而这些对象在运行时会表现出不同的行为。
# 多态的好处是，我们只需要定义父类，就可以让子类去继承，而不需要修改源代码。
# 多态的实现，就是利用“鸭子类型”，只要某个对象看起来像鸭子，走起路来像鸭子，那它就可以被看做是鸭子。


# 静态语言 vs 动态语言
# 对于静态语言（例如Java）来说，如果需要传入Animal类型，则传入的对象必须是Animal类型或者它的子类，否则，将无法调用run()方法。
# 对于Python这样的动态语言来说，则不一定需要传入Animal类型。我们只需要保证传入的对象有一个run()方法就可以了：
# class Timer(object):
#     def run(self):
#         print('Start...')
# 这就是动态语言的“鸭子类型”，它并不要求严格的继承体系，一个对象只要“看起来像鸭子，走起路来像鸭子”，那它就可以被看做是鸭子。
# Python的“file-like object“就是一种鸭子类型。对真正的文件对象，它有一个read()方法，返回其内容。
# 但是，许多对象，只要有read()方法，都被视为“file-like object“。许多函数接收的参数就是“file-like object“，你不一定要传入真正的文件对象，完全可以传入任何实现了read()方法的对象。

# 获取对象信息
# 使用type()函数可以获取对象的类型：
# 基本类型都可以用type()判断：
    # >>> type(123)
    # <class 'int'>
# 一个变量指向函数或者类，也可以用type()判断：
    # >>> type(abs)
    # <class 'builtin_function_or_method'>
    # >>> type(a)
    # <class '__main__.Animal'>
# type()函数返回的是什么类型呢？它返回对应的Class类型。如果我们要在if语句中判断，就需要比较两个变量的type类型是否相同：
    # if type(obj) == type(123):
    #     print('123')
    # else:
    #     print('not 123')
# 判断基本数据类型可以直接写int，str等，但如果要判断一个对象是否是函数怎么办？可以使用types模块中定义的常量：
    # >>> import types
    # >>> def fn():
    # ...     pass
    # ...
    # >>> type(fn)==types.FunctionType
    # True
    # >>> type(abs)==types.BuiltinFunctionType
    # True
    # >>> type(lambda x: x)==types.LambdaType
    # True
    # >>> type((x for x in range(10)))==types.GeneratorType
    # True

# 使用isinstance()函数可以判断一个对象是否是一个类的实例，isinstance()判断的是一个对象是否是该类型本身，或者位于该类型的父继承链上。
# 例如：
    # >>> isinstance('abc', str)
    # True
    # >>> isinstance(123, int)
    # True
    # >>> isinstance(bart, Student)
    # True
    # >>> isinstance(fn, types.FunctionType)

# 使用dir()
# 如果要获得一个对象的所有属性和方法，可以使用dir()函数，它返回一个包含字符串的list，比如，获得一个str对象的所有属性和方法：
    # >>> dir('ABC')
    # ['__add__', '__class__',..., '__subclasshook__', 'capitalize', 'casefold',..., 'zfill']
    # 类似__xxx__的属性和方法在Python中都是有特殊用途的，比如__len__方法返回长度。
    # 在Python中，如果你调用len()函数试图获取一个对象的长度，实际上，在len()函数内部，它自动去调用该对象的__len__()方法，所以，下面的代码是等价的：
    # >>> len('ABC')
    # 3
    # >>> 'ABC'.__len__()
    # 3
# 仅仅把属性和方法列出来是不够的，配合getattr()、setattr()以及hasattr()，我们可以直接操作一个对象的状态：
    # >>> hasattr(bart, 'name')  # 判断对象是否有name属性
    # True
    # >>> getattr(bart, 'name')  # 获取name属性的值
    # 'Bart Simpson'
    # >>> setattr(bart, 'age', 18)  # 设置age属性的值
    # >>> getattr(bart, 'age')
    # 18
    # 如果试图获取不存在的属性，会抛出AttributeError的错误：
    # >>> getattr(bart, 'gender')
    # AttributeError: 'Student' object has no attribute 'gender'
    # 可以传入一个default参数，如果属性不存在，就返回默认值：
    # >>> getattr(bart, 'gender', 'unknown') # 获取不到gender属性，返回默认值'unknown'
    # 也可以获得对象的方法：
    # >>> hasattr(bart, 'print_score')
    # True
    # >>> getattr(bart, 'print_score')
    # <bound method Student.print_score of <__main__.Student object at 0x7f1a700a3438>>
    # >>> fn = getattr(obj, 'power') # 获取属性'power'并赋值到变量fn
    # >>> fn(2) # 调用fn函数，计算obj的2次方，调用fn()与调用obj.power()是一样的

# 实例属性和类属性
# 由于Python是动态语言，根据类创建的实例可以任意绑定属性。
    # 给实例绑定属性的方法是通过实例变量，或者通过self变量：
    # class Student(object):
    #     def __init__(self, name):
    #         self.name = name
    #
    # s = Student('Bob')
    # s.score = 90
# 如果Student类本身需要绑定一个属性可以直接在class中定义属性，这种属性是类属性，归Student类所有：
    # class Student(object):
    #     count = 0  # 类属性