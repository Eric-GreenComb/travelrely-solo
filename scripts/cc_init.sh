echo "POST instantiate chaincode on peer1 of Org1"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "content-type: application/json" \
  -d "{
	\"chaincodeName\":\"ticket\",
	\"chaincodeVersion\":\"v0\",
	\"chaincodeType\": \"golang\",
	\"args\":[\"\"],
  \"username\":\"admin\",
  \"orgname\":\"Org1\"  
}"
echo
echo