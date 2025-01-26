## openssl
- openssl常用方法：https://ruanhao.cc/blog/openssl.html#orgd1822e1
- 系列：(待详细看 C++)
  - 1、OpenSSL 入门：密码学基础知识：https://linux.cn/article-11810-1.html
  - 2、如何使用 OpenSSL：哈希值、数字签名等：
- openssl的 openssl.cnf配置文件详解：https://blog.csdn.net/yeyuningzi/article/details/135294481
- openssl 使用命令帮助：https://www.cnblogs.com/gavin11/p/14302045.html
- 实现基于国密算法SM2SSL证书的https加密：https://developer.baidu.com/article/details/2891669
- 使用 OpenSSL 制作一个包含 SAN（Subject Alternative Name）的证书：https://blog.csdn.net/u011599033/article/details/116454007
- openssl - 创建 CA、生成密钥、颁发、吊销证书等：https://www.soulchild.cn/post/2391

### 语法
官方文档：https://docs.openssl.org/3.4/man1/

语法
```
openssl [OPTION] [OPTION]
```
参数说明：

第一个OPTION选择使用openssl的哪种功能，参数说明如下：
- ca：openssl ca用于模拟一个CA行为的工具，能够签发证书请求文件以及生成CRL列表。在自签发证书、吊销证书、生成CRL列表等时使用。
- genrsa：openssl genrsa用于生成rsa私钥文件，可以指定私钥长度以及给私钥文件增加密码保护。
- rsa：openssl rsa用于处理RSA私钥、格式转换和打印信息等。
- req：openssl req用于生成证书请求，让第三方权威机构CA来签发，生成需要的证书。
- x509：openssl x509 是证书处理工具，用于显示证书内容、转换格式等场景。
- pkcs7/pkcs8/pkcs12：openssl pkcs7/pkcs8/pkcs12是私钥文件的转换工具，例如，在JDBC开启ssl功能后的配置客户端时用于生成keyStore。
- verify：openssl verify用于对证书的有效性进行验证，verify指令会沿着证书链一直向上验证，直到一个自签名的CA。

第二个OPTION根据第一个参数OPTION选择，结合实际需要配置。

### openssl 证书签发
- https://blog.csdn.net/weixin_40304387/article/details/127800427
- https://blog.csdn.net/m0_49946916/article/details/109280115
- https://blog.csdn.net/Solyutian/article/details/84033765

### openssl 公私钥匹配验证
- https://www.cnblogs.com/Rocky_/p/17851928.html

### openssl解析pkcs7文件
- https://support.kaspersky.com/KSMG/2.0/zh-Hans/239056.htm

### openssl实现数字签名
- https://cloud.tencent.com/developer/article/1756191