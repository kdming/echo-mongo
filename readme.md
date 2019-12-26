golang + echo + mongo 框架  
看了很多他人搭建的架构，吸取经验，搭建一个完善的架构  
理解：
1. Dao层可以封装curd，可以随时更换dao层
2. 多个目录嵌套不是坏事情，对于go这门语言来说
3. 工具类不要和service层掺和到一块，这样更清晰了
4. 如果一个service有很复杂的计算，可以把计算单独写一个文件  
5. 一个接口可能调用多个方法的话，把它放到service层
6. 不要吝啬struct的声明
7. 方法名可以简写，curd，共用函数名起名需有意义
8. 不要吝啬error捕获，语言特性如此


##### 项目目录结构如下
```$xslt
├─ config
│  ├─ config.go 项目配置
│
├─ dao 数据库封装层
│ 
├─ middleware 存放web中间件
│  ├─ jwt 权限验证
│─ models 模型层
│─ pkg 工具包
│  ├─ file 文件管理
│─ routers 路由
│  ├─ api 接口
│     ├─ v1 版本
│  ├─ router.go 总路由管理
│─ service 服务层
│─ main.go 项目入口
 │

```
