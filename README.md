# Clash配置简易服务端

用于统计vps流量使用量和vps释放时间。

## 使用方法

1. 下载release二进制文件，在vps（默认amd64 linux系统）中mkdir工作目录。

2. 在工作目录mkdir config目录，把本项目的config.toml上传到config目录，修改config内容（如下）。

```toml
[General]
  DefaultTraffic = 100 # 这是你的vps的总流量，单位GB
  Offset = 0 # 这是已经使用的流量，单位GB，当开启每月重置时，到点会重置为0
  NetworkCardName = "eth0" # 这是你的vps使用的网卡，只会统计这一个网卡的流量
  StartTraffic = 0 # 这个是用于记录最新一次启动程序时，网卡已使用流量
  ExpireTime = "2024-05-12 13:50:52.4085838 +0800 CST" # 这是你的vps释放时间，在clash for windows里会显示成日期
  ClashPath = "./clash_config.yml" # 这是你的clash配置的位置
  ResetMonthly = false # 这是每个月是否重置已使用流量为0并将下次重置时间设为执行日期的月份+1，当到达“ExpireTime”的日期的零时时执行重置（注意服务器时区设置，如果下个月没有当前的日，则会挑选最后一天）
```

3. 上传你的clash配置文件并在config.toml填写你的clash配置位置。
4. "chmod +x main"并"./main -r"启动程序，其中"-r"代表记录最新网卡已使用流量并修改vps释放时间（默认为当日的一个月后），"-rn"代表只记录最新网卡已使用流量，不会修改vps释放时间。
5. 端口默认8080，你可以使用nginx来反代到域名。