# Домашнее задание

Есть сервис, который по запросу отдаёт данные рекламного баннера. Одним из компонентов ответа является кликовая ссылка (см. пример ниже).
Запрос к сервису представляет собой JSON объект такого вида:
```json
{
  "id": "123",
  "ua": "Baikal PC",
  "cpc": 0.13145,
  "redirect_uri": "http://example.com"
}
```
Поля запроса являются необязательными, в случае отсутствия они будут заполнены сервисом по своему усмотрению.

Пример такого запроса:
```shell
POST https://demo.kadam.net/go-interview/homework/banner
Content-Type: application/json

{
  "id": "123",
  "ua": "iPhone",
}
```
и ответа:
```json
{
  "image": "https://media.moddb.com/images/groups/1/3/2392/cat.jpg",
  "width": 1600,
  "height": 1200,
  "click_uri": "http://localhost:8080/Y25iXWV3dBADCQkLDHQJKNUEBoO2XkR6ARknHyVfXVYkFhFLCB83DjgJExsgTwUKBEIxAzkCXQ==",
  "debug": {
    "click_data": {
      "id": "123",
      "ua": "iPhone",
      "cpc": 0.6868230728671094,
      "redirect_uri": "https://www.ardanlabs.com/blog/"
    }
  }
}
```

Все прочие поля, кроме `click_uri` нам неинтересны. Тело кликовой ссылки представляет собой данные, последовательно закодированные:
* base64
* encrypt (см алгоритм ниже)
* protobuf marshal пакета Click из click.proto (*)

Шифрование пакета необходимо для защиты от подделки кликовой ссылки и делается с помощью `xor` функции
```go
for i := 0; i < len(raw); i++ {
    raw[i] = raw[i] ^ sec[i%len(sec)]
}
```
с ключом `imSoVerySafe`.

Когда человек нажимает на рекламный баннер - происходит переход по кликовой ссылке.
Нужно написать HTTP сервер их обрабатывающий:
1. Декодировать пакет click из тела кликовой ссылки;
2. Проверить уникальность клика по click.id - можно использовать обычный map или любую СУБД/хранилище;
3. Проверить соответсвие заголовка User-Agent в запросе и click.ua в декодированном пакете;
4. Записать все поля click в таблицу clicks любой реляционной СУБД. Если хотя бы одна из предыдущих проверок не прошла - записать true в колонку isSuspicious;
5. Перенаправить пользователя на click.redirectUri.

Таблица clicks должна содержать теже поля, что описаны в click пакете + колонку isSuspicious.

Материалы для размышления:
1) (*) https://developers.google.com/protocol-buffers/docs/gotutorial;
2) https://en.wikipedia.org/wiki/One-time_pad - в задании версия без random.
