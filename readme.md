#### golang echo + mongo 框架  

##### 看了很多他人搭建的架构，吸取经验，搭建一个完善的架构  


1. Dao层可以封装curd，可以随时更换dao层
2. 多个目录嵌套不是坏事情，对于go这门语言来说
3. 工具类不要和service层掺和到一块，这样更清晰了
4. 如果一个service有很复杂的计算，可以把计算单独写一个文件  
5. 一个接口可能调用多个方法的话，把它放到service层
6. 不要吝啬struct的声明
7. 方法名可以简写，curd，共用函数名起名需有意义
8. 不要吝啬error捕获，语言特性如此

---

##### 项目目录结构如下:

```$xslt
├─ config
│  ├─ config.go 项目配置
│
├─ dao 
│  ├─ database.go 数据库连接
│  ├─ curd.go curd封装
│ 
├─ middleware 存放web中间件
│  ├─ jwt 权限验证
│ 
│─ models 模型层
│  ├─ user.go 用户model
│ 
│─ pkg 工具包
│  ├─ file 文件管理
│  ├─ e echo工具
│  ├─ encrypt 加密解密
│
│─ routers 路由
│  ├─ api 接口
│     ├─ v1 版本
│  ├─ router.go 总路由管理
│─ service 服务层
│  ├─ user_service 用户service
│
│─ main.go 项目入口

```

#### 启动命令
##### 1.手工编译启动
```cassandraql
go build main.go
sh run.sh
```
##### 2.系统编译启动
```cassandraql
sh run.sh build
```

Ps:   
run.sh中端口号需根项目端口号一支，使用nohup部署  
每次启动都会先杀死已有进程