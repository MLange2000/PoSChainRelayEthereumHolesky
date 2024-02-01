import Web3 from 'web3';
import { AbiItem } from 'web3-utils';
import { TransactionReceipt } from 'web3-core';
import { Contract } from 'web3-eth-contract';
import { logger } from '../src/utils/logger.js';
import fs from 'fs';
import BeaconChainClient from './beaconChainClient.js';
import data from '../../build/contracts/Eth2ChainRelay_512_NoStorage_Client.json' with { type:"json" };

let abi = data.abi;
let bytecode = data.bytecode;

export default class VerilayClient {
    relayContract: Contract;

    beaconChainClient: BeaconChainClient;

    senderAccount: string;

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
        this.beaconChainClient = new BeaconChainClient(_sourceUrl);
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
        return deployedContract.options.address;
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
            updateData.latestSlot - 10000, // ???
        );
    };

    public updateSyncCommittee = async (
        _syncCommitteePeriod: number,
    ): Promise<string> => {
        const updateData = await this.beaconChainClient.getCommitteeUpdateData(_syncCommitteePeriod);
        logger.info(this.relayContract.options.address);
        const receipt: TransactionReceipt = await this.relayContract.methods.submitUpdate(updateData).send({ from: this.senderAccount, gasLimit: 39000000, contractAddress: this.relayContract.options.address});
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
