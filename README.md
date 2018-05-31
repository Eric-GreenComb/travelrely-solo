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

## 访问URL

1. Health

    [http://localhost:4000/health](http://localhost:4000/health)

2. Query Get Block by BlockNumber

    [http://localhost:4000/channels/mychannel/blocks/4?peer=peer1&username=admin&orgname=org1](http://localhost:4000/channels/mychannel/blocks/4?peer=peer1&username=admin&orgname=org1)

3. Query Get Transaction by Transaction ID
    [http://localhost:4000/channels/mychannel/transactions/eb0fb438d4dc6a90095b78f66f2f5766c88a18eac88faa55ea38260256fb83ed?peer=peer1&username=admin&orgname=org1](http://localhost:4000/channels/mychannel/transactions/eb0fb438d4dc6a90095b78f66f2f5766c88a18eac88faa55ea38260256fb83ed?peer=peer1&username=admin&orgname=org1)

4. Query to fetch all Installed/instantiated chaincodes

    [http://localhost:4000/chaincodes?peer=peer1&username=Jim&orgname=org1&type=installed&channel=mychannel1](http://localhost:4000/chaincodes?peer=peer1&username=Jim&orgname=org1&type=installed&channel=mychannel1)

5. Query to fetch channels

    [http://localhost:4000/channels?peer=peer1&username=Jim&orgname=org1](http://localhost:4000/channels?peer=peer1&username=Jim&orgname=org1)

6. Query for Channel Information

    [http://localhost:4000/channels/mychannel1?peer=peer1&username=Jim&orgname=org1](http://localhost:4000/channels/mychannel1?peer=peer1&username=Jim&orgname=org1)

7. Query for Channel Height

    [http://localhost:4000/channels/mychannel1/height?peer=peer1&username=Jim&orgname=org1](http://localhost:4000/channels/mychannel1/height?peer=peer1&username=Jim&orgname=org1)
