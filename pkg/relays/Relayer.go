package relays

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"strconv"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"net/http"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/prysmaticlabs/prysm/shared/bls"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
)

type relayer struct {
	bscClient        *ethclient.Client
	ethholClient     *ethclient.Client
	Verilay32        *Main
	privateKey       *ecdsa.PrivateKey
	batchSize        uint
	ethChainID       *big.Int
	bscChainID       *big.Int
}

func NewRelayer(c Config) (*relayer, error) {
	bscClient, err := ethclient.Dial(c.BSCHost)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %w", err)
	}

	ethholClient, err := ethclient.Dial(c.EthereumHost)
	if err != nil {
		return nil, fmt.Errorf("dial bsc client: %w", err)
	}

	VerilayContract, err := NewMain(common.HexToAddress(c.ContractAddress), bscClient)
	if err != nil {
		return nil, fmt.Errorf("new relay contract: %w", err)
	}

	privateKey, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("private key hex to ecdsa: %w", err)
	}

	ethChainID, err := bscClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get source chain id: %w", err)
	}

	bscChainID, err := ethholClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get source chain id: %w", err)
	}

	return &relayer{
		bscClient:        bscClient,
		ethholClient:        ethholClient,
		Verilay32: 	      VerilayContract,
		privateKey:       privateKey,
		batchSize:        c.BatchSize,
		ethChainID:       ethChainID,
		bscChainID:       bscChainID,
	}, nil
}

func (r *relayer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := r.relay(ctx); err != nil {
				log.Errorf("relay: %v", err)
			}
		}
	}
}

func fetchData(url string, data *[]byte) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Fetch Error :-S", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Fetch Error :-S", err)
	}
	*data = body
}
	
func fetchValidatorData(indicesOrPubkey string, validatorSet *[]string) {
	requestBody, err := json.Marshal(map[string]string{"indicesOrPubkey": indicesOrPubkey})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	response, err := http.Post("https://holesky.beaconcha.in/api/v1/validator", "application/json", strings.NewReader(string(requestBody)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}	
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(data)
	for _, v := range data["data"].([]interface{}) {
		validator := v.(map[string]interface{})["pubkey"].(string)
		*validatorSet = append(*validatorSet, validator)
	}
}
	
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (r *relayer) relay(ctx context.Context) error {
	//getting data / setting data for relay update
	const urlcurrentSyncCommitee = "https://holesky.beaconcha.in/api/v1/sync_committee/latest"
	const urlnextSyncCommitee = "https://holesky.beaconcha.in/api/v1/sync_committee/latest"
	const urlslotfinal = "https://holesky.beaconcha.in/api/v1/epoch/finalized/slots"
	const urlslot = "https://holesky.beaconcha.in/api/v1/slot/"

	var _validatorSetInformation []byte
	var _nextValidatorSetInformation []byte
	var _finalizedEpochData []byte

	//getting sync_committe
	fetchData(urlcurrentSyncCommitee, &_validatorSetInformation)
	fetchData(urlnextSyncCommitee, &_nextValidatorSetInformation)
	fetchData(urlslotfinal, &_finalizedEpochData)

	var _currentValidatorSet []byte
	var _nextValidatorSet []byte
	var _validatorSetInformationSplitted [][]string
	var _nextValidatorSetInformationSplitted [][]string

	time.Sleep(60000 * time.Millisecond)

	fmt.Println("Now seperating Indexes of Validators")

	// Splitting validator set information
	for i := 0; i < len(_validatorSetInformation); i += 100 {
		chunk := _validatorSetInformation[i:min(i+100, len(_validatorSetInformation))]
		chunknext := _nextValidatorSetInformation[i:min(i+100, len(_nextValidatorSetInformation))]
		_validatorSetInformationSplitted = append(_validatorSetInformationSplitted, chunk)
		_nextValidatorSetInformationSplitted = append(_nextValidatorSetInformationSplitted, chunknext)
	}

	for i := 0; i < 7; i++ {
		indicesOrPubkey := strings.Join(_validatorSetInformationSplitted[i], ",")
		fetchValidatorData(indicesOrPubkey, &_currentValidatorSet)
	}

	time.Sleep(60000 * time.Millisecond)

	for i := 0; i < 7; i++ {
		indicesOrPubkey := strings.Join(_nextValidatorSetInformationSplitted[i], ",")
		fetchValidatorData(indicesOrPubkey, &_nextValidatorSet)
	}

	//fetch latest finalized slot
	fmt.Println("Now Getting Finalized Slot and Root")
	_finalizedData := _finalizedEpochData[len(_finalizedEpochData)-1]

	fmt.Println(_currentValidatorSet)
	
	chainRelayUpdate := ChainRelayUpdate
      {
        signature: "0xa79b5387e04ecfe315f4acac0b490f3dbe188cd4e004855f06708ab1935d834efc5c57d0bd3a34d54c33439fbd394406086588c1c2ca7f53185daa6a0af17e65bdeff12860b7703df1960df33e50bc4cc4978459a8c70925ad0c459953b70c05",
        participants: [true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false, true, false],
        latestBlockRoot: "0x4087870d7603db08be22daba18362f32dc57c8575d1c615a4361e583aecfeafe",
        signingDomain: "0x01000000cf7b10b712a4684db6e4d3b7f46c2e3272f12b22c1af641ec12ad4d2",
        latestSlot: 0,
        latestSlotBranch: ["0x0000000000000000000000000000000000000000000000000000000000000000", "0x7444b1f8ef0425d6a23089958f777caf7247cea5069d109856366bb5cf5fe96c", "0xd13b9159e1591c3a417cf0369c6032e29a9182cbec64407c1f73f4a0b5801ca6"],
        finalizedBlockRoot: "0x4087870d7603db08be22daba18362f32dc57c8575d1c615a4361e583aecfeafe", 
        finalizingBranch: ['0x0000000000000000000000000000000000000000000000000000000000000000',
        '0x504e8dfadf27180c21c80f1886b202de4c1c89ee2977562433abbbc92ea194e6',
        '0x923ebc2107d5288dad0afd4187e0fe6caae2df92fcf6252ce7ce7feea56e7152',
        '0xfeb3abafe78c2a78ed08962a29cb063c00ba6b3ae84ea7ec84bd8700182197dd',
        '0xc78009fdf07fc56a11f122370658a353aaa542ed63e44c4bc15ff4cd105ab33c',
        '0x164afa0f9d14bd27fceed6cdfdc3acc9d662b29895d7e03d4bc4fb163c7207aa'],
        finalizedStateRoot: "0x89ede5efc4c097ba778e0a6197c49a184c08ccc812139f4cabef1247d18649d8",
        finalizedStateRootBranch: ['0x0000000000000000000000000000000000000000000000000000000000000000','0xf5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b','0xd13b9159e1591c3a417cf0369c6032e29a9182cbec64407c1f73f4a0b5801ca6']
       ,
        finalizedSlot: 0, 
        finalizedSlotBranch: ["0x0000000000000000000000000000000000000000000000000000000000000000", "0x7444b1f8ef0425d6a23089958f777caf7247cea5069d109856366bb5cf5fe96c", "0xd13b9159e1591c3a417cf0369c6032e29a9182cbec64407c1f73f4a0b5801ca6"],
        stateRoot: "0x89ede5efc4c097ba778e0a6197c49a184c08ccc812139f4cabef1247d18649d8",
        stateRootBranch: ['0x0000000000000000000000000000000000000000000000000000000000000000','0xf5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b','0xd13b9159e1591c3a417cf0369c6032e29a9182cbec64407c1f73f4a0b5801ca6']       
    }

    syncCommitteeUpdate := SyncCommitteeUpdate
    {
      nextNextValidatorSet: _currentValidatorSet,
      nextNextValidatorSetAggregate: AggregatePublicKeys(_currentValidatorSet),
      nextNextValidatorSetBranch: []
    }
	
	currentHeight, err := r.VerilayContract.GetFinalizedSlot(nil)
	if err != nil {
		return fmt.Errorf("get current height: %w", err)
	}

	//blockNumber, err := r.ethholClient.BlockNumber(ctx)
	slotNumber, err := _finalizedSlot
	if err != nil {
		return fmt.Errorf("blocknumber: %w", err)
	}

	unsignedHeaders := make([][]byte, 0)
	signedHeaders := make([][]byte, 0)

	//diff := blockNumber - currentHeight.Uint64()
	diff:= _finalizedSlot - currentHeight.Uint64()
	// if diff > uint64(r.batchSize) {
	// 	diff = uint64(r.batchSize)
	// }
	targetHeight := currentHeight.Uint64() + diff

	for i := currentHeight.Uint64(); i < targetHeight; i++ {
		
		//update calls for different slots
		//TODO look after completed epoch
		var _slotData []byte
		fetchData(urlslot+strconv.Itoa(i), &_slotData)
		RelayUpdate := ChainRelayUpdate{
			signature = _slotData.signature,
			participants = ,
			latestBlockRoot = _slotData.blockroot,
			signingDomain = _slotData.syncaggregate_signature, //syncaggregate_signature
			latestSlot:,
			latestSlotBranch,
			finalizedBlockRoot = 
			finalizedStateRootBranch: ['0x0000000000000000000000000000000000000000000000000000000000000000','0xf5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b','0xd13b9159e1591c3a417cf0369c6032e29a9182cbec64407c1f73f4a0b5801ca6']
       ,
        	finalizedSlot: 0, 
        	finalizedSlotBranch: ["0x0000000000000000000000000000000000000000000000000000000000000000", "0x7444b1f8ef0425d6a23089958f777caf7247cea5069d109856366bb5cf5fe96c", "0xd13b9159e1591c3a417cf0369c6032e29a9182cbec64407c1f73f4a0b5801ca6"],
        	stateRoot: _slotData.stateroot,
        	stateRootBranch: ['0x0000000000000000000000000000000000000000000000000000000000000000','0xf5a5fd42d16a20302798ef6ed309979b43003d2320d9f0e8ea9831a92759fb4b','0xd13b9159e1591c3a417cf0369c6032e29a9182cbec64407c1f73f4a0b5801ca6']
		}
		
		
		
		
		
		
		
		header, err := r.ethholClient.HeaderByNumber(ctx, big.NewInt(int64(i+1)))
		if err != nil {
			return fmt.Errorf("header by number: %w", err)
		}

		signedHeader := new(bytes.Buffer)
		err = rlp.Encode(signedHeader, header)
		if err != nil {
			return fmt.Errorf("rlp encode signed header: %w", err)
		}
		signedHeaders = append(signedHeaders, signedHeader.Bytes())

		unsignedHeader, err := EncodeUnsignedHeaderToRLP(header, r.bscChainID)
		if err != nil {
			return fmt.Errorf("rlp encode unsigned header: %w", err)
		}
		unsignedHeaders = append(unsignedHeaders, unsignedHeader)
	}

	log.Infof("Submitting headers from height %d to %d", currentHeight.Int64(), targetHeight)

	transactOpts, err := bind.NewKeyedTransactorWithChainID(r.privateKey, r.ethChainID)
	if err != nil {
		return fmt.Errorf("keyed transactor with chain id: %w", err)
	}

	tx, err := r.VerilayContract.SubmitBlockHeaderBatch(transactOpts, unsignedHeaders, signedHeaders)
	if err != nil {
		return fmt.Errorf("submit block header: %w", err)
	}

	_, err = bind.WaitMined(ctx, r.bscClient, tx)
	if err != nil {
		return fmt.Errorf("wait mined: %w", err)
	}

	return nil
}
