echo "GET query chaincode on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/ticket?peer=peer1.org1.example.com&fcn=queryObj&args=%5B%22obj20180101001%22%5D&username=admin&orgname=Org1" \
  -H "content-type: application/json"
echo
echo

echo "GET query chaincode on peer1 of Org2"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/ticket?peer=peer1.org1.example.com&fcn=queryObj&args=%5B%22obj20180101001%22%5D&username=admin&orgname=Org2" \
  -H "content-type: application/json"
echo
echo