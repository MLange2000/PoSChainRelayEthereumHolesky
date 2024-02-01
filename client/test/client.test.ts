import { describe, before } from 'mocha';
import { expect, assert } from 'chai';
import Web3 from 'web3';

import { logger } from '../src/utils/logger.js';
import VerilayClient from '../src/index.js';

describe('Deploy and update relay from historic state', () => {
    const targetUrl = 'https://bsc-testnet.publicnode.com';
    const sourceUrl = 'https://lodestar-holesky.chainsafe.io/'; //lodestar
    //const sourceUrl = 'http://localhost:9596'; //lodestar https://lodestar-holesky.chainsafe.io/
    let syncCommitteePeriod = 105;
    let verilayClient: VerilayClient;

    before(async () => {
        logger.setSettings({ minLevel: 'info', name: 'client test setup' });

        const web3 = new Web3(targetUrl);
        //const [account] = await web3.eth.getAccounts();
        //expect(account).to.exist;
        verilayClient = new VerilayClient(targetUrl, sourceUrl, "xyz");
    });

    it('should deploy the relay contract', async () => {
        logger.setSettings({ minLevel: 'info', name: 'deploy relay contract' });

        const relayContractAddress = await verilayClient.deployContractAtPeriod(
            syncCommitteePeriod++,
            170,
            0,
        );
        logger.info(relayContractAddress);
        expect(relayContractAddress).to.exist;
    });

    const updateRelayContract = (i: number) => {
        it(`should update relay contract with new sync committee period, ${i}. iteration`, async () => {
            logger.setSettings({ minLevel: 'info', name: `update relay, ${i} iteration` });
    
            const txHash = await verilayClient.updateSyncCommittee(syncCommitteePeriod+2);
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
