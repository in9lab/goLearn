# 编译Dockerfile, 并打tag
docker build -t changsiheng/gin-server:week3 .

# 推送镜像至dockerhub中
docker push changsiheng/gin-server:week3

# DockerHub
changsiheng/gin-server

# 启动http-server
docker run -p 8888:8080 --name gin-server -d changsiheng/gin-server:week3

# 通过nsenter查看IP
nsenter -t$( docker inspect $(docker ps | grep gin-server | awk '{print $1}') | grep -i \"pid\": | awk -F: '{print $2}' | awk -F, '{print $1}')  -n ip a
