package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	commontypes "github.com/dymensionxyz/dymension/v3/x/common/types"
)

type DelayedAckHooks interface {
	AfterPacketStatusUpdated(ctx sdk.Context, packet *commontypes.RollappPacket, oldPacketKey string, newPacketKey string) error
	AfterPacketDeleted(ctx sdk.Context, rollappPacket *commontypes.RollappPacket)
}

type MultiDelayedAckHooks []DelayedAckHooks

var _ DelayedAckHooks = MultiDelayedAckHooks{}

func NewMultiDelayedAckHooks(hooks ...DelayedAckHooks) MultiDelayedAckHooks {
	return hooks
}

func (h MultiDelayedAckHooks) AfterPacketStatusUpdated(ctx sdk.Context, packet *commontypes.RollappPacket, oldPacketKey string, newPacketKey string) error {
	for i := range h {
		err := h[i].AfterPacketStatusUpdated(ctx, packet, oldPacketKey, newPacketKey)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h MultiDelayedAckHooks) AfterPacketDeleted(ctx sdk.Context, rollappPacket *commontypes.RollappPacket) {
	for i := range h {
		h[i].AfterPacketDeleted(ctx, rollappPacket)
	}
}

type BaseDelayedAckHook struct{}

var _ DelayedAckHooks = BaseDelayedAckHook{}

func (b BaseDelayedAckHook) AfterPacketStatusUpdated(ctx sdk.Context, packet *commontypes.RollappPacket, oldPacketKey string, newPacketKey string) error {
	return nil
}

func (b BaseDelayedAckHook) AfterPacketDeleted(sdk.Context, *commontypes.RollappPacket) {
}
