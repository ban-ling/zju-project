# zju-project

# 机制：TruSD信任框架+反馈增强信任（DFET）计算

（我认为反馈机制是相当精彩的做法，以及这次的题目没有关于效率的要求）

## 机制1解释

时间窗（Dt）用于衡量满意和不满意服务的数量。时间窗口为在时间窗口内进行计数。经过一段时间后，窗口向右滑动一个
时间帧，从而放弃在最后一个时间帧中显示的服务行为。因此
，随着时间的推移，窗口忘记了旧框架的体验，而是增加了新
时间框架的体验。根据信任管理场景，窗口长度可以缩短或更
长。图中显示了该时间窗方案的一个示例场景。5.三个时间单
位ðDu1；杜2和杜3给出了Þ，每个Þ包含几个时间框架。因此，
每个时间单位对应于一个评价矩阵。使用等式(5)我们可以推导
三组ssdD1; D2; D3.从三组ssd中得到的(D1; D2; D3)，我们
可以根据给定的时间窗口Dt来计算正或负服务行为的实例数。
例如，如果我们设置Dt = 20，我们推导出UðPiÞ = 5;SðPiÞ =
15.如果我们设置Dt = 32，我们推导出UðPiÞ = 7;SðPiÞ = 25. 在确定了UðP的值后iÞ和SðPiÞ在时间窗口内，DTD Dnð P i资
源P的Þi位于0和1之间的值被定义为：
DnðPiÞ =
(
PiÞ
!
;
(6)
其中，入是一个正整数，和UðPi Þ 0.SðPiÞ是资源P的令人满
意的服务的总数i和UðPiÞ是资源P中不令人满意的服务的总数i
在时间窗口Dt期间。在特殊情况下，如果是SðPi Þ0和UðPiÞ=为
0，我们设置为DnðPiÞ = 1.如果UðPiÞ+ SðPiÞ=0，这意味着资
源P之间没有交互作用i和用户在时间窗口Dt。我们设置DnðPiÞ
= 0:5.
DTD是基于直接的交互作用而建立的
在用户和服务资源之间，并根据令人满意和不令人满意的服务
的数量进行评估

PiÞ 随着
不满意服务数量的增加，迅速接近0，这表明了严格的惩罚特征
DFET用于此类服务。DFET的严格惩罚特性可以有效防止来自累
积可信度较高的恶意节点的突然攻击。入的值决定了惩罚的程
度。入值越小，表示惩罚值越大。为了减轻可能由网络流量问
题导致的一些错误的服务声明的影响，我们建议为入设置一个
稍大一点的值，比如入=3。确定一种有效的方法来确定这个新
参数值得我们今后进一步研究。
图6直观地描述了DTD的趋势iÞ和UðPiÞ（其中为入= 3）。从
无花果。6，我们发现DTD的值随着UðP的增加而迅速趋近于0iÞ . 例如，信任值为0.9,1不满意10满意服务，信任值为0.6614,2不
满意10满意服务，信任值为0.4498,4不满意10满意服务。我们
发现，当不令人满意的服务数量达到4个时，DTD的值就小于0.5 。在这种情况下，尽管令人满意的服务的数量远远大于不令人
满意的服务的数量，但来自该资源的服务将被认为是不可信的
。图中的案例。6直观上表示DFET对不满意服务的严格惩罚特征
，即信任容易丧失，难以恢复。


## 机制1
![img_2.png](img_2.png)

(DFET模式是一种混合信任评估方法[1]，它超越
了最近关注服务站点[4]、[5]、[7]、[8]、[17]的单边信任因
素的信任方案。为了全面评价云服务，基于实时QoS证据、用户
反馈和监控器反馈这两个互补因素计算可信度，可以在一定程
度上解决单因素信任计算模型造成的缺陷问题。这种扩展的信
任评估模式更符合信任关系的基本属性，因此，它更符合用户
[1]、[15]的期望。根据定义2，DFET模式也是一种基于TTP [1]
的方法。云监视器充当TTP，它包括许多已注册的服务。与其他
信任模式，如在用户方面采用信任机制相比，这种基于TTP的信
任模式在监视器上采用信任机制，有效地减少用户的负担。)


## 机制2
我们假设有一个名为T的信任函数，它根据上下 文收集一些参数，并返回一个信任值tvP对于提供程序节点P: t u P←T（参
数1参数2, …).为了找出DHT层中提供者的引用持有人，计算了以下哈希操 作：h(节点P| | i ) . 哈希值计算每个i，其中i∈{0,1,2，……，c}。（c是每
个设备的参考支架数量）。每个计算出的散列值都显示了其中一个参考支架装置提供程序。每个设备都有自己的内部信任列表TL，其中包含它之前通信过的任何节 点的信任值。在请求者从节点接收到该服务后P，它将增加电视P到其内部信
任列表TL。然后，对于进一步的通信，请求者节点将使用此列表。 由于这项工作的目的是在从提供者节点请求服务之前提供一个决策，因
此请求者节点需要征求参考持有人节点的意见。根据他们的意见，计算出一 个聚合的信任值ATP。然后利用一种决策算法
将ATP与阈值e进行比较。然后就决定了它是否可靠。

## 论证2
我们的框架位于一个基于分布式散列表（DHT）的P2P网络之上。我们假 设DHT将保留框架中每个设备的信任值。我们没有指定DHT技术，它可以是任何类型，如CAN，Chord等。操作它有一个动态的
结构，节点可以随时连接和离开网络。由于我们为P2P网络中的服务提供提供了一个信任框架，因此一个设备将请求一个服务，而另一个设备将提供该服务。只要它是一个P2P网络，所有的设备都可以同时有不同的角色。在我们的系统模型中，设备可以具有以下角色：提供者、请求者和参考持有人（RH）。提供者是将查询其信任的设备，
而请求者是希望与提供者设备通信的设备，如果该设备是真实的。引用持有
人是持有提供者节点的信任值的节点。在我们的框架中，我们提供了一种信任机制，旨在为新的通信提供准确的信息，特别是在新通信中，一个节点希
望从另一个之前没有通信的节点接收服务。其主要目标是对以前没有接触过
的设备做出准确的决定。在某些情况下，信任一个未知节点可能是毁灭性的。其中一个例子如图所示。3.假设在智能家居中，有一个温度传感器需要从
互联网上进行更新。然后，传感器将具有请求者的角色，并命名为R。它首先
从一个索引表（不在我们的范围内）学习供应商列表PL，并学习智能手表和 摄像头可以提供这项服务（第0点）。然而，传感器没有与任何一个通信。假设它从提供商列表中选择了智能手表作为提供商P。然后，R首先需要与智能
手表的参考持有人进行交流，并听取他们的意见。为了到达参考持有人，请求者需要通过DHT地址的查找功能来学习参考持有人的通信地址X=查找(X)（点1），其中X是通过应用散列计算出的对象，如下：h(节点P| | i ) . 然后，R通过参考查询协议（点2）与所有引用持有者进行通信。然后，它计算一个聚 合的信任值，将此值与一个阈值进行比较，并决定是否信任。如果R决定信任，则需要通过DHT地址的查找功能来了解提供者P的通信地址P=查找(P)（点3 ）。然后，它通过服务提供协议（第4点）请求服务。当R到达服务后，它会 对服务进行评估，并提供一个反馈报告。然后通过反馈报告协议（第5点）将报告发送给所有参考资料持有人。如果智能手表是恶意的，它可以发送一个 恶意软件，这将是毁灭性的，不可能恢复。因此，我们提出了对这种陌生人 通信的信任框架，以减少这种破坏的可能性。 我们假设攻击者的主要目标是获得对网络的控制，以提供尽可能多的恶 意服务。他的目标是操纵良性节点来信任，并要求恶意节点提供服务。我们允许对手从N个设备中捕获k个
并设法将大多数良性节点指向恶意节点。如果一个
攻击者设法成为一个参考持有者，他可以提供积极的
参考了他的合作者，而负面的良性节点则是为了操纵。攻击者还可以损害合
法的设备，并阻止它们提供服务的能力。我们认为他们可以合作，他们之间
有完美的即时交流。他们可以控制自己的行为，他们可以决定撒谎，但他们
无法接触到其他设备的信息。
受破坏的节点可以提供开关攻击、选择性攻击、坏口齿攻击和选票填充
攻击[15]。开关攻击利用了更重视最近关于信任系统的建议的特性。攻击者
可以在一段时间内表现良好，并重新获得信任。我们的TruSD模型通过给予以
前的信任值更多的权重来处理这个问题。当恶意节点提供不诚实的建议来降
低诚实方的可信度时，就会发生坏话攻击。当恶意设备具有请求者或引用持
有人角色时，它们会对良性设备发送负面反馈。类似地，选票填充攻击是指
操纵被盗设备的声誉，以提高可信度。在我们的系统中，恶意设备正在对被
捕获的节点发送正反馈。TruSD可以抵抗这些攻击，因为请求者将自己的经验
与参考持有者的建议进行比较。然后，系统可以很容易地注意到恶意节点操
作的试验。此外，选择性攻击还提供了伪装概率，即当恶意节点具有提供者
角色时，它们在请求之间的表现交替不佳。这种攻击的影响在TruSD中减弱了
，因为信任决策不仅考虑一个，而且考虑几个参考持有人的建议提供。我们
不考虑干扰和DDoS攻击。



## 其他

### 提高信誉机制：
1.身份验证：用户和服务提供者在注册之前需要进行身份验证，以确保其唯一性和真实性。可以使用数字身份证明、第三方认证机构或网络身份验证协议来完成验证过程。身份验证信息将被加密存储在区块链上，以确保隐私和数据安全。（通过强制要求用户和服务提供者进行身份验证，可以减少虚假身份和欺诈行为的风险。使用数字身份证明、第三方认证机构或网络身份验证协议能够确保参与者的唯一性和真实性，并将身份验证信息存储在区块链上以确保数据的安全性。）
2.用户评价系统：用户可以对完成的服务进行评价和评论，评价可以包括服务质量、交付时间、沟通能力等方面。这些评价将被记录在区块链上，并与服务提供者的身份关联起来。用户评价可以作为其他用户选择服务提供者的参考依据，有助于建立可信度。（用户评价和评论可以提供对服务质量、交付时间、沟通能力等方面的反馈信息。这些评价将被记录在区块链上并与服务提供者的身份关联起来，这种透明性和公开的评价系统可以降低恶意行为的风险，同时激励服务提供者提供良好的服务质量。）
3.服务履约智能合约：使用智能合约规定服务的履约条件和标准，例如交付时间、服务质量等。当服务提供者按照合约规定提供服务并得到用户确认后，其信誉度将得到提高。如果服务提供者未能履约，用户可以提出争议并启动仲裁流程，由信誉评估委员会进行审查。去中心化信誉评估委员会：（使用智能合约规定服务的履约条件和标准，可以确保服务提供者按照约定的条件提供服务，这种机制可以促使服务提供者遵守合约，并为用户提供一种途径来解决争议。）
4.引入由多个独立实体组成的信誉评估委员会，用于对用户评价和争议进行审查和仲裁。该委员会由具有良好信誉度和专业背景的实体组成，在网络中发挥公正、客观和权威的作用。委员会的决策将对参与者的信誉度产生影响，以确保审查过程的公正性和可靠性。

### 降低信誉机制：
1.行为分析：通过分析用户的行为模式和活动，可以识别出与正常行为明显偏离的异常行为。例如，频繁改变身份、异常的访问模式、大量无效操作等。
2.网络流量分析：实时监控和分析网络流量，以检测异常的流量模式，如异常的数据包大小、流量峰值等。这对于检测DDoS攻击等恶意行为非常有用。
3.数据异常检测：应用机器学习和统计分析技术，利用历史数据来训练模型，以便检测出与正常数据分布明显不同的异常数据点。
发现异常后进行降低信誉。当然，对于攻击还是要处理的。对于女巫攻击，采用身份验证机制和去中心化身份认证协议，确保每个实体的唯一性。
对于日蚀攻击，采用多节点随机选择和分布式审计机制，避免一个实体控制整个网络的节点选择过程。
对于DDoS攻击，使用分布式共识算法和反DDoS措施，保护网络的正常运行。

### 高信誉奖励机制：
1.可以设计奖励机制来鼓励信誉高的实体，例如通过发放代币、优惠券或其他形式的奖励来回报他们的良好行为。
2.推荐和宣传：将信誉高的实体推荐给其他用户，提供更多的业务机会和曝光机会，从而促进其业务增长。
3.优先权和特殊待遇：给予信誉高的实体一定的优先权和特殊待遇，例如在任务分配、资源分配等方面给予优先考虑。
4.参与决策：邀请信誉高的实体参与项目决策或共治过程，使其能够在系统中发挥更大的作用。

### 低信誉惩罚机制：
1.可以设计惩罚机制来对信誉低的实体进行惩罚，例如降低其在平台上的排名、限制其业务机会或加强对其行为的监管。
2.仲裁和争议解决：通过建立信誉评估委员会或类似的机构来处理用户之间的争议，并对恶意行为进行惩罚。
3.合约解除：在履约智能合约中设定相应条款，当实体未按照合约规定提供服务时，可以解除合约并采取相应的法律措施。
4.公开曝光：对于严重的恶意行为，公开曝光其行为和信誉低下的情况，以警示其他参与者避免与其交互。
