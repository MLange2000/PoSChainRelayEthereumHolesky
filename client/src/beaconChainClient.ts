//old imports without update of lodestar to an esm project
import { Api, getClient} from '@lodestar/api';
import { config } from '@lodestar/config/default';
import { createBeaconConfig, BeaconConfig } from '../node_modules/@lodestar/config/lib/beaconConfig.js';
import { toHexString } from '@chainsafe/ssz';
import { DOMAIN_SYNC_COMMITTEE } from '@lodestar/params';
import { ssz, altair } from '@lodestar/types';
//import { isValidMerkleBranch } from '../node_modules/@chainsafe/lodestar-light-client/lib/utils/verifyMerkleBranch.js';
import { verifyMerkleBranch } from '@lodestar/utils';
import { computeSyncPeriodAtEpoch, computeSyncPeriodAtSlot } from '@lodestar/state-transition';
import { createNodeFromProof, computeDescriptor } from '@chainsafe/persistent-merkle-tree';
import { createSingleProof } from '../node_modules/@chainsafe/persistent-merkle-tree/lib/proof/single.js';
import {Tree} from "@chainsafe/persistent-merkle-tree";

import { ChainRelayUpdate} from './types.js';
import { logger } from './utils/logger.js';

export default class BeaconChainClient {
    clientNode: Api;

    static readonly FINALIZED_ROOT_INDEX = 105;

    // static readonly CURRENT_SYNC_COMMITTEE_INDEX = 54;
    static readonly NEXT_SYNC_COMMITTEE_INDEX = 55;

    static readonly SLOT_INDEX = 8;

    static readonly STATE_ROOT_INDEX = 11;

    static readonly SLOTS_PER_EPOCH = 32;

    constructor(_url: string) {
        this.clientNode = getClient({ baseUrl: _url }, {config});
    }

    private static fromGindex = (gindex: number): [number, number] => {
        const depth = Math.floor(Math.log2(gindex));
        const firstIndex = 2 ** depth;
        return [depth, gindex % firstIndex];
    };

    private static uint8ArrayToBooleanArray(uint8Array: Uint8Array): boolean[] {
        const booleanArray: boolean[] = [];
      
        for (let i = 0; i < uint8Array.length; i++) {
          const byte = uint8Array[i];
      
          for (let j = 7; j >= 0; j--) {
            const bit = (byte >> j) & 1;
            booleanArray.push(bit === 1);
          }
        }
      
        return booleanArray;
    }
    
    public getCommitteeUpdateData = async (
        _syncCommitteePeriod: number,
    ): Promise<ChainRelayUpdate> => {
        logger.debug('Moin! A getCommitee Update started...');
        const [prevCommitteeUpdate, committeeUpdate] = (await this.clientNode.lightclient.getUpdates(_syncCommitteePeriod - 1, 2)).response!
        //logger.debug('finalized slot: ', prevCommitteeUpdate.data);
        const finalizedBlockHeader = committeeUpdate.data.attestedHeader;
        const finalizedSlot = finalizedBlockHeader.beacon.slot;
        const finalizedStateRoot = finalizedBlockHeader.beacon.stateRoot;
        const finalizedBlockRoot = (await this.clientNode.beacon.getBlockRoot(finalizedSlot)).response!;

        const latestBlockHeader = committeeUpdate.data.finalizedHeader;
        const latestSlot = latestBlockHeader.beacon.slot;
        const latestBlockRoot = ssz.phase0.BeaconBlockHeader.hashTreeRoot(latestBlockHeader.beacon);
        const latestStateRoot = latestBlockHeader.beacon.stateRoot;

        const { nextSyncCommitteeBranch } = committeeUpdate.data;
        const { nextSyncCommittee } = committeeUpdate.data;

        
        logger.debug('latest block root: ', toHexString(latestBlockRoot));
        //const stateRootBranch = new Tree(ssz.phase0.BeaconBlockHeader.value_toTree(latestBlockHeader.beacon)).getSingleProof(BigInt(BeaconChainClient.STATE_ROOT_INDEX));
        const stateRootBranch = new Tree(ssz.phase0.BeaconBlockHeader.toView(latestBlockHeader.beacon).node).getSingleProof(BigInt(BeaconChainClient.STATE_ROOT_INDEX));
        logger.info(
            'stateRootBranch valid: ',
            verifyMerkleBranch(
                latestStateRoot.valueOf() as Uint8Array,
                stateRootBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.STATE_ROOT_INDEX),
                latestBlockRoot.valueOf() as Uint8Array,
            ),
        );

        //const latestSlotBranch = new Tree(ssz.phase0.BeaconBlockHeader.value_toTree(latestBlockHeader.beacon)).getSingleProof(BigInt(BeaconChainClient.SLOT_INDEX));
        const latestSlotBranch = new Tree(ssz.phase0.BeaconBlockHeader.toView(latestBlockHeader.beacon).node).getSingleProof(BigInt(BeaconChainClient.SLOT_INDEX));
        logger.debug('latestSlotBranch: ', latestSlotBranch.map(toHexString));

        logger.info(
            'latestSlotBranch valid: ',
            verifyMerkleBranch(
                ssz.Slot.hashTreeRoot(latestSlot),
                latestSlotBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.SLOT_INDEX),
                latestBlockRoot.valueOf() as Uint8Array,
            ),
        );

        //const finalizedStateRootBranch = new Tree(ssz.phase0.BeaconBlockHeader.value_toTree(finalizedBlockHeader.beacon)).getSingleProof(BigInt(BeaconChainClient.STATE_ROOT_INDEX));
        const finalizedStateRootBranch = new Tree(ssz.phase0.BeaconBlockHeader.toView(finalizedBlockHeader.beacon).node).getSingleProof(BigInt(BeaconChainClient.STATE_ROOT_INDEX));
        
        logger.debug('finalizedStateRootBranch: ', finalizedStateRootBranch.map(toHexString));

        logger.info(
            'finalizedStateRootBranch valid: ',
            verifyMerkleBranch(
                finalizedStateRoot.valueOf() as Uint8Array,
                finalizedStateRootBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.STATE_ROOT_INDEX),
                finalizedBlockRoot.data.root.valueOf() as Uint8Array,
            ),
        );

        //const finalizedSlotBranch = new Tree(ssz.phase0.BeaconBlockHeader.value_toTree(finalizedBlockHeader.beacon)).getSingleProof(BigInt(BeaconChainClient.SLOT_INDEX));
        const finalizedSlotBranch = new Tree(ssz.phase0.BeaconBlockHeader.toView(finalizedBlockHeader.beacon).node).getSingleProof(BigInt(BeaconChainClient.SLOT_INDEX));
        logger.debug('finalizedSlotBranch: ', finalizedSlotBranch.map(toHexString));

        logger.info(
            'finalizedSlotBranch valid: ',
            verifyMerkleBranch(
                ssz.Slot.hashTreeRoot(finalizedSlot),
                finalizedSlotBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.SLOT_INDEX),
                finalizedBlockRoot.data.root.valueOf() as Uint8Array,
            ),
        );
        

        logger.debug(
            'nextSyncCommitteeBranch valid: ',
            verifyMerkleBranch(
                ssz.altair.SyncCommittee.hashTreeRoot(nextSyncCommittee),
                Array.from(nextSyncCommitteeBranch).map((val) => val.valueOf() as Uint8Array),
                ...BeaconChainClient.fromGindex(BeaconChainClient.NEXT_SYNC_COMMITTEE_INDEX),
                finalizedStateRoot.valueOf() as Uint8Array,
            ),
        );

        logger.debug(
            'finalizingBranch valid: ',
            verifyMerkleBranch(
                finalizedBlockRoot.data.root.valueOf() as Uint8Array,
                Array.from(committeeUpdate.data.finalityBranch).map((val) => val.valueOf() as Uint8Array),
                ...BeaconChainClient.fromGindex(BeaconChainClient.FINALIZED_ROOT_INDEX),
                latestStateRoot.valueOf() as Uint8Array,
            ),
        );

        logger.debug('state root: ', toHexString(latestStateRoot));
        logger.debug('body root: ', toHexString(latestBlockHeader.beacon.bodyRoot));

        logger.debug('finalized block state root: ', toHexString(finalizedStateRoot));

        const { genesisValidatorsRoot } = (await this.clientNode.beacon.getGenesis()).response!.data;
        const beaconConfig: BeaconConfig = createBeaconConfig(config, genesisValidatorsRoot);
        const domain = beaconConfig.getDomain(latestSlot,DOMAIN_SYNC_COMMITTEE);

        const syncCommittee = prevCommitteeUpdate.data.nextSyncCommittee.pubkeys;
        const syncCommitteeBranch = prevCommitteeUpdate.data.nextSyncCommitteeBranch;
        const syncCommitteeAggregate = prevCommitteeUpdate.data.nextSyncCommittee.aggregatePubkey;

        // const { BeaconState } = beaconConfig.getForkTypes(latestSlot);
        // const path = ['blockRoots'];
        // const { gindex } = BeaconState.getPathInfo(path);
        // logger.debug('gindex: ', gindex);
        // const proof = (await this.clientNode.lightclient.getStateProof('head',[path])).data;
        // const [history, historyBranch] = createSingleProof(createNodeFromProof(proof), gindex);
        // logger.debug('history root: ', toHexString(history));
            
        return {
            signature: toHexString(committeeUpdate.data.syncAggregate.syncCommitteeSignature),
            participants: BeaconChainClient.uint8ArrayToBooleanArray(committeeUpdate.data.syncAggregate.syncCommitteeBits.uint8Array),
            latestBlockRoot: toHexString(latestBlockRoot),
            signingDomain: toHexString(domain),
            stateRoot: toHexString(latestStateRoot),
            stateRootBranch: stateRootBranch.map(toHexString),
            latestSlot,
            latestSlotBranch: latestSlotBranch.map(toHexString),
            finalizedBlockRoot: toHexString(finalizedBlockRoot.data.root),
            finalizingBranch: Array.from(committeeUpdate.data.finalityBranch).map(toHexString),
            finalizedSlot,
            finalizedSlotBranch: finalizedSlotBranch.map(toHexString),
            finalizedStateRoot: toHexString(finalizedStateRoot),
            finalizedStateRootBranch: finalizedStateRootBranch.map(toHexString),
            syncCommittee: Array.from(syncCommittee).map(toHexString),
            syncCommitteeAggregate: toHexString(syncCommitteeAggregate),
            syncCommitteeBranch: Array.from(syncCommitteeBranch).map(toHexString),
        };
    };

    public getBlockUpdateDataForEpoch = async (
        _epoch: number,
    ): Promise<ChainRelayUpdate> => {
        const { genesisValidatorsRoot } = (await this.clientNode.beacon.getGenesis()).response!.data;
        const beaconConfig: BeaconConfig = createBeaconConfig(config, genesisValidatorsRoot);

        const finalizedSlot = BeaconChainClient.SLOTS_PER_EPOCH * _epoch;

        // TODO: check reasoning
        // the earliest finalization block appears after two epochs
        const latestSlot = finalizedSlot + BeaconChainClient.SLOTS_PER_EPOCH * 2;

        const signedLatestBlockHeader = (await this.clientNode.beacon.getBlockHeader(latestSlot)).response!.data.header;
        const latestBlockHeader = signedLatestBlockHeader.message;
        //const stateRootBranch = ssz.phase0.BeaconBlockHeader.createTreeBackedFromStruct(latestBlockHeader).tree.getSingleProof(BigInt(BeaconChainClient.STATE_ROOT_INDEX));
        const stateRootBranch = new Tree(ssz.phase0.BeaconBlockHeader.value_toTree(latestBlockHeader)).getSingleProof(BigInt(BeaconChainClient.STATE_ROOT_INDEX));
        const latestStateRoot = latestBlockHeader.stateRoot;
        const latestBlockRoot = ssz.phase0.BeaconBlockHeader.hashTreeRoot(latestBlockHeader);

        const { BeaconState } = beaconConfig.getForkTypes(latestSlot);
        const domain = beaconConfig.getDomain(latestSlot, DOMAIN_SYNC_COMMITTEE);

        // const proof = (await this.clientNode.lightclient.getStateProof(toHexString(latestStateRoot), [['finalized_checkpoint']]));
        logger.debug(
            'stateRootBranch valid: ',
            verifyMerkleBranch(
                latestStateRoot.valueOf() as Uint8Array,
                stateRootBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.STATE_ROOT_INDEX),
                latestBlockRoot.valueOf() as Uint8Array,
            ),
        );

        const latestSlotBranch = new Tree(ssz.phase0.BeaconBlockHeader.value_toTree(latestBlockHeader)).getSingleProof(BigInt(BeaconChainClient.SLOT_INDEX));
        logger.debug(
            'latestSlotBranch valid: ',
            verifyMerkleBranch(
                ssz.Slot.hashTreeRoot(latestSlot),
                latestSlotBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.SLOT_INDEX),
                latestBlockRoot.valueOf() as Uint8Array,
            ),
        );
        const FINALIZED_STATE_ROOT_GINDEX = ssz.capella.BeaconBlock.getPathInfo(['finalizedCheckpoint', 'root']).gindex;
        const descriptor = computeDescriptor([FINALIZED_STATE_ROOT_GINDEX]); //here eventually indexes from above to bigin
        let path = ['finalizedCheckpoint', 'root'];
        const finalityProof = (await this.clientNode.proof.getStateProof(toHexString(latestStateRoot), descriptor)).response!.data; //here eventually indexes from above to bigin
        const [finalizedBlockRoot, finalityBranch] = createSingleProof(createNodeFromProof(finalityProof), BigInt(BeaconChainClient.FINALIZED_ROOT_INDEX));

        logger.debug(
            'finalizingBranch valid: ',
            verifyMerkleBranch(
                finalizedBlockRoot.valueOf() as Uint8Array,
                finalityBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.FINALIZED_ROOT_INDEX),
                latestStateRoot.valueOf() as Uint8Array,
            ),
        );

        logger.debug('computed finalizedBlockRoot: ', toHexString((await this.clientNode.beacon.getBlockRoot(finalizedSlot)).response!.data.root));
        logger.debug('actual finalizedBlockRoot: ', toHexString(finalizedBlockRoot));
        const finalizedBlockHeader = (await this.clientNode.beacon.getBlockHeader(toHexString(finalizedBlockRoot))).response!.data.header.message;
        const finalizedStateRoot = finalizedBlockHeader.stateRoot;
        // const finalizedSlot = finalizedBlockHeader.slot;
        const finalizedStateRootBranch = new Tree(ssz.phase0.BeaconBlockHeader.value_toTree(finalizedBlockHeader)).getSingleProof(BigInt(BeaconChainClient.STATE_ROOT_INDEX));
        logger.debug('finalizedStateRoot: ', toHexString(finalizedStateRoot));

        logger.debug(
            'finalizedStateRootBranch valid: ',
            verifyMerkleBranch(
                finalizedStateRoot.valueOf() as Uint8Array,
                finalizedStateRootBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.STATE_ROOT_INDEX),
                finalizedBlockRoot.valueOf() as Uint8Array,
            ),
        );

        const finalizedSlotBranch = new Tree(ssz.phase0.BeaconBlockHeader.value_toTree(finalizedBlockHeader)).getSingleProof(BigInt(BeaconChainClient.SLOT_INDEX));
        // logger.debug('finalizedSlotBranch: ', finalizedSlotBranch.map(toHexString));

        logger.debug(
            'finalizedSlotBranch valid: ',
            verifyMerkleBranch(
                ssz.Slot.hashTreeRoot(finalizedSlot),
                finalizedSlotBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.SLOT_INDEX),
                finalizedBlockRoot.valueOf() as Uint8Array,
            ),
        );

        path = ['nextSyncCommittee'];
        const { gindex } = BeaconState.getPathInfo(path);
        logger.debug('gindex: ', gindex);

        logger.debug('finalized slot: ', finalizedSlot);
        logger.debug('finalizedStateRoot: ', toHexString(finalizedStateRoot));
        const NEXT_SYNC_COMMITTEE_GINDEX = ssz.capella.BeaconBlock.getPathInfo(['nextSyncCommittee']).gindex;
        const descriptor_syn_commitee = computeDescriptor([NEXT_SYNC_COMMITTEE_GINDEX]); //here eventually indexes from above to bigint
        const nextSyncCommitteeProof = (await this.clientNode.proof.getStateProof(toHexString(finalizedStateRoot), descriptor_syn_commitee)).response!.data;
        const [nextSyncCommitteeRoot, nextSyncCommitteeBranch] = createSingleProof(createNodeFromProof(nextSyncCommitteeProof), gindex);

        const syncPeriod = computeSyncPeriodAtSlot(finalizedSlot);
        const { nextSyncCommittee } = (await this.clientNode.lightclient.getUpdates(syncPeriod, 0)).response![0].data;

        logger.debug('roots equal: ', toHexString(nextSyncCommitteeRoot) === toHexString(ssz.altair.SyncCommittee.hashTreeRoot(nextSyncCommittee)));

        const { syncCommitteeSignature, syncCommitteeBits } = ((await this.clientNode.beacon.getBlock(finalizedSlot)).response!.data.message as altair.BeaconBlock).body.syncAggregate;

        logger.debug(
            'nextSyncCommitteeBranch valid: ',
            verifyMerkleBranch(
                nextSyncCommitteeRoot,
                nextSyncCommitteeBranch,
                ...BeaconChainClient.fromGindex(BeaconChainClient.NEXT_SYNC_COMMITTEE_INDEX),
                finalizedStateRoot.valueOf() as Uint8Array,
            ),
        );

        return {
            signature: toHexString(syncCommitteeSignature),
            participants: BeaconChainClient.uint8ArrayToBooleanArray(syncCommitteeBits.uint8Array),
            latestBlockRoot: toHexString(latestBlockRoot),
            signingDomain: toHexString(domain),
            stateRoot: toHexString(latestStateRoot),
            stateRootBranch: stateRootBranch.map(toHexString),
            latestSlot,
            latestSlotBranch: latestSlotBranch.map(toHexString),
            finalizedBlockRoot: toHexString(finalizedBlockRoot),
            finalizingBranch: Array.from(finalityBranch).map(toHexString),
            finalizedSlot,
            finalizedSlotBranch: finalizedSlotBranch.map(toHexString),
            finalizedStateRoot: toHexString(finalizedStateRoot),
            finalizedStateRootBranch: finalizedStateRootBranch.map(toHexString),
            syncCommittee: Array.from(nextSyncCommittee.pubkeys).map(toHexString),
            syncCommitteeAggregate: toHexString(nextSyncCommittee.aggregatePubkey),
            syncCommitteeBranch: Array.from(nextSyncCommitteeBranch).map(toHexString),
        };
    };

    public getBlockUpdateForEpoch = async (
        _epoch: number,
    ): Promise<ChainRelayUpdate> => {
        const syncCommitteePeriod = computeSyncPeriodAtEpoch(_epoch);
        return this.getCommitteeUpdateData(syncCommitteePeriod);
    };

    public getLatestSyncCommitteePeriod = async (): Promise<number> => {
        // TODO: could be computed w/o node interaction
        const latestBlockHeader = (await this.clientNode.beacon.getBlockHeader('head')).response!.data;
        const latestSlot = latestBlockHeader.header.message.slot;
        return computeSyncPeriodAtSlot(latestSlot);
    };
}
