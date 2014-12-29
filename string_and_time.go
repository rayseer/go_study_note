package main

import(
    "fmt"
    "strings"
    "time"
    "reflect"
)

func main(){
    // func Unix(sec int64, nsec int64) Time
    // func (t Time) Add(d Duration) Time
    // t := time.Unix() //t 为实例
    // t := time.Now() //t 为实例
    // t.Add()
    // 调用实例一般为两个步骤
    // 第一步获取type实例
    // 第二部调用实例的方法
    str := "I love you"

    //repeat 10 times
    repeat_count := 10
    repeat_string := strings.Repeat(str, repeat_count)
    fmt.Print(repeat_string)

    //用空格分隔
    split_char := " "
    split_string_array := strings.Split(str, split_char)
    fmt.Println(split_string_array)

    //打印现在的时间
    time_now := time.Now()
    fmt.Println("Print Time:", time_now.String())
    time_arr := strings.Split(time_now.String(), ".")
    fmt.Println("Split Time 0 :", time_arr[0])
    fmt.Println("Timestamp :", time_now.Unix())

    //RFC规范 https://www.ietf.org/rfc/rfc3339.txt
    //RFC3339     = "2006-01-02T15:04:05Z07:00"
    //RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    t := time.Now().Local()
    fmt.Println("Time Local:", t.Format("20060102150405"))
    fmt.Println("Time Local:", t.Format("2006010215"))
    fmt.Println("Time Local:", t.Format("2006-01-02"))
    fmt.Println("Time Local:", t.Format("2006-01-02 15:04:05"))
    //打印pkg const
    fmt.Println("Time Local:", t.Format(time.RFC3339))

    //打印现在的时区
    fmt.Println("Time Location:", t.Location())


    //time.Now().In()
    //TODO:
    //直接使用time.LoadLocation感觉是有问题的
    //因为他属于type Location struct

    //大概明白了。
    //package 可调用大部分方法，除了struct方法
    //package 可调用全部的Const
    fmt.Println(time.UTC.String())
    china_local, err:= time.LoadLocation("Asia/Shanghai")
    fmt.Println("type:", reflect.TypeOf(china_local))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(china_local.String())
    fmt.Println(time.Now().In(china_local).String())

    fmt.Println(time.Monday)

    //用反射判断类型
    fmt.Println("type:", reflect.TypeOf(time.Hour))

    //获取Json
    timeJson, _ := time.Now().MarshalJSON()
    fmt.Println("Time Json:",string(timeJson))

}
