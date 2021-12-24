package main

import (
	"constraints"
	"fmt"
)

//https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md

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
		
	//常规函数
	sum1 := addInt(m1)
	sum2 := addFloat(m2)
	fmt.Println(sum1)
	fmt.Println(sum2)
	
	//泛型函数
	fmt.Println(addWithGeneric(m1))
	fmt.Println(addWithGeneric(m2))
	
	//泛型结构体初始化
	t1 := &GenericStruct[bool]{
		Value:true,
		Elem:"hello",
	}

	t2 := &GenericStruct[string]{
		Value:"hello",
		Elem:"hello",
	}
	
	//泛型结构体方法
	fmt.Println(t1.GetValue())
	fmt.Println(t2.GetValue())
	
	
	intSlice := []int{1,2,3,4,5}
	floatSlice := []float64{1.1,2.2,3.3, 4.4}
	fmt.Println(Max(intSlice...))
	fmt.Println(Max(floatSlice...))

	sliceP := []int32{1,2,3,4}
	var e int32  = 2 
	fmt.Println(Scale(sliceP, e))

	//泛型接口使用
	var _  = New[*Vertex, *FromTo](nil)
	//这样不行，泛型参数必须是指针类型，我猜是因为 方法接收者类型是指针
	//var _  = New[Vertex, FromTo](nil)
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


/**
方法与函数区别，支持反向函数
不支持泛型方法，
func (e *GenericStruct[T])Print[A any](a A){
	fmt.Println(a)
}
*/

/**
泛型接口
*/
type GenericInterFace[E MyNumber] interface{
	Max(e []E)E
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


func Scale[S ~[]E, E constraints.Integer](s S, e E)S{
	res := make(S, len(s))
	for i,v := range s{
		res[i] = v*e
	}
	return res
}

// ##################################泛型接口######################################################
// NodeConstraint is the type constraint for graph nodes:
// they must have an Edges method that returns the Edge's
// that connect to this Node.
//顶点接口，有方法可以获取链接到该顶点的所有边
type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

// EdgeConstraint is the type constraint for graph edges:
// they must have a Nodes method that returns the two Nodes
// that this edge connects.
//边接口， 有方法可以获取这条边链接的两个顶点
type EdgeConstraint[Node any] interface {
	Nodes() (from, to Node)
}

// Graph is a graph composed of nodes and edges.
//图结构体， 由顶点和边组成
type Graph[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] struct { }

// New returns a new graph given a list of nodes.
//泛型函数 构造图 入参nodes需要满足NodeConstraint约束
func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]] (nodes []Node) *Graph[Node, Edge] {
	return nil 
}

// ShortestPath returns the shortest path between two nodes,
// as a list of edges.
func (g *Graph[Node, Edge]) ShortestPath(from, to Node) []Edge {
	return nil 
}

//顶点
// Vertex is a node in a graph.
type Vertex struct {  }

//顶点有方法返回链接到这个顶点的边集合
// Edges returns the edges connected to v.
func (v *Vertex) Edges() []*FromTo { 
	return nil 
 }

 //边
// FromTo is an edge in a graph.
type FromTo struct {  }

//边有一个方法  返回链接的两个顶点
// Nodes returns the nodes that ft connects.
func (ft *FromTo) Nodes() (*Vertex, *Vertex) { return nil, nil  }

