# go-kakaoskill

- `golang` 용 **카카오 i 오픈빌더 스킬 v2** 라이브러리입니다.

- 마땅한 라이브러리가 보이지 않아서 만들었습니다.

- [카카오 i 오픈빌더 도움말](https://i.kakao.com/docs/skill-build) 을 기반으로 제작하였습니다.

- [**MIT license**](LICENSE) 하에 배포됩니다.

## installation

1. 먼저 [Go](https://golang.org/) (**1.13 이상**) 설치해주신 후, 다음 커맨드를 입력하여 **go-kakaoskill** 을 설치합니다.
```shell
$ go get -v "github.com/RyuaNerin/go-kakaoskill"
```

2. 다음과 같이 import 하여 사용할 수 있습니다.
```go
import skill "github.com/RyuaNerin/go-kakaoskill/v2" // go modules 사용 (GO111MODULE=on or $GOPATH 밖)
import skill "github.com/RyuaNerin/go-kakaoskill" // go modules 사용 안함
```

## Quick start

- `skill.F` `skill.H` 두 함수를 사용해서 `http.HandlerFunc` 로 변환해주세요.

```go
// example.go
package main

import (
	"net/http"

	skill "github.com/RyuaNerin/go-kakaoskill/v2"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/simple1", skill.F(simpleFunc))
	mux.Handle("/simple2", skill.H(new(simpleHandler)))

	mh := skill.NewMuxHelper(mux, "")
	mh.F("/simple3", simpleFunc)
	mh.H("/simple4", new(simpleHandler))

	server := http.Server{
		Handler: mux,
	}

	server.ListenAndServe()
}

func simpleFunc(ctx *skill.Context) {
	ctx.WriteSimpleText("simple text")
}

type simpleHandler int

func (h simpleHandler) Handle(ctx *skill.Context) {
	ctx.WriteSimpleText("simple text")
}
```

```shell
$ go run example.go
```

## Usage

### Response Code
- `200 OK`
    - 아래 호출 시
        - `Context.WriteSimpleText`
        - `Context.WriteSimpleImage`
        - `Context.WriteResponse`
- `204 No Content`
    - `Context.Write~~~` 시리즈 함수가 호출되지 않고 핸들러가 종료되었을 때
- `400 Bad Request`
    - `Method` 가 `POST` 가 아닐 때
- `406 Not Acceptable`
    - `skill.AllowedRemoteAddr` 가 설정되어 있고, 목록에 없는 아이피일 때
- `500 Internal Error`
    - payload 파싱에 실패했을 때. (panic)

### 카카오톡 스킬 서버 IP

- `skill.AllowedRemoteAddr` 에서 접근을 허용할 IP 를 지정해줄 수 있습니다.
- `skill.AllowedRemoteAddr` 를 비워주시면 모든 IP 대역에서의 접근이 허용됩니다.
- 기본값으로 아래 IP 가 선언되어 있습니다.
	- 219.249.231.40
	- 219.249.231.41
	- 219.249.231.42
- `nginx` 의 `proxy_pass` 와 같은 서비스 사용시 아래와 같이 설정해주세요
    - `skill.FP` `skill.HP` 함수를 사용해주세요.
        - Request Header 에서 찾으면 해당 헤더의 값을, 없으면 `Request.RemoteAddr` 필드가 사용됩니다.
    - 예시
        - nginx.conf
        ```nginx
        ...

        location ... {
            proxy_pass ...;
            proxy_set_header X-Real-IP $remote_addr;
        }

        ...
        ```

        - main.go
        ```go
        // example.go
        package main

        import (
            "net/http"

            skill "github.com/RyuaNerin/go-kakaoskill/v2"
        )

        func main() {
            mux := http.NewServeMux()
            mux.HandleFunc("/simple1", skill.FP("X-Real-IP", simpleFunc))
            mux.Handle("/simple2", skill.HP("X-Real-IP", new(simpleHandler)))

            mh := skill.NewMuxHelper(mux, "X-Real-IP")
            mh.F("/simple3", simpleFunc)
            mh.H("/simple4", new(simpleHandler))

            server := http.Server{
                Handler: mux,
            }

            server.ListenAndServe()
        }

        func simpleFunc(ctx *skill.Context) {
            ctx.WriteSimpleText("simple text")
        }

        type simpleHandler int

        func (h simpleHandler) Handle(ctx *skill.Context) {
            ctx.WriteSimpleText("simple text")
        }
        ```