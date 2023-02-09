# music_player
go music player service 
http use gin and store use sqlite by gorm
request data by gout
## [how to run]
```
go run main.go actions.go
```
then will works on :8080

## [build]
```
go build 
```
## [docker]
```
docker built -t xxx .
```
Also you can find it in Dockerfile
### [docker run]
```
docker run -i -t --rm -p 8080:8080 xxx
```