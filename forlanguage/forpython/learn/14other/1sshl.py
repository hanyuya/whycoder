import paramiko
import socket
import threading

def forward_tunnel(local_port, remote_host, remote_port, transport):
    # 创建一个本地套接字
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.bind(('127.0.0.1', local_port))
    server.listen(100)

    while True:
        # 接受来自本地的连接
        client_socket, addr = server.accept()
        print(f"Accepted connection from {addr}")

        # 创建一个新的线程来处理转发
        thread = threading.Thread(target=handle_tunnel, args=(client_socket, remote_host, remote_port, transport))
        thread.start()

def handle_tunnel(client_socket, remote_host, remote_port, transport):
    # 通过SSH连接到远程主机
    remote_socket = transport.open_channel("direct-tcpip", (remote_host, remote_port), client_socket.getpeername())
    
    if remote_socket is None:
        print("Failed to create remote socket")
        client_socket.close()
        return

    # 转发数据
    while True:
        # 从客户端读取数据
        data = client_socket.recv(1024)
        if len(data) == 0:
            break
        remote_socket.send(data)

    client_socket.close()
    remote_socket.close()

def create_ssh_tunnel(local_port, remote_host, remote_port, ssh_host, ssh_user, ssh_password):
    # 创建SSH客户端
    client = paramiko.SSHClient()
    client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    client.connect(ssh_host, username=ssh_user, password=ssh_password)

    # 获取SSH连接的传输对象
    transport = client.get_transport()

    # 启动转发隧道
    forward_tunnel(local_port, remote_host, remote_port, transport)

if __name__ == '__main__':
    # 示例参数
    local_port = 63417
    remote_host = "10.254.80.109"
    remote_port = 1234
    ssh_host = "10.235.187.143"
    ssh_user = "pict"
    ssh_password = "PicT1!2@3#4$"

    create_ssh_tunnel(local_port, remote_host, remote_port, ssh_host, ssh_user, ssh_password)