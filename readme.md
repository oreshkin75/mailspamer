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
serverPort: 
externalServer: 
mailconf:
  mail: 
  password: 
  smtp: 
database:
  user: 
  password: 
  dbname: 
  sslmode: 
```