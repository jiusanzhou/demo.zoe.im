curl 'https://jgchat.net/api/v1/user/realname_verify' \
  -H 'authority: jgchat.net' \
  -H 'accept: application/json, text/plain, */*' \
  -H 'authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcmsiLCJleHAiOjE2MTU4MDAzMzAsImlhdCI6MTYwOTc1MjMzMCwiaXNzIjoiYXJrIiwianRpIjoiYjQ5ZWFkY2YtMjMxNS00ZWJiLThlZDctNTNkOGZlMzE1YThkIiwibmJmIjoxNjA5NzUyMzI5LCJzdWIiOiIxMzMyOCIsInR5cCI6ImFjY2VzcyJ9.bCvq_j95Vbyis-vBprXTjgMc4GZduTV03nCVryWahClhhpUKnMwZ1kC1u39pBX6FVu4l3TETv0aB9zk0To-GVg' \
  -H 'user-agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36' \
  -H 'content-type: application/json;charset=UTF-8' \
  -H 'origin: https://jgchat.net' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://jgchat.net/explore' \
  -H 'accept-language: zh-CN,zh;q=0.9' \
  -H 'cookie: _UV=1609752990123746; NUV=1609776000000; ued_ping_tk185=1,1609752161907; DIFF=1609752390506; ONLINE_TIME=1609752390499; ued_ping_ssid2=160975299012374616097540796195661609752158530|7' \
  --data-binary '{"realname":"张舒琪","card_number":"440582199812205568"}' \
  --compressed