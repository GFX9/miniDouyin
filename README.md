# miniDouyin# 1 组织结构```bash├── biz # 除dal外，其余子文件夹由IDL生成│   ├── dal # 数据库相关操作│   │   ├── init.go # 初始化Gorm连接│   │   └── pg  # pg定义的结构体和接口函数│   │       ├── DbInterface.go│   │       ├── DbLogic.go│   │       ├── init.go # 初始化pg连接│   │       └── user.go # user结构体│   ├── handler # 路由函数等接口│   │   ├── miniDouyin│   │   │   └── api│   │   │       └── mini_douyin.go│   │   └── ping.go│   ├── model # IDL生成│   │   └── miniDouyin│   │       └── api│   │           └── minidouyin.go│   └── router # IDL生成，路由注册│       ├── miniDouyin│       │   └── api│       │       ├── middleware.go│       │       └── minidouyin.go│       └── register.go├── build.sh├── createDB.sql # Mysql建表├── createDB_pg.sql # PG建表├── go.mod├── go.sum├── main.go # 主函数入口├── minidouyin.thrift # IDL├── router.go # 猜测：自定义路由注册├── router_gen.go # 猜测：自定义路由函数└── script    └── bootstrap.sh```