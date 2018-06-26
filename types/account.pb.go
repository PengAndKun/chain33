// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	account.proto
	blockchain.proto
	common.proto
	config.proto
	db.proto
	evmcontract.proto
	executor.proto
	executorTrade.proto
	p2p.proto
	pbft.proto
	rpc.proto
	statistic.proto
	transaction.proto
	wallet.proto

It has these top-level messages:
	Account
	ReceiptExecAccountTransfer
	ReceiptAccountTransfer
	ReqBalance
	Accounts
	ReqTokenBalance
	ReqAccountTokenAssets
	TokenAsset
	ReplyAccountTokenAssets
	Header
	Block
	Blocks
	BlockPid
	BlockDetails
	Headers
	HeadersPid
	BlockOverview
	BlockDetail
	Receipts
	PrivacyKV
	PrivacyKVToken
	ReceiptsAndPrivacyKV
	ReceiptCheckTxList
	ChainStatus
	ReqBlocks
	MempoolSize
	ReplyBlockHeight
	BlockBody
	IsCaughtUp
	IsNtpClockSync
	BlockChainQuery
	Reply
	ReqString
	ReplyString
	ReplyStrings
	ReqInt
	Int64
	ReqHash
	ReplyHash
	ReqNil
	ReqHashes
	ReplyHashes
	KeyValue
	TxHash
	Config
	Log
	MemPool
	Consensus
	Wallet
	Store
	BlockChain
	P2P
	Rpc
	Exec
	LeafNode
	InnerNode
	MAVLProof
	StoreNode
	LocalDBSet
	LocalDBList
	LocalDBGet
	LocalReplyValue
	StoreSet
	StoreSetWithSync
	StoreGet
	StoreReplyValue
	EVMContractObject
	EVMContractData
	EVMContractState
	EVMContractAction
	ReceiptEVMContract
	EVMContractDataCmd
	EVMContractStateCmd
	ReceiptEVMContractCmd
	CheckEVMAddrReq
	CheckEVMAddrResp
	EstimateEVMGasReq
	EstimateEVMGasResp
	EvmDebugReq
	EvmDebugResp
	Genesis
	CoinsAction
	CoinsGenesis
	CoinsTransfer
	CoinsTransferToExec
	CoinsWithdraw
	Hashlock
	HashlockAction
	HashlockLock
	HashlockUnlock
	HashlockSend
	HashRecv
	Hashlockquery
	Ticket
	TicketAction
	TicketMiner
	TicketBind
	TicketOpen
	TicketGenesis
	TicketClose
	TicketList
	TicketInfos
	ReplyTicketList
	ReplyWalletTickets
	ReceiptTicket
	ReceiptTicketBind
	ExecTxList
	Query
	Norm
	NormAction
	NormPut
	RetrievePara
	Retrieve
	RetrieveAction
	BackupRetrieve
	PreRetrieve
	PerformRetrieve
	CancelRetrieve
	ReqRetrieveInfo
	RetrieveQuery
	TokenAction
	TokenPreCreate
	TokenFinishCreate
	TokenRevokeCreate
	Token
	ReqTokens
	ReplyTokens
	ReceiptToken
	TokenRecv
	ReplyAddrRecvForTokens
	ArrayConfig
	StringConfig
	Int32Config
	ConfigItem
	ModifyConfig
	ManageAction
	ReceiptConfig
	ReplyConfig
	PrivacyAction
	Public2Privacy
	Privacy2Privacy
	Privacy2Public
	UTXOGlobalIndex
	KeyInput
	PrivacyInput
	KeyOutput
	PrivacyOutput
	GroupUTXOGlobalIndex
	LocalUTXOItem
	ReqUTXOPubKeys
	PublicKeyData
	GroupUTXOPubKey
	ResUTXOPubKeys
	ReqPrivacyToken
	AmountDetail
	ReplyPrivacyAmounts
	ReplyUTXOsOfAmount
	ReceiptPrivacyOutput
	AmountsOfUTXO
	TokenNamesOfUTXO
	Trade
	TradeForSell
	TradeForBuy
	TradeForRevokeSell
	TradeForBuyLimit
	TradeForSellMarket
	TradeForRevokeBuy
	SellOrder
	BuyLimitOrder
	ReceiptBuyBase
	ReceiptSellBase
	ReceiptTradeBuyMarket
	ReceiptTradeBuyLimit
	ReceiptTradeBuyRevoke
	ReceiptTradeSell
	ReceiptSellMarket
	ReceiptTradeRevoke
	ReqAddrTokens
	ReqTokenSellOrder
	ReqTokenBuyOrder
	ReplyBuyOrder
	ReplySellOrder
	ReplySellOrders
	ReplyBuyOrders
	ReplyTradeOrder
	ReplyTradeOrders
	P2PGetPeerInfo
	P2PPeerInfo
	P2PVersion
	P2PVerAck
	P2PPing
	P2PPong
	P2PGetAddr
	P2PAddr
	P2PAddrList
	P2PExternalInfo
	P2PGetBlocks
	P2PGetMempool
	P2PInv
	Inventory
	P2PGetData
	P2PTx
	P2PBlock
	BroadCastData
	P2PGetHeaders
	P2PHeaders
	InvData
	InvDatas
	Peer
	PeerList
	NodeNetInfo
	Operation
	Checkpoint
	Entry
	ViewChange
	Summary
	Result
	Request
	RequestClient
	RequestPrePrepare
	RequestPrepare
	RequestCommit
	RequestCheckpoint
	RequestViewChange
	RequestAck
	RequestNewView
	ClientReply
	TotalFee
	ReqGetTotalCoins
	ReplyGetTotalCoins
	IterateRangeByStateHash
	TicketStatistic
	TicketMinerInfo
	CreateTx
	UnsignTx
	SignedTx
	Transaction
	Transactions
	RingSignature
	RingSignatureItem
	Signature
	AddrOverview
	ReqAddr
	ReqPrivacy
	HexTx
	ReplyTxInfo
	ReqTxList
	ReplyTxList
	TxHashList
	ReplyTxInfos
	ReceiptLog
	Receipt
	ReceiptData
	TxResult
	TransactionDetail
	TransactionDetails
	ReqAddrs
	WalletTxDetail
	WalletTxDetails
	WalletAccountStore
	WalletAccountPrivacy
	WalletPwHash
	WalletStatus
	WalletAccounts
	WalletAccount
	WalletUnLock
	GenSeedLang
	GetSeedByPw
	SaveSeedByPw
	ReplySeed
	ReqWalletSetPasswd
	ReqNewAccount
	MinerFlag
	ReqWalletTransactionList
	ReqWalletImportPrivKey
	ReqWalletSendToAddress
	ReqWalletSetFee
	ReqWalletSetLabel
	ReqWalletMergeBalance
	ReqStr
	ReplyStr
	ReqTokenPreCreate
	ReqTokenFinishCreate
	ReqTokenRevokeCreate
	ReqSellToken
	ReqBuyToken
	ReqRevokeSell
	ReqModifyConfig
	ReqSignRawTx
	ReplySignRawTx
	ReportErrEvent
	Int32
	ReqPub2Pri
	ReqPri2Pri
	ReqPri2Pub
	ReqCreateUTXOs
	ReplyPrivacyPkPair
	ReqPrivBal4AddrToken
	ReplyPrivacyBalance
	PrivacyDBStore
	UTXO
	UTXOHaveTxHash
	UTXOs
	UTXOHaveTxHashs
	ReqUTXOGlobalIndex
	UTXOBasic
	UTXOIndex4Amount
	ResUTXOGlobalIndex
	FTXOsSTXOsInOneTx
	ReqCreateTransaction
	RealKeyInput
	UTXOBasics
	CreateTransactionCache
	ReqCacheTxList
	ReplyCacheTxList
	ReqPrivacyAccount
	ReqPPrivacyAccount
	ReplyPrivacyAccount
	ReqCreateCacheTxKey
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Account 的信息
type Account struct {
	// coins标识，目前只有0 一个值
	Currency int32 `protobuf:"varint,1,opt,name=currency" json:"currency,omitempty"`
	// 账户可用余额
	Balance int64 `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	// 账户冻结余额
	Frozen int64 `protobuf:"varint,3,opt,name=frozen" json:"frozen,omitempty"`
	// 账户的地址
	Addr string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetCurrency() int32 {
	if m != nil {
		return m.Currency
	}
	return 0
}

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// 账户余额改变的一个交易回报（合约内）
type ReceiptExecAccountTransfer struct {
	// 合约地址
	ExecAddr string `protobuf:"bytes,1,opt,name=execAddr" json:"execAddr,omitempty"`
	// 转移前
	Prev *Account `protobuf:"bytes,2,opt,name=prev" json:"prev,omitempty"`
	// 转移后
	Current *Account `protobuf:"bytes,3,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptExecAccountTransfer) Reset()                    { *m = ReceiptExecAccountTransfer{} }
func (m *ReceiptExecAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptExecAccountTransfer) ProtoMessage()               {}
func (*ReceiptExecAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReceiptExecAccountTransfer) GetExecAddr() string {
	if m != nil {
		return m.ExecAddr
	}
	return ""
}

func (m *ReceiptExecAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptExecAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 账户余额改变的一个交易回报（coins内）
type ReceiptAccountTransfer struct {
	// 转移前
	Prev *Account `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	// 转移后
	Current *Account `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptAccountTransfer) Reset()                    { *m = ReceiptAccountTransfer{} }
func (m *ReceiptAccountTransfer) String() string            { return proto.CompactTextString(m) }
func (*ReceiptAccountTransfer) ProtoMessage()               {}
func (*ReceiptAccountTransfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReceiptAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

// 查询一个地址列表在某个执行器中余额
type ReqBalance struct {
	// 地址列表
	Addresses []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	// 执行器名称
	Execer string `protobuf:"bytes,2,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqBalance) Reset()                    { *m = ReqBalance{} }
func (m *ReqBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqBalance) ProtoMessage()               {}
func (*ReqBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReqBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

// Account 的列表
type Accounts struct {
	Acc []*Account `protobuf:"bytes,1,rep,name=acc" json:"acc,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Accounts) GetAcc() []*Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

type ReqTokenBalance struct {
	Addresses   []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	TokenSymbol string   `protobuf:"bytes,2,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
	Execer      string   `protobuf:"bytes,3,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqTokenBalance) Reset()                    { *m = ReqTokenBalance{} }
func (m *ReqTokenBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenBalance) ProtoMessage()               {}
func (*ReqTokenBalance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReqTokenBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqTokenBalance) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *ReqTokenBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

type ReqAccountTokenAssets struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Execer  string `protobuf:"bytes,2,opt,name=execer" json:"execer,omitempty"`
}

func (m *ReqAccountTokenAssets) Reset()                    { *m = ReqAccountTokenAssets{} }
func (m *ReqAccountTokenAssets) String() string            { return proto.CompactTextString(m) }
func (*ReqAccountTokenAssets) ProtoMessage()               {}
func (*ReqAccountTokenAssets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReqAccountTokenAssets) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ReqAccountTokenAssets) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

type TokenAsset struct {
	Symbol  string   `protobuf:"bytes,1,opt,name=symbol" json:"symbol,omitempty"`
	Account *Account `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
}

func (m *TokenAsset) Reset()                    { *m = TokenAsset{} }
func (m *TokenAsset) String() string            { return proto.CompactTextString(m) }
func (*TokenAsset) ProtoMessage()               {}
func (*TokenAsset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TokenAsset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TokenAsset) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ReplyAccountTokenAssets struct {
	TokenAssets []*TokenAsset `protobuf:"bytes,1,rep,name=tokenAssets" json:"tokenAssets,omitempty"`
}

func (m *ReplyAccountTokenAssets) Reset()                    { *m = ReplyAccountTokenAssets{} }
func (m *ReplyAccountTokenAssets) String() string            { return proto.CompactTextString(m) }
func (*ReplyAccountTokenAssets) ProtoMessage()               {}
func (*ReplyAccountTokenAssets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReplyAccountTokenAssets) GetTokenAssets() []*TokenAsset {
	if m != nil {
		return m.TokenAssets
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "types.Account")
	proto.RegisterType((*ReceiptExecAccountTransfer)(nil), "types.ReceiptExecAccountTransfer")
	proto.RegisterType((*ReceiptAccountTransfer)(nil), "types.ReceiptAccountTransfer")
	proto.RegisterType((*ReqBalance)(nil), "types.ReqBalance")
	proto.RegisterType((*Accounts)(nil), "types.Accounts")
	proto.RegisterType((*ReqTokenBalance)(nil), "types.ReqTokenBalance")
	proto.RegisterType((*ReqAccountTokenAssets)(nil), "types.ReqAccountTokenAssets")
	proto.RegisterType((*TokenAsset)(nil), "types.TokenAsset")
	proto.RegisterType((*ReplyAccountTokenAssets)(nil), "types.ReplyAccountTokenAssets")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcb, 0x4e, 0xc2, 0x40,
	0x14, 0xcd, 0x50, 0x5e, 0xbd, 0x44, 0x8d, 0x93, 0x88, 0x0d, 0x71, 0xd1, 0xcc, 0xaa, 0x0b, 0xc3,
	0x42, 0xbe, 0x00, 0x12, 0x17, 0x6e, 0x58, 0x8c, 0xfc, 0x40, 0x19, 0x2e, 0x09, 0x01, 0xdb, 0x32,
	0x33, 0x18, 0xeb, 0x07, 0xf8, 0xdd, 0x66, 0xa6, 0xb7, 0x80, 0x22, 0x86, 0x1d, 0xe7, 0x3e, 0xce,
	0x39, 0xf7, 0x30, 0x85, 0xab, 0x54, 0xa9, 0x7c, 0x97, 0xd9, 0x61, 0xa1, 0x73, 0x9b, 0xf3, 0x96,
	0x2d, 0x0b, 0x34, 0x62, 0x0d, 0x9d, 0x71, 0x55, 0xe7, 0x03, 0xe8, 0xaa, 0x9d, 0xd6, 0x98, 0xa9,
	0x32, 0x62, 0x31, 0x4b, 0x5a, 0x72, 0x8f, 0x79, 0x04, 0x9d, 0x79, 0xba, 0x49, 0x33, 0x85, 0x51,
	0x23, 0x66, 0x49, 0x20, 0x6b, 0xc8, 0xfb, 0xd0, 0x5e, 0xea, 0xfc, 0x13, 0xb3, 0x28, 0xf0, 0x0d,
	0x42, 0x9c, 0x43, 0x33, 0x5d, 0x2c, 0x74, 0xd4, 0x8c, 0x59, 0x12, 0x4a, 0xff, 0x5b, 0x7c, 0x31,
	0x18, 0x48, 0x54, 0xb8, 0x2a, 0xec, 0xf3, 0x07, 0x2a, 0x12, 0x9e, 0xe9, 0x34, 0x33, 0x4b, 0xd4,
	0xce, 0x00, 0xba, 0xb2, 0x5b, 0x63, 0x7e, 0x6d, 0x8f, 0xb9, 0x80, 0x66, 0xa1, 0xf1, 0xdd, 0xab,
	0xf7, 0x9e, 0xae, 0x87, 0xde, 0xfd, 0x90, 0x18, 0xa4, 0xef, 0xf1, 0x04, 0x3a, 0x95, 0x61, 0xeb,
	0xbd, 0x9c, 0x8e, 0xd5, 0x6d, 0xb1, 0x84, 0x3e, 0xf9, 0xf8, 0xed, 0xa1, 0xd6, 0x61, 0x97, 0xe9,
	0x34, 0xfe, 0xd7, 0x99, 0x00, 0x48, 0xdc, 0x4e, 0x28, 0xaa, 0x07, 0x08, 0x5d, 0x0c, 0x68, 0x0c,
	0x9a, 0x88, 0xc5, 0x41, 0x12, 0xca, 0x43, 0xc1, 0x05, 0xe9, 0xae, 0x45, 0xed, 0x49, 0x43, 0x49,
	0x48, 0x3c, 0x42, 0x97, 0x78, 0x0d, 0x8f, 0x21, 0x48, 0x95, 0xf2, 0xbb, 0xa7, 0xaa, 0xae, 0x25,
	0x56, 0x70, 0x23, 0x71, 0x3b, 0xcb, 0xd7, 0x98, 0x5d, 0x26, 0x1b, 0x43, 0xcf, 0xba, 0xe9, 0xd7,
	0xf2, 0x6d, 0x9e, 0x6f, 0x48, 0xfb, 0xb8, 0x74, 0x64, 0x2c, 0xf8, 0x61, 0xec, 0x05, 0xee, 0x24,
	0x6e, 0xeb, 0x00, 0xdd, 0xc2, 0xd8, 0x18, 0xb4, 0xc6, 0x3d, 0x16, 0xe2, 0xa7, 0xbf, 0xb1, 0x86,
	0x67, 0x6f, 0x9c, 0x02, 0x1c, 0x08, 0xdc, 0x94, 0xa9, 0xdc, 0x54, 0xeb, 0x84, 0x5c, 0xee, 0xf4,
	0x86, 0xcf, 0xe5, 0x4e, 0x6d, 0x31, 0x85, 0x7b, 0x89, 0xc5, 0xa6, 0xfc, 0xc3, 0xdc, 0x88, 0xee,
	0xad, 0x20, 0x45, 0x79, 0x4b, 0x44, 0x87, 0x41, 0x79, 0x3c, 0x35, 0x6f, 0xfb, 0x6f, 0x66, 0xf4,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x89, 0x5a, 0xee, 0x69, 0x44, 0x03, 0x00, 0x00,
}
