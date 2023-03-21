#### msgcenter
一个基于`Go gin`的websocket消息分发中心。通过http `POST` 请求将消息分发到订阅主题的websocket。具体使用方法参考下面说明。

### 运行
- `make gen_denpendence`
- `make download-dependence`
- `make run`

### 使用
- `msgcenter -query=true` 获取`访问token`
- `msgcenter -add=<token>` 添加`访问token`
- `msgcenter -del=<token>` 删除`访问token`
- `msgcenter -upd=<old_token,new_token>` 更新`访问token`

### websocket 客户端, 在浏览器终端中运行下面的代码
```js
const socket = new WebSocket('ws://localhost:8001/?token=testToken&&topic=hello');
socket.addEventListener('open', event => {
    console.log('Connected to server');
    socket.send('Hello, server!');
});
socket.addEventListener('message', event => {
    console.log(`Received message: ${event.data}`);
    socket.close();
});
```

### 发送消息
- `curl -X POST -v -H "Auth-Token: testToken" -d '{"msg": "hello topic"}' "http://localhost:8001/topic/hello"`
- 订阅了`hello`主题的websocket就会接收到`{"msg": "hello topic"}`消息


### 启用https
- 修改配置文件`EnableTLS: true`
- 程序启动会生成默认的证书和私钥文件，放置在`~/.config/msgcenter/cert.pem` 和 `~/.config/msgcenter/key.pem`。 默认证书仅支持本地回环地址访问。需要支持外部网络访问需要替换为自己的证书和私钥。

### 注意：
- 一个websocket连接仅支持订阅一个主题。需要订阅多个主题需要启用多个websocket连接。
- websocket连接不要向消息中心发送消息。发现消息通过http `POST` 请求进行发送。

### 参考
- [go安装依赖包（go get, go module）](https://blog.csdn.net/weixin_41519463/article/details/103501485)
- [Golang设置代理](https://developer.aliyun.com/article/879662)
- [Gin 解决跨域问题跨域配置](https://juejin.cn/post/6871583587062415367)
- [Go 入门指南](https://learnku.com/docs/the-way-to-go)
