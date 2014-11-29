# 전송하기

## GO 언에 수준 메시지 전송

````
package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	subProcess := exec.Command("/tg/bin/telegram-cli", "-k", "/tg/server.pub", "-W")

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdin.Close()

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	subProcess.Start()

	command := fmt.Sprintf("msg %s %s\nsafe_quit\n", "L_Kim", "nice to meet you")
	fmt.Println(command)
	io.WriteString(stdin, command)
	subProcess.Wait()
}
````

## 메시지를 커맨드 라인 인자로 바꿔 봅시다.

````
package main

import (
        "fmt"
        "io"
        "os"
        "os/exec"
)

func main() {
	var text string
	if len(os.Args) > 1 {
		text = os.Args[1]
	} else {
		text = "Hello GoBot"
	}
        subProcess := exec.Command("/tg/bin/telegram-cli", "-k", "/tg/server.pub", "-W")

        stdin, err := subProcess.StdinPipe()
        if err != nil {
                fmt.Println(err)
        }
        defer stdin.Close()

        subProcess.Stdout = os.Stdout
        subProcess.Stderr = os.Stderr

        subProcess.Start()

        command := fmt.Sprintf("msg %s %s\nsafe_quit\n", "L_Kim", text)
        fmt.Println(command)
        io.WriteString(stdin, command)
        subProcess.Wait()
}
````
