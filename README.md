# go-log 日志说明





### Step 1

首先在你的工程目录下新建**conf**文件夹, 以及**config.yaml**, 目录如:

```
├── conf
│   └── config.yaml
└── main.go
```



**config.yaml**

```
logs:
  level: debug                          # 日志等级 默认为debug
  encode: console                       # json or console 默认为console
  levelPort: 5053                       # log level服务端口 默认为不开启
  levelPattern: /handle/level           # log level uri 默认为不开启
  output: file,console                  # 输出方式控制台和文件
  file:
    path: ./logs/location.log           # 文件输出路径 默认当前logs/app.log
    maxSize: 100                        # 日志最大容量 (单位/M) 默认100M
    maxBackups: 10                      # 保留的最大旧日志文件数 默认10
    maxAge: 28                          # 旧日志超过多少时间删除 (单位/天) 默认30天
    encode: json                        # 文件的编码方式
  initFields:                           # log 初始字段, 一般用户日志集中化过滤
    company: wayz                       # 字段自定义, 示例为公司名称
    app: yourapp                        # 字段自定义, 示例为应用名称
```






**PS:**
1. ./conf/config.yaml 路径不可更改
2. 以上所有字段皆可不填
3. file.path 字段支持环境变量 如: path: ./logs/{hostname}/foo.log


### Step 2

在代码中引入logs包, 常规打印即可, 如下:

```
func main() {
	for {
		logs.Debug("debug message", RandomString(6))
		logs.Debugf("debug something happened %s", RandomString(6))
		logs.Debugw("Debugw", "url", RandomString(6), "attempt", 3)

		logs.Info("info message", RandomString(6))
		logs.Infof("info something happened %s", RandomString(6))
		logs.Infow("Infow", "url", RandomString(6), "attempt", 3)

		time.Sleep(time.Second * 2)
	}
}
```



- xxx   普通日志方法
- xxxf  format打印方法
- xxxw with打印方法

