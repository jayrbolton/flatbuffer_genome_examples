// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Sample

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Potion struct {
	_tab flatbuffers.Table
}

func GetRootAsPotion(buf []byte, offset flatbuffers.UOffsetT) *Potion {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Potion{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Potion) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Potion) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Potion) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Potion) HealingAmount() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Potion) MutateHealingAmount(n int16) bool {
	return rcv._tab.MutateInt16Slot(6, n)
}

func PotionStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PotionAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func PotionAddHealingAmount(builder *flatbuffers.Builder, healingAmount int16) {
	builder.PrependInt16Slot(1, healingAmount, 0)
}
func PotionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
