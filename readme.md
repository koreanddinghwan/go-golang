```sh
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /albums                   --> main.getAlbums (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on localhost:10020
[GIN] 2023/10/29 - 18:52:52 | 200 |       157.5µs |       127.0.0.1 | GET      "/albums"
[GIN] 2023/10/29 - 18:52:54 | 200 |      92.708µs |       127.0.0.1 | GET      "/albums"
[GIN] 2023/10/29 - 18:52:55 | 200 |      47.083µs |       127.0.0.1 | GET      "/albums"
[GIN] 2023/10/29 - 18:52:55 | 200 |      40.375µs |       127.0.0.1 | GET      "/albums"
```

- 캐싱되는걸 직관적으로 확인이 가능함.
- gin/echo 모두 node의 express와 같은 느낌. 
nestjs처럼 아키텍처 문제를 전적으로 해결하는 프레임워크가 아닌, 최소한의 웹서버역할만을 하며  메모리/CPU점유가 보다 낮음.
