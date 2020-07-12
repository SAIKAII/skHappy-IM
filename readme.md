# skHappy-IM
## 工具
- 数据库：MySQL5.7
- ORM：GORM
- 配置：Viper
- RPC：GRPC
- JSON编解码：json-iterator
- 认证：jwt-go
- 数据传输协议：Protobuf3
- redis：redigo
- 长连接：go-conn-manager
- 单元测试：goconvey

## 设计
程序当前分成两块，一个logic，一个tcp_conn，作为不同进程运行。logic以RPC接口的方式提供服务，tcp_conn作为logic的辅助，协助logic完成一些无法完成的工作。
### logic
- 提供用户注册、添加好友、删除好友、获取好友信息、罗列所有好友、创建群组、加入群组、退出群组、解散群组、发送单对单消息、发送群组消息、更新个人信息、更改密码功能。

### tcp_conn
- 提供用户登录服务，登录后保持连接以便服务端推送消息到用户端。