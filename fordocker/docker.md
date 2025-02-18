## docker pull
```bash
docker pull kindest/node:v1.22.7
```


## docker image
```bash
# 列出image
docker image list
# 查询批量删除image
docker image rm $(docker image ls -q redis)
docker rmi --force $(docker image ls -q redis)
```

## docker run
```bash
# 检查os信息
docker run --rm <image_id> cat /etc/os-release
# docker run检查
docker run -it <image_id> sh
# 后台特权运行运行镜像
docker run --name openssl340 -d --privileged alpine:3.19.4.1 sleep 99999
docker run --name openssl341 -d --privileged alpine-aarch64:3.19.4.1 sleep 99999
docker run -d --name openssl340 --privileged alpine-glibc:alpine-3.15.11_glibc-2.33 sleep 99999
docker run -d --name openssl341 --privileged alpine-glibc:alpine-3.15.11_glibc-2.33 sleep 99999
```

## docker cp
```bash
docker cp myfile.txt my_container:/tmp/
```

## docker container_id_or_name
```bash
docker start <container_id_or_name>
docker exec -it <container_id_or_name> sh
docker rm <container_id_or_name>
docker stop <container_id_or_name>
```

## docker save load
```bash
## 在 Dockerfile 文件所在目录构建镜像
docker build -t my_image_name:v3 .
docker save my_image_name:v3 -o filename
docker save my_image_name:v3 | gzip > my_image_name:v3.tar.gz
docker load -i my_image_name:v3.tar.gz
```