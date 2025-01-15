import paramiko

# 利用paramiko实现ssh功能
    # 1 利用用户名密码登录
def useUserPwd():
    ssh_client = paramiko.SSHClient()
    #允许连接不在know_hosts中的主机
    ssh_client.set_missing_host_key_policy(paramiko.AutoAddPolicy)
    #连接服务器
    ssh_client.connect(hostname="10.235.187.143",port=22,username="ubuntu",password="UbuntU1!2@3#4$")
    #执行指定命令,返回三个结果，标准输入，标准输出，标准错误，标准输出和标准错误不同时存在
    stdin,stdout,stderr = ssh_client.exec_command("ls -al")
    #读取输出结果
    std_out = stdout.read()
    std_err = stderr.read()
    if len(std_out) == 0:
        print(std_err)
    else:
        print(std_out)
    #关闭服务
    ssh_client.close()

    # 2 RSA免密登录
def useRsa():
    private_key = paramiko.RSAKey.from_private_key_file('/root/.ssh/id_rsa')
    # 创建SSH对象
    ssh = paramiko.SSHClient()
    # 允许连接不在know_hosts文件中的主机
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    # 连接服务器
    ssh.connect(hostname='10.235.187.143', port=22, username='root', pkey=private_key)
    
    # 执行命令
    stdin, stdout, stderr = ssh.exec_command('ls -al')
    # 获取命令结果
    result = stdout.read()
    
    # 关闭连接
    ssh.close()

# 利用paramiko.Transport实现ssh功能
    # 1 利用用户名密码登录
def useUserTrans():
    transport = paramiko.Transport(('10.235.187.143', 22))
    transport.connect(username='ubuntu', password='UbuntU1!2@3#4$')
    
    ssh = paramiko.SSHClient()
    ssh._transport = transport
    stdin, stdout, stderr = ssh.exec_command('ls -al')
    print(stdout.read())
    
    transport.close()

def useRsaTrans():
    private_key = paramiko.RSAKey.from_private_key_file('/root/.ssh/id_rsa')  #私钥地址  公钥也在同目录下 id_rsa.pub
    transport = paramiko.Transport(('10.235.187.143', 22))
    transport.connect(username='root', pkey=private_key)
    ssh = paramiko.SSHClient()
    ssh._transport = transport
    stdin, stdout, stderr = ssh.exec_command('ls -al')
    std_out = stdout.read()
    std_err = stderr.read()
    if len(std_out) == 0:
        print(std_err)
    else:
        print(std_out)
    transport.close()


if __name__ == '__main__':
    useUserTrans()