本次作业操作内容如下：
1. 使用gin构建http-server
2. 调用老师编写的metrics.go程序
3. 读懂metrics.go作业流程
4. 制作DockerFile, 最新版本是v11
docker pull changsiheng/gin-server:v11
5. 制作gin-server.yaml用于k8s环境部署
6. 安装helm工具
7. 通过loki仓库安装promethus, grafana, Loki
8. 通过kubectl edit svc 暴露 promethus, grafana 的端口
9. 访问应用gin-server的/hello路由
10. 通过promethus查看ginserver_execution_latency_seconds_bucket指标