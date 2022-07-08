Что выведет программа? Объяснить вывод программы.

```go
package main
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
После вывода чисел от 0 до 9,которые main получил из канала, выскочит DeadLock,	потому что for n := range ch ждеь значения из канала до тех пор,
пока он не будет закрыт.
```