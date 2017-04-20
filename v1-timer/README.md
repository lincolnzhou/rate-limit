# v1-timer
使用channel接住所有的request，之后再通过time.tick来控制每个请求的处理时间

请求的处理无法形成并发，必须一个个通过时间窗口去执行

## 测试结果
```
➜  rate-limit ab -n 100 -c 10 http://localhost:7000/health
This is ApacheBench, Version 2.3 <$Revision: 1757674 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            7000

Document Path:          /health
Document Length:        2 bytes

Concurrency Level:      10
Time taken for tests:   19.517 seconds
Complete requests:      100
Failed requests:        0
Total transferred:      11800 bytes
HTML transferred:       200 bytes
Requests per second:    5.12 [#/sec] (mean)
Time per request:       1951.720 [ms] (mean)
Time per request:       195.172 [ms] (mean, across all concurrent requests)
Transfer rate:          0.59 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.1      0       1
Processing:     1 1861 438.6   2000    2005
Waiting:        1 1861 438.6   2000    2005
Total:          1 1862 438.6   2000    2005

Percentage of the requests served within a certain time (ms)
  50%   2000
  66%   2001
  75%   2001
  80%   2001
  90%   2003
  95%   2004
  98%   2005
  99%   2005
 100%   2005 (longest request)
```