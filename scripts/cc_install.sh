# Language defaults to "golang"
LANGUAGE="golang"

echo "POST Install chaincode on Org1"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.org1.travelrely.com\"],
	\"chaincodeName\":\"msisdn\",
	\"chaincodePath\":\"github.com/msisdn\",
	\"chaincodeType\": \"golang\",
	\"chaincodeVersion\":\"v0\",
  \"username\":\"admin\",
  \"orgname\":\"Org1\"  
}"
echo
echo

echo "POST Install chaincode on Org2"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.org2.travelrely.com\"],
	\"chaincodeName\":\"msisdn\",
	\"chaincodePath\":\"github.com/msisdn\",
	\"chaincodeType\": \"golang\",
	\"chaincodeVersion\":\"v0\",
  \"username\":\"admin\",
  \"orgname\":\"Org2\"  
}"
echo
echo

echo "POST Install chaincode on Org3"
echo
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.org3.travelrely.com\"],
	\"chaincodeName\":\"msisdn\",
	\"chaincodePath\":\"github.com/msisdn\",
	\"chaincodeType\": \"golang\",
	\"chaincodeVersion\":\"v0\",
  \"username\":\"admin\",
  \"orgname\":\"Org3\"  
}"
echo
echo