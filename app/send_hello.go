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
