# assetchain


资产链是一条包括区块链基础功能和上层应用的区块链平台，具有高兼容性的特点，支持区块链互操作能力，实现异构区块链的跨链互联功能，支持让云设施接管平台中原有的弹性、韧性、安全性、可观测性等大量非功能特性。

资产链基于自主研发、完全开源的区块链底层技术，拥有模块插件化的开发框架。

资产链开源的基本理念是创新、开放、自由、共享、协同、民主、绿色。通过协同开发、协同作业、协作生产，开源加速了各行业的技术传播和创新力量的激发。在开源的同时将区块链技术逐渐标准化，可大幅降低开发者、用户对于区块链的学习成本与接入成本。

## 基础设施

资产链采用自主研发的区块链底层开发平台，是一套支持共识算法、数据库、虚拟机等可插拔，且易升级的区块链架构。

资产链支持EVM（以太坊虚拟机）等通用智能合约虚拟机，支持相关的智能合约，实现了对以太坊的兼容，解决了不同厂商上传合约的兼容性问题。

同时，资产链按照节点的出块数量来发放共享券，共享券的应用场景很多，可以有效提高效率、实现实时交易。

项目使用模块插件化的区块链底层开发框架，资产链保留了核心功能，将扩展功能以插件方式加进来。基于插件的设计可以将扩展功能从系统框架中剥离出来，降低框架的复杂度，让框架更加容易实现。

同时，模块插件化的设计架构也能更便捷的搭建区块链和应用。单一的模块并不能体现出其优势，如果有不同的模块，就可以搭建出不同属性的应用，为高校、企业和各级部门的搭链和功能模块创建，提供个性化选择。这样的方式可以让上链用户将更多的精力投入到业务中，而非区块链底层技术的研发中。


## 认证

传统方式中，用户的注册和身份管理完全依赖于单一中心的注册机构。随着分布式账本技术的出现，分布式多中心的身份注册、标识和管理成为可能。

认证可以实现对个人,单位或企业主体身份的身份认证, 并将个人,单位或企业同一个区块链中地址绑定。

认证过程可允许职能部门或银行等机构对用户的身份进行背书，确保真实性。在此基础上，增加数字签名之后即可当做一种数字证书。具体实施可以和能提供CA能力的公司联合对用户的身份进行认证. 

使用类似的原理, 可以由政府职能部门将单位或企业的已经进行了认证的信息放在区块链上, 并将单位或企业和公钥相绑定。 单位或企业掌握对应的私钥, 可以在资产链上进行交易。 

用户在完成主体身份认证后, 可以开展很多其他的应用, 如下面介绍的数据资产的确权和交易, 完成资产的溯源等 。 区块链上的发布交易受私钥控制，只有持有私钥的用户或企业或单位才有权修改。

上链用户将根据其信用评级、资质评定等通过一级审批和二级审批。审批后用户将被发放数字身份和可验证声明签发过程以及身份认证。身份认证将通过哈希值加密保存在服务器，监管方将保存上链用户密钥，确保监管力度，防止审批混乱和审批不当。同时，可根据级别，选择不同的监管力度级别和加密保存级别。

## 资产的编码确权和交易

为解决资产编码不统一的问题, 资产编码服务将按照标准(DB33/T 2227.1—2019)进行实施, 提供全局的资产编码服务。资产编码作为资产本体关键字，用于访问存储在计算机文件中资产的特性以及记录资产的移动。资产将包括 实物资产, 软件资产和数据资产。

资产主体将所有权数据上链登记，使用 NFT确权信息在区块链上永久保存，以证明该主体对资产的所有权, 以便资产进行可能的交易和流通。在具体实施上, 登记的NFT确权资产维度不宜过大，应作为本资产主体进行交易的可验证、可定价、可核算的最小单位，便于在交易过程中，对确权数据进行确认。同时，也可以作为后续定价、交易以及评价的最小单元进行管理。

针对数据资产, 登记数据资产应包含数据的摘要信息，在不泄漏数据隐私的前提下，保证可以对确权的数据完整性和一致性进行验证，确保不会出现与其他确权数据混淆的情况。数据资产确权码记录标识数据主体信息、签名信息、数据摘要信息、时间戳信息组合的唯一性，便于后续数据交易过程中，对该数据确权登记信息的使用。

资产在登记确权和交易信息, 需要能进行溯源. 资产链于已登记的资产NFT确权记录和交易记录能支持方便、快速的检索处理。

## 存证

传统企业一般采用中心化的记账方式，在这种情况下，数据容易遭到黑客、内部人员的无痕篡改，不利于企业的数据安全。区块链为企业提供了一种分布式记账方式。上链数据可增不可减，数据、操作、资产、文件无法伪造和否认。数据修正公开透明、留痕，链上数据均生成唯一身份哈希，事件发展、业务流程、资产转移等都与相关数据哈希绑定，并通过时间戳前后排序，可追溯可验真。信息经私钥签名后，用户身份确定，而通过分布式校验，可确保其无法假冒和被盗用。同时数据加密上链，用户隐私得到保障。

依托于区块链的不可篡改性质，可以构建比传统方式更有可信度的存证溯源系统核心系统。不同的存证溯源业务都可以在存证溯源核心系统基础上进行二次开发。为存证溯源业务开发提供一套可配置的基础组件（如认证服务，组织管理，权限管理等），快速完成业务开发成为一种可能。

## 共享券

区块链技术通过共享券让每个创造价值的人分享价值，以提高每个参与者的动力。区块链技术推动人类群体协作可以总结为两个方面：一方面是降低信任成本，另一方面是提高参与动力。

共享券合约提供的功能可以分为两组:
 1. 发行相关: 预创建, 审核发行, 增发, 销毁；
 2. 流转相关: 转账, 转账到合约, 提款回合约。

为了保证共享券的真实有效，需要配置超级管理员，通过审核员和审核名单进行人工审核，才可以进行交易。


##  钱包

区块链钱包（Block Chain Wallet）是密钥的管理工具。

钱包中包含成对的私钥和公钥，用户用私钥来签名交易，从而证明该用户拥有交易的输出权。    

输出的交易信息则存储在区块链中。
利用钱包功能，用户可放心在其中保管自己的福利券、共享券等重要内容。钱包里内的变化可溯源。

##  跨链桥

随着区块链技术的迅速发展，各种公有链、联盟链、私有链层出不穷，区块链之间的互通性极大程度的限制了区块链的应用空间。

跨链就是实现一个链到另一个链的通信协议，使原本存储在特定区块链上的资产可以转换成为另一条链上的资产，从而实现价值的流通。当然也可以将其理解为不同资产持有人之间的一种兑换行为。这个过程实际并不改变每条区块链上的价值总额，但这个过程将映射在不同的链上进行保存，不可篡改。通过跨链桥服务，可以保证资产链的兼容性。

无论是公有链还是私有链，跨链技术都是实现价值互联网的关键。它是把区块链从分散的价值孤岛中拯救出来的良药，也是区块链向外拓展和连接的桥梁。

##  IM服务
现有集中存储服务的中心化IM产品都伴随着严重的信息泄露风险和隐私安全问题，对去中心化IM服务的需求应运而生。

IM服务使用的是区块链技术，依托于节点进行建设，去除集中存储服务，从而减少了信息泄露风险。聊天文字、语音私聊、图片、文件等信息加密传输，和不同的好友实现在不同的服务器上交流传输，信息内容通过用户私钥解密才可以查看。通讯录好友关系全部加密上链，保证第三方无法窃取其他用户信息，进而保护用户信息安全，减少聊天隐私泄露的隐患。通过手机验证与人脸识别等多重验证方式，实现聊天双方的实名认证。

而作为企业政府级别的IM服务，在此基础上加入监管部门和自主服务器选择，管理员可以对该服务器下的账号进行审核及封号操作，也可以对群聊人数设定限制，进一步加强风控管理能力。和传统中心化聊天不同，这样的去中心化聊天具有高度安全性，并可提高用户自主性。服务器节点的维护支持工作将由杭州复杂美科技有限公司负责。监管部门保存密钥，聊天信息加密保存在自主选择的服务器上，通过两者分开保管，进一步保证了去中心化聊天的合法性和安全性，提高用户对于产品的信任。

## CA证书

作为密码技术的集大成解决方案，PKI/CA的解决方案的完整性和优越性是普遍的一个共识。非对称算法解决了密钥传递的问题，对称算法提高加密运算效率，哈希算法解决了完整性验证的问题的同时提高了数字效率，CA作为第三方为公钥的所有者背书，解决公钥的持有者证明问题，形成一个解决信息安全机密性、完整性、抗抵赖的完整解决方案。

CA作为具备电子认证资格第三方公信力服务机构，天然具备账户托管的优势，可以实现智能合约交易的托管中心。分布式记账模式的建立，解决CA中心的中心化的陷阱问题。

在区块链的具体实现中，电子认证联盟区块链可以结合国家密码管理局国产算法的优势，形成国家自主可控、可监管、公信力的联盟区块链，在区块链的竞争中实现巨大的先发优势。

## 应用
根据服务层已有的功能，可以直接插拔下列应用模块插件，或将功能接入已有的资产云产品中。

### 溯源

传统市场上，假冒、劣质商品盛行，资产管理难以追溯，消费者难以分辨，相关工作人员工作量繁杂重复，效率低下。与此同时，传统的防伪技术过于专业，不易仿冒的同时也难以鉴别。互联网行业的快速发展，更增加了各行各业溯源的难度，也增加了溯源需求，因此统一管理十分重要。

#### 资产溯源
（一）无形资产溯源
NFT即非同质代币，本质上是数字内容的代币化。 数字内容本身具有独特性，NFT也是独一无二的。 
无形资产可以通过NFT进行发布，将其价值数字化，这可以在保留其价值的同时，明确其所有权，并且实现可溯源，不可更改以及透明灵活。这解决了无形资产在传统市场上“流动性差”、“作者追续权/版权”的问题。 目前，NFT已经为传统艺术品市场带来了无法想象的经济效益。

（二）有形资产溯源
通过将有形资产和共享券等标签相结合，确保有形资产可溯源，提高资产信息的及时性和准确性，明确有形资产的价值、数量、位置、状态、所有人信息等，满足多种应用场景。

#### 商品/农产品溯源

一物一码标识，全过程流转的信息上链存证，不可篡改，确保商品/农产品的唯一性。用户可轻松查询商品的整个生命周期。

参与资产链的用户，在流通环节之间可以实现信息数据透明、及时、高效地共享，有效降低成本。

商品从生产到销售，各环节将信息签名上链存证，一旦出现纠纷，可以快速锁定出现问题的环节。

通过区块链技术链接生产、运输、仓储、交易等全过程，形成具有信用价值的数据链，为生产者和消费者提供可信的信用认证。

#### 商城/福利券

区块链社交电商将区块链与分布式社交、数字资产相结合，构建集成了商品数字化+商品溯源+区块链提货券+数字资产交易模块的数字商务全新业态。

平台上每一个产品、每一笔交易都会生成独一无二的哈希值，不能删减、不能篡改交易平台的数据为商家和用户的真实交易奠定基础。在分布式社交上进行价值和信息的传递，有利于增加商家和用户之间的黏性，使商品推广流转。

区块链社交电商实现了去中心化的管理，免除了平台收取的额外费用，商户的运营成本进一步下降，小微企业也能谋求更广阔的发展机会。分布的社交又增加了商品推广途径，商家不需要持续增加运营成本，也能有效提高商品推广度。

区块链电商将区块链和供应链、销售链相结合，实现资产数字化。商品可按需拆分流转，通过社交网络转让赠送。双方不需要实时买卖收货，线下提货也增加了商品的流转途径，有助于构建绿色物流，开辟消费新时尚，搭建数字经济生态，真正用产品打造公平的交易平台。

通过区块链浏览器可验证产品真伪，溯源其生产、销售流程，保障商品安全和质量。平台提供多方式验证链上数据，实现查询便捷，溯源可靠。

基于部署在区块链上的可自动执行智能合约，达到平台自治的目的。整个分布式系统中的所有节点能够在去中心化的环境中，相互信任，实现安全交换数据，无需向中心管理机构求证数据的真实性。

### 电子合同/文件存储

区块链本身具备的存证溯源的功能，有利于实现电子合同和文件存储的存证和溯源。

电子合同可以通过身份认证（数字化签名认证打包加密形成哈希值），自动执行的智能合约以及去中心化来保证内容的真实性。

文件可将文本内容加密，形成独一无二的哈希值，任何变动都会被节点记录，实现真正的去中心化，透明灵活，提高工作效率，保证合同文件的真实可信。

### IM应用模块

目前市面上已有的IM应用大多避免不了平台信息泄露，安全性不足，用户连接成本过高，内容创造者无法得到应有的收益等问题。

区块链IM即时通讯工具及其相应功能模块利用其去中心化、公开透明，让每个用户均可参与数据库记录的特点，保证了聊天通讯内容的安全性。内容创造者通过其发布内容时间等明确不可更改的记录，可以有效避免被抄袭后难以维护自身权益的问题。同时，给予服务层的身份认证功能，可以保证IM应用模块的安全可控。

IM应用模块包括：人员架构，客户管理，进销存，OKR，评分考核，协同日程等。扩展IM功能，通过集成政企服务中常用工具网站等应用，可以方便用户快速找到入口，真正实现快速迭代，提高扩展能力。

### 其余各类资产应用

数字资产的应用经过几年发展，已有长足进展，但仍存在一些根本性的问题，在技术层面，围绕数据的“可用”、“可享”、“可管”、“可信”等问题更为突出。
作为新兴技术，区块链在诸多资产应用领域具有较大应用潜力。在基础设施方面，区块链可以实现赋能信息基础设施、智慧交通、能源电力等领域，提升城市管理的智能化、精准化水平。在数据资源方面，区块链有望打破原有数据流通共享壁垒，提供高质量数据共享保障，提升数据管控能力，提高数据安全保护能力。在智能应用方面，区块链将围绕惠民服务、精准治理、生态宜居、产业经济等应用场景，催生新型资产应用服务。

