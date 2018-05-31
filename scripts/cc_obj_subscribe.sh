echo "POST invoke chaincode on peers of Org1"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/msisdn \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.example.com"],
	"fcn":"subscribe",
	"args":["13810167616","uuid1234","eki2","bcuser_id","bcuser_key"],
  "username":"admin",
  "orgname":"Org1"  
}')
echo "Transacton ID is $TRX_ID"
echo

# echo "POST invoke chaincode on peers of Org1"
# echo
# TRX_ID=$(curl -s -X POST \
#   http://localhost:4000/channels/mychannel/chaincodes/msisdn \
#   -H "content-type: application/json" \
#   -d '{
# 	"peers": ["peer0.org1.example.com","peer1.org1.example.com"],
# 	"fcn":"unsubscribe",
# 	"args":["13810167616","uuid1234","bcuser_id","bcuser_key"],
#   "username":"admin",
#   "orgname":"Org1"  
# }')
# echo "Transacton ID is $TRX_ID"
# echo