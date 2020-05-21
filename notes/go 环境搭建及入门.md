## Go 环境搭建及入门



### 1. Windows下Go环境搭建 - 安装和配置SDK

#### 1.1 下载SDK

1）Go的官网为：golang.org，需要梯子。

2）Go语言中文网：studygolang.com，

3）选择对应版本：根据电脑操作系统和位数选择

![image-20200520095140432](http://wesub.ifree258.top/image-20200520095140432.png)



#### 1.2 Windows下安装SDK

1）注意：安装路径不要有中文或特殊符号如空格等

2）SDK建议目录：Windows我安装在D:\programs

3）zip解压就行，解压后的go目录结构

![image-20200520095634114](http://wesub.ifree258.top/image-20200520095634114.png)

4）测试go是否安装成功：DOS窗口输入go version

![image-20200520100209218](http://wesub.ifree258.top/image-20200520100209218.png)

显示‘go’不是内部或外部命令。。。意料之中，需要配置环境变量



#### 1.3 Windows下配置环境变量

Go开发中需要配置的环境变量：

| 环境变量 | 说明                                                  |
| -------- | ----------------------------------------------------- |
| GOROOT   | 指定SDK的安装路径 D:\programs\go1.12.windows-amd64\go |
| Path     | 添加SDK的/bin目录                                     |
| GOPATH   | 工作目录，将来我们的go项目的工作路径 D:\projects      |

![image-20200520100805129](http://wesub.ifree258.top/image-20200520100805129.png)



再来测试一下，注意：需要重新打开dos终端，环境变量才会生效！

![image-20200520100937465](http://wesub.ifree258.top/image-20200520100937465.png)



Linux 和 Mac 大同小异，时间原因在此就不介绍了。



### 2. Go快速开发入门



#### 2.1 开发工具

1）VSCode 精简方便 https://code.visualstudio.com/download

2）Goland JetBrains产品

这里先使用VSCode方便入门。工具的下载都很傻瓜式就不过介绍了。



#### 2.2 开发需求

开发一个hello.go程序，输出“Hello，Go”



#### 2.3 开发步骤

1）先看下开发程序、项目时，go的目录结构如何组织

![image-20200520103650022](http://wesub.ifree258.top/image-20200520103650022.png)

2）main目录下新建一个hello.go 文件，代码如下：

![image-20200520103827374](http://wesub.ifree258.top/image-20200520103827374.png)

说明：

- go文件的后缀是 .go

- package main

  表示该hello.go 文件所在的包是main，**在go中，每个文件都必须归属于一个包！**

- import "fmt"

  表示引入一个包，包名fmt，引入该包后就可以使用该包中的函数，比如fmt.Println

- func main(){

  }

  func 是一个关键字，表示一个函数

  main 是函数名，是一个主函数，，即我们程序的入口。

3）通过 go build 命令对该 go 文件进行编译，生成 .exe 文件

![image-20200520105537881](http://wesub.ifree258.top/image-20200520105537881.png)

4）运行hello.exe 即可

![image-20200520105614209](http://wesub.ifree258.top/image-20200520105614209.png)

5）注意：使用 go run 命令可以直接运行go程序（类似执行一个脚本文件的形式）

![image-20200520105857474](http://wesub.ifree258.top/image-20200520105857474.png)



#### 2.4 go执行流程分析

![image-20200520110007110](http://wesub.ifree258.top/image-20200520110007110.png)



#### 2.5 go开发注意事项

1. Go源文件以 “go” 为扩展名
2. Go程序的入口是 main() 函数
3. Go语言严格区分大小写
4. Go方法由一条条语句构成，**每个语句后不需要分号**（Go编译器会在每行后自动加分号），这体现了Go的简洁性
5. Go编译器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一行，否则报错

![image-20200521151001203](http://wesub.ifree258.top/image-20200521151001203.png)

6. Go语言定义的变量或者导入的包如果没有使用到，代码不能编译通过

![image-20200521151102031](http://wesub.ifree258.top/image-20200521151102031.png)

7. 编码风格：

```go
func main(){
    ...
}
上面这种写法正确
-------------
下面这种写法是错误的!GO编译不通过
func main()
{
    ...
}
```









































