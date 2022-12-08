package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	// "math"
	"math/rand"
	"os"
	"time"

	pt "github.com/assetcloud/assetchain/plugin/dapp/pos33/types"
	"github.com/assetcloud/chain/common"
	"github.com/assetcloud/chain/common/address"
	"github.com/assetcloud/chain/common/crypto"
	grpc "github.com/assetcloud/chain/rpc/grpcclient"
	jrpc "github.com/assetcloud/chain/rpc/jsonclient"
	rpctypes "github.com/assetcloud/chain/rpc/types"
	"github.com/assetcloud/chain/system/crypto/none"
	ctypes "github.com/assetcloud/chain/system/dapp/coins/types"
	"github.com/assetcloud/chain/types"
	_ "github.com/assetcloud/plugin/plugin"
)

//
// test auto generate tx and send to the node
//

var rootKey, noUseKey crypto.PrivKey

const ethID = pt.EthAddrID

func init() {
	rand.Seed(time.Now().UnixNano())
	rootKey = HexToPrivkey("CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944")
	noUseKey = HexToPrivkey("c207ae827ccd7dfbe2b0d24d639d131300f26f8a176cf7f3752b5b454d62ed1f")
}

var rpcURL = flag.String("u", "http://localhost:9901", "rpc url")
var grpcURL = flag.String("g", "127.0.0.1:9902", "grpc url")

var pnodes = flag.Bool("p", false, "only print node private keys")
var ini = flag.Bool("i", false, "send init tx")
var maxacc = flag.Int("a", 10000, "max account")
var maxaccF = flag.Int("m", 1000000, "max account in a file")

// var rn = flag.Int("r", 3000, "sleep in Microsecond")
var conf = flag.String("c", "assetchain.toml", "chain config file")
var useGrpc = flag.Bool("G", false, "if use grpc")
var sign = flag.Bool("s", true, "signature tx")
var accFile = flag.String("f", "acc.dat", "acc file")
var noUseTx = flag.Bool("n", false, "send no use tx")

var gClient types.ChainClient
var jClient *jrpc.JSONClient
var config *types.ChainConfig

func main() {
	flag.Parse()
	config = types.NewChainConfig(types.MergeCfg(types.ReadFile(*conf), ""))
	rand.Seed(time.Now().Unix())
	// config.EnableCheckFork(false)

	privCh := runLoadAccounts(*accFile, *maxacc)
	if *pnodes {
		return
	}
	if privCh == nil {
		log.Fatal("NO account !!!")
		return
	}

	gclient, err := grpc.NewMainChainClient(config, *grpcURL)
	if err != nil {
		log.Println(err)
	}
	gClient = gclient

	jclient, err := jrpc.NewJSONClient(*rpcURL)
	if err != nil {
		log.Println(err)
	}
	jClient = jclient

	if *ini {
		log.Println("@@@@@@@ send init txs", *ini)
		runSendInitTxs(privCh)
	} else {
		var privs []crypto.PrivKey
		for {
			priv, ok := <-privCh
			if !ok {
				break
			}
			privs = append(privs, priv)

		}
		run(privs)
	}
}

// Int is int64
type Int int64

// Marshal Int to []byte
func (i Int) Marshal() []byte {
	b := make([]byte, 16)
	n := binary.PutVarint(b, int64(i))
	return b[:n]
}

// Unmarshal []byte to Int
func (i *Int) Unmarshal(b []byte) (int, error) {
	a, n := binary.Varint(b)
	*i = Int(a)
	return n, nil
}

type pp struct {
	i int
	p crypto.PrivKey
}

// Tx is alise types.Transaction
type Tx = types.Transaction

func run(privs []crypto.PrivKey) {
	hch := make(chan int64, 10)
	ch := generateTxs(privs, hch)
	tch := time.NewTicker(time.Second * 10).C
	h, err := gClient.GetLastHeader(context.Background(), &types.ReqNil{})
	if err != nil {
		panic(err)
	}
	height := h.Height
	hch <- height
	i := 0
	max := 256
	txs := make([]*types.Transaction, max)
	ntx := 0
	sleepd := time.Millisecond * 20
	t := time.Now()
	for tx := range ch {
		select {
		case <-tch:
			if *useGrpc {
				data := &types.ReqNil{}
				h, err := gClient.GetLastHeader(context.Background(), data)
				if err != nil {
					panic(err)
				}
				height = h.Height
			} else {
				var res rpctypes.Header
				err := jClient.Call("Chain.GetLastHeader", nil, &res)
				if err != nil {
					panic(err)
				}
				height = res.Height
			}
		default:
		}
		hch <- height

		txs[i] = tx
		i++

		ntx++
		if ntx%1000 == 0 {
			if ntx%10000 == 0 {
				log.Println(ntx, "... txs sent")
				time.Sleep(time.Second * 1)
			}

			d := time.Since(t)
			if d > time.Second*1 {
				sleepd -= time.Millisecond * 10
			} else if d < time.Second*1 {
				sleepd += time.Millisecond * 10
			}

			t = time.Now()
		}

		if i == max {
			sendTxs(txs)
			time.Sleep(sleepd)
			i = 0
		}
	}
}

func generateTxs(privs []crypto.PrivKey, hch <-chan int64) chan *Tx {
	N := 4
	l := len(privs) - 1
	ch := make(chan *Tx, N)
	f := func() {
		for {
			i := rand.Intn(len(privs))
			signer := privs[l-i]
			if *noUseTx {
				ch <- newNoUseTx()
			} else {
				ch <- newTxWithTxHeight(signer, 1, address.PubKeyToAddr(address.DefaultID, privs[i].PubKey().Bytes()), <-hch)
			}
		}
	}
	for i := 0; i < N; i++ {
		go f()
	}
	return ch
}

func sendTxs(txs []*Tx) error {
	ts := &types.Transactions{Txs: txs}
	if *useGrpc {
		rs, err := gClient.SendTransactions(context.Background(), ts)
		if err != nil {
			return err
		}
		for _, r := range rs.ReplyList {
			if string(r.Msg) == types.ErrMemFull.Error() {
				time.Sleep(time.Second * 5)
			}
			if !r.IsOk {
				log.Println(string(r.Msg))
			}
		}
	} else {
		panic("can't go here")
		// return errors.New("not support grpc in batch send txs")
		// err = jClient.Call("Chain.SendTransaction", &rpctypes.RawParm{Data: common.ToHex(types.Encode(tx))}, &txHash)
	}
	return nil
}

// _, ok := err.(*json.InvalidUnmarshalError)
func sendTx(tx *Tx) error {
	var err error
	if *useGrpc {
		_, err = gClient.SendTransaction(context.Background(), tx)
	} else {
		var txHash string
		err = jClient.Call("Chain.SendTransaction", &rpctypes.RawParm{Data: common.ToHex(types.Encode(tx))}, &txHash)
	}
	if err != nil {
		// _, ok := err.(*json.InvalidUnmarshalError)
		// if ok {
		// 	break
		// }
		// if err == types.ErrFeeTooLow {
		// 	log.Println("@@@ rpc error: ", err, tx.From(), tx.Fee)
		// 	tx.Fee *= 2
		// 	continue
		// }
		log.Println("@@@ rpc error: ", err, common.HashHex(tx.Hash()))
		if err.Error() == types.ErrMemFull.Error() {
			time.Sleep(time.Second)
		}
	}
	return nil
}

func runSendInitTxs(privCh chan crypto.PrivKey) {
	ch := make(chan *Tx, 16)
	go runGenerateInitTxs(privCh, ch)
	max := 256
	i := 0
	txs := make([]*types.Transaction, max)
	ntx := 0
	for {
		tx, ok := <-ch
		if !ok {
			sendTxs(txs[:i])
			break
		}
		txs[i] = tx
		i++
		if i == max {
			sendTxs(txs)
			i = 0
		}
		ntx++
	}
	log.Println("send ntx:", ntx)
}

func newTxWithTxHeight(priv crypto.PrivKey, amount int64, to string, height int64) *Tx {
	act := &ctypes.CoinsAction{Value: &ctypes.CoinsAction_Transfer{Transfer: &types.AssetsTransfer{Cointoken: "AS", Amount: amount}}, Ty: ctypes.CoinsActionTransfer}
	payload := types.Encode(act)
	tx, err := types.CreateFormatTx(config, "coins", payload)
	if err != nil {
		panic(err)
	}
	tx.To = to
	tx.Expire = height + 20 + types.TxHeightFlag
	tx.ChainID = 9991
	tx.Fee = 1000000
	tx.Nonce = rand.Int63()
	if *sign {
		signID := types.EncodeSignID(types.SECP256K1, ethID)
		tx.Sign(signID, priv)
	}
	return tx
}

func RandStringBytes(n int) []byte {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-=_+=/<>!@#$%^&"
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func newNoUseTx() *Tx {
	execAddr := address.ExecAddress("user.write")
	//构造存证交易
	tx := &types.Transaction{Execer: []byte("user.write")}
	tx.To = execAddr
	tx.Fee = 1e6
	tx.Nonce = time.Now().UnixNano()
	//tx.Expire = types.TxHeightFlag + txHeight
	tx.Expire = 0
	tx.Payload = RandStringBytes(32)
	//交易签名
	//tx.Sign(types.SECP256K1, priv)
	// tx.Signature = &types.Signature{Ty: none.ID, Pubkey: priv.PubKey().Bytes()}
	tx.Signature = &types.Signature{Ty: none.ID, Pubkey: noUseKey.PubKey().Bytes()}
	tx.ChainID = 9991
	return tx
}

func newTx(priv crypto.PrivKey, amount int64, to string) *Tx {
	act := &ctypes.CoinsAction{Value: &ctypes.CoinsAction_Transfer{Transfer: &types.AssetsTransfer{Cointoken: "AS", Amount: amount}}, Ty: ctypes.CoinsActionTransfer}
	payload := types.Encode(act)
	tx, err := types.CreateFormatTx(config, "coins", payload)
	if err != nil {
		panic(err)
	}
	tx.Fee = 1000000
	tx.To = to
	tx.ChainID = 9991
	tx.Nonce = rand.Int63()
	if *sign {
		signID := types.EncodeSignID(types.SECP256K1, ethID)
		tx.Sign(signID, priv)
	}
	return tx
}

//HexToPrivkey ： convert hex string to private key
func HexToPrivkey(key string) crypto.PrivKey {
	cr, err := crypto.Load(types.GetSignName("", types.SECP256K1), -1)
	if err != nil {
		panic(err)
	}
	bkey, err := common.FromHex(key)
	if err != nil {
		panic(err)
	}
	priv, err := cr.PrivKeyFromBytes(bkey)
	if err != nil {
		panic(err)
	}
	return priv
}

func runGenerateInitTxs(privCh chan crypto.PrivKey, ch chan *Tx) {
	for {
		priv, ok := <-privCh
		if !ok {
			log.Println("privCh is closed")
			close(ch)
			return
		}
		m := 100 * types.DefaultCoinPrecision
		to := address.PubKeyToAddr(ethID, priv.PubKey().Bytes())
		ch <- newTx(rootKey, m, to)
	}
}
func generateInitTxs(n int, privs []crypto.PrivKey, ch chan *Tx, done chan struct{}) {
	for _, priv := range privs {
		select {
		case <-done:
			return
		default:
		}

		m := 100 * types.DefaultCoinPrecision
		ch <- newTx(rootKey, m, address.PubKeyToAddr(ethID, priv.PubKey().Bytes()))
	}
	log.Println(n, len(privs))
}

//
func runGenerateAccounts(max int, pksCh chan []crypto.PrivKey) {
	log.Println("generate accounts begin:")

	t := time.Now()
	goN := 16
	ch := make(chan pp, goN)
	done := make(chan struct{}, 1)

	c, _ := crypto.Load(types.GetSignName("", types.SECP256K1), -1)
	for i := 0; i < goN; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					priv, _ := c.GenKey()
					ch <- pp{i: 0, p: priv}
				}
			}
		}()
	}
	all := 0
	var pks []crypto.PrivKey
	for {
		p := <-ch
		pks = append(pks, p.p)
		l := len(pks)
		if l%1000 == 0 && l > 0 {
			all += 1000
			log.Println("generate acc:", all, *maxaccF, l)
			if l%(*maxaccF) == 0 {
				pksCh <- pks
				log.Println(time.Since(t))
				pks = nil
			}
		}
		if all == max {
			close(done)
			break
		}
	}
	log.Println("generate accounts end", time.Since(t))
	close(pksCh)
}

func loadAccounts(filename string, max int) []crypto.PrivKey {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		i := 0
		ss := strings.Split(filename, ".")
		pksCh := make(chan []crypto.PrivKey, 10)

		go runGenerateAccounts(max, pksCh)
		for {
			privs := <-pksCh
			if privs == nil {
				log.Println("xxxxxx")
				return nil
			}
			i++
			fname := ss[0] + fmt.Sprintf("%d", i) + "." + ss[1]
			f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			var data []byte
			data = append(data, Int(len(privs)).Marshal()...)
			f.Write(data)
			for _, p := range privs {
				f.Write(p.Bytes())
			}
			log.Println("write acc file", fname)
		}
	}
	var l Int
	n, err := l.Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}
	ln := int(l)
	b = b[n:]

	if max < ln {
		ln = max
	}

	privs := make([]crypto.PrivKey, ln)

	n = 32
	c, _ := crypto.Load(types.GetSignName("", types.SECP256K1), -1)
	for i := 0; i < ln; i++ {
		p := b[:n]
		priv, err := c.PrivKeyFromBytes(p)
		if err != nil {
			log.Fatal(err)
		}
		privs[i] = priv
		b = b[n:]
		if i%10000 == 0 {
			log.Println("load acc:", i)
		}
		if i%1000 == 0 {
			log.Println("account: ", address.PubKeyToAddr(ethID, priv.PubKey().Bytes()))
		}
	}
	log.Println("loadAccount: ", len(privs))
	return privs
}

func runLoadAccounts(filename string, max int) chan crypto.PrivKey {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		i := 0
		ss := strings.Split(filename, ".")
		pksCh := make(chan []crypto.PrivKey, 1)

		go runGenerateAccounts(max, pksCh)
		for {
			privs := <-pksCh
			if privs == nil {
				log.Println("xxxxxx")
				return nil
			}
			i++
			fname := ss[0] + fmt.Sprintf("%d", i) + "." + ss[1]
			f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			var data []byte
			data = append(data, Int(len(privs)).Marshal()...)
			f.Write(data)
			for _, p := range privs {
				f.Write(p.Bytes())
			}
			log.Println("write acc file", fname)
		}
	}
	var l Int
	n, err := l.Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}
	ln := int(l)
	b = b[n:]

	if max < ln {
		ln = max
	}

	privCh := make(chan crypto.PrivKey, 1024)

	go func() {
		n = 32
		c, _ := crypto.Load(types.GetSignName("", types.SECP256K1), -1)
		for i := 0; i < ln; i++ {
			p := b[:n]
			priv, err := c.PrivKeyFromBytes(p)
			if err != nil {
				log.Fatal(err)
			}
			privCh <- priv
			b = b[n:]
			if i%1000 == 0 {
				log.Println("account: ", i, " ", address.PubKeyToAddr(ethID, priv.PubKey().Bytes()))
			}
		}
		close(privCh)
	}()
	return privCh
}
