// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package identitymanager

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// IdentitymanagerABI is the input ABI used to generate the binding from.
const IdentitymanagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"MNS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameMetaIdentityManager\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"metaID\",\"type\":\"bytes32\"}],\"name\":\"CreateMetaID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"oldMetaID\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"newMetaID\",\"type\":\"bytes32\"}],\"name\":\"UpdateMetaID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"metaID\",\"type\":\"bytes32\"}],\"name\":\"DeleteMetaID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMetadiumNameServiceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_metaID\",\"type\":\"bytes32\"},{\"name\":\"_sig\",\"type\":\"bytes\"},{\"name\":\"_metaPackage\",\"type\":\"bytes\"}],\"name\":\"createMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_metaID\",\"type\":\"bytes32\"},{\"name\":\"_sig\",\"type\":\"bytes\"},{\"name\":\"_metaPackage\",\"type\":\"bytes\"}],\"name\":\"deleteMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldMetaID\",\"type\":\"bytes32\"},{\"name\":\"_newMetaID\",\"type\":\"bytes32\"},{\"name\":\"_sig\",\"type\":\"bytes\"},{\"name\":\"_metaPackage\",\"type\":\"bytes\"}],\"name\":\"updateMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_metaID\",\"type\":\"uint256\"}],\"name\":\"getRecoveryData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"ecrecovery\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ecverify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"getAddressFromMetaPackage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// IdentitymanagerBin is the compiled bytecode used for deploying new contracts.
const IdentitymanagerBin = `6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061259c806100536000396000f3006080604052600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630c1f2e0f146100f65780631ea663b51461014d5780632f745c591461019657806339cdde32146101f75780636352211e146102a6578063702fecd11461031357806370a08231146103e8578063711718221461043f57806376845d261461048257806377d32e94146105575780638da5cb5b1461060e578063ade266fe14610665578063c520a1e014610748578063c87b56dd146107f1578063e9e7b9ea14610897578063eaeaa3d9146108ca578063f2fde38b146108fd575b600080fd5b34801561010257600080fd5b5061010b610940565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561015957600080fd5b5061017860048036038101908080359060200190929190505050610966565b60405180826000191660001916815260200191505060405180910390f35b3480156101a257600080fd5b506101e1600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061096d565b6040518082815260200191505060405180910390f35b34801561020357600080fd5b5061028c6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b4d565b604051808215151515815260200191505060405180910390f35b3480156102b257600080fd5b506102d160048036038101908080359060200190929190505050610c60565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561031f57600080fd5b506103ce6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610e0b565b604051808215151515815260200191505060405180910390f35b3480156103f457600080fd5b50610429600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506112c3565b6040518082815260200191505060405180910390f35b34801561044b57600080fd5b50610480600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061149a565b005b34801561048e57600080fd5b5061053d6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611539565b604051808215151515815260200191505060405180910390f35b34801561056357600080fd5b506105cc6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611979565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561061a57600080fd5b50610623611a71565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561067157600080fd5b5061072e60048036038101908080356000191690602001909291908035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611a96565b604051808215151515815260200191505060405180910390f35b34801561075457600080fd5b506107af600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506120c2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107fd57600080fd5b5061081c600480360381019080803590602001909291905050506121d1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561085c578082015181840152602081019050610841565b50505050905090810190601f1680156108895780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156108a357600080fd5b506108ac6123d3565b60405180826000191660001916815260200191505060405180910390f35b3480156108d657600080fd5b506108df6123f7565b60405180826000191660001916815260200191505060405180910390f35b34801561090957600080fd5b5061093e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061241b565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000919050565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d2020dd7f4d657461494400000000000000000000000000000000000000000000000000006040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015610a2957600080fd5b505af1158015610a3d573d6000803e3d6000fd5b505050506040513d6020811015610a5357600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff16632f745c5985856040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610b0957600080fd5b505af1158015610b1d573d6000803e3d6000fd5b505050506040513d6020811015610b3357600080fd5b810190808051906020019092919050505091505092915050565b6000606084604051808260001916600019168152602001915050604051809103902094506040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250905080856040518083805190602001908083835b602083101515610be05780518252602082019150602081019050602083039250610bbb565b6001836020036101000a03801982511681845116808217855250505050505090500182600019166000191681526020019250505060405180910390209450610c288585611979565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16149150509392505050565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d2020dd7f4d657461494400000000000000000000000000000000000000000000000000006040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015610d1c57600080fd5b505af1158015610d30573d6000803e3d6000fd5b505050506040513d6020811015610d4657600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015610dc857600080fd5b505af1158015610ddc573d6000803e3d6000fd5b505050506040513d6020811015610df257600080fd5b8101908080519060200190929190505050915050919050565b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360d6c7cf7f4d6574616469756d4964656e746974794d616e61676572000000000000000000336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015610efd57600080fd5b505af1158015610f11573d6000803e3d6000fd5b505050506040513d6020811015610f2757600080fd5b81019080805190602001909291905050501515610f4357600080fd5b610f4c846120c2565b9150610f59868684610b4d565b1515610f6457600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d2020dd7f4d657461494400000000000000000000000000000000000000000000000000006040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561101d57600080fd5b505af1158015611031573d6000803e3d6000fd5b505050506040513d602081101561104757600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff16634f558e7987600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156110cd57600080fd5b505af11580156110e1573d6000803e3d6000fd5b505050506040513d60208110156110f757600080fd5b810190808051906020019092919050505015151561111457600080fd5b8073ffffffffffffffffffffffffffffffffffffffff1663d3fc9864838860019004876040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156111da5780820151818401526020810190506111bf565b50505050905090810190601f1680156112075780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561122857600080fd5b505af115801561123c573d6000803e3d6000fd5b505050506040513d602081101561125257600080fd5b8101908080519060200190929190505050151561126e57600080fd5b85600019168273ffffffffffffffffffffffffffffffffffffffff167f0afbed7f307c1b94f44fd781ff78fa8f85831c62848a1c53db8d28f575cf1aa160405160405180910390a36001925050509392505050565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d2020dd7f4d657461494400000000000000000000000000000000000000000000000000006040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561137f57600080fd5b505af1158015611393573d6000803e3d6000fd5b505050506040513d60208110156113a957600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff166370a08231846040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561145757600080fd5b505af115801561146b573d6000803e3d6000fd5b505050506040513d602081101561148157600080fd5b8101908080519060200190929190505050915050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156114f557600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360d6c7cf7f4d6574616469756d4964656e746974794d616e61676572000000000000000000336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b15801561162b57600080fd5b505af115801561163f573d6000803e3d6000fd5b505050506040513d602081101561165557600080fd5b8101908080519060200190929190505050151561167157600080fd5b61167a846120c2565b9150611687868684610b4d565b151561169257600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d2020dd7f4d657461494400000000000000000000000000000000000000000000000000006040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561174b57600080fd5b505af115801561175f573d6000803e3d6000fd5b505050506040513d602081101561177557600080fd5b810190808051906020019092919050505090508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16636352211e88600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561181257600080fd5b505af1158015611826573d6000803e3d6000fd5b505050506040513d602081101561183c57600080fd5b810190808051906020019092919050505073ffffffffffffffffffffffffffffffffffffffff1614151561186f57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166342966c6887600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156118e257600080fd5b505af11580156118f6573d6000803e3d6000fd5b505050506040513d602081101561190c57600080fd5b8101908080519060200190929190505050151561192857600080fd5b85600019168273ffffffffffffffffffffffffffffffffffffffff167f4e8bb03e2af31dac5f75505a69e10ce88e9d1239cc1859e1e79a7d360fbf061b60405160405180910390a350509392505050565b600080600080604185511415156119935760009350611a68565b602085015192506040850151915060ff6041860151169050601b8160ff1610156119be57601b810190505b601b8160ff16141580156119d65750601c8160ff1614155b156119e45760009350611a68565b600186828585604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af1158015611a5b573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360d6c7cf7f4d6574616469756d4964656e746974794d616e61676572000000000000000000336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015611b8857600080fd5b505af1158015611b9c573d6000803e3d6000fd5b505050506040513d6020811015611bb257600080fd5b81019080805190602001909291905050501515611bce57600080fd5b611bd7846120c2565b9150611be4868684610b4d565b1515611bef57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d2020dd7f4d657461494400000000000000000000000000000000000000000000000000006040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015611ca857600080fd5b505af1158015611cbc573d6000803e3d6000fd5b505050506040513d6020811015611cd257600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff16634f558e7988600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015611d5857600080fd5b505af1158015611d6c573d6000803e3d6000fd5b505050506040513d6020811015611d8257600080fd5b81019080805190602001909291905050501515611d9e57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16634f558e7987600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015611e1157600080fd5b505af1158015611e25573d6000803e3d6000fd5b505050506040513d6020811015611e3b57600080fd5b8101908080519060200190929190505050151515611e5857600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166342966c6888600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015611ecb57600080fd5b505af1158015611edf573d6000803e3d6000fd5b505050506040513d6020811015611ef557600080fd5b81019080805190602001909291905050501515611f1157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff1663d3fc9864838860019004876040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611fd7578082015181840152602081019050611fbc565b50505050905090810190601f1680156120045780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561202557600080fd5b505af1158015612039573d6000803e3d6000fd5b505050506040513d602081101561204f57600080fd5b8101908080519060200190929190505050151561206b57600080fd5b856000191687600019168373ffffffffffffffffffffffffffffffffffffffff167fde4d89fa4d58496bf0ceb285c9696695ffef28d55d9c47936ef981f124aa2c8c60405160405180910390a45050949350505050565b600080600080601692508285511115156120db57600080fd5b600190505b60158110156121b6576008600182030260ff7f010000000000000000000000000000000000000000000000000000000000000002868381518110151561212257fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166bffffffffffffffffffffffff19169060020a90048217915080806001019150506120e0565b816c0100000000000000000000000090049350505050919050565b60606000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d2020dd7f4d657461494400000000000000000000000000000000000000000000000000006040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b15801561228e57600080fd5b505af11580156122a2573d6000803e3d6000fd5b505050506040513d60208110156122b857600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff1663c87b56dd846040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561233a57600080fd5b505af115801561234e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561237857600080fd5b81019080805164010000000081111561239057600080fd5b828101905060208101848111156123a657600080fd5b81518560018202830111640100000000821117156123c357600080fd5b5050929190505050915050919050565b7f4d6574616469756d4964656e746974794d616e6167657200000000000000000081565b7f4d6574614944000000000000000000000000000000000000000000000000000081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561247657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156124b257600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058206bdb6a8d061073a46d58048a451946dc8b7eafcaef36e54c8ba2a346cd09a6720029`

// DeployIdentitymanager deploys a new Ethereum contract, binding an instance of Identitymanager to it.
func DeployIdentitymanager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Identitymanager, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentitymanagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentitymanagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Identitymanager{IdentitymanagerCaller: IdentitymanagerCaller{contract: contract}, IdentitymanagerTransactor: IdentitymanagerTransactor{contract: contract}, IdentitymanagerFilterer: IdentitymanagerFilterer{contract: contract}}, nil
}

// Identitymanager is an auto generated Go binding around an Ethereum contract.
type Identitymanager struct {
	IdentitymanagerCaller     // Read-only binding to the contract
	IdentitymanagerTransactor // Write-only binding to the contract
	IdentitymanagerFilterer   // Log filterer for contract events
}

// IdentitymanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentitymanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitymanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentitymanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitymanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentitymanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitymanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitymanagerSession struct {
	Contract     *Identitymanager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentitymanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentitymanagerCallerSession struct {
	Contract *IdentitymanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IdentitymanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentitymanagerTransactorSession struct {
	Contract     *IdentitymanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IdentitymanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentitymanagerRaw struct {
	Contract *Identitymanager // Generic contract binding to access the raw methods on
}

// IdentitymanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentitymanagerCallerRaw struct {
	Contract *IdentitymanagerCaller // Generic read-only contract binding to access the raw methods on
}

// IdentitymanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentitymanagerTransactorRaw struct {
	Contract *IdentitymanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentitymanager creates a new instance of Identitymanager, bound to a specific deployed contract.
func NewIdentitymanager(address common.Address, backend bind.ContractBackend) (*Identitymanager, error) {
	contract, err := bindIdentitymanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identitymanager{IdentitymanagerCaller: IdentitymanagerCaller{contract: contract}, IdentitymanagerTransactor: IdentitymanagerTransactor{contract: contract}, IdentitymanagerFilterer: IdentitymanagerFilterer{contract: contract}}, nil
}

// NewIdentitymanagerCaller creates a new read-only instance of Identitymanager, bound to a specific deployed contract.
func NewIdentitymanagerCaller(address common.Address, caller bind.ContractCaller) (*IdentitymanagerCaller, error) {
	contract, err := bindIdentitymanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentitymanagerCaller{contract: contract}, nil
}

// NewIdentitymanagerTransactor creates a new write-only instance of Identitymanager, bound to a specific deployed contract.
func NewIdentitymanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentitymanagerTransactor, error) {
	contract, err := bindIdentitymanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentitymanagerTransactor{contract: contract}, nil
}

// NewIdentitymanagerFilterer creates a new log filterer instance of Identitymanager, bound to a specific deployed contract.
func NewIdentitymanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentitymanagerFilterer, error) {
	contract, err := bindIdentitymanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentitymanagerFilterer{contract: contract}, nil
}

// bindIdentitymanager binds a generic wrapper to an already deployed contract.
func bindIdentitymanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentitymanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identitymanager *IdentitymanagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identitymanager.Contract.IdentitymanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identitymanager *IdentitymanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identitymanager.Contract.IdentitymanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identitymanager *IdentitymanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identitymanager.Contract.IdentitymanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identitymanager *IdentitymanagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identitymanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identitymanager *IdentitymanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identitymanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identitymanager *IdentitymanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identitymanager.Contract.contract.Transact(opts, method, params...)
}

// MNS is a free data retrieval call binding the contract method 0x0c1f2e0f.
//
// Solidity: function MNS() constant returns(address)
func (_Identitymanager *IdentitymanagerCaller) MNS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "MNS")
	return *ret0, err
}

// MNS is a free data retrieval call binding the contract method 0x0c1f2e0f.
//
// Solidity: function MNS() constant returns(address)
func (_Identitymanager *IdentitymanagerSession) MNS() (common.Address, error) {
	return _Identitymanager.Contract.MNS(&_Identitymanager.CallOpts)
}

// MNS is a free data retrieval call binding the contract method 0x0c1f2e0f.
//
// Solidity: function MNS() constant returns(address)
func (_Identitymanager *IdentitymanagerCallerSession) MNS() (common.Address, error) {
	return _Identitymanager.Contract.MNS(&_Identitymanager.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_Identitymanager *IdentitymanagerCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_Identitymanager *IdentitymanagerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Identitymanager.Contract.BalanceOf(&_Identitymanager.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_Identitymanager *IdentitymanagerCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Identitymanager.Contract.BalanceOf(&_Identitymanager.CallOpts, _owner)
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(hash bytes32, sig bytes) constant returns(address)
func (_Identitymanager *IdentitymanagerCaller) Ecrecovery(opts *bind.CallOpts, hash [32]byte, sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "ecrecovery", hash, sig)
	return *ret0, err
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(hash bytes32, sig bytes) constant returns(address)
func (_Identitymanager *IdentitymanagerSession) Ecrecovery(hash [32]byte, sig []byte) (common.Address, error) {
	return _Identitymanager.Contract.Ecrecovery(&_Identitymanager.CallOpts, hash, sig)
}

// Ecrecovery is a free data retrieval call binding the contract method 0x77d32e94.
//
// Solidity: function ecrecovery(hash bytes32, sig bytes) constant returns(address)
func (_Identitymanager *IdentitymanagerCallerSession) Ecrecovery(hash [32]byte, sig []byte) (common.Address, error) {
	return _Identitymanager.Contract.Ecrecovery(&_Identitymanager.CallOpts, hash, sig)
}

// Ecverify is a free data retrieval call binding the contract method 0x39cdde32.
//
// Solidity: function ecverify(message bytes32, sig bytes, signer address) constant returns(bool)
func (_Identitymanager *IdentitymanagerCaller) Ecverify(opts *bind.CallOpts, message [32]byte, sig []byte, signer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "ecverify", message, sig, signer)
	return *ret0, err
}

// Ecverify is a free data retrieval call binding the contract method 0x39cdde32.
//
// Solidity: function ecverify(message bytes32, sig bytes, signer address) constant returns(bool)
func (_Identitymanager *IdentitymanagerSession) Ecverify(message [32]byte, sig []byte, signer common.Address) (bool, error) {
	return _Identitymanager.Contract.Ecverify(&_Identitymanager.CallOpts, message, sig, signer)
}

// Ecverify is a free data retrieval call binding the contract method 0x39cdde32.
//
// Solidity: function ecverify(message bytes32, sig bytes, signer address) constant returns(bool)
func (_Identitymanager *IdentitymanagerCallerSession) Ecverify(message [32]byte, sig []byte, signer common.Address) (bool, error) {
	return _Identitymanager.Contract.Ecverify(&_Identitymanager.CallOpts, message, sig, signer)
}

// GetAddressFromMetaPackage is a free data retrieval call binding the contract method 0xc520a1e0.
//
// Solidity: function getAddressFromMetaPackage(b bytes) constant returns(address)
func (_Identitymanager *IdentitymanagerCaller) GetAddressFromMetaPackage(opts *bind.CallOpts, b []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "getAddressFromMetaPackage", b)
	return *ret0, err
}

// GetAddressFromMetaPackage is a free data retrieval call binding the contract method 0xc520a1e0.
//
// Solidity: function getAddressFromMetaPackage(b bytes) constant returns(address)
func (_Identitymanager *IdentitymanagerSession) GetAddressFromMetaPackage(b []byte) (common.Address, error) {
	return _Identitymanager.Contract.GetAddressFromMetaPackage(&_Identitymanager.CallOpts, b)
}

// GetAddressFromMetaPackage is a free data retrieval call binding the contract method 0xc520a1e0.
//
// Solidity: function getAddressFromMetaPackage(b bytes) constant returns(address)
func (_Identitymanager *IdentitymanagerCallerSession) GetAddressFromMetaPackage(b []byte) (common.Address, error) {
	return _Identitymanager.Contract.GetAddressFromMetaPackage(&_Identitymanager.CallOpts, b)
}

// GetRecoveryData is a free data retrieval call binding the contract method 0x1ea663b5.
//
// Solidity: function getRecoveryData(_metaID uint256) constant returns(bytes32)
func (_Identitymanager *IdentitymanagerCaller) GetRecoveryData(opts *bind.CallOpts, _metaID *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "getRecoveryData", _metaID)
	return *ret0, err
}

// GetRecoveryData is a free data retrieval call binding the contract method 0x1ea663b5.
//
// Solidity: function getRecoveryData(_metaID uint256) constant returns(bytes32)
func (_Identitymanager *IdentitymanagerSession) GetRecoveryData(_metaID *big.Int) ([32]byte, error) {
	return _Identitymanager.Contract.GetRecoveryData(&_Identitymanager.CallOpts, _metaID)
}

// GetRecoveryData is a free data retrieval call binding the contract method 0x1ea663b5.
//
// Solidity: function getRecoveryData(_metaID uint256) constant returns(bytes32)
func (_Identitymanager *IdentitymanagerCallerSession) GetRecoveryData(_metaID *big.Int) ([32]byte, error) {
	return _Identitymanager.Contract.GetRecoveryData(&_Identitymanager.CallOpts, _metaID)
}

// NameMetaID is a free data retrieval call binding the contract method 0xeaeaa3d9.
//
// Solidity: function nameMetaID() constant returns(bytes32)
func (_Identitymanager *IdentitymanagerCaller) NameMetaID(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "nameMetaID")
	return *ret0, err
}

// NameMetaID is a free data retrieval call binding the contract method 0xeaeaa3d9.
//
// Solidity: function nameMetaID() constant returns(bytes32)
func (_Identitymanager *IdentitymanagerSession) NameMetaID() ([32]byte, error) {
	return _Identitymanager.Contract.NameMetaID(&_Identitymanager.CallOpts)
}

// NameMetaID is a free data retrieval call binding the contract method 0xeaeaa3d9.
//
// Solidity: function nameMetaID() constant returns(bytes32)
func (_Identitymanager *IdentitymanagerCallerSession) NameMetaID() ([32]byte, error) {
	return _Identitymanager.Contract.NameMetaID(&_Identitymanager.CallOpts)
}

// NameMetaIdentityManager is a free data retrieval call binding the contract method 0xe9e7b9ea.
//
// Solidity: function nameMetaIdentityManager() constant returns(bytes32)
func (_Identitymanager *IdentitymanagerCaller) NameMetaIdentityManager(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "nameMetaIdentityManager")
	return *ret0, err
}

// NameMetaIdentityManager is a free data retrieval call binding the contract method 0xe9e7b9ea.
//
// Solidity: function nameMetaIdentityManager() constant returns(bytes32)
func (_Identitymanager *IdentitymanagerSession) NameMetaIdentityManager() ([32]byte, error) {
	return _Identitymanager.Contract.NameMetaIdentityManager(&_Identitymanager.CallOpts)
}

// NameMetaIdentityManager is a free data retrieval call binding the contract method 0xe9e7b9ea.
//
// Solidity: function nameMetaIdentityManager() constant returns(bytes32)
func (_Identitymanager *IdentitymanagerCallerSession) NameMetaIdentityManager() ([32]byte, error) {
	return _Identitymanager.Contract.NameMetaIdentityManager(&_Identitymanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Identitymanager *IdentitymanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Identitymanager *IdentitymanagerSession) Owner() (common.Address, error) {
	return _Identitymanager.Contract.Owner(&_Identitymanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Identitymanager *IdentitymanagerCallerSession) Owner() (common.Address, error) {
	return _Identitymanager.Contract.Owner(&_Identitymanager.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_Identitymanager *IdentitymanagerCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_Identitymanager *IdentitymanagerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Identitymanager.Contract.OwnerOf(&_Identitymanager.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_Identitymanager *IdentitymanagerCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Identitymanager.Contract.OwnerOf(&_Identitymanager.CallOpts, _tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_Identitymanager *IdentitymanagerCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_Identitymanager *IdentitymanagerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Identitymanager.Contract.TokenOfOwnerByIndex(&_Identitymanager.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_Identitymanager *IdentitymanagerCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Identitymanager.Contract.TokenOfOwnerByIndex(&_Identitymanager.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_Identitymanager *IdentitymanagerCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_Identitymanager *IdentitymanagerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Identitymanager.Contract.TokenURI(&_Identitymanager.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_Identitymanager *IdentitymanagerCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Identitymanager.Contract.TokenURI(&_Identitymanager.CallOpts, _tokenId)
}

// CreateMetaID is a paid mutator transaction binding the contract method 0x702fecd1.
//
// Solidity: function createMetaID(_metaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactor) CreateMetaID(opts *bind.TransactOpts, _metaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.contract.Transact(opts, "createMetaID", _metaID, _sig, _metaPackage)
}

// CreateMetaID is a paid mutator transaction binding the contract method 0x702fecd1.
//
// Solidity: function createMetaID(_metaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerSession) CreateMetaID(_metaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.CreateMetaID(&_Identitymanager.TransactOpts, _metaID, _sig, _metaPackage)
}

// CreateMetaID is a paid mutator transaction binding the contract method 0x702fecd1.
//
// Solidity: function createMetaID(_metaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactorSession) CreateMetaID(_metaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.CreateMetaID(&_Identitymanager.TransactOpts, _metaID, _sig, _metaPackage)
}

// DeleteMetaID is a paid mutator transaction binding the contract method 0x76845d26.
//
// Solidity: function deleteMetaID(_metaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactor) DeleteMetaID(opts *bind.TransactOpts, _metaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.contract.Transact(opts, "deleteMetaID", _metaID, _sig, _metaPackage)
}

// DeleteMetaID is a paid mutator transaction binding the contract method 0x76845d26.
//
// Solidity: function deleteMetaID(_metaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerSession) DeleteMetaID(_metaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.DeleteMetaID(&_Identitymanager.TransactOpts, _metaID, _sig, _metaPackage)
}

// DeleteMetaID is a paid mutator transaction binding the contract method 0x76845d26.
//
// Solidity: function deleteMetaID(_metaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactorSession) DeleteMetaID(_metaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.DeleteMetaID(&_Identitymanager.TransactOpts, _metaID, _sig, _metaPackage)
}

// SetMetadiumNameServiceAddress is a paid mutator transaction binding the contract method 0x71171822.
//
// Solidity: function setMetadiumNameServiceAddress(_addr address) returns()
func (_Identitymanager *IdentitymanagerTransactor) SetMetadiumNameServiceAddress(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Identitymanager.contract.Transact(opts, "setMetadiumNameServiceAddress", _addr)
}

// SetMetadiumNameServiceAddress is a paid mutator transaction binding the contract method 0x71171822.
//
// Solidity: function setMetadiumNameServiceAddress(_addr address) returns()
func (_Identitymanager *IdentitymanagerSession) SetMetadiumNameServiceAddress(_addr common.Address) (*types.Transaction, error) {
	return _Identitymanager.Contract.SetMetadiumNameServiceAddress(&_Identitymanager.TransactOpts, _addr)
}

// SetMetadiumNameServiceAddress is a paid mutator transaction binding the contract method 0x71171822.
//
// Solidity: function setMetadiumNameServiceAddress(_addr address) returns()
func (_Identitymanager *IdentitymanagerTransactorSession) SetMetadiumNameServiceAddress(_addr common.Address) (*types.Transaction, error) {
	return _Identitymanager.Contract.SetMetadiumNameServiceAddress(&_Identitymanager.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Identitymanager *IdentitymanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Identitymanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Identitymanager *IdentitymanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identitymanager.Contract.TransferOwnership(&_Identitymanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Identitymanager *IdentitymanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identitymanager.Contract.TransferOwnership(&_Identitymanager.TransactOpts, newOwner)
}

// UpdateMetaID is a paid mutator transaction binding the contract method 0xade266fe.
//
// Solidity: function updateMetaID(_oldMetaID bytes32, _newMetaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactor) UpdateMetaID(opts *bind.TransactOpts, _oldMetaID [32]byte, _newMetaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.contract.Transact(opts, "updateMetaID", _oldMetaID, _newMetaID, _sig, _metaPackage)
}

// UpdateMetaID is a paid mutator transaction binding the contract method 0xade266fe.
//
// Solidity: function updateMetaID(_oldMetaID bytes32, _newMetaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerSession) UpdateMetaID(_oldMetaID [32]byte, _newMetaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.UpdateMetaID(&_Identitymanager.TransactOpts, _oldMetaID, _newMetaID, _sig, _metaPackage)
}

// UpdateMetaID is a paid mutator transaction binding the contract method 0xade266fe.
//
// Solidity: function updateMetaID(_oldMetaID bytes32, _newMetaID bytes32, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactorSession) UpdateMetaID(_oldMetaID [32]byte, _newMetaID [32]byte, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.UpdateMetaID(&_Identitymanager.TransactOpts, _oldMetaID, _newMetaID, _sig, _metaPackage)
}

// IdentitymanagerCreateMetaIDIterator is returned from FilterCreateMetaID and is used to iterate over the raw logs and unpacked data for CreateMetaID events raised by the Identitymanager contract.
type IdentitymanagerCreateMetaIDIterator struct {
	Event *IdentitymanagerCreateMetaID // Event containing the contract specifics and raw log

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
func (it *IdentitymanagerCreateMetaIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentitymanagerCreateMetaID)
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
		it.Event = new(IdentitymanagerCreateMetaID)
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
func (it *IdentitymanagerCreateMetaIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentitymanagerCreateMetaIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentitymanagerCreateMetaID represents a CreateMetaID event raised by the Identitymanager contract.
type IdentitymanagerCreateMetaID struct {
	Owner  common.Address
	MetaID [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreateMetaID is a free log retrieval operation binding the contract event 0x0afbed7f307c1b94f44fd781ff78fa8f85831c62848a1c53db8d28f575cf1aa1.
//
// Solidity: e CreateMetaID(owner indexed address, metaID indexed bytes32)
func (_Identitymanager *IdentitymanagerFilterer) FilterCreateMetaID(opts *bind.FilterOpts, owner []common.Address, metaID [][32]byte) (*IdentitymanagerCreateMetaIDIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var metaIDRule []interface{}
	for _, metaIDItem := range metaID {
		metaIDRule = append(metaIDRule, metaIDItem)
	}

	logs, sub, err := _Identitymanager.contract.FilterLogs(opts, "CreateMetaID", ownerRule, metaIDRule)
	if err != nil {
		return nil, err
	}
	return &IdentitymanagerCreateMetaIDIterator{contract: _Identitymanager.contract, event: "CreateMetaID", logs: logs, sub: sub}, nil
}

// WatchCreateMetaID is a free log subscription operation binding the contract event 0x0afbed7f307c1b94f44fd781ff78fa8f85831c62848a1c53db8d28f575cf1aa1.
//
// Solidity: e CreateMetaID(owner indexed address, metaID indexed bytes32)
func (_Identitymanager *IdentitymanagerFilterer) WatchCreateMetaID(opts *bind.WatchOpts, sink chan<- *IdentitymanagerCreateMetaID, owner []common.Address, metaID [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var metaIDRule []interface{}
	for _, metaIDItem := range metaID {
		metaIDRule = append(metaIDRule, metaIDItem)
	}

	logs, sub, err := _Identitymanager.contract.WatchLogs(opts, "CreateMetaID", ownerRule, metaIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentitymanagerCreateMetaID)
				if err := _Identitymanager.contract.UnpackLog(event, "CreateMetaID", log); err != nil {
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

// IdentitymanagerDeleteMetaIDIterator is returned from FilterDeleteMetaID and is used to iterate over the raw logs and unpacked data for DeleteMetaID events raised by the Identitymanager contract.
type IdentitymanagerDeleteMetaIDIterator struct {
	Event *IdentitymanagerDeleteMetaID // Event containing the contract specifics and raw log

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
func (it *IdentitymanagerDeleteMetaIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentitymanagerDeleteMetaID)
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
		it.Event = new(IdentitymanagerDeleteMetaID)
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
func (it *IdentitymanagerDeleteMetaIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentitymanagerDeleteMetaIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentitymanagerDeleteMetaID represents a DeleteMetaID event raised by the Identitymanager contract.
type IdentitymanagerDeleteMetaID struct {
	Owner  common.Address
	MetaID [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeleteMetaID is a free log retrieval operation binding the contract event 0x4e8bb03e2af31dac5f75505a69e10ce88e9d1239cc1859e1e79a7d360fbf061b.
//
// Solidity: e DeleteMetaID(owner indexed address, metaID indexed bytes32)
func (_Identitymanager *IdentitymanagerFilterer) FilterDeleteMetaID(opts *bind.FilterOpts, owner []common.Address, metaID [][32]byte) (*IdentitymanagerDeleteMetaIDIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var metaIDRule []interface{}
	for _, metaIDItem := range metaID {
		metaIDRule = append(metaIDRule, metaIDItem)
	}

	logs, sub, err := _Identitymanager.contract.FilterLogs(opts, "DeleteMetaID", ownerRule, metaIDRule)
	if err != nil {
		return nil, err
	}
	return &IdentitymanagerDeleteMetaIDIterator{contract: _Identitymanager.contract, event: "DeleteMetaID", logs: logs, sub: sub}, nil
}

// WatchDeleteMetaID is a free log subscription operation binding the contract event 0x4e8bb03e2af31dac5f75505a69e10ce88e9d1239cc1859e1e79a7d360fbf061b.
//
// Solidity: e DeleteMetaID(owner indexed address, metaID indexed bytes32)
func (_Identitymanager *IdentitymanagerFilterer) WatchDeleteMetaID(opts *bind.WatchOpts, sink chan<- *IdentitymanagerDeleteMetaID, owner []common.Address, metaID [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var metaIDRule []interface{}
	for _, metaIDItem := range metaID {
		metaIDRule = append(metaIDRule, metaIDItem)
	}

	logs, sub, err := _Identitymanager.contract.WatchLogs(opts, "DeleteMetaID", ownerRule, metaIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentitymanagerDeleteMetaID)
				if err := _Identitymanager.contract.UnpackLog(event, "DeleteMetaID", log); err != nil {
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

// IdentitymanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Identitymanager contract.
type IdentitymanagerOwnershipTransferredIterator struct {
	Event *IdentitymanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IdentitymanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentitymanagerOwnershipTransferred)
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
		it.Event = new(IdentitymanagerOwnershipTransferred)
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
func (it *IdentitymanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentitymanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentitymanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Identitymanager contract.
type IdentitymanagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Identitymanager *IdentitymanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentitymanagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identitymanager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentitymanagerOwnershipTransferredIterator{contract: _Identitymanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Identitymanager *IdentitymanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentitymanagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identitymanager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentitymanagerOwnershipTransferred)
				if err := _Identitymanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// IdentitymanagerUpdateMetaIDIterator is returned from FilterUpdateMetaID and is used to iterate over the raw logs and unpacked data for UpdateMetaID events raised by the Identitymanager contract.
type IdentitymanagerUpdateMetaIDIterator struct {
	Event *IdentitymanagerUpdateMetaID // Event containing the contract specifics and raw log

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
func (it *IdentitymanagerUpdateMetaIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentitymanagerUpdateMetaID)
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
		it.Event = new(IdentitymanagerUpdateMetaID)
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
func (it *IdentitymanagerUpdateMetaIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentitymanagerUpdateMetaIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentitymanagerUpdateMetaID represents a UpdateMetaID event raised by the Identitymanager contract.
type IdentitymanagerUpdateMetaID struct {
	Owner     common.Address
	OldMetaID [32]byte
	NewMetaID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateMetaID is a free log retrieval operation binding the contract event 0xde4d89fa4d58496bf0ceb285c9696695ffef28d55d9c47936ef981f124aa2c8c.
//
// Solidity: e UpdateMetaID(owner indexed address, oldMetaID indexed bytes32, newMetaID indexed bytes32)
func (_Identitymanager *IdentitymanagerFilterer) FilterUpdateMetaID(opts *bind.FilterOpts, owner []common.Address, oldMetaID [][32]byte, newMetaID [][32]byte) (*IdentitymanagerUpdateMetaIDIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var oldMetaIDRule []interface{}
	for _, oldMetaIDItem := range oldMetaID {
		oldMetaIDRule = append(oldMetaIDRule, oldMetaIDItem)
	}
	var newMetaIDRule []interface{}
	for _, newMetaIDItem := range newMetaID {
		newMetaIDRule = append(newMetaIDRule, newMetaIDItem)
	}

	logs, sub, err := _Identitymanager.contract.FilterLogs(opts, "UpdateMetaID", ownerRule, oldMetaIDRule, newMetaIDRule)
	if err != nil {
		return nil, err
	}
	return &IdentitymanagerUpdateMetaIDIterator{contract: _Identitymanager.contract, event: "UpdateMetaID", logs: logs, sub: sub}, nil
}

// WatchUpdateMetaID is a free log subscription operation binding the contract event 0xde4d89fa4d58496bf0ceb285c9696695ffef28d55d9c47936ef981f124aa2c8c.
//
// Solidity: e UpdateMetaID(owner indexed address, oldMetaID indexed bytes32, newMetaID indexed bytes32)
func (_Identitymanager *IdentitymanagerFilterer) WatchUpdateMetaID(opts *bind.WatchOpts, sink chan<- *IdentitymanagerUpdateMetaID, owner []common.Address, oldMetaID [][32]byte, newMetaID [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var oldMetaIDRule []interface{}
	for _, oldMetaIDItem := range oldMetaID {
		oldMetaIDRule = append(oldMetaIDRule, oldMetaIDItem)
	}
	var newMetaIDRule []interface{}
	for _, newMetaIDItem := range newMetaID {
		newMetaIDRule = append(newMetaIDRule, newMetaIDItem)
	}

	logs, sub, err := _Identitymanager.contract.WatchLogs(opts, "UpdateMetaID", ownerRule, oldMetaIDRule, newMetaIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentitymanagerUpdateMetaID)
				if err := _Identitymanager.contract.UnpackLog(event, "UpdateMetaID", log); err != nil {
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
