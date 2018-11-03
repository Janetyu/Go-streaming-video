## 常用命令

#### build 

- 最常用的go command之一，编译go文件
- 跨平台编译： env GOOS=linux GOARCH=amd64 go build

#### install

- 也是编译，与build最大的区别是编译后会将输出文件打包成库放在pkg下
- 常用与本地打包编译的命令：go install

#### get

- 用于获取go的第三方包，通常会默认从git repo上pull最新的版本
- 常用命令如：go get -u github.com/go-sql-driver/mysql（从github上获取mysql的driver病安装到本地）， -u 指始终获取最新版本的包

#### fmt

- 类似与C中的 lint ，统一代码风格和排版
- 常用命令如： go fmt

#### test

- 运行当前包目录下的tests
- 常用命令如：go test 或 go test -v 等
- 文件名为：xxx_test.go
- xxx 的部分一般为 xxx_test.go 所要测试的代码文件名，注：Go并没有特别要求xxx的部分必须是要测试的文件名。

##### test的写法

- 每一个 test 文件须 `import` 一个 `testing` 

- `test` 文件下的每一个 `test case` 均必须以 `Test` 开头并且符合 `TestXxxx` 形式，否则 `go test` 会直接跳过测试不执行

- `test case` 的入参为 `t *testing.T` 或 `b *testing.B`

- `t.Errorf` 为打印错误信息，并且当前 `test case` 会被跳过

- `t.SkipNow()`为跳过当前的`test`，并且直接按PASS处理继续下一个`test`

-  Go 的 test 不会保证多个 TestXxx 是顺序执行，但通常会按顺序执行

- 使用 t.Run 来执行 `subtests` 可以做到控制 test 输出以及 test 的顺序

  - ```go
    func TestPrint(t *testing.T) {
        t.Run("a1", func(t *testing.T) {fmt.Println("a1")})
        t.Run("a2", func(t *testing.T) {fmt.Println("a2")})
        t.Run("a3", func(t *testing.T) {fmt.Println("a3")})
    }
    ```

- 使用 TestMain 作为初始化test，并且使用 m.Run() 来调用其他 tests 可以完成一些需要初始化操作的 testing，比如数据库连接， 文件打开， REST服务登录等

  - ```go
    func TestMain(m *testing.M) {
    	fmt.Println("test main first")
    	m.Run()
    }
    ```

  - 如果没有在TestMain中调用m.Run()，则除了TestMain以外的其他tests都不会被执行

##### Test之benchmark

- benchmark函数一般以Benchmark开头

- benchmark的case一般会跑b.N次，而且每次执行都会如此

- 在执行过程中会根据实际case的执行时间是否稳定，会增加b.N的次数以达到稳态

  - ```go
    func BenchmarkAll(b *testing.B) {
    	for n := 0; n < b.N; n++ {
    		Print1to20()
    	}
    }
    ```

    `go test -bench=.` 输入命令进行测试，此命令只会跑带Bench开头的测试