package main

import "fmt"
//import "gowork/person"

import "person/cmd"
//在主函数中创建结构体的数值
func main(){
	fmt.Println("---")
	//包的名字，结构体的名字
	p1 := person.Person{"坚持",22,"女"}
	fmt.Println(p1)
	fmt.Println(p1.Name,p1.age,p1.Sex)
}
//CreateFile main.go: The system cannot find the file specified.
