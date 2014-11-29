# 처음 만나는 텔레그램 CLI와 GO

컨테이너가 제대로 동작하는지 확인해 봅시다.

# 쉘 스크립트

매번 docker run을 입력하기 귀찮기 때문에 아래와 같이 쉘 스크립트를 만듭니다.

````
cat > run.sh
#!/bin/bash
docker run --rm -it -v /Users/lameduck/project/GO-Bot/app:/app gobot
[CTRL + D]

chmod +x run.sh
````

이제 부터는 `./run.sh`로 들어갈 수 있습니다.


# Hello Go Bot

아래의 파일을 로컬의 app 디렉토리에 `hello.go`로 저장합니다.

````
package main

import "fmt"

func main() {
	    fmt.Println("Hello, Go Bot")
}
````

## GO 프로그램 실행

컨테이너에서 다음과 같이 입력합니다.

````
cd /app
go run hello.go
````

## GO 바이너리 빌드

실행 파일을 만들기 위해서는 다음과 같이 입력합니다.

````
go build hello.go
````

# 텔레그램 CLI

````
cd /tg
bin/telegram-cli
````

최초 수행시에 패스 워드 인증이 나옵니다. 다음과 같이 진행합니다.

````
# bin/telegram-cli 
change_user_group: can't find the user telegramd to switch to
Telegram-cli version 1.1.1, Copyright (C) 2013-2014 Vitaly Valtman
Telegram-cli comes with ABSOLUTELY NO WARRANTY; for details type `show_license'.
This is free software, and you are welcome to redistribute it
under certain conditions; type `show_license' for details.
I: config dir=[/root/.telegram-cli]
Telephone number (with '+' sign): +82103097####
Code from sms (if you did not receive an SMS and want to be called, type "call"): 24876
User L Kim offline (was online [2014/11/28 22:49:03])
User L Kim online
> 
````

사용자 이름 사이의 공백은 _으로 처리되는 것을 볼 수 있습니다. 이름이 영어이면 user_info로 질의할 수 있습니다. (한글인 경우 아직 쉘 설정이 문제가 있네요.)

````
> user_info L_KIM
> user_info L_Kim
User L Kim updated name
User L Kim (#73270521):
	real name: L Kim
	phone: 82103097####
	online
````

user_info에 넣었던 L_Kim으로도 메시지를 보낼 수 있고 #넘버 앞에 user를 붙여서도 메시지를 보낼 수 있는 것을 볼 수 있습니다.

````
> msg L_Kim hello telegram
[23:32]  L Kim <<< hello telegram
> msg user#73270521 Nice to meet you
[23:33]  L Kim <<< Nice to meet you
````

우리는 이 특성을 이용할 것입니다.
