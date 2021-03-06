# Go语言学习笔记

- 基本语法、函数式编程、面向接口、并发编程、分布式爬虫实战 全面掌握Go语言
- 学习代码笔记链接：https://gitee.com/laosuaidami/learnGo
- 学习视频：https://coding.imooc.com/class/180.html

### 第1章 课程介绍 
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_01_install_env**

欢迎大家来到深度讲解Go语言的课堂。本课程将从基本语法讲起，逐渐深入，帮助同学深度理解Go语言面向接口，函数式编程，错误处理，测试，并行计算等元素，并带领大家实现一个分布式爬虫的实战项目。

- 1-1 Google资深工程师深度讲解go语言 试看
- 1-2 安装与环境
- 1-3 国内镜像配置
- 1-4 IntelliJ Idea 的安装和配置
- 1-5 vscode 的安装和配置

### 第2章 基础语法 
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_02_var_const_func_pointer**

量，常量，类型，选择，循环，函数，指针，本章节带领大家学习一门新语言所需的必备语法知识。让大家对Go语言有一个初步的认识！

- 2-1 变量定义
- 2-2 内建变量类型
- 2-3 常量与枚举
- 2-4 条件语句
- 2-5 循环
- 2-6 函数
- 2-7 指针
### 第3章 内建容器 
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_03_array_slicen_map_string**

本章节我们来学习数组，切片，Map和字符串。在Go语言中，我们一般不直接使用数组，而是使用切片来管理线性表结构，它的语法类似python的list，不过更强大哦。当然，Map和字符串的学习也是必不可少。掌握至此，我们就可以写一些简单的算法了，刷刷leetcode不在话下，我们就来试一试。...

- 3-1 数组
- 3-2 切片的概念
- 3-3 切片的操作
- 3-4 Map
- 3-5 Map例题
- 3-6 字符和字符串处理
### 第4章 面向“对象” 
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_04_struct**

Go语言没有class，只有struct。我们来看看struct如何使用，Go语言给结构体定义类似方法或者成员函数的做法非常有特色。我们还将学习Go语言的包的概念，以及如何封装，如何扩展已有类型等。我们还将学习GOPATH和Go语言项目的目录结构，如何从网上下载依赖包等一系列项目相关的知识。我们将以“树”的结构和遍历作为贯穿本章...

- 4-1 结构体和方法
- 4-2 包和封装
- 4-3 扩展已有类型
- 4-4 使用内嵌来扩展已有类型
### 第5章 Go语言的依赖管理
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_05_gopath_govendor_gomod**

关于Go语言的依赖管理大有可用，只讲核心的，只学有用的，把时间投资在最有价值的学习上。

- 5-1 依赖管理
- 5-2 GOPATH 和 GOVENDOR
- 5-3 go mod的使用
- 5-4 目录的整理
### 第6章 面向接口
**https://github.com/guohewei123/LearnGo/tree/master/chapter_06_interface**

这一章我们从duck typing的概念开始学起，还将探讨其他语言中对duck typing的支持，由此引出接口的概念。我们将深入理解Go语言接口的内部实现以及使用接口实现组合的模式。

- 6-1 接口的概念
- 6-2 duck typing的概念 试看
- 6-3 接口的定义和实现
- 6-4 接口的值类型
- 6-5 接口的组合
- 6-6 常用系统接口
### 第7章 函数式编程
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_07_func**

在其他通用语言中，函数式编程是“高级”概念，但对于Go语言却非常基本。本章我们将讲解函数式编程的概念并且比较其他语言函数式编程的实现方法。我们将重点理解闭包。这章中我们将采用多样的例题来帮助大家更好的理解闭包，函数作为一等公民等及其常见概念和应用方法。...

- 7-1 函数式编程
- 7-2 函数式编程例一
- 7-3 函数式编程例二
### 第8章 错误处理和资源管理
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_08_defer_panic_recover**

这将是本课程最“无聊”的一章，但却是区分出优秀软件工程师的关键能力。Go语言独特的defer/panic/recover，以及错误机制，在社区有着广泛的争论。我们来深入理解Go语言的错误处理机制，看看Go语言如何区分错误以及异常。最后，我们实现一个Web应用微型项目，采用商业服务的错误处理思路，结合函数式编程，来演示Go语言错误...

- 8-1 defer调用
- 8-2 错误处理概念
- 8-3 服务器统一出错处理
- 8-4 panic和recover
- 8-5 服务器统一出错处理2
### 第9章 测试与性能调优
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_09_test_pprof**

Go语言的测试不同于其他如junit，Go语言采用“表格驱动测试”的理念。我们将学习和体会这样的理念，并用Go语言的测试支持库来实践表格驱动测试，并做代码覆盖和性能检测，通过內建的性能调优工具来优化我们之前的算法。最后演示了对http服务器的多种粒度的测试。...

- 9-1 测试
- 9-2 代码覆盖率和性能测试
- 9-3 使用pprof进行性能调优
- 9-4 测试http服务器（上）
- 9-5 测试http服务器（下）
- 9-6 生成文档和示例代码
- 9-7 测试总结
### 第10章 Goroutine
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_10_goroutine**

这一章开始我们进入并发编程。我们讲解Goroutine，协程的概念，以及背后的Go语言调度器。

- 10-1 goroutine
- 10-2 go语言的调度器

### 第11章 Channel
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_11_channel**

Channel是Goroutine之间通信的桥梁，它和函数一样是一等公民。在介绍完Channel的语法及运行方式后，我们将采用数个例题来演示Go语言并发编程中最常见的任务极其解决模式。

- 11-1 channel
- 11-2 使用Channel等待任务结束
- 11-3 使用Channel进行树的遍历
- 11-4 用select进行调度
- 11-5 传统同步机制
- 11-6 sync包的使用

### 第12章 迷宫的广度优先搜索
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_12_example_maze**

这章我们将综合运用学过的知识实现一个广度优先算法来解迷宫，为接下来的实战项目做好技术和算法上的准备。广度优先算法不仅是面试和工作中常用的技术，而且实现上相比大部分其它算法更为复杂，是检验是否熟练掌握一门语言的经典例题。让我们来试一试吧。...

- 12-1 迷宫_算法
- 12-2 迷宫代码实现

### 第13章 http及其他标准库
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_13_http_gin**

这里我们简要介绍一下Go语言中非常重要而且封装良好的http标准库，回顾并实现http客户端和服务器。我们还介绍了Go语言中其他的标准库。

- 13-1 http标准库
- 13-2 其它标准库
- 13-3 gin 框架介绍
- 13-4 为gin增加middleware

### 第14章 开始实战项目
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_14_reptile_project_start**

至此为止，恭喜同学完成了这门课Go语言部分的学习。接下来我们来进入实战项目。本章将介绍项目的具体内容，课题的选择，技术选型，总体架构，以及实现步骤。

- 14-1 爬虫项目介绍
- 14-2 爬虫的法律风险
- 14-3 新爬虫的选择
- 14-4 总体算法
- 14-5 模拟相亲网站上线啦！
### 第15章 单任务版爬虫
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_15_reptile_project_signle**

在考虑性能之前我们首先应该考虑正确性。单任务版爬虫确保我们能够正确爬取我们所需的信息。我们应用了之前练习的广度优先算法，抽象出Parser和Fetcher，学习正则表达式，成功实现并运行单任务版爬虫。

- 15-1 获得初始页面内容
- 15-2 正则表达式
- 15-3 提取城市和url
- 15-4 单任务版爬虫的架构
- 15-5 Engine 与 Parser
- 15-6 测试CityListParser
- 15-7 城市解析器
- 15-8 用户信息解析器（上）
- 15-9 用户信息解析器（下）
- 15-10 单任务版爬虫性能
### 第16章 并发版爬虫
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_16_reptile_project_concurrency**

为了提升爬虫性能，我们抽象出Worker的概念，并添加调度器，实现并发版爬虫。我们应用接口的概念，完成了由简至复杂的多个调度器的实现。同学可以在实战项目中更真实的体会并学习Go语言并发编程的多种模式。

- 16-1 并发版爬虫架构
- 16-2 简单调度器
- 16-3 并发调度器
- 16-4 队列实现调度器
- 16-5 重构和总结
- 16-6 更多城市
- 16-7 更多用户与去重
### 第17章 数据存储和展示
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_17_reptile_project_display_storage**

是时候检验我们项目的成果了。我们将采用Docker+ElasticSearch来存储我们爬取的信息。在简单了解Docker和ElasticSearch后，我们将使用ElasticSearch的Go语言客户端将爬取数据写入。之后我们使用Go语言的模板引擎迅速实现前端网页展示。至此，我们已经可以尝试自己喜欢的搜索条件去查看数据啦。...

- 17-1 ItemSaver的架构
- 17-2 Docker和ElasticSearch介绍
- 17-3 Docker的安装和使用
- 17-4 ElasticSearch入门
- 17-5 向ElasticSearch存储数据
- 17-6 完整爬虫的运行与数据存储
- 17-7 添加URL与ID
- 17-8 重构与运行
- 17-9 标准模板库介绍
- 17-10 实现前端展示页面
- 17-11 完善前端展示
### 第18章 分布式爬虫
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_18_reptile_project_distributed**

本章在简要介绍分布式概念后，将我们的并发爬虫改写成分布式。我们在很少改动的情况下，加入jsonrpc客户/服务端，实现并部署分布式爬虫。最后探讨实战项目的更多改进方案。

- 18-1 分布式系统简介
- 18-2 分布式爬虫架构
- 18-3 jsonrpc的使用
- 18-4 ItemSaver服务
- 18-5 整合ItemSaver服务
- 18-6 解析器的序列化
- 18-7 实现爬虫服务
- 18-8 完整分布式爬虫的运行
- 18-9 使用连接池链接爬虫集群
- 18-10 实战项目总结
- 18-11 进一步的工作
### 第19章 课程总结
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_19_summary**

感谢同学们学到这里，恭喜同学们给自己的技术栈加上了非常重要的Go语言技能。希望同学们带着这门课上学到的知识，更好的参与到项目中去，共同推动Go语言的发展。

- 19-1 体会Go语言的设计
- 19-2 课程总结

### 第20章 断言assertion和反射reflect
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_20_assertion_reflect**
- 20-1 断言assertion使用
- 20-2 反射reflect使用


### 第21章 context包试用
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_21_context** 

### 第22章 Go 中 io 包的使用方法
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_22_io**
- 22-1 io包的基础使用

### 第23章 Go 中文件的操作
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_23_op_file**

### 第24章 Go fmt 打印
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_24_fmt**

### 第25章 Go log zap 库使用
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_25_log_zap**

### 第26章 Go gorm 访问 mysql 库使用
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_26_gorm**

### 第27章 GRPC 学习
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_27_grpc**

### 第28章 配置文件读取 库viper学习
**学习笔记：https://github.com/guohewei123/LearnGo/tree/master/chapter_28_config_file_viper**
