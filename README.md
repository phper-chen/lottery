
## 使用方法：
* 放在GOPATH下，开启go mod模式，要求go1.3以上
* 进入lottery/local/，执行docker-compose up -d启动mysql和redis
* 进入mysql执行lottery.sql文件，建表和预插入奖品数据
* 进入lottery/Makefile，执行make启动service服务
* 在浏览器中打开frontend/lottery.html前端界面

## 已实现：
* 粗略的临时组装了后端开发框架
* 抽奖核心逻辑，解决数据的并发竞态问题
* 服务启动后可以压测该接口
* wrk -t10 -c100 -d5 "http://localhost:8080/lottery/draw"
* 实现抽奖轮盘前端完整的渲染和抽奖功能
* 时间问题，其他功能缺少校验，未能对接上前端
* 可以导出所有获奖记录CSV
* 其他进来得及实现粗略的接口

## 未实现：

* 个人抽奖次数限制
* 用户session处理