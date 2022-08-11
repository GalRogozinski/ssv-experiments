// Code generated by fastssz. DO NOT EDIT.
// Hash: a368e5af1b248b4e6b433e96bc69d5a864f25b397b12a7a3ce488d095c214d6e
package ssv

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the PartialSignature object
func (p *PartialSignature) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the PartialSignature object to a target array
func (p *PartialSignature) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	return
}

// UnmarshalSSZ ssz unmarshals the PartialSignature object
func (p *PartialSignature) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the PartialSignature object
func (p *PartialSignature) SizeSSZ() (size int) {
	size = 4
	return
}

// HashTreeRoot ssz hashes the PartialSignature object
func (p *PartialSignature) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PartialSignature object with a hasher
func (p *PartialSignature) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

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
	offset := int(4)

	return
}

// UnmarshalSSZ ssz unmarshals the PartialSignatures object
func (p *PartialSignatures) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the PartialSignatures object
func (p *PartialSignatures) SizeSSZ() (size int) {
	size = 4
	return
}

// HashTreeRoot ssz hashes the PartialSignatures object
func (p *PartialSignatures) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the PartialSignatures object with a hasher
func (p *PartialSignatures) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

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
	offset := int(4)

	return
}

// UnmarshalSSZ ssz unmarshals the SignedPartialSignatures object
func (s *SignedPartialSignatures) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SignedPartialSignatures object
func (s *SignedPartialSignatures) SizeSSZ() (size int) {
	size = 4
	return
}

// HashTreeRoot ssz hashes the SignedPartialSignatures object
func (s *SignedPartialSignatures) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SignedPartialSignatures object with a hasher
func (s *SignedPartialSignatures) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SignedPartialSignatures object
func (s *SignedPartialSignatures) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
