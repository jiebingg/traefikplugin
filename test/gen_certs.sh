openssl genrsa -out ca.key 2048

# 生成自签名根证书: ca.crt
openssl req -new -key ca.key -x509 -days 3650 -out ca.crt -subj /C=CN/ST=Singapore/O="Localhost Ltd"/CN="Localhost Root"

# >>>>>>>>>>>>>>>>>> 服务器证书 <<<<<<<<<<<<<<<<<<<<<<
# 生成服务器证书私钥: ca.key
openssl genrsa -out server.key 2048

# 生成服务器证书请求: server.csr
openssl req -new -nodes -key server.key -out server.csr -subj /C=CN/ST=Singapore/L=Singapore/O="Localhost Server"/CN=localhost

# 签名服务器证书: server.crt
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt

# >>>>>>>>>>>>>>>>>> 客户端证书 <<<<<<<<<<<<<<<<<<<<<<
# 生成客户端证书私钥: ca.key
openssl genrsa -out client.key 2048

# 生成客户端证书请求: client.csr
openssl req -new -nodes -key client.key -out client.csr -subj /C=CN/ST=Singapore/L=Singapore/O="Localhost Client"/CN=localhost

# 签名客户端证书: client.crt
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt