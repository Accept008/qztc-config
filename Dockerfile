# build
FROM registry.cn-hangzhou.aliyuncs.com/sonicmoving/golang:1.9 AS build
WORKDIR /opt/qztc-config
ADD main /opt/qztc-config
ADD /toml/old.toml /opt/qztc-config
ADD /toml/qa.toml /opt/qztc-config

CMD  ["./main","-config=./qa.toml"]


