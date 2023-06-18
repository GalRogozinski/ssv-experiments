// Code generated by fastssz. DO NOT EDIT.
// Hash: 26e205b5b9d92e91ddd76b466e5237d03d78c8dac1c7a71e4f9d110fce40091c
// Version: 0.1.2
package qbft

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the InputData object
func (i *InputData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(i)
}

// MarshalSSZTo ssz marshals the InputData object to a target array
func (i *InputData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Data'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(i.Data)

	// Field (0) 'Data'
	if size := len(i.Data); size > 4259840 {
		err = ssz.ErrBytesLengthFn("InputData.Data", size, 4259840)
		return
	}
	dst = append(dst, i.Data...)

	return
}

// UnmarshalSSZ ssz unmarshals the InputData object
func (i *InputData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Data'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Data'
	{
		buf = tail[o0:]
		if len(buf) > 4259840 {
			return ssz.ErrBytesLength
		}
		if cap(i.Data) == 0 {
			i.Data = make([]byte, 0, len(buf))
		}
		i.Data = append(i.Data, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the InputData object
func (i *InputData) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Data'
	size += len(i.Data)

	return
}

// HashTreeRoot ssz hashes the InputData object
func (i *InputData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(i)
}

// HashTreeRootWith ssz hashes the InputData object with a hasher
func (i *InputData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Data'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(i.Data))
		if byteLen > 4259840 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.PutBytes(i.Data)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (4259840+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the InputData object
func (i *InputData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(i)
}
