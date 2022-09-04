本周完成内容：

1. 修改http-server为 services1, services2, services3， 在dockerHub中的changsiheng/gin-server-8081:v1, gin-server-8082:v1, gin-server-8083:v1
2. 代码中的Mesh链路为services1 --> services2 --> services3
3. serviceMesh文件中gin-server-deployment和gin-server-services是部署使用的deployment和services
4. 生成TLS使用证书
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt

5. 绑定证书
## 注意, 需要在istio ingress的空间下进行相同操作，否则会出现SSL错误
kubectl create -n tracing secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt

kubectl create -n istio-system secret tls cncamp-credential --key=cncamp.io.key --cert=cncamp.io.crt


7. 配置访问规则, 规则内容存放在istio-specs.yaml，分别有路由和安全相关的配置
8. 安装Jaeger通过jaeger.yaml进行
10. curl测试链路响应
for i in {1..1000}; do curl --resolve httpsserver.cncamp.io:443:10.43.55.73 https://httpsserver.cncamp.io/healthz -k; done;

11.Jaeger查看安装结果, 但我的结果不知道为啥没有像老师的示例中，可以很分明的看出每个方法的用时
![微信图片_20220814234751](https://user-images.githubusercontent.com/16226183/184544765-7a9c99a1-e8c3-4634-84ef-f99cc11ef09b.png)
