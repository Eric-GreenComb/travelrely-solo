echo "POST invoke chaincode on peers of Org1"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/ticket \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.example.com","peer1.org1.example.com"],
	"fcn":"createObj",
	"args":["obj20180101001","this is a 1"],
  "username":"admin",
  "orgname":"Org1"  
}')
echo "Transacton ID is $TRX_ID"
echo