## 容器

### docker

**什么是docker**
```
Docker是一个容器化平台，它以容器的形式将您的应用程序及其所有依赖项打包在一起，以确保您的应用程序在任何环境中无缝运行
```

#### docker架构

**镜像**
相当于一个最小的root文件系统

**容器**  
镜像和容器的关系，就像面向对象里面类和实例的关系一样，镜像是静态的，容器是启动后的动态的

**仓库**  
仓库可以看成一个代码控制中心，用来保存镜像  

#### docker常见面试题
```
什么是docker镜像？
docker容器的源代码，用于创建（build）容器

什么是docker容器？
包含程序运行的所有依赖项，作为独立进程运行

docker容器有几种状态？
运行、已停止、重新启动、已退出

DockerFile的常见指令
FROM   指定基础镜像
LABEL  功能为镜像指定标签
RUN    运行指定命令
CMD    容器启动时要运行的命令

DockerFile中的COPY和ADD命令的区别？
其中COPY的SRC只能是本地文件，其他用法一致

容器退出后，数据是否会丢失？
容器退出后，docker ps看不到，但是docker ps -a可以查看，并且可以重启，只有删除容器才会清除数据
```

### kubernetes

一般我们称呼其为k8s，容器编排工具

#### k8s的基本概念

**Node**
```
Node作为集群节点中的工作节点，运行真正的应用程序，在Node上k8s管理运行的最小单元是Pod
Node上运行着k8s的kubelet、kube-proxy服务进程，这些服务负责Pod的创建、监控、重启、销毁、以及负载均衡
```

**Pod**
```
k8s最基本的操作单元，包含一个或者多个紧密相关的容器，一个pod的多个容器应用通常是紧密耦合的
每个pod里运行着一个特殊的被称之为Pause的容器，其他容器为业务容器，这些容器共享Pause容器的网络栈和Volume挂载卷
同一个Pod里的容器之间仅需通过localhost就能互相通信
k8s为Pod设计了一套独特的网络配置，为每个Pod分配一个IP地址，使用Pod名作为容器间通信的主机名等
```

**Service**
```
一个Service可以看作是提供相同服务的Pod的对外访问接口，Service作用于哪些Pod是通过Label Selector来定义的
Service拥有一个指定的名字，拥有虚拟IP和端口号，提供某种远程服务能力
如果Service要提供外网服务，需要指定公共IP和NodePort，或外部负载均衡器
NodePort，系统会在Kubernetes集群中的每个Node上打开一个主机的真实端口，这样，能够访问Node的客户端就能通过这个端口访问到内部的Service了
```

**Volume**
```
Volume是Pod中能够被多个容器访问的共享目录
```

**Label**
```
Label以key/value的形式附加到各种对象上，如Pod、Service、RC、Node等，以识别这些对象，管理关联关系等，如Service和Pod的关联关系
```

**RC（Replication Controller）**
```
目标Pod的定义；
目标Pod需要运行的副本数量；
要监控的目标Pod标签（Lable）；
Kubernetes通过RC中定义的Lable筛选出对应的Pod实例，并实时监控其状态和数量，如果实例数量少于定义的副本数量（Replicas），则会根据RC中定义的Pod模板来创建一个新的Pod，然后将此Pod调度到合适的Node上启动运行，直到Pod实例数量达到预定目标。
```

#### k8s的架构

```
Kubernetes将集群中的机器划分为一个Master节点和一群工作节点（Node）
Master节点上运行着集群管理相关的一组进程etcd、API Server、Controller Manager、Scheduler，后三个组件构成了Kubernetes的总控中心
这些进程实现了整个集群的资源管理、Pod调度、弹性伸缩、安全控制、系统监控和纠错等管理功能，并且全都是自动完成
在每个Node上运行Kubelet、Proxy、Docker daemon三个组件，负责对本节点上的Pod的生命周期进行管理，以及实现服务代理的功能
```

流程  
通过Kubectl提交一个创建RC的请求，该请求通过API Server被写入etcd中，此时Controller Manager通过API Server的监听资源变化的接口监听到这个RC事件，分析之后，发现当前集群中还没有它所对应的Pod实例，于是根据RC里的Pod模板定义生成一个Pod对象，通过API Server写入etcd，接下来，此事件被Scheduler发现，它立即执行一个复杂的调度流程，为这个新Pod选定一个落户的Node，然后通过API Server讲这一结果写入到etcd中，随后，目标Node上运行的Kubelet进程通过API Server监测到这个“新生的”Pod，并按照它的定义，启动该Pod并任劳任怨地负责它的下半生，直到Pod的生命结束。

随后，我们通过Kubectl提交一个新的映射到该Pod的Service的创建请求，Controller Manager会通过Label标签查询到相关联的Pod实例，然后生成Service的Endpoints信息，并通过API Server写入到etcd中，接下来，所有Node上运行的Proxy进程通过API Server查询并监听Service对象与其对应的Endpoints信息，建立一个软件方式的负载均衡器来实现Service访问到后端Pod的流量转发功能


* etcd  
用于持久化存储集群中所有的资源对象，如Node、Service、Pod、RC、Namespace等；API Server提供了操作etcd的封装接口API，这些API基本上都是集群中资源对象的增删改查及监听资源变化的接口。

* API Server  
提供了资源对象的唯一操作入口，其他所有组件都必须通过它提供的API来操作资源数据，通过对相关的资源数据“全量查询”+“变化监听”，这些组件可以很“实时”地完成相关的业务功能。

* Controller Manager  
集群内部的管理控制中心，其主要目的是实现Kubernetes集群的故障检测和恢复的自动化工作，比如根据RC的定义完成Pod的复制或移除，以确保Pod实例数符合RC副本的定义；根据Service与Pod的管理关系，完成服务的Endpoints对象的创建和更新；其他诸如Node的发现、管理和状态监控、死亡容器所占磁盘空间及本地缓存的镜像文件的清理等工作也是由Controller Manager完成的。

* Scheduler  
集群中的调度器，负责Pod在集群节点中的调度分配。

* Kubelet  
负责本Node节点上的Pod的创建、修改、监控、删除等全生命周期管理，同时Kubelet定时“上报”本Node的状态信息到API Server里。

* Proxy  
实现了Service的代理与软件模式的负载均衡器。  

客户端通过Kubectl命令行工具或Kubectl Proxy来访问Kubernetes系统，在Kubernetes集群内部的客户端可以直接使用Kuberctl命令管理集群。Kubectl Proxy是API Server的一个反向代理，在Kubernetes集群外部的客户端可以通过Kubernetes Proxy来访问API Server。

API Server内部有一套完备的安全机制，包括认证、授权和准入控制等相关模块。