# 대화해보아요

````
function on_msg_receive(msg)
        print("Name : ", msg.from.print_name)
        print("Phone : ", msg.from.phone)
        print("Msg Num : ", msg.id)
        print("Msg Text: ", msg.text)

        if (msg.text == "pic") then
                local file = io.popen('go run /app/send_image.go', 'r')
                local output = file:read('*all')
                file:close()
                print(output)
        end
end

function ok_cb(extra, success, result)
end
````

아쉽게도 telegram-cli가 제공하는 환경에서는 GO언어로 대화형 환경을 만들기가 쉽지 않습니다. 그래서 lua 언어의 도움을 받겠습니다.

위의 내용을 route.lua 정도로 저장합시다.

````
/tg/bin/telegram-cli -s /app/route.lua
````

pic이란 커맨드를 인식합니다.


## 다음

[마무리](https://github.com/GDG-Korea/GO-Bot/tree/5-end)
