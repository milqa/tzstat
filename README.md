* Для старта сервиса `docker-compose up`
* Для старта loader-a `go run ./cmd/load/main.go -job get -rps 50 -clients 10`
  * job: может быть get или post
  * rps: количество запросов в секунду для одного клиента
  * clients: количество клиентов
* grafana: http://localhost:3000/
  * дашборд tzstat: http://localhost:3000/d/mgHjIWc4k/tzstat?orgId=1&refresh=5s

 
 
Какие запросы принимает:

POST:
  
```
curl -X POST http://localhost:8080/api/stat/ -d '{"datetime": "2022-12-12T21:00:00Z", "value": 10}'
```
  
GET:

```
curl -X GET http://localhost:8080/api/stat/
```
или
```
curl -X GET http://localhost:8080/api/stat?date_to=2022-12-12T21:00:00Z
```

Что еще можно добавить:
* Много чего
* нормальный грейсфул шатдаун
* децимал значения
* тесты апи
* swagger/grpc
* переделать main
* добавить хоть какую-то базу
* логов добавить
* и еще кучу всего