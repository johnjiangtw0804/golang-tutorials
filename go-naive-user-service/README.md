<p>

# Go-naive-user-service
這是一個簡單的go user service api. 在 "local" 上專案目錄下可以使用以下步驟去啟動server. <br>
1. go mod init Go-naive-user-service<br>
2. go build<br>
3. ./Go-naive-user-service<br>

## Bonus
以下操作可以使用 Docker 將應用程式容器化 <br>
1. docker image build -t go-naive-v01 .
2. docker image ls (optional, just to make sure the image has been created) <br>
3. docker run -dp 8080:8080 -it go-naive-v01 <br>
使用完成後用以下指令Stop container<br>
   docker stop \<containerId\>

(optional) <br>
docker ps <to look for process status> <br>
docker docker exec -it \<containerId\> /bin/bash (run bash on container to interact with the container incase you need anything)

## Client操作方法
### Get user

Request example
```bash
curl --request GET \
  --url http://localhost:8080/user/<user_id>
```

Success example

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 26 Jul 2021 15:39:09 GMT
Content-Length: xxx
Connection: close

{
  "user_id": "foobar",
	"age": 30
}
```

### Create user

Request example

```bash
curl --request POST \
  --url http://localhost:8080/user \
  --header 'content-type: application/json' \
  --data '{"user_id": "foobar","age": 30}'
```

Success example

```bash
HTTP/1.1 204 No Content
```

### Delete user

Request example

```bash
curl --request DELETE \
  --url http://localhost:8080/user/<user_id>
```
Success example

```bash
HTTP/1.1 204 No Content
```
<p>
