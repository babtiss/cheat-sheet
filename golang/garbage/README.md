## Как работает Garbage Collection в Go?
- Garbage Collection - это процесс освобождения места в памяти, которое больше не используется. В документации указано следующее:
```
GC выполняется конкурентно (concurrent), одновременно с потоками мутатора (mutator),
в точном соответствии с типом (этот принцип также известен как чувствительность к типу),
допускается парааллельная выполнение нескольких потоков GC.
Это конкурентная пометка и очистка (mark-sweep), при которой используется барьер записи (write barrier).
При этом в процессе ничего не генерируется и не сжимается.
Освобождение памяти выполняется на основе размера, выделенного для каждой программы Р,
чтобы в общем случае минимизировать фрагментацию и избежать блокировок.
```

## Алгоритм

В основе работы GC Go лежит "трехцветный алгоритм". Официальное название "трехцветный алгоритм пометки и очистки".
Использует барьер памяти. Главный принцип алгоритма трехцветной пометки и очистки состоит в разделении объектов, находящихся в куче, на три набора, в соответствии с "цветом".
Условно разделяются на 3 цвета:
- черные объекты - гарантированно не имеют указателей на белые объекты;
- серые объекты - могут иметь указатели на белые объекты;
- белые объекты - могут иметь указатели на черные объекты; на них могут ссылаться некоторые серые объекты.
Краткий алгоритм:

1. Все объекты сначала белые;
2. Идет перебор "корневых" объектов, помечаются как серые. Корневые - это объекты к которым можно обращаться напрямую, например глобальные переменные, элементы в стеке и т.д.
3. Идет перебор серых объектов, проверяются ссылки на другие объекты и помечаются на черные объекты. Если есть ссылка на белый объект, то белый становится серым.
4. Продолжается до тех пор, пока не будут перебраны все серые объекты.
5. Оставшиеся после перебора белые объекты считаются недостижимыми и занимаемая ими область памяти может быть освобождена.
- Есть еще Мутатор - это приложение, работающее во время сборки мусора. Вызывает функцию барьера записи. Выполняется каждый раз, когда меняется указатель в куче. После изменения указателя объект считается достижимым и помечается как серый.