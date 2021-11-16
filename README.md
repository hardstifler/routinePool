# routinePool
简易任务池 [pool.go](https://github.com/hardstifler/routinePool/edit/main/routonPool.go)  

并发执行方法 [parallel.go](https://github.com/hardstifler/routinePool/edit/main/parallel.go)

通过ldflags参数在编译时注入版本等信息
go build -ldflags "-X github.com/hardstifler/routinePool/version.OS=darwin -X 'github.com/hardstifler/routinePool/version.BuildTimeStamp=`date`' -X 'github.com/hardstifler/routinePool/version.GitVersion=`git rev-parse --short HEAD`'"
