## Условие
Схема БД состоит из четырех таблиц:

1) `Product(maker, model, type)`
2) `PC(code, model, speed, ram, hd, cd, price)`
3) `Laptop(code, model, speed, ram, hd, price, screen)`
4) `Printer(code, model, color, type, price)`

Таблица Product представляет производителя (maker), номер модели (model) и тип ('PC' - ПК, 'Laptop' - ПК-блокнот или 'Printer' - принтер). 
Предполагается, что номера моделей в таблице Product уникальны для всех производителей и типов продуктов. В таблице PC для каждого ПК, однозначно определяемого уникальным кодом – code, указаны модель – model (внешний ключ к таблице Product), скорость - speed (процессора в мегагерцах), объем памяти - ram (в мегабайтах), размер диска - hd (в гигабайтах), скорость считывающего устройства - cd (например, '4x') и цена - price (в долларах). Таблица Laptop аналогична таблице РС за исключением того, что вместо скорости CD содержит размер экрана -screen (в дюймах). В таблице Printer для каждой модели принтера указывается, является ли он цветным - color ('y', если цветной), тип принтера - type (лазерный – 'Laser', струйный – 'Jet' или матричный – 'Matrix') и цена - price.

## TASK 1
Найдите номер модели, скорость и размер жесткого диска для всех ПК стоимостью менее 500 дол. Вывести: model, speed и hd

Решение:
```sql
SELECT model, speed, hd FROM PC
WHERE price < 500
```

## TASK 2
Найдите производителей принтеров. Вывести: maker

Неверное решение:
```sql
SELECT DISTINCT maker FROM Product
JOIN Printer ON Product.model=Printer.model
```
- такое решение может давать верный результат, но само по себе оно не верно. Если изучить таблицы, то решить можно проще.

Решение
```sql
SELECT DISTINCT maker FROM Product
WHERE type = 'Printer'
```

## TASK 3
Найдите номер модели, объем памяти и размеры экранов ПК-блокнотов, цена которых превышает 1000 дол.
```sql
SELECT model, ram, screen FROM Laptop
WHERE price > 1000
```

## TASK 4
Найдите все записи таблицы Printer для цветных принтеров.
```sql
SELECT * FROM Printer
WHERE color = 'y'
```

## TASK 5
Найдите номер модели, скорость и размер жесткого диска ПК, имеющих 12x или 24x CD и цену менее 600 дол.
```sql
SELECT model, speed, hd FROM PC
WHERE price < 600 and (cd = '12x' or cd = '24x')
```

