# Сервис сокращения ссылок
 
# Запуск

~~~zsh
# Клонирование репозитория
git clone https://github.com/realPointer/url-shortener && cd url-shortener

# Запуск in-memory
make compose-up

# Запуск PostgreSQL
make compose-up-postgres
~~~

# Swagger

После запуска приложения доступна Swagger-документация по адресу [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

![swagger](https://github.com/realPointer/url-shortener/assets/50529632/855006ce-1018-41cd-b7f8-dbaeae0d2d0a)


# HTTP Запросы

### Сохранение оригинального URL и возвращение сокращённого

~~~zsh
curl -X POST "http://localhost:8080/v1/shortener" \
       -H 'Content-Type: application/json' \
       -d '{"url": "{originalURL}"}'
~~~

Пример ответа:
~~~json
{"shortURL":"fE54KN4v-4"}
~~~
---

### Получение оригинальной ссылки

~~~zsh
curl "http://localhost:8080/v1/shortener/{shortURL}"
~~~

Пример ответа:
~~~json
{"originalURL":"ya.ru"}
~~~

# gRPC

![image](https://github.com/realPointer/url-shortener/assets/50529632/91103f34-30c5-405e-b4e3-43d7f3bf245b)
