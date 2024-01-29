import { describe, before } from 'mocha';
import { expect, assert } from 'chai';
import Web3 from 'web3';

import { logger } from '../src/utils/logger';
import VerilayClient from '../src/index';

describe('Deploy and update relay from historic state', () => {
    const targetUrl = 'https://bsc-testnet.publicnode.com';
    const sourceUrl = 'http://localhost:9596'; //lodestar
    let syncCommitteePeriod = 1004;
    let verilayClient: VerilayClient;

    before(async () => {
        logger.setSettings({ minLevel: 'info', name: 'client test setup' });

        const web3 = new Web3(targetUrl);
        //const [account] = await web3.eth.getAccounts();
        //expect(account).to.exist;
        verilayClient = new VerilayClient(targetUrl, sourceUrl, "c97790a0e16de2f90568fbf578cfbf04540a562bf4bff01cc76b002d64b55aea");
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

});
