package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ModuleName = "margin"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route
	QuerierRoute = ModuleName

	// RouterKey is the msg router key
	RouterKey = ModuleName
)

func NewMTP(signer string, collateralAsset string, borrowAsset string) *MTP {
	return &MTP{
		Address:          signer,
		CollateralAsset:  collateralAsset,
		CollateralAmount: sdk.ZeroUint(),
		LiabilitiesP:     sdk.ZeroUint(),
		LiabilitiesI:     sdk.ZeroUint(),
		CustodyAsset:     borrowAsset,
		CustodyAmount:    sdk.ZeroUint(),
		Leverage:         sdk.ZeroUint(),
		MtpHealth:        sdk.ZeroDec(),
	}
}

func (mtp MTP) Validate() error {
	if mtp.CollateralAsset == "" {
		return sdkerrors.Wrap(ErrMTPInvalid, "no asset specified")
	}
	if mtp.Address == "" {
		return sdkerrors.Wrap(ErrMTPInvalid, "no address specified")
	}
	if mtp.Id == 0 {
		return sdkerrors.Wrap(ErrMTPInvalid, "no id specified")
	}

	return nil
}

func GetSettlementAsset() string {
	return "rowan"
}
