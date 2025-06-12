判断HTTP请求Body内容类型并处理数据
---
ch022_http_body


测试验证：
```shell
# 
curl -X POST http://localhost:8080/ -H "Content-Type: application/json" -d '{"name": "Andy", "age": 30}'

curl -X POST http://localhost:8080/ -H "Content-Type: application/x-www-form-urlencoded" -d "name=Andy&age=30"

curl -X POST http://localhost:8080/ -H "Content-Type: multipart/form-data" -F "file=@/Users/andyron/Downloads/file.txt"

```

## ref

https://www.oryoy.com/news/golang-shi-zhan-ru-he-gao-xiao-pan-duan-http-qing-qiu-body-nei-rong-lei-xing-bing-chu-li-shu-ju.html