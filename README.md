# 账号数据管理服务

## 项目简介
#### share服务用于数据存储（redis，mongodb)。

------

```
|-- Dockerfile
|-- share-sa.yaml
|-- share.yaml
|-- api
|-- cmd
|-- configs
|   |-- bindreward.json         # 绑定赠送配置文件
|   `-- littleacclimit.txt      # 小号限制配置文件
|-- docker-compose.yml
|-- internal
|   |-- code                    # 服务自定义错误码
|   |   |-- share.ecode.go    # 自定义错误码绑定说明
|   |   |-- share.proto       # 自定义错误码
|   |   `-- errmsg.go           # 错误码的proto文件
|   |-- config                  # 加载自定义配置文件逻辑
|   |   |-- bindreward.go
|   |   |-- config.go
|   |   `-- littleacclimit.go
|   |-- dao
|   |   |-- redis.go            # 绑定的redis处理逻辑
|   |   |-- mongo.go            # mongodb初始化逻辑
|   |   `-- dao.go              # dao数据访问层
|   |-- di
|   |-- model                   # 定义的相关redis数据处理结构
|   |   |-- limit.go            # 小号限制rediskey定义
|   |   |-- http.go             # facebook，line的数据结构
|   |   |-- share.go
|   |   |-- model.go
|   |-- server
|   `-- service
|       |-- handler.go          # service数据流处理
|       |-- service.go          # 接口所在地方
|       `-- service_test.go     # 接口单元测试
`-- pkg                         # 服务启动逻辑包
    `-- pkg.go
```

