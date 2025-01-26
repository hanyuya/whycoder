## pod
[亲和性](https://kubernetes.io/zh-cn/docs/concepts/scheduling-eviction/assign-pod-node/)

## service
[k8s service章节](https://kubernetes.io/zh-cn/docs/concepts/services-networking/service)

### service类型：

- 详解k8s的4种Service类型:
    - <https://blog.csdn.net/weixin_40274679/article/details/107887678>
    - <https://www.cnblogs.com/guyouyin123/p/15592539.html>


- [K8s 中 Service、Endpoints、Pod 之间的关系](https://blog.csdn.net/catoop/article/details/122072608)
- [k8s的service和endpoint](https://blog.csdn.net/jjt_zaj/article/details/122432291)
- [Kubernetes网络篇——Service网络(上)](https://morningspace.github.io/tech/k8s-net-service-1)

### 删除策略
Kubernetes 中的对象是如何删除的：Finalizers 字段介绍：https://blog.csdn.net/cr7258/article/details/123928301
- Finalizer 和 DeletionTimestamp 是紧密相关的。
- 当资源对象被标记为要删除时，Kubernetes 会在 DeletionTimestamp 字段中设置一个时间戳，并在 Finalizer 列表中添加 Finalizers。
- 一旦所有 Finalizers 都执行完毕，Kubernetes 就会删除该对象，并将 DeletionTimestamp 字段设置为 null。



## kind

[官方文档](https://kind.sigs.k8s.io/)

实现原理：Kind 使用一个docker 容器来模拟一个 node，在 docker 容器里面跑 systemd ，并用 systemd 托管 kubelet 以及 containerd，然后通过容器内部的 kubelet 把其他 K8s 组件，比如 kube-apiserver、etcd 等跑起来，最后在部署上 CNI 整个集群就完成了。Kind 内部也是使用的 kubeadm 进行部署。

- [Kubernetes教程(十五)---使用 kind 在本地快速部署一个 k8s集群](https://www.lixueduan.com/posts/kubernetes/15-kind-kubernetes-in-docker/)
- [Kubernetes入门：使用Kind快速安装、部署k8s集群](https://juejin.cn/post/7322672734768578599)
- [k8s--kind 搭建 k8s 集群](https://www.cnblogs.com/zouzou-busy/p/16388186.html)

## 二次开发
### kubebuilder
从 Operator 理念的提出到现在已经有了很多工具可以帮助我们快速低成本的开发，其中最常用的就是 CoreOS 开源的 operator-sdk 和 k8s sig 小组维护的 kubebuilder。

使用实战指南“使用kubebuilder开发operator”：
- 从环境搭建开始构建：https://lailin.xyz/post/operator-01-overview.html
- 会介绍client-go：https://podsbook.com/posts/kubernetes/operator/
- k8s编程operator中的一个章节：https://blog.csdn.net/Peerless__/article/details/128287145
- kubebuilder-标记注释：https://blog.csdn.net/a1023934860/article/details/122859248

kubebuilder实战系列课程：
- kubebuilder实战之一：准备工作：https://xinchen.blog.csdn.net/article/details/113035349
- kubebuilder实战之二：初次体验kubebuilder：https://xinchen.blog.csdn.net/article/details/113089414
- kubebuilder实战之三：基础知识速览：https://xinchen.blog.csdn.net/article/details/113815479
- kubebuilder实战之四：operator需求说明和设计：https://xinchen.blog.csdn.net/article/details/113822065
- kubebuilder实战之五：operator编码：https://xinchen.blog.csdn.net/article/details/113836090
- kubebuilder实战之六：构建部署运行：https://xinchen.blog.csdn.net/article/details/113840999
- kubebuilder实战之七：webhook：https://xinchen.blog.csdn.net/article/details/113922328
- kubebuilder实战之八：知识点小记：https://xinchen.blog.csdn.net/article/details/114215218

### controller机制
client-go：https://isekiro.com/categories/client-go/

k8s编程operator系列：
- k8s编程operator——(1) client-go基础部分：https://blog.csdn.net/Peerless__/article/details/127814142
- k8s编程operator——(2) client-go中的informer：https://blog.csdn.net/Peerless__/article/details/128002585
- k8s编程operator——(3) 自定义资源CRD：https://blog.csdn.net/Peerless__/article/details/128056647
- k8s编程operator——(4) kubebuilder & controller-runtime源码分析：https://blog.csdn.net/Peerless__/article/details/128287145

controller-runtime源码分析：

- Kubernetes Controller开发利器:controller-runtime：https://maao.cloud/2021/02/26/Kubernetes-Controller%E5%BC%80%E5%8F%91%E5%88%A9%E5%99%A8-controller-runtime/
- Kubernetes Controller runtime 详解：https://www.geekgame.site/post/k8s/extensions/controller_runtime/
- controller-runtime源码分析：https://qiankunli.github.io/2020/08/10/controller_runtime.html

### fakeclient
FakeClient update on status subresource：https://github.com/kubernetes-sigs/controller-runtime/issues/2362