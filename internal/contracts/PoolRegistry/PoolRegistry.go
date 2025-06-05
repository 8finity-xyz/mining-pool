// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PoolRegistry

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// PoolRegistryMetaData contains all meta data concerning the PoolRegistry contract.
var PoolRegistryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPoolOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolAlreadyCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PoolNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"PoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INFINITY\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POW\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"_notifyReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountToLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToLock_\",\"type\":\"uint256\"}],\"name\":\"changeAmountToLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"createPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"submitsCost\",\"type\":\"uint256\"}],\"name\":\"finalizeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedInfinity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastSubmit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unfinalizedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockInfinity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"fee\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "PoolRegistry",
}

// PoolRegistry is an auto generated Go binding around an Ethereum contract.
type PoolRegistry struct {
	abi abi.ABI
}

// NewPoolRegistry creates a new instance of PoolRegistry.
func NewPoolRegistry() *PoolRegistry {
	parsed, err := PoolRegistryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &PoolRegistry{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *PoolRegistry) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackINFINITY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f48f44b.
//
// Solidity: function INFINITY() view returns(address)
func (poolRegistry *PoolRegistry) PackINFINITY() []byte {
	enc, err := poolRegistry.abi.Pack("INFINITY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackINFINITY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f48f44b.
//
// Solidity: function INFINITY() view returns(address)
func (poolRegistry *PoolRegistry) UnpackINFINITY(data []byte) (common.Address, error) {
	out, err := poolRegistry.abi.Unpack("INFINITY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPOW is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4296bf6.
//
// Solidity: function POW() view returns(address)
func (poolRegistry *PoolRegistry) PackPOW() []byte {
	enc, err := poolRegistry.abi.Pack("POW")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPOW is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa4296bf6.
//
// Solidity: function POW() view returns(address)
func (poolRegistry *PoolRegistry) UnpackPOW(data []byte) (common.Address, error) {
	out, err := poolRegistry.abi.Unpack("POW", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (poolRegistry *PoolRegistry) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := poolRegistry.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (poolRegistry *PoolRegistry) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := poolRegistry.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackNotifyReward is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf1518500.
//
// Solidity: function _notifyReward(uint256 poolId, uint256 reward) returns()
func (poolRegistry *PoolRegistry) PackNotifyReward(poolId *big.Int, reward *big.Int) []byte {
	enc, err := poolRegistry.abi.Pack("_notifyReward", poolId, reward)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAmountToLock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9c6f4e5d.
//
// Solidity: function amountToLock() view returns(uint256)
func (poolRegistry *PoolRegistry) PackAmountToLock() []byte {
	enc, err := poolRegistry.abi.Pack("amountToLock")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackAmountToLock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9c6f4e5d.
//
// Solidity: function amountToLock() view returns(uint256)
func (poolRegistry *PoolRegistry) UnpackAmountToLock(data []byte) (*big.Int, error) {
	out, err := poolRegistry.abi.Unpack("amountToLock", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackChangeAmountToLock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d50e68f.
//
// Solidity: function changeAmountToLock(uint256 amountToLock_) returns()
func (poolRegistry *PoolRegistry) PackChangeAmountToLock(amountToLock *big.Int) []byte {
	enc, err := poolRegistry.abi.Pack("changeAmountToLock", amountToLock)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClaim is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x99016142.
//
// Solidity: function claim(uint256 poolId, address miner, uint256 totalReward, bytes signature) returns()
func (poolRegistry *PoolRegistry) PackClaim(poolId *big.Int, miner common.Address, totalReward *big.Int, signature []byte) []byte {
	enc, err := poolRegistry.abi.Pack("claim", poolId, miner, totalReward, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClosePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66805de5.
//
// Solidity: function closePool() returns()
func (poolRegistry *PoolRegistry) PackClosePool() []byte {
	enc, err := poolRegistry.abi.Pack("closePool")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCreatePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49c585dd.
//
// Solidity: function createPool(uint16 fee, string name, string url) returns()
func (poolRegistry *PoolRegistry) PackCreatePool(fee uint16, name string, url string) []byte {
	enc, err := poolRegistry.abi.Pack("createPool", fee, name, url)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (poolRegistry *PoolRegistry) PackEip712Domain() []byte {
	enc, err := poolRegistry.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (poolRegistry *PoolRegistry) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := poolRegistry.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, err

}

// PackFinalizeReward is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52056aa1.
//
// Solidity: function finalizeReward(uint256 submitsCost) returns()
func (poolRegistry *PoolRegistry) PackFinalizeReward(submitsCost *big.Int) []byte {
	enc, err := poolRegistry.abi.Pack("finalizeReward", submitsCost)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetPoolId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcaa9a08d.
//
// Solidity: function getPoolId(address owner) view returns(uint256 poolId)
func (poolRegistry *PoolRegistry) PackGetPoolId(owner common.Address) []byte {
	enc, err := poolRegistry.abi.Pack("getPoolId", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetPoolId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcaa9a08d.
//
// Solidity: function getPoolId(address owner) view returns(uint256 poolId)
func (poolRegistry *PoolRegistry) UnpackGetPoolId(data []byte) (*big.Int, error) {
	out, err := poolRegistry.abi.Unpack("getPoolId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (poolRegistry *PoolRegistry) PackInitialize(owner common.Address) []byte {
	enc, err := poolRegistry.abi.Pack("initialize", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackLastPoolId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa657e579.
//
// Solidity: function lastPoolId() view returns(uint256)
func (poolRegistry *PoolRegistry) PackLastPoolId() []byte {
	enc, err := poolRegistry.abi.Pack("lastPoolId")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackLastPoolId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa657e579.
//
// Solidity: function lastPoolId() view returns(uint256)
func (poolRegistry *PoolRegistry) UnpackLastPoolId(data []byte) (*big.Int, error) {
	out, err := poolRegistry.abi.Unpack("lastPoolId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackLockedInfinity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3cb25ecc.
//
// Solidity: function lockedInfinity(address ) view returns(uint256)
func (poolRegistry *PoolRegistry) PackLockedInfinity(arg0 common.Address) []byte {
	enc, err := poolRegistry.abi.Pack("lockedInfinity", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackLockedInfinity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3cb25ecc.
//
// Solidity: function lockedInfinity(address ) view returns(uint256)
func (poolRegistry *PoolRegistry) UnpackLockedInfinity(data []byte) (*big.Int, error) {
	out, err := poolRegistry.abi.Unpack("lockedInfinity", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (poolRegistry *PoolRegistry) PackOwner() []byte {
	enc, err := poolRegistry.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (poolRegistry *PoolRegistry) UnpackOwner(data []byte) (common.Address, error) {
	out, err := poolRegistry.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPoolById is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfdc55e8b.
//
// Solidity: function poolById(uint256 ) view returns(address owner, uint256 lastSubmit, uint256 unfinalizedReward, uint256 finalizedReward, uint256 totalReward, uint16 fee, string name, string url)
func (poolRegistry *PoolRegistry) PackPoolById(arg0 *big.Int) []byte {
	enc, err := poolRegistry.abi.Pack("poolById", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// PoolByIdOutput serves as a container for the return parameters of contract
// method PoolById.
type PoolByIdOutput struct {
	Owner             common.Address
	LastSubmit        *big.Int
	UnfinalizedReward *big.Int
	FinalizedReward   *big.Int
	TotalReward       *big.Int
	Fee               uint16
	Name              string
	Url               string
}

// UnpackPoolById is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfdc55e8b.
//
// Solidity: function poolById(uint256 ) view returns(address owner, uint256 lastSubmit, uint256 unfinalizedReward, uint256 finalizedReward, uint256 totalReward, uint16 fee, string name, string url)
func (poolRegistry *PoolRegistry) UnpackPoolById(data []byte) (PoolByIdOutput, error) {
	out, err := poolRegistry.abi.Unpack("poolById", data)
	outstruct := new(PoolByIdOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LastSubmit = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.UnfinalizedReward = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.FinalizedReward = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.TotalReward = abi.ConvertType(out[4], new(big.Int)).(*big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.Name = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Url = *abi.ConvertType(out[7], new(string)).(*string)
	return *outstruct, err

}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (poolRegistry *PoolRegistry) PackProxiableUUID() []byte {
	enc, err := poolRegistry.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (poolRegistry *PoolRegistry) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := poolRegistry.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (poolRegistry *PoolRegistry) PackRenounceOwnership() []byte {
	enc, err := poolRegistry.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (poolRegistry *PoolRegistry) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := poolRegistry.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUnlockInfinity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcae9a4e.
//
// Solidity: function unlockInfinity() returns()
func (poolRegistry *PoolRegistry) PackUnlockInfinity() []byte {
	enc, err := poolRegistry.abi.Pack("unlockInfinity")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpdatePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95947875.
//
// Solidity: function updatePool(uint8 fee, string name, string url) returns()
func (poolRegistry *PoolRegistry) PackUpdatePool(fee uint8, name string, url string) []byte {
	enc, err := poolRegistry.abi.Pack("updatePool", fee, name, url)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (poolRegistry *PoolRegistry) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := poolRegistry.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PoolRegistryEIP712DomainChanged represents a EIP712DomainChanged event raised by the PoolRegistry contract.
type PoolRegistryEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const PoolRegistryEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (PoolRegistryEIP712DomainChanged) ContractEventName() string {
	return PoolRegistryEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (poolRegistry *PoolRegistry) UnpackEIP712DomainChangedEvent(log *types.Log) (*PoolRegistryEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != poolRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoolRegistryEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := poolRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoolRegistryInitialized represents a Initialized event raised by the PoolRegistry contract.
type PoolRegistryInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const PoolRegistryInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (PoolRegistryInitialized) ContractEventName() string {
	return PoolRegistryInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (poolRegistry *PoolRegistry) UnpackInitializedEvent(log *types.Log) (*PoolRegistryInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != poolRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoolRegistryInitialized)
	if len(log.Data) > 0 {
		if err := poolRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoolRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the PoolRegistry contract.
type PoolRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const PoolRegistryOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (PoolRegistryOwnershipTransferred) ContractEventName() string {
	return PoolRegistryOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (poolRegistry *PoolRegistry) UnpackOwnershipTransferredEvent(log *types.Log) (*PoolRegistryOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != poolRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoolRegistryOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := poolRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoolRegistryPoolClosed represents a PoolClosed event raised by the PoolRegistry contract.
type PoolRegistryPoolClosed struct {
	PoolId *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const PoolRegistryPoolClosedEventName = "PoolClosed"

// ContractEventName returns the user-defined event name.
func (PoolRegistryPoolClosed) ContractEventName() string {
	return PoolRegistryPoolClosedEventName
}

// UnpackPoolClosedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PoolClosed(uint256 indexed poolId)
func (poolRegistry *PoolRegistry) UnpackPoolClosedEvent(log *types.Log) (*PoolRegistryPoolClosed, error) {
	event := "PoolClosed"
	if log.Topics[0] != poolRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoolRegistryPoolClosed)
	if len(log.Data) > 0 {
		if err := poolRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoolRegistryPoolCreated represents a PoolCreated event raised by the PoolRegistry contract.
type PoolRegistryPoolCreated struct {
	PoolId *big.Int
	Owner  common.Address
	Fee    uint16
	Name   string
	Url    string
	Raw    *types.Log // Blockchain specific contextual infos
}

const PoolRegistryPoolCreatedEventName = "PoolCreated"

// ContractEventName returns the user-defined event name.
func (PoolRegistryPoolCreated) ContractEventName() string {
	return PoolRegistryPoolCreatedEventName
}

// UnpackPoolCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PoolCreated(uint256 indexed poolId, address indexed owner, uint16 fee, string name, string url)
func (poolRegistry *PoolRegistry) UnpackPoolCreatedEvent(log *types.Log) (*PoolRegistryPoolCreated, error) {
	event := "PoolCreated"
	if log.Topics[0] != poolRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoolRegistryPoolCreated)
	if len(log.Data) > 0 {
		if err := poolRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoolRegistryPoolUpdated represents a PoolUpdated event raised by the PoolRegistry contract.
type PoolRegistryPoolUpdated struct {
	PoolId *big.Int
	Fee    uint16
	Name   string
	Url    string
	Raw    *types.Log // Blockchain specific contextual infos
}

const PoolRegistryPoolUpdatedEventName = "PoolUpdated"

// ContractEventName returns the user-defined event name.
func (PoolRegistryPoolUpdated) ContractEventName() string {
	return PoolRegistryPoolUpdatedEventName
}

// UnpackPoolUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PoolUpdated(uint256 indexed poolId, uint16 fee, string name, string url)
func (poolRegistry *PoolRegistry) UnpackPoolUpdatedEvent(log *types.Log) (*PoolRegistryPoolUpdated, error) {
	event := "PoolUpdated"
	if log.Topics[0] != poolRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoolRegistryPoolUpdated)
	if len(log.Data) > 0 {
		if err := poolRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoolRegistryRewardClaimed represents a RewardClaimed event raised by the PoolRegistry contract.
type PoolRegistryRewardClaimed struct {
	Miner  common.Address
	PoolId *big.Int
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const PoolRegistryRewardClaimedEventName = "RewardClaimed"

// ContractEventName returns the user-defined event name.
func (PoolRegistryRewardClaimed) ContractEventName() string {
	return PoolRegistryRewardClaimedEventName
}

// UnpackRewardClaimedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RewardClaimed(address miner, uint256 poolId, uint256 amount)
func (poolRegistry *PoolRegistry) UnpackRewardClaimedEvent(log *types.Log) (*PoolRegistryRewardClaimed, error) {
	event := "RewardClaimed"
	if log.Topics[0] != poolRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoolRegistryRewardClaimed)
	if len(log.Data) > 0 {
		if err := poolRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoolRegistryUpgraded represents a Upgraded event raised by the PoolRegistry contract.
type PoolRegistryUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const PoolRegistryUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (PoolRegistryUpgraded) ContractEventName() string {
	return PoolRegistryUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (poolRegistry *PoolRegistry) UnpackUpgradedEvent(log *types.Log) (*PoolRegistryUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != poolRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoolRegistryUpgraded)
	if len(log.Data) > 0 {
		if err := poolRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poolRegistry.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (poolRegistry *PoolRegistry) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["OnlyPoolOwner"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackOnlyPoolOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["PoolAlreadyCreated"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackPoolAlreadyCreatedError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["PoolNotFound"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackPoolNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], poolRegistry.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return poolRegistry.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PoolRegistryAddressEmptyCode represents a AddressEmptyCode error raised by the PoolRegistry contract.
type PoolRegistryAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func PoolRegistryAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (poolRegistry *PoolRegistry) UnpackAddressEmptyCodeError(raw []byte) (*PoolRegistryAddressEmptyCode, error) {
	out := new(PoolRegistryAddressEmptyCode)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the PoolRegistry contract.
type PoolRegistryECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func PoolRegistryECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (poolRegistry *PoolRegistry) UnpackECDSAInvalidSignatureError(raw []byte) (*PoolRegistryECDSAInvalidSignature, error) {
	out := new(PoolRegistryECDSAInvalidSignature)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the PoolRegistry contract.
type PoolRegistryECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func PoolRegistryECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (poolRegistry *PoolRegistry) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*PoolRegistryECDSAInvalidSignatureLength, error) {
	out := new(PoolRegistryECDSAInvalidSignatureLength)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the PoolRegistry contract.
type PoolRegistryECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func PoolRegistryECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (poolRegistry *PoolRegistry) UnpackECDSAInvalidSignatureSError(raw []byte) (*PoolRegistryECDSAInvalidSignatureS, error) {
	out := new(PoolRegistryECDSAInvalidSignatureS)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the PoolRegistry contract.
type PoolRegistryERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func PoolRegistryERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (poolRegistry *PoolRegistry) UnpackERC1967InvalidImplementationError(raw []byte) (*PoolRegistryERC1967InvalidImplementation, error) {
	out := new(PoolRegistryERC1967InvalidImplementation)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryERC1967NonPayable represents a ERC1967NonPayable error raised by the PoolRegistry contract.
type PoolRegistryERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func PoolRegistryERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (poolRegistry *PoolRegistry) UnpackERC1967NonPayableError(raw []byte) (*PoolRegistryERC1967NonPayable, error) {
	out := new(PoolRegistryERC1967NonPayable)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryFailedCall represents a FailedCall error raised by the PoolRegistry contract.
type PoolRegistryFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func PoolRegistryFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (poolRegistry *PoolRegistry) UnpackFailedCallError(raw []byte) (*PoolRegistryFailedCall, error) {
	out := new(PoolRegistryFailedCall)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryInvalidInitialization represents a InvalidInitialization error raised by the PoolRegistry contract.
type PoolRegistryInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func PoolRegistryInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (poolRegistry *PoolRegistry) UnpackInvalidInitializationError(raw []byte) (*PoolRegistryInvalidInitialization, error) {
	out := new(PoolRegistryInvalidInitialization)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryNotInitializing represents a NotInitializing error raised by the PoolRegistry contract.
type PoolRegistryNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func PoolRegistryNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (poolRegistry *PoolRegistry) UnpackNotInitializingError(raw []byte) (*PoolRegistryNotInitializing, error) {
	out := new(PoolRegistryNotInitializing)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryOnlyPoolOwner represents a OnlyPoolOwner error raised by the PoolRegistry contract.
type PoolRegistryOnlyPoolOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyPoolOwner()
func PoolRegistryOnlyPoolOwnerErrorID() common.Hash {
	return common.HexToHash("0x8fb524e1111beb7efd49d58a0599fb1c4de67f8fc190117867a8b51a5c4526ba")
}

// UnpackOnlyPoolOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyPoolOwner()
func (poolRegistry *PoolRegistry) UnpackOnlyPoolOwnerError(raw []byte) (*PoolRegistryOnlyPoolOwner, error) {
	out := new(PoolRegistryOnlyPoolOwner)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "OnlyPoolOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the PoolRegistry contract.
type PoolRegistryOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func PoolRegistryOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (poolRegistry *PoolRegistry) UnpackOwnableInvalidOwnerError(raw []byte) (*PoolRegistryOwnableInvalidOwner, error) {
	out := new(PoolRegistryOwnableInvalidOwner)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the PoolRegistry contract.
type PoolRegistryOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func PoolRegistryOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (poolRegistry *PoolRegistry) UnpackOwnableUnauthorizedAccountError(raw []byte) (*PoolRegistryOwnableUnauthorizedAccount, error) {
	out := new(PoolRegistryOwnableUnauthorizedAccount)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryPoolAlreadyCreated represents a PoolAlreadyCreated error raised by the PoolRegistry contract.
type PoolRegistryPoolAlreadyCreated struct {
	PoolId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PoolAlreadyCreated(uint256 poolId)
func PoolRegistryPoolAlreadyCreatedErrorID() common.Hash {
	return common.HexToHash("0x5ed2a88c1b4d15826afad80201d629201e7fb76a26c5f75108a70e21b8d0c98d")
}

// UnpackPoolAlreadyCreatedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PoolAlreadyCreated(uint256 poolId)
func (poolRegistry *PoolRegistry) UnpackPoolAlreadyCreatedError(raw []byte) (*PoolRegistryPoolAlreadyCreated, error) {
	out := new(PoolRegistryPoolAlreadyCreated)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "PoolAlreadyCreated", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryPoolNotFound represents a PoolNotFound error raised by the PoolRegistry contract.
type PoolRegistryPoolNotFound struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PoolNotFound(address owner)
func PoolRegistryPoolNotFoundErrorID() common.Hash {
	return common.HexToHash("0x6a34f98c401d042c461255fbda2d2ee3d9d52350d9c5ba09af49b05e6182d0d7")
}

// UnpackPoolNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PoolNotFound(address owner)
func (poolRegistry *PoolRegistry) UnpackPoolNotFoundError(raw []byte) (*PoolRegistryPoolNotFound, error) {
	out := new(PoolRegistryPoolNotFound)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "PoolNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the PoolRegistry contract.
type PoolRegistryUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func PoolRegistryUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (poolRegistry *PoolRegistry) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*PoolRegistryUUPSUnauthorizedCallContext, error) {
	out := new(PoolRegistryUUPSUnauthorizedCallContext)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRegistryUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the PoolRegistry contract.
type PoolRegistryUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func PoolRegistryUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (poolRegistry *PoolRegistry) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*PoolRegistryUUPSUnsupportedProxiableUUID, error) {
	out := new(PoolRegistryUUPSUnsupportedProxiableUUID)
	if err := poolRegistry.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
