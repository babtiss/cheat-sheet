# База С++

## INDEX:
- [Типы данных](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#%D1%82%D0%B8%D0%BF%D1%8B-%D0%B4%D0%B0%D0%BD%D0%BD%D1%8B%D1%85)
    - [Логический тип](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#логический-тип)
    - [Символьный тип](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#символьный-тип)
    - [Числовой тип](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#числовой-тип)
    - [Void](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#void)
    - [Спецификатор auto](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#спецификатор-auto)
    - [nan и inf](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#nan-и-inf)
- [Указатели](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#указатели)
  - [Динамические переменные](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#динамические-переменные)
  - [Освобождения памяти](https://github.com/babtiss/cheat-sheet/tree/master/c++/base/types#освобождение-памяти)

## Типы данных
В языке С++ все переменные имеют определенный тип данных. 
Тип данных присваивается переменной при ее объявлении или инициализации. 

### Логический тип
- `bool`: Может принимать одну из двух значений true (истина) и false (ложь). Размер занимаемой памяти для этого типа точно не определен.
```c++
bool a = true
```
### Символьный тип
- `char`: представляет один символ в кодировке ASCII. Занимает в памяти 1 байт (8 бит). Может хранить любое значение из диапазона от -128 до 127, либо от 0 до 255

- `signed char`: представляет один символ. Занимает в памяти 1 байт (8 бит). Может хранить любой значение из диапазона от -128 до 127

- `unsigned char`: представляет один символ. Занимает в памяти 1 байт (8 бит). Может хранить любой значение из диапазона от 0 до 255

- `wchar_t`: представляет расширенный символ. На Windows занимает в памяти 2 байта (16 бит), на Linux - 4 байта (32 бита). Может хранить любой значение из диапазона от 0 до 65 535 (при 2 байтах), либо от 0 до 4 294 967 295 (для 4 байт)

- `char16_t`: представляет один символ в кодировке Unicode. Занимает в памяти 2 байта (16 бит). Может хранить любой значение из диапазона от 0 до 65 535

- `char32_t`: представляет один символ в кодировке Unicode. Занимает в памяти 4 байта (32 бита). Может хранить любой значение из диапазона от 0 до 4 294 967 295

Стоит учитывать, что для вывода на консоль символов `wchar_t` следует использовать не `std::cout`, а` поток std::wcout:`

При этом поток `std::wcout` может работать как с `char`, так и с `wchar_t`. А поток `std::cout` для переменной `wchar_t` вместо символа будет выводить его числовой код.
```c++
int main()
{
    char a = 'H';
    wchar_t b = 'e';
    std::wcout << a << b << '\n';
    return 0;
}
```

### Числовой тип
- `short`: представляет целое число в диапазоне от –32768 до 32767. Занимает в памяти 2 байта (16 бит).
Данный тип также имеет синонимы `short int`, `signed short int`, `signed short`.

- `unsigned short`: представляет целое число в диапазоне от 0 до 65535. Занимает в памяти 2 байта (16 бит).
Данный тип также имеет синоним `unsigned short int`.

- `int`: представляет целое число. В зависимости от архитектуры процессора может занимать 2 байта (16 бит) или 4 байта (32 бита). Диапазон предельных значений соответственно также может варьироваться от –32768 до 32767 (при 2 байтах) или от −2 147 483 648 до 2 147 483 647 (при 4 байтах). Но в любом случае размер должен быть больше или равен размеру типа short и меньше или равен размеру типа long
Данный тип имеет синонимы `signed int` и `signed`.

- `unsigned int`: представляет положительное целое число. В зависимости от архитектуры процессора может занимать 2 байта (16 бит) или 4 байта (32 бита), и из-за этого диапазон предельных значений может меняться: от 0 до 65535 (для 2 байт), либо от 0 до 4 294 967 295 (для 4 байт).
В качестве синонима этого типа может использоваться unsigned

- `long`: представляет целое число в диапазоне от −2 147 483 648 до 2 147 483 647. Занимает в памяти 4 байта (32 бита).
У данного типа также есть синонимы `long int`, `signed long int` и `signed long`

- `unsigned long`: представляет целое число в диапазоне от 0 до 4 294 967 295. Занимает в памяти 4 байта (32 бита).
Имеет синоним `unsigned long int`.

- `long long`: представляет целое число в диапазоне от −9 223 372 036 854 775 808 до +9 223 372 036 854 775 807. Занимает в памяти, как правило, 8 байт (64 бита).
Имеет синонимы `long long int`, `signed long long int` и `signed long long`.

- `unsigned long long`: представляет целое число в диапазоне от 0 до 18 446 744 073 709 551 615. Занимает в памяти, как правило, 8 байт (64 бита).
Имеет синоним `unsigned long long int`.

- `float`: представляет вещественное число ординарной точности с плавающей точкой в диапазоне +/- 3.4E-38 до 3.4E+38. В памяти занимает 4 байта (32 бита)

- `double`: представляет вещественное число двойной точности с плавающей точкой в диапазоне +/- 1.7E-308 до 1.7E+308. В памяти занимает 8 байт (64 бита)

- `long double`: представляет вещественное число двойной точности с плавающей точкой не менее 8 байт (64 бит). В зависимости от размера занимаемой памяти может отличаться диапазон допустимых значений.

```c++
unsigned short b= 10;
int c = -30;
float a = -10.45;
double b = 0.00105;
long double c = 30.890045;
```

В С++ есть оператор `sizeof()`, который возвращает размер памяти в байтах, которую занимает переменная:
```c++
int main()
{
    long double number = 2;
    std::cout << "sizeof(number) =" << sizeof(number);
    return 0;
}
```
### Void
- `void`: тип без значения

### Спецификатор auto
Иногда бывает трудно определить тип выражения. И согласно последним стандартам можно предоставить компилятору самому выводить тип объекта. И для этого применяется спецификатор auto. При этом если мы определяем переменную со спецификатором auto, эта переменная должна быть обязательно инициализирована каким-либо значением:
```c++
auto number = 5;
```

### nan и inf

Есть две специальные категории чисел типа с плавающей запятой:

`inf` (или «бесконечность», от англ «infinity»), которая может быть либо положительной, либо отрицательной.

`nan` (или «не число», от англ «not a number»). Их есть несколько видов.

```c++
double zero = 0.0;
double posinf = 5.0 / zero; // положительная бесконечность
double neginf = -5.0 / zero; // отрицательная бесконечность
double nan = zero / zero; // не число (математически некорректно)
```


## Указатели

При выполнении любой программы, все необходимые для ее работы данные должны быть алоцированы в оперативной памяти компьютера.
Для обращения к переменным, находящимся в памяти, используются специальные адреса, 
которые записываются в шестнадцатеричном виде, например 0x100 или 0x200.

Если переменных в памяти потребуется слишком большое количество, которое не сможет вместить в себя сама аппаратная часть,
произойдет перегрузка системы или её зависание.

Поэтому нужно удалять неиспользуемые переменные.

### Динамические переменные
Выделение памяти осуществляется с помощью оператора new и имеет вид: `тип_данных *имя_указателя = new тип_данных;`.

Инициализация значения, находящегося по адресу указателя выполняется схожим образом,
только в конце ставятся круглые скобки с нужным значением:
`тип данных *имя_указателя = new тип_данных(значение);`.
```c++
int main()
{
    int *a = new int; // Объявление указателя для переменной типа int
    int *b = new int(5); // Инициализация указателя

    *a = 10;
    *b = *a + *b;

    cout << "b is " << *b << endl;

    delete b;
    delete a;

    return 0;
}
```
После удачного выполнения такой операции, в оперативной памяти компьютера происходит выделение диапазона ячеек,
необходимого для хранения переменной типа `int`.

### Освобождения памяти
При использовании оператора `delete` для указателя, знак `*` не используется.
```c++
int main()
{
    // Выделение памяти
    int *a = new int;
    int *b = new int;
    float *c = new float;

    // ... Любые действия программы

    // Освобождение выделенной памяти
    delete c;
    delete b;
    delete a;

    return 0;
}
```