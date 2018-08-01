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
const IdentitymanagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"MNS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MID\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameMetaIdentityManager\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nameMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"metaID\",\"type\":\"bytes32\"}],\"name\":\"CreateMetaID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"oldMetaID\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"newMetaID\",\"type\":\"bytes32\"}],\"name\":\"UpdateMetaID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"oldMetaID\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"newMetaID\",\"type\":\"bytes32\"}],\"name\":\"RestoreMetaID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"metaID\",\"type\":\"bytes32\"}],\"name\":\"DeleteMetaID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_metaID\",\"type\":\"bytes32\"},{\"name\":\"_sig\",\"type\":\"bytes\"},{\"name\":\"_metaPackage\",\"type\":\"bytes\"}],\"name\":\"createMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_metaID\",\"type\":\"bytes32\"},{\"name\":\"_timestamp\",\"type\":\"bytes\"},{\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"deleteMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldMetaID\",\"type\":\"bytes32\"},{\"name\":\"_newMetaID\",\"type\":\"bytes32\"},{\"name\":\"_sig\",\"type\":\"bytes\"},{\"name\":\"_metaPackage\",\"type\":\"bytes\"}],\"name\":\"updateMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oldMetaID\",\"type\":\"bytes32\"},{\"name\":\"_newMetaID\",\"type\":\"bytes32\"},{\"name\":\"_oldAddress\",\"type\":\"address\"},{\"name\":\"_sig\",\"type\":\"bytes\"},{\"name\":\"_metaPackage\",\"type\":\"bytes\"}],\"name\":\"restoreMetaID\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"bytes32\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"bytes32\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"bytes32\"}],\"name\":\"tokenURIAsBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"ecrecovery\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ecverify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ecverifyWithTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"getAddressFromMetaPackage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMetadiumNameServiceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMetaIDAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IdentitymanagerBin is the compiled bytecode used for deploying new contracts.
const IdentitymanagerBin = `6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061330d806100536000396000f30060806040526004361061011d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630c1f2e0f1461012257806326336ce7146101795780632f745c591461027c57806339cdde32146102e5578063702fecd11461039457806370a082311461046957806371171822146104c057806371966d141461050357806376845d261461054657806377d32e941461061b5780637dd56411146106d25780638ab38747146107435780638da5cb5b14610838578063a37dd7af1461088f578063ade266fe14610939578063afbb8d5114610a1c578063c520a1e014610ac6578063d7a7ac6814610b6f578063e9e7b9ea14610bc6578063eaeaa3d914610bf9578063f2fde38b14610c2c575b600080fd5b34801561012e57600080fd5b50610137610c6f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561018557600080fd5b5061026260048036038101908080356000191690602001909291908035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610c95565b604051808215151515815260200191505060405180910390f35b34801561028857600080fd5b506102c7600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061128d565b60405180826000191660001916815260200191505060405180910390f35b3480156102f157600080fd5b5061037a6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611452565b604051808215151515815260200191505060405180910390f35b3480156103a057600080fd5b5061044f6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611545565b604051808215151515815260200191505060405180910390f35b34801561047557600080fd5b506104aa600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a30565b6040518082815260200191505060405180910390f35b3480156104cc57600080fd5b50610501600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611be9565b005b34801561050f57600080fd5b50610544600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611c88565b005b34801561055257600080fd5b506106016004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611e1a565b604051808215151515815260200191505060405180910390f35b34801561062757600080fd5b506106906004803603810190808035600019169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061215a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106de57600080fd5b506107016004803603810190808035600019169060200190929190505050612252565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561074f57600080fd5b5061081e6004803603810190808035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506123e3565b604051808215151515815260200191505060405180910390f35b34801561084457600080fd5b5061084d612623565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561089b57600080fd5b506108be6004803603810190808035600019169060200190929190505050612648565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108fe5780820151818401526020810190506108e3565b50505050905090810190601f16801561092b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561094557600080fd5b50610a0260048036038101908080356000191690602001909291908035600019169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050612830565b604051808215151515815260200191505060405180910390f35b348015610a2857600080fd5b50610a4b6004803603810190808035600019169060200190929190505050612e27565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a8b578082015181840152602081019050610a70565b50505050905090810190601f168015610ab85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610ad257600080fd5b50610b2d600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061300f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610b7b57600080fd5b50610b8461311e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015610bd257600080fd5b50610bdb613144565b60405180826000191660001916815260200191505060405180910390f35b348015610c0557600080fd5b50610c0e613168565b60405180826000191660001916815260200191505060405180910390f35b348015610c3857600080fd5b50610c6d600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061318c565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614158015610d465750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b1515610d5157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360d6c7cf7f4d6574616469756d4964656e746974794d616e61676572000000000000000000336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015610e3e57600080fd5b505af1158015610e52573d6000803e3d6000fd5b505050506040513d6020811015610e6857600080fd5b81019080805190602001909291905050501515610e8457600080fd5b82600060559050808251141515610e9a57600080fd5b610ea38561300f565b92508673ffffffffffffffffffffffffffffffffffffffff16610ec58a612252565b73ffffffffffffffffffffffffffffffffffffffff16141515610ee757600080fd5b610ef2888785611452565b1515610efd57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634f558e7989600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015610f9257600080fd5b505af1158015610fa6573d6000803e3d6000fd5b505050506040513d6020811015610fbc57600080fd5b8101908080519060200190929190505050151515610fd957600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342966c688a600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561106e57600080fd5b505af1158015611082573d6000803e3d6000fd5b505050506040513d602081101561109857600080fd5b810190808051906020019092919050505015156110b457600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d3fc9864848a60019004886040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561119c578082015181840152602081019050611181565b50505050905090810190601f1680156111c95780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b1580156111ea57600080fd5b505af11580156111fe573d6000803e3d6000fd5b505050506040513d602081101561121457600080fd5b8101908080519060200190929190505050151561123057600080fd5b876000191689600019168473ffffffffffffffffffffffffffffffffffffffff167ffbec1ceb80a2d30df4afb2beb79b98f019a582e77fa600eb96ce2e1d1f66cf0a60405160405180910390a46001935050505095945050505050565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415801561133c5750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b151561134757600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632f745c5984846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561140c57600080fd5b505af1158015611420573d6000803e3d6000fd5b505050506040513d602081101561143657600080fd5b8101908080519060200190929190505050600102905092915050565b600060606040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a333200000000815250905080856040518083805190602001908083835b6020831015156114c557805182526020820191506020810190506020830392506114a0565b6001836020036101000a0380198251168184511680821785525050505050509050018260001916600019168152602001925050506040518091039020945061150d858561215a565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16149150509392505050565b600080600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156115f65750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b151561160157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360d6c7cf7f4d6574616469756d4964656e746974794d616e61676572000000000000000000336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156116ee57600080fd5b505af1158015611702573d6000803e3d6000fd5b505050506040513d602081101561171857600080fd5b8101908080519060200190929190505050151561173457600080fd5b8260006055905080825114151561174a57600080fd5b6117538561300f565b9250611760878785611452565b151561176b57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634f558e7988600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561180057600080fd5b505af1158015611814573d6000803e3d6000fd5b505050506040513d602081101561182a57600080fd5b810190808051906020019092919050505015151561184757600080fd5b600061185284611a30565b14151561185e57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d3fc9864848960019004886040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561194657808201518184015260208101905061192b565b50505050905090810190601f1680156119735780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561199457600080fd5b505af11580156119a8573d6000803e3d6000fd5b505050506040513d60208110156119be57600080fd5b810190808051906020019092919050505015156119da57600080fd5b86600019168373ffffffffffffffffffffffffffffffffffffffff167f0afbed7f307c1b94f44fd781ff78fa8f85831c62848a1c53db8d28f575cf1aa160405160405180910390a3600193505050509392505050565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614158015611adf5750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b1515611aea57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015611ba757600080fd5b505af1158015611bbb573d6000803e3d6000fd5b505050506040513d6020811015611bd157600080fd5b81019080805190602001909291905050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611c4457600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611ce357600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630d2020dd7f4d657461494400000000000000000000000000000000000000000000000000006040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808260001916600019168152602001915050602060405180830381600087803b158015611d9c57600080fd5b505af1158015611db0573d6000803e3d6000fd5b505050506040513d6020811015611dc657600080fd5b8101908080519060200190929190505050600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614158015611ecb5750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b1515611ed657600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360d6c7cf7f4d6574616469756d4964656e746974794d616e61676572000000000000000000336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015611fc357600080fd5b505af1158015611fd7573d6000803e3d6000fd5b505050506040513d6020811015611fed57600080fd5b8101908080519060200190929190505050151561200957600080fd5b61201285612252565b9050612020858585846123e3565b151561202b57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342966c6886600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156120c057600080fd5b505af11580156120d4573d6000803e3d6000fd5b505050506040513d60208110156120ea57600080fd5b8101908080519060200190929190505050151561210657600080fd5b84600019168173ffffffffffffffffffffffffffffffffffffffff167f4e8bb03e2af31dac5f75505a69e10ce88e9d1239cc1859e1e79a7d360fbf061b60405160405180910390a360019150509392505050565b600080600080604185511415156121745760009350612249565b602085015192506040850151915060ff6041860151169050601b8160ff16101561219f57601b810190505b601b8160ff16141580156121b75750601c8160ff1614155b156121c55760009350612249565b600186828585604051600081526020016040526040518085600019166000191681526020018460ff1660ff1681526020018360001916600019168152602001826000191660001916815260200194505050505060206040516020810390808403906000865af115801561223c573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b60008073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156123015750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b151561230c57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e83600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156123a157600080fd5b505af11580156123b5573d6000803e3d6000fd5b505050506040513d60208110156123cb57600080fd5b81019080805190602001909291905050509050919050565b600060606000806040805190810160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525092508651602060ff160191506030600a8381151561243a57fe5b04017f01000000000000000000000000000000000000000000000000000000000000000283600285510381518110151561247057fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506030600a838115156124ae57fe5b06017f0100000000000000000000000000000000000000000000000000000000000000028360018551038151811015156124e457fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508288886040518084805190602001908083835b60208310151561254c5780518252602082019150602081019050602083039250612527565b6001836020036101000a038019825116818451168082178552505050505050905001836000191660001916815260200182805190602001908083835b6020831015156125ad5780518252602082019150602081019050602083039250612588565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051809103902090506125e8818761215a565b73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16149350505050949350505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6060600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156126f85750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b151561270357600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663108f96df83600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561279857600080fd5b505af11580156127ac573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156127d657600080fd5b8101908080516401000000008111156127ee57600080fd5b8281019050602081018481111561280457600080fd5b815185600182028301116401000000008211171561282157600080fd5b50509291905050509050919050565b600080600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141580156128e15750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b15156128ec57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360d6c7cf7f4d6574616469756d4964656e746974794d616e61676572000000000000000000336040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b1580156129d957600080fd5b505af11580156129ed573d6000803e3d6000fd5b505050506040513d6020811015612a0357600080fd5b81019080805190602001909291905050501515612a1f57600080fd5b82600060559050808251141515612a3557600080fd5b612a3e8561300f565b92508273ffffffffffffffffffffffffffffffffffffffff16612a6089612252565b73ffffffffffffffffffffffffffffffffffffffff16141515612a8257600080fd5b612a8d878785611452565b1515612a9857600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634f558e7988600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015612b2d57600080fd5b505af1158015612b41573d6000803e3d6000fd5b505050506040513d6020811015612b5757600080fd5b8101908080519060200190929190505050151515612b7457600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342966c6889600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015612c0957600080fd5b505af1158015612c1d573d6000803e3d6000fd5b505050506040513d6020811015612c3357600080fd5b81019080805190602001909291905050501515612c4f57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d3fc9864848960019004886040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015612d37578082015181840152602081019050612d1c565b50505050905090810190601f168015612d645780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015612d8557600080fd5b505af1158015612d99573d6000803e3d6000fd5b505050506040513d6020811015612daf57600080fd5b81019080805190602001909291905050501515612dcb57600080fd5b866000191688600019168473ffffffffffffffffffffffffffffffffffffffff167fde4d89fa4d58496bf0ceb285c9696695ffef28d55d9c47936ef981f124aa2c8c60405160405180910390a460019350505050949350505050565b6060600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614158015612ed75750600073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b1515612ee257600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c87b56dd83600190046040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b158015612f7757600080fd5b505af1158015612f8b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015612fb557600080fd5b810190808051640100000000811115612fcd57600080fd5b82810190506020810184811115612fe357600080fd5b815185600182028301116401000000008211171561300057600080fd5b50509291905050509050919050565b6000806000806016925082855111151561302857600080fd5b600190505b6015811015613103576008600182030260ff7f010000000000000000000000000000000000000000000000000000000000000002868381518110151561306f57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166bffffffffffffffffffffffff19169060020a900482179150808060010191505061302d565b816c0100000000000000000000000090049350505050919050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f4d6574616469756d4964656e746974794d616e6167657200000000000000000081565b7f4d6574614944000000000000000000000000000000000000000000000000000081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156131e757600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561322357600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a723058205e990a91899e4e03fc330999dd91f4ce4ed49f64804ed5a1037140c9093c39d30029`

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

// MID is a free data retrieval call binding the contract method 0xd7a7ac68.
//
// Solidity: function MID() constant returns(address)
func (_Identitymanager *IdentitymanagerCaller) MID(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "MID")
	return *ret0, err
}

// MID is a free data retrieval call binding the contract method 0xd7a7ac68.
//
// Solidity: function MID() constant returns(address)
func (_Identitymanager *IdentitymanagerSession) MID() (common.Address, error) {
	return _Identitymanager.Contract.MID(&_Identitymanager.CallOpts)
}

// MID is a free data retrieval call binding the contract method 0xd7a7ac68.
//
// Solidity: function MID() constant returns(address)
func (_Identitymanager *IdentitymanagerCallerSession) MID() (common.Address, error) {
	return _Identitymanager.Contract.MID(&_Identitymanager.CallOpts)
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

// EcverifyWithTimestamp is a free data retrieval call binding the contract method 0x8ab38747.
//
// Solidity: function ecverifyWithTimestamp(message bytes32, timestamp bytes, sig bytes, signer address) constant returns(bool)
func (_Identitymanager *IdentitymanagerCaller) EcverifyWithTimestamp(opts *bind.CallOpts, message [32]byte, timestamp []byte, sig []byte, signer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "ecverifyWithTimestamp", message, timestamp, sig, signer)
	return *ret0, err
}

// EcverifyWithTimestamp is a free data retrieval call binding the contract method 0x8ab38747.
//
// Solidity: function ecverifyWithTimestamp(message bytes32, timestamp bytes, sig bytes, signer address) constant returns(bool)
func (_Identitymanager *IdentitymanagerSession) EcverifyWithTimestamp(message [32]byte, timestamp []byte, sig []byte, signer common.Address) (bool, error) {
	return _Identitymanager.Contract.EcverifyWithTimestamp(&_Identitymanager.CallOpts, message, timestamp, sig, signer)
}

// EcverifyWithTimestamp is a free data retrieval call binding the contract method 0x8ab38747.
//
// Solidity: function ecverifyWithTimestamp(message bytes32, timestamp bytes, sig bytes, signer address) constant returns(bool)
func (_Identitymanager *IdentitymanagerCallerSession) EcverifyWithTimestamp(message [32]byte, timestamp []byte, sig []byte, signer common.Address) (bool, error) {
	return _Identitymanager.Contract.EcverifyWithTimestamp(&_Identitymanager.CallOpts, message, timestamp, sig, signer)
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

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(_tokenId bytes32) constant returns(_owner address)
func (_Identitymanager *IdentitymanagerCaller) OwnerOf(opts *bind.CallOpts, _tokenId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(_tokenId bytes32) constant returns(_owner address)
func (_Identitymanager *IdentitymanagerSession) OwnerOf(_tokenId [32]byte) (common.Address, error) {
	return _Identitymanager.Contract.OwnerOf(&_Identitymanager.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x7dd56411.
//
// Solidity: function ownerOf(_tokenId bytes32) constant returns(_owner address)
func (_Identitymanager *IdentitymanagerCallerSession) OwnerOf(_tokenId [32]byte) (common.Address, error) {
	return _Identitymanager.Contract.OwnerOf(&_Identitymanager.CallOpts, _tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId bytes32)
func (_Identitymanager *IdentitymanagerCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId bytes32)
func (_Identitymanager *IdentitymanagerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) ([32]byte, error) {
	return _Identitymanager.Contract.TokenOfOwnerByIndex(&_Identitymanager.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId bytes32)
func (_Identitymanager *IdentitymanagerCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) ([32]byte, error) {
	return _Identitymanager.Contract.TokenOfOwnerByIndex(&_Identitymanager.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xafbb8d51.
//
// Solidity: function tokenURI(_tokenId bytes32) constant returns(string)
func (_Identitymanager *IdentitymanagerCaller) TokenURI(opts *bind.CallOpts, _tokenId [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xafbb8d51.
//
// Solidity: function tokenURI(_tokenId bytes32) constant returns(string)
func (_Identitymanager *IdentitymanagerSession) TokenURI(_tokenId [32]byte) (string, error) {
	return _Identitymanager.Contract.TokenURI(&_Identitymanager.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xafbb8d51.
//
// Solidity: function tokenURI(_tokenId bytes32) constant returns(string)
func (_Identitymanager *IdentitymanagerCallerSession) TokenURI(_tokenId [32]byte) (string, error) {
	return _Identitymanager.Contract.TokenURI(&_Identitymanager.CallOpts, _tokenId)
}

// TokenURIAsBytes is a free data retrieval call binding the contract method 0xa37dd7af.
//
// Solidity: function tokenURIAsBytes(_tokenId bytes32) constant returns(bytes)
func (_Identitymanager *IdentitymanagerCaller) TokenURIAsBytes(opts *bind.CallOpts, _tokenId [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Identitymanager.contract.Call(opts, out, "tokenURIAsBytes", _tokenId)
	return *ret0, err
}

// TokenURIAsBytes is a free data retrieval call binding the contract method 0xa37dd7af.
//
// Solidity: function tokenURIAsBytes(_tokenId bytes32) constant returns(bytes)
func (_Identitymanager *IdentitymanagerSession) TokenURIAsBytes(_tokenId [32]byte) ([]byte, error) {
	return _Identitymanager.Contract.TokenURIAsBytes(&_Identitymanager.CallOpts, _tokenId)
}

// TokenURIAsBytes is a free data retrieval call binding the contract method 0xa37dd7af.
//
// Solidity: function tokenURIAsBytes(_tokenId bytes32) constant returns(bytes)
func (_Identitymanager *IdentitymanagerCallerSession) TokenURIAsBytes(_tokenId [32]byte) ([]byte, error) {
	return _Identitymanager.Contract.TokenURIAsBytes(&_Identitymanager.CallOpts, _tokenId)
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
// Solidity: function deleteMetaID(_metaID bytes32, _timestamp bytes, _sig bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactor) DeleteMetaID(opts *bind.TransactOpts, _metaID [32]byte, _timestamp []byte, _sig []byte) (*types.Transaction, error) {
	return _Identitymanager.contract.Transact(opts, "deleteMetaID", _metaID, _timestamp, _sig)
}

// DeleteMetaID is a paid mutator transaction binding the contract method 0x76845d26.
//
// Solidity: function deleteMetaID(_metaID bytes32, _timestamp bytes, _sig bytes) returns(bool)
func (_Identitymanager *IdentitymanagerSession) DeleteMetaID(_metaID [32]byte, _timestamp []byte, _sig []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.DeleteMetaID(&_Identitymanager.TransactOpts, _metaID, _timestamp, _sig)
}

// DeleteMetaID is a paid mutator transaction binding the contract method 0x76845d26.
//
// Solidity: function deleteMetaID(_metaID bytes32, _timestamp bytes, _sig bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactorSession) DeleteMetaID(_metaID [32]byte, _timestamp []byte, _sig []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.DeleteMetaID(&_Identitymanager.TransactOpts, _metaID, _timestamp, _sig)
}

// RestoreMetaID is a paid mutator transaction binding the contract method 0x26336ce7.
//
// Solidity: function restoreMetaID(_oldMetaID bytes32, _newMetaID bytes32, _oldAddress address, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactor) RestoreMetaID(opts *bind.TransactOpts, _oldMetaID [32]byte, _newMetaID [32]byte, _oldAddress common.Address, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.contract.Transact(opts, "restoreMetaID", _oldMetaID, _newMetaID, _oldAddress, _sig, _metaPackage)
}

// RestoreMetaID is a paid mutator transaction binding the contract method 0x26336ce7.
//
// Solidity: function restoreMetaID(_oldMetaID bytes32, _newMetaID bytes32, _oldAddress address, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerSession) RestoreMetaID(_oldMetaID [32]byte, _newMetaID [32]byte, _oldAddress common.Address, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.RestoreMetaID(&_Identitymanager.TransactOpts, _oldMetaID, _newMetaID, _oldAddress, _sig, _metaPackage)
}

// RestoreMetaID is a paid mutator transaction binding the contract method 0x26336ce7.
//
// Solidity: function restoreMetaID(_oldMetaID bytes32, _newMetaID bytes32, _oldAddress address, _sig bytes, _metaPackage bytes) returns(bool)
func (_Identitymanager *IdentitymanagerTransactorSession) RestoreMetaID(_oldMetaID [32]byte, _newMetaID [32]byte, _oldAddress common.Address, _sig []byte, _metaPackage []byte) (*types.Transaction, error) {
	return _Identitymanager.Contract.RestoreMetaID(&_Identitymanager.TransactOpts, _oldMetaID, _newMetaID, _oldAddress, _sig, _metaPackage)
}

// SetMetaIDAddress is a paid mutator transaction binding the contract method 0x71966d14.
//
// Solidity: function setMetaIDAddress(_addr address) returns()
func (_Identitymanager *IdentitymanagerTransactor) SetMetaIDAddress(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Identitymanager.contract.Transact(opts, "setMetaIDAddress", _addr)
}

// SetMetaIDAddress is a paid mutator transaction binding the contract method 0x71966d14.
//
// Solidity: function setMetaIDAddress(_addr address) returns()
func (_Identitymanager *IdentitymanagerSession) SetMetaIDAddress(_addr common.Address) (*types.Transaction, error) {
	return _Identitymanager.Contract.SetMetaIDAddress(&_Identitymanager.TransactOpts, _addr)
}

// SetMetaIDAddress is a paid mutator transaction binding the contract method 0x71966d14.
//
// Solidity: function setMetaIDAddress(_addr address) returns()
func (_Identitymanager *IdentitymanagerTransactorSession) SetMetaIDAddress(_addr common.Address) (*types.Transaction, error) {
	return _Identitymanager.Contract.SetMetaIDAddress(&_Identitymanager.TransactOpts, _addr)
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

// IdentitymanagerRestoreMetaIDIterator is returned from FilterRestoreMetaID and is used to iterate over the raw logs and unpacked data for RestoreMetaID events raised by the Identitymanager contract.
type IdentitymanagerRestoreMetaIDIterator struct {
	Event *IdentitymanagerRestoreMetaID // Event containing the contract specifics and raw log

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
func (it *IdentitymanagerRestoreMetaIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentitymanagerRestoreMetaID)
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
		it.Event = new(IdentitymanagerRestoreMetaID)
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
func (it *IdentitymanagerRestoreMetaIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentitymanagerRestoreMetaIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentitymanagerRestoreMetaID represents a RestoreMetaID event raised by the Identitymanager contract.
type IdentitymanagerRestoreMetaID struct {
	Owner     common.Address
	OldMetaID [32]byte
	NewMetaID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRestoreMetaID is a free log retrieval operation binding the contract event 0xfbec1ceb80a2d30df4afb2beb79b98f019a582e77fa600eb96ce2e1d1f66cf0a.
//
// Solidity: e RestoreMetaID(owner indexed address, oldMetaID indexed bytes32, newMetaID indexed bytes32)
func (_Identitymanager *IdentitymanagerFilterer) FilterRestoreMetaID(opts *bind.FilterOpts, owner []common.Address, oldMetaID [][32]byte, newMetaID [][32]byte) (*IdentitymanagerRestoreMetaIDIterator, error) {

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

	logs, sub, err := _Identitymanager.contract.FilterLogs(opts, "RestoreMetaID", ownerRule, oldMetaIDRule, newMetaIDRule)
	if err != nil {
		return nil, err
	}
	return &IdentitymanagerRestoreMetaIDIterator{contract: _Identitymanager.contract, event: "RestoreMetaID", logs: logs, sub: sub}, nil
}

// WatchRestoreMetaID is a free log subscription operation binding the contract event 0xfbec1ceb80a2d30df4afb2beb79b98f019a582e77fa600eb96ce2e1d1f66cf0a.
//
// Solidity: e RestoreMetaID(owner indexed address, oldMetaID indexed bytes32, newMetaID indexed bytes32)
func (_Identitymanager *IdentitymanagerFilterer) WatchRestoreMetaID(opts *bind.WatchOpts, sink chan<- *IdentitymanagerRestoreMetaID, owner []common.Address, oldMetaID [][32]byte, newMetaID [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Identitymanager.contract.WatchLogs(opts, "RestoreMetaID", ownerRule, oldMetaIDRule, newMetaIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentitymanagerRestoreMetaID)
				if err := _Identitymanager.contract.UnpackLog(event, "RestoreMetaID", log); err != nil {
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
