package qbft

import (
	"bytes"
	"github.com/pkg/errors"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// UponCommit checks if message is unique and transitions the state accordingly.
// Assumes commit message is valid!
func UponCommit(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if !uniqueSignerForRound(state, signedMessage) {
		return errors.New("duplicate message")
	}
	AddMessage(state, signedMessage)
	if IsCommitQuorum(state, share) {
		state.Decided = true
	}

	return nil
}

// CreateCommitMessage returns commit message
// This function is called each time state.PreparedRound == state.Round becomes true
func CreateCommitMessage(state *types.QBFT) (*types.QBFTMessage, error) {
	// TODO implement
	return &types.QBFTMessage{
		MsgType:    types.CommitMessageType,
		Round:      state.Round,
		Height:     state.Height,
		Identifier: state.Identifier, // TODO
		Root:       state.ProposalAcceptedForCurrentRound.Message.Root,
		DataRound:  state.Round, // TODO for future version omit this and keep only round?
	}, nil
}

func IsCommitQuorum(state *types.QBFT, share *types.Share) bool {
	commits := MessagesByRoundAndType(state, state.Round, types.CommitMessageType)
	return UniqueSignerQuorum(share.Quorum, commits)
}

// isValidCommit returns nil if commit message (not a decided message) is valid for state
func isValidCommit(state *types.QBFT, share *types.Share, signedMessage *types.QBFTSignedMessage) error {
	if err := baseCommitValidation(share, state.Height, signedMessage); err != nil {
		return err
	}

	if state.ProposalAcceptedForCurrentRound == nil {
		return errors.New("no proposal accepted for round")
	}

	if len(signedMessage.Signers) != 1 {
		return errors.New("msg allows 1 signer")
	}

	if signedMessage.Message.Round != state.Round {
		return errors.New("wrong msg round")
	}

	if !bytes.Equal(state.ProposalAcceptedForCurrentRound.Message.Root[:], signedMessage.Message.Root[:]) {
		return errors.New("proposed data mismatch")
	}

	return nil
}

// baseCommitValidation returns true if commit message (which can be a decided message as well) is valid
func baseCommitValidation(share *types.Share, height uint64, signedMessage *types.QBFTSignedMessage) error {
	if signedMessage.Message.MsgType != types.CommitMessageType {
		return errors.New("commit msg type is wrong")
	}
	if signedMessage.Message.Height != height {
		return errors.New("wrong msg height")
	}

	if err := signedMessage.Validate(); err != nil {
		return errors.Wrap(err, "signed commit invalid")
	}

	// verify signature
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
