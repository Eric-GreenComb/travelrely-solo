# Travelrely Solo

## 第一步 启动 fabric 1.0 网络

sudo ./run_fabric.sh -m restart

## 第二步 启动 node express http server

### npm install

npm install

### install forever

sudo npm install forever -g

### start node express

PORT=4000 forever start app

curl 127.0.0.1:4000

forever stop app

## 第三步 初始化 fabric 1.0 channel ,发布 chaincode

cd scripts

./enroll.sh

./channel.sh

./cc_install.sh

./cc_init.sh

## 第四步 运行测试数据

./cc_obj_subscribe.sh

./cc_obj_query.sh

./cc_obj_unsubscribe.sh

## 删除网络

sudo ./run_fabric.sh -m down

## 重新启动网络

sudo ./run_fabric.sh -m stop

sudo ./run_fabric.sh -m start

PORT=4000 forever start app

cd script

./cc_obj_query.sh

## 启动多节点

### 节点

114.215.82.68(外) 10.66.181.165（内） ca1,order,org1

115.28.51.50(外) 10.66.182.46（内） ca2,org2

118.190.137.46(外) 10.30.181.162（内） ca3,org3

### start

进入artifacts目录

- 114.215.82.68

sudo docker-compose up --no-deps ca.org1.travelrely.com orderer.travelrely.com peer0.org1.travelrely.com

sudo docker-compose up -d --no-deps ca.org1.travelrely.com orderer.travelrely.com peer0.org1.travelrely.com

- 115.28.51.50

sudo docker-compose up --no-deps ca.org2.travelrely.com peer0.org2.travelrely.com

sudo docker-compose up -d --no-deps ca.org2.travelrely.com peer0.org2.travelrely.com

- 118.190.137.46

sudo docker-compose up --no-deps ca.org3.travelrely.com peer0.org3.travelrely.com

sudo docker-compose up -d --no-deps ca.org3.travelrely.com peer0.org3.travelrely.com

sudo docker rm $(sudo docker ps -aq)