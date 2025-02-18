
## kubebuilder

### 创建一个项目：
```bash
mkdir podsbook
cd podsbook
```
### 初始化一个项目
```bash
kubebuilder init --domain test.com
kubebuilder init --domain test.com --repo github.com/podsbook
--domain test.com 我们的项目的域名
--repo xxx 是仓库地址，同时也是 go mode中的repo地址
```
### 创建一个api
```bash
kubebuilder create api --group apps --version v1 --kind CertCenter
```
### 自定义crd重新生成
```bash
make manifests generate
```