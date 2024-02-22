import { describe, before } from 'mocha';
import { expect, assert } from 'chai';
import Web3 from 'web3';

import { logger } from '../src/utils/logger.js';
import VerilayClient from '../src/index.js';

describe('Deploy and update relay from historic state', () => {
    const targetUrl = 'https://bsc-testnet.publicnode.com';
    const sourceUrl = 'https://lodestar-holesky.chainsafe.io/'; //lodestar
    //const sourceUrl = 'http://localhost:9596'; //lodestar https://lodestar-holesky.chainsafe.io/
    let syncCommitteePeriod = 128;
    let verilayClient: VerilayClient;

    before(async () => {
        logger.setSettings({ minLevel: 'info', name: 'client test setup' });

        const web3 = new Web3(targetUrl);
        //const [account] = await web3.eth.getAccounts();
        //expect(account).to.exist;
        verilayClient = new VerilayClient(targetUrl, sourceUrl, "WalletPrivateKey");
    });

    it('should deploy the relay contract', async () => {
        logger.setSettings({ minLevel: 'info', name: 'deploy relay contract' });

        const relayContractAddress = await verilayClient.deployContractAtPeriod(
            syncCommitteePeriod,
            170,
            0,
        );
        logger.info(relayContractAddress);
        expect(relayContractAddress).to.exist;
    });

    it('should initialize sync-commitee', async () => {
        logger.setSettings({ minLevel: 'info', name: 'initialize current sync-commitee...' });

        const txHash = await verilayClient.initializeSyncCommitee(
            "current",
            syncCommitteePeriod
        );
        expect(txHash).to.exist;

        logger.setSettings({ minLevel: 'info', name: 'initialize next sync-commitee...' });

    });

    it('should initialize next sync-commitee', async () => {
        logger.setSettings({ minLevel: 'info', name: 'initialize current sync-commitee...' });

        const txHash = await verilayClient.initializeSyncCommitee(
            "next",
            syncCommitteePeriod++
        );
        expect(txHash).to.exist;

        logger.setSettings({ minLevel: 'info', name: 'initialize next sync-commitee...' });

    });

    it('should check if random block in source-net can be validated with its sync_aggregate-information', async () => {
        logger.setSettings({ minLevel: 'info', name: 'getting sync-commitee-information...' });

        const txHash = await verilayClient.createAndVerifyBlockProof(988914);
        expect(txHash).to.exist;

        logger.setSettings({ minLevel: 'info', name: 'initialize next sync-commitee...' });

    });

    const updateRelayContract = (i: number) => {
        it(`should update relay contract with new sync committee period, ${i}. iteration`, async () => {
            logger.setSettings({ minLevel: 'info', name: `update relay, ${i} iteration` });
    
            const txHash = await verilayClient.updateSyncCommittee(syncCommitteePeriod);
            expect(txHash).to.exist;
            logger.info(`${i}. update txHash: ${txHash}`);
        });
    };

    describe('multiple updates', () => {
        for (let i = 1; i <= 1; i++) {
            updateRelayContract(i);
        }  
    });

});
