# Потокобезопасность
Потоковая безопасность — это концепция программирования, применимая к многопоточным программам.
Код потокобезопасен, если он функционирует исправно при использовании его из нескольких потоков одновременно.
В частности, он должен обеспечивать правильный доступ нескольких потоков к разделяемым данным.

## INDEX:
- [Проблема DataRace](https://github.com/babtiss/cheat-sheet/tree/master/golang/multithreading/thread%20safety#%D0%BF%D1%80%D0%B8%D0%BC%D0%B5%D1%80-%D0%BF%D1%80%D0%BE%D0%B1%D0%BB%D0%B5%D0%BC%D1%8B)
- [Пакет sync](https://github.com/babtiss/cheat-sheet/tree/master/golang/multithreading/thread%20safety#%D0%BF%D1%80%D0%B8%D0%BC%D0%B8%D1%82%D0%B8%D0%B2%D1%8B-sync)
    - [type Mutex](https://github.com/babtiss/cheat-sheet/tree/master/golang/multithreading/thread%20safety#syncmutex--mutually-exclusive-lock)
    - [type Once](https://github.com/babtiss/cheat-sheet/tree/master/golang/multithreading/thread%20safety#synconce)
    - [type WaitGroup](https://github.com/babtiss/cheat-sheet/tree/master/golang/multithreading/thread%20safety#syncwaitgroup)
    - [type Map](https://github.com/babtiss/cheat-sheet/tree/master/golang/multithreading/thread%20safety#syncmap)
- [Пакет atomic]()

## Пример проблемы DataRace
Гонка горутин иногда приводит к изменениям состояния любого значения, хранящегося в адресе памяти без какого-либо соблюдения порядка.

```go
package main

import (
    "fmt"
    "time"
)

type Counter struct {
    Count int
}

func updateC(c *Counter, cur int) {
    c.Count += cur
    fmt.Printf("New Count is %v\n", c.Count)
}

func main() {
    totalCur := 50
    c := Counter{1000000}
    for i := 0; i <= totalCur; i++ {
        go updateC(&c, int(i))
    }
    time.Sleep(time.Second * 1)  // this is just so that we don't need channels
}
```

output:
```go
New Count is 1001199
New Count is 1001030
New Count is 1001206
...
New Count is 1001176
New Count is 1000957
New Count is 1001005
```

Это происходит из-за того, что операция `c.Count += cur` состоит из нескольких операций, которые выполняются не атомарно:
```go
*addr += delta
return *addr
```
Атомарность должен поддерживать сам разработчик иначе может возникать такая проблема.

Научимся решать эту и многие другие проблемы.

## Примитивы `sync`

### `sync.Mutex` — Mutually Exclusive Lock.
Мьютекс — это блокировка взаимного исключения. Нулевое значение для мьютекса — это разблокированный мьютекс.

```go
type Mutex struct {
    state int32
    sema  uint32
}
```
Mutex используется для того, чтобы заблокировать остальные потоки и поставить их на ожидание в очереди,
когда одна горутина обращается к значению внутри адреса памяти.
Это гарантирует, что не будет случайного доступа и изменения значений.

```go
func (m *Mutex) Lock()
// Lock: только одна процедура go читает/записывает одновременно, получая блокировку.

func (m *Mutex) TryLock() bool
// TryLock: пытается заблокировать m и сообщает, удалось ли это.

func (m *Mutex) Unlock()
// Unlock: разблокирует m.
// Получим ошибку времени выполнения, если m не заблокирован для чтения при входе в RUnlock.
```

Пример использования:
```go
func updateC(c *Counter, cur int) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.Count += cur
    fmt.Printf("New Count is %v\n", c.Count)
}
```

Обязательно нужно использовать размьючивание `defer c.mu.Unlock()`. Причина этого в том, что нам нужно избежать взаимоблокировки.

Взаимоблокировки — это уязвимые места мьютексов, которых следует избегать.
Иначе блокировка будет реализована на неопределенный срок и что все остальные горутины вообще не смогут получить к ней доступ.

### `sync.RWMutex` - Reader and Writer Mutually Exclusive Lock.
RWМьютекс — это блокировка взаимного исключения чтения/записи.
Блокировка может удерживаться произвольным числом читателей или одним писателем.

```go
func (rw *RWMutex) Lock()

func (rw *RWMutex) Unlock()

func (rw *RWMutex) RLock()
// несколько подпрограмм могут читать одновременно, но не записывать, получая блокировку.

func (rw *RWMutex) RUnlock()
// RUnlock отменяет одиночный вызов RLock; это не влияет на других одновременных читателей.
// Получим ошибку времени выполнения, если rw не заблокирован для чтения при входе в RUnlock.

func (rw *RWMutex) RLocker() Locker
// RLocker возвращает интерфейс Locker, который реализует методы Lock и Unlock, вызывая rw.RLock и rw.RUnlock.

func (rw * RWMutex ) TryLock() bool

func (rw * RWMutex ) TryRLock() bool
```

Пример использования:
```go
type SafeDict struct {
    data  map[string]int
    *sync.RWMutex
}

func (d *SafeDict) Get(key string) (int, bool) {
    d.RLock()
    defer d.RUnlock()
    old_value, ok := d.data[key]
    return old_value, ok
}

```

### `sync.Once`
Once - это объект, который будет выполнять ровно одно действие.

```go
func (o *Once) Do(f func())
// Do вызывает функцию f только тогда, когда Do вызывается впервые для данного экземпляра Once.
```

Пример использования:
```go
func main() {
    var once sync.Once
    onceBody := func() {
        fmt.Println("Only once")
    }
    done := make(chan bool)
    for i := 0; i < 10; i++ {
        go func() {
            once.Do(onceBody)
            done <- true
        }()
    }
    for i := 0; i < 10; i++ {
        <-done
    }
}

output:
"Only once"
```


### `sync.WaitGroup`
Группа ожидания - предназначена для создания точки, в которой программа дожидается окончания всех горутин в группе.

Основная горутина вызывает Add, чтобы установить количество ожидаемых горутин.
Затем запускается каждая горутина и по завершении вызывает Done.
В то же время, Wait можно использовать для блокировки, пока не закончатся все горутины.

```go
func (wg *WaitGroup) Add(delta int)
// Add добавляет дельту (может быть отрицательной) к счетчику группы ожидания.
// Если счетчик становится равным нулю: все горутины, заблокированные в ожидании, освобождаются.
// Если счетчик становится отрицательным: получим панику.

func (wg * WaitGroup ) Done()
// Done уменьшает значение счетчика WaitGroup на единицу.

func (wg *WaitGroup) Wait()
// Блокировка горутины, пока счетчик WaitGroup не станет равным нулю.
```

> Вызовы Add должны происходить до оператора, создающего горутину или другое ожидаемое событие.

Пример использования:
```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    funcWithWaitGroup()
}

func funcWithWaitGroup() {
    wg := sync.WaitGroup{}
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println(i)
            wg.Done()
        }(i)
    }
    wg.Wait()
    fmt.Println("exit")
}
```


### `sync.Map`
Карта похожа на `map[interface{}]interface{}` в Go,
но безопасна для одновременного использования несколькими горутинами без дополнительной блокировки или координации.
Загрузки, сохранения и удаления выполняются в амортизированном постоянном времени.

```go
type Map struct {
    mu sync.Mutex

    read atomic.Value // указатель на структуру readOnly
    dirty map[interface{}]*entry
    misses int
}

type readOnly struct {
    m       map[interface{}]*entry
    amended bool
}
```
`read` — указатель на структуру `readOnly`, в данной структуре хранится часть данных `sync.Map`,
используемая для проверки наличия ключа, либо же чтения.
Поэтому при доступе к `read` мьютексы не нужны (параллельное чтение мапы не является проблемой).

`dirty` — мапа, в которой хранится другая часть данных,
используемая для добавления новых элементов.
Поэтому при доступе к `dirty` задействуется мьютекс `mu`.

Таким образом, `sync.Map` имеет две внутренние мапы (`read` и `dirty`)
и благодаря этому пытается избежать использования мьютексов при чтении.
Далее рассмотрим, как изнутри происходят различные операции с `sync.Map`.

Тип карты оптимизирован для двух распространенных случаев использования:
1) когда запись для данного ключа записывается только один раз, но читается много раз, как в кэшах, которые только увеличиваются
2) когда несколько горутин читают, записывают и перезаписывать записи для непересекающихся наборов ключей.
В этих двух случаях использование карты может значительно уменьшить количество конфликтов за блокировку,
по сравнению с картой Go в паре с отдельным `Mutex` или `RWMutex`.

```go
func (m *Map) Delete(key any)
// Удалить значение ключа

func (m *Map) Load(key any) (value any, ok bool)
// Load возвращает значение, хранящееся в карте для ключа, или nil, если значение отсутствует.
// Результат ok указывает, было ли найдено значение в map.

func (m *Map) Store(key, value any)
// Store устанавливает значение для ключа.

func (m *Map) LoadAndDelete(key any) (value any, loaded bool)
// LoadAndDelete удаляет значение ключа, возвращая предыдущее значение, если оно есть.
// Загруженный результат сообщает, присутствовал ли ключ.

func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)
// LoadOrStore возвращает существующее значение для ключа, если оно присутствует.
// В противном случае он сохраняет и возвращает заданное значение.
// Загруженный результат истинен, если значение было загружено, и ложно, если сохранено.

func (m *Map) Range(f func(key, value any) bool)
// Range последовательно вызывает f для каждого ключа и значения, представленных на карте.
// Если f возвращает false, диапазон останавливает итерацию.
```

## Пакет `atomic`
Пакет atomic предоставляет низкоуровневые примитивы атомарной памяти, полезные для реализации алгоритмов синхронизации.

