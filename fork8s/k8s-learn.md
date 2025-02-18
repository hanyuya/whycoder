## GV & GVK & GVR

### GV: Api Group & Version

- API Group 是相关 API 功能的集合
- 每个 Group 拥有一或多个 Versions

### GVK: Group Version Kind

- 每个 GV 都包含 N 个 api 类型，称之为 Kinds，不同 Version 同一个 Kinds 可能不同

### GVR: Group Version Resource

- Resource 是 Kind 的对象标识，一般来 Kind 和 Resource 是 1:1 的，但是有时候存在 1:n 的关系，不过对于 Operator 来说都是 1:1 的关系

举个🌰，我们在 k8s 中的 yaml 文件都有下面这么两行，例如上篇文章我们部署的 nginx deployment
```yaml
apiVersion: apps/v1 # 这个是 GV，G 是 apps，V 是 v1
kind: Deployment    # 这个就是 Kind
sepc:               # 加上下放的 spec 就是 Resource了
  ...
```
根据 GVK K8s 就能找到你到底要创建什么类型的资源，根据你定义的 Spec 创建好资源之后就成为了 Resource，也就是 GVR。GVK/GVR 就是 K8s 资源的坐标，是我们创建/删除/修改/读取资源的基础。