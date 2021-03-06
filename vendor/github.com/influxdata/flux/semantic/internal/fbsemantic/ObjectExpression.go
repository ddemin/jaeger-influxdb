// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ObjectExpression struct {
	_tab flatbuffers.Table
}

func GetRootAsObjectExpression(buf []byte, offset flatbuffers.UOffsetT) *ObjectExpression {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ObjectExpression{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ObjectExpression) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ObjectExpression) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ObjectExpression) Loc(obj *SourceLocation) *SourceLocation {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(SourceLocation)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ObjectExpression) With(obj *IdentifierExpression) *IdentifierExpression {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(IdentifierExpression)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ObjectExpression) Properties(obj *Property, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ObjectExpression) PropertiesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ObjectExpression) TypType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ObjectExpression) MutateTypType(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func (rcv *ObjectExpression) Typ(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func ObjectExpressionStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ObjectExpressionAddLoc(builder *flatbuffers.Builder, loc flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(loc), 0)
}
func ObjectExpressionAddWith(builder *flatbuffers.Builder, with flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(with), 0)
}
func ObjectExpressionAddProperties(builder *flatbuffers.Builder, properties flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(properties), 0)
}
func ObjectExpressionStartPropertiesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ObjectExpressionAddTypType(builder *flatbuffers.Builder, typType byte) {
	builder.PrependByteSlot(3, typType, 0)
}
func ObjectExpressionAddTyp(builder *flatbuffers.Builder, typ flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(typ), 0)
}
func ObjectExpressionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
