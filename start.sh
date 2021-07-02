# HLF_VERSION=1.4.6
# docker pull hyperledger/fabric-peer:${HLF_VERSION} \
#     && docker pull hyperledger/fabric-orderer:${HLF_VERSION} \
#     && docker pull hyperledger/fabric-ca:${HLF_VERSION} \
#     && docker pull hyperledger/fabric-ccenv:${HLF_VERSION}
# cd fabric-sdk-py/
# docker-compose -f test/fixtures/docker-compose-2orgs-4peers-tls.yaml up
# docker-compose -f test/fixtures/ca/docker-compose.yml up
# pip3 install virtualenv; make venv
# source venv/bin/activate
# make install
# tox -e py3 -- test/integration/e2e_test.py # Run specified test case
# deactive

# sudo -s
# conda activate fabric

# docker-compose -f test/fixtures/docker-compose-2orgs-4peers-tls.yaml up
#   


# 更改相关的链码
# docker cp mycc2.0 cli:/opt/gopath/src/github.com/hyperledger/fabric/examples/chaincode/go
# peer chaincode install -n mycc -v 2.0 -p github.com/hyperledger/fabric/examples/chaincode/go/mycc2.0
# peer chaincode upgrade -o orderer.example.com:7050 -C mychannel -n mycc -v 2.0 -c '{"Args":["init"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
# nohup python3 manage.py runserver 0.0.0.0:8000 &> server.log &

cd fabric-sdk-py/
docker-compose -f test/fixtures/docker-compose-2orgs-4peers-tls.yaml up -d
docker-compose -f test/fixtures/ca/docker-compose.yml up -d

cd paper_system/
python creat_connection.py
python connect_chaincode.py

cd /home/fabric3/paper_system/fabric-sdk-py/paper_sever
nohup python3 manage.py runserver 0.0.0.0:8000 &> server.log &





