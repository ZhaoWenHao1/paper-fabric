from hfc.fabric import Client
from hfc.fabric_ca.caservice import ca_service
import asyncio

# cli = Client(net_profile="../test/fixtures/network.json")
cli = Client(net_profile="/home/fabric3/paper_system/fabric-sdk-py/test/fixtures/network.json")

# 声明相关的中心容器的配置

print(cli.organizations)  # orgs in the network
print(cli.peers)  # peers in the network
print(cli.orderers)  # orderers in the network
print(cli.CAs)  # ca nodes in the networkpyth
# 检查所有的配置文件，网络资源，检查资源

org1_admin = cli.get_user(org_name='org1.example.com', name='Admin') 
# get the admin user from local path（SDK将从本地路径加载有效凭据）


"""
要使用CA，必须启动CA服务器。例如，
$ docker-compose -f test/fixtures/ca/docker-compose.yml up
"""


# 从Fabric CA获取Credentail
casvc = ca_service(target="http://127.0.0.1:7054")
# 以管理员身份注册到Fabric CA；
adminEnrollment = casvc.enroll("admin", "adminpw") 
# 当地拥有管理员的环境
# secret = adminEnrollment.register("user1") 
# 注册用户user1到CA
# user1Enrollment = casvc.enroll("user1", secret) 
# 注册新用户user1并获得本地证书；
# user1ReEnrollment = casvc.reenroll(user1Enrollment) 
# 重新注册user1
RevokedCerts, CRL = adminEnrollment.revoke("user1") 
# 撤销user1



"""
使用网络操作频道
创建一个新频道并加入
使用sdk创建一个新频道，并允许对等方加入。
"""

loop = asyncio.get_event_loop()
org1_admin = cli.get_user(org_name='org1.example.com', name='Admin')
orderer_admin = cli.get_user(org_name='orderer.example.com', name='Admin')



# 创建一个新的通道
response = loop.run_until_complete(cli.channel_create(
            orderer='orderer.example.com',
            channel_name='businesschannel',
            requestor=org1_admin,
            config_yaml='/home/fabric3/paper_system/fabric-sdk-py/test/fixtures/e2e_cli/',
            channel_profile='TwoOrgsChannel'
            ))
print(response == True)




# 将peer加入到channel中

responses = loop.run_until_complete(cli.channel_join(
               requestor=org1_admin,
               channel_name='businesschannel',
               peers=['peer0.org1.example.com',
                      'peer1.org1.example.com'],
               orderer='orderer.example.com'
               ))
print(responses)

# MSP：在区块链网络中用于颁发和验证证书和身份的一组加密机制和协议
# 将来自不同的MSP协议的peer加入到channel中
org2_admin = cli.get_user(org_name='org2.example.com', name='Admin')

# 对org2.example.com的peer进行操作，org2_admin作为请求方
responses = loop.run_until_complete(cli.channel_join(
               requestor=org2_admin,
               channel_name='businesschannel',
               peers=['peer0.org2.example.com',
                      'peer1.org2.example.com'],
               orderer='orderer.example.com'
               ))
print(responses)





 