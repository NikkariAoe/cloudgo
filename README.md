## curl 测试

```shell
$ curl -v http://127.0.0.1:9090
* Rebuilt URL to: http://127.0.0.1:9090/
* timeout on name lookup is not supported
*   Trying 127.0.0.1...
* TCP_NODELAY set
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to 127.0.0.                               1 (127.0.0.1) port 9090 (#0)
> GET / HTTP/1.1
> Host: 127.0.0.1:9090
> User-Agent: curl/7.51.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 14 Nov 2017 12:26:43 GMT
< Content-Length: 14
< Content-Type: text/plain; charset=utf-8
<
{ [14 bytes data]
* Curl_http_done: called premature == 0
100    14  100    14    0     0     14      0  0:00:01 --:--:--  0:00:01 14000Hello astaxie!
* Connection #0 to host 127.0.0.1 left intact

```



## ab测试

> hey工具包是由go实现的类似ab工具的负载测试工具，以下都是基于hey的测试例子,其中 -n表示发起的请求数，-c表示并发量，-cpus表示使用的cpu核数

```shell
$ ./hey.exe -n 1000 -c 1000 -cpus 2 http://127.0.0.1:9090
Summary:
  Total:        1.1093 secs
  Slowest:      1.0362 secs
  Fastest:      0.2672 secs
  Average:      0.8964 secs
  Requests/sec: 901.4745
  Total data:   14000 bytes
  Size/request: 14 bytes

Response time histogram:
  0.267 [1]     |
  0.344 [5]     |
  0.421 [5]     |
  0.498 [17]    |∎∎
  0.575 [18]    |∎∎
  0.652 [10]    |∎
  0.729 [5]     |
  0.806 [22]    |∎∎
  0.882 [211]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.959 [428]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.036 [278]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎

Latency distribution:
  10% in 0.8431 secs
  25% in 0.8741 secs
  50% in 0.9182 secs
  75% in 0.9652 secs
  90% in 0.9872 secs
  95% in 0.9942 secs
  99% in 1.0047 secs

Details (average, fastest, slowest):
  DNS+dialup:    0.5690 secs, 0.2282 secs, 0.9082 secs
  DNS-lookup:    0.0000 secs, 0.0000 secs, 0.0000 secs
  req write:     0.0056 secs, 0.0000 secs, 0.1251 secs
  resp wait:     0.1626 secs, 0.0225 secs, 0.5244 secs
  resp read:     0.0768 secs, 0.0000 secs, 0.1256 secs

Status code distribution:
  [200] 1000 responses


$ ./hey.exe -n 1000 -c 100 http://127.0.0.1:9090
Summary:
  Total:        0.4588 secs
  Slowest:      0.2086 secs
  Fastest:      0.0005 secs
  Average:      0.0396 secs
  Requests/sec: 2179.4821
  Total data:   14000 bytes
  Size/request: 14 bytes

Response time histogram:
  0.000 [1]     |
  0.021 [486]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.042 [188]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  0.063 [101]   |∎∎∎∎∎∎∎∎
  0.084 [87]    |∎∎∎∎∎∎∎
  0.105 [32]    |∎∎∎
  0.125 [78]    |∎∎∎∎∎∎
  0.146 [13]    |∎
  0.167 [3]     |
  0.188 [10]    |∎
  0.209 [1]     |

Latency distribution:
  10% in 0.0065 secs
  25% in 0.0140 secs
  50% in 0.0220 secs
  75% in 0.0595 secs
  90% in 0.1071 secs
  95% in 0.1206 secs
  99% in 0.1756 secs

Details (average, fastest, slowest):
  DNS+dialup:    0.0009 secs, 0.0000 secs, 0.1126 secs
  DNS-lookup:    0.0000 secs, 0.0000 secs, 0.0000 secs
  req write:     0.0006 secs, 0.0000 secs, 0.0645 secs
  resp wait:     0.0296 secs, 0.0000 secs, 0.1311 secs
  resp read:     0.0068 secs, 0.0000 secs, 0.0570 secs

Status code distribution:
  [200] 1000 responses

```

