// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNumParticipants\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumParticipants\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"DisputeResolutionBegan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"DisputeResolutionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipientAddr\",\"type\":\"address\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AVAILABLE_NODE_TYPES\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"BeginDisputeResolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_MIN_SLASH_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_STAKING_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"EndDisputeResolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAvailableNodeTypes\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GetCurrentAccountStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentDisputeWatchtowers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentSequencers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentSequencersInProbation\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentStakingThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentWatchtowers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtowerAddr\",\"type\":\"address\"}],\"name\":\"GetDisputedSequencerAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"GetDisputedWatchtowerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"GetIsSequencerInProbation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMaxNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMaxNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMaxNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMaxNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMinNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMinNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMinNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMinNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSlashPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_SEQUENCER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_VALIDATOR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_WATCHTOWER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumNumParticipants\",\"type\":\"uint256\"}],\"name\":\"SetMaxNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumNumSequencers\",\"type\":\"uint256\"}],\"name\":\"SetMaxNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumNumValidators\",\"type\":\"uint256\"}],\"name\":\"SetMaxNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumNumWatchtowers\",\"type\":\"uint256\"}],\"name\":\"SetMaxNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumNumParticipants\",\"type\":\"uint256\"}],\"name\":\"SetMinNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumNumSequencers\",\"type\":\"uint256\"}],\"name\":\"SetMinNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumNumValidators\",\"type\":\"uint256\"}],\"name\":\"SetMinNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumNumWatchtowers\",\"type\":\"uint256\"}],\"name\":\"SetMinNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPercentage\",\"type\":\"uint256\"}],\"name\":\"SetSlashPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"SetStakingMinThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressDisputedSequencerToWatchtower\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressDisputedSequencerToWatchtowerExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressDisputedWatchtowerToSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressDisputedWatchtowerToSequencerExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToDisputedWatchtowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsDisputedWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsSequencerInProbationAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToNodeType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToParticipantIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToSequencerInProbationIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToSequencerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToSequencerToDisputedWatchtower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToValidatorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToWatchtowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_dispute_watchtowers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minStakingThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_sequencers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_sequencers_in_probation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_slashPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_watchtowers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeType\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260405180606001604052806040518060400160405280600981526020017f73657175656e636572000000000000000000000000000000000000000000000081525081526020016040518060400160405280600a81526020017f7761746368746f7765720000000000000000000000000000000000000000000081525081526020016040518060400160405280600981526020017f76616c696461746f7200000000000000000000000000000000000000000000008152508152506000906003620000d092919062000160565b50348015620000de57600080fd5b506040516200633d3803806200633d83398181016040528101906200010491906200028d565b808211156200014a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001419062000381565b60405180910390fd5b8160098190555080600a819055505050620006fa565b828054828255906000526020600020908101928215620001ad579160200282015b82811115620001ac5782518290816200019b919062000613565b509160200191906001019062000181565b5b509050620001bc9190620001c0565b5090565b5b80821115620001e45760008181620001da9190620001e8565b50600101620001c1565b5090565b508054620001f6906200040c565b6000825580601f106200020a57506200022b565b601f0160209004906000526020600020908101906200022a91906200022e565b5b50565b5b80821115620002495760008160009055506001016200022f565b5090565b600080fd5b6000819050919050565b620002678162000252565b81146200027357600080fd5b50565b60008151905062000287816200025c565b92915050565b60008060408385031215620002a757620002a66200024d565b5b6000620002b78582860162000276565b9250506020620002ca8582860162000276565b9150509250929050565b600082825260208201905092915050565b7f4d696e207061727469636970616e7473206e756d6265722063616e206e6f742060008201527f62652067726561746572207468616e206d6178206e756d206f6620706172746960208201527f636970616e747300000000000000000000000000000000000000000000000000604082015250565b600062000369604783620002d4565b91506200037682620002e5565b606082019050919050565b600060208201905081810360008301526200039c816200035a565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200042557607f821691505b6020821081036200043b576200043a620003dd565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004a57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000466565b620004b1868362000466565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620004f4620004ee620004e88462000252565b620004c9565b62000252565b9050919050565b6000819050919050565b6200051083620004d3565b620005286200051f82620004fb565b84845462000473565b825550505050565b600090565b6200053f62000530565b6200054c81848462000505565b505050565b5b8181101562000574576200056860008262000535565b60018101905062000552565b5050565b601f821115620005c3576200058d8162000441565b620005988462000456565b81016020851015620005a8578190505b620005c0620005b78562000456565b83018262000551565b50505b505050565b600082821c905092915050565b6000620005e860001984600802620005c8565b1980831691505092915050565b6000620006038383620005d5565b9150826002028217905092915050565b6200061e82620003a3565b67ffffffffffffffff8111156200063a5762000639620003ae565b5b6200064682546200040c565b6200065382828562000578565b600060209050601f8311600181146200068b576000841562000676578287015190505b620006828582620005f5565b865550620006f2565b601f1984166200069b8662000441565b60005b82811015620006c5578489015182556001820191506020850194506020810190506200069e565b86831015620006e55784890151620006e1601f891682620005d5565b8355505b6001600288020188555050505b505050505050565b615c33806200070a6000396000f3fe6080604052600436106104895760003560e01c8063a145012611610255578063c70966fb11610144578063dea79774116100c1578063eec9f52e11610085578063eec9f52e146113bf578063f196f824146113ea578063f76a1ee014611427578063f90ecacc14611452578063f9a43d4c1461148f578063fa70f5e5146114cc57610489565b8063dea79774146112b4578063e028aadf146112f1578063e2b5769d1461132e578063e387a7ed1461136b578063ebe4ad6f1461139657610489565b8063d147c38711610108578063d147c387146111a7578063d3be81a8146111e4578063d551832f1461120f578063d79f10d21461123a578063d80ff2b91461127757610489565b8063c70966fb146110c0578063c795c077146110fd578063c8fa5fe914611128578063c96be4cb14611153578063d14055171461117c57610489565b8063b8889f02116101d2578063c2ccacdf11610196578063c2ccacdf14610fd7578063c36352e114611002578063c402d43c1461102d578063c48c5f3714611058578063c6199c1f1461109557610489565b8063b8889f0214610eca578063bb5a26ff14610f07578063bd9ed0b314610f44578063be31929a14610f6f578063bfc5aeee14610fac57610489565b8063b17ca89c11610219578063b17ca89c14610dbd578063b1da3b1314610de8578063b41327ff14610e25578063b61b127f14610e62578063b737cd9e14610e9f57610489565b8063a145012614610cc2578063a17e8b3d14610ced578063aa90aa6614610d2a578063af6da36e14610d67578063aff1ea7614610d9257610489565b806346f45b8d1161037c5780637b339263116102f957806398abb25e116102bd57806398abb25e14610b8a57806398ff822d14610bc757806399a7c7e414610c045780639a083e8214610c2f5780639f9f710114610c5a578063a0f14bff14610c8557610489565b80637b33926314610a8f5780637dceceb814610acc578063855598b714610b095780638d08414814610b3457806390e7c86214610b5f57610489565b8063584060bb11610340578063584060bb146109965780636b2f7635146109d35780636edf2481146109fc57806372c8af2d14610a27578063760118b614610a6457610489565b806346f45b8d146108aa57806351184694146108c6578063560854f6146108f157806357a87ffd1461092e5780635804e51c1461095957610489565b80631f2743b61161040a5780632def6620116103ce5780632def6620146107d75780633603da84146107ee578063409b9ece1461081957806340b007531461084457806342ca58d51461088157610489565b80631f2743b6146106cc57806324208990146106f757806325d9b8c0146107345780632b3cfe0e146107715780632c404e9c146107ae57610489565b806319908df81161045157806319908df81461059b57806319ae6764146105d85780631b27dbc3146106155780631cb9bdca146106525780631efcf4d11461068f57610489565b806302b751991461048e578063065ae171146104cb578063068946e8146105085780630e533c8e146105335780630fb18ccc14610570575b600080fd5b34801561049a57600080fd5b506104b560048036038101906104b09190614a83565b611509565b6040516104c29190614ac9565b60405180910390f35b3480156104d757600080fd5b506104f260048036038101906104ed9190614a83565b611521565b6040516104ff9190614aff565b60405180910390f35b34801561051457600080fd5b5061051d611541565b60405161052a9190614ac9565b60405180910390f35b34801561053f57600080fd5b5061055a60048036038101906105559190614a83565b611547565b6040516105679190614ac9565b60405180910390f35b34801561057c57600080fd5b5061058561155f565b6040516105929190614ac9565b60405180910390f35b3480156105a757600080fd5b506105c260048036038101906105bd9190614b46565b611565565b6040516105cf9190614b82565b60405180910390f35b3480156105e457600080fd5b506105ff60048036038101906105fa9190614a83565b6115a4565b60405161060c9190614c36565b60405180910390f35b34801561062157600080fd5b5061063c60048036038101906106379190614b46565b611644565b6040516106499190614b82565b60405180910390f35b34801561065e57600080fd5b5061067960048036038101906106749190614a83565b611683565b6040516106869190614ac9565b60405180910390f35b34801561069b57600080fd5b506106b660048036038101906106b19190614a83565b61169b565b6040516106c39190614aff565b60405180910390f35b3480156106d857600080fd5b506106e16116bb565b6040516106ee9190614d64565b60405180910390f35b34801561070357600080fd5b5061071e60048036038101906107199190614a83565b611794565b60405161072b9190614b82565b60405180910390f35b34801561074057600080fd5b5061075b60048036038101906107569190614a83565b6117fd565b6040516107689190614ac9565b60405180910390f35b34801561077d57600080fd5b5061079860048036038101906107939190614a83565b611815565b6040516107a59190614ac9565b60405180910390f35b3480156107ba57600080fd5b506107d560048036038101906107d09190614b46565b61182d565b005b3480156107e357600080fd5b506107ec611896565b005b3480156107fa57600080fd5b50610803611981565b6040516108109190614e44565b60405180910390f35b34801561082557600080fd5b5061082e611a0f565b60405161083b9190614ac9565b60405180910390f35b34801561085057600080fd5b5061086b60048036038101906108669190614a83565b611a1e565b6040516108789190614aff565b60405180910390f35b34801561088d57600080fd5b506108a860048036038101906108a39190614b46565b611a3e565b005b6108c460048036038101906108bf9190614f9b565b611aa7565b005b3480156108d257600080fd5b506108db611b12565b6040516108e89190614ac9565b60405180910390f35b3480156108fd57600080fd5b5061091860048036038101906109139190614b46565b611b17565b6040516109259190614ac9565b60405180910390f35b34801561093a57600080fd5b50610943611b2a565b6040516109509190614ac9565b60405180910390f35b34801561096557600080fd5b50610980600480360381019061097b9190614a83565b611b39565b60405161098d9190614b82565b60405180910390f35b3480156109a257600080fd5b506109bd60048036038101906109b89190614b46565b611b6c565b6040516109ca9190614b82565b60405180910390f35b3480156109df57600080fd5b506109fa60048036038101906109f59190614a83565b611bab565b005b348015610a0857600080fd5b50610a11611c0d565b604051610a1e9190614ac9565b60405180910390f35b348015610a3357600080fd5b50610a4e6004803603810190610a499190614a83565b611c17565b604051610a5b9190614aff565b60405180910390f35b348015610a7057600080fd5b50610a79611c37565b604051610a869190614ac9565b60405180910390f35b348015610a9b57600080fd5b50610ab66004803603810190610ab19190614a83565b611c41565b604051610ac39190614aff565b60405180910390f35b348015610ad857600080fd5b50610af36004803603810190610aee9190614a83565b611c61565b604051610b009190614ac9565b60405180910390f35b348015610b1557600080fd5b50610b1e611c79565b604051610b2b9190614c36565b60405180910390f35b348015610b4057600080fd5b50610b49611cb2565b604051610b569190614ac9565b60405180910390f35b348015610b6b57600080fd5b50610b74611cb8565b604051610b819190614ac9565b60405180910390f35b348015610b9657600080fd5b50610bb16004803603810190610bac9190614b46565b611cbe565b604051610bbe9190614ac9565b60405180910390f35b348015610bd357600080fd5b50610bee6004803603810190610be99190614a83565b611cd1565b604051610bfb9190614aff565b60405180910390f35b348015610c1057600080fd5b50610c19611d27565b604051610c269190614e44565b60405180910390f35b348015610c3b57600080fd5b50610c44611db5565b604051610c519190614c36565b60405180910390f35b348015610c6657600080fd5b50610c6f611dee565b604051610c7c9190614c36565b60405180910390f35b348015610c9157600080fd5b50610cac6004803603810190610ca79190614a83565b611e27565b604051610cb99190614ac9565b60405180910390f35b348015610cce57600080fd5b50610cd7611e3f565b604051610ce49190614ac9565b60405180910390f35b348015610cf957600080fd5b50610d146004803603810190610d0f9190614a83565b611e45565b604051610d219190614aff565b60405180910390f35b348015610d3657600080fd5b50610d516004803603810190610d4c9190614a83565b611e57565b604051610d5e9190614b82565b60405180910390f35b348015610d7357600080fd5b50610d7c611ec0565b604051610d899190614ac9565b60405180910390f35b348015610d9e57600080fd5b50610da7611ec6565b604051610db49190614ac9565b60405180910390f35b348015610dc957600080fd5b50610dd2611ed2565b604051610ddf9190614ac9565b60405180910390f35b348015610df457600080fd5b50610e0f6004803603810190610e0a9190614a83565b611edc565b604051610e1c9190614b82565b60405180910390f35b348015610e3157600080fd5b50610e4c6004803603810190610e479190614b46565b611f0e565b604051610e599190614c36565b60405180910390f35b348015610e6e57600080fd5b50610e896004803603810190610e849190614b46565b611fba565b604051610e969190614ac9565b60405180910390f35b348015610eab57600080fd5b50610eb4611fcd565b604051610ec19190614e44565b60405180910390f35b348015610ed657600080fd5b50610ef16004803603810190610eec9190614a83565b61205b565b604051610efe9190614ac9565b60405180910390f35b348015610f1357600080fd5b50610f2e6004803603810190610f299190614a83565b612073565b604051610f3b9190614aff565b60405180910390f35b348015610f5057600080fd5b50610f596120c9565b604051610f669190614ac9565b60405180910390f35b348015610f7b57600080fd5b50610f966004803603810190610f919190614a83565b6120d3565b604051610fa39190614aff565b60405180910390f35b348015610fb857600080fd5b50610fc16120f3565b604051610fce9190614ac9565b60405180910390f35b348015610fe357600080fd5b50610fec6120fd565b604051610ff99190614ac9565b60405180910390f35b34801561100e57600080fd5b50611017612107565b6040516110249190614e44565b60405180910390f35b34801561103957600080fd5b50611042612195565b60405161104f9190614e44565b60405180910390f35b34801561106457600080fd5b5061107f600480360381019061107a9190614b46565b612223565b60405161108c9190614b82565b60405180910390f35b3480156110a157600080fd5b506110aa612262565b6040516110b79190614ac9565b60405180910390f35b3480156110cc57600080fd5b506110e760048036038101906110e29190614a83565b612268565b6040516110f49190614aff565b60405180910390f35b34801561110957600080fd5b50611112612288565b60405161111f9190614ac9565b60405180910390f35b34801561113457600080fd5b5061113d61228e565b60405161114a9190614ac9565b60405180910390f35b34801561115f57600080fd5b5061117a60048036038101906111759190614a83565b612298565b005b34801561118857600080fd5b50611191612385565b60405161119e9190614e44565b60405180910390f35b3480156111b357600080fd5b506111ce60048036038101906111c99190614b46565b612413565b6040516111db9190614ac9565b60405180910390f35b3480156111f057600080fd5b506111f9612426565b6040516112069190614ac9565b60405180910390f35b34801561121b57600080fd5b50611224612430565b6040516112319190614ac9565b60405180910390f35b34801561124657600080fd5b50611261600480360381019061125c9190614a83565b612436565b60405161126e9190614ac9565b60405180910390f35b34801561128357600080fd5b5061129e60048036038101906112999190614b46565b61247f565b6040516112ab9190614ac9565b60405180910390f35b3480156112c057600080fd5b506112db60048036038101906112d69190614b46565b612492565b6040516112e89190614ac9565b60405180910390f35b3480156112fd57600080fd5b5061131860048036038101906113139190614b46565b6124a5565b6040516113259190614ac9565b60405180910390f35b34801561133a57600080fd5b5061135560048036038101906113509190614a83565b6124b8565b6040516113629190614aff565b60405180910390f35b34801561137757600080fd5b5061138061250e565b60405161138d9190614ac9565b60405180910390f35b3480156113a257600080fd5b506113bd60048036038101906113b89190614a83565b612514565b005b3480156113cb57600080fd5b506113d4612563565b6040516113e19190614ac9565b60405180910390f35b3480156113f657600080fd5b50611411600480360381019061140c9190614b46565b61256d565b60405161141e9190614b82565b60405180910390f35b34801561143357600080fd5b5061143c6125ac565b6040516114499190614ac9565b60405180910390f35b34801561145e57600080fd5b5061147960048036038101906114749190614b46565b6125b2565b6040516114869190614b82565b60405180910390f35b34801561149b57600080fd5b506114b660048036038101906114b19190614b46565b6125f1565b6040516114c39190614ac9565b60405180910390f35b3480156114d857600080fd5b506114f360048036038101906114ee9190614a83565b612604565b6040516115009190614aff565b60405180910390f35b60246020528060005260406000206000915090505481565b60236020528060005260406000206000915054906101000a900460ff1681565b600a5481565b60166020528060005260406000206000915090505481565b600e5481565b6011818154811061157557600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760205280600052604060002060009150905080546115c390615013565b80601f01602080910402602001604051908101604052809291908181526020018280546115ef90615013565b801561163c5780601f106116115761010080835404028352916020019161163c565b820191906000526020600020905b81548152906001019060200180831161161f57829003601f168201915b505050505081565b601a818154811061165457600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60136020528060005260406000206000915090505481565b60156020528060005260406000206000915054906101000a900460ff1681565b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561178b5783829060005260206000200180546116fe90615013565b80601f016020809104026020016040519081016040528092919081815260200182805461172a90615013565b80156117775780601f1061174c57610100808354040283529160200191611777565b820191906000526020600020905b81548152906001019060200180831161175a57829003601f168201915b5050505050815260200190600101906116df565b50505050905090565b6000602060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b601c6020528060005260406000206000915090505481565b601d6020528060005260406000206000915090505481565b61184c3373ffffffffffffffffffffffffffffffffffffffff16612624565b1561188c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161188390615090565b60405180910390fd5b8060018190555050565b6118b53373ffffffffffffffffffffffffffffffffffffffff16612624565b156118f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118ec90615090565b60405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411611977576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196e906150fc565b60405180910390fd5b61197f612647565b565b60606017805480602002602001604051908101604052809291908181526020018280548015611a0557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116119bb575b5050505050905090565b6000611a1961281a565b905090565b60186020528060005260406000206000915054906101000a900460ff1681565b611a5d3373ffffffffffffffffffffffffffffffffffffffff16612624565b15611a9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a9490615090565b60405180910390fd5b8060028190555050565b611ac63373ffffffffffffffffffffffffffffffffffffffff16612624565b15611b06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611afd90615090565b60405180910390fd5b611b0f8161283f565b50565b600181565b600081600c81905550600c549050919050565b6000611b34612ae7565b905090565b601e6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60148181548110611b7c57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611bb481612b04565b611bbd33612c0b565b611bc73382612d12565b8073ffffffffffffffffffffffffffffffffffffffff167f14ebe573ff6e00d7293f5d831eb94cade703155cd0910eb5316a7b2f17b14c3560405160405180910390a250565b6000600f54905090565b60216020528060005260406000206000915054906101000a900460ff1681565b6000600b54905090565b60126020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915090505481565b6040518060400160405280600981526020017f73657175656e636572000000000000000000000000000000000000000000000081525081565b600b5481565b60105481565b6000816009819055506009549050919050565b6000602360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60606011805480602002602001604051908101604052809291908181526020018280548015611dab57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611d61575b5050505050905090565b6040518060400160405280600981526020017f76616c696461746f72000000000000000000000000000000000000000000000081525081565b6040518060400160405280600a81526020017f7761746368746f7765720000000000000000000000000000000000000000000081525081565b60196020528060005260406000206000915090505481565b60015481565b6000611e5082612ec2565b9050919050565b6000601e60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600f5481565b670de0b6b3a764000081565b6000601054905090565b602080528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008181548110611f1e57600080fd5b906000526020600020016000915090508054611f3990615013565b80601f0160208091040260200160405190810160405280929190818152602001828054611f6590615013565b8015611fb25780601f10611f8757610100808354040283529160200191611fb2565b820191906000526020600020905b815481529060010190602001808311611f9557829003601f168201915b505050505081565b600081600b81905550600b549050919050565b6060601a80548060200260200160405190810160405280929190818152602001828054801561205157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612007575b5050505050905090565b60066020528060005260406000206000915090505481565b6000601260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000600d54905090565b601f6020528060005260406000206000915054906101000a900460ff1681565b6000600954905090565b6000600e54905090565b6060602280548060200260200160405190810160405280929190818152602001828054801561218b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612141575b5050505050905090565b6060601480548060200260200160405190810160405280929190818152602001828054801561221957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116121cf575b5050505050905090565b6017818154811061223357600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60046020528060005260406000206000915054906101000a900460ff1681565b600c5481565b6000600a54905090565b6122b73373ffffffffffffffffffffffffffffffffffffffff16612624565b156122f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122ee90615090565b60405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411612379576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612370906150fc565b60405180910390fd5b61238281612f18565b50565b6060600380548060200260200160405190810160405280929190818152602001828054801561240957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116123bf575b5050505050905090565b6000816010819055506010549050919050565b6000600854905090565b60025481565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600081600d81905550600d549050919050565b600081600f81905550600f549050919050565b600081600a81905550600a549050919050565b6000601860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60085481565b61251d81613277565b8073ffffffffffffffffffffffffffffffffffffffff167f7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c0660405160405180910390a250565b6000600c54905090565b6003818154811061257d57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600d5481565b602281815481106125c257600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081600e81905550600e549050919050565b601b6020528060005260406000206000915054906101000a900460ff1681565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b61265033613575565b61268f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612686906151b4565b60405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806008600082825461272a9190615203565b9250508190555061273a336135cb565b612743336138ca565b156127525761275133613920565b5b61275b33613bd6565b1561276a5761276933613c2c565b5b61277333613ee2565b156127825761278133613f38565b5b3373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156127c8573d6000803e3d6000fd5b503373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f758260405161280f9190614ac9565b60405180910390a250565b600080600154111561283057600154905061283c565b670de0b6b3a764000090505b90565b612848816141ee565b8061285857506128578161428b565b5b80612868575061286781614328565b5b6128a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161289e906152a9565b60405180910390fd5b6128af61281a565b3410156128f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128e890615361565b60405180910390fd5b6128fb33826143c5565b61293a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612931906153f3565b60405180910390fd5b612944338261443e565b612983576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161297a90615485565b60405180910390fd5b61298d33826144b7565b6129cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129c390615517565b60405180910390fd5b34600860008282546129de9190615537565b9250508190555034600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612a349190615537565b92505081905550612a4433614530565b612a4d816141ee565b15612a6057612a5b33614689565b612a96565b612a698161428b565b15612a7c57612a7733614790565b612a95565b612a8581614328565b15612a9457612a9333614897565b5b5b5b3373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d34604051612adc9190614ac9565b60405180910390a250565b60008060025411612afb5760019050612b01565b60025490505b90565b6001601560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601480549050601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506014819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6001601b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601a80549050601c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550601a819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80601e60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001601f60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555081602060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001602160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000601560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000612f7b826064612f6c612ae7565b612f76919061558d565b61499e565b9050600060011515602160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615150361303e57601e60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506130fa565b60011515602160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515036130f957602060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b5b600082846131089190615203565b905080600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555082600860008282546131609190615203565b925050819055508173ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f193505050501580156131ad573d6000803e3d6000fd5b508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f10cc080888f97c96419786385100283aa8ae09efd09b75964b955212f42847d4838660405161320d9291906155e7565b60405180910390a361321e85612ec2565b156132705761322c85613277565b8473ffffffffffffffffffffffffffffffffffffffff167f7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c0660405160405180910390a25b5050505050565b61328081612ec2565b6132bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132b6906156a8565b60405180910390fd5b601480549050601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410613345576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161333c9061573a565b60405180910390fd5b6000601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600160148054905061339d9190615203565b905080821461348c576000601482815481106133bc576133bb61575a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080601484815481106133fe576133fd61575a565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000601560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000601660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550601480548061353b5761353a615789565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60095460038054905011613614576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161360b90615850565b60405180910390fd5b600380549050600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541061369a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613691906158e2565b60405180910390fd5b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600060016003805490506136f29190615203565b90508082146137e1576000600382815481106137115761371061575a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080600384815481106137535761375261575a565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060038054806138905761388f615789565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000601260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b601180549050601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054106139a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161399d90615974565b60405180910390fd5b6000601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600060016011805490506139fe9190615203565b9050808214613aed57600060118281548110613a1d57613a1c61575a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060118481548110613a5f57613a5e61575a565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000601360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506011805480613b9c57613b9b615789565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000601860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b601780549050601960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410613cb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ca990615a06565b60405180910390fd5b6000601960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060006001601780549050613d0a9190615203565b9050808214613df957600060178281548110613d2957613d2861575a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060178481548110613d6b57613d6a61575a565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082601960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000601860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000601960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506017805480613ea857613ea7615789565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000602360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b602280549050602460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410613fbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613fb590615a98565b60405180910390fd5b6000602460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600060016022805490506140169190615203565b9050808214614105576000602282815481106140355761403461575a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080602284815481106140775761407661575a565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082602460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000602360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000602460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060228054806141b4576141b3615789565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000816040516020016142019190615af4565b604051602081830303815290604052805190602001206040518060400160405280600981526020017f73657175656e636572000000000000000000000000000000000000000000000081525060405160200161425d9190615af4565b60405160208183030381529060405280519060200120036142815760019050614286565b600090505b919050565b60008160405160200161429e9190615af4565b604051602081830303815290604052805190602001206040518060400160405280600a81526020017f7761746368746f776572000000000000000000000000000000000000000000008152506040516020016142fa9190615af4565b604051602081830303815290604052805190602001200361431e5760019050614323565b600090505b919050565b60008160405160200161433b9190615af4565b604051602081830303815290604052805190602001206040518060400160405280600981526020017f76616c696461746f7200000000000000000000000000000000000000000000008152506040516020016143979190615af4565b60405160208183030381529060405280519060200120036143bb57600190506143c0565b600090505b919050565b60006143d0826141ee565b80156144255750601260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b156144335760009050614438565b600190505b92915050565b60006144498261428b565b801561449e5750601860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b156144ac57600090506144b1565b600190505b92915050565b60006144c282614328565b80156145175750602360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b15614525576000905061452a565b600190505b92915050565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16614686576001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600380549050600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506003819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6001601260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601180549050601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506011819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6001601860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601780549050601960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506017819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6001602360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550602280549050602460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506022819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600061271082846149af919061558d565b10156149f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149e790615b7d565b60405180910390fd5b61271082846149ff919061558d565b614a099190615bcc565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000614a5082614a25565b9050919050565b614a6081614a45565b8114614a6b57600080fd5b50565b600081359050614a7d81614a57565b92915050565b600060208284031215614a9957614a98614a1b565b5b6000614aa784828501614a6e565b91505092915050565b6000819050919050565b614ac381614ab0565b82525050565b6000602082019050614ade6000830184614aba565b92915050565b60008115159050919050565b614af981614ae4565b82525050565b6000602082019050614b146000830184614af0565b92915050565b614b2381614ab0565b8114614b2e57600080fd5b50565b600081359050614b4081614b1a565b92915050565b600060208284031215614b5c57614b5b614a1b565b5b6000614b6a84828501614b31565b91505092915050565b614b7c81614a45565b82525050565b6000602082019050614b976000830184614b73565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015614bd7578082015181840152602081019050614bbc565b83811115614be6576000848401525b50505050565b6000601f19601f8301169050919050565b6000614c0882614b9d565b614c128185614ba8565b9350614c22818560208601614bb9565b614c2b81614bec565b840191505092915050565b60006020820190508181036000830152614c508184614bfd565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b6000614ca082614b9d565b614caa8185614c84565b9350614cba818560208601614bb9565b614cc381614bec565b840191505092915050565b6000614cda8383614c95565b905092915050565b6000602082019050919050565b6000614cfa82614c58565b614d048185614c63565b935083602082028501614d1685614c74565b8060005b85811015614d525784840389528151614d338582614cce565b9450614d3e83614ce2565b925060208a01995050600181019050614d1a565b50829750879550505050505092915050565b60006020820190508181036000830152614d7e8184614cef565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b614dbb81614a45565b82525050565b6000614dcd8383614db2565b60208301905092915050565b6000602082019050919050565b6000614df182614d86565b614dfb8185614d91565b9350614e0683614da2565b8060005b83811015614e37578151614e1e8882614dc1565b9750614e2983614dd9565b925050600181019050614e0a565b5085935050505092915050565b60006020820190508181036000830152614e5e8184614de6565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b614ea882614bec565b810181811067ffffffffffffffff82111715614ec757614ec6614e70565b5b80604052505050565b6000614eda614a11565b9050614ee68282614e9f565b919050565b600067ffffffffffffffff821115614f0657614f05614e70565b5b614f0f82614bec565b9050602081019050919050565b82818337600083830152505050565b6000614f3e614f3984614eeb565b614ed0565b905082815260208101848484011115614f5a57614f59614e6b565b5b614f65848285614f1c565b509392505050565b600082601f830112614f8257614f81614e66565b5b8135614f92848260208601614f2b565b91505092915050565b600060208284031215614fb157614fb0614a1b565b5b600082013567ffffffffffffffff811115614fcf57614fce614a20565b5b614fdb84828501614f6d565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061502b57607f821691505b60208210810361503e5761503d614fe4565b5b50919050565b7f4f6e6c7920454f412063616e2063616c6c2066756e6374696f6e000000000000600082015250565b600061507a601a83614ba8565b915061508582615044565b602082019050919050565b600060208201905081810360008301526150a98161506d565b9050919050565b7f4f6e6c79207374616b65722063616e2063616c6c2066756e6374696f6e000000600082015250565b60006150e6601d83614ba8565b91506150f1826150b0565b602082019050919050565b60006020820190508181036000830152615115816150d9565b9050919050565b7f53656e6465722068617320746f2062652070617274206f66207374616b696e6760008201527f20706f6c6c20696e206f7264657220746f20756e7374616b652069747320736860208201527f6172652e20556e7374616b652072656a65637465642e00000000000000000000604082015250565b600061519e605683614ba8565b91506151a98261511c565b606082019050919050565b600060208201905081810360008301526151cd81615191565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061520e82614ab0565b915061521983614ab0565b92508282101561522c5761522b6151d4565b5b828203905092915050565b7f50726f7669646564206e6f646520747970652068617320746f206d617463682060008201527f617661696c61626c65206e6f6465207479706573000000000000000000000000602082015250565b6000615293603483614ba8565b915061529e82615237565b604082019050919050565b600060208201905081810360008301526152c281615286565b9050919050565b7f496e737566696369656e74207374616b696e6720616d6f756e742070726f766960008201527f6465642e2048617320746f206265206c6172676572207468616e207374616b6960208201527f6e67206d696e696d756d207468726573686f6c64000000000000000000000000604082015250565b600061534b605483614ba8565b9150615356826152c9565b606082019050919050565b6000602082019050818103600083015261537a8161533e565b9050919050565b7f53656e64657220697320616c72656164792073657175656e6365722e2052656a60008201527f656374696e67207374616b6520726571756573742e0000000000000000000000602082015250565b60006153dd603583614ba8565b91506153e882615381565b604082019050919050565b6000602082019050818103600083015261540c816153d0565b9050919050565b7f53656e64657220697320616c7265616479207761746368746f7765722e20526560008201527f6a656374696e67207374616b6520726571756573742e00000000000000000000602082015250565b600061546f603683614ba8565b915061547a82615413565b604082019050919050565b6000602082019050818103600083015261549e81615462565b9050919050565b7f53656e64657220697320616c72656164792076616c696461746f722e2052656a60008201527f656374696e67207374616b6520726571756573742e0000000000000000000000602082015250565b6000615501603583614ba8565b915061550c826154a5565b604082019050919050565b60006020820190508181036000830152615530816154f4565b9050919050565b600061554282614ab0565b915061554d83614ab0565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115615582576155816151d4565b5b828201905092915050565b600061559882614ab0565b91506155a383614ab0565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156155dc576155db6151d4565b5b828202905092915050565b60006040820190506155fc6000830185614aba565b6156096020830184614aba565b9392505050565b7f416464726573732068617320746f20626520696e2070726f626174696f6e206960008201527f6e206f7264657220746f2064656c6574652069742066726f6d2074686520707260208201527f6f626174696f6e2073657175656e63657273206c6973742e0000000000000000604082015250565b6000615692605883614ba8565b915061569d82615610565b606082019050919050565b600060208201905081810360008301526156c181615685565b9050919050565b7f6d616c6963696f7573207061727469636970616e7420696e646578206f75742060008201527f6f662072616e676520696e206d617070696e6700000000000000000000000000602082015250565b6000615724603383614ba8565b915061572f826156c8565b604082019050919050565b6000602082019050818103600083015261575381615717565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f5374616b696e67207061727469636970616e74732063616e2774206265206c6560008201527f7373207468616e20746865206d696e696d756d2072657175697265642070617260208201527f7469636970616e74206e756d0000000000000000000000000000000000000000604082015250565b600061583a604c83614ba8565b9150615845826157b8565b606082019050919050565b600060208201905081810360008301526158698161582d565b9050919050565b7f7061727469636970616e7420696e646578206f7574206f662072616e6765206960008201527f6e206d617070696e670000000000000000000000000000000000000000000000602082015250565b60006158cc602983614ba8565b91506158d782615870565b604082019050919050565b600060208201905081810360008301526158fb816158bf565b9050919050565b7f73657175656e63657220696e646578206f7574206f662072616e676520696e2060008201527f6d617070696e6700000000000000000000000000000000000000000000000000602082015250565b600061595e602783614ba8565b915061596982615902565b604082019050919050565b6000602082019050818103600083015261598d81615951565b9050919050565b7f7761746368746f77657220696e646578206f7574206f662072616e676520696e60008201527f206d617070696e67000000000000000000000000000000000000000000000000602082015250565b60006159f0602883614ba8565b91506159fb82615994565b604082019050919050565b60006020820190508181036000830152615a1f816159e3565b9050919050565b7f76616c696461746f7220696e646578206f7574206f662072616e676520696e2060008201527f6d617070696e6700000000000000000000000000000000000000000000000000602082015250565b6000615a82602783614ba8565b9150615a8d82615a26565b604082019050919050565b60006020820190508181036000830152615ab181615a75565b9050919050565b600081905092915050565b6000615ace82614b9d565b615ad88185615ab8565b9350615ae8818560208601614bb9565b80840191505092915050565b6000615b008284615ac3565b915081905092915050565b7f416d6f756e7420697320746f6f206c6f7720746f2063616c63756c617465206660008201527f6565000000000000000000000000000000000000000000000000000000000000602082015250565b6000615b67602283614ba8565b9150615b7282615b0b565b604082019050919050565b60006020820190508181036000830152615b9681615b5a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000615bd782614ab0565b9150615be283614ab0565b925082615bf257615bf1615b9d565b5b82820490509291505056fea2646970667358221220ed72ae35029160920103a18afc44e955721baca5ab7c785ca012fe7be610ebf964736f6c634300080f0033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, minNumParticipants *big.Int, maxNumParticipants *big.Int) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, minNumParticipants, maxNumParticipants)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// AVAILABLENODETYPES is a free data retrieval call binding the contract method 0xb41327ff.
//
// Solidity: function AVAILABLE_NODE_TYPES(uint256 ) view returns(string)
func (_Staking *StakingCaller) AVAILABLENODETYPES(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "AVAILABLE_NODE_TYPES", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AVAILABLENODETYPES is a free data retrieval call binding the contract method 0xb41327ff.
//
// Solidity: function AVAILABLE_NODE_TYPES(uint256 ) view returns(string)
func (_Staking *StakingSession) AVAILABLENODETYPES(arg0 *big.Int) (string, error) {
	return _Staking.Contract.AVAILABLENODETYPES(&_Staking.CallOpts, arg0)
}

// AVAILABLENODETYPES is a free data retrieval call binding the contract method 0xb41327ff.
//
// Solidity: function AVAILABLE_NODE_TYPES(uint256 ) view returns(string)
func (_Staking *StakingCallerSession) AVAILABLENODETYPES(arg0 *big.Int) (string, error) {
	return _Staking.Contract.AVAILABLENODETYPES(&_Staking.CallOpts, arg0)
}

// DEFAULTMINSLASHPERCENTAGE is a free data retrieval call binding the contract method 0x51184694.
//
// Solidity: function DEFAULT_MIN_SLASH_PERCENTAGE() view returns(uint256)
func (_Staking *StakingCaller) DEFAULTMINSLASHPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEFAULT_MIN_SLASH_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTMINSLASHPERCENTAGE is a free data retrieval call binding the contract method 0x51184694.
//
// Solidity: function DEFAULT_MIN_SLASH_PERCENTAGE() view returns(uint256)
func (_Staking *StakingSession) DEFAULTMINSLASHPERCENTAGE() (*big.Int, error) {
	return _Staking.Contract.DEFAULTMINSLASHPERCENTAGE(&_Staking.CallOpts)
}

// DEFAULTMINSLASHPERCENTAGE is a free data retrieval call binding the contract method 0x51184694.
//
// Solidity: function DEFAULT_MIN_SLASH_PERCENTAGE() view returns(uint256)
func (_Staking *StakingCallerSession) DEFAULTMINSLASHPERCENTAGE() (*big.Int, error) {
	return _Staking.Contract.DEFAULTMINSLASHPERCENTAGE(&_Staking.CallOpts)
}

// DEFAULTSTAKINGTHRESHOLD is a free data retrieval call binding the contract method 0xaff1ea76.
//
// Solidity: function DEFAULT_STAKING_THRESHOLD() view returns(uint256)
func (_Staking *StakingCaller) DEFAULTSTAKINGTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEFAULT_STAKING_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTSTAKINGTHRESHOLD is a free data retrieval call binding the contract method 0xaff1ea76.
//
// Solidity: function DEFAULT_STAKING_THRESHOLD() view returns(uint256)
func (_Staking *StakingSession) DEFAULTSTAKINGTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.DEFAULTSTAKINGTHRESHOLD(&_Staking.CallOpts)
}

// DEFAULTSTAKINGTHRESHOLD is a free data retrieval call binding the contract method 0xaff1ea76.
//
// Solidity: function DEFAULT_STAKING_THRESHOLD() view returns(uint256)
func (_Staking *StakingCallerSession) DEFAULTSTAKINGTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.DEFAULTSTAKINGTHRESHOLD(&_Staking.CallOpts)
}

// GetAvailableNodeTypes is a free data retrieval call binding the contract method 0x1f2743b6.
//
// Solidity: function GetAvailableNodeTypes() view returns(string[])
func (_Staking *StakingCaller) GetAvailableNodeTypes(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetAvailableNodeTypes")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAvailableNodeTypes is a free data retrieval call binding the contract method 0x1f2743b6.
//
// Solidity: function GetAvailableNodeTypes() view returns(string[])
func (_Staking *StakingSession) GetAvailableNodeTypes() ([]string, error) {
	return _Staking.Contract.GetAvailableNodeTypes(&_Staking.CallOpts)
}

// GetAvailableNodeTypes is a free data retrieval call binding the contract method 0x1f2743b6.
//
// Solidity: function GetAvailableNodeTypes() view returns(string[])
func (_Staking *StakingCallerSession) GetAvailableNodeTypes() ([]string, error) {
	return _Staking.Contract.GetAvailableNodeTypes(&_Staking.CallOpts)
}

// GetCurrentAccountStakedAmount is a free data retrieval call binding the contract method 0xd79f10d2.
//
// Solidity: function GetCurrentAccountStakedAmount(address addr) view returns(uint256)
func (_Staking *StakingCaller) GetCurrentAccountStakedAmount(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentAccountStakedAmount", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAccountStakedAmount is a free data retrieval call binding the contract method 0xd79f10d2.
//
// Solidity: function GetCurrentAccountStakedAmount(address addr) view returns(uint256)
func (_Staking *StakingSession) GetCurrentAccountStakedAmount(addr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetCurrentAccountStakedAmount(&_Staking.CallOpts, addr)
}

// GetCurrentAccountStakedAmount is a free data retrieval call binding the contract method 0xd79f10d2.
//
// Solidity: function GetCurrentAccountStakedAmount(address addr) view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentAccountStakedAmount(addr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetCurrentAccountStakedAmount(&_Staking.CallOpts, addr)
}

// GetCurrentDisputeWatchtowers is a free data retrieval call binding the contract method 0xb737cd9e.
//
// Solidity: function GetCurrentDisputeWatchtowers() view returns(address[])
func (_Staking *StakingCaller) GetCurrentDisputeWatchtowers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentDisputeWatchtowers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentDisputeWatchtowers is a free data retrieval call binding the contract method 0xb737cd9e.
//
// Solidity: function GetCurrentDisputeWatchtowers() view returns(address[])
func (_Staking *StakingSession) GetCurrentDisputeWatchtowers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentDisputeWatchtowers(&_Staking.CallOpts)
}

// GetCurrentDisputeWatchtowers is a free data retrieval call binding the contract method 0xb737cd9e.
//
// Solidity: function GetCurrentDisputeWatchtowers() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentDisputeWatchtowers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentDisputeWatchtowers(&_Staking.CallOpts)
}

// GetCurrentParticipants is a free data retrieval call binding the contract method 0xd1405517.
//
// Solidity: function GetCurrentParticipants() view returns(address[])
func (_Staking *StakingCaller) GetCurrentParticipants(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentParticipants")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentParticipants is a free data retrieval call binding the contract method 0xd1405517.
//
// Solidity: function GetCurrentParticipants() view returns(address[])
func (_Staking *StakingSession) GetCurrentParticipants() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentParticipants(&_Staking.CallOpts)
}

// GetCurrentParticipants is a free data retrieval call binding the contract method 0xd1405517.
//
// Solidity: function GetCurrentParticipants() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentParticipants() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentParticipants(&_Staking.CallOpts)
}

// GetCurrentSequencers is a free data retrieval call binding the contract method 0x99a7c7e4.
//
// Solidity: function GetCurrentSequencers() view returns(address[])
func (_Staking *StakingCaller) GetCurrentSequencers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentSequencers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentSequencers is a free data retrieval call binding the contract method 0x99a7c7e4.
//
// Solidity: function GetCurrentSequencers() view returns(address[])
func (_Staking *StakingSession) GetCurrentSequencers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentSequencers(&_Staking.CallOpts)
}

// GetCurrentSequencers is a free data retrieval call binding the contract method 0x99a7c7e4.
//
// Solidity: function GetCurrentSequencers() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentSequencers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentSequencers(&_Staking.CallOpts)
}

// GetCurrentSequencersInProbation is a free data retrieval call binding the contract method 0xc402d43c.
//
// Solidity: function GetCurrentSequencersInProbation() view returns(address[])
func (_Staking *StakingCaller) GetCurrentSequencersInProbation(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentSequencersInProbation")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentSequencersInProbation is a free data retrieval call binding the contract method 0xc402d43c.
//
// Solidity: function GetCurrentSequencersInProbation() view returns(address[])
func (_Staking *StakingSession) GetCurrentSequencersInProbation() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentSequencersInProbation(&_Staking.CallOpts)
}

// GetCurrentSequencersInProbation is a free data retrieval call binding the contract method 0xc402d43c.
//
// Solidity: function GetCurrentSequencersInProbation() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentSequencersInProbation() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentSequencersInProbation(&_Staking.CallOpts)
}

// GetCurrentStakedAmount is a free data retrieval call binding the contract method 0xd3be81a8.
//
// Solidity: function GetCurrentStakedAmount() view returns(uint256)
func (_Staking *StakingCaller) GetCurrentStakedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentStakedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakedAmount is a free data retrieval call binding the contract method 0xd3be81a8.
//
// Solidity: function GetCurrentStakedAmount() view returns(uint256)
func (_Staking *StakingSession) GetCurrentStakedAmount() (*big.Int, error) {
	return _Staking.Contract.GetCurrentStakedAmount(&_Staking.CallOpts)
}

// GetCurrentStakedAmount is a free data retrieval call binding the contract method 0xd3be81a8.
//
// Solidity: function GetCurrentStakedAmount() view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentStakedAmount() (*big.Int, error) {
	return _Staking.Contract.GetCurrentStakedAmount(&_Staking.CallOpts)
}

// GetCurrentStakingThreshold is a free data retrieval call binding the contract method 0x409b9ece.
//
// Solidity: function GetCurrentStakingThreshold() view returns(uint256)
func (_Staking *StakingCaller) GetCurrentStakingThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentStakingThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakingThreshold is a free data retrieval call binding the contract method 0x409b9ece.
//
// Solidity: function GetCurrentStakingThreshold() view returns(uint256)
func (_Staking *StakingSession) GetCurrentStakingThreshold() (*big.Int, error) {
	return _Staking.Contract.GetCurrentStakingThreshold(&_Staking.CallOpts)
}

// GetCurrentStakingThreshold is a free data retrieval call binding the contract method 0x409b9ece.
//
// Solidity: function GetCurrentStakingThreshold() view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentStakingThreshold() (*big.Int, error) {
	return _Staking.Contract.GetCurrentStakingThreshold(&_Staking.CallOpts)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc36352e1.
//
// Solidity: function GetCurrentValidators() view returns(address[])
func (_Staking *StakingCaller) GetCurrentValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc36352e1.
//
// Solidity: function GetCurrentValidators() view returns(address[])
func (_Staking *StakingSession) GetCurrentValidators() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentValidators(&_Staking.CallOpts)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc36352e1.
//
// Solidity: function GetCurrentValidators() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentValidators() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentValidators(&_Staking.CallOpts)
}

// GetCurrentWatchtowers is a free data retrieval call binding the contract method 0x3603da84.
//
// Solidity: function GetCurrentWatchtowers() view returns(address[])
func (_Staking *StakingCaller) GetCurrentWatchtowers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentWatchtowers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentWatchtowers is a free data retrieval call binding the contract method 0x3603da84.
//
// Solidity: function GetCurrentWatchtowers() view returns(address[])
func (_Staking *StakingSession) GetCurrentWatchtowers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentWatchtowers(&_Staking.CallOpts)
}

// GetCurrentWatchtowers is a free data retrieval call binding the contract method 0x3603da84.
//
// Solidity: function GetCurrentWatchtowers() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentWatchtowers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentWatchtowers(&_Staking.CallOpts)
}

// GetDisputedSequencerAddrs is a free data retrieval call binding the contract method 0xaa90aa66.
//
// Solidity: function GetDisputedSequencerAddrs(address watchtowerAddr) view returns(address)
func (_Staking *StakingCaller) GetDisputedSequencerAddrs(opts *bind.CallOpts, watchtowerAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetDisputedSequencerAddrs", watchtowerAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDisputedSequencerAddrs is a free data retrieval call binding the contract method 0xaa90aa66.
//
// Solidity: function GetDisputedSequencerAddrs(address watchtowerAddr) view returns(address)
func (_Staking *StakingSession) GetDisputedSequencerAddrs(watchtowerAddr common.Address) (common.Address, error) {
	return _Staking.Contract.GetDisputedSequencerAddrs(&_Staking.CallOpts, watchtowerAddr)
}

// GetDisputedSequencerAddrs is a free data retrieval call binding the contract method 0xaa90aa66.
//
// Solidity: function GetDisputedSequencerAddrs(address watchtowerAddr) view returns(address)
func (_Staking *StakingCallerSession) GetDisputedSequencerAddrs(watchtowerAddr common.Address) (common.Address, error) {
	return _Staking.Contract.GetDisputedSequencerAddrs(&_Staking.CallOpts, watchtowerAddr)
}

// GetDisputedWatchtowerAddr is a free data retrieval call binding the contract method 0x24208990.
//
// Solidity: function GetDisputedWatchtowerAddr(address sequencerAddr) view returns(address)
func (_Staking *StakingCaller) GetDisputedWatchtowerAddr(opts *bind.CallOpts, sequencerAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetDisputedWatchtowerAddr", sequencerAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDisputedWatchtowerAddr is a free data retrieval call binding the contract method 0x24208990.
//
// Solidity: function GetDisputedWatchtowerAddr(address sequencerAddr) view returns(address)
func (_Staking *StakingSession) GetDisputedWatchtowerAddr(sequencerAddr common.Address) (common.Address, error) {
	return _Staking.Contract.GetDisputedWatchtowerAddr(&_Staking.CallOpts, sequencerAddr)
}

// GetDisputedWatchtowerAddr is a free data retrieval call binding the contract method 0x24208990.
//
// Solidity: function GetDisputedWatchtowerAddr(address sequencerAddr) view returns(address)
func (_Staking *StakingCallerSession) GetDisputedWatchtowerAddr(sequencerAddr common.Address) (common.Address, error) {
	return _Staking.Contract.GetDisputedWatchtowerAddr(&_Staking.CallOpts, sequencerAddr)
}

// GetIsSequencerInProbation is a free data retrieval call binding the contract method 0xa17e8b3d.
//
// Solidity: function GetIsSequencerInProbation(address sequencerAddr) view returns(bool)
func (_Staking *StakingCaller) GetIsSequencerInProbation(opts *bind.CallOpts, sequencerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetIsSequencerInProbation", sequencerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsSequencerInProbation is a free data retrieval call binding the contract method 0xa17e8b3d.
//
// Solidity: function GetIsSequencerInProbation(address sequencerAddr) view returns(bool)
func (_Staking *StakingSession) GetIsSequencerInProbation(sequencerAddr common.Address) (bool, error) {
	return _Staking.Contract.GetIsSequencerInProbation(&_Staking.CallOpts, sequencerAddr)
}

// GetIsSequencerInProbation is a free data retrieval call binding the contract method 0xa17e8b3d.
//
// Solidity: function GetIsSequencerInProbation(address sequencerAddr) view returns(bool)
func (_Staking *StakingCallerSession) GetIsSequencerInProbation(sequencerAddr common.Address) (bool, error) {
	return _Staking.Contract.GetIsSequencerInProbation(&_Staking.CallOpts, sequencerAddr)
}

// GetMaxNumParticipants is a free data retrieval call binding the contract method 0xc8fa5fe9.
//
// Solidity: function GetMaxNumParticipants() view returns(uint256)
func (_Staking *StakingCaller) GetMaxNumParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMaxNumParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumParticipants is a free data retrieval call binding the contract method 0xc8fa5fe9.
//
// Solidity: function GetMaxNumParticipants() view returns(uint256)
func (_Staking *StakingSession) GetMaxNumParticipants() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumParticipants(&_Staking.CallOpts)
}

// GetMaxNumParticipants is a free data retrieval call binding the contract method 0xc8fa5fe9.
//
// Solidity: function GetMaxNumParticipants() view returns(uint256)
func (_Staking *StakingCallerSession) GetMaxNumParticipants() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumParticipants(&_Staking.CallOpts)
}

// GetMaxNumSequencers is a free data retrieval call binding the contract method 0xc2ccacdf.
//
// Solidity: function GetMaxNumSequencers() view returns(uint256)
func (_Staking *StakingCaller) GetMaxNumSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMaxNumSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumSequencers is a free data retrieval call binding the contract method 0xc2ccacdf.
//
// Solidity: function GetMaxNumSequencers() view returns(uint256)
func (_Staking *StakingSession) GetMaxNumSequencers() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumSequencers(&_Staking.CallOpts)
}

// GetMaxNumSequencers is a free data retrieval call binding the contract method 0xc2ccacdf.
//
// Solidity: function GetMaxNumSequencers() view returns(uint256)
func (_Staking *StakingCallerSession) GetMaxNumSequencers() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumSequencers(&_Staking.CallOpts)
}

// GetMaxNumValidators is a free data retrieval call binding the contract method 0x6edf2481.
//
// Solidity: function GetMaxNumValidators() view returns(uint256)
func (_Staking *StakingCaller) GetMaxNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMaxNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumValidators is a free data retrieval call binding the contract method 0x6edf2481.
//
// Solidity: function GetMaxNumValidators() view returns(uint256)
func (_Staking *StakingSession) GetMaxNumValidators() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumValidators(&_Staking.CallOpts)
}

// GetMaxNumValidators is a free data retrieval call binding the contract method 0x6edf2481.
//
// Solidity: function GetMaxNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) GetMaxNumValidators() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumValidators(&_Staking.CallOpts)
}

// GetMaxNumWatchtowers is a free data retrieval call binding the contract method 0xb17ca89c.
//
// Solidity: function GetMaxNumWatchtowers() view returns(uint256)
func (_Staking *StakingCaller) GetMaxNumWatchtowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMaxNumWatchtowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumWatchtowers is a free data retrieval call binding the contract method 0xb17ca89c.
//
// Solidity: function GetMaxNumWatchtowers() view returns(uint256)
func (_Staking *StakingSession) GetMaxNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumWatchtowers(&_Staking.CallOpts)
}

// GetMaxNumWatchtowers is a free data retrieval call binding the contract method 0xb17ca89c.
//
// Solidity: function GetMaxNumWatchtowers() view returns(uint256)
func (_Staking *StakingCallerSession) GetMaxNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumWatchtowers(&_Staking.CallOpts)
}

// GetMinNumParticipants is a free data retrieval call binding the contract method 0xbfc5aeee.
//
// Solidity: function GetMinNumParticipants() view returns(uint256)
func (_Staking *StakingCaller) GetMinNumParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMinNumParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNumParticipants is a free data retrieval call binding the contract method 0xbfc5aeee.
//
// Solidity: function GetMinNumParticipants() view returns(uint256)
func (_Staking *StakingSession) GetMinNumParticipants() (*big.Int, error) {
	return _Staking.Contract.GetMinNumParticipants(&_Staking.CallOpts)
}

// GetMinNumParticipants is a free data retrieval call binding the contract method 0xbfc5aeee.
//
// Solidity: function GetMinNumParticipants() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinNumParticipants() (*big.Int, error) {
	return _Staking.Contract.GetMinNumParticipants(&_Staking.CallOpts)
}

// GetMinNumSequencers is a free data retrieval call binding the contract method 0x760118b6.
//
// Solidity: function GetMinNumSequencers() view returns(uint256)
func (_Staking *StakingCaller) GetMinNumSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMinNumSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNumSequencers is a free data retrieval call binding the contract method 0x760118b6.
//
// Solidity: function GetMinNumSequencers() view returns(uint256)
func (_Staking *StakingSession) GetMinNumSequencers() (*big.Int, error) {
	return _Staking.Contract.GetMinNumSequencers(&_Staking.CallOpts)
}

// GetMinNumSequencers is a free data retrieval call binding the contract method 0x760118b6.
//
// Solidity: function GetMinNumSequencers() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinNumSequencers() (*big.Int, error) {
	return _Staking.Contract.GetMinNumSequencers(&_Staking.CallOpts)
}

// GetMinNumValidators is a free data retrieval call binding the contract method 0xeec9f52e.
//
// Solidity: function GetMinNumValidators() view returns(uint256)
func (_Staking *StakingCaller) GetMinNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMinNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNumValidators is a free data retrieval call binding the contract method 0xeec9f52e.
//
// Solidity: function GetMinNumValidators() view returns(uint256)
func (_Staking *StakingSession) GetMinNumValidators() (*big.Int, error) {
	return _Staking.Contract.GetMinNumValidators(&_Staking.CallOpts)
}

// GetMinNumValidators is a free data retrieval call binding the contract method 0xeec9f52e.
//
// Solidity: function GetMinNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinNumValidators() (*big.Int, error) {
	return _Staking.Contract.GetMinNumValidators(&_Staking.CallOpts)
}

// GetMinNumWatchtowers is a free data retrieval call binding the contract method 0xbd9ed0b3.
//
// Solidity: function GetMinNumWatchtowers() view returns(uint256)
func (_Staking *StakingCaller) GetMinNumWatchtowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMinNumWatchtowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNumWatchtowers is a free data retrieval call binding the contract method 0xbd9ed0b3.
//
// Solidity: function GetMinNumWatchtowers() view returns(uint256)
func (_Staking *StakingSession) GetMinNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.GetMinNumWatchtowers(&_Staking.CallOpts)
}

// GetMinNumWatchtowers is a free data retrieval call binding the contract method 0xbd9ed0b3.
//
// Solidity: function GetMinNumWatchtowers() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.GetMinNumWatchtowers(&_Staking.CallOpts)
}

// GetSlashPercentage is a free data retrieval call binding the contract method 0x57a87ffd.
//
// Solidity: function GetSlashPercentage() view returns(uint256)
func (_Staking *StakingCaller) GetSlashPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetSlashPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashPercentage is a free data retrieval call binding the contract method 0x57a87ffd.
//
// Solidity: function GetSlashPercentage() view returns(uint256)
func (_Staking *StakingSession) GetSlashPercentage() (*big.Int, error) {
	return _Staking.Contract.GetSlashPercentage(&_Staking.CallOpts)
}

// GetSlashPercentage is a free data retrieval call binding the contract method 0x57a87ffd.
//
// Solidity: function GetSlashPercentage() view returns(uint256)
func (_Staking *StakingCallerSession) GetSlashPercentage() (*big.Int, error) {
	return _Staking.Contract.GetSlashPercentage(&_Staking.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0xbb5a26ff.
//
// Solidity: function IsSequencer(address addr) view returns(bool)
func (_Staking *StakingCaller) IsSequencer(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "IsSequencer", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0xbb5a26ff.
//
// Solidity: function IsSequencer(address addr) view returns(bool)
func (_Staking *StakingSession) IsSequencer(addr common.Address) (bool, error) {
	return _Staking.Contract.IsSequencer(&_Staking.CallOpts, addr)
}

// IsSequencer is a free data retrieval call binding the contract method 0xbb5a26ff.
//
// Solidity: function IsSequencer(address addr) view returns(bool)
func (_Staking *StakingCallerSession) IsSequencer(addr common.Address) (bool, error) {
	return _Staking.Contract.IsSequencer(&_Staking.CallOpts, addr)
}

// IsValidator is a free data retrieval call binding the contract method 0x98ff822d.
//
// Solidity: function IsValidator(address addr) view returns(bool)
func (_Staking *StakingCaller) IsValidator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "IsValidator", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0x98ff822d.
//
// Solidity: function IsValidator(address addr) view returns(bool)
func (_Staking *StakingSession) IsValidator(addr common.Address) (bool, error) {
	return _Staking.Contract.IsValidator(&_Staking.CallOpts, addr)
}

// IsValidator is a free data retrieval call binding the contract method 0x98ff822d.
//
// Solidity: function IsValidator(address addr) view returns(bool)
func (_Staking *StakingCallerSession) IsValidator(addr common.Address) (bool, error) {
	return _Staking.Contract.IsValidator(&_Staking.CallOpts, addr)
}

// IsWatchtower is a free data retrieval call binding the contract method 0xe2b5769d.
//
// Solidity: function IsWatchtower(address addr) view returns(bool)
func (_Staking *StakingCaller) IsWatchtower(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "IsWatchtower", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWatchtower is a free data retrieval call binding the contract method 0xe2b5769d.
//
// Solidity: function IsWatchtower(address addr) view returns(bool)
func (_Staking *StakingSession) IsWatchtower(addr common.Address) (bool, error) {
	return _Staking.Contract.IsWatchtower(&_Staking.CallOpts, addr)
}

// IsWatchtower is a free data retrieval call binding the contract method 0xe2b5769d.
//
// Solidity: function IsWatchtower(address addr) view returns(bool)
func (_Staking *StakingCallerSession) IsWatchtower(addr common.Address) (bool, error) {
	return _Staking.Contract.IsWatchtower(&_Staking.CallOpts, addr)
}

// NODESEQUENCER is a free data retrieval call binding the contract method 0x855598b7.
//
// Solidity: function NODE_SEQUENCER() view returns(string)
func (_Staking *StakingCaller) NODESEQUENCER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "NODE_SEQUENCER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NODESEQUENCER is a free data retrieval call binding the contract method 0x855598b7.
//
// Solidity: function NODE_SEQUENCER() view returns(string)
func (_Staking *StakingSession) NODESEQUENCER() (string, error) {
	return _Staking.Contract.NODESEQUENCER(&_Staking.CallOpts)
}

// NODESEQUENCER is a free data retrieval call binding the contract method 0x855598b7.
//
// Solidity: function NODE_SEQUENCER() view returns(string)
func (_Staking *StakingCallerSession) NODESEQUENCER() (string, error) {
	return _Staking.Contract.NODESEQUENCER(&_Staking.CallOpts)
}

// NODEVALIDATOR is a free data retrieval call binding the contract method 0x9a083e82.
//
// Solidity: function NODE_VALIDATOR() view returns(string)
func (_Staking *StakingCaller) NODEVALIDATOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "NODE_VALIDATOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NODEVALIDATOR is a free data retrieval call binding the contract method 0x9a083e82.
//
// Solidity: function NODE_VALIDATOR() view returns(string)
func (_Staking *StakingSession) NODEVALIDATOR() (string, error) {
	return _Staking.Contract.NODEVALIDATOR(&_Staking.CallOpts)
}

// NODEVALIDATOR is a free data retrieval call binding the contract method 0x9a083e82.
//
// Solidity: function NODE_VALIDATOR() view returns(string)
func (_Staking *StakingCallerSession) NODEVALIDATOR() (string, error) {
	return _Staking.Contract.NODEVALIDATOR(&_Staking.CallOpts)
}

// NODEWATCHTOWER is a free data retrieval call binding the contract method 0x9f9f7101.
//
// Solidity: function NODE_WATCHTOWER() view returns(string)
func (_Staking *StakingCaller) NODEWATCHTOWER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "NODE_WATCHTOWER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NODEWATCHTOWER is a free data retrieval call binding the contract method 0x9f9f7101.
//
// Solidity: function NODE_WATCHTOWER() view returns(string)
func (_Staking *StakingSession) NODEWATCHTOWER() (string, error) {
	return _Staking.Contract.NODEWATCHTOWER(&_Staking.CallOpts)
}

// NODEWATCHTOWER is a free data retrieval call binding the contract method 0x9f9f7101.
//
// Solidity: function NODE_WATCHTOWER() view returns(string)
func (_Staking *StakingCallerSession) NODEWATCHTOWER() (string, error) {
	return _Staking.Contract.NODEWATCHTOWER(&_Staking.CallOpts)
}

// AddressDisputedSequencerToWatchtower is a free data retrieval call binding the contract method 0xb1da3b13.
//
// Solidity: function _addressDisputedSequencerToWatchtower(address ) view returns(address)
func (_Staking *StakingCaller) AddressDisputedSequencerToWatchtower(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressDisputedSequencerToWatchtower", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressDisputedSequencerToWatchtower is a free data retrieval call binding the contract method 0xb1da3b13.
//
// Solidity: function _addressDisputedSequencerToWatchtower(address ) view returns(address)
func (_Staking *StakingSession) AddressDisputedSequencerToWatchtower(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.AddressDisputedSequencerToWatchtower(&_Staking.CallOpts, arg0)
}

// AddressDisputedSequencerToWatchtower is a free data retrieval call binding the contract method 0xb1da3b13.
//
// Solidity: function _addressDisputedSequencerToWatchtower(address ) view returns(address)
func (_Staking *StakingCallerSession) AddressDisputedSequencerToWatchtower(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.AddressDisputedSequencerToWatchtower(&_Staking.CallOpts, arg0)
}

// AddressDisputedSequencerToWatchtowerExists is a free data retrieval call binding the contract method 0x72c8af2d.
//
// Solidity: function _addressDisputedSequencerToWatchtowerExists(address ) view returns(bool)
func (_Staking *StakingCaller) AddressDisputedSequencerToWatchtowerExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressDisputedSequencerToWatchtowerExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressDisputedSequencerToWatchtowerExists is a free data retrieval call binding the contract method 0x72c8af2d.
//
// Solidity: function _addressDisputedSequencerToWatchtowerExists(address ) view returns(bool)
func (_Staking *StakingSession) AddressDisputedSequencerToWatchtowerExists(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressDisputedSequencerToWatchtowerExists(&_Staking.CallOpts, arg0)
}

// AddressDisputedSequencerToWatchtowerExists is a free data retrieval call binding the contract method 0x72c8af2d.
//
// Solidity: function _addressDisputedSequencerToWatchtowerExists(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressDisputedSequencerToWatchtowerExists(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressDisputedSequencerToWatchtowerExists(&_Staking.CallOpts, arg0)
}

// AddressDisputedWatchtowerToSequencer is a free data retrieval call binding the contract method 0x5804e51c.
//
// Solidity: function _addressDisputedWatchtowerToSequencer(address ) view returns(address)
func (_Staking *StakingCaller) AddressDisputedWatchtowerToSequencer(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressDisputedWatchtowerToSequencer", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressDisputedWatchtowerToSequencer is a free data retrieval call binding the contract method 0x5804e51c.
//
// Solidity: function _addressDisputedWatchtowerToSequencer(address ) view returns(address)
func (_Staking *StakingSession) AddressDisputedWatchtowerToSequencer(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.AddressDisputedWatchtowerToSequencer(&_Staking.CallOpts, arg0)
}

// AddressDisputedWatchtowerToSequencer is a free data retrieval call binding the contract method 0x5804e51c.
//
// Solidity: function _addressDisputedWatchtowerToSequencer(address ) view returns(address)
func (_Staking *StakingCallerSession) AddressDisputedWatchtowerToSequencer(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.AddressDisputedWatchtowerToSequencer(&_Staking.CallOpts, arg0)
}

// AddressDisputedWatchtowerToSequencerExists is a free data retrieval call binding the contract method 0xbe31929a.
//
// Solidity: function _addressDisputedWatchtowerToSequencerExists(address ) view returns(bool)
func (_Staking *StakingCaller) AddressDisputedWatchtowerToSequencerExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressDisputedWatchtowerToSequencerExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressDisputedWatchtowerToSequencerExists is a free data retrieval call binding the contract method 0xbe31929a.
//
// Solidity: function _addressDisputedWatchtowerToSequencerExists(address ) view returns(bool)
func (_Staking *StakingSession) AddressDisputedWatchtowerToSequencerExists(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressDisputedWatchtowerToSequencerExists(&_Staking.CallOpts, arg0)
}

// AddressDisputedWatchtowerToSequencerExists is a free data retrieval call binding the contract method 0xbe31929a.
//
// Solidity: function _addressDisputedWatchtowerToSequencerExists(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressDisputedWatchtowerToSequencerExists(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressDisputedWatchtowerToSequencerExists(&_Staking.CallOpts, arg0)
}

// AddressToDisputedWatchtowerIndex is a free data retrieval call binding the contract method 0x25d9b8c0.
//
// Solidity: function _addressToDisputedWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToDisputedWatchtowerIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToDisputedWatchtowerIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToDisputedWatchtowerIndex is a free data retrieval call binding the contract method 0x25d9b8c0.
//
// Solidity: function _addressToDisputedWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToDisputedWatchtowerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToDisputedWatchtowerIndex(&_Staking.CallOpts, arg0)
}

// AddressToDisputedWatchtowerIndex is a free data retrieval call binding the contract method 0x25d9b8c0.
//
// Solidity: function _addressToDisputedWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToDisputedWatchtowerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToDisputedWatchtowerIndex(&_Staking.CallOpts, arg0)
}

// AddressToIsDisputedWatchtower is a free data retrieval call binding the contract method 0xfa70f5e5.
//
// Solidity: function _addressToIsDisputedWatchtower(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsDisputedWatchtower(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsDisputedWatchtower", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsDisputedWatchtower is a free data retrieval call binding the contract method 0xfa70f5e5.
//
// Solidity: function _addressToIsDisputedWatchtower(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsDisputedWatchtower(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsDisputedWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToIsDisputedWatchtower is a free data retrieval call binding the contract method 0xfa70f5e5.
//
// Solidity: function _addressToIsDisputedWatchtower(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsDisputedWatchtower(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsDisputedWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToIsParticipant is a free data retrieval call binding the contract method 0xc70966fb.
//
// Solidity: function _addressToIsParticipant(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsParticipant(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsParticipant", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsParticipant is a free data retrieval call binding the contract method 0xc70966fb.
//
// Solidity: function _addressToIsParticipant(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsParticipant(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsParticipant(&_Staking.CallOpts, arg0)
}

// AddressToIsParticipant is a free data retrieval call binding the contract method 0xc70966fb.
//
// Solidity: function _addressToIsParticipant(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsParticipant(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsParticipant(&_Staking.CallOpts, arg0)
}

// AddressToIsSequencer is a free data retrieval call binding the contract method 0x7b339263.
//
// Solidity: function _addressToIsSequencer(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsSequencer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsSequencer is a free data retrieval call binding the contract method 0x7b339263.
//
// Solidity: function _addressToIsSequencer(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsSequencer(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsSequencer(&_Staking.CallOpts, arg0)
}

// AddressToIsSequencer is a free data retrieval call binding the contract method 0x7b339263.
//
// Solidity: function _addressToIsSequencer(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsSequencer(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsSequencer(&_Staking.CallOpts, arg0)
}

// AddressToIsSequencerInProbationAddr is a free data retrieval call binding the contract method 0x1efcf4d1.
//
// Solidity: function _addressToIsSequencerInProbationAddr(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsSequencerInProbationAddr(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsSequencerInProbationAddr", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsSequencerInProbationAddr is a free data retrieval call binding the contract method 0x1efcf4d1.
//
// Solidity: function _addressToIsSequencerInProbationAddr(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsSequencerInProbationAddr(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsSequencerInProbationAddr(&_Staking.CallOpts, arg0)
}

// AddressToIsSequencerInProbationAddr is a free data retrieval call binding the contract method 0x1efcf4d1.
//
// Solidity: function _addressToIsSequencerInProbationAddr(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsSequencerInProbationAddr(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsSequencerInProbationAddr(&_Staking.CallOpts, arg0)
}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsValidator(&_Staking.CallOpts, arg0)
}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsValidator(&_Staking.CallOpts, arg0)
}

// AddressToIsWatchtower is a free data retrieval call binding the contract method 0x40b00753.
//
// Solidity: function _addressToIsWatchtower(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsWatchtower(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsWatchtower", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsWatchtower is a free data retrieval call binding the contract method 0x40b00753.
//
// Solidity: function _addressToIsWatchtower(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsWatchtower(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToIsWatchtower is a free data retrieval call binding the contract method 0x40b00753.
//
// Solidity: function _addressToIsWatchtower(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsWatchtower(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToNodeType is a free data retrieval call binding the contract method 0x19ae6764.
//
// Solidity: function _addressToNodeType(address ) view returns(string)
func (_Staking *StakingCaller) AddressToNodeType(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToNodeType", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressToNodeType is a free data retrieval call binding the contract method 0x19ae6764.
//
// Solidity: function _addressToNodeType(address ) view returns(string)
func (_Staking *StakingSession) AddressToNodeType(arg0 common.Address) (string, error) {
	return _Staking.Contract.AddressToNodeType(&_Staking.CallOpts, arg0)
}

// AddressToNodeType is a free data retrieval call binding the contract method 0x19ae6764.
//
// Solidity: function _addressToNodeType(address ) view returns(string)
func (_Staking *StakingCallerSession) AddressToNodeType(arg0 common.Address) (string, error) {
	return _Staking.Contract.AddressToNodeType(&_Staking.CallOpts, arg0)
}

// AddressToParticipantIndex is a free data retrieval call binding the contract method 0xb8889f02.
//
// Solidity: function _addressToParticipantIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToParticipantIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToParticipantIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToParticipantIndex is a free data retrieval call binding the contract method 0xb8889f02.
//
// Solidity: function _addressToParticipantIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToParticipantIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToParticipantIndex(&_Staking.CallOpts, arg0)
}

// AddressToParticipantIndex is a free data retrieval call binding the contract method 0xb8889f02.
//
// Solidity: function _addressToParticipantIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToParticipantIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToParticipantIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerInProbationIndex is a free data retrieval call binding the contract method 0x0e533c8e.
//
// Solidity: function _addressToSequencerInProbationIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToSequencerInProbationIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToSequencerInProbationIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToSequencerInProbationIndex is a free data retrieval call binding the contract method 0x0e533c8e.
//
// Solidity: function _addressToSequencerInProbationIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToSequencerInProbationIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerInProbationIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerInProbationIndex is a free data retrieval call binding the contract method 0x0e533c8e.
//
// Solidity: function _addressToSequencerInProbationIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToSequencerInProbationIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerInProbationIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerIndex is a free data retrieval call binding the contract method 0x1cb9bdca.
//
// Solidity: function _addressToSequencerIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToSequencerIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToSequencerIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToSequencerIndex is a free data retrieval call binding the contract method 0x1cb9bdca.
//
// Solidity: function _addressToSequencerIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToSequencerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerIndex is a free data retrieval call binding the contract method 0x1cb9bdca.
//
// Solidity: function _addressToSequencerIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToSequencerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerToDisputedWatchtower is a free data retrieval call binding the contract method 0x2b3cfe0e.
//
// Solidity: function _addressToSequencerToDisputedWatchtower(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToSequencerToDisputedWatchtower(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToSequencerToDisputedWatchtower", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToSequencerToDisputedWatchtower is a free data retrieval call binding the contract method 0x2b3cfe0e.
//
// Solidity: function _addressToSequencerToDisputedWatchtower(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToSequencerToDisputedWatchtower(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerToDisputedWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToSequencerToDisputedWatchtower is a free data retrieval call binding the contract method 0x2b3cfe0e.
//
// Solidity: function _addressToSequencerToDisputedWatchtower(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToSequencerToDisputedWatchtower(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerToDisputedWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToStakedAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToStakedAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToStakedAmount(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToStakedAmount(&_Staking.CallOpts, arg0)
}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToStakedAmount(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToStakedAmount(&_Staking.CallOpts, arg0)
}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToValidatorIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToValidatorIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToValidatorIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToValidatorIndex(&_Staking.CallOpts, arg0)
}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToValidatorIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToValidatorIndex(&_Staking.CallOpts, arg0)
}

// AddressToWatchtowerIndex is a free data retrieval call binding the contract method 0xa0f14bff.
//
// Solidity: function _addressToWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToWatchtowerIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToWatchtowerIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToWatchtowerIndex is a free data retrieval call binding the contract method 0xa0f14bff.
//
// Solidity: function _addressToWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToWatchtowerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToWatchtowerIndex(&_Staking.CallOpts, arg0)
}

// AddressToWatchtowerIndex is a free data retrieval call binding the contract method 0xa0f14bff.
//
// Solidity: function _addressToWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToWatchtowerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToWatchtowerIndex(&_Staking.CallOpts, arg0)
}

// DisputeWatchtowers is a free data retrieval call binding the contract method 0x1b27dbc3.
//
// Solidity: function _dispute_watchtowers(uint256 ) view returns(address)
func (_Staking *StakingCaller) DisputeWatchtowers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_dispute_watchtowers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DisputeWatchtowers is a free data retrieval call binding the contract method 0x1b27dbc3.
//
// Solidity: function _dispute_watchtowers(uint256 ) view returns(address)
func (_Staking *StakingSession) DisputeWatchtowers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.DisputeWatchtowers(&_Staking.CallOpts, arg0)
}

// DisputeWatchtowers is a free data retrieval call binding the contract method 0x1b27dbc3.
//
// Solidity: function _dispute_watchtowers(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) DisputeWatchtowers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.DisputeWatchtowers(&_Staking.CallOpts, arg0)
}

// MaximumNumParticipants is a free data retrieval call binding the contract method 0x068946e8.
//
// Solidity: function _maximumNumParticipants() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumParticipants is a free data retrieval call binding the contract method 0x068946e8.
//
// Solidity: function _maximumNumParticipants() view returns(uint256)
func (_Staking *StakingSession) MaximumNumParticipants() (*big.Int, error) {
	return _Staking.Contract.MaximumNumParticipants(&_Staking.CallOpts)
}

// MaximumNumParticipants is a free data retrieval call binding the contract method 0x068946e8.
//
// Solidity: function _maximumNumParticipants() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumParticipants() (*big.Int, error) {
	return _Staking.Contract.MaximumNumParticipants(&_Staking.CallOpts)
}

// MaximumNumSequencers is a free data retrieval call binding the contract method 0x0fb18ccc.
//
// Solidity: function _maximumNumSequencers() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumSequencers is a free data retrieval call binding the contract method 0x0fb18ccc.
//
// Solidity: function _maximumNumSequencers() view returns(uint256)
func (_Staking *StakingSession) MaximumNumSequencers() (*big.Int, error) {
	return _Staking.Contract.MaximumNumSequencers(&_Staking.CallOpts)
}

// MaximumNumSequencers is a free data retrieval call binding the contract method 0x0fb18ccc.
//
// Solidity: function _maximumNumSequencers() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumSequencers() (*big.Int, error) {
	return _Staking.Contract.MaximumNumSequencers(&_Staking.CallOpts)
}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingSession) MaximumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MaximumNumValidators(&_Staking.CallOpts)
}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MaximumNumValidators(&_Staking.CallOpts)
}

// MaximumNumWatchtowers is a free data retrieval call binding the contract method 0x90e7c862.
//
// Solidity: function _maximumNumWatchtowers() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumWatchtowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumWatchtowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumWatchtowers is a free data retrieval call binding the contract method 0x90e7c862.
//
// Solidity: function _maximumNumWatchtowers() view returns(uint256)
func (_Staking *StakingSession) MaximumNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.MaximumNumWatchtowers(&_Staking.CallOpts)
}

// MaximumNumWatchtowers is a free data retrieval call binding the contract method 0x90e7c862.
//
// Solidity: function _maximumNumWatchtowers() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.MaximumNumWatchtowers(&_Staking.CallOpts)
}

// MinStakingThreshold is a free data retrieval call binding the contract method 0xa1450126.
//
// Solidity: function _minStakingThreshold() view returns(uint256)
func (_Staking *StakingCaller) MinStakingThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minStakingThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakingThreshold is a free data retrieval call binding the contract method 0xa1450126.
//
// Solidity: function _minStakingThreshold() view returns(uint256)
func (_Staking *StakingSession) MinStakingThreshold() (*big.Int, error) {
	return _Staking.Contract.MinStakingThreshold(&_Staking.CallOpts)
}

// MinStakingThreshold is a free data retrieval call binding the contract method 0xa1450126.
//
// Solidity: function _minStakingThreshold() view returns(uint256)
func (_Staking *StakingCallerSession) MinStakingThreshold() (*big.Int, error) {
	return _Staking.Contract.MinStakingThreshold(&_Staking.CallOpts)
}

// MinimumNumParticipants is a free data retrieval call binding the contract method 0xc6199c1f.
//
// Solidity: function _minimumNumParticipants() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumParticipants is a free data retrieval call binding the contract method 0xc6199c1f.
//
// Solidity: function _minimumNumParticipants() view returns(uint256)
func (_Staking *StakingSession) MinimumNumParticipants() (*big.Int, error) {
	return _Staking.Contract.MinimumNumParticipants(&_Staking.CallOpts)
}

// MinimumNumParticipants is a free data retrieval call binding the contract method 0xc6199c1f.
//
// Solidity: function _minimumNumParticipants() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumParticipants() (*big.Int, error) {
	return _Staking.Contract.MinimumNumParticipants(&_Staking.CallOpts)
}

// MinimumNumSequencers is a free data retrieval call binding the contract method 0x8d084148.
//
// Solidity: function _minimumNumSequencers() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumSequencers is a free data retrieval call binding the contract method 0x8d084148.
//
// Solidity: function _minimumNumSequencers() view returns(uint256)
func (_Staking *StakingSession) MinimumNumSequencers() (*big.Int, error) {
	return _Staking.Contract.MinimumNumSequencers(&_Staking.CallOpts)
}

// MinimumNumSequencers is a free data retrieval call binding the contract method 0x8d084148.
//
// Solidity: function _minimumNumSequencers() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumSequencers() (*big.Int, error) {
	return _Staking.Contract.MinimumNumSequencers(&_Staking.CallOpts)
}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingSession) MinimumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MinimumNumValidators(&_Staking.CallOpts)
}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MinimumNumValidators(&_Staking.CallOpts)
}

// MinimumNumWatchtowers is a free data retrieval call binding the contract method 0xf76a1ee0.
//
// Solidity: function _minimumNumWatchtowers() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumWatchtowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumWatchtowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumWatchtowers is a free data retrieval call binding the contract method 0xf76a1ee0.
//
// Solidity: function _minimumNumWatchtowers() view returns(uint256)
func (_Staking *StakingSession) MinimumNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.MinimumNumWatchtowers(&_Staking.CallOpts)
}

// MinimumNumWatchtowers is a free data retrieval call binding the contract method 0xf76a1ee0.
//
// Solidity: function _minimumNumWatchtowers() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.MinimumNumWatchtowers(&_Staking.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0xf196f824.
//
// Solidity: function _participants(uint256 ) view returns(address)
func (_Staking *StakingCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0xf196f824.
//
// Solidity: function _participants(uint256 ) view returns(address)
func (_Staking *StakingSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Participants(&_Staking.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0xf196f824.
//
// Solidity: function _participants(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Participants(&_Staking.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x19908df8.
//
// Solidity: function _sequencers(uint256 ) view returns(address)
func (_Staking *StakingCaller) Sequencers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_sequencers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencers is a free data retrieval call binding the contract method 0x19908df8.
//
// Solidity: function _sequencers(uint256 ) view returns(address)
func (_Staking *StakingSession) Sequencers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Sequencers(&_Staking.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x19908df8.
//
// Solidity: function _sequencers(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Sequencers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Sequencers(&_Staking.CallOpts, arg0)
}

// SequencersInProbation is a free data retrieval call binding the contract method 0x584060bb.
//
// Solidity: function _sequencers_in_probation(uint256 ) view returns(address)
func (_Staking *StakingCaller) SequencersInProbation(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_sequencers_in_probation", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencersInProbation is a free data retrieval call binding the contract method 0x584060bb.
//
// Solidity: function _sequencers_in_probation(uint256 ) view returns(address)
func (_Staking *StakingSession) SequencersInProbation(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.SequencersInProbation(&_Staking.CallOpts, arg0)
}

// SequencersInProbation is a free data retrieval call binding the contract method 0x584060bb.
//
// Solidity: function _sequencers_in_probation(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) SequencersInProbation(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.SequencersInProbation(&_Staking.CallOpts, arg0)
}

// SlashPercentage is a free data retrieval call binding the contract method 0xd551832f.
//
// Solidity: function _slashPercentage() view returns(uint256)
func (_Staking *StakingCaller) SlashPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_slashPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashPercentage is a free data retrieval call binding the contract method 0xd551832f.
//
// Solidity: function _slashPercentage() view returns(uint256)
func (_Staking *StakingSession) SlashPercentage() (*big.Int, error) {
	return _Staking.Contract.SlashPercentage(&_Staking.CallOpts)
}

// SlashPercentage is a free data retrieval call binding the contract method 0xd551832f.
//
// Solidity: function _slashPercentage() view returns(uint256)
func (_Staking *StakingCallerSession) SlashPercentage() (*big.Int, error) {
	return _Staking.Contract.SlashPercentage(&_Staking.CallOpts)
}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingCaller) StakedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_stakedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingSession) StakedAmount() (*big.Int, error) {
	return _Staking.Contract.StakedAmount(&_Staking.CallOpts)
}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingCallerSession) StakedAmount() (*big.Int, error) {
	return _Staking.Contract.StakedAmount(&_Staking.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, arg0)
}

// Watchtowers is a free data retrieval call binding the contract method 0xc48c5f37.
//
// Solidity: function _watchtowers(uint256 ) view returns(address)
func (_Staking *StakingCaller) Watchtowers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_watchtowers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Watchtowers is a free data retrieval call binding the contract method 0xc48c5f37.
//
// Solidity: function _watchtowers(uint256 ) view returns(address)
func (_Staking *StakingSession) Watchtowers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Watchtowers(&_Staking.CallOpts, arg0)
}

// Watchtowers is a free data retrieval call binding the contract method 0xc48c5f37.
//
// Solidity: function _watchtowers(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Watchtowers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Watchtowers(&_Staking.CallOpts, arg0)
}

// BeginDisputeResolution is a paid mutator transaction binding the contract method 0x6b2f7635.
//
// Solidity: function BeginDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingTransactor) BeginDisputeResolution(opts *bind.TransactOpts, sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "BeginDisputeResolution", sequencerAddr)
}

// BeginDisputeResolution is a paid mutator transaction binding the contract method 0x6b2f7635.
//
// Solidity: function BeginDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingSession) BeginDisputeResolution(sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.BeginDisputeResolution(&_Staking.TransactOpts, sequencerAddr)
}

// BeginDisputeResolution is a paid mutator transaction binding the contract method 0x6b2f7635.
//
// Solidity: function BeginDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingTransactorSession) BeginDisputeResolution(sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.BeginDisputeResolution(&_Staking.TransactOpts, sequencerAddr)
}

// EndDisputeResolution is a paid mutator transaction binding the contract method 0xebe4ad6f.
//
// Solidity: function EndDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingTransactor) EndDisputeResolution(opts *bind.TransactOpts, sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "EndDisputeResolution", sequencerAddr)
}

// EndDisputeResolution is a paid mutator transaction binding the contract method 0xebe4ad6f.
//
// Solidity: function EndDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingSession) EndDisputeResolution(sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.EndDisputeResolution(&_Staking.TransactOpts, sequencerAddr)
}

// EndDisputeResolution is a paid mutator transaction binding the contract method 0xebe4ad6f.
//
// Solidity: function EndDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingTransactorSession) EndDisputeResolution(sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.EndDisputeResolution(&_Staking.TransactOpts, sequencerAddr)
}

// SetMaxNumParticipants is a paid mutator transaction binding the contract method 0xe028aadf.
//
// Solidity: function SetMaxNumParticipants(uint256 maximumNumParticipants) returns(uint256)
func (_Staking *StakingTransactor) SetMaxNumParticipants(opts *bind.TransactOpts, maximumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMaxNumParticipants", maximumNumParticipants)
}

// SetMaxNumParticipants is a paid mutator transaction binding the contract method 0xe028aadf.
//
// Solidity: function SetMaxNumParticipants(uint256 maximumNumParticipants) returns(uint256)
func (_Staking *StakingSession) SetMaxNumParticipants(maximumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumParticipants(&_Staking.TransactOpts, maximumNumParticipants)
}

// SetMaxNumParticipants is a paid mutator transaction binding the contract method 0xe028aadf.
//
// Solidity: function SetMaxNumParticipants(uint256 maximumNumParticipants) returns(uint256)
func (_Staking *StakingTransactorSession) SetMaxNumParticipants(maximumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumParticipants(&_Staking.TransactOpts, maximumNumParticipants)
}

// SetMaxNumSequencers is a paid mutator transaction binding the contract method 0xf9a43d4c.
//
// Solidity: function SetMaxNumSequencers(uint256 maximumNumSequencers) returns(uint256)
func (_Staking *StakingTransactor) SetMaxNumSequencers(opts *bind.TransactOpts, maximumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMaxNumSequencers", maximumNumSequencers)
}

// SetMaxNumSequencers is a paid mutator transaction binding the contract method 0xf9a43d4c.
//
// Solidity: function SetMaxNumSequencers(uint256 maximumNumSequencers) returns(uint256)
func (_Staking *StakingSession) SetMaxNumSequencers(maximumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumSequencers(&_Staking.TransactOpts, maximumNumSequencers)
}

// SetMaxNumSequencers is a paid mutator transaction binding the contract method 0xf9a43d4c.
//
// Solidity: function SetMaxNumSequencers(uint256 maximumNumSequencers) returns(uint256)
func (_Staking *StakingTransactorSession) SetMaxNumSequencers(maximumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumSequencers(&_Staking.TransactOpts, maximumNumSequencers)
}

// SetMaxNumValidators is a paid mutator transaction binding the contract method 0xdea79774.
//
// Solidity: function SetMaxNumValidators(uint256 maximumNumValidators) returns(uint256)
func (_Staking *StakingTransactor) SetMaxNumValidators(opts *bind.TransactOpts, maximumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMaxNumValidators", maximumNumValidators)
}

// SetMaxNumValidators is a paid mutator transaction binding the contract method 0xdea79774.
//
// Solidity: function SetMaxNumValidators(uint256 maximumNumValidators) returns(uint256)
func (_Staking *StakingSession) SetMaxNumValidators(maximumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumValidators(&_Staking.TransactOpts, maximumNumValidators)
}

// SetMaxNumValidators is a paid mutator transaction binding the contract method 0xdea79774.
//
// Solidity: function SetMaxNumValidators(uint256 maximumNumValidators) returns(uint256)
func (_Staking *StakingTransactorSession) SetMaxNumValidators(maximumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumValidators(&_Staking.TransactOpts, maximumNumValidators)
}

// SetMaxNumWatchtowers is a paid mutator transaction binding the contract method 0xd147c387.
//
// Solidity: function SetMaxNumWatchtowers(uint256 maximumNumWatchtowers) returns(uint256)
func (_Staking *StakingTransactor) SetMaxNumWatchtowers(opts *bind.TransactOpts, maximumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMaxNumWatchtowers", maximumNumWatchtowers)
}

// SetMaxNumWatchtowers is a paid mutator transaction binding the contract method 0xd147c387.
//
// Solidity: function SetMaxNumWatchtowers(uint256 maximumNumWatchtowers) returns(uint256)
func (_Staking *StakingSession) SetMaxNumWatchtowers(maximumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumWatchtowers(&_Staking.TransactOpts, maximumNumWatchtowers)
}

// SetMaxNumWatchtowers is a paid mutator transaction binding the contract method 0xd147c387.
//
// Solidity: function SetMaxNumWatchtowers(uint256 maximumNumWatchtowers) returns(uint256)
func (_Staking *StakingTransactorSession) SetMaxNumWatchtowers(maximumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumWatchtowers(&_Staking.TransactOpts, maximumNumWatchtowers)
}

// SetMinNumParticipants is a paid mutator transaction binding the contract method 0x98abb25e.
//
// Solidity: function SetMinNumParticipants(uint256 minimumNumParticipants) returns(uint256)
func (_Staking *StakingTransactor) SetMinNumParticipants(opts *bind.TransactOpts, minimumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMinNumParticipants", minimumNumParticipants)
}

// SetMinNumParticipants is a paid mutator transaction binding the contract method 0x98abb25e.
//
// Solidity: function SetMinNumParticipants(uint256 minimumNumParticipants) returns(uint256)
func (_Staking *StakingSession) SetMinNumParticipants(minimumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumParticipants(&_Staking.TransactOpts, minimumNumParticipants)
}

// SetMinNumParticipants is a paid mutator transaction binding the contract method 0x98abb25e.
//
// Solidity: function SetMinNumParticipants(uint256 minimumNumParticipants) returns(uint256)
func (_Staking *StakingTransactorSession) SetMinNumParticipants(minimumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumParticipants(&_Staking.TransactOpts, minimumNumParticipants)
}

// SetMinNumSequencers is a paid mutator transaction binding the contract method 0xb61b127f.
//
// Solidity: function SetMinNumSequencers(uint256 minimumNumSequencers) returns(uint256)
func (_Staking *StakingTransactor) SetMinNumSequencers(opts *bind.TransactOpts, minimumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMinNumSequencers", minimumNumSequencers)
}

// SetMinNumSequencers is a paid mutator transaction binding the contract method 0xb61b127f.
//
// Solidity: function SetMinNumSequencers(uint256 minimumNumSequencers) returns(uint256)
func (_Staking *StakingSession) SetMinNumSequencers(minimumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumSequencers(&_Staking.TransactOpts, minimumNumSequencers)
}

// SetMinNumSequencers is a paid mutator transaction binding the contract method 0xb61b127f.
//
// Solidity: function SetMinNumSequencers(uint256 minimumNumSequencers) returns(uint256)
func (_Staking *StakingTransactorSession) SetMinNumSequencers(minimumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumSequencers(&_Staking.TransactOpts, minimumNumSequencers)
}

// SetMinNumValidators is a paid mutator transaction binding the contract method 0x560854f6.
//
// Solidity: function SetMinNumValidators(uint256 minimumNumValidators) returns(uint256)
func (_Staking *StakingTransactor) SetMinNumValidators(opts *bind.TransactOpts, minimumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMinNumValidators", minimumNumValidators)
}

// SetMinNumValidators is a paid mutator transaction binding the contract method 0x560854f6.
//
// Solidity: function SetMinNumValidators(uint256 minimumNumValidators) returns(uint256)
func (_Staking *StakingSession) SetMinNumValidators(minimumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumValidators(&_Staking.TransactOpts, minimumNumValidators)
}

// SetMinNumValidators is a paid mutator transaction binding the contract method 0x560854f6.
//
// Solidity: function SetMinNumValidators(uint256 minimumNumValidators) returns(uint256)
func (_Staking *StakingTransactorSession) SetMinNumValidators(minimumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumValidators(&_Staking.TransactOpts, minimumNumValidators)
}

// SetMinNumWatchtowers is a paid mutator transaction binding the contract method 0xd80ff2b9.
//
// Solidity: function SetMinNumWatchtowers(uint256 minimumNumWatchtowers) returns(uint256)
func (_Staking *StakingTransactor) SetMinNumWatchtowers(opts *bind.TransactOpts, minimumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMinNumWatchtowers", minimumNumWatchtowers)
}

// SetMinNumWatchtowers is a paid mutator transaction binding the contract method 0xd80ff2b9.
//
// Solidity: function SetMinNumWatchtowers(uint256 minimumNumWatchtowers) returns(uint256)
func (_Staking *StakingSession) SetMinNumWatchtowers(minimumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumWatchtowers(&_Staking.TransactOpts, minimumNumWatchtowers)
}

// SetMinNumWatchtowers is a paid mutator transaction binding the contract method 0xd80ff2b9.
//
// Solidity: function SetMinNumWatchtowers(uint256 minimumNumWatchtowers) returns(uint256)
func (_Staking *StakingTransactorSession) SetMinNumWatchtowers(minimumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumWatchtowers(&_Staking.TransactOpts, minimumNumWatchtowers)
}

// SetSlashPercentage is a paid mutator transaction binding the contract method 0x42ca58d5.
//
// Solidity: function SetSlashPercentage(uint256 newPercentage) returns()
func (_Staking *StakingTransactor) SetSlashPercentage(opts *bind.TransactOpts, newPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetSlashPercentage", newPercentage)
}

// SetSlashPercentage is a paid mutator transaction binding the contract method 0x42ca58d5.
//
// Solidity: function SetSlashPercentage(uint256 newPercentage) returns()
func (_Staking *StakingSession) SetSlashPercentage(newPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetSlashPercentage(&_Staking.TransactOpts, newPercentage)
}

// SetSlashPercentage is a paid mutator transaction binding the contract method 0x42ca58d5.
//
// Solidity: function SetSlashPercentage(uint256 newPercentage) returns()
func (_Staking *StakingTransactorSession) SetSlashPercentage(newPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetSlashPercentage(&_Staking.TransactOpts, newPercentage)
}

// SetStakingMinThreshold is a paid mutator transaction binding the contract method 0x2c404e9c.
//
// Solidity: function SetStakingMinThreshold(uint256 newThreshold) returns()
func (_Staking *StakingTransactor) SetStakingMinThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetStakingMinThreshold", newThreshold)
}

// SetStakingMinThreshold is a paid mutator transaction binding the contract method 0x2c404e9c.
//
// Solidity: function SetStakingMinThreshold(uint256 newThreshold) returns()
func (_Staking *StakingSession) SetStakingMinThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetStakingMinThreshold(&_Staking.TransactOpts, newThreshold)
}

// SetStakingMinThreshold is a paid mutator transaction binding the contract method 0x2c404e9c.
//
// Solidity: function SetStakingMinThreshold(uint256 newThreshold) returns()
func (_Staking *StakingTransactorSession) SetStakingMinThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetStakingMinThreshold(&_Staking.TransactOpts, newThreshold)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address slashAddr) returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts, slashAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash", slashAddr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address slashAddr) returns()
func (_Staking *StakingSession) Slash(slashAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, slashAddr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address slashAddr) returns()
func (_Staking *StakingTransactorSession) Slash(slashAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, slashAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string nodeType) payable returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, nodeType string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", nodeType)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string nodeType) payable returns()
func (_Staking *StakingSession) Stake(nodeType string) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, nodeType)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string nodeType) payable returns()
func (_Staking *StakingTransactorSession) Stake(nodeType string) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, nodeType)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactorSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// StakingDisputeResolutionBeganIterator is returned from FilterDisputeResolutionBegan and is used to iterate over the raw logs and unpacked data for DisputeResolutionBegan events raised by the Staking contract.
type StakingDisputeResolutionBeganIterator struct {
	Event *StakingDisputeResolutionBegan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingDisputeResolutionBeganIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDisputeResolutionBegan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingDisputeResolutionBegan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingDisputeResolutionBeganIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDisputeResolutionBeganIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDisputeResolutionBegan represents a DisputeResolutionBegan event raised by the Staking contract.
type StakingDisputeResolutionBegan struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolutionBegan is a free log retrieval operation binding the contract event 0x14ebe573ff6e00d7293f5d831eb94cade703155cd0910eb5316a7b2f17b14c35.
//
// Solidity: event DisputeResolutionBegan(address indexed account)
func (_Staking *StakingFilterer) FilterDisputeResolutionBegan(opts *bind.FilterOpts, account []common.Address) (*StakingDisputeResolutionBeganIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DisputeResolutionBegan", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingDisputeResolutionBeganIterator{contract: _Staking.contract, event: "DisputeResolutionBegan", logs: logs, sub: sub}, nil
}

// WatchDisputeResolutionBegan is a free log subscription operation binding the contract event 0x14ebe573ff6e00d7293f5d831eb94cade703155cd0910eb5316a7b2f17b14c35.
//
// Solidity: event DisputeResolutionBegan(address indexed account)
func (_Staking *StakingFilterer) WatchDisputeResolutionBegan(opts *bind.WatchOpts, sink chan<- *StakingDisputeResolutionBegan, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DisputeResolutionBegan", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDisputeResolutionBegan)
				if err := _Staking.contract.UnpackLog(event, "DisputeResolutionBegan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeResolutionBegan is a log parse operation binding the contract event 0x14ebe573ff6e00d7293f5d831eb94cade703155cd0910eb5316a7b2f17b14c35.
//
// Solidity: event DisputeResolutionBegan(address indexed account)
func (_Staking *StakingFilterer) ParseDisputeResolutionBegan(log types.Log) (*StakingDisputeResolutionBegan, error) {
	event := new(StakingDisputeResolutionBegan)
	if err := _Staking.contract.UnpackLog(event, "DisputeResolutionBegan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDisputeResolutionEndedIterator is returned from FilterDisputeResolutionEnded and is used to iterate over the raw logs and unpacked data for DisputeResolutionEnded events raised by the Staking contract.
type StakingDisputeResolutionEndedIterator struct {
	Event *StakingDisputeResolutionEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingDisputeResolutionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDisputeResolutionEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingDisputeResolutionEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingDisputeResolutionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDisputeResolutionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDisputeResolutionEnded represents a DisputeResolutionEnded event raised by the Staking contract.
type StakingDisputeResolutionEnded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolutionEnded is a free log retrieval operation binding the contract event 0x7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c06.
//
// Solidity: event DisputeResolutionEnded(address indexed account)
func (_Staking *StakingFilterer) FilterDisputeResolutionEnded(opts *bind.FilterOpts, account []common.Address) (*StakingDisputeResolutionEndedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DisputeResolutionEnded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingDisputeResolutionEndedIterator{contract: _Staking.contract, event: "DisputeResolutionEnded", logs: logs, sub: sub}, nil
}

// WatchDisputeResolutionEnded is a free log subscription operation binding the contract event 0x7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c06.
//
// Solidity: event DisputeResolutionEnded(address indexed account)
func (_Staking *StakingFilterer) WatchDisputeResolutionEnded(opts *bind.WatchOpts, sink chan<- *StakingDisputeResolutionEnded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DisputeResolutionEnded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDisputeResolutionEnded)
				if err := _Staking.contract.UnpackLog(event, "DisputeResolutionEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeResolutionEnded is a log parse operation binding the contract event 0x7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c06.
//
// Solidity: event DisputeResolutionEnded(address indexed account)
func (_Staking *StakingFilterer) ParseDisputeResolutionEnded(log types.Log) (*StakingDisputeResolutionEnded, error) {
	event := new(StakingDisputeResolutionEnded)
	if err := _Staking.contract.UnpackLog(event, "DisputeResolutionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Staking contract.
type StakingSlashedIterator struct {
	Event *StakingSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlashed represents a Slashed event raised by the Staking contract.
type StakingSlashed struct {
	Account          common.Address
	NewAmount        *big.Int
	SlashedAmount    *big.Int
	FeeRecipientAddr common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x10cc080888f97c96419786385100283aa8ae09efd09b75964b955212f42847d4.
//
// Solidity: event Slashed(address indexed account, uint256 newAmount, uint256 slashedAmount, address indexed feeRecipientAddr)
func (_Staking *StakingFilterer) FilterSlashed(opts *bind.FilterOpts, account []common.Address, feeRecipientAddr []common.Address) (*StakingSlashedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var feeRecipientAddrRule []interface{}
	for _, feeRecipientAddrItem := range feeRecipientAddr {
		feeRecipientAddrRule = append(feeRecipientAddrRule, feeRecipientAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Slashed", accountRule, feeRecipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlashedIterator{contract: _Staking.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x10cc080888f97c96419786385100283aa8ae09efd09b75964b955212f42847d4.
//
// Solidity: event Slashed(address indexed account, uint256 newAmount, uint256 slashedAmount, address indexed feeRecipientAddr)
func (_Staking *StakingFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakingSlashed, account []common.Address, feeRecipientAddr []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var feeRecipientAddrRule []interface{}
	for _, feeRecipientAddrItem := range feeRecipientAddr {
		feeRecipientAddrRule = append(feeRecipientAddrRule, feeRecipientAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Slashed", accountRule, feeRecipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlashed)
				if err := _Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashed is a log parse operation binding the contract event 0x10cc080888f97c96419786385100283aa8ae09efd09b75964b955212f42847d4.
//
// Solidity: event Slashed(address indexed account, uint256 newAmount, uint256 slashedAmount, address indexed feeRecipientAddr)
func (_Staking *StakingFilterer) ParseSlashed(log types.Log) (*StakingSlashed, error) {
	event := new(StakingSlashed)
	if err := _Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts, account []common.Address) (*StakingStakedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Staking contract.
type StakingUnstakedIterator struct {
	Event *StakingUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstaked represents a Unstaked event raised by the Staking contract.
type StakingUnstaked struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) FilterUnstaked(opts *bind.FilterOpts, account []common.Address) (*StakingUnstakedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unstaked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakedIterator{contract: _Staking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingUnstaked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unstaked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstaked)
				if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) ParseUnstaked(log types.Log) (*StakingUnstaked, error) {
	event := new(StakingUnstaked)
	if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}