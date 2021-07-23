package main

// 这部分配置随代码发布，不能修改
var testchainconfig = `
TestNet=false


[blockchain]
maxFetchBlockNum=128
timeoutSeconds=5
batchBlockNum=128
driver="leveldb"
isStrongConsistency=false
disableShard=true
onChainTimeout=1

[p2p]
enable=true
msgCacheSize=10240
driver="leveldb"

[p2p.sub.gossip]
serverStart=true

[p2p.sub.dht]
#bootstraps是内置不能修改的引导节点

[mempool]


[mempool.sub.score]
poolCacheSize=102400
minTxFee=100000
maxTxNumPerAccount=100
timeParam=1      #时间占价格比例
priceConstant=1544  #手续费相对于时间的一个合适的常量,取当前unxi时间戳前四位数,排序时手续费高1e-5~=快1s
pricePower=1     #常量比例

[mempool.sub.price]
poolCacheSize=102400

[consensus]
name="pos33"
minerstart=true
genesisBlockTime=1604449783
genesis="163eimvjqQacXh6evrmRpRxXRCxKqHUT88"
minerExecs=["pos33"]

[consensus.sub.pos33]
genesisBlockTime=1611627559
checkFutureBlockHeight=1500000
# listenPort="10901"
# bootPeers=["/ip4/183.129.226.76/tcp/10901/p2p/16Uiu2HAmErmNhtS145Lv5fe9FWrHSrNjPkp1eMLeLgi6t3sdr1of"]

[[consensus.sub.pos33.genesis]]
minerAddr="19AqkFqEEB4kRxfnYomuAx4TgxeHw53X99"
returnAddr="163eimvjqQacXh6evrmRpRxXRCxKqHUT88"
count=10000

[mver.consensus]
fundKeyAddr="14wDnMse45jomBWx1AuHB6KDhVz662YJMx"
powLimitBits="0x1f00ffff"
maxTxNumber=6000

[mver.consensus.ForkChainParamV1]
maxTxNumber=6000

[mver.consensus.ForkChainParamV2]
powLimitBits="0x1f2fffff"

[mver.consensus.ForkTicketFundAddrV1]
fundKeyAddr="14wDnMse45jomBWx1AuHB6KDhVz662YJMx"

[mver.consensus.pos33]
coinReward=18
coinDevFund=12
ticketPrice=10000
retargetAdjustmentFactor=4
futureBlockTime=5
ticketFrozenTime=43200
ticketWithdrawTime=10
ticketMinerWaitTime=7200
targetTimespan=2160
targetTimePerBlock=15

[store]

[exec]
[exec.sub.pos33]
ForkTicketId=0
ForkTicketVrf=0

[exec.sub.token]
#配置一个空值，防止配置文件被覆盖
tokenApprs=[13jXcaJB3Sg7oBB2FAZf3VbDB58gHn7PPd]
[exec.sub.relay]
genesis="13jXcaJB3Sg7oBB2FAZf3VbDB58gHn7PPd"

[exec.sub.manage]
superManager=[
    "13jXcaJB3Sg7oBB2FAZf3VbDB58gHn7PPd", 
]

[exec.sub.paracross]
nodeGroupFrozenCoins=0
#平行链共识停止后主链等待的高度
paraConsensusStopBlocks=30000

[exec.sub.autonomy]
total="13jXcaJB3Sg7oBB2FAZf3VbDB58gHn7PPd"
useBalance=false

[health]
listenAddr="localhost:8809"
checkInterval=1
unSyncMaxTimes=2
`
