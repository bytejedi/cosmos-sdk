package types

import (
	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
)

// CommitPacket return the hash of commitment bytes
// TODO: no specification for packet commitment currently,
// make it spec compatible once we have it
func CommitPacket(packet exported.PacketI) []byte {
	buf := sdk.Uint64ToBigEndian(packet.GetTimeoutHeight())
	buf = append(buf, packet.GetData()...)
	return tmhash.Sum(buf)
}

// CommitAcknowledgement returns the hash of commitment bytes
func CommitAcknowledgement(data []byte) []byte {
	return tmhash.Sum(data)
}

var _ exported.PacketI = Packet{}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	Data []byte `json:"data" yaml:"data"` // Actual opaque bytes transferred directly to the application module

	Sequence           uint64 `json:"sequence" yaml:"sequence"`                       // number corresponds to the order of sends and receives, where a Packet with an earlier sequence number must be sent and received before a Packet with a later sequence number.
	SourcePort         string `json:"source_port" yaml:"source_port"`                 // identifies the port on the sending chain.
	SourceChannel      string `json:"source_channel" yaml:"source_channel"`           // identifies the channel end on the sending chain.
	DestinationPort    string `json:"destination_port" yaml:"destination_port"`       // identifies the port on the receiving chain.
	DestinationChannel string `json:"destination_channel" yaml:"destination_channel"` // identifies the channel end on the receiving chain.
	TimeoutHeight      uint64 `json:"timeout_height" yaml:"timeout_height"`           // block height after which the packet times out
}

// NewPacket creates a new Packet instance
func NewPacket(
	data []byte,
	sequence uint64, sourcePort, sourceChannel,
	destinationPort, destinationChannel string,
	timeoutHeight uint64,
) Packet {
	return Packet{
		Data:               data,
		Sequence:           sequence,
		SourcePort:         sourcePort,
		SourceChannel:      sourceChannel,
		DestinationPort:    destinationPort,
		DestinationChannel: destinationChannel,
		TimeoutHeight:      timeoutHeight,
	}
}

// GetSequence implements PacketI interface
func (p Packet) GetSequence() uint64 { return p.Sequence }

// GetSourcePort implements PacketI interface
func (p Packet) GetSourcePort() string { return p.SourcePort }

// GetSourceChannel implements PacketI interface
func (p Packet) GetSourceChannel() string { return p.SourceChannel }

// GetDestPort implements PacketI interface
func (p Packet) GetDestPort() string { return p.DestinationPort }

// GetDestChannel implements PacketI interface
func (p Packet) GetDestChannel() string { return p.DestinationChannel }

// GetData implements PacketI interface
func (p Packet) GetData() []byte { return p.Data }

// GetTimeoutHeight implements PacketI interface
func (p Packet) GetTimeoutHeight() uint64 { return p.TimeoutHeight }

// ValidateBasic implements PacketI interface
func (p Packet) ValidateBasic() error {
	if err := host.DefaultPortIdentifierValidator(p.SourcePort); err != nil {
		return sdkerrors.Wrapf(
			ErrInvalidPacket,
			sdkerrors.Wrapf(err, "invalid source port ID: %s", p.SourcePort).Error(),
		)
	}
	if err := host.DefaultPortIdentifierValidator(p.DestinationPort); err != nil {
		return sdkerrors.Wrapf(
			ErrInvalidPacket,
			sdkerrors.Wrapf(err, "invalid destination port ID: %s", p.DestinationPort).Error(),
		)
	}
	if err := host.DefaultChannelIdentifierValidator(p.SourceChannel); err != nil {
		return sdkerrors.Wrapf(
			ErrInvalidPacket,
			sdkerrors.Wrapf(err, "invalid source channel ID: %s", p.SourceChannel).Error(),
		)
	}
	if err := host.DefaultChannelIdentifierValidator(p.DestinationChannel); err != nil {
		return sdkerrors.Wrapf(
			ErrInvalidPacket,
			sdkerrors.Wrapf(err, "invalid destination channel ID: %s", p.DestinationChannel).Error(),
		)
	}
	if p.Sequence == 0 {
		return sdkerrors.Wrap(ErrInvalidPacket, "packet sequence cannot be 0")
	}
	if p.TimeoutHeight == 0 {
		return sdkerrors.Wrap(ErrInvalidPacket, "packet timeout cannot be 0")
	}
	if len(p.Data) == 0 {
		return sdkerrors.Wrap(ErrInvalidPacket, "packet data bytes cannot be empty")
	}
	return nil
}
