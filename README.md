### <center>用于处理定时任务的go程序</center>

##### 配置文件
+ 配置文件位于conf/conf.json中
+ 各配置项说明
	- logPath：日志文件路径
	- tasks：要执行的定时任务列表
		- url：定时任务url地址(支持相对地址，若为相对url，则需要执行命令时传入rootUrl参数)
		- timeunit：定时任务间隔时间单位
		- interval：定时任务间隔多久执行一次
		- immediateExe：是否启动程序后立即执行此url定时任务
		- method：url的调用方式（get、post）
		- data：get或post的参数

##### 执行方法：
+ 编缉conf/conf.json配置文件
	- conf.json中不可有注释代码
+ 执行go程序
	- 方法1：编缉为可执行程序并执行
	
		`go build timedtask.go`
		
		`./timedtask` 或 `./timedtask -rootUrl=http://www.mamian.net/`
		
		注：若采用 nohup ./timedtask &  会产生大量进程
	- 方法2：直接运行代码
	
		`go run timedtask.go`
