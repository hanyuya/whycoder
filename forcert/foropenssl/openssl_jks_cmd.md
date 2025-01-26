[openssl man手册](https://docs.openssl.org/3.4/man1/)

## cmp cmd

### cmp ir
```bash
./openssl cmp -cmd ir -server https://127.0.0.1:20809 -path cmp/cmpRequest.do -tls_used -recipient "/C=CN" -subject "/C=CN/CN=testcase" -cert certificate.cer -key key.pem -newkey private.key -trusted test.cer -ignore_keyusage -extracertsout capubs.crt -certout mycert.crt -days 1500 -san_nodefault -no_proxy 127.0.0.1

./openssl cmp -cmd ir -server "https://demo.one.digicert.com/iot/api/v1/cmp" -tls_used -recipient "/C=CN" -cert cos-ecc.pem -key cos-ecc.key -newkey mykey.pem -newkeypass pass:wdl10027679 -subject "/CN=x" -trusted trustcert.crt -certout mycert.pem -ignore_keyusage

./openssl cmp -cmd ir -server "https://demo.one.digicert.com/iot/api/v1/cmp" -tls_used -recipient "/C=CN" -cert cos-ecc.crt -key cos-ecc.key -newkey mykey.pem -subject "/CN=x" -sans "10.235.13.12" -days 3650 -trusted trustcert.crt -certout mycert.pem -extracertsout mycertca.pem -ignore_keyusage -proxy http://192.168.184.16:80

openssl cmp -cmd ir -server "https://demo.one.digicert.com/iot/api/v1/cmp" -recipient "/C=CN" -cert cos-ecc.crt -key cos-ecc.key -newkey mykey.pem -subject "/CN=x" -sans "10.235.13.12" -days 3650 -trusted trustcert.cr -certout mycert.pem -extracertsout mycertca.pem -ignore_keyusage -proxy http://192.168.184.16:80

./openssl cmp -cmd ir -server "https://demo.one.digicert.com/iot/api/v1/cmp" -tls_used -cert cos-ecc.crt -key cos-ecc.key -newkey newkey.pem -subject "/CN=test" -days 3650 -srvcert mediate.pem -certout mycert.pem -extracertsout mycertca.pem -ignore_keyusage -san_nodefault -proxy http://192.168.111.111:80
```
### cmp kur
```bash
./openssl cmp -cmd kur -server https://demo.one.digicert.com/iot/api/v1/cmp -tls_used -cert mycert.pem -key newkey.pem -newkey newkey.pem -srvcert mycertca.pem -certout mycert.pem -extracertsout mycertca.pem

./openssl cmp -cmd kur -server "https://127.0.0.1:20809" -tls_used -cert mycert.pem -key newkey.pem -newkey newkey.pem -srvcert mycertca.pem -certout mycert.pem -extracertsout mycertca.pem -sans "191.168.8.118,10.239.160.118" -no_proxy 127.0.0.1
```
### cmp rr
```bash
./openssl cmp -cmd rr -server "https://10.239.133.216:20809" -tls_used -oldcert mycert.pem -cert mycert.pem -key newkey.pem -srvcert mycertca.pem -no_proxy 10.239.133.216

./openssl cmp -cmd rr -server "https://demo.one.digicert.com/iot/api/v1/cmp" -tls_used -oldcert mycert.pem -cert mycert.pem -key newkey.pem -srvcert mycertca.pem -revreason -1
```

## ocsp查询
```bash
./openssl ocsp -issuer 签发CA证书 -cert 证书 -url "ocsp服务器" -respout ocsp_resp.der -no_nonce

./openssl ocsp -issuer mediate.pem -cert mycert.pem -url "http://ocsp.demo.one.digicert.com" -respout ocsp_resp.der -no_nonce
```
## 公私钥生成 cmd

### 生成私钥
```bash
# 生成ras私钥
./openssl genrsa -out private.key 2048
# 生成ecdsa私钥
./openssl ecparam -genkey -name prime256v1 -out ec_private.key
```
### 查看私钥
```bash
# 查看rsa私钥
./openssl rsa -in private.key -text -noout
# 查看ecdsa私钥
./openssl ec -in ec_private.key -text -noout
```

### 生成公钥
```bash
# 根据ras私钥生成公钥
./openssl rsa -in private.key -inform pem -pubout -out pubkey.pem
# 根据ecdsa私钥生成公钥
./openssl ec -in ec_private.key -pubout -out ec_pubkey.pem
```

### 查看公钥
```bash
# 查看rsa公钥
./openssl rsa -pubin -in pubkey.pem -text -noout
# 查看ecdsa公钥
./openssl ec -pubin -in ec_pubkey.pem -text -noout
```

## 生成CA证书
```bash
./openssl req -new -x509 -key root.key -out root.crt -days 3650 -subj /CN=test
```

## CSR
[req-openssl man手册](https://docs.openssl.org/3.4/man1/openssl-req/#options)

### 生成CSR
```bash
# 根据私钥直接生成csr
./openssl req -new -out cert.csr -key private.key
# 根据私钥直接生成带subject的csr
./openssl req -new -key root.key -out root.csr -subj "/C=CN/CN=rootCA"
# 根据私钥直接生成带subject的csr
./openssl req -new -key mid.key -out mid.csr -subj "/C=CN/CN=midCA"
# 根据私钥直接生成带subject的csr
./openssl req -new -key server.key -out server.csr -subj "/C=CN/CN=server"
# 根据私钥直接生成带subject的csr，并指定openssl config配置文件中的v3_req字段
./openssl req -new -key private.key -out cert.csr -config ./conf/openssl.conf -reqexts v3_req -subj "/C=BD/ST=Test/L=Test/O=Test/OU=Test/CN=Test/emailAddress=Test"
# -nodes -keyout 生成csr
./openssl req -new -nodes -keyout cert.key -out cert.csr -config ./conf/openssl.conf -reqexts SAN -subj "/C=CN/CN=testdns"
# 生成rsa:2048私钥的同时生成csr
./openssl req -new -newkey rsa:2048 -nodes -keyout cert.key -out cert.csr -subj "/C=CN/CN=testdns"
# 生成rsa:2048私钥的同时编辑配置文件SAN字段带san信息生成crs
./openssl req -new -newkey rsa:2048 -nodes -keyout cert.key -out cert.csr -subj "/C=BD/ST=Test/L=Test/O=Test/OU=Test/CN=Test/emailAddress=Test" -reqexts SAN -config <(cat ./openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:test.com"))
```
### 查看CSR

```bash
./openssl req -in cert.csr -noout -text
```
### 根据CSR生成证书

#### 生成根CA证书
```bash
# 根据私钥及csr生成根证书
./openssl x509 -req -in root.csr -out root.crt -signkey root.key -days 3650
# 根据私钥及csr生成根证书，证书信息使用ca.ext文件中的v3_ca字段
./openssl x509 -req -extfile ../conf/rootca.ext -extensions v3_ca -in root.csr -signkey root.key -out root.crt -days 3650
```
```
# ../conf/rootca.ext
[ v3_ca ]
subjectKeyIdentifier=hash
authorityKeyIdentifier=keyid:always,issuer
basicConstraints = critical,CA:true
```
#### 生成中间CA证书
```bash
# 根据私钥及csr并使用根证书签发中间CA证书，证书信息使用ca.ext文件中的v3_mid_ca字段
./openssl x509 -req -extfile ../cnf/ca.ext -extensions v3_mid_ca -in mid.csr -out mid.crt -CA root.crt -CAkey root.key -CAcreateserial -days 3650
```
```
# ../cnf/midca.ext
[ v3_mid_ca ]
# Extensions for a typical intermediate CA (`man x509v3_config`).
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid:always,issuer
basicConstraints = critical, CA:true, pathlen:0
keyUsage = critical, digitalSignature, cRLSign, keyCertSign
```

#### 生成服务器证书
```bash
# 使用root CA 签发服务端证书
./openssl x509 -req -in server.csr -out server.crt -CA root.crt -CAkey root.key -CAcreateserial -days 3650
# 使用root CA 签发服务端证书
./openssl x509 -req -in server.csr -out server.crt -signkey server.key -CA root.crt -CAkey root.key -CAcreateserial -days 3650
# 使用mid CA 签发服务端证书，同时使用cert.ext配置文件中的v3_server字段信息
./openssl x509 -req -extfile ../cnf/cert.ext -extensions v3_server -days 365 -CA mid.crt -CAkey mid.key -CAcreateserial -in server.csr -out server.crt
```
#### 生成客户端证书
```bash
# 使用root CA 签发客户端证书
./openssl x509 -req -in client.csr -out client.crt -CA root.crt -CAkey root.key -CAcreateserial -days 3650
# 使用root CA 签发客户端证书
./openssl x509 -req -in client.csr -out client.crt -signkey client.key -CA root.crt -CAkey root.key -CAcreateserial -days 3650
# 使用mid CA 签发客户端证书，同时使用cert.ext配置文件中的v3_client字段信息
./openssl x509 -req -extfile ../cnf/cert.ext -extensions v3_client -days 365 -CA mid.crt -CAkey mid.key -CAcreateserial -in client.csr -out client.crt
```

### 验证证书是否是由CA签发
```bash
./openssl verify -CAfile root.crt mid.crt 
```

### 生成证书链
```bash
cat mid.crt root.crt > ca-chain.crt
```

### 查看证书
[x509 openssl man手册](https://docs.openssl.org/3.4/man1/openssl-x509)
```bash
./openssl x509 -in cert.crt -text -noout/serial/startdate/enddate/dates/subject/issuer/ext subjectAltName
```
## 证书转换

### pfx p12

#### 将pem crt证书转换为pfx证书
```bash
./openssl pkcs12 -export -out cert.pfx -inkey cert.key -in cert.crt
```
#### 将pem crt证书转换为p12证书
```bash
./openssl pkcs12 -export -clcerts -in cert.crt -inkey cert.key -out cert.p12
```

### pkcs7
    
#### 解析pkcs7文件——转换为pem格式
```bash
./openssl pkcs7 -in source.p7b -inform PEM -print_certs -out cert.pem
```
#### 解析pkcs7文件——转换为der格式：
```bash
./openssl pkcs7 -in source.p7b -inform DER -print_certs -out cert.pem
```

## 公私钥匹配
```bash
# 公私钥匹配——私钥
./openssl pkey -in private.key -pubout -outform pem | sha256sum
# 公私钥匹配——公钥证书
./openssl x509 -in certificate.crt -pubkey -noout -outform pem | sha256sum
# 公私钥匹配——CSR
./openssl req -in CSR.csr -pubkey -noout -outform pem | sha256sum
```


## jks证书
    
### 生成jks证书：
```bash
keytool -genkey -alias jwt -keyalg RSA -keystore jwt.jks

keytool -genkeypair -alias jwt -keyalg RSA -keypass 123456 -storepass 123456 -keyalg RSA -keysize 2048 -validity 3650 -keystore jwt.jks
    
参数注解：
-genkeypair 生成密钥对
-alias 别名
-keyalt 采用公钥算法，默认是DSA
-keypass 指定生成密钥的密码
-storepass 指定访问密钥库的密码
-keysize 密钥长度(DSA算法对应的默认算法是sha1withDSA，不支持2048长度，此时需指定RSA)
-keystore 指定keystore文件 jwt.jks
```

### jks查看证书内容

```bash
keytool -list -v -keystore jwt.jks
```

### jks导出证书cer
```bash
keytool -alias jwt -exportcert -keystore jwt.jks -file jwt.cer
```

### jks导入证书p12
```bash
keytool -importkeystore -srckeystore jwt.jks -destkeystore jwt.p12 -deststoretype pkcs12
```

### 转换为pem证书
```bash
openssl pkcs12 -in jwt.p12 -out jwt.pem -nodes
```

### jks导入证书pfx
```bash
keytool -v -importkeystore -srckeystore jwt.jks -srcstoretype jks -srcstorepass 123456 -destkeystore jwt.pfx -deststoretype pkcs12 -deststorepass 123456 -destkeypass 123456
```
### 提取私钥
```bash
openssl pkcs12 -in jwt.pfx -nocerts -nodes -out jwt.key
```
### 提取证书
```bash
keytool -export -alias jwt -file jwt.cer -keystore jwt.jks
```

## p12证书
    
### 生成p12证书
```bash
./openssl pkcs12 -export -clcerts -in cert.crt -inkey cert.key -out cert.p12
```
### 查看p12证书信息
```bash
./openssl pkcs12 -info -in cert.p12

./openssl pkcs12 -in cert.p12 -nokeys -info

./openssl pkcs12 -in cert.p12 -nocerts -nodes
```

### 打印证书链的详细信息
```bash
# 包括每个证书的版本、序列号、签发者信息：
openssl pkcs12 -in jwt.p12 -nokeys -clcerts
```

### 显示私钥的详细信息
```bash
# 包括版本、算法、长度：
./openssl pkcs12 -in jwt.p12 -noenc -nocerts
```

### 导出p12私钥
```bash
./openssl pkcs12 -in jwt.p12 -noenc -nocerts -out key.pem
```

### 导出p12证书
```bash
./openssl pkcs12 -in jwt.p12 -clcerts -nokeys -out jwt.pem
```

## pfx证书

[pkcs12 openssl man手册](https://docs.openssl.org/1.1.1/man1/pkcs12/)

### 查看pfx证书信息
```bash
openssl pkcs12 -info -in cert.pfx

openssl pkcs12 -in cert.pfx -nokeys -info

openssl pkcs12 -in cert.pfx -nocerts -nodes
```
### 生成pfx证书
```bash
./openssl pkcs12 -export -out cert.pfx -inkey cert.key -in cert.crt
```
### 导出pfx证书
```bash
openssl pkcs12 -in test.pfx -clcerts -nokeys -passin pass:123456 -out cert.pem
```
### 导出pfx私钥
```bash
openssl pkcs12 -in test.pfx -nocerts -nodes -passin pass:123456 -out jwt.pem

# 参数解析：
    -clcerts：Only output client certificates (not CA certificates).
    -cacerts：Only output CA certificates (not client certificates).
    -nocerts：No certificates at all will be output.
    -nokeys：No private keys will be output.
    -nodes: Don't encrypt the private keys at all.
    -info: Output additional information about the PKCS#12 file structure,algorithms used and iteration counts.
    -passin password: Pass phrase source to decrypt any input private keys with.
```

## 国密证书
    
### 铜锁openssl
```bash
# 生成sm2私钥
./openssl-tongsuo ecparam -genkey -name SM2 -out sm2.key
# 生成sm2公钥
./openssl ec -in sm2key.pem -pubout -out sm2PubKey.pem
# 生成sm2 CSR
./openssl-tongsuo req -new -key sm2.key -out sm2.csr -config ./conf/openssl-sm2.conf -subj "/C=CN/ST=BB/O=CC/OU=DD/CN=server"
# 签发sm2证书
./openssl-tongsuo x509 -req -in sm2.csr -CA ./ca/sm2-ca.crt -CAkey ./ca/sm2-ca.key -days 5475 -extfile ./conf/cert.ext -extensions v3_server -out sm2.crt -CAcreateserial
# 查看sm2证书
./openssl-tongsuo x509 -in sm2.crt -text -noout
```
### openssl
```bash
# 生成根CA私钥
./openssl ecparam -genkey -name SM2 -out sm2cert-ca.key
# 生成根CA公钥
./openssl ec -in sm2cert-ca.key -pubout -out sm2cert-ca-pub.pem
# 生成根CA CSR
./openssl req -new -key sm2cert-ca.key -out sm2cert-ca.csr -subj "/C=CN/ST=BB/O=CC/OU=DD/CN=server"
# 签发根CA证书
./openssl x509 -req -in sm2cert-ca.csr -signkey sm2cert-ca.key -out sm2cert-ca.pem -extfile ../conf/rootca.ext -extensions v3_ca -days 7300
# 生成证书私钥
./openssl ecparam -genkey -name SM2 -out sm2key.pem
# 生成公钥
./openssl ec -in sm2key.pem -pubout -out sm2PubKey.pem
# 生成CSR
./openssl req -new -key sm2key.pem -out sm2cert.csr -subj "/C=CN/ST=BB/O=CC/OU=DD/CN=test"
# 使用根CA证书签发证书 
./openssl x509 -req -in sm2cert.csr -CA ./ca/sm2cert-ca.pem -CAkey ./ca/sm2cert-ca.key -days 3650 -extfile ./conf/cert.ext -extensions v3_server -out sm2cert.pem -CAcreateserial
# 查看证书
./openssl x509 -in sm2cert.pem -text -noout
```
## pkcs1转换为pkcs8
```bash
./openssl pkcs8 -topk8 -inform PEM -in private.pem -outform pem -nocrypt -out pkcs8.pem
```

## EC证书
```bash
# 生成私钥
./openssl ecparam -name secp256k1 -genkey -out ec.key

./openssl ecparam -in secp256k1.pem -genkey -out secp256k1-key.pem

# pem/der
./openssl ecparam -name secp256k1 -genkey -out secp256k1-key.pem

./openssl ecparam -name secp256k1 -genkey -outform DER -out secp256k1-key.der

# 生成公钥：pem/der
./openssl ec -in secp256k1-key.pem -pubout -out ecpubkey.pem

./openssl ec -in secp256k1-key.pem -pubout -outform DER -out ecpubkey.der
```

## CRL
[crl openssl man手册](https://docs.openssl.org/3.4/man1/openssl-crl)

[ca openssl man手册](https://docs.openssl.org/3.4/man1/openssl-ca)
### 转换pem格式crl到der格式
```bash
openssl crl -in crl.pem -outform DER -out crl.der
```
### 转换der格式crl到pem格式
```bash
openssl crl -inform DER -in crl.der -out crl.pem
```
### 查看crl信息
```bash
openssl crl -in crl.pem -text -noout
```
### 签发证书
```bash
openssl genrsa -out cert2.key 2048
openssl req -new -key cert2.key -out cert2.csr -subj "/C=CN/CN=test"
openssl ca -config ./conf/openssl.conf -in cert.csr -out cert.crt
openssl ca -config ./conf/openssl.conf -extensions v3_req -days 3650 -in cert2.csr -out cert2.crt
```
### 吊销证书
```bash
openssl ca -config ./conf/openssl.conf -revoke cert.crt
```
### 更新吊销列表
```bash
openssl ca -config ./conf/openssl.conf -gencrl -out crl.pem
```

## 公私钥加解密

### 公钥加密私钥解密

#### 公钥加密
```bash
openssl pkeyutl -encrypt -inkey pubkey.pem -pubin -in test.txt -out test-encrypted.bin
openssl pkeyutl -encrypt -in test.txt -inkey pubkey.pem -pubin -out test-encrypted.bin
```
#### 私钥解密
```bash
openssl pkeyutl -decrypt -in test.txt -inkey private.key -out test-decrypted.txt
```

### 从证书中提取公钥
```bash
openssl x509 -in ca.crt -pubkey -noout -out public.key
```

### 证书指纹
```bash
openssl x509 -in certificate.crt -MD5 -fingerprint -noout
openssl x509 -in certificate.crt -sha1 -fingerprint -noout
openssl x509 -in certificate.crt -sha256 -fingerprint -noout
openssl x509 -in certificate.crt -sha384 -fingerprint -noout
openssl x509 -in certificate.crt -sha512 -fingerprint -noout
# hash 算法
    -md4            to use the md4 message digest algorithm
    -md5            to use the md5 message digest algorithm
    -ripemd160      to use the ripemd160 message digest algorithm
    -sha            to use the sha message digest algorithm
    -sha1           to use the sha1 message digest algorithm
    -sha224         to use the sha224 message digest algorithm
    -sha256         to use the sha256 message digest algorithm
    -sha384         to use the sha384 message digest algorithm
    -sha512         to use the sha512 message digest algorithm
    -whirlpool      to use the whirlpool message digest algorithm
```

### 私钥签名公钥验证签名：dgst

#### 私钥生成hash并签名
```bash
openssl dgst -sha256 -sign private.key -out signature.bin test.txt
```
#### 可读性更好的签名
```bash
openssl enc -base64 -in signature.bin -out signature.bin.base64
openssl enc -base64 -d -in signature.bin.base64 -out signature.bin
```
#### 公钥签名验证hash
```bash
openssl dgst -sha256 -verify public.key -signature signature.bin test.txt
```

#### 计算文本hash值
```bash
openssl sha1 -out singature.bin test.txt
openssl sha256 -out hash.bin test.txt
# 获取hash值,去除空格和冒号
openssl dgst -sha256 test.txt | cut -d= -f2 | tr -d '[:space:]:' | xxd -r -p > hash.bin
```

### 私钥签名公钥验证签名：pkeyutl

#### 私钥直接对数据进行签名
```bash
openssl pkeyutl -sign -in hash.bin -inkey private.key -out signature.bin
```
#### 公钥验证签名并恢复文本
```bash
openssl pkeyutl -verifyrecover -in signature.bin -inkey private.key -pubin
```
#### 公钥验证签名
```bash
openssl pkeyutl -verify -in test.txt -sigfile signature.bin -inkey public.key -pubin
```
#### Sign some data using an SM2(7) private key and a specific ID 
```bash
openssl pkeyutl -sign -in file -inkey sm2.key -out sig -rawin -digest sm3 -pkeyopt distid:someid
```
#### Verify some data using an SM2(7) certificate and a specific ID
```bash
openssl pkeyutl -verify -certin -in file -inkey sm2.cert -sigfile sig -rawin -digest sm3 -pkeyopt distid:someid
```

### 证书签名验证

1、从证书中提取数字签名
```bash
openssl x509 -in certificate.crt -text -noout -certopt ca_default -certopt no_validity -certopt no_serial -certopt no_subject -certopt no_extensions -certopt no_signame | grep -v 'Signature Algorithm' | grep -v 'Signature Value' | tr -d '[:space:]:' | xxd -r -p > signature.bin
```
2、获取签名中的指纹
```bash
openssl pkeyutl -verifyrecover -inkey public.key -pubin -in signature.bin -out signature-decrypted.bin
```
3、获取使用的hash算法
```bash
openssl asn1parse -inform der -in signature-decrypted.bin | grep '4:d=2  hl=2 l=   9 prim: OBJECT' | awk 'BEGIN {FS = ":"} {print $4}'
```
4、asn1解析获取证书body
```bash
openssl asn1parse -in certificate.crt -strparse 4 -out cert-body.bin -noout
```
5、使用公钥对数字签名进行验证
```bash
openssl dgst -sha256 -verify public.key -signature signature.bin cert-body.bin
# 检查证书的 ASN.1 结构
openssl asn1parse -in certificate.crt
# 提取证书签名
openssl asn1parse -in certificate.crt -strparse 671 -out signature2.bin
# 获取证书指纹
openssl x509 -in certificate.crt -sha256 -fingerprint -noout | cut -d= -f2 | tr -d ':' | xxd -r -p > fingerprint.bin
```

## 算法套件
```bash
openssl ciphers -V | column -t | less
```