const Eth2ChainRelay_512 = artifacts.require("Eth2ChainRelay_512");
import { bls12_381 as bls } from '@noble/curves/bls12-381';

//Edited by Max Lange
//Added support for ethereum

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

const urlcurrentSyncCommitee = "https://holesky.beaconcha.in/api/v1/sync_committee/latest";
const urlnextSyncCommitee = "https://holesky.beaconcha.in/api/v1/sync_committee/latest";
const urlslotlatest = "https://holesky.beaconcha.in/api/v1/slot/latest";
const urlepoch = "https://holesky.beaconcha.in/api/v1/epoch/"
const urlslotfinal = "https://holesky.beaconcha.in/api/v1/epoch/finalized/slots"

var _validatorSetInformation;
var _nextValidatorSetInformation;
var _currentSlotData;
var _finalizedEpochData;
var _latestSlotOldValidatorSet;

fetch(urlcurrentSyncCommitee).then(function(response) {
  return response.json();
}).then(function(data) {
  console.log("Got Validator Set")
  console.log(data.data);
  _validatorSetInformation = data.data; 
  console.log(urlepoch + (_validatorSetInformation.start_epoch-1).toString()+"/slots")
  fetch(urlepoch + (_validatorSetInformation.start_epoch-1).toString()+"/slots").then(function(response) {
    return response.json();
  }).then(function(data) {
    console.log("Got Last Slot With Old Validator Set")
    console.log(data.data);
    _latestSlotOldValidatorSet = data.data; 
  }).catch(function(err) {
    console.log('Fetch Error :-S', err);
  });
}).catch(function(err) {
  console.log('Fetch Error :-S', err);
});

fetch(urlnextSyncCommitee).then(function(response) {
  return response.json();
}).then(function(data) {
  console.log("Got Validator Set")
  console.log(data.data);
  _nextValidatorSetInformation = data.data; 
}).catch(function(err) {
  console.log('Fetch Error :-S', err);
});

fetch(urlslotlatest).then(function(response) {
  return response.json();
}).then(function(data) {
  console.log("Got Latest Slot Data")
  console.log(data.data);
  _currentSlotData = data.data; 
}).catch(function(err) {
  console.log('Fetch Error :-S', err);
});

fetch(urlslotfinal).then(function(response) {
  return response.json();
}).then(function(data) {
  console.log("Got Finalized Epoch Data")
  console.log(data.data);
  _finalizedEpochData = data.data; 
}).catch(function(err) {
  console.log('Fetch Error :-S', err);
});



var _currentValidatorSet = []

var _nextValidatorSet = []

var _validatorSetInformationSplitted = [];
var _nextValidatorSetInformationSplitted = [];

sleep(60000).then(() => { 

// fetch("https://holesky.beaconcha.in/api/v1/validator", {
//   method: "POST",
//   body: JSON.stringify({
//     indicesOrPubkey: data.validators
//   }),
//   headers: {
//     "Content-type": "application/json; charset=UTF-8"
//   }
//   }).then(function(response) {
//     response.json()
//     }
//     )

//fetch validator-public-keys
console.log("Now seperating Indexes of Validators")



for (let i = 0; i < _validatorSetInformation.length; i += 100) {
  const chunk = _validatorSetInformation.validators.slice(i, i + 100);
  const chunknext = _nextValidatorSetInformation.validators.slice(i, i + 100);
  _validatorSetInformationSplitted.push(chunk);
  _nextValidatorSetInformationSplitted.push(chunknext);
}

for (let i = 0; i < 7; i++) {
  fetch('https://holesky.beaconcha.in/api/v1/validator', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ indicesOrPubkey: _validatorSetInformationSplitted[i].toString() })
  })
    .then(response => response.json())
    .then(function(data) {
      console.log(data)
      for (var k = 0; k < data.data.length; k++) {
        _currentValidatorSet.push(data.data[k].pubkey)
      }
    })
    .catch(error => console.error(error));
}

// console.log("Now Getting Public-Adresses of Validators")
// for (let i = 0; i < _validatorSetInformation.validators.length; i++) { 
//   fetch("https://holesky.beaconcha.in/api/v1/validator/"+_validatorSetInformation.validators[i]).then(function(response) {
//     return _validatorSetInformation.validators[i] = response.pubkey;
//   }).then(function(data) {
//     console.log(data);
//     _validatorSetInformation = data 
//   }).catch(function(err) {
//     console.log('Fetch Error :-S', err);
//   });
// }

});

sleep(60000).then(() => { 
  for (let i = 0; i < 7; i++) {
    fetch('https://holesky.beaconcha.in/api/v1/validator', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ indicesOrPubkey: _nextValidatorSetInformationSplitted[i].toString() })
    })
      .then(response => response.json())
      .then(function(data) {
        console.log(data)
        for (var k = 0; k < data.data.length; k++) {
          _nextValidatorSet.push(data.data[k].pubkey)
        }
      })
      .catch(error => console.error(error));
  }

//fetch latest finalized slot
console.log("Now Getting Finalized Slot and Root")
_finalizedData = _finalizedEpochData[_finalizedEpochData.length-1]

_currentValidatorSet = _validatorSetInformation.validators



const _currentValidatorSetAggregate = bls.aggregatePublicKeys(_validatorSetInformation.validators)


//trusting period eventually epoch-end-time -> at the moment: period in between
module.exports = function (deployer) {
  deployer.deploy(Eth2ChainRelay_512, 16, _validatorSetInformation.period, _finalizedData.blockroot, _finalizedData.exec_state_root, _finalizedData.slot, _currentSlotData.slot, _latestSlotOldValidatorSet.slot, {gas: 8000000})
  .then(instance => {
    instance.initializeCurrent(_currentValidatorSet, _currentValidatorSetAggregate, {gas: 250000001})
    //I would definiteley change it here, as well
    instance.initializeNext(_currentValidatorSet, _currentValidatorSetAggregate, {gas: 250000002})
  });
};

});
