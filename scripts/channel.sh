echo
echo "POST request Create channel  ..."
echo
curl -s -X POST \
  http://localhost:4000/channels \
  -H "content-type: application/json" \
  -d '{
	"channelName":"mychannel",
	"channelConfigPath":"../artifacts/channel/mychannel.tx",
  "username":"admin",
  "orgname":"Org1"
}'
echo
echo
sleep 10
echo "POST request Join channel on Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.travelrely.com"],
  "username":"admin",
  "orgname":"Org1"
}'
echo
echo

echo "POST request Join channel on Org2"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org2.travelrely.com"],
  "username":"admin",
  "orgname":"Org2"  
}'
echo
echo

echo "POST request Join channel on Org3"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org3.travelrely.com"],
  "username":"admin",
  "orgname":"Org3"  
}'
echo
echo