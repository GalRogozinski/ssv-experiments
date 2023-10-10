package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch/qbft"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// UponPrepare process a pre-validated prepare message
func UponPrepare(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if !uniqueSignerForRound(state, signedMessage) {
		return errors.New("duplicate message")
	}
	AddMessage(state, signedMessage)

	if PrepareQuorum(state, share) {
		state.PreparedRound = state.Round
	}

	return nil
}

// CreatePrepareMessage returns unsigned prepare message
func CreatePrepareMessage(state *types.QBFT) *types.QBFTMessage {
	return &types.QBFTMessage{
		MsgType:                  types.PrepareMessageType,
		Round:                    state.Round,
		Height:                   state.Height,
		Identifier:               nil,
		Root:                     state.ProposalAcceptedForCurrentRound.Root,
		DataRound:                0,
		RoundChangeJustification: nil,
		PrepareJustification:     nil,
	}
}

func PrepareQuorum(state *types.QBFT, share *types.Share) bool {
	all := MessagesByRoundAndType(state, state.Round, types.PrepareMessageType)
	return UniqueSignerQuorum(share.Quorum, all)
}

// ValidateSignedPrepare returns nil if signed prepare message has a valid signature and correct number of signers
func ValidateSignedPrepare(signedMessage *qbft.SignedMessage, share types.Share) error {
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

// isValidPrepare returns nil if prepare message is valid for state's height and round
func isValidPrepare(qbftMessage *types.QBFTMessage, state *types.QBFT) error {
	if qbftMessage.MsgType != types.PrepareMessageType {
		return errors.New("prepare msg type is wrong")
	}

	if qbftMessage.Height != state.Height {
		return errors.New("wrong msg height")
	}
	if qbftMessage.Round != state.Round {
		return errors.New("wrong msg round")
	}

	if err := qbftMessage.Validate(); err != nil {
		return errors.Wrap(err, "prepareData invalid")
	}

	if !bytes.Equal(qbftMessage.Root[:], state.ProposalAcceptedForCurrentRound.Root[:]) {
		return errors.New("proposed data mistmatch")
	}

	return nil
}
