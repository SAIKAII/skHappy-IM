### 用户相关
- 用户填入账户、密码进行注册
- 用户填入账户密码进行登录

### 聊天相关
- 用户A发送消息到用户B
  1. A的消息到达服务器后，对消息进行保存，并在成功保存后发生ack给A;起一个goroutine把消息发送给B，并确保消息已经到达B。分情况：若B在线，
消息发送给B，B在获取消息后把本地last_ack_id更新为当前消息的id（数据库中自增长的那个id）;若B不在线，不发送消息，等待B上线后自己拉去消息
，属于B的并且<last_ack_id的消息都拉取。 B存在于redis中则表明在线，否则不在线。
  2. 关于last_ack_id，A与B之间一直递增，A与B的last_ack_id与A与C的last_ack_id不是同一个。cache中暂存各个last_ack_id，不需要每次
通讯都去数据库中获取。redis中存=>key:A&B value:last_ack_id。