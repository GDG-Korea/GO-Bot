# 이미지도 보냅니다

전송 메시지만 바꾸면 가능합니다.

## Messaging

````
msg <peer> Text - sends message to this peer
fwd <user> <msg-seqno> - forward message to user. You can see message numbers starting client with -N
chat_with_peer <peer> starts one on one chat session with this peer. /exit or /quit to end this mode.
add_contact <phone-number> <first-name> <last-name> - tries to add contact to contact-list by phone
rename_contact <user> <first-name> <last-name> - tries to rename contact. If you have another device it will be a fight
mark_read <peer> - mark read all received messages with peer
delete_msg <msg-seqno> - deletes message (not completly, though)
restore_msg <msg-seqno> - restores delete message. Impossible for secret chats. Only possible short time (one hour, I think) after deletion
````

## Multimedia

````
send_photo <peer> <photo-file-name> - sends photo to peer
send_video <peer> <video-file-name> - sends video to peer
send_text <peer> <text-file-name> - sends text file as plain messages
load_photo/load_video/load_video_thumb/load_audio/load_document/load_document_thumb <msg-seqno> - loads photo/video/audio/document to download dir
view_photo/view_video/view_video_thumb/view_audio/view_document/view_document_thumb <msg-seqno> - loads photo/video to download dir and starts system default viewer
fwd_media <msg-seqno> send media in your message. Use this to prevent sharing info about author of media (though, it is possible to determine user_id from media itself, it is not possible get access_hash of this user)
set_profile_photo <photo-file-name> - sets userpic. Photo should be square, or server will cut biggest central square part
````

## 이미지 예제

`app` 디렉토리에 `gonuts.png`파일이 있습니다.
![짱 귀여운 이미지](https://github.com/GDG-Korea/GO-Bot/blob/3-img/app/gonuts.png)

## 다음

[대화해보아요](https://github.com/GDG-Korea/GO-Bot/tree/4-interactive)
