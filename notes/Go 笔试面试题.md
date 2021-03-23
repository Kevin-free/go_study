# Go 笔试面试题



#### 1、`=` 和 `:=` 的区别？

`=` 仅赋值

`:=`声明和赋值

```go
var i int
i = 10
// equals
i := 10
```



#### 2、指针的作用？

指针用来保存变量的地址。

例如

```go
var x = 5
var p *int = &x //p 是 x 的地址
fmt.Println("x = %d", *p) //x 可以用 *p 访问
```

- `*` 运算符，也称为解引用运算符，用于访问地址中的值。
- `&` 运算符，也称为地址运算符，用于返回变量的地址。



#### 3、Go 允许多个返回值吗？

允许

```go
func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("A", "B")
   fmt.Println(a, b) // B A
}
```



#### 4、Go 有异常类型吗？

Go 没有异常类型，只有错误类型（Error），通常使用返回值来表示异常状态。

```go
f, err := os.Open("test.txt")
if err != nil {
    log.Fatal(err)
}
```



#### 5、什么是协程（Goroutine）











