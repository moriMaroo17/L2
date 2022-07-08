Что выведет программа? Объяснить вывод программы.
```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}

```

Ответ:
```
Будет выведено сообщение 'error'. Так как тип переменной 'err' это интерфейс error, то хоть поле data и будет nil, то в поле itable будет записано значение *customError
```