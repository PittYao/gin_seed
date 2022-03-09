# gin的单体结构项目
- [x] [gin](https://github.com/gin-gonic/gin)
- [x] [zap](https://github.com/uber-go/zap)
- [x] [gorm](https://gorm.io/zh_CN/docs/)
- [x] [redis](https://github.com/go-redis/redis/v8)
- [x] [db2gorm](https://github.com/qmhball/db2gorm)
- [x] [swagger](https://github.com/swaggo/gin-swagger)
- [x] [gout](https://github.com/guonaihong/gout)

# 项目结构
```
├─app   // 代码主体
│  ├─common // 常用包
│  │  ├─db // mysql包
│  │  ├─ginconfig // gin配置
│  │  ├─global // 全局对象
│  │  ├─globalkey // 全局常量
│  │  ├─redis // redis
│  │    └─redis_test // redis 使用示例
│  │  ├─request // 分页请求体
│  │  ├─response // 统一响应体
│  │  ├─tool // 工具包
│  │  └─zap // zap日志包
│  │ 
│  ├─etc // 配置文件
│  │  └─config.yaml // 配置文件
│  │  └─config.go // 配置文件实例
│  ├─internal // 业务代码目录
│  │  ├─dao 
│  │  ├─httpclient // 三方httpclient示例
│  │  ├─middleware // 中间件
│  │  ├─model // 生成model
│  │  │  ├─ffmpegtemplate
│  │  │  └─recordone
│  │  ├─routes // 路由
│  │  │  └─docs // swagger
│  │  └─service 
│  │      └─example_service
│  ├─script // 脚本
│  │  └─gen // db生成struct程序
│  └─statics // 静态资源
│      ├─css
│      ├─html
│      └─js
├─deploy //  部署资料
│  └─sql 
└─logs // 日志文件夹

```

# 配置文件
- 路径 app/etc/config.yaml
- 可修改的地方
  - Name: 应用名称 [ 也是日志文件的前缀 ]
  - ListenOn: 服务端口
  - Mysql.Url: 数据库地址 格式[ username:password@tcp(ip:port)/dbName ]
  - Swagger.Title: swagger文档名称
  - Swagger.Desc: 文档描述
  - Zap.Format: 开发时用console 正式时用json
  - Zap.Dir: 存放日志文件夹名
    
# swagger
- 访问地址: http://localhost:8008/doc/index.html
- 在 internal/service下写swagger注解
- 生成doc文档:
```
  # 在app/main.go下 控制台执行生成下面命令
  # swag init --dir ./ --generalInfo internal/routes/routes.go --propertyStrategy snakecase --output internal/routes/docs
  # 重启程序后生效 [访问 http://localhost:8008/doc/index.html ]
```
- 为安全初次访问时有密码验证:
  - 配置文件config.yaml中 
    - BasicAuth.UserName
    - BasicAuth.Password
    
# db2gorm
- 数据库生成对应的struct
  - 案例:app/script/gen/mysql_test.go
  - dsn配置db地址
  - 生成指定单表
  - 生成指定整个Db
  - 生成路径指定在 app/internal/model

# zap
- 日志默认在项目的路径下logs存储,文件名称 applicationName_info.log
- 在app/common/zap针对gorm v2 实现自定义接口来适配
- 默认开发时配置 
  - Zap.Format: console 日志在控制台打印
- 正式环境时配置
  - Zap.Format: json 日志在控制台、文件都打印

# 开发
- 路由配置
  - app/internal/routes/routes.go
- service
  - app/internal/service
- dao 
  - app/internal/dao
- 分页查询示例
  - app/internal/service/example_service/onestream/RecordOnePage