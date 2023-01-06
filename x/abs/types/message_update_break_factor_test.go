package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/terra-money/core/testutil/sample"
)

func TestMsgUpdateBreakFactor_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateBreakFactor
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateBreakFactor{
				Creator: "invalid_address",
				Value:   "0",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateBreakFactor{
				Creator: sample.AccAddress(),
				Value:   "0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
