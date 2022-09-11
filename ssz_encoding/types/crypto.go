package types

type DomainType [4]byte

var (
	Shifu      = DomainType{0x0, 0x0, 0x0, 0x1}
	SSVMainnet = DomainType{0x1, 0x0, 0x0, 0x0}
)
