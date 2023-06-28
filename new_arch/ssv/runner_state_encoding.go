// Code generated by fastssz. DO NOT EDIT.
// Hash: 45a3c29b3f3990f9c75792a5175cad6af3e33aa51586d13f04caa1434bafb0b2
// Version: 0.1.3
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
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		offset += 4
		offset += s.PartialSignatures[ii].SizeSSZ()
	}

	// Offset (1) 'DecidedValue'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(s.DecidedValue)

	// Field (2) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if dst, err = s.StartingDuty.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (0) 'PartialSignatures'
	if size := len(s.PartialSignatures); size > 256 {
		err = ssz.ErrListTooBigFn("State.PartialSignatures", size, 256)
		return
	}
	{
		offset = 4 * len(s.PartialSignatures)
		for ii := 0; ii < len(s.PartialSignatures); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += s.PartialSignatures[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		if dst, err = s.PartialSignatures[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (1) 'DecidedValue'
	if size := len(s.DecidedValue); size > 8388608 {
		err = ssz.ErrBytesLengthFn("State.DecidedValue", size, 8388608)
		return
	}
	dst = append(dst, s.DecidedValue...)

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
	var o0, o1 uint64

	// Offset (0) 'PartialSignatures'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 72 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'DecidedValue'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (2) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if err = s.StartingDuty.UnmarshalSSZ(buf[8:72]); err != nil {
		return err
	}

	// Field (0) 'PartialSignatures'
	{
		buf = tail[o0:o1]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		s.PartialSignatures = make([]*types.SignedPartialSignatureMessages, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if s.PartialSignatures[indx] == nil {
				s.PartialSignatures[indx] = new(types.SignedPartialSignatureMessages)
			}
			if err = s.PartialSignatures[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (1) 'DecidedValue'
	{
		buf = tail[o1:]
		if len(buf) > 8388608 {
			return ssz.ErrBytesLength
		}
		if cap(s.DecidedValue) == 0 {
			s.DecidedValue = make([]byte, 0, len(buf))
		}
		s.DecidedValue = append(s.DecidedValue, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the State object
func (s *State) SizeSSZ() (size int) {
	size = 72

	// Field (0) 'PartialSignatures'
	for ii := 0; ii < len(s.PartialSignatures); ii++ {
		size += 4
		size += s.PartialSignatures[ii].SizeSSZ()
	}

	// Field (1) 'DecidedValue'
	size += len(s.DecidedValue)

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
	{
		subIndx := hh.Index()
		num := uint64(len(s.PartialSignatures))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range s.PartialSignatures {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	// Field (1) 'DecidedValue'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(s.DecidedValue))
		if byteLen > 8388608 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(s.DecidedValue)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (8388608+31)/32)
	}

	// Field (2) 'StartingDuty'
	if s.StartingDuty == nil {
		s.StartingDuty = new(types.Duty)
	}
	if err = s.StartingDuty.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the State object
func (s *State) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
