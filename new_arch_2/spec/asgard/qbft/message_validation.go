package qbft

import (
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// ValidateMessage returns nil if message valid for state
func ValidateMessage(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if err := signedMessage.Validate(); err != nil {
		return err
	}

	if signedMessage.Message.Round < state.Round {
		return errors.New("past round")
	}

	if signedMessage.Message.Height != state.Height {
		return errors.New("wrong message height")
	}

	switch signedMessage.Message.MsgType {
	case types.ProposalMessageType:
		return isValidProposal(state, share, signedMessage)
	case types.PrepareMessageType:
		// TODO validSignedPrepareForHeightRoundAndRoot
		return nil
	case types.CommitMessageType:
		// TODO validateCommit
		return nil
	case types.RoundChangeMessageType:
		// TODO validRoundChangeForData
		return nil
	default:
		return errors.New("unknown message type")
	}
}
