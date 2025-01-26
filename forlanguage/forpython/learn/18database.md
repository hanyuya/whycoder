## 使用SQLite

SQLite是一种嵌入式数据库，它的数据库就是一个文件。由于SQLite本身是C写的，而且体积很小，所以，经常被集成到各种应用程序中，甚至在iOS和Android的App中都可以集成。

Python就内置了SQLite3，所以，在Python中使用SQLite，不需要安装任何东西，直接使用。

在使用SQLite前，我们先要搞清楚几个概念：

表是数据库中存放关系数据的集合，一个数据库里面通常都包含多个表，比如学生的表，班级的表，学校的表，等等。表和表之间通过外键关联。

要操作关系数据库，首先需要连接到数据库，一个数据库连接称为`Connection`；

连接到数据库后，需要打开游标，称之为Cursor，通过Cursor执行SQL语句，然后，获得执行结果。

Python定义了一套操作数据库的API接口，任何数据库要连接到Python，只需要提供符合Python标准的数据库驱动即可。

由于SQLite的驱动内置在Python标准库中，所以我们可以直接来操作SQLite数据库。

我们在Python交互式命令行实践一下：
```python
# 导入SQLite驱动:
>>> import sqlite3
# 连接到SQLite数据库
# 数据库文件是test.db
# 如果文件不存在，会自动在当前目录创建:
>>> conn = sqlite3.connect('test.db')
# 创建一个Cursor:
>>> cursor = conn.cursor()
# 执行一条SQL语句，创建user表:
>>> cursor.execute('create table user (id varchar(20) primary key, name varchar(20))')
<sqlite3.Cursor object at 0x10f8aa260>
# 继续执行一条SQL语句，插入一条记录:
>>> cursor.execute('insert into user (id, name) values (\'1\', \'Michael\')')
<sqlite3.Cursor object at 0x10f8aa260>
# 通过rowcount获得插入的行数:
>>> cursor.rowcount
1
# 提交事务:
>>> conn.commit()
# 关闭Cursor:
>>> cursor.close()
# 关闭Connection:
>>> conn.close()
```
我们再试试查询记录：
```python
>>> conn = sqlite3.connect('test.db')
>>> cursor = conn.cursor()
# 执行查询语句:
>>> cursor.execute('select * from user where id=?', ('1',))
<sqlite3.Cursor object at 0x10f8aa340>
# 获得查询结果集:
>>> values = cursor.fetchall()
>>> values
[('1', 'Michael')]
>>> cursor.close()
>>> conn.close()
```
使用Python的DB-API时，只要搞清楚`Connection`和`Cursor`对象，打开后一定记得关闭，就可以放心地使用。

使用`Cursor`对象执行`insert`，`update`，`delete`语句时，执行结果由`rowcount`返回影响的行数，就可以拿到执行结果。

使用Cursor对象执行select语句时，通过`fetchall()`可以拿到结果集。结果集是一个`list`，每个元素都是一个`tuple`，对应一行记录。

如果SQL语句带有参数，那么需要把参数按照位置传递给`execute()`方法，有几个?占位符就必须对应几个参数，例如：
```python
cursor.execute('select * from user where name=? and pwd=?', ('abc', 'password'))
```
SQLite支持常见的标准SQL语句以及几种常见的数据类型。具体文档请参阅SQLite官方网站。

### 小结
- 在Python中操作数据库时，要先导入数据库对应的驱动，然后，通过Connection对象和Cursor对象操作数据。
- 要确保打开的Connection对象和Cursor对象都正确地被关闭，否则，资源就会泄露。
- 如何才能确保出错的情况下也关闭掉Connection对象和Cursor对象呢？请回忆try:...except:...finally:...的用法。

## 使用MySQL

MySQL是Web世界中使用最广泛的数据库服务器。SQLite的特点是轻量级、可嵌入，但不能承受高并发访问，适合桌面和移动应用。而MySQL是为服务器端设计的数据库，能承受高并发访问，同时占用的内存也远远大于SQLite。

此外，**MySQL内部有多种数据库引擎，最常用的引擎是支持数据库事务的InnoDB**。

安装MySQL
可以直接从MySQL官方网站下载最新的Community Server 8.x版本。MySQL是跨平台的，选择对应的平台下载安装文件，安装即可。

安装时，MySQL会提示输入root用户的口令，请务必记清楚。如果怕记不住，就把口令设置为password。

在Windows上，安装时请选择UTF-8编码，以便正确地处理中文。

在Mac或Linux上，需要编辑MySQL的配置文件，把数据库默认的编码全部改为UTF-8。MySQL的配置文件默认存放在`/etc/my.cnf`或者`/etc/mysql/my.cnf`：
```python
[client]
default-character-set = utf8mb4

[mysqld]
default-storage-engine = INNODB
character-set-server = utf8mb4
collation-server = utf8_general_ci
```
重启MySQL后，可以通过MySQL的客户端命令行检查编码：
```python
$ mysql -u root -p
Enter password: 
Welcome to the MySQL monitor...
...

mysql> show variables like '%char%';
+--------------------------+--------------------------------------+
| Variable_name            | Value                                |
+--------------------------+--------------------------------------+
| character_set_client     | utf8mb4                              |
| character_set_connection | utf8mb4                              |
| character_set_database   | utf8mb4                              |
| character_set_filesystem | binary                               |
| character_set_results    | utf8mb4                              |
| character_set_server     | utf8mb4                              |
| character_set_system     | utf8mb3                              |
| character_sets_dir       | /usr/local/mysql-8.x/share/charsets/ |
+--------------------------+--------------------------------------+
8 rows in set (0.00 sec)
```
看到utf8mb4字样就表示编码设置正确。

> 注意
> 
>如果MySQL的版本<5.5.3，则只能把编码设置为utf8，utf8mb4支持最新的Unicode标准，可以显示emoji字符，但utf8无法显示emoji字符。

### 用Docker启动MySQL
如果不想安装MySQL，还可以以Docker的方式快速启动MySQL。

首先安装Docker Desktop，然后在命令行输入：
```bash
$ docker run -e MYSQL_ROOT_PASSWORD=password -p 3306:3306 --name mysql-8.4 -v ./mysql-data:/var/lib/mysql mysql:8.4 --mysql-native-password=ON --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```
上述命令详细参数如下：
- `-e MYSQL_ROOT_PASSWORD=password`：传入root用户口令的环境变量，密码是password；
- `-p 3306:3306`：在本机3306端口监听；
- `--name mysql-8.4`：启动后容器的名称为mysql-8.4，可任意设置；
- `-v ./mysql-data:/var/lib/mysql`：把当前目录`./mysql-data`映射到容器目录`/var/lib/mysql`，此目录存放MySQL数据库文件，避免容器停止后数据丢失；
- `mysql:8.4`：启动镜像名称为mysql:8.4；
- `--mysql-native-password=ON`：表示启用明文口令；
- `--character-set-server=utf8mb4`：表示启用utf8mb4作为字符集；
- `--collation-server=utf8mb4_unicode_ci`：表示启用utf8mb4作为排序规则。
运行命令后可看到如下输出：
```python
2024-07-11 02:44:05+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.4.1-1.el9 started.
...
2024-07-11T02:44:16.874162Z 0 [System] [MY-015015] [Server] MySQL Server - start.
...
2024-07-11T02:44:17.120017Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
2024-07-11T02:44:17.561242Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
...
2024-07-11T02:44:17.868691Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '8.4.1'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server - GPL.
```
看到最后一行ready for connections表示启动成功。

### 安装MySQL驱动
由于MySQL服务器以独立的进程运行，并通过网络对外服务，所以，需要支持Python的MySQL驱动来连接到MySQL服务器。MySQL官方提供了`mysql-connector-python`驱动：
```python
$ pip install mysql-connector-python 
```
我们演示如何连接到MySQL服务器的test数据库：
```python
# 导入MySQL驱动:
>>> import mysql.connector
# 注意把password设为你的root口令:
>>> conn = mysql.connector.connect(user='root', password='password', database='test')
>>> cursor = conn.cursor()
# 创建user表:
>>> cursor.execute('create table user (id varchar(20) primary key, name varchar(20))')
# 插入一行记录，注意MySQL的占位符是%s:
>>> cursor.execute('insert into user (id, name) values (%s, %s)', ['1', 'Michael'])
>>> cursor.rowcount
1
# 提交事务:
>>> conn.commit()
>>> cursor.close()
# 运行查询:
>>> cursor = conn.cursor()
>>> cursor.execute('select * from user where id = %s', ('1',))
>>> values = cursor.fetchall()
>>> values
[('1', 'Michael')]
# 关闭Cursor和Connection:
>>> cursor.close()
True
>>> conn.close()
```
由于Python的DB-API定义都是通用的，所以，操作MySQL的数据库代码和SQLite类似。

### 小结
- 执行INSERT等操作后要调用`commit()`提交事务；
- MySQL的SQL占位符是`%s`。

## 使用SQLAlchemy

数据库表是一个二维表，包含多行多列。把一个表的内容用Python的数据结构表示出来的话，可以用一个list表示多行，list的每一个元素是tuple，表示一行记录，比如，包含id和name的user表：
```python
[
    ('1', 'Michael'),
    ('2', 'Bob'),
    ('3', 'Adam')
]
```
Python的DB-API返回的数据结构就是像上面这样表示的。

但是用tuple表示一行很难看出表的结构。如果把一个tuple用class实例来表示，就可以更容易地看出表的结构来：
```python
class User(object):
    def __init__(self, id, name):
        self.id = id
        self.name = name

[
    User('1', 'Michael'),
    User('2', 'Bob'),
    User('3', 'Adam')
]
```
这就是传说中的ORM技术：Object-Relational Mapping，把关系数据库的表结构映射到对象上。是不是很简单？

但是由谁来做这个转换呢？所以ORM框架应运而生。

在Python中，最有名的ORM框架是SQLAlchemy。我们来看看SQLAlchemy的用法。

首先通过pip安装SQLAlchemy：
```python
$ pip install sqlalchemy
```
然后，利用上次我们在MySQL的test数据库中创建的user表，用SQLAlchemy来试试：

第一步，导入SQLAlchemy，并初始化DBSession：
```python
# 导入:
from sqlalchemy import Column, String, create_engine
from sqlalchemy.orm import sessionmaker
from sqlalchemy.orm import declarative_base

# 创建对象的基类:
Base = declarative_base()

# 定义User对象:
class User(Base):
    # 表的名字:
    __tablename__ = 'user'

    # 表的结构:
    id = Column(String(20), primary_key=True)
    name = Column(String(20))

# 初始化数据库连接:
engine = create_engine('mysql+mysqlconnector://root:password@localhost:3306/test')
# 创建DBSession类型:
DBSession = sessionmaker(bind=engine)
```
以上代码完成SQLAlchemy的初始化和具体每个表的class定义。如果有多个表，就继续定义其他class，例如School：
```python
class School(Base):
    __tablename__ = 'school'
    id = ...
    name = ...
```
`create_engine()`用来初始化数据库连接。SQLAlchemy用一个字符串表示连接信息：
```python
'数据库类型+数据库驱动名称://用户名:口令@机器地址:端口号/数据库名'
```
你只需要根据需要替换掉用户名、口令等信息即可。

下面，我们看看如何向数据库表中添加一行记录。

由于有了ORM，我们向数据库表中添加一行记录，可以视为添加一个`User`对象：
```python
# 创建session对象:
session = DBSession()
# 创建新User对象:
new_user = User(id='5', name='Bob')
# 添加到session:
session.add(new_user)
# 提交即保存到数据库:
session.commit()
# 关闭session:
session.close()
```
可见，关键是获取session，然后把对象添加到session，最后提交并关闭。DBSession对象可视为当前数据库连接。

如何从数据库表中查询数据呢？有了ORM，查询出来的可以不再是tuple，而是User对象。SQLAlchemy提供的查询接口如下：
```python
# 创建Session:
session = DBSession()
# 创建Query查询，filter是where条件，最后调用one()返回唯一行，如果调用all()则返回所有行:
user = session.query(User).filter(User.id=='5').one()
# 打印类型和对象的name属性:
print('type:', type(user))
print('name:', user.name)
# 关闭Session:
session.close()
```
运行结果如下：
```python
type: <class '__main__.User'>
name: Bob
```
可见，ORM就是把数据库表的行与相应的对象建立关联，互相转换。

由于关系数据库的多个表还可以用外键实现一对多、多对多等关联，相应地，ORM框架也可以提供两个对象之间的一对多、多对多等功能。

例如，如果一个User拥有多个Book，就可以定义一对多关系如下：
```python
class User(Base):
    __tablename__ = 'user'

    id = Column(String(20), primary_key=True)
    name = Column(String(20))
    # 一对多:
    books = relationship('Book')

class Book(Base):
    __tablename__ = 'book'

    id = Column(String(20), primary_key=True)
    name = Column(String(20))
    # “多”的一方的book表是通过外键关联到user表的:
    user_id = Column(String(20), ForeignKey('user.id'))
```
当我们查询一个User对象时，该对象的books属性将返回一个包含若干个Book对象的list。

### 小结
- ORM框架的作用就是把数据库表的一行记录与一个对象互相做自动转换。
- 正确使用ORM的前提是了解关系数据库的原理。