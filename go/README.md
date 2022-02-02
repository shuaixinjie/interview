## go语法

### 数据类型
```
布尔型 bool
整型 int8、16、32、64 uint8、16、32、64
浮点型 float32、float64、complex32、complex64
其他数字类型 byte(uint8)、rune(int32)、uint\int、uintptr
string
其他 pointer、array、struct、channel、func、slice、interface、map
```

### 变量和常量
```
var identifier type
var identifier = value
identifier := value (局部变量可以这么做) 

const identifier [type] = value
iota 常量计数
```

**写到这儿我发现go的语法真的没有什么好写的，特性语法糖都很少，就不写了，直接上考点**

### 常问问题

#### new和make的区别
```
new 的作用是初始化一个指向类型的指针(*Type )，使用new函数来分配空间。传递给new 函数的是一个类型，不是一个值。返回值是 指向这个新分配的零值的指针。
make 的作用是为 slice，map 或 chan 初始化并返回引用(Type)。 第一个参数是一个类型，第二个参数是长度。
并且，map和channel是必须用make初始化之后才可以使用的，而slice即使为nil也可以使用内建的append函数
```

#### go的switch特别之处
```
单个case中，可以出现多个结果选项。
只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case
```

#### go的for特别之处
```
go的break可以选择终止哪一个for，这和goto的功能如出一辙
```

#### go中的引用类型
```
slice、map、channel、interface
引用类型即本身就是传递的指针，所以需要注意如果在结构体中，是返回该拷贝，还是该引用
```

#### channel特性
```
给一个 nil channel 发送数据，造成永远阻塞
从一个 nil channel 接收数据，造成永远阻塞
给一个已经关闭的 channel 发送数据，引起 panic
从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
无缓冲的channel是同步的，而有缓冲的channel是非同步的
```

#### go的panic场景
```
nil pointer
index out of range
division by zero
the panic was called
```



