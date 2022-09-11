// Code generated by fastssz. DO NOT EDIT.
// Hash: 1ae1a9f70cce6e91b5775e06c9a395bc086172baf803937c022ab7cb96bc3171
package ssv

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/ssz_encoding/qbft"
)

// MarshalSSZ ssz marshals the PartialSignature object
func (p *PartialSignature) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the PartialSignature object to a target array
func (p *PartialSignature) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(140)

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, p.Slot)

	// Field (1) 'Signature'
	dst = append(dst, p.Signature[:]...)

	// Field (2) 'SigningRoot'
	dst = append(dst, p.SigningRoot[:]...)

	// Offset (3) 'Justification'
	dst = ssz.WriteOffset(dst, offset)
	if p.Justification == nil {
		p.Justification = new(qbft.SignedMessageHeader)
	}
	offset += p.Justification.SizeSSZ()

	// Field (3) 'Justification'
	if dst, err = p.Justification.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the PartialSignature object
func (p *PartialSignature) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 140 {
		return ssz.ErrSize
	}

	tail := buf
	var o3 uint64

	// Field (0) 'Slot'
	p.Slot = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Signature'
	copy(p.Signature[:], buf[8:104])

	// Field (2) 'SigningRoot'
	copy(p.SigningRoot[:], buf[104:136])

	// Offset (3) 'Justification'
	if o3 = ssz.ReadOffset(buf[136:140]); o3 > size {
		return ssz.ErrOffset
	}

	if o3 < 140 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (3) 'Justification'
	{
		buf = tail[o3:]
		if p.Justification == nil {
			p.Justification = new(qbft.SignedMessageHeader)
		}
		if err = p.Justification.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the PartialSignature object
func (p *PartialSignature) SizeSSZ() (size int) {
	size = 140

	// Field (3) 'Justification'
	if p.Justification == nil {
		p.Justification = new(qbft.SignedMessageHeader)
	}
	size += p.Justification.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the PartialSignature object
func (p *PartialSignature) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PartialSignature object with a hasher
func (p *PartialSignature) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(p.Slot)

	// Field (1) 'Signature'
	hh.PutBytes(p.Signature[:])

	// Field (2) 'SigningRoot'
	hh.PutBytes(p.SigningRoot[:])

	// Field (3) 'Justification'
	if err = p.Justification.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the PartialSignature object
func (p *PartialSignature) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(p)
}

// MarshalSSZ ssz marshals the PartialSignatures object
func (p *PartialSignatures) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the PartialSignatures object to a target array
func (p *PartialSignatures) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Field (0) 'Type'
	dst = ssz.MarshalUint64(dst, uint64(p.Type))

	// Offset (1) 'PartialSignatures'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(p.PartialSignatures); ii++ {
		offset += 4
		offset += p.PartialSignatures[ii].SizeSSZ()
	}

	// Field (1) 'PartialSignatures'
	if size := len(p.PartialSignatures); size > 13 {
		err = ssz.ErrListTooBigFn("PartialSignatures.PartialSignatures", size, 13)
		return
	}
	{
		offset = 4 * len(p.PartialSignatures)
		for ii := 0; ii < len(p.PartialSignatures); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += p.PartialSignatures[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(p.PartialSignatures); ii++ {
		if dst, err = p.PartialSignatures[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the PartialSignatures object
func (p *PartialSignatures) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o1 uint64

	// Field (0) 'Type'
	p.Type = PartialSigMsgType(ssz.UnmarshallUint64(buf[0:8]))

	// Offset (1) 'PartialSignatures'
	if o1 = ssz.ReadOffset(buf[8:12]); o1 > size {
		return ssz.ErrOffset
	}

	if o1 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'PartialSignatures'
	{
		buf = tail[o1:]
		num, err := ssz.DecodeDynamicLength(buf, 13)
		if err != nil {
			return err
		}
		p.PartialSignatures = make([]*PartialSignature, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if p.PartialSignatures[indx] == nil {
				p.PartialSignatures[indx] = new(PartialSignature)
			}
			if err = p.PartialSignatures[indx].UnmarshalSSZ(buf); err != nil {
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

// SizeSSZ returns the ssz encoded size in bytes for the PartialSignatures object
func (p *PartialSignatures) SizeSSZ() (size int) {
	size = 12

	// Field (1) 'PartialSignatures'
	for ii := 0; ii < len(p.PartialSignatures); ii++ {
		size += 4
		size += p.PartialSignatures[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the PartialSignatures object
func (p *PartialSignatures) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PartialSignatures object with a hasher
func (p *PartialSignatures) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Type'
	hh.PutUint64(uint64(p.Type))

	// Field (1) 'PartialSignatures'
	{
		subIndx := hh.Index()
		num := uint64(len(p.PartialSignatures))
		if num > 13 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range p.PartialSignatures {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 13)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the PartialSignatures object
func (p *PartialSignatures) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(p)
}

// MarshalSSZ ssz marshals the SignedPartialSignatures object
func (s *SignedPartialSignatures) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SignedPartialSignatures object to a target array
func (s *SignedPartialSignatures) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(108)

	// Offset (0) 'PartialSignatures'
	dst = ssz.WriteOffset(dst, offset)
	offset += s.PartialSignatures.SizeSSZ()

	// Field (1) 'Signature'
	dst = append(dst, s.Signature[:]...)

	// Field (2) 'Signer'
	dst = ssz.MarshalUint64(dst, s.Signer)

	// Field (0) 'PartialSignatures'
	if dst, err = s.PartialSignatures.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the SignedPartialSignatures object
func (s *SignedPartialSignatures) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 108 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'PartialSignatures'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 108 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Signature'
	copy(s.Signature[:], buf[4:100])

	// Field (2) 'Signer'
	s.Signer = ssz.UnmarshallUint64(buf[100:108])

	// Field (0) 'PartialSignatures'
	{
		buf = tail[o0:]
		if err = s.PartialSignatures.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedPartialSignatures object
func (s *SignedPartialSignatures) SizeSSZ() (size int) {
	size = 108

	// Field (0) 'PartialSignatures'
	size += s.PartialSignatures.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the SignedPartialSignatures object
func (s *SignedPartialSignatures) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedPartialSignatures object with a hasher
func (s *SignedPartialSignatures) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'PartialSignatures'
	if err = s.PartialSignatures.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Signature'
	hh.PutBytes(s.Signature[:])

	// Field (2) 'Signer'
	hh.PutUint64(s.Signer)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedPartialSignatures object
func (s *SignedPartialSignatures) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
