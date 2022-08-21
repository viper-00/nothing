# Nothing

Nothing 是一个用 Go 编写的简单的 Linux 系统监控工具。Nothing 主要支持监控小型服务器、家用电脑和树莓派（Raspberry Pi）等设备，现在已扩展成为 Linux 系统级监控程序。Nothing 还支持自定义时间序列数据收集，可用来收集传感器或程序的输入输出和性能数据。

当前，Nothing 以一定时间为间隔，收集以下数据指标：
* 系统信息
* CPU 信息
* 内存/内存回收资源（SWAP）信息
* 一小时前（默认值）或自定义时间区间的CPU/内存使用图表
* 磁盘使用和使用图表
* 网络使用信息和图表
* 服务运行情况
* 系统进程按 CPU 或内存使用率排序
* 自定义指标（时间序列数据）收集和图表
* 在指定的时间或数据区间显示指标数据
* 指定时间加载进程数据
* 设置警报：CPU/内存/内存回收资源（SWAP）/磁盘/服务
  * 通过邮箱、Slack 或叮叮等工具推送系统警报给用户
* 自定义数据保留周期

### 屏幕截图

<img src="https://github.com/dhamith93/SyMon/blob/master/png/01.png?raw=true" width="400" />
<img src="https://github.com/dhamith93/SyMon/blob/master/png/02.png?raw=true" width="400" />
<img src="https://github.com/dhamith93/SyMon/blob/master/png/03.png?raw=true" width="400" />
<img src="https://github.com/dhamith93/SyMon/blob/master/png/04.png?raw=true" width="400" />
<img src="https://github.com/dhamith93/SyMon/blob/master/png/05.png?raw=true" width="400" />
<img src="https://github.com/dhamith93/SyMon/blob/master/png/06.png?raw=true" width="400" />
<img src="https://github.com/dhamith93/SyMon/blob/master/png/07.png?raw=true" width="400" />

## 自定义

组件之间通信（agents、collector、frontend 和 alert processor）是使用 gRPC 来实现。如果需要，可以设置自定义服务来读取数据或将数据推送到任何组件。[API](internal/api/api.proto)、[Alert API](internal/alertapi/alertapi.proto)。

## 组成部分

### Agent

### Collector

### Alert Processor

### Client

## 安装和使用

### 支持传输层安全（TLS）

### Collector

### Agents

### 自定义指标

### Alerts

### Clients

### API 文档

## LICENSE

MIT