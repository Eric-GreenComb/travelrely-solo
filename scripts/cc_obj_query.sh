echo "GET query chaincode on peer0 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/msisdn?peer=peer0.org1.travelrely.com&fcn=msisdn_state&args=%5B%2213810167616%22%5D&username=admin&orgname=Org1" \
  -H "content-type: application/json"
echo
echo

echo "GET query chaincode on peer0 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/msisdn?peer=peer0.org1.travelrely.com&fcn=get_msisdn_history&args=%5B%2213810167616%22%5D&username=admin&orgname=Org1" \
  -H "content-type: application/json"
echo
echo

echo "GET query chaincode on peer0 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/msisdn?peer=peer0.org1.travelrely.com&fcn=asset_info&args=%5B%22uuid1234%22%5D&username=admin&orgname=Org1" \
  -H "content-type: application/json"
echo
echo