# WebSocket

> // Написал пока вкратце, TODO: дописать

- WebSocket — протокол связи поверх TCP-соединения, предназначенный для обмена сообщениями между браузером и веб-сервером в режиме реального времени.
В настоящее время в W3C осуществляется стандартизация API Web Sockets.

### Открытие канала webSocket
Для установления соединения WebSocket клиент и сервер используют протокол, похожий на HTTP. Клиент формирует особый HTTP-запрос,
на который сервер отвечает определенным образом.

Протокол Web Socket определяет две URI-схемы, ws: (нешифрованное соединение) и wss: (шифрованное соединение).
