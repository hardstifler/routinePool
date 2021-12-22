# routinePool
简易任务池 [pool.go](https://github.com/hardstifler/routinePool/blob/main/pool.go)  

并发执行方法 [parallel.go](https://github.com/hardstifler/routinePool/blob/main/parallel.go)

golang 1.18Beta 泛型初体验 [generic.go](https://github.com/hardstifler/routinePool/blob/main/generic.go)


# ldflags -X 
main.go中输出了几个版本号信息 [mian.go](https://github.com/hardstifler/routinePool/blob/main/cmd/main.go)

我们通过通过ldflags参数在编译时注入版本等信息 
编译命令如下  
go build -ldflags 

"-X github.com/hardstifler/routinePool/version.OS=darwin

-X 'github.com/hardstifler/routinePool/version.BuildTimeStamp=`date`' 

-X 'github.com/hardstifler/routinePool/version.GitVersion=`git rev-parse --short HEAD`'" 

-o ./cmd . 
