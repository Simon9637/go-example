#go-example
本项目基于golang语言编写，是go语言学习的一个简单例子。

##说明
1. 本项目实现go 的api router（http）。可用于构建基于go语言的后端程序。
2. 本项目使用Viper， 可动态修改配置文件并实现程序热加载。
3. 本项目中实现zk cluster的初始化连接以及相关操作，包括zk client， register node, get service node and it's node data.zk初始化可使用https://github.com/Simon9637/zk-cluster中docker-compose.yml进行初始化。
4. test目录中有amqp协议的producer consumer 测试用例。测试前可使用https://github.com/Simon9637/docker-rabbitmq进行初始化。
5. go 中其他再陆续加入。。。。。。

###1. Viper(github.com/spf13/viper)
Viper 是国外大神 spf13 编写的开源配置解决方案，具有如下特性:

设置默认值
可以读取如下格式的配置文件：JSON、TOML、YAML、HCL
监控配置文件改动，并热加载配置文件
从环境变量读取配置
从远程配置中心读取配置（etcd/consul），并监控变动
从命令行 flag 读取配置
从缓存中读取配置
支持直接设置配置项的值
Viper 配置读取顺序：

viper.Set() 所设置的值
命令行 flag
环境变量
配置文件
配置中心：etcd/consul
默认值
从上面这些特性来看，Viper 毫无疑问是非常强大的，而且 Viper 用起来也很方便，在初始化配置文件后，读取配置只需要调用 viper.GetString()、viper.GetInt() 和 viper.GetBool() 等函数即可。

Viper 也可以非常方便地读取多个层级的配置，比如这样一个 YAML 格式的配置：
common:
  database:
    name: test
    host: 127.0.0.1
如果要读取 host 配置，执行 viper.GetString("common.database.host") 即可。