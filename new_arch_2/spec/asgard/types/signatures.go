package types

import (
	ssz "github.com/ferranbt/fastssz"
)

func VerifyObjectSignature(
	signature [96]byte,
	rootObj ssz.HashRoot,
	domain Domain,
	signatureType SignatureType,
	signers []*Signer) error {
	r, err := rootObj.HashTreeRoot()
	if err != nil {
		return err
	}
	return VerifySignature(signature, r, domain, signatureType, signers)
}

func VerifySignature(
	signature [96]byte,
	root [32]byte,
	domain Domain,
	signatureType SignatureType,
	signers []*Signer) error {
	panic("implement")
}
