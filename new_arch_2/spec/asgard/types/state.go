package types

type QBFT struct {
	Identifier []byte `ssz-max:"56"` // concatenation of domain, public key, and role

	Round  uint64
	Height uint64

	PreparedRound uint64

	// ProposalAcceptedForCurrentRound holds the accepted proposal for the current round
	ProposalAcceptedForCurrentRound *QBFTMessage // TODO this can simply be the fullData or the root

	// Messages is a unified (to all message type) container slice, simple and easy to serialize.
	// All messages in the container are verified and authenticated
	Messages []*QBFTMessage `ssz-max:"256"`

	// Stopped when true, can't process any messages
	Stopped bool

	// Decided when true, the QBFT instance finished and may stop processing messages
	Decided bool
}

func (qbft *QBFT) DecidedValue() *ConsensusData {
	panic("implement")
}

type State struct {
	// PartialSignatures holds partial BLS signatures
	PartialSignatures []*SignedPartialSignatureMessages `ssz-max:"256"`
	// DecidedValue holds the decided value set after consensus phase
	QBFT         *QBFT
	StartingDuty *Duty
}

func NewState(duty *Duty) *State {
	return &State{
		StartingDuty:      duty,
		PartialSignatures: []*SignedPartialSignatureMessages{},
	}
}
