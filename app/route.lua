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
