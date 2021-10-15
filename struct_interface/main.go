package main
 
import "fmt"
 
type InterfaceParent interface{
	runParent()
}
 
type InterfaceChild interface{
	InterfaceParent
	runChild()
}
 
type StructParent struct{
	name string
}
 
func NewStructParent(name string) *StructParent{
	return &StructParent{name:name}
}
 
type StructChild struct {
	InterfaceChild
	*StructParent
	age int
}
 
func NewStructChild(ic InterfaceChild,sp *StructParent,age int) *StructChild{
	return &StructChild{
		InterfaceChild: ic,
		StructParent:   sp,
		age:            age,
	}
}
 
type TestStruct struct {
}

func NewTestStruct() *TestStruct{
	return &TestStruct{}
}

func (t *TestStruct)runParent(){
	fmt.Println("TestStruct====runParent()")
}
func (t *TestStruct)runChild(){
	fmt.Println("TestStruct====runChild()")
}

func main(){
	test := NewTestStruct() // 空结构体
	parent := NewStructParent("张三") // {name}
	child := NewStructChild(test,parent,10) // {InterfaceChild{},StructParent{},age}
    fmt.Println(child.name,child.age)
	fmt.Printf("child: %#v \n", child)
	child.runParent()
	child.runChild()
	fmt.Println("##################################")
	child.InterfaceChild.runChild()
	child.InterfaceChild.runParent()
}