# 生成SSL证书(使用golearn.com域名进行测试)， 证书存储在ssl目录中
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=golearn.com/O=golearn" -addext "subjectAltName = DNS:golearn.com"

# 添加ssl证书到k8s中
kubectl create secret tls golearn-tls --cert=./tls.crt --key=./tls.key


# nginx-ingress-deployment.yaml 安装nginx-ingress控制器
kubectl apply nginx-ingress-deployment.yaml


# go-server-pod.yaml进行配置pod, service, ingress组件
kubectl apply go-server-pod.yaml

# 通过curl进行检查
curl -H "Host: golearn.com" https://10.43.82.48/healthz -k