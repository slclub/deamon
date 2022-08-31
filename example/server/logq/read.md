### 项目名 glog
使用 github.com/slclub/glog

### 优秀点：
当出现panic 或者warn 错误可以准确的追踪代码路径以及行号；
可以自定义很多中写法; 且没有使用lock 锁；利用环形队列方式的；

为里快速构建项目我们只用封里两种；

日志可以按文件大小，和日期多种方式 自动切换生成日志文件； 

可以设置默认保留多少天的日志

### 缺点：
为里性能的，牺牲里实时性，有时log会过几秒中才会从内存，刷到日志中

巨量测试未层实验过

### 接口说明

Info(args ...interface{})

InfoF(format string, args ...interface{})

Debug(args ...interface{})
DebugF(format string, args ...interface{})


Warn  <同上>

WarnF <同上>

Error  <同上>
