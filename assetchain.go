package main

import (
	"fmt"

	"github.com/assetcloud/assetchain/version"
)

var assetchain = fmt.Sprintf(`
TestNet=false
version="%s"

[crypto]
enableTypes=[]    #设置启用的加密插件名称，不配置启用所有
[crypto.enableHeight]  #配置已启用插件的启用高度，不配置采用默认高度0， 负数表示不启用
bls=0
btcscript=0

[crypto.sub.secp256k1eth] 
evmChainID=898
coinsPrecision=1e4

[address]
defaultDriver="eth"
[address.enableHeight]
eth=0
btc=0
btcMultiSign=0

[blockchain]
defCacheSize=128
maxFetchBlockNum=128
timeoutSeconds=5
batchBlockNum=128
driver="leveldb"
isStrongConsistency=false
# 分片存储中每个大块包含的区块数，固定参数
chunkblockNum=1000
# blockchain模块保留的区块数，指定最新的reservedBlockNum个区块不参与分片
reservedBlockNum=300000

[p2p]
enable=true
msgCacheSize=10240
driver="leveldb"

[p2p.sub.gossip]
serverStart=true

#[p2p.sub.dht]


[p2p.sub.dht.broadcast]
# 区块哈希广播最小大小 100KB
minLtBlockSize=100
#关闭交易批量广播功能, 后续全网升级后再开启
disableBatchTx=true
#关闭轻广播方案, 后续全网升级后再开启
disableLtBlock=true

# price 模式
[mempool]
name="price"
poolCacheSize=102400
minTxFeeRate=1
maxTxFee=1000000000
isLevelFee=true

[mempool.sub.score]
poolCacheSize=102400
minTxFee=1
maxTxNumPerAccount=100
timeParam=1      #时间占价格比例
priceConstant=1544  #手续费相对于时间的一个合适的常量,取当前unxi时间戳前四位数,排序时手续费高1e-5~=快1s
pricePower=1     #常量比例

[mempool.sub.price]
poolCacheSize=102400

[consensus]
name="ticket"
minerstart=true
genesisBlockTime=1703667518
genesis="0x94e67ceba77c19c683db4e24675b1f5c29e0226a" # TODO.address.1 genesis
genesisAmount=202005201314
minerExecs=["ticket", "autonomy"]
enableBestBlockCmp=true

[mver.consensus]
fundKeyAddr = "0xabb92bbb064e57463c7032ceb5a4372ad208a047" # TODO.address.1 user1_addr
powLimitBits = "0x1f00ffff"
maxTxNumber = 1500

[mver.consensus.ticket]
coinReward = 18
coinDevFund = 12
ticketPrice = 10000
retargetAdjustmentFactor = 4
futureBlockTime = 15
ticketFrozenTime = 43200
ticketWithdrawTime = 172800
ticketMinerWaitTime = 7200
targetTimespan = 2160
targetTimePerBlock = 15

[mver.consensus.ticket.ForkChainParamV2]
coinReward = 5
coinDevFund = 3
targetTimespan = 720
targetTimePerBlock = 5
ticketPrice = 3000

[mver.consensus.ForkTicketFundAddrV1]
fundKeyAddr = "0xabb92bbb064e57463c7032ceb5a4372ad208a047" # TODO.address.1 user1_addr

[consensus.sub.ticket] 
genesisBlockTime=1703667518

[[consensus.sub.ticket.genesis]]
minerAddr="0xbd75dfb9315199d6bb40db35e482007837a5e068" # TODO.address.1 user9_addr
returnAddr="0x0422121d85ac44609e83f343b696a96fce176f3b" # TODO.address.1 user8_addr
count=20000

 
[[consensus.sub.ticket.genesis]]
minerAddr="0xdd51224232045bf9e85686ae4e91695740d86ae8" # TODO.address.1 user4_addr
returnAddr="0x1fcb4ce2c94e883009b5ba29a8bbfa3b221ebb4e" # TODO.address.1   user5_addr 
count=3000 

[[consensus.sub.ticket.genesis]]
minerAddr="0xb2ee80a1cc031a0ec4291491aacf89c74dae2c4b" # TODO.address.1  user6_addr
returnAddr="0x077081f846a96d0d6bb9e9cf3c47ff8cd708f3d9" # TODO.address.1 user7_addr
count=4733

[store]
name="kvmvccmavl"
driver="leveldb"
storedbVersion="2.0.0"

[wallet]
minFee=1
driver="leveldb"
signType="secp256k1"

[exec]
proxyExecAddress="0x0000000000000000000000000000000000200005"
[exec.sub.coins]
#允许evm执行器操作coins
friendExecer=["evm"]

[exec.sub.token]
#配置一个空值，防止配置文件被覆盖
tokenApprs = ["0xd09d60dbc1d572cf01f58ffb87866c1fee0b4394", "0xc6a307aa74f71a3796f2240c13fe7670eeed7a42"] # TODO.address.1 tokanManage manager3_addr

[exec.sub.relay]
genesis="0x01f0ddbaf9c73510b23ba51c81e931f7488f2422" # TODO.address.1 relayManage

[exec.sub.manage]
superManager=[
    "0xfde1e130268184de2da955302bb3ce271c5738d8",  # TODO.address.1 manager1
	"0x14e92e09bf973dc2d842b042609173003878ab3d", # manager2
	"0xa3d3da443260436f32b37a0fd29f1eaf0094e06b", # manager3
	"0x94bb77a68cfff2706a19b580c46734b862b464f7", # 辅助manager1
]
#自治合约执行器名字
autonomyExec="autonomy"

[exec.sub.paracross]
nodeGroupFrozenCoins=0
#平行链共识停止后主链等待的高度
paraConsensusStopBlocks=30000
#配置平行链资产跨链交易的高度列表，title省略user.p,不同title使用,分割，不同hit高度使用"."分割，
#不同ignore高度区间用"."分割，区间内部使用"-"分割，hit高度在ignore范围内，为平行链自身的高度，不是主链高度
## para.hit.10.50.250, para.ignore.1-100.200-300
paraCrossAssetTxHeightList=[]


[exec.sub.autonomy]
total="0x36e041f189753673079191f91568c3c5f9ad7c26" # TODO.address.1 autonomyManage
useBalance=false

[mver.autonomy]
#最小委员会数量
minBoards=3
#最大委员会数量
maxBoards=40
#公示一周时间，以区块高度计算
publicPeriod=120960
#单张票价
ticketPrice=3000
#重大项目公示金额阈值
largeProjectAmount=1000000
#创建者消耗金额bty
proposalAmount=500
#董事会成员赞成率，百分比，可修改
boardApproveRatio=51
#全体持票人参与率，百分比
pubAttendRatio=75
#全体持票人赞成率，百分比
pubApproveRatio=66
#全体持票人否决率，百分比
pubOpposeRatio=33
#提案开始结束最小周期高度
startEndBlockPeriod=720
#提案高度 结束高度最大周期 100W
propEndBlockPeriod=1000000
#最小董事会赞成率
minBoardApproveRatio=50
#最大董事会赞成率
maxBoardApproveRatio=66
#最小全体持票人否决率
minPubOpposeRatio=33
#最大全体持票人否决率
maxPubOpposeRatio=50
#可以调整，但是可能需要进行范围的限制：参与率最低设置为 50， 最高设置为 80，赞成率，最低 50.1，最高80
#不能设置太低和太高，太低就容易作弊，太高则有可能很难达到
#最小全体持票人参与率
minPubAttendRatio=50
#最大全体持票人参与率
maxPubAttendRatio=80
#最小全体持票人赞成率
minPubApproveRatio=50
#最大全体持票人赞成率
maxPubApproveRatio=80
#最小公示周期
minPublicPeriod=120960
#最大公示周期
maxPublicPeriod=241920
#最小重大项目阈值(coin)
minLargeProjectAmount=1000000
#最大重大项目阈值(coin)
maxLargeProjectAmount=3000000
#最小提案金(coin)
minProposalAmount=20
#最大提案金(coin)
maxProposalAmount=2000	
#每个时期董事会审批最大额度300万
maxBoardPeriodAmount =3000000
#时期为一个月
boardPeriod=518400
#4w高度，大概2天 (未生效)
itemWaitBlockNumber=40000

[exec.sub.evm]
addressDriver="eth"
ethMapFromExecutor="coins"
#title的币种名称
ethMapFromSymbol="AS" 
#当前最大为200万
evmGasLimit=200000000

[exec.sub.evm.preCompile]
# 激活合token-erc20 的合约管理地址，必须配置管理员地址
superManager=["0xfde1e130268184de2da955302bb3ce271c5738d8","0x93b8a430152d390eaa8683909b3ae0745b1b5e24"] # TODO.address.1 manager1 manager2_addr


#系统中所有的fork,默认用chain的测试网络的
#但是我们可以替换
[fork.system]
ForkChainParamV1= 0
ForkCheckTxDup=0
ForkBlockHash= 0
ForkMinerTime= 0
ForkTransferExec= 0
ForkExecKey= 0
ForkTxGroup= 0
ForkResetTx0= 0
ForkWithdraw= 0
ForkExecRollback= 0
ForkCheckBlockTime=0
ForkTxHeight= 0
ForkTxGroupPara= 0
ForkChainParamV2= 0
ForkMultiSignAddress=0
ForkStateDBSet=0
ForkLocalDBAccess=0
ForkBlockCheck=0
ForkBase58AddressCheck=0
ForkEnableParaRegExec=0
ForkCacheDriver=0
ForkTicketFundAddrV1=0
#fork for 6.4
ForkRootHash=0 
#eth address key format fork
ForkFormatAddressKey=0
ForkCheckEthTxSort=0
ForkProxyExec=0

[fork.sub.evm]
Enable=0
ForkEVMABI=0
ForkEVMYoloV1=0
ForkEVMState=0
ForkEVMFrozen=0
ForkEVMTxGroup=0
ForkEVMKVHash=0
ForkEVMMixAddress=0
ForkIntrinsicGas=0
ForkEVMAddressInit=0
ForkEvmExecNonce=0
ForkEvmExecNonceV2=0

[fork.sub.rollup]
Enable=0


[fork.sub.none]
ForkUseTimeDelay=0

[fork.sub.coins]
Enable=0
ForkFriendExecer=0

[fork.sub.ticket]
Enable=0
ForkTicketId = 0
ForkTicketVrf = 0

[fork.sub.retrieve]
Enable=0
ForkRetrive=0
ForkRetriveAsset=0

[fork.sub.hashlock]
Enable=0
ForkBadRepeatSecret=0

[fork.sub.manage]
Enable=0
ForkManageExec=0
ForkManageAutonomyEnable=0

[fork.sub.token]
Enable=0
ForkTokenBlackList= 0
ForkBadTokenSymbol= 0
ForkTokenPrice= 300000
ForkTokenSymbolWithNumber=0
ForkTokenCheck= 0
ForkTokenEvm=0

[fork.sub.trade]
Enable=0
ForkTradeBuyLimit= 0
ForkTradeAsset= 0
ForkTradeID = 0
ForkTradeFixAssetDB=0
ForkTradePrice=0

[fork.sub.paracross]
Enable=0
ForkParacrossWithdrawFromParachain=0
ForkParacrossCommitTx=0
ForkLoopCheckCommitTxDone=0
#fork for 6.4
ForkParaAssetTransferRbk=0    
ForkParaSelfConsStages=0
ForkParaSupervision=0
ForkParaAutonomySuperGroup=0
#仅平行链适用
ForkParaFullMinerHeight=-1
ForkParaRootHash=-1
ForkParaFreeRegister=0
ForkParaCheckTx=0

[fork.sub.multisig]
Enable=0

[fork.sub.autonomy]
Enable=0
ForkAutonomyDelRule=0
ForkAutonomyEnableItem=0

[fork.sub.unfreeze]
Enable=0
ForkTerminatePart=0
ForkUnfreezeIDX= 0

[fork.sub.store-kvmvccmavl]
ForkKvmvccmavl=0

[health]
checkInterval=1
unSyncMaxTimes=2

`, version.Version)
