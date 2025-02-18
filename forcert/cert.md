## 证书
- 证书系列基础课程：https://www.ssldragon.com/zh/category/ssl-basics-zh/
- 证书介绍：https://juejin.cn/post/6844904143719497735、https://juejin.cn/post/6844904143719497735
- OpenSSL SAN 证书：https://blog.csdn.net/baishitongtian/article/details/119544269

### 证书类型
- 怎么区分DV、OV、EV证书：https://blog.csdn.net/liuzehn/article/details/107807248
- DV, OV, IV, and EV Certificates：https://www.ssl.com/article/dv-ov-and-ev-certificates/#ftoc-heading-6

### 证书相关算法
- 比较 ECDSA 与 RSA：简单指南：https://www.ssl.com/zh-CN/%E6%96%87%E7%AB%A0%EF%BC%8C/%E6%AF%94%E8%BE%83-ecdsa-%E5%92%8C-rsa-%E7%9A%84%E7%AE%80%E5%8D%95%E6%8C%87%E5%8D%97/
- OpenSSL 命令行生成Elliptic Curve (EC) algorithms(椭圆曲线算法)私钥、公钥：https://steemit.com/openssl/@oflyhigh/openssl-elliptic-curve-ec-algorithms

### 证书字段
- web安全 | CN(Comman Name) SAN FQDNS是什么？：https://blog.csdn.net/Sunny_Ducky/article/details/104605111
 ——根据RFC 6125（在2011年发布），验证器必须首先检查SAN，如果存在SAN，则不应检查CN。

### PKI体系架构
- https://support.huawei.com/enterprise/zh/doc/EDOC1100127227/35e686f4

### 密钥生成方法
- https://help.aliyun.com/document_detail/64281.html

### rsa算法详解
- https://blog.csdn.net/qq_40587575/article/details/79542842

### rsa介绍系列博客
- http://www.shangyang.me/2017/05/24/encrypt-rsa-signature/

### pkcs
#### 公钥密码学标准
> 公钥密码学标准：https://zh.wikipedia.org/wiki/%E5%85%AC%E9%92%A5%E5%AF%86%E7%A0%81%E5%AD%A6%E6%A0%87%E5%87%86
>
> 公钥加密标准（Public Key Cryptography Standards, PKCS），此一标准的设计与发布皆由RSA信息安全公司所制定。
> 
> RSA信息安全公司旗下的RSA实验室为了发扬公开密钥技术的使用，便发展了一系列的公开密钥密码编译标准。只不过，虽然该标准具有相当大的象征性，也被信息界的产业所认同；但是，若RSA公司认为有必要，这些标准的内容仍然可能会更动。所幸，这些变动并不大；此外，这几年RSA公司也与其他组织（比较知名的有IETF、PKIX）将标准的制定透过standards track程序来达成。

||版本|名称|简介|
|---|---|---|---|
|PKCS #1|	2.1|	RSA密码编译标准（RSA Cryptography Standard）|	定义了RSA的数理基础、公/私钥格式，以及加/解密、签/验章的流程。1.5版本曾经遭到攻击[1]。|
|PKCS #2|	-|	弃用|	原本是用以规范RSA加密摘要的转换方式，现已被纳入PKCS#1之中。|
|PKCS #3|	1.4|	DH密钥协议标准（Diffie-Hellman key agreement Standard）|	规范以DH密钥协议为基础的密钥协议标准。其功能，可以让两方透过金议协议，拟定一把会议密钥(Session key)。|
|PKCS #4|	-	|弃用|	原本用以规范转换RSA密钥的流程。已被纳入PKCS#1之中。|
|PKCS #5|	2.0|	密码基植加密标准（Password-based Encryption Standard）|	参见RFC 2898与PBKDF2。|
|PKCS #6|	1.5|	证书扩展语法标准（Extended-Certificate Syntax Standard）|	将原本X.509的证书格式标准加以扩展。|
|PKCS #7|	1.5|	密码消息语法标准（Cryptographic Message Syntax Standard）|	参见RFC 2315。规范了以公开密钥基础设施（PKI）所产生之签名/密文之格式。其目的一样是为了拓展数字证书的应用。其中，包含了S/MIME与CMS。主要用来存储签名或者加密后的数据,如证书或者CRL。PKCS7可以用二进制的DER格式存储，也可以使用PEM格式存储。Windows中PKCS7通常使用.p7b文件扩展名。|
|PKCS #8|	1.2|	私钥消息表示标准（Private-Key Information Syntax Standard）|	Apache读取证书私钥的标准。主要用来存储私钥。私钥首先会使用PKCS #5的标准进行加密，然后Base64编码，转换成为PEM格式存储。|
|PKCS #9|	2.0|	选择属性格式（Selected Attribute Types）|	定义PKCS#6、7、8、10的选择属性格式。|
|PKCS #10|	1.7	|证书申请标准（Certification Request Standard）|	参见RFC 2986。规范了向证书中心申请证书之CSR（certificate signing request）的格式。|
|PKCS #11|	2.20|	密码设备标准接口（Cryptographic Token Interface (Cryptoki)）|	定义了密码设备的应用程序接口（API）之规格。|
|PKCS #12|	1.0|	个人消息交换标准（Personal Information Exchange Syntax Standard）|	定义了包含私钥与公钥证书（public key certificate）的文件格式。私钥采密码(password)保护。PKCS12文件的扩展名是.p12或者.pfx。PKCS12可以看做是PKCS7的扩展，PKCS12可以额外存储私钥。在PKCS12中可以存储证书、私钥、CRL。|
|PKCS #13|	–	|椭圆曲线密码学标准（Elliptic curve cryptography Standard）	|制定中。规范以椭圆曲线密码学为基础所发展之密码技术应用。椭圆曲线密码学是新的密码学技术，其强度与效率皆比现行以指数运算为基础之密码学算法来的优秀。然而，该算法的应用尚不普及。|
|PKCS #14|	–	|拟随机数产生器标准（Pseudo-random Number Generation）|	制定中。规范拟随机数产生器的使用与设计。|
|PKCS #15|	1.1	|密码设备消息格式标准（Cryptographic Token Information Format Standard）|	定义了密码设备内部数据的组织结构。|

#### RSA加密算法， PKCS#1 和PKCS#8区别是什么？
> 1、https://www.cnblogs.com/aworkstory/p/rsa-pkcs1-pkcs8.html
> 
> 2、https://cloud.tencent.com/developer/article/2390956
> 
> PKCS1是标准RSA秘钥对标准规范，但是都是裸奔的；
> 
> PKCS8是对加密后的秘钥进行了描述，就是说P8格式的秘钥不是裸奔了


### https单双向认证及自签CA和证书
- https://help.aliyun.com/zh/api-gateway/traditional-api-gateway/user-guide/mutual-tls-authentication#section-50r-0yc-3ra
