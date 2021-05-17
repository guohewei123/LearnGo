# 
## 分布式系统
### 1. 多个节点
- 容错性
- 可扩展性（性能）
- 固有分布性
### 2. 消息传递
- 节点具有私有存储
- 易于开发
- 可扩展性（功能）
- 对比：并行计算
- 消息传递方法
    - REST 接口
    - RPC
    - 中间件
  ![](images/c833ca03.png)

- 一般消息传递的方法
    - 对外: REST 
    - 模块内部: RPC 
    - 模块之间: 中间件， REST 
### 3. 完成特定需求

## 分布式架构 vs 微服务架构

- 分布式：指导节点之间如何通信
- 微服务：鼓励按业务划分模块
- 微服务架构通过分布式架构来实现

## 多层架构 vs 微服务架构
- 微服务架构具有更多的 "服务"
- 微服务通常需要配合自动化测试，部署，服务发现等

## 并发版的问题
![](images/cee3acc0.png)
### 1. 限流问题
- 单节点能够承受的流量有限
- 将Worker放到不同的节点
![](images/3c2685cd.png)
### 2. 去重问题
- 单节点能够承受的去重数据量有限
- 无法保存之前去重结果
- 基于Key-Value Store (如Redis) 进行分布式去重
  ![](images/1f5b918d.png)
  ![](images/bee6e1d4.png)
### 3. 数据存储问题
- 存储部分的结构，技术栈和爬虫部分区别很大
- 进一步优化需要特殊的ElasticSearch技术背景
- 固有分布式
![](images/1d893b11.png)

## 本课程架构
- 解决限流问题（理论上）
- 解决存储问题
- 分布式去重
![](images/00dd6b16.png)


## RPC
![](images/78196ed7.png)
- jsonrpc(本课程采用) demo代码：chapter_18/rpc
- grpc(使用protobuf)
- Thrift

## 自由协议
![](images/a7e20431.png)
![](images/2fcec60f.png)
- docker/libchan
- NATS streaming
- gociruit
- 根据自己需求

## 解析器的序列化/反序列化
- 解析器原先的定义为函数
- 需要处理函数的序列化/反序列化 
![](images/eb195a8d.png)


## 总结
![](images/3169e767.png)

![](images/aa25a9b2.png)

![](images/9df3df66.png)

![](images/0da288d2.png)

![](images/57401dc2.png)

![](images/642ef76b.png)

![](images/13c3bdd1.png)

![](images/be8859d4.png)

![](images/87f65394.png)

![](images/96b82d31.png)

![](images/96025bdc.png)

![](images/815a0874.png)

![](images/139f9f30.png)

![](images/863b8d77.png)

![](images/ddbff895.png)

![](images/0e6e3857.png)

![](images/02fe97b1.png)

![](images/f3a57e74.png)

![](images/9b91288b.png)

![](images/0178d838.png)

![](images/c3b91c90.png)

![](images/879409e5.png)
