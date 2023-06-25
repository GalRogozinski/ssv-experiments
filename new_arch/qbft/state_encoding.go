// Code generated by fastssz. DO NOT EDIT.
// Hash: 1af11f59872834ead318275d1c725a886ae4c33a73cff9fb39076b79c3ca20d5
// Version: 0.1.3
package qbft

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the State object
func (s *State) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the State object to a target array
func (s *State) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(28)

	// Field (0) 'Round'
	dst = ssz.MarshalUint64(dst, s.Round)

	// Field (1) 'Height'
	dst = ssz.MarshalUint64(dst, s.Height)

	// Field (2) 'PreparedRound'
	dst = ssz.MarshalUint64(dst, s.PreparedRound)

	// Offset (3) 'Messages'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(s.Messages); ii++ {
		offset += 4
		offset += s.Messages[ii].SizeSSZ()
	}

	// Field (3) 'Messages'
	if size := len(s.Messages); size > 256 {
		err = ssz.ErrListTooBigFn("State.Messages", size, 256)
		return
	}
	{
		offset = 4 * len(s.Messages)
		for ii := 0; ii < len(s.Messages); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.Messages[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.Messages); ii++ {
		if dst, err = s.Messages[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the State object
func (s *State) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 28 {
		return ssz.ErrSize
	}

	tail := buf
	var o3 uint64

	// Field (0) 'Round'
	s.Round = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Height'
	s.Height = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'PreparedRound'
	s.PreparedRound = ssz.UnmarshallUint64(buf[16:24])

	// Offset (3) 'Messages'
	if o3 = ssz.ReadOffset(buf[24:28]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 28 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (3) 'Messages'
	{
		buf = tail[o3:]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.Messages = make([]*SignedMessage, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.Messages[indx] == nil {
				s.Messages[indx] = new(SignedMessage)
			}
			if err = s.Messages[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the State object
func (s *State) SizeSSZ() (size int) {
	size = 28

	// Field (3) 'Messages'
	for ii := 0; ii < len(s.Messages); ii++ {
		size += 4
		size += s.Messages[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the State object
func (s *State) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the State object with a hasher
func (s *State) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Round'
	hh.PutUint64(s.Round)

	// Field (1) 'Height'
	hh.PutUint64(s.Height)

	// Field (2) 'PreparedRound'
	hh.PutUint64(s.PreparedRound)

	// Field (3) 'Messages'
	{
		subIndx := hh.Index()
		num := uint64(len(s.Messages))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.Messages {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the State object
func (s *State) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
