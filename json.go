package main

import(
    "encoding/json"
    "fmt"
    "os"
)

//struct name and struct 对象 名字大写才可以访问
type Person struct {
    Name string
    Age int
    email string
}

func main(){
    //TODO: &的用法 
    //这里用不用都可以成功？
    person := &Person{Name: "tom", Age: 30, email: "tom@g.cn"}
    fmt.Println(person)

    personByte, _ := json.Marshal(person)
    fmt.Println(string(personByte))
    os.Stdout.Write(personByte)


    //TODO: 不需要Struct，解析Json to []Byte or Slice

    //确实Struct是一个比较好的方案
    //因为Go是静态语言，变量需要有类型，所以Struct相当于定义了变量的类型的作用
    //如果不使用Struct，返回的是一个{}interface 相当于一个泛型，需要遍历，猜测变量的类型，然后放入到类型里面
    //https://github.com/bitly/go-simplejson  就是使用这个做法的一个Lib



}
