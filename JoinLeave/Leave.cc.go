{{/*
        Made by Rhyker/Ranger (779096217853886504)

    Trigger Type: `Join Message in channel`
©️ Dynamic 2021
MIT License
*/}}


{{$embed := cembed
            "author" (sdict "url" (.User.AvatarURL "4096") "name" "User Left" "icon_url" (.User.AvatarURL "1024"))
            "description" (print  .User.Username "\n[" .User.String "]" " : " "`" .User.ID "`" "\nAccount created " "**" currentUserAgeHuman "** ago" "\nWe now have" "** " .Guild.MemberCount " **" "members")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 16711680
            }}
{{sendMessage nil $embed}}
