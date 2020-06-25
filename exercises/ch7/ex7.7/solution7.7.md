## Question

请解释为什么默认值20.0没写°C，而帮助消息中却包含°C

```go
$ go run tempflag.go -h
Usage of /var/folders/g2/_pbmdskj3wscvjk679m80tfr0000gp/T/go-build836957387/b001/exe/tempflag:
  -temp value
    	the temperature (default 20°C)
exit status 2
```

## Answer

Celsius定义了String()方法:

```go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```