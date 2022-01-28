# gobase 一个基于Go的脚手架工具

> 基于Go的微服务实践，最重要的就是两个先决条件：一是完整的流水线基础设施，二是一个完善的脚手架代码框架

流水线构想

1. 能够通过web页面进行配置，类似CDH页面部署
2. 能够支持离线部署，环境脚本安装
3. 一些流水线配置和流水线模板

需求拆分

- 基础环境
  - docker
  - k8s
  - 实体机
- 基础设施
  - gitlab
  - harbor
  - gitlab-runner/jenkins/argo
- 基础服务
  - mysql/postgres
  - mongo
  - redis
  - etcd/zookeeper
  - elastic search
  - rabbitmq/kafka
  - snowflake
  - clickhouse

逻辑设计

假设 5 台服务器 Mem 64G, Cpu 16 core 32 Thread, Disk 4T

脚手架代码

1. 包含基础的微服务，如：权限中心、配置中心、日志中心、监控中心
2. 包含一些微服务中公共代码的sdk封装
3. 包含一些工程管理的工具，如，代码生成，服务重启，部署打包，检查更新