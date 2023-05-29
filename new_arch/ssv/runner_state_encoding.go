// Code generated by fastssz. DO NOT EDIT.
// Hash: 98f69add302950d5cbfd80372d2d7b72649dc665651ca7e190fc3c794d829023
// Version: 0.1.2
package ssv

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/new_arch/types"
)

// MarshalSSZ ssz marshals the State object
func (s *State) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the State object to a target array
func (s *State) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(72)

	// Offset (0) 'PartialSignatures'
	dst = ssz.WriteOffset(dst, offset)
	if s.PartialSignatures == nil {
		s.PartialSignatures = new(PartialSignatureContainer)
	}
	offset += s.PartialSignatures.SizeSSZ()

	// Field (1) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if dst, err = s.StartingDuty.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (2) 'DecidedData'
	dst = ssz.WriteOffset(dst, offset)
	if s.DecidedData == nil {
		s.DecidedData = new(types.ConsensusData)
	}
	offset += s.DecidedData.SizeSSZ()

	// Field (0) 'PartialSignatures'
	if dst, err = s.PartialSignatures.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'DecidedData'
	if dst, err = s.DecidedData.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the State object
func (s *State) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 72 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o2 uint64

	// Offset (0) 'PartialSignatures'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 72 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if err = s.StartingDuty.UnmarshalSSZ(buf[4:68]); err != nil {
		return err
	}

	// Offset (2) 'DecidedData'
	if o2 = ssz.ReadOffset(buf[68:72]); o2 > size || o0 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'PartialSignatures'
	{
		buf = tail[o0:o2]
		if s.PartialSignatures == nil {
			s.PartialSignatures = new(PartialSignatureContainer)
		}
		if err = s.PartialSignatures.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'DecidedData'
	{
		buf = tail[o2:]
		if s.DecidedData == nil {
			s.DecidedData = new(types.ConsensusData)
		}
		if err = s.DecidedData.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the State object
func (s *State) SizeSSZ() (size int) {
	size = 72

	// Field (0) 'PartialSignatures'
	if s.PartialSignatures == nil {
		s.PartialSignatures = new(PartialSignatureContainer)
	}
	size += s.PartialSignatures.SizeSSZ()

	// Field (2) 'DecidedData'
	if s.DecidedData == nil {
		s.DecidedData = new(types.ConsensusData)
	}
	size += s.DecidedData.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the State object
func (s *State) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the State object with a hasher
func (s *State) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'PartialSignatures'
	if err = s.PartialSignatures.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if err = s.StartingDuty.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'DecidedData'
	if err = s.DecidedData.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the State object
func (s *State) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
