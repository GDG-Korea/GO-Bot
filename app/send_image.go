package main

import (
        "fmt"
        "io"
        "os"
        "os/exec"
)

func main() {
        img := "/app/gonuts.png"
        subProcess := exec.Command("/tg/bin/telegram-cli", "-k", "/tg/server.pub", "-W")

        stdin, err := subProcess.StdinPipe()
        if err != nil {
                fmt.Println(err)
        }
        defer stdin.Close()

        subProcess.Stdout = os.Stdout
        subProcess.Stderr = os.Stderr

        subProcess.Start()

        command := fmt.Sprintf("send_photo %s %s\nsafe_quit\n", "L_Kim", img)
        fmt.Println(command)
        io.WriteString(stdin, command)
        subProcess.Wait()
}
