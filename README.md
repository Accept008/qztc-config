# 客户声连码有效期检查的服务

## 镜像打包上传
    先打包main,再进行镜像打包
    
    sudo root
    GOOS=linux go build -o main main.go
    docker build -t qztc-config:1.0.1 .
    docker run --rm -ti qztc-config:1.0.1
    