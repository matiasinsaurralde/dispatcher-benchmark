package dispatcher

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"C"

	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Dispatcher) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Mode":
			z.Mode, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Dispatcher) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "Mode"
	err = en.Append(0x81, 0xa4, 0x4d, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Mode)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Dispatcher) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Mode"
	o = append(o, 0x81, 0xa4, 0x4d, 0x6f, 0x64, 0x65)
	o = msgp.AppendInt(o, z.Mode)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Dispatcher) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Mode":
			z.Mode, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Dispatcher) Msgsize() (s int) {
	s = 1 + 5 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NestedObject) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "string_field":
			z.NestedStringField, err = dc.ReadString()
			if err != nil {
				return
			}
		case "int_field":
			z.NestedIntField, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z NestedObject) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "string_field"
	err = en.Append(0x82, 0xac, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.NestedStringField)
	if err != nil {
		return
	}
	// write "int_field"
	err = en.Append(0xa9, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.NestedIntField)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NestedObject) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "string_field"
	o = append(o, 0x82, 0xac, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64)
	o = msgp.AppendString(o, z.NestedStringField)
	// string "int_field"
	o = append(o, 0xa9, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64)
	o = msgp.AppendInt(o, z.NestedIntField)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NestedObject) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcmr uint32
	zcmr, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcmr > 0 {
		zcmr--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "string_field":
			z.NestedStringField, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "int_field":
			z.NestedIntField, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NestedObject) Msgsize() (s int) {
	s = 1 + 13 + msgp.StringPrefixSize + len(z.NestedStringField) + 10 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Object) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "message":
			z.Message, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ts":
			z.Timestamp, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "nested_object":
			var zwht uint32
			zwht, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zwht > 0 {
				zwht--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "string_field":
					z.NestedObject.NestedStringField, err = dc.ReadString()
					if err != nil {
						return
					}
				case "int_field":
					z.NestedObject.NestedIntField, err = dc.ReadInt()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Object) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "name"
	err = en.Append(0x84, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "message"
	err = en.Append(0xa7, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Message)
	if err != nil {
		return
	}
	// write "ts"
	err = en.Append(0xa2, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Timestamp)
	if err != nil {
		return
	}
	// write "nested_object"
	// map header, size 2
	// write "string_field"
	err = en.Append(0xad, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x82, 0xac, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.NestedObject.NestedStringField)
	if err != nil {
		return
	}
	// write "int_field"
	err = en.Append(0xa9, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.NestedObject.NestedIntField)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Object) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "name"
	o = append(o, 0x84, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "message"
	o = append(o, 0xa7, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Message)
	// string "ts"
	o = append(o, 0xa2, 0x74, 0x73)
	o = msgp.AppendInt(o, z.Timestamp)
	// string "nested_object"
	// map header, size 2
	// string "string_field"
	o = append(o, 0xad, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x82, 0xac, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64)
	o = msgp.AppendString(o, z.NestedObject.NestedStringField)
	// string "int_field"
	o = append(o, 0xa9, 0x69, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64)
	o = msgp.AppendInt(o, z.NestedObject.NestedIntField)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Object) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "message":
			z.Message, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ts":
			z.Timestamp, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "nested_object":
			var zcua uint32
			zcua, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zcua > 0 {
				zcua--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "string_field":
					z.NestedObject.NestedStringField, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "int_field":
					z.NestedObject.NestedIntField, bts, err = msgp.ReadIntBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Object) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.StringPrefixSize + len(z.Message) + 3 + msgp.IntSize + 14 + 1 + 13 + msgp.StringPrefixSize + len(z.NestedObject.NestedStringField) + 10 + msgp.IntSize
	return
}
