# URL shortener (Сервис сокращения размеров ссылки)
<h2>Запуск приложения осуществляется из директории cmd/app:</h2>

#### go run main.go path_to_conf

*параметром командной строки необходимо указать абсолютный путь к расположению файла конфигурации
Например: **path_to_conf = /home/tatun/go/src/shortener/pkg/app/config/config.toml**

## :briefcase: В качестве хранилища используется **BoltDB**

# Возможности сервиса
## 1. Добавление ссылки в хранилище по url: http://localhost:8080/a/?url=длинная_ссылка

В результате осуществляется добавление ссылки в хранилище, а также генерируется последовательность из 8 символов, которые возвращаются в теле ответа в следующем формате **Short URL = /^[a-z0-9]{8}$/**

## 2. Переход по сокращенной ссылки по url: http://localhost:8080/s/короткая_ссылка

В результате данной операции происходит редирект на исходную (длинную) ссылку
