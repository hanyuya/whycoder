import unittest
# python 是动态语言，不需要声明变量类型，变量类型由值来决定。

# 基本数据类型
    # int, float, bool, str, None
    # 数字类型 int, float
    # 布尔类型 bool —— True/False
    # 字符串类型 str —— 字符串是以单引号'或双引号"括起来的任意文本，比如'abc'，"xyz"等等。
    # —— s.replace('a', 'b') 替换字符串中的a为b
    # 空值 None

# 格式化输出
    # 1、%实现
    # 使用%运算符可以格式化输出字符串，用法如下：
    # %d 整数
    # %f 浮点数
    # %s 字符串
    # %x 十六进制整数
    # 举例：
    # print("Hello, %s!" % "world")  # 输出 "Hello, world!"
    # print("Age: %d" % 25)  # 输出 "Age: 25"
    # print("Price: %f" % 3.14)  # 输出 "Price: 3.140000"
    # print("Flag: %s" % True)  # 输出 "Flag: True"
    # print("Hex: %x" % 255)  # 输出 "Hex: ff"

# 2、format实现
    # 使用format()方法可以格式化输出字符串，用法如下：
    # format(value, format_spec)
    # value: 要格式化的值
    # format_spec: 格式说明符，用于指定格式
    # 举例：
    # print("Hello, {}!".format("world"))  # 输出 "Hello, world!"
    # print("Age: {}".format(25))  # 输出 "Age: 25"
    # print("Price: {:.2f}".format(3.1415926))  # 输出 "Price: 3.14"
    # print("Flag: {}".format(True))  # 输出 "Flag: True"
    # print("Hex: {:x}".format(255))  # 输出 "Hex: ff"

# 3、f-string实现
    # 3.1 基本用法
    # 在3.6版本后，Python引入了f-string，可以直接在字符串中嵌入变量，语法如下：
    # f"string {variable} string"
    # 举例：
    # name = "world"
    # print(f"Hello, {name}!")  # 输出 "Hello, world!"
    # 3.2 格式化
    # 除了变量外，还可以用大括号{}来格式化输出，语法如下：
    # f"string {variable:format_spec} string"
    # 举例：
    # pi = 3.1415926
    # print(f"Pi is {pi:.2f}")  # 输出 "Pi is 3.14"
    # 3.3 表达式
    # 还可以用圆括号()来执行表达式，语法如下：
    # f"string {expression} string"
    # 举例：
    # x = 10
    # y = 5
    # print(f"The result is {x + y}")  # 输出 "The result is 15"

# list and tuple
    # list: 一种有序的集合，可以随时添加和删除其中的元素。len/pop()/pop(i)/l.sort()
    # tuple: 有序列表叫元组, tuple和list非常类似，但是tuple一旦初始化就不能修改。因为tuple不可变，所以代码更安全。tuple中的元素类型也可以不相同。

# dict
    # 字典，是一种无序的键值对集合(dict内部存放的顺序和key放入的顺序是没有关系的)。
    # dict用大括号{}表示，键值对之间用冒号:分隔，每个键值对之间用逗号,分隔。
    # dict的key必须是不可变对象
    # list是可变的，就不能作为key
    # 举例：
    # d = {'name': 'Alice', 'age': 25, 'city': 'Beijing'}
    # d.get('name') : 获取键值
    # d['not_exist'] : 获取不存在的键值会报错
    # 判断键是否存在：
    # 1、d.get('not_exist', 'default') : 获取不存在的键值，返回默认值/None
    # 2、 'xxx' in d : 判断键是否存在

# set
    # 集合，是一种无序不重复元素的集。
    # set用大括号{}表示，元素之间用逗号,分隔。
    # 提供一个list作为输入集合，返回一个新的set，重复的元素会自动被过滤。
    # s.add(key) : 添加元素
    # s.remove(key) : 删除元素
    # set可以进行集合运算，比如union、intersection、difference等。
    # 举例：
    # s1 = {1, 2, 3}


# if else
    # 分支结构，if/elif/else，条件判断语句。
    # 结构如下：
    # if <条件判断1>:
    #     <执行1>
    # elif <条件判断2>:
    #     <执行2>
    # elif <条件判断3>:
    #     <执行3>
    # else:
    #     <执行4>

# match
# 匹配结构，类似switch语句，可以匹配不同的值，并执行相应的操作。
# match语句除了可以匹配简单的单个值外，还可以匹配多个值、匹配一定范围，并且把匹配后的值绑定到变量
    # match age:
    #     case x if x < 10:
    #         print(f'< 10 years old: {x}')
# match语句还可以匹配列表
    #     args = ['gcc', 'hello.c', 'world.c']
    #     # args = ['clean']
    #     # args = ['gcc']
    #     match args:
    #         # 如果仅出现gcc，报错:
    #         case ['gcc']:
    #             print('gcc: missing source file(s).')
    #         # 出现gcc，且至少指定了一个文件:
    #         case ['gcc', file1, *files]:
    #             print('gcc compile: ' + file1 + ', ' + ', '.join(files))
    #         # 仅出现clean:
    #         case ['clean']:
    #             print('clean')
    #         case _:
    #             print('invalid command.')
    # 结构如下：
    # match <匹配值>:
    #     case <匹配值1>:
    #         <执行1>
    #     case <匹配值2>:
    #         <执行2>
    #    ...
    #     case <匹配值n>:
    #         <执行n>
    #     case _:
    #         <执行默认>

# for
    # 循环结构，用于遍历序列，比如列表、元组、字符串。
    # Python的循环有两种，一种是for...in循环，依次把list或tuple中的每个元素迭代出来
    # for <变量> in <序列>:
    #     <执行>
    # 第二种循环是while循环，只要条件满足，就不断循环，条件不满足时退出循环。
    # while <条件>:
    #     <执行>
    #
    # break语句可以提前退出循环
    # continue语句可以跳过当前循环，继续下一次循环
#

class TestBasic(unittest.TestCase):
    def test_basic(self):
        self.assertEqual(1+1, 2)

if __name__ == '__main__':
    unittest.main()