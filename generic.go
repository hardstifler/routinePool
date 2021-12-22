package main

import (
	"constraints"
	"fmt"
)

/**
https://go.dev/doc/tutorial/generics
golang 1.18 泛型初探
*/

//定义类型约束集合
type MyNumber interface {
	int |int32|float64
}
func main() {
	var m1 map[string]int = map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	var m2 map[string]float64 = map[string]float64{
		"1": 1.1,
		"2": 2.2,
		"3": 3.3,
	}

	sum1 := addInt(m1)
	sum2 := addFloat(m2)
	fmt.Println(sum1)
	fmt.Println(sum2)

	fmt.Println(addWithGeneric(m1))
	fmt.Println(addWithGeneric(m2))
	
	//泛型结构体，泛型参数暂时不可以省略，不是很爽
	t1 := &GenericStruct[bool]{
		Value:true,
		Elem:"hello",
	}

	t2 := &GenericStruct[string]{
		Value:"hello",
		Elem:"hello",
	}

	fmt.Println(t1.GetValue())
	fmt.Println(t2.GetValue())

	intSlice := []int{1,2,3,4,5}
	floatSlice := []float64{1.1,2.2,3.3, 4.4}
	fmt.Println(Max(intSlice...))
	fmt.Println(Max(floatSlice...))

	sliceP := []int32{1,2,3,4}
	var e int32  = 2 
	fmt.Println(Scale(sliceP, e))
}

func addInt(args map[string]int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func addFloat(args map[string]float64) float64 {
	var sum float64
	for _, v := range args {
		sum += v
	}
	return sum
}

//golang map key 必须是可以比较的； 次数定义K类型为可比较类型， V类型为上面定义的类型集合 
func addWithGeneric[K comparable, V MyNumber](args map[K]V) V{
	var sum V
	for _, v := range args {
		sum += v
	}
	return sum
}

//结构体中如何使用泛型
type GenericStruct[T any] struct{
	Value T  `json:"value"`
	Elem string  `json:"elem"`
}

//泛型结构体的方法
func (e *GenericStruct[T])GetValue()T{
	return e.Value
}

//类型别名中使用泛型
type GenericSlice[T any] []T 

type GenericMaps[K comparable, V any] map[K]V 

//泛型函数
func Max[T MyNumber](elem ...T)T{
	var max T 
	for _,v := range elem {
		if v > max{
			max = v 
		}
	}
	return max
}

//泛型函数 ~波浪线表示底层类型是E slice即可 
//因此  type point[E] []E类型 在此处是合法的
func Scale[S ~[]E, E constraints.Integer](s S, e E)S{
	res := make(S, len(s))
	for i,v := range s{
		res[i] = v*e
	}
	return res
}
