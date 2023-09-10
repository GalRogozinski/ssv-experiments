// Code generated by fastssz. DO NOT EDIT.
// Hash: 2c1a28f2c42cd618e2ea7c3ab19cabb74d8766dd8ce33f5afb1b072cdee1ba80
// Version: 0.1.2
package qbft

import (
	ssz "github.com/ferranbt/fastssz"
	"ssv-experiments/new_arch_2/spec/asgard/types"
)

// MarshalSSZ ssz marshals the ProcessMessageTest object
func (p *ProcessMessageTest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the ProcessMessageTest object to a target array
func (p *ProcessMessageTest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(12)

	// Offset (0) 'Pre'
	dst = ssz.WriteOffset(dst, offset)
	if p.Pre == nil {
		p.Pre = new(types.QBFT)
	}
	offset += p.Pre.SizeSSZ()

	// Offset (1) 'Post'
	dst = ssz.WriteOffset(dst, offset)
	if p.Post == nil {
		p.Post = new(types.QBFT)
	}
	offset += p.Post.SizeSSZ()

	// Offset (2) 'Messages'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(p.Messages); ii++ {
		offset += 4
		offset += p.Messages[ii].SizeSSZ()
	}

	// Field (0) 'Pre'
	if dst, err = p.Pre.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'Post'
	if dst, err = p.Post.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (2) 'Messages'
	if size := len(p.Messages); size > 256 {
		err = ssz.ErrListTooBigFn("ProcessMessageTest.Messages", size, 256)
		return
	}
	{
		offset = 4 * len(p.Messages)
		for ii := 0; ii < len(p.Messages); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += p.Messages[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(p.Messages); ii++ {
		if dst, err = p.Messages[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ProcessMessageTest object
func (p *ProcessMessageTest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 12 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1, o2 uint64

	// Offset (0) 'Pre'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 12 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'Post'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Offset (2) 'Messages'
	if o2 = ssz.ReadOffset(buf[8:12]); o2 > size || o1 > o2 {
		return ssz.ErrOffset
	}

	// Field (0) 'Pre'
	{
		buf = tail[o0:o1]
		if p.Pre == nil {
			p.Pre = new(types.QBFT)
		}
		if err = p.Pre.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'Post'
	{
		buf = tail[o1:o2]
		if p.Post == nil {
			p.Post = new(types.QBFT)
		}
		if err = p.Post.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (2) 'Messages'
	{
		buf = tail[o2:]
		num, err := ssz.DecodeDynamicLength(buf, 256)
		if err != nil {
			return err
		}
		p.Messages = make([]*types.QBFTSignedMessage, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if p.Messages[indx] == nil {
				p.Messages[indx] = new(types.QBFTSignedMessage)
			}
			if err = p.Messages[indx].UnmarshalSSZ(buf); err != nil {
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

// SizeSSZ returns the ssz encoded size in bytes for the ProcessMessageTest object
func (p *ProcessMessageTest) SizeSSZ() (size int) {
	size = 12

	// Field (0) 'Pre'
	if p.Pre == nil {
		p.Pre = new(types.QBFT)
	}
	size += p.Pre.SizeSSZ()

	// Field (1) 'Post'
	if p.Post == nil {
		p.Post = new(types.QBFT)
	}
	size += p.Post.SizeSSZ()

	// Field (2) 'Messages'
	for ii := 0; ii < len(p.Messages); ii++ {
		size += 4
		size += p.Messages[ii].SizeSSZ()
	}

	return
}

// HashTreeRoot ssz hashes the ProcessMessageTest object
func (p *ProcessMessageTest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the ProcessMessageTest object with a hasher
func (p *ProcessMessageTest) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pre'
	if err = p.Pre.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Post'
	if err = p.Post.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (2) 'Messages'
	{
		subIndx := hh.Index()
		num := uint64(len(p.Messages))
		if num > 256 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range p.Messages {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 256)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ProcessMessageTest object
func (p *ProcessMessageTest) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(p)
}
