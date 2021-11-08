# xk6-youzan-nsq
Pressure test for youzan nsq

## 编译k6
### 本地编译
```shell
git clone git@github.com:zhqqqy/xk6-youzan-nsq.git
xk6 build --with github.com/zhqqqy/xk6-youzan-nsq=.
```

### 在线编译
```shell
xk6 build --with github.com/zhqqqy/xk6-youzan-nsq@latest
```
### 交叉编译
```shell
GOOS=linux GOARCH=amd64 xk6 build --with github.com/zhqqqy/xk6-youzan-nsq=.
```

## 数据同步到influxdb
```shell
./k6 run  --duration 30s --out influxdb=http://10.2.1.6:8086/myk6db - < scripts/test_json.js
```
