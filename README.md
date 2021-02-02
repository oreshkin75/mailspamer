# Руководство по использованию mailSpamer
## Требования
* PostgreSQL
## Библиотеки
* github.com/lib/pq
* gopkg.in/yaml.v3
* gopkg.in/gomail.v2
## Конфигурационный файл
Конфигурационный файл указывается как первый аргумент командной строки.
Файл должен быть в формате .yaml.
```
serverPort: 888
externalServer: http://externalserver.com
mailconf:
  mail: example@mail.com
  password: examplepass
  smtp: smtp.test.com
database:
  user: postgre
  password: testpass
  dbname: emailDb
  sslmode: false
```
