import Web3 from 'web3';
import { AbiItem } from 'web3-utils';
import { TransactionReceipt } from 'web3-core';
import { Contract } from 'web3-eth-contract';
import { logger } from '../src/utils/logger.js';
import { BaseTrie } from 'merkle-patricia-tree';
import fs from 'fs';
import BeaconChainClient from './beaconChainClient.js';
import data from '../../build/contracts/Eth2ChainRelay_512.json' with { type:"json" };
import RLP from 'rlp';

let abi = data.abi;
let bytecode = data.bytecode;

export default class VerilayClient {
    relayContract: Contract;

    beaconChainClient: BeaconChainClient;

    senderAccount: string;

    sourceUrl: string;

    constructor(
        _targetUrl: string,
        _sourceUrl: string,
        _account: string,
        _Contractaddress?: string,
    ) {
        const web3 = new Web3(_targetUrl);
        const signer = web3.eth.accounts.privateKeyToAccount(
            '0x' + _account,
          );
        web3.eth.accounts.wallet.add(signer);
        this.senderAccount = signer.address;
        this.relayContract = new web3.eth.Contract(abi as AbiItem[], _Contractaddress);
        this.sourceUrl = _sourceUrl;
        this.beaconChainClient = new BeaconChainClient(_sourceUrl);
    }

    private encodedHeaderFromBlock = (block: any) => {
        return RLP.encode([
          block.parentHash,
          block.sha3Uncles,
          block.miner,
          block.stateRoot,
          block.transactionsRoot,
          !!block.receiptsRoot,
          block.logsBloom,
          Web3.utils.toHex(Web3.utils.toBN(block.difficulty)),
          Web3.utils.toHex(Web3.utils.toBN(block.number)),
          Web3.utils.toHex(Web3.utils.toBN(block.gasLimit)),
          Web3.utils.toHex(Web3.utils.toBN(block.gasUsed)),
          Web3.utils.toHex(Web3.utils.toBN(block.timestamp)),
          block.extraData,
          !!block.mixHash,
          Web3.utils.toHex(Web3.utils.toBN(block.nonce)),
        ]);
    }

    /*function encodeTransaction(tx) {
        return ethers.utils.RLP.encode([
          ethers.utils.hexlify(ethers.BigNumber.from(tx.nonce)),
          ethers.utils.hexlify(ethers.BigNumber.from(tx.gasPrice)),
          ethers.utils.hexlify(tx.gasLimit),
          tx.to,
          ethers.utils.hexlify(tx.value),
          tx.data,
          ethers.utils.hexlify(ethers.BigNumber.from(tx.v)),
          tx.r,
          tx.s,
        ]);
      }
      
      function encodeLogs(logs) {
        const encodedLogs = [];
        for (let i = 0; i < logs.length; i++) {
          encodedLogs.push([logs[i].address, logs[i].topics, logs[i].data]);
        }
        return encodedLogs;
      }
      
      function encodeTransactionReceipt(receipt) {
        return ethers.utils.RLP.encode([
          ethers.utils.hexlify(receipt.status),
          ethers.utils.hexlify(receipt.cumulativeGasUsed),
          receipt.logsBloom,
          encodeLogs(receipt.logs),
        ]);
      }
      
    async function encodedTransactionMerkleProof(block, index) {
        const trie = new Trie();
      
        for (let i = 0; i < block.transactions.length; i++) {
          const rlpTx = encodeTransaction(
            await ethers.provider.getTransaction(block.transactions[i].hash)
          );
          const key = ethers.utils.RLP.encode(
            ethers.utils.hexlify(ethers.BigNumber.from(i))
          );
          await trie.put(key, rlpTx);
        }
      
        const key = ethers.utils.RLP.encode(
          ethers.utils.hexlify(ethers.BigNumber.from(index))
        );
        return ethers.utils.RLP.encode(await Trie.createProof(trie, key));
      }
      
    async function encodedTransactionReceiptMerkleProof(block, index) {
        const trie = new Trie();
      
        for (let i = 0; i < block.transactions.length; i++) {
          const rlpReceipt = encodeTransactionReceipt(
            await ethers.provider.getTransactionReceipt(block.transactions[i].hash)
          );
          const key = ethers.utils.RLP.encode(
            ethers.utils.hexlify(ethers.BigNumber.from(i))
          );
          await trie.put(key, rlpReceipt);
        }
      
        const key = ethers.utils.RLP.encode(
          ethers.utils.hexlify(ethers.BigNumber.from(index))
        );
        return ethers.utils.RLP.encode(await Trie.createProof(trie, key));
    }*/
    
    public createAndVerifyBlockProof = async (_blockNumber: number): Promise<string> => {
        const holeskyNet = new Web3("https://ethereum-holesky.publicnode.com")
        const block = await holeskyNet.eth.getBlock(_blockNumber);
        //getting signature and sync-commitee_bits of specific block
        const {sync_committee_bits, sync_committee_signature} = await this.beaconChainClient.getBlockSignature(_blockNumber, this.sourceUrl);
        const receipt: TransactionReceipt = await this.relayContract.methods.verifyBlockSignature(sync_committee_signature, block.hash, sync_committee_bits).send({ from: this.senderAccount, gasLimit: 39000000, contractAddress: this.relayContract.options.address});
        
        return receipt.transactionHash;
    }
    
    public deployContract = async (
        _signatureThreshold: number,
        _trustingPeriod: number,
        _finalizedBlockRoot: string,
        _finalizedStateRoot: string,
        _finalizedSlot: number,
        _latestSlot: number,
        _latestSlotWithValidatorSetChange: number,
    ): Promise<string> => {

        const deploy = this.relayContract.deploy({
            data: bytecode,
            arguments: [
                _signatureThreshold,
                _trustingPeriod,
                Web3.utils.numberToHex(_finalizedBlockRoot),
                Web3.utils.numberToHex(_finalizedStateRoot),
                _finalizedSlot,
                _latestSlot,
                _latestSlotWithValidatorSetChange,
            ],
        })
        
        const deployedContract = await deploy.send({
            from: this.senderAccount,
            gas: await deploy.estimateGas(),
        });

        this.relayContract = deployedContract;
        this.relayContract.options.address = deployedContract.options.address;
        logger.info(this.relayContract.options.address);
        
        
        return deployedContract.options.address;
    };

    public initializeSyncCommitee = async (
        _currentOrNext: string,
        _syncCommitteePeriod: number,
    ): Promise<string> => {
        const updateData = await this.beaconChainClient.getCommitteeUpdateData(_syncCommitteePeriod);
        if(_currentOrNext == "current") {
            const receipt: TransactionReceipt = await this.relayContract.methods.initializeCurrent(updateData.syncCommittee, updateData.syncCommitteeAggregate).send({ from: this.senderAccount, gasLimit: 39000000, contractAddress: this.relayContract.options.address});
            return receipt.transactionHash;
        } else {
            if(_currentOrNext == "next") {
                const receipt: TransactionReceipt = await this.relayContract.methods.initializeNext(updateData.syncCommittee, updateData.syncCommitteeAggregate).send({ from: this.senderAccount, gasLimit: 39000000, contractAddress: this.relayContract.options.address});
                return receipt.transactionHash;
            }
        }
        
        return "";
    };
    
    public deployContractAtPeriod = async (
        _syncCommitteePeriod: number,
        _signatureThreshold: number,
        _trustingPeriod: number,
    ): Promise<string> => {
        const updateData = await this.beaconChainClient.getCommitteeUpdateData(_syncCommitteePeriod);

        return this.deployContract(
            _signatureThreshold,
            _trustingPeriod,
            updateData.finalizedBlockRoot,
            updateData.finalizedStateRoot,
            updateData.finalizedSlot,
            updateData.latestSlot,
            updateData.latestSlot - 10000, // ??? Potential TODO: find out how to calculate or hardcode
        );
    };

    public updateSyncCommittee = async (
        _syncCommitteePeriod: number,
    ): Promise<string> => {
        const publicNodeData = await this.beaconChainClient.getCommitteeUpdateData(_syncCommitteePeriod);
        //fetching sync-commitee data in seperate variable
        const updateData = {
                signature: publicNodeData.signature,
                participants: publicNodeData.participants,
                latestBlockRoot: publicNodeData.latestBlockRoot,
                signingDomain: publicNodeData.signingDomain,
                stateRoot: publicNodeData.stateRoot,
                stateRootBranch: publicNodeData.stateRootBranch,
                latestSlot: publicNodeData.latestSlot,
                latestSlotBranch: publicNodeData.latestSlotBranch,
                finalizedBlockRoot: publicNodeData.finalizedBlockRoot,
                finalizingBranch: publicNodeData.finalizingBranch,
                finalizedSlot: publicNodeData.finalizedSlot,
                finalizedSlotBranch: publicNodeData.finalizedSlotBranch,
                finalizedStateRoot: publicNodeData.finalizedStateRoot,
                finalizedStateRootBranch: publicNodeData.finalizedStateRootBranch,
            }
        const syncCommitee = {
                nextNextValidatorSet: publicNodeData.syncCommittee,
                nextNextValidatorSetAggregate: publicNodeData.syncCommitteeAggregate,
                nextNextValidatorSetBranch: publicNodeData.syncCommitteeBranch
            }
        
        logger.info(this.relayContract.options.address);
        const receipt: TransactionReceipt = await this.relayContract.methods.submitUpdate(updateData
            , syncCommitee
            ).send({ from: this.senderAccount, gasLimit: 39000000, contractAddress: this.relayContract.options.address});
        return receipt.transactionHash;
    };

    public updateEpoch = async (
        _epoch: number,
    ): Promise<string> => {
        const updateData = await this.beaconChainClient.getBlockUpdateDataForEpoch(_epoch);

        const receipt: TransactionReceipt = await this.relayContract.methods.submitUpdate(updateData).send({ from: this.senderAccount, gasLimit: 39000000, contractAddress: this.relayContract.options.address});
        return receipt.transactionHash;
    };

    public setRelayContractAddress = (_address: string) => {
        this.relayContract.options.address = _address;
    };
}
