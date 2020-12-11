package main

import (
	"CSPTest/config"
	"CSPTest/key"
	"CSPTest/simulate"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 测试用
var cfg *config.Config
var reqcfg *config.RequstConfig
var ecms *simulate.ECMArray
var fks *key.FillKeyArray

func main() {
	start := time.Now()
	var k uint64
	for i := 0; i < 100000000; i++ {
		k = k + uint64(i)
	}
	fmt.Println(time.Since(start), k)
	start = time.Now()
	//var k uint64
	k = 0
	for i := 0; i < 100000000; i++ {
		k = atomic.AddUint64(&k, uint64(i))
	}
	fmt.Println(time.Since(start), k)
	var mut sync.Mutex
	start = time.Now()
	//var k uint64
	k = 0
	for i := 0; i < 100000000; i++ {
		mut.Lock()
		k = k + uint64(i)
		mut.Unlock()
	}
	fmt.Println(time.Since(start), k)
	return

	cfg = config.ParseConfig("config.json")
	reqcfg = config.ParseRequestConfig("request.json")
	ecms = simulate.CreateECMArray(cfg.EcmNum, cfg, reqcfg) // 密码机个数，cdg,recfg
	defer ecms.Destory()
	fks = key.CreateFillKeyArray(cfg.CardNum, cfg.FKeyLen, cfg.IDPrefix, ecms) // 充注密钥数，密钥长度，ecms
	defer fks.Destory()

	ecmcards := key.CreateCardIDByECM(cfg.CardNum, cfg.EcmNum, cfg.IDPrefix)
	sks := key.CreateSessKeyArrayByECM(cfg.Parallel, cfg.TimesPara, ecmcards,
		cfg.HttpUrl, reqcfg)
	defer sks.Destory()

	fmt.Println(fks.GetCardID())
	fmt.Println(key.CreateCardID(cfg.CardNum, cfg.IDPrefix))
	fmt.Println(fks.GetCardIDByECM())
	fmt.Println(key.CreateCardIDByECM(cfg.CardNum, cfg.EcmNum, cfg.IDPrefix))

	return
}
