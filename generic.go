package main

import "fmt"

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
	//此处的泛型约束不能省略，否则报错，后续应该会有类型默认的类型推导
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

func addWithGeneric[K comparable, V MyNumber](args map[K]V) V{
	var sum V
	for _, v := range args {
		sum += v
	}
	return sum
}

type GenericStruct[T any] struct{
	Value T  `json:"value"`
	Elem string  `json:"elem"`
}

func (e *GenericStruct[T])GetValue()T{
	return e.Value
}

type GenericSlice[T any] []T 

type GenericMaps[K comparable, V any] map[K]V 

func Max[T MyNumber](elem ...T)T{
	var max T 
	for _,v := range elem {
		if v > max{
			max = v 
		}
	}
	return max
}
