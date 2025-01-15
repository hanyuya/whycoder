-- 登录超级用户
-- sudo -i -u postgres

-- 登录到 PostgreSQL
-- psql -h 127.0.0.1 -U test -d testdb
-- psql -h 127.0.0.1 -U testcp -d testcpdb

-- 查看配置文件路径 -- show config_file
-- 显示hba_file路径 -- show hba_file
-- 列出所有用户及其角色属性 -- \du
-- 查看所有数据库 -- \l
-- 查看当前数据库 -- \c
-- 切换到指定数据库 -- \c testdb
-- 显示所有表 -- \dt
-- 显示所有列 -- \d table_name
-- 显示所有函数 -- \df
-- 退出 -- \q

-- 显示所有数据库
SELECT datname FROM pg_database;

-- 显示当前数据库
SELECT current_database();

-- 创建新数据库
CREATE DATABASE testdb;

-- 删除数据库
DROP DATABASE testdb;

-- 清空数据库中的所有表，但不删除表结构
TRUNCATE TABLE testdb CASCADE;

-- 查询所有的标签
SELECT tablename FROM pg_tables WHERE schemaname = 'public';


-- 创建新用户
CREATE USER testcp WITH PASSWORD 'testcp';

-- 创建数据库并指定所有者
CREATE DATABASE testcpdb OWNER testcp;

-- 授予权限
GRANT ALL PRIVILEGES ON DATABASE testdb TO test;


-- 查询查看连接到数据库的会话：
SELECT pid, usename, datname, state, query 
FROM pg_stat_activity 
WHERE datname = 'testdb';

-- 开特定会话
SELECT pg_terminate_backend(pid)
FROM pg_stat_activity
WHERE datname = 'testdb'
AND pid <> pg_backend_pid();

-- 查看当前会话的pid
select pg_backend_pid();

-- 删除某个表的某条记录
DELETE FROM table_name WHERE id = 1;

-- 查询所有表格
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public';

-- 查询tb_config表格中的所有数据
SELECT * FROM tb_config;

-- 查询tb_config表格id为cmp_cmpServerAddress的记录
SELECT * FROM tb_config WHERE id = 'test';

-- 更新tb_config表格id为cmp_cmpServerAddress的记录
UPDATE tb_config SET item_value = 'https' WHERE id = 'test';
UPDATE tb_config SET item_value = 'https_no_auth' WHERE id = 'test';

