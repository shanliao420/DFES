package data_server

import (
	"context"
	"github.com/shanliao420/DFES/encryption"
	pb "github.com/shanliao420/DFES/gateway/proto"
	idGenerator "github.com/shanliao420/DFES/id-generator"
	"github.com/shanliao420/DFES/utils"
	"log"
	"os"
	"path/filepath"
)

const (
	DefaultKeyStorePath = "./data/key/"
	DefaultServerName   = "shanliao-data-node-1"
)

var (
	registerClient pb.RegistryClient
	dataService    *DataServer
)

func Init() {
	dataService = NewDataServer(nil, nil, DefaultServerName)
	registerClient = pb.NewRegistryClient(utils.NewGrpcClient(RegistryAddr))
	pri, pub := getKey(DefaultKeyStorePath)
	dataService.privateKey = pri
	dataService.publicKey = pub
	serviceCnt, err := registerClient.GetHistoryAllServiceCnt(context.Background(), nil)
	if err != nil {
		log.Fatalln("init data service in id node err:", err)
	}
	log.Println("init mate server id node -> ", serviceCnt.GetServiceCnt())
	dataService.idGen = idGenerator.NewSnowflakeIdGenerator(serviceCnt.ServiceCnt)
}

func SetDataServerName(serverName string) {
	dataService.serverName = serverName
	dataService.storePath = DefaultDataStorePath + serverName + "/"
}

func getKey(path string) ([]byte, []byte) {
	priName := dataService.serverName + ".private.key"
	pubName := dataService.serverName + ".public.key"
	pri, err := os.ReadFile(filepath.Join(path, priName))
	if err != nil {
		pri, pub := (&encryption.AsymmetricEncryptor{}).GenerateKey(encryption.RSA)
		_ = utils.CreateDirIfNotExist(path)
		err = os.WriteFile(filepath.Join(path, priName), pri, 0700)
		err = os.WriteFile(filepath.Join(path, pubName), pub, 0700)
		if err != nil {
			log.Println("key save error:", err)
		}
		return pri, pub
	}
	pub, err := os.ReadFile(filepath.Join(path, pubName))
	if err != nil {
		pri, pub := (&encryption.AsymmetricEncryptor{}).GenerateKey(encryption.RSA)
		_ = utils.CreateDirIfNotExist(path)
		err = os.WriteFile(priName, pri, 0700)
		err = os.WriteFile(pubName, pub, 0700)
		if err != nil {
			log.Println("key save error:", err)
		}
		return pri, pub
	}
	return pri, pub
}
