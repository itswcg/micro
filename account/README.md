### account
统一用户中心(UIM)

#### 功能
account服务提供以下功能:  
1. 账户信息  
2. 用户认证  
3. 权限管理  

http:

* 注册
* 登录
* 修改资料
* 获取用户资料
* github授权登录

rpc:

* 查询用户资料


##### Mysql
```
create database account

CREATE TABLE `user_0` (
  `mid` int(11) NOT NULL,
  `name` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  `sex` tinyint(4) NOT NULL DEFAULT 0,
  `face` varchar(512) NOT NULL DEFAULT '',
  `email` varchar(128) NOT NULL DEFAULT '',
  `phone` varchar(11) NOT NULL DEFAULT '',
  `join_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  PRIMARY KEY (`mid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

#### Todo

* 认证选型(基于cookie,session?基于token?基于jwt?)
* migrate工具
* token 认证, rpc认证?