package main 


import (
  "fmt"
  "github.com/hardstifler/routinePool/version"
 )


func main(){
  	fmt.Println(version.GitVersion)
	fmt.Println(version.BuildTimeStamp)
	fmt.Println(version.OS)
}
