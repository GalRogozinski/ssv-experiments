// Code generated by fastssz. DO NOT EDIT.
// Hash: 63590a45d15e1b0c43c3d7deb9b4248a5b114b8163317ec0687516bc815415b6
// Version: 0.1.2
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
	offset := int(140)

	// Field (0) 'Role'
	dst = ssz.MarshalUint64(dst, d.Role)

	// Field (1) 'ValidatorPK'
	dst = append(dst, d.ValidatorPK[:]...)

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, d.Slot)

	// Field (3) 'DomainData'
	dst = append(dst, d.DomainData[:]...)

	// Field (4) 'ValidatorIndex'
	dst = ssz.MarshalUint64(dst, d.ValidatorIndex)

	// Field (5) 'CommitteeIndex'
	dst = ssz.MarshalUint64(dst, d.CommitteeIndex)

	// Field (6) 'CommitteeLength'
	dst = ssz.MarshalUint64(dst, d.CommitteeLength)

	// Field (7) 'CommitteesAtSlot'
	dst = ssz.MarshalUint64(dst, d.CommitteesAtSlot)

	// Field (8) 'ValidatorCommitteeIndex'
	dst = ssz.MarshalUint64(dst, d.ValidatorCommitteeIndex)

	// Offset (9) 'ValidatorSyncCommitteeIndices'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(d.ValidatorSyncCommitteeIndices) * 8

	// Field (9) 'ValidatorSyncCommitteeIndices'
	if size := len(d.ValidatorSyncCommitteeIndices); size > 13 {
		err = ssz.ErrListTooBigFn("Duty.ValidatorSyncCommitteeIndices", size, 13)
		return
	}
	for ii := 0; ii < len(d.ValidatorSyncCommitteeIndices); ii++ {
		dst = ssz.MarshalUint64(dst, d.ValidatorSyncCommitteeIndices[ii])
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Duty object
func (d *Duty) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 140 {
		return ssz.ErrSize
	}

	tail := buf
	var o9 uint64

	// Field (0) 'Role'
	d.Role = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'ValidatorPK'
	copy(d.ValidatorPK[:], buf[8:56])

	// Field (2) 'Slot'
	d.Slot = ssz.UnmarshallUint64(buf[56:64])

	// Field (3) 'DomainData'
	copy(d.DomainData[:], buf[64:96])

	// Field (4) 'ValidatorIndex'
	d.ValidatorIndex = ssz.UnmarshallUint64(buf[96:104])

	// Field (5) 'CommitteeIndex'
	d.CommitteeIndex = ssz.UnmarshallUint64(buf[104:112])

	// Field (6) 'CommitteeLength'
	d.CommitteeLength = ssz.UnmarshallUint64(buf[112:120])

	// Field (7) 'CommitteesAtSlot'
	d.CommitteesAtSlot = ssz.UnmarshallUint64(buf[120:128])

	// Field (8) 'ValidatorCommitteeIndex'
	d.ValidatorCommitteeIndex = ssz.UnmarshallUint64(buf[128:136])

	// Offset (9) 'ValidatorSyncCommitteeIndices'
	if o9 = ssz.ReadOffset(buf[136:140]); o9 > size {
		return ssz.ErrOffset
	}

	if o9 < 140 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (9) 'ValidatorSyncCommitteeIndices'
	{
		buf = tail[o9:]
		num, err := ssz.DivideInt2(len(buf), 8, 13)
		if err != nil {
			return err
		}
		d.ValidatorSyncCommitteeIndices = ssz.ExtendUint64(d.ValidatorSyncCommitteeIndices, num)
		for ii := 0; ii < num; ii++ {
			d.ValidatorSyncCommitteeIndices[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Duty object
func (d *Duty) SizeSSZ() (size int) {
	size = 140

	// Field (9) 'ValidatorSyncCommitteeIndices'
	size += len(d.ValidatorSyncCommitteeIndices) * 8

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

	// Field (3) 'DomainData'
	hh.PutBytes(d.DomainData[:])

	// Field (4) 'ValidatorIndex'
	hh.PutUint64(d.ValidatorIndex)

	// Field (5) 'CommitteeIndex'
	hh.PutUint64(d.CommitteeIndex)

	// Field (6) 'CommitteeLength'
	hh.PutUint64(d.CommitteeLength)

	// Field (7) 'CommitteesAtSlot'
	hh.PutUint64(d.CommitteesAtSlot)

	// Field (8) 'ValidatorCommitteeIndex'
	hh.PutUint64(d.ValidatorCommitteeIndex)

	// Field (9) 'ValidatorSyncCommitteeIndices'
	{
		if size := len(d.ValidatorSyncCommitteeIndices); size > 13 {
			err = ssz.ErrListTooBigFn("Duty.ValidatorSyncCommitteeIndices", size, 13)
			return
		}
		subIndx := hh.Index()
		for _, i := range d.ValidatorSyncCommitteeIndices {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(d.ValidatorSyncCommitteeIndices))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(13, numItems, 8))
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Duty object
func (d *Duty) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(d)
}
