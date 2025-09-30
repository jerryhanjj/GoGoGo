# fmt.Println 调用 String() 方法的源码分析

## 1. 调用链路概览

```
fmt.Println(p) 
    ↓
fmt.Fprintln(os.Stdout, p)
    ↓  
fmt.Fprint(w, p)
    ↓
pp.doPrint(args)
    ↓
pp.printArg(arg)
    ↓
pp.handleMethods(verb)
    ↓
检查 fmt.Stringer 接口
    ↓
调用 p.String()
```

## 2. 关键源码分析

### 2.1 fmt.Println 的实现

在 Go 源码 `src/fmt/print.go` 中：

```go
// Println formats using the default formats for its operands and writes to standard output.
func Println(a ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintln(a)
    n, err = w.Write(p.buf)
    p.free()
    return
}
```

### 2.2 核心格式化逻辑

```go
// doPrintln is like doPrint but always adds a space between arguments and a newline after.
func (p *pp) doPrintln(a []interface{}) {
    for argNum, arg := range a {
        if argNum > 0 {
            p.buf.writeByte(' ')
        }
        p.printArg(arg, 'v')  // 关键：这里调用 printArg
    }
    p.buf.writeByte('\n')
}
```

### 2.3 参数处理函数 printArg

```go
// printArg formats the argument and writes it to the buffer.
func (p *pp) printArg(arg interface{}, verb rune) {
    p.arg = arg
    p.value = reflect.Value{}

    if arg == nil {
        switch verb {
        case 'T', 'v':
            p.fmt.padString(nilAngleString)
        default:
            p.badVerb(verb)
        }
        return
    }

    // 关键部分：处理特殊方法（包括String方法）
    if p.handleMethods(verb) {
        return
    }
    
    // 如果没有特殊方法，使用默认格式化
    p.printValue(reflect.ValueOf(arg), verb, 0)
}
```

### 2.4 最关键的 handleMethods 函数

```go
// handleMethods 处理格式化方法，包括 String() 方法
func (p *pp) handleMethods(verb rune) (handled bool) {
    if p.erroring {
        return
    }
    
    // 避免递归调用
    if verb == 'w' {
        // ... 处理 error wrapping
    }
    
    // 处理 error 接口
    if verb != 'T' {
        if formatter, ok := p.arg.(error); ok && verb == 'v' || verb == 's' || verb == 'q' {
            // ... 处理 error
        }
    }
    
    // 🔥 关键部分：检查 fmt.Stringer 接口
    if verb != 'T' && verb != 'p' {
        if stringer, ok := p.arg.(Stringer); ok {
            p.fmt.fmtS(stringer.String())  // 调用 String() 方法！
            return true
        }
    }
    
    // 处理 GoStringer 接口（用于 %#v）
    if verb == 'v' {
        if goStringer, ok := p.arg.(GoStringer); ok {
            p.fmt.fmtS(goStringer.GoString())
            return true
        }
    }
    
    return false
}
```

### 2.5 fmt.Stringer 接口定义

在 `src/fmt/print.go` 中定义：

```go
// Stringer is implemented by any value that has a String method,
// which defines the "native" format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
type Stringer interface {
    String() string
}
```

## 3. 类型断言的工作原理

当执行 `if stringer, ok := p.arg.(Stringer); ok` 时：

1. **运行时类型检查**：Go运行时检查 `p.arg` 的具体类型是否实现了 `Stringer` 接口
2. **方法集匹配**：检查该类型的方法集是否包含 `String() string` 方法
3. **接口转换**：如果匹配成功，将 `p.arg` 转换为 `Stringer` 接口类型
4. **方法调用**：通过接口调用 `String()` 方法

## 4. 为什么会自动调用

总结一下为什么 `fmt.Println(p)` 会自动调用 `p.String()`：

1. **设计哲学**：Go的fmt包设计时考虑了类型的"自然表示"
2. **接口约定**：如果类型实现了 `String()` 方法，说明它知道如何表示自己
3. **优先级机制**：fmt包优先使用类型自定义的字符串表示
4. **自动发现**：通过类型断言自动发现并调用合适的方法

## 5. 与你的代码的对应关系

在你的 `stringer.go` 中：

```go
type stPerson struct {
    Name string
    Age  int
}

func (p stPerson) String() string {
    return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

func StringerDemo() {
    p := stPerson{Name: "Alice", Age: 30}
    fmt.Println(p) // 这里会触发上面的整个调用链
}
```

执行流程：
1. `fmt.Println(p)` 调用
2. 最终到达 `handleMethods()` 函数
3. 执行类型断言：`if stringer, ok := p.(Stringer); ok`
4. 因为 `stPerson` 实现了 `String()` 方法，断言成功
5. 调用 `stringer.String()`，即你定义的 `String()` 方法
6. 返回格式化后的字符串并输出

这就是为什么你的 `String()` 方法会被自动调用的完整原理！