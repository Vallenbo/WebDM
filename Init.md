# 概述

一个基于gin框架的docker可视化工具，该工具可以帮助您更轻松地管理容器和镜像。该工具采用了最新的技术，包括Elemnt- plus和Vue3，以提供直观易用的用户界面。该工具的主要功能包括：

1、显示所有正在运行的容器和镜像的列表，包括容器的状态、名称、ID、端口映射等信息。

2、允许您在容器上执行命令，并查看命令输出的结果。

3、提供一个可视化的界面来创建和删除容器和镜像，包括图形化的表单和选项。

4、支持容器之间的通信，方便您进行测试和调试。

5、除此之外，该工具也提供了一些额外的功能，比如监控服务器的资源使用情况、查看容器的日志文件等。



# 介绍

我开发了一个名为“DockerWeb_admin”的Web应用程序，

这个可视化的工具软件，类似于naivcat 工具

设计时考虑的，面向个人开发者使用的工具，

 

简单地介绍一下我采用的技术栈，包括后端和前端的技术，大致阐述它们的优势。

**后端方面**

我使用了golang语言+ Gin + Gorm +websocket实现了后端API接口以及与交互，这两者都是独立并且强大的框架，

 

**前端方面**，我选择使用Vue3.2 + Element Plus + EChart来实现UI组件和图表展示，

Vue3.2具有更高的性能和可测试性，

Element Plus则提供了大量的UI组件，

EChart提供了易于配置和高度可定制的数据可视化解决方案。

​	

 主要的目的在简化 Docker 的经常性任务，

如构建、启动、停止容器、搜索、拉取镜像等，从而节省用户时间，

里面的功能有，通过本地项目，编写dockerfile文件构建起一个临时测试的环境，

通过将整个环境打包成镜像的方式，传给需要使用的人，其他人不再需要搭建环境就能启动，这个项目，这样就可以节省重复部署环境的时间。

查看服务的各项基础硬件性能，如cpu 、 内存、 网络io、 硬盘io

当然也可以查看docker的版本信息

同时提供聚合 Docker 容器内部的信息报告。



# 目录介绍

```python
web # 为此项目 前端项目内容

其他 文件为后端所需文件佳
```



# 需要添加的第三方 配置信息

```python
# docker添加监听端口
vim /usr/lib/systemd/system/docker.service

ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock -H tcp://0.0.0.0:2375
```



```python
service docker restart
systemctl daemon-reload
ss -tupln | grep 2375
```

