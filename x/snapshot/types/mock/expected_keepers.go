// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/x/snapshot/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"sync"
)

// Ensure, that StakingKeeperMock does implement types.StakingKeeper.
// If this is not the case, regenerate this file with moq.
var _ types.StakingKeeper = &StakingKeeperMock{}

// StakingKeeperMock is a mock implementation of types.StakingKeeper.
//
// 	func TestSomethingThatUsesStakingKeeper(t *testing.T) {
//
// 		// make and configure a mocked types.StakingKeeper
// 		mockedStakingKeeper := &StakingKeeperMock{
// 			GetLastTotalPowerFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int {
// 				panic("mock out the GetLastTotalPower method")
// 			},
// 			IterateBondedValidatorsByPowerFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, fn func(index int64, validator stakingtypes.ValidatorI) (stop bool))  {
// 				panic("mock out the IterateBondedValidatorsByPower method")
// 			},
// 			PowerReductionFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int {
// 				panic("mock out the PowerReduction method")
// 			},
// 			ValidatorFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) stakingtypes.ValidatorI {
// 				panic("mock out the Validator method")
// 			},
// 		}
//
// 		// use mockedStakingKeeper in code that requires types.StakingKeeper
// 		// and then make assertions.
//
// 	}
type StakingKeeperMock struct {
	// GetLastTotalPowerFunc mocks the GetLastTotalPower method.
	GetLastTotalPowerFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int

	// IterateBondedValidatorsByPowerFunc mocks the IterateBondedValidatorsByPower method.
	IterateBondedValidatorsByPowerFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, fn func(index int64, validator stakingtypes.ValidatorI) (stop bool))

	// PowerReductionFunc mocks the PowerReduction method.
	PowerReductionFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int

	// ValidatorFunc mocks the Validator method.
	ValidatorFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) stakingtypes.ValidatorI

	// calls tracks calls to the methods.
	calls struct {
		// GetLastTotalPower holds details about calls to the GetLastTotalPower method.
		GetLastTotalPower []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
		}
		// IterateBondedValidatorsByPower holds details about calls to the IterateBondedValidatorsByPower method.
		IterateBondedValidatorsByPower []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Fn is the fn argument value.
			Fn func(index int64, validator stakingtypes.ValidatorI) (stop bool)
		}
		// PowerReduction holds details about calls to the PowerReduction method.
		PowerReduction []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
		}
		// Validator holds details about calls to the Validator method.
		Validator []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Addr is the addr argument value.
			Addr github_com_cosmos_cosmos_sdk_types.ValAddress
		}
	}
	lockGetLastTotalPower              sync.RWMutex
	lockIterateBondedValidatorsByPower sync.RWMutex
	lockPowerReduction                 sync.RWMutex
	lockValidator                      sync.RWMutex
}

// GetLastTotalPower calls GetLastTotalPowerFunc.
func (mock *StakingKeeperMock) GetLastTotalPower(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int {
	if mock.GetLastTotalPowerFunc == nil {
		panic("StakingKeeperMock.GetLastTotalPowerFunc: method is nil but StakingKeeper.GetLastTotalPower was just called")
	}
	callInfo := struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetLastTotalPower.Lock()
	mock.calls.GetLastTotalPower = append(mock.calls.GetLastTotalPower, callInfo)
	mock.lockGetLastTotalPower.Unlock()
	return mock.GetLastTotalPowerFunc(ctx)
}

// GetLastTotalPowerCalls gets all the calls that were made to GetLastTotalPower.
// Check the length with:
//     len(mockedStakingKeeper.GetLastTotalPowerCalls())
func (mock *StakingKeeperMock) GetLastTotalPowerCalls() []struct {
	Ctx github_com_cosmos_cosmos_sdk_types.Context
} {
	var calls []struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
	}
	mock.lockGetLastTotalPower.RLock()
	calls = mock.calls.GetLastTotalPower
	mock.lockGetLastTotalPower.RUnlock()
	return calls
}

// IterateBondedValidatorsByPower calls IterateBondedValidatorsByPowerFunc.
func (mock *StakingKeeperMock) IterateBondedValidatorsByPower(ctx github_com_cosmos_cosmos_sdk_types.Context, fn func(index int64, validator stakingtypes.ValidatorI) (stop bool)) {
	if mock.IterateBondedValidatorsByPowerFunc == nil {
		panic("StakingKeeperMock.IterateBondedValidatorsByPowerFunc: method is nil but StakingKeeper.IterateBondedValidatorsByPower was just called")
	}
	callInfo := struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
		Fn  func(index int64, validator stakingtypes.ValidatorI) (stop bool)
	}{
		Ctx: ctx,
		Fn:  fn,
	}
	mock.lockIterateBondedValidatorsByPower.Lock()
	mock.calls.IterateBondedValidatorsByPower = append(mock.calls.IterateBondedValidatorsByPower, callInfo)
	mock.lockIterateBondedValidatorsByPower.Unlock()
	mock.IterateBondedValidatorsByPowerFunc(ctx, fn)
}

// IterateBondedValidatorsByPowerCalls gets all the calls that were made to IterateBondedValidatorsByPower.
// Check the length with:
//     len(mockedStakingKeeper.IterateBondedValidatorsByPowerCalls())
func (mock *StakingKeeperMock) IterateBondedValidatorsByPowerCalls() []struct {
	Ctx github_com_cosmos_cosmos_sdk_types.Context
	Fn  func(index int64, validator stakingtypes.ValidatorI) (stop bool)
} {
	var calls []struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
		Fn  func(index int64, validator stakingtypes.ValidatorI) (stop bool)
	}
	mock.lockIterateBondedValidatorsByPower.RLock()
	calls = mock.calls.IterateBondedValidatorsByPower
	mock.lockIterateBondedValidatorsByPower.RUnlock()
	return calls
}

// PowerReduction calls PowerReductionFunc.
func (mock *StakingKeeperMock) PowerReduction(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int {
	if mock.PowerReductionFunc == nil {
		panic("StakingKeeperMock.PowerReductionFunc: method is nil but StakingKeeper.PowerReduction was just called")
	}
	callInfo := struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
	}{
		Ctx: ctx,
	}
	mock.lockPowerReduction.Lock()
	mock.calls.PowerReduction = append(mock.calls.PowerReduction, callInfo)
	mock.lockPowerReduction.Unlock()
	return mock.PowerReductionFunc(ctx)
}

// PowerReductionCalls gets all the calls that were made to PowerReduction.
// Check the length with:
//     len(mockedStakingKeeper.PowerReductionCalls())
func (mock *StakingKeeperMock) PowerReductionCalls() []struct {
	Ctx github_com_cosmos_cosmos_sdk_types.Context
} {
	var calls []struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
	}
	mock.lockPowerReduction.RLock()
	calls = mock.calls.PowerReduction
	mock.lockPowerReduction.RUnlock()
	return calls
}

// Validator calls ValidatorFunc.
func (mock *StakingKeeperMock) Validator(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) stakingtypes.ValidatorI {
	if mock.ValidatorFunc == nil {
		panic("StakingKeeperMock.ValidatorFunc: method is nil but StakingKeeper.Validator was just called")
	}
	callInfo := struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Addr github_com_cosmos_cosmos_sdk_types.ValAddress
	}{
		Ctx:  ctx,
		Addr: addr,
	}
	mock.lockValidator.Lock()
	mock.calls.Validator = append(mock.calls.Validator, callInfo)
	mock.lockValidator.Unlock()
	return mock.ValidatorFunc(ctx, addr)
}

// ValidatorCalls gets all the calls that were made to Validator.
// Check the length with:
//     len(mockedStakingKeeper.ValidatorCalls())
func (mock *StakingKeeperMock) ValidatorCalls() []struct {
	Ctx  github_com_cosmos_cosmos_sdk_types.Context
	Addr github_com_cosmos_cosmos_sdk_types.ValAddress
} {
	var calls []struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Addr github_com_cosmos_cosmos_sdk_types.ValAddress
	}
	mock.lockValidator.RLock()
	calls = mock.calls.Validator
	mock.lockValidator.RUnlock()
	return calls
}
