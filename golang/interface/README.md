# Интерфейс
Интерфейс в GO - это определенный _контракт_ между абстракцией и реализацией поведения у объектов.
В Go объект реализует интерфейс, когда он реализует все его методы.
> Да-да, утиная типизация би лайк

### Определение простого интерфейса
```go
type ExampleInterface interface {
    ExampleFunc()
}
```

### Пример когда объект реализует интерфейс
```go
type ExampleInterface interface {
    ExampleFunc()
}
type A struct {
    value int
}
func (a A) ExampleFunc() {
    return
}
func UseInterface(example ExampleInterface)  {
    return
}
func main()  {
    a := A{value: 1}
    a.ExampleFunc()
    UseInterface(a)
}
// Метод UseInterface() - ожидает на вход объект, который реализует данный интерфейс ExampleInterface
// Как раз объект A реализует такой интерфейс, т.к. имеет описанный метод ExampleFunc()
```

### Дженерики
> я бы написал тут что-то, но я не юзал их

### Пустой интерфейс в Golang
Интерфейс - это два компонента: это набор методов и тип.

Тип `interface{}` - это интерфейс, не имеющий методов.
> Пустому интерфейсу удовлетворяет любой тип.
