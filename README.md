# assetchain
Yuan Chain Project
# pos33 运行
# chain33开发文档
https://chain.33.cn/document/60
## 运行节点
	$ ./assetchain
或者
	
	$ nohup ./assetchain &

## 挖矿
1. 创建mining 挖矿账户
	
		$ bash wallet-init.sh
	
2. 给mining账户打币作为手续费
3. mining 账户bls绑定

		$ ./assetchain-cli pos33 blsbind

3. 给委托账户打币用来委托挖矿
4.  委托账户抵押到合约
	
		$ bash deposit.sh 委托账户 数量
5. 委托挖矿
	
		$ bash entrust.sh 挖矿账号 委托账号 委托私钥 数量

