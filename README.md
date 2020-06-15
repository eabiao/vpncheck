# vpncheck-go

定时检测主机连通状况，通过任务栏图标显示连通状态

参考了以下项目的实现
- https://github.com/xilp/systray
- https://github.com/cratonica/trayhost

生成内嵌图标用到的工具
- https://github.com/cratonica/2goarray

生成内嵌图标的命令
```
cat icon/on.ico |2goarray.exe iconOnData main >icondata.go
cat icon/off.ico |2goarray.exe iconOffData main >>icondata.go
```