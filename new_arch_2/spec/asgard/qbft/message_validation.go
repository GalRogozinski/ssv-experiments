package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// ValidateSignedMessage returns nil if signed message has a valid signature and correct number of signers
// It may check more conditions depending on the message type
func ValidateSignedMessage(signedMessage *types.QBFTSignedMessage, share types.Share) error {
	if len(signedMessage.Signers) != 1 {
		return errors.New("msg allows 1 signer")
	}
	if err := types.VerifyObjectSignature(
		signedMessage.Signature,
		signedMessage,
		share.Domain,
		types.QBFTSignatureType,
		share.Cluster,
	); err != nil {
		return err
	}
	return nil
}

// ValidateMessage returns nil if message valid for state
func ValidateMessage(state *types.QBFT, share *types.Share, qbftMessage *types.QBFTMessage) error {
	if err := qbftMessage.Validate(); err != nil {
		return err
	}

	if qbftMessage.Round < state.Round {
		return errors.New("past round")
	}

	if qbftMessage.Height != state.Height {
		return errors.New("wrong message height")
	}

	if !bytes.Equal(qbftMessage.Identifier, state.Identifier) {
		return errors.New("wrong message identifier")
	}

	switch qbftMessage.MsgType {
	case types.ProposalMessageType:
		return ValidateProposal(state, share, qbftMessage)
	case types.PrepareMessageType:
		if state.ProposalAcceptedForCurrentRound == nil {
			return errors.New("no proposal accepted for round")
		}
		return isValidPrepare(qbftMessage, state)
	case types.CommitMessageType:
		return isValidCommit(state, share, qbftMessage)
	case types.RoundChangeMessageType:
		// TODO validRoundChangeForData
		return nil
	default:
		return errors.New("unknown message type")
	}
}
