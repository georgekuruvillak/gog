// Code generated by protoc-gen-gogo.
// source: message.proto
// DO NOT EDIT!

/*
	Package protobuf is a generated protocol buffer package.

	It is generated from these files:
		message.proto

	It has these top-level messages:
		UserMessage
		Join
		ForwardJoin
		Disconnect
		Shuffle
		ShuffleReply
*/
package protobuf

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io "io"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"

import fmt "fmt"
import strings "strings"
import reflect "reflect"

import fmt1 "fmt"
import strings1 "strings"
import code_google_com_p_gogoprotobuf_proto1 "code.google.com/p/gogoprotobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect1 "reflect"

import fmt2 "fmt"
import bytes "bytes"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// User defined messages.
type UserMessage struct {
	Uuid             []byte   `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	Payload          [][]byte `protobuf:"bytes,2,rep,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *UserMessage) Reset()      { *m = UserMessage{} }
func (*UserMessage) ProtoMessage() {}

func (m *UserMessage) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *UserMessage) GetPayload() [][]byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// The Join request.
type Join struct {
	Uuid             []byte `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Join) Reset()      { *m = Join{} }
func (*Join) ProtoMessage() {}

func (m *Join) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

// The ForwardJoin request.
type ForwardJoin struct {
	Uuid             []byte  `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	Ttl              *uint32 `protobuf:"varint,2,req,name=ttl" json:"ttl,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ForwardJoin) Reset()      { *m = ForwardJoin{} }
func (*ForwardJoin) ProtoMessage() {}

func (m *ForwardJoin) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *ForwardJoin) GetTtl() uint32 {
	if m != nil && m.Ttl != nil {
		return *m.Ttl
	}
	return 0
}

// The Disconnect request.
type Disconnect struct {
	Uuid             []byte `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Disconnect) Reset()      { *m = Disconnect{} }
func (*Disconnect) ProtoMessage() {}

func (m *Disconnect) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

// The Shuffle request.
type Shuffle struct {
	Uuid             []byte   `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	Candidates       [][]byte `protobuf:"bytes,2,rep,name=candidates" json:"candidates,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Shuffle) Reset()      { *m = Shuffle{} }
func (*Shuffle) ProtoMessage() {}

func (m *Shuffle) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *Shuffle) GetCandidates() [][]byte {
	if m != nil {
		return m.Candidates
	}
	return nil
}

// The ShuffleReply.
type ShuffleReply struct {
	Uuid             []byte   `protobuf:"bytes,1,req,name=uuid" json:"uuid,omitempty"`
	Candidates       [][]byte `protobuf:"bytes,2,rep,name=candidates" json:"candidates,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ShuffleReply) Reset()      { *m = ShuffleReply{} }
func (*ShuffleReply) ProtoMessage() {}

func (m *ShuffleReply) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *ShuffleReply) GetCandidates() [][]byte {
	if m != nil {
		return m.Candidates
	}
	return nil
}

func init() {
}
func (m *UserMessage) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid, data[index:postIndex]...)
			index = postIndex
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload, make([]byte, postIndex-index))
			copy(m.Payload[len(m.Payload)-1], data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Join) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *ForwardJoin) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid, data[index:postIndex]...)
			index = postIndex
		case 2:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ttl = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Disconnect) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Shuffle) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid, data[index:postIndex]...)
			index = postIndex
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Candidates = append(m.Candidates, make([]byte, postIndex-index))
			copy(m.Candidates[len(m.Candidates)-1], data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *ShuffleReply) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = append(m.Uuid, data[index:postIndex]...)
			index = postIndex
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Candidates = append(m.Candidates, make([]byte, postIndex-index))
			copy(m.Candidates[len(m.Candidates)-1], data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *UserMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UserMessage{`,
		`Uuid:` + valueToStringMessage(this.Uuid) + `,`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Join) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Join{`,
		`Uuid:` + valueToStringMessage(this.Uuid) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ForwardJoin) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ForwardJoin{`,
		`Uuid:` + valueToStringMessage(this.Uuid) + `,`,
		`Ttl:` + valueToStringMessage(this.Ttl) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Disconnect) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Disconnect{`,
		`Uuid:` + valueToStringMessage(this.Uuid) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Shuffle) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Shuffle{`,
		`Uuid:` + valueToStringMessage(this.Uuid) + `,`,
		`Candidates:` + fmt.Sprintf("%v", this.Candidates) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ShuffleReply) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ShuffleReply{`,
		`Uuid:` + valueToStringMessage(this.Uuid) + `,`,
		`Candidates:` + fmt.Sprintf("%v", this.Candidates) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *UserMessage) Size() (n int) {
	var l int
	_ = l
	if m.Uuid != nil {
		l = len(m.Uuid)
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Payload) > 0 {
		for _, b := range m.Payload {
			l = len(b)
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Join) Size() (n int) {
	var l int
	_ = l
	if m.Uuid != nil {
		l = len(m.Uuid)
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *ForwardJoin) Size() (n int) {
	var l int
	_ = l
	if m.Uuid != nil {
		l = len(m.Uuid)
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Ttl != nil {
		n += 1 + sovMessage(uint64(*m.Ttl))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Disconnect) Size() (n int) {
	var l int
	_ = l
	if m.Uuid != nil {
		l = len(m.Uuid)
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Shuffle) Size() (n int) {
	var l int
	_ = l
	if m.Uuid != nil {
		l = len(m.Uuid)
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Candidates) > 0 {
		for _, b := range m.Candidates {
			l = len(b)
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *ShuffleReply) Size() (n int) {
	var l int
	_ = l
	if m.Uuid != nil {
		l = len(m.Uuid)
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Candidates) > 0 {
		for _, b := range m.Candidates {
			l = len(b)
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedUserMessage(r randyMessage, easy bool) *UserMessage {
	this := &UserMessage{}
	v1 := r.Intn(100)
	this.Uuid = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Uuid[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(100)
		this.Payload = make([][]byte, v2)
		for i := 0; i < v2; i++ {
			v3 := r.Intn(100)
			this.Payload[i] = make([]byte, v3)
			for j := 0; j < v3; j++ {
				this.Payload[i][j] = byte(r.Intn(256))
			}
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessage(r, 3)
	}
	return this
}

func NewPopulatedJoin(r randyMessage, easy bool) *Join {
	this := &Join{}
	v4 := r.Intn(100)
	this.Uuid = make([]byte, v4)
	for i := 0; i < v4; i++ {
		this.Uuid[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessage(r, 2)
	}
	return this
}

func NewPopulatedForwardJoin(r randyMessage, easy bool) *ForwardJoin {
	this := &ForwardJoin{}
	v5 := r.Intn(100)
	this.Uuid = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.Uuid[i] = byte(r.Intn(256))
	}
	v6 := r.Uint32()
	this.Ttl = &v6
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessage(r, 3)
	}
	return this
}

func NewPopulatedDisconnect(r randyMessage, easy bool) *Disconnect {
	this := &Disconnect{}
	v7 := r.Intn(100)
	this.Uuid = make([]byte, v7)
	for i := 0; i < v7; i++ {
		this.Uuid[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessage(r, 2)
	}
	return this
}

func NewPopulatedShuffle(r randyMessage, easy bool) *Shuffle {
	this := &Shuffle{}
	v8 := r.Intn(100)
	this.Uuid = make([]byte, v8)
	for i := 0; i < v8; i++ {
		this.Uuid[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		v9 := r.Intn(100)
		this.Candidates = make([][]byte, v9)
		for i := 0; i < v9; i++ {
			v10 := r.Intn(100)
			this.Candidates[i] = make([]byte, v10)
			for j := 0; j < v10; j++ {
				this.Candidates[i][j] = byte(r.Intn(256))
			}
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessage(r, 3)
	}
	return this
}

func NewPopulatedShuffleReply(r randyMessage, easy bool) *ShuffleReply {
	this := &ShuffleReply{}
	v11 := r.Intn(100)
	this.Uuid = make([]byte, v11)
	for i := 0; i < v11; i++ {
		this.Uuid[i] = byte(r.Intn(256))
	}
	if r.Intn(10) != 0 {
		v12 := r.Intn(100)
		this.Candidates = make([][]byte, v12)
		for i := 0; i < v12; i++ {
			v13 := r.Intn(100)
			this.Candidates[i] = make([]byte, v13)
			for j := 0; j < v13; j++ {
				this.Candidates[i][j] = byte(r.Intn(256))
			}
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMessage(r, 3)
	}
	return this
}

type randyMessage interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMessage(r randyMessage) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringMessage(r randyMessage) string {
	v14 := r.Intn(100)
	tmps := make([]rune, v14)
	for i := 0; i < v14; i++ {
		tmps[i] = randUTF8RuneMessage(r)
	}
	return string(tmps)
}
func randUnrecognizedMessage(r randyMessage, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldMessage(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldMessage(data []byte, r randyMessage, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateMessage(data, uint64(key))
		v15 := r.Int63()
		if r.Intn(2) == 0 {
			v15 *= -1
		}
		data = encodeVarintPopulateMessage(data, uint64(v15))
	case 1:
		data = encodeVarintPopulateMessage(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateMessage(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateMessage(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateMessage(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateMessage(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *UserMessage) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *UserMessage) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uuid != nil {
		data[i] = 0xa
		i++
		i = encodeVarintMessage(data, i, uint64(len(m.Uuid)))
		i += copy(data[i:], m.Uuid)
	}
	if len(m.Payload) > 0 {
		for _, b := range m.Payload {
			data[i] = 0x12
			i++
			i = encodeVarintMessage(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Join) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Join) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uuid != nil {
		data[i] = 0xa
		i++
		i = encodeVarintMessage(data, i, uint64(len(m.Uuid)))
		i += copy(data[i:], m.Uuid)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *ForwardJoin) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ForwardJoin) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uuid != nil {
		data[i] = 0xa
		i++
		i = encodeVarintMessage(data, i, uint64(len(m.Uuid)))
		i += copy(data[i:], m.Uuid)
	}
	if m.Ttl != nil {
		data[i] = 0x10
		i++
		i = encodeVarintMessage(data, i, uint64(*m.Ttl))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Disconnect) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Disconnect) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uuid != nil {
		data[i] = 0xa
		i++
		i = encodeVarintMessage(data, i, uint64(len(m.Uuid)))
		i += copy(data[i:], m.Uuid)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Shuffle) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Shuffle) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uuid != nil {
		data[i] = 0xa
		i++
		i = encodeVarintMessage(data, i, uint64(len(m.Uuid)))
		i += copy(data[i:], m.Uuid)
	}
	if len(m.Candidates) > 0 {
		for _, b := range m.Candidates {
			data[i] = 0x12
			i++
			i = encodeVarintMessage(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *ShuffleReply) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ShuffleReply) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Uuid != nil {
		data[i] = 0xa
		i++
		i = encodeVarintMessage(data, i, uint64(len(m.Uuid)))
		i += copy(data[i:], m.Uuid)
	}
	if len(m.Candidates) > 0 {
		for _, b := range m.Candidates {
			data[i] = 0x12
			i++
			i = encodeVarintMessage(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64Message(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Message(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMessage(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *UserMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&protobuf.UserMessage{` + `Uuid:` + valueToGoStringMessage(this.Uuid, "byte"), `Payload:` + fmt1.Sprintf("%#v", this.Payload), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Join) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&protobuf.Join{` + `Uuid:` + valueToGoStringMessage(this.Uuid, "byte"), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *ForwardJoin) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&protobuf.ForwardJoin{` + `Uuid:` + valueToGoStringMessage(this.Uuid, "byte"), `Ttl:` + valueToGoStringMessage(this.Ttl, "uint32"), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Disconnect) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&protobuf.Disconnect{` + `Uuid:` + valueToGoStringMessage(this.Uuid, "byte"), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Shuffle) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&protobuf.Shuffle{` + `Uuid:` + valueToGoStringMessage(this.Uuid, "byte"), `Candidates:` + fmt1.Sprintf("%#v", this.Candidates), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *ShuffleReply) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings1.Join([]string{`&protobuf.ShuffleReply{` + `Uuid:` + valueToGoStringMessage(this.Uuid, "byte"), `Candidates:` + fmt1.Sprintf("%#v", this.Candidates), `XXX_unrecognized:` + fmt1.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect1.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect1.Indirect(rv).Interface()
	return fmt1.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringMessage(e map[int32]code_google_com_p_gogoprotobuf_proto1.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings1.Join(ss, ",") + "}"
	return s
}
func (this *UserMessage) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*UserMessage)
	if !ok {
		return fmt2.Errorf("that is not of type *UserMessage")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *UserMessage but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *UserMessagebut is not nil && this == nil")
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return fmt2.Errorf("Uuid this(%v) Not Equal that(%v)", this.Uuid, that1.Uuid)
	}
	if len(this.Payload) != len(that1.Payload) {
		return fmt2.Errorf("Payload this(%v) Not Equal that(%v)", len(this.Payload), len(that1.Payload))
	}
	for i := range this.Payload {
		if !bytes.Equal(this.Payload[i], that1.Payload[i]) {
			return fmt2.Errorf("Payload this[%v](%v) Not Equal that[%v](%v)", i, this.Payload[i], i, that1.Payload[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *UserMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*UserMessage)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return false
	}
	if len(this.Payload) != len(that1.Payload) {
		return false
	}
	for i := range this.Payload {
		if !bytes.Equal(this.Payload[i], that1.Payload[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Join) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Join)
	if !ok {
		return fmt2.Errorf("that is not of type *Join")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *Join but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *Joinbut is not nil && this == nil")
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return fmt2.Errorf("Uuid this(%v) Not Equal that(%v)", this.Uuid, that1.Uuid)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Join) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Join)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ForwardJoin) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ForwardJoin)
	if !ok {
		return fmt2.Errorf("that is not of type *ForwardJoin")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *ForwardJoin but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *ForwardJoinbut is not nil && this == nil")
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return fmt2.Errorf("Uuid this(%v) Not Equal that(%v)", this.Uuid, that1.Uuid)
	}
	if this.Ttl != nil && that1.Ttl != nil {
		if *this.Ttl != *that1.Ttl {
			return fmt2.Errorf("Ttl this(%v) Not Equal that(%v)", *this.Ttl, *that1.Ttl)
		}
	} else if this.Ttl != nil {
		return fmt2.Errorf("this.Ttl == nil && that.Ttl != nil")
	} else if that1.Ttl != nil {
		return fmt2.Errorf("Ttl this(%v) Not Equal that(%v)", this.Ttl, that1.Ttl)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *ForwardJoin) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ForwardJoin)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return false
	}
	if this.Ttl != nil && that1.Ttl != nil {
		if *this.Ttl != *that1.Ttl {
			return false
		}
	} else if this.Ttl != nil {
		return false
	} else if that1.Ttl != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Disconnect) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Disconnect)
	if !ok {
		return fmt2.Errorf("that is not of type *Disconnect")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *Disconnect but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *Disconnectbut is not nil && this == nil")
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return fmt2.Errorf("Uuid this(%v) Not Equal that(%v)", this.Uuid, that1.Uuid)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Disconnect) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Disconnect)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Shuffle) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Shuffle)
	if !ok {
		return fmt2.Errorf("that is not of type *Shuffle")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *Shuffle but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *Shufflebut is not nil && this == nil")
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return fmt2.Errorf("Uuid this(%v) Not Equal that(%v)", this.Uuid, that1.Uuid)
	}
	if len(this.Candidates) != len(that1.Candidates) {
		return fmt2.Errorf("Candidates this(%v) Not Equal that(%v)", len(this.Candidates), len(that1.Candidates))
	}
	for i := range this.Candidates {
		if !bytes.Equal(this.Candidates[i], that1.Candidates[i]) {
			return fmt2.Errorf("Candidates this[%v](%v) Not Equal that[%v](%v)", i, this.Candidates[i], i, that1.Candidates[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *Shuffle) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Shuffle)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return false
	}
	if len(this.Candidates) != len(that1.Candidates) {
		return false
	}
	for i := range this.Candidates {
		if !bytes.Equal(this.Candidates[i], that1.Candidates[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ShuffleReply) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ShuffleReply)
	if !ok {
		return fmt2.Errorf("that is not of type *ShuffleReply")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt2.Errorf("that is type *ShuffleReply but is nil && this != nil")
	} else if this == nil {
		return fmt2.Errorf("that is type *ShuffleReplybut is not nil && this == nil")
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return fmt2.Errorf("Uuid this(%v) Not Equal that(%v)", this.Uuid, that1.Uuid)
	}
	if len(this.Candidates) != len(that1.Candidates) {
		return fmt2.Errorf("Candidates this(%v) Not Equal that(%v)", len(this.Candidates), len(that1.Candidates))
	}
	for i := range this.Candidates {
		if !bytes.Equal(this.Candidates[i], that1.Candidates[i]) {
			return fmt2.Errorf("Candidates this[%v](%v) Not Equal that[%v](%v)", i, this.Candidates[i], i, that1.Candidates[i])
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt2.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *ShuffleReply) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ShuffleReply)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Uuid, that1.Uuid) {
		return false
	}
	if len(this.Candidates) != len(that1.Candidates) {
		return false
	}
	for i := range this.Candidates {
		if !bytes.Equal(this.Candidates[i], that1.Candidates[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}