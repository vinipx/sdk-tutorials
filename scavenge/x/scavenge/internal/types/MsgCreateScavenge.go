package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgCreateScavenge
// ------------------------------------------------------------------------------
var _ sdk.Msg = &MsgCreateScavenge{}

// MsgCreateScavenge - struct for unjailing jailed validator
type MsgCreateScavenge struct {
	Creator        sdk.AccAddress `json:"creator" yaml:"creator"`               // address of the scavenger creator
	Description    string         `json:"description" yaml:"description"`       // description of the scavenge
	SolutionHash   string         `json:"solutionHash" yaml:"solutionHash"`     // solution hash of the scavenge
	CoinReward     sdk.Coins      `json:"rewardCoins" yaml:"rewardCoins"`       // coin reward of the scavenger
	NFTRewardDenom string         `json:"NFTRewardDenom" yaml:"NFTRewardDenom"` // NFT reward of the scavenger
	NFTRewardID    string         `json:"NFTRewardID" yaml:"NFTRewardID"`       // NFT reward of the scavenger
}

// NewMsgCreateScavenge creates a new MsgCreateScavenge instance
func NewMsgCreateScavenge(creator sdk.AccAddress, description, solutionHash string, coinReward sdk.Coins, nftRewardDenom, nftRewardID string) MsgCreateScavenge {
	return MsgCreateScavenge{
		Creator:        creator,
		Description:    description,
		SolutionHash:   solutionHash,
		CoinReward:     coinReward,
		NFTRewardDenom: nftRewardDenom,
		NFTRewardID:    nftRewardID,
	}
}

// CreateScavengeConst is CreateScavenge Constant
const CreateScavengeConst = "CreateScavenge"

// nolint
func (msg MsgCreateScavenge) Route() string { return RouterKey }
func (msg MsgCreateScavenge) Type() string  { return CreateScavengeConst }
func (msg MsgCreateScavenge) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateScavenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateScavenge) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionScavengerHash can't be empty")
	}
	if (!msg.CoinReward.IsValid() || !msg.CoinReward.IsZero()) && (msg.NFTRewardDenom == "" || msg.NFTRewardID == "") {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Must include a valid NFT Reward or a Coin Reward")
	}
	return nil
}
