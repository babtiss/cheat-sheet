# Основы TCP/IP
*Стек протоколов Transmission Control Protocol/Internet Protocol (протоколы управления передачей/Интернет-протоколы) — представляют сетевую модель, описывающую процесс передачи цифровых данных.
Она названа по двум главным протоколам, по этой модели построена глобальная сеть — интернет.*

Как и OSI, модель имеет деление на уровни, внутри которых действуют определенные протоколы и выполняются собственные функции.

![Image](https://github.com/babtiss/cheat-sheet/blob/master/base/Models.jpg)


### Прикладной уровень
Протоколы прикладного уровня действуют для большинства приложений, они предоставляют услуги пользователю или обмениваются данными с нижними уровнями по уже установленным соединениям.
Здесь для большинства приложений созданы свои протоколы, например:
* HTTP для передачи гипертекста по сети,
* SMTP для передачи почты,
* FTP для передачи файлов,
* протокол назначения IP-адресов DHCP и прочие.

### Транспортный уровень
Протоколы транспортного уровня (Transport layer) могут решать проблему негарантированной доставки сообщений («дошло ли сообщение до адресата?»), а также гарантировать правильную последовательность прихода данных.
В стеке TCP/IP транспортные протоколы определяют, для какого именно приложения предназначены эти данные.

* TCP (протокол управления передачей) — надежный, он обеспечивает передачу информации, проверяя дошла ли она, насколько полным является объем полученной информации и т.д. TCP дает возможность двум хостам производить обмен пакетами через установку соединения. Он предоставляет услугу для приложений, повторно запрашивает потерянную информацию, устраняет дублирующие пакеты, регулируя загруженность сети. TCP гарантирует получение и сборку информации у адресата в правильном порядке.

* UDP (протокол пользовательских датаграмм) — ненадежный, он занимается передачей автономных датаграмм. UDP не гарантирует, что всех датаграммы дойдут до получателя. Датаграммы уже содержат всю необходимую информацию, чтобы дойти до получателя, но они все равно могут быть потеряны или доставлены в порядке отличном от порядка при отправлении.

### Межсетевой уровень (internet layer)
Каждая индивидуальная сеть называется локальной, глобальная сеть интернет позволяет объединить все локальные сети. За объединение локальных сетей в глобальную отвечает сетевой уровень. Он регламентирует передачу информации по множеству локальных сетей, благодаря чему открывается возможность взаимодействия разных сетей.

Межсетевое взаимодействие — это основной принцип построения интернета. Локальные сети по всему миру объединены в глобальную, а передачу данных между этими сетями осуществляют магистральные и пограничные маршрутизаторы.

### Канальный уровень
Канальный уровень (англ. Link layer) описывает способ кодирования данных для передачи пакета данных на физическом уровне (то есть специальные последовательности бит, определяющих начало и конец пакета данных, а также обеспечивающие помехоустойчивость). Ethernet, например, в полях заголовка пакета содержит указание того, какой машине или машинам в сети предназначен этот пакет.

Примеры протоколов канального уровня — Ethernet, IEEE 802.11 WLAN, SLIP, Token Ring, ATM и MPLS.