// Code generated by fastssz. DO NOT EDIT.
// Hash: 8f95237b8156dac8f4798dd2cbaf2f2e928846dbebb53794bd61eba1064a770f
// Version: 0.1.3
package types

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Duty object
func (d *Duty) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Duty object to a target array
func (d *Duty) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Role'
	dst = ssz.MarshalUint64(dst, d.Role)

	// Field (1) 'ValidatorPK'
	dst = append(dst, d.ValidatorPK[:]...)

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, d.Slot)

	return
}

// UnmarshalSSZ ssz unmarshals the Duty object
func (d *Duty) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'Role'
	d.Role = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ValidatorPK'
	copy(d.ValidatorPK[:], buf[8:56])

	// Field (2) 'Slot'
	d.Slot = ssz.UnmarshallUint64(buf[56:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Duty object
func (d *Duty) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the Duty object
func (d *Duty) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Duty object with a hasher
func (d *Duty) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Role'
	hh.PutUint64(d.Role)

	// Field (1) 'ValidatorPK'
	hh.PutBytes(d.ValidatorPK[:])

	// Field (2) 'Slot'
	hh.PutUint64(d.Slot)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Duty object
func (d *Duty) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(d)
}