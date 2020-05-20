# pass

[![License](https://img.shields.io/badge/license-MIT-blue.svg?label=License&maxAge=86400)](./LICENSE)
[![codebeat][1]][2]

[1]: https://codebeat.co/badges/d7a74a17-b99c-4d59-a07d-20685cf169e4 "Codebeat badge"
[2]: https://codebeat.co/projects/github-com-karpovdl-pass-master "Codebeat"

* pass --version, -v
* pass --help, -h

To build an application for windows make script/build.cmd

## run

Application pass - application for pass data
* pass run, pass r - run server
* pass run --help, -h

pass run -h - help command
* port, p - port to use listen server
Default: 9000
* pprof, pf - active pprof mode
If `flag` [true], then active pprof, otherwise [false]
Default: false
* pprof_port, pfp - port to use pprof listen server
Default: 9001

### run tm

* pass run tm, pass r tm - run aplication for pass data to telegram
* pass run tm --help, -h

* bot_token, bt - telegram bot token uuid
* channel_id, cid - telegram channel id
* message, m - telegram message

Configuration telegram.json file
```
{
    "bot_token": "{BOT_TOKEN}",
    "channel_id": {CHANNEL_ID},
    "message": "{MESSAGE}"
}
```

### sample run

Sample run server without parameters
```
pass.exe run
```

Sample run with active pprof
```
pass.exe ^
  "run" ^
  "--port=9000" ^
  "--pprof=true" ^
  "--pprof_port=9001"
```

Sample pass data to telegram from browser
```
http://localhost:9000/tm/?bt={BOT_TOKEN}&cid={CHANNEL_ID}&m={MESSAGE}
```

Sample pass data to telegram from browser if exist configuration telegram.json file
```
http://localhost:9000/tm/?m={MESSAGE}
```

Sample pass data to telegram from curl
```
curl -X GET http://localhost:9000/tm/?bt={BOT_TOKEN}&cid={CHANNEL_ID}&m={MESSAGE}
```

Sample pass data to telegram from curl if exist configuration telegram.json file
```
curl -X GET http://localhost:9000/tm/?m={MESSAGE}
```

Sample run application with command for pass data only to telegram
```
pass.exe ^
  "run" ^
  "tm" ^
  "--bt={BOT_TOKEN}" ^
  "--cid={CHANNEL_ID}" ^
  "-m={MESSAGE}"
```

Sample run application with command for pass data only to telegram if exist configuration telegram.json file
```
pass.exe ^
  "run" ^
  "tm" ^
  "-m={MESSAGE}"
```
