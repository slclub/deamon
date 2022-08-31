守护进程

### server
- 守护进程 服务端
- server 守护进程  引入deamon.server 可以直接监听 服务 
- 可以直接当deamon 服务运行
- 运行模式 docker in docker
- 运行：docker run -p 10801:10801 --privileged=true -v /var/run/docker.sock:/var/run/docker.sock -v /usr/bin/docker:/usr/bin/docker --name tiandi-deamon -d tiandi-deamon:latest

### client
- 守护进程客户端  demo
- deamon.client 模块可以被其他项目直接 引入使用

### deamon.logger 日志注入
- 可以使用自己的日志 注入 到 监听中
- 注入的自定义日志 系统 需要实现deamon.logger.Logger 接口
- 注入方法deamon.logger.Log(自己的日志系统)