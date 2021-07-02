# _*_ coding: utf-8 _*_
from hfc.protos.peer.proposal_response_pb2 import Response
from django.http import HttpResponse
from django.http import JsonResponse
from django.shortcuts import render
from rest_framework.parsers import JSONParser
import sys
# from state1_client import state1_client
# from kvdb import DB
import json
# import utils.util as util
import os 
# import trace_pb2_grpc as trace_service
# import trace_pb2
# import grpc
import asyncio
from hfc.fabric import Client
# sys.path.append("/home/fabric3/paper_system/fabric-sdk-py")
# import new_type_cross.listener as listener

# reload(sys)
# sys.setdefaultencoding('utf-8')

def home(request):
    return render(request, 'home.html')

loop = asyncio.get_event_loop()


cli = Client(net_profile="/home/fabric3/paper_system/fabric-sdk-py/test/fixtures/network.json")
org1_admin = cli.get_user('org1.example.com', 'Admin')

find_reduestor = {
    'org1_admin':org1_admin
}

businesschannel = cli.new_channel('businesschannel')
gopath_bak = os.environ.get('GOPATH', '')
gopath = os.path.normpath(os.path.join(
                      os.path.dirname(os.path.realpath('__file__')),
                      '../test/fixtures/chaincode'
                     ))
os.environ['GOPATH'] = os.path.abspath(gopath)

find_channel = {
    'businesschannel':businesschannel
}


policy = {
    'identities': [
        {'role': {'name': 'member', 'mspId': 'Org1MSP'}},
        {'role': {'name': 'admin', 'mspId': 'Org1MSP'}},
    ],
    'policy': {
        '1-of': [
            {'signed-by': 0}, {'signed-by': 1},
        ]
    }
}

# def updateNewPaper(request):
#     requestor_str = request.GET['requestor']
#     requestor = find_reduestor[requestor_str]
#     channel_name = request.GET['channel_name']
#     peers = request.GET.getlist('peers')
#     args = request.GET.getlist('args')
#     cc_name = request.GET['cc_name']
#     cc_version = request.GET['cc_version']
#     cc_endorsement_policy = policy
#     collections_config = None
#     transient_map = None
#     wait_for_event = True
#     response = loop.run_until_complete(cli.chaincode_upgrade(
#             requestor=requestor,
#             channel_name=channel_name,
#             peers=peers,
#             args=args,
#             cc_name=cc_name,
#             cc_version=cc_version,
#             cc_endorsement_policy=cc_endorsement_policy,
#             collections_config=collections_config,
#             transient_map=transient_map, 
#             wait_for_event=wait_for_event 
#     ))
#     return HttpResponse(response)

def transCode(args):
    i = 0
    while i < len(args):
        if type(args[i]) == str:
            args[i] = args[i].encode('utf-8')
        i = i + 1
    return args
    


def addConference(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    function = 'addConference'
    cc_name = 'paper'
    
    # for a in args:
    #     print(a, end=' ')
    
    # i = 0
    # while i < len(args):
    #     if type(args[i]) == str:
    #         args[i] = args[i].encode('utf-8')
    #     i = i + 1

    args = transCode(args)


    # cc_version = request.GET['cc_version']
    # cc_endorsement_policy = policy
    # collections_config = None
    # transient_map = None
    # wait_for_event = True
    # response = loop.run_until_complete(cli.chaincode_invoke(
            # requestor=requestor,
    #         channel_name=channel_name,
    #         peers=peers,
    #         fcn = fcn,
    #         args=args,
    #         cc_name=cc_name,
    #         # cc_version=cc_version,
    #         # cc_endorsement_policy=cc_endorsement_policy,
    #         # collections_config=collections_config,
    #         transient_map=transient_map, 
    #         wait_for_event=wait_for_event
    # ))
    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    return HttpResponse(response)


def addJournal(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'addJournal'
    cc_name = 'paper'

    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    return HttpResponse(response)

def addSoftware(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'addSoftware'
    cc_name = 'paper'
    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    return HttpResponse(response)


def addPatent(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'addPatent'
    cc_name = 'paper'

    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    return HttpResponse(response)


def traceBackwardPaper(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'traceBackwardPaper'
    cc_name = 'paper'

    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    print(response)
    sys.stdout.flush()
    return HttpResponse(response)

def traceBackwardConOrPatent(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'traceBackwardConOrPatent'
    cc_name = 'paper'

    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    return HttpResponse(response)

def traceBackwardConOrPatentFromTxid(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'traceBackwardConOrPatentFromTxid'
    cc_name = 'paper'

    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    return HttpResponse(response)


def traceBackwardPaperFromTxid(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'traceBackwardPaperFromTxid'
    cc_name = 'paper'

    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    return HttpResponse(response)


def traceBackwardAllPaperFromHashTitle(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'traceBackwardPaperFromTxid'
    cc_name = 'paper'
    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = 'traceBackwardPaper',
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    response_json = json.loads(response)
    txid = response_json['this_tx_id']
    args[0] = txid
    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    # print(response)
    # print(type(response))
    response_json = json.loads(response)
    print(response_json)
    print(response_json['last_tx_id'])
    result = []
    result.append(response_json)
    


    while response_json['last_tx_id'] != "":
        args[0] = response_json['last_tx_id']
        response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
        response_json = json.loads(response)
        result.append(response_json)
    result_json = json.dumps(result, ensure_ascii=False)
    return HttpResponse(result_json)



def traceBackwardAllConOrPatentFromHashTitle(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'traceBackwardConOrPatentFromTxid'
    cc_name = 'paper'
    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = 'traceBackwardConOrPatent',
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    response_json = json.loads(response)
    print(response_json)
    txid = response_json['this_tx_id']
    args[0] = txid
    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    # print(response)
    # print(type(response))
    response_json = json.loads(response)
    print(response_json)
    print(response_json['last_tx_id'])
    result = []
    result.append(response_json)
    


    while response_json['last_tx_id'] != "":
        args[0] = response_json['last_tx_id']
        response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
        response_json = json.loads(response)
        result.append(response_json)
    result_json = json.dumps(result, ensure_ascii=False)
    return HttpResponse(result_json)


def GetPaper(request):
    requestor_str = 'org1_admin'
    requestor = find_reduestor[requestor_str]
    channel_name = 'businesschannel'
    peers = ['peer0.org1.example.com']
    args = request.GET.getlist('args')
    args = transCode(args)
    function = 'GetPaper'
    cc_name = 'paper'

    response = loop.run_until_complete(cli.chaincode_invoke(
               requestor=requestor,
               channel_name=channel_name,
               peers=peers,
               fcn = function,
               args=args,
               cc_name='paper',
               transient_map=None, # optional, for private data
               wait_for_event=True, # for being sure chaincode invocation has been commited in the ledger, default is on tx event
               #cc_pattern='^invoked*' # if you want to wait for chaincode event and you have a `stub.SetEvent("invoked", value)` in your chaincode
               ))
    return HttpResponse(response)







