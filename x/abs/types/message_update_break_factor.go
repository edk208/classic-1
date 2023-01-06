package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBreakFactor = "update_break_factor"

var _ sdk.Msg = &MsgUpdateBreakFactor{}

func NewMsgUpdateBreakFactor(creator string, value string) *MsgUpdateBreakFactor {
	return &MsgUpdateBreakFactor{
		Creator: creator,
		Value:   value,
	}
}

func (msg *MsgUpdateBreakFactor) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBreakFactor) Type() string {
	return TypeMsgUpdateBreakFactor
}

func (msg *MsgUpdateBreakFactor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBreakFactor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBreakFactor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	value, err := sdk.NewDecFromStr(msg.Value)

	if err != nil || value.IsNegative() || value.GT(sdk.OneDec()) {
		return sdkerrors.Wrapf(ErrBreakFactorOutOfBounds, "invalid breakfactor value (%s)", err)
	}

	return nil
}
