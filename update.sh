cd /home/fabric3/paper_system/fabric-sdk-py/test/fixtures/chaincode/src/github.com/paper_chaincode/
docker cp paper cli:/opt/gopath/src/github.com/hyperledger/fabric/examples/chaincode/go
peer chaincode install -n paper -v 2.0 -p github.com/hyperledger/fabric/examples/chaincode/go/paper2.0
peer chaincode upgrade -o orderer.example.com:7050 -C mychannel -n paper -v 2.0 -c '{"Args":["init"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
# nohup python3 manage.py runserver 0.0.0.0:8000 &> server.log &