Tan盘数据服务
===

负责响应接口服务的定位请求并对接口服务提供数据的传输  
本项目为《分布式对象存储——原理、架构及 GO 语言实现》一书的实践
## 主要流程
数据服务启动后,连接RabbitMQ,创建Exchange,每5秒向Exchange发送心跳消息,接口服务通过队列绑定exchange,获取数据服务的心跳消息(ListenAddress),并监听apiServer经由exchange发来的定位消息,如果本地存在指定对象,则将地址发送到临时消息队列

## Auther
[BuTn](https://github.com/kimmosc2)

## 环境变量
RabbitMQ地址:RABBITMQ_SERVER  
本地监听地址:LISTEN_ADDRESS  
本地存储路径:STORAGE_ROOT  

## 其他组件
[接口服务](https://github.com/kimmosc2/upload-interface)