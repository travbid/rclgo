package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsUInt16 struct {
	data    *cwrap.StdMsgs_MsgUInt16
	MsgType MessageTypeSupport
}

func (msg *StdMsgsUInt16) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsUInt16) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsUInt16) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgUInt16()
	msg.MsgType = GetMessageTypeFromStdMsgsUInt16()
}

func (msg *StdMsgsUInt16) SetUInt16(data uint16) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsUInt16) GetUInt16() uint16 {
	return uint16(msg.data.Data())
}

func (msg *StdMsgsUInt16) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsUInt16) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgUInt16(msg.data)
}

func GetMessageTypeFromStdMsgsUInt16() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgUInt16()
	return MessageTypeSupport{ret}
}
