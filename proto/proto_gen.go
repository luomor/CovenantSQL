package proto

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z DatabaseID) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendString(o, string(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DatabaseID) Msgsize() (s int) {
	s = hsp.StringPrefixSize + len(string(z))
	return
}

// MarshalHash marshals for hash
func (z *Envelope) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 4
	o = append(o, 0x84, 0x84)
	o = hsp.AppendString(o, z.Version)
	o = append(o, 0x84)
	o = hsp.AppendInt64(o, int64(z.TTL))
	o = append(o, 0x84)
	o = hsp.AppendInt64(o, int64(z.Expire))
	o = append(o, 0x84)
	if z.NodeID == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.NodeID.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Envelope) Msgsize() (s int) {
	s = 1 + 8 + hsp.StringPrefixSize + len(z.Version) + 4 + hsp.Int64Size + 7 + hsp.Int64Size + 7
	if z.NodeID == nil {
		s += hsp.NilSize
	} else {
		s += z.NodeID.Msgsize()
	}
	return
}

// MarshalHash marshals for hash
func (z *FindNeighborReq) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 4
	o = append(o, 0x84, 0x84)
	if oTemp, err := z.NodeID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x84)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Roles)))
	for za0001 := range z.Roles {
		if oTemp, err := z.Roles[za0001].MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x84)
	o = hsp.AppendInt(o, z.Count)
	o = append(o, 0x84)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FindNeighborReq) Msgsize() (s int) {
	s = 1 + 7 + z.NodeID.Msgsize() + 6 + hsp.ArrayHeaderSize
	for za0001 := range z.Roles {
		s += z.Roles[za0001].Msgsize()
	}
	s += 6 + hsp.IntSize + 9 + z.Envelope.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *FindNeighborResp) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 3
	o = append(o, 0x83, 0x83)
	o = hsp.AppendArrayHeader(o, uint32(len(z.Nodes)))
	for za0001 := range z.Nodes {
		if oTemp, err := z.Nodes[za0001].MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x83)
	o = hsp.AppendString(o, z.Msg)
	o = append(o, 0x83)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FindNeighborResp) Msgsize() (s int) {
	s = 1 + 6 + hsp.ArrayHeaderSize
	for za0001 := range z.Nodes {
		s += z.Nodes[za0001].Msgsize()
	}
	s += 4 + hsp.StringPrefixSize + len(z.Msg) + 9 + z.Envelope.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *FindNodeReq) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	if oTemp, err := z.NodeID.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x82)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FindNodeReq) Msgsize() (s int) {
	s = 1 + 7 + z.NodeID.Msgsize() + 9 + z.Envelope.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *FindNodeResp) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 3
	o = append(o, 0x83, 0x83)
	if z.Node == nil {
		o = hsp.AppendNil(o)
	} else {
		if oTemp, err := z.Node.MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x83)
	o = hsp.AppendString(o, z.Msg)
	o = append(o, 0x83)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *FindNodeResp) Msgsize() (s int) {
	s = 1 + 5
	if z.Node == nil {
		s += hsp.NilSize
	} else {
		s += z.Node.Msgsize()
	}
	s += 4 + hsp.StringPrefixSize + len(z.Msg) + 9 + z.Envelope.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *PingReq) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	if oTemp, err := z.Node.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x82)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PingReq) Msgsize() (s int) {
	s = 1 + 5 + z.Node.Msgsize() + 9 + z.Envelope.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *PingResp) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	o = hsp.AppendString(o, z.Msg)
	o = append(o, 0x82)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PingResp) Msgsize() (s int) {
	s = 1 + 4 + hsp.StringPrefixSize + len(z.Msg) + 9 + z.Envelope.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *UploadMetricsReq) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	o = hsp.AppendArrayHeader(o, uint32(len(z.MFBytes)))
	for za0001 := range z.MFBytes {
		o = hsp.AppendBytes(o, z.MFBytes[za0001])
	}
	o = append(o, 0x82)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UploadMetricsReq) Msgsize() (s int) {
	s = 1 + 8 + hsp.ArrayHeaderSize
	for za0001 := range z.MFBytes {
		s += hsp.BytesPrefixSize + len(z.MFBytes[za0001])
	}
	s += 9 + z.Envelope.Msgsize()
	return
}

// MarshalHash marshals for hash
func (z *UploadMetricsResp) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 2
	o = append(o, 0x82, 0x82)
	o = hsp.AppendString(o, z.Msg)
	o = append(o, 0x82)
	if oTemp, err := z.Envelope.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UploadMetricsResp) Msgsize() (s int) {
	s = 1 + 4 + hsp.StringPrefixSize + len(z.Msg) + 9 + z.Envelope.Msgsize()
	return
}
