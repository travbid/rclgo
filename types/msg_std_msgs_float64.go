package types

import (
	"unsafe"

	cwrap "github.com/richardrigby/rclgo/internal"
)

type StdMsgsFloat64 struct {
	data    *cwrap.StdMsgs_MsgFloat64
	MsgType MessageTypeSupport
}

func (msg *StdMsgsFloat64) GetMessage() MessageTypeSupport {
	return msg.MsgType
}

func (msg *StdMsgsFloat64) GetData() MessageData {
	return MessageData{unsafe.Pointer(msg.data)}
}

func (msg *StdMsgsFloat64) InitMessage() {
	msg.data = cwrap.InitStdMsgsMsgFloat64()
	msg.MsgType = GetMessageTypeFromStdMsgsFloat64()
}

func (msg *StdMsgsFloat64) SetFloat64(data float64) {
	//TODO: to implement the setter
	msg.data.Set(data)
}

func (msg *StdMsgsFloat64) GetFloat64() float64 {
	return float64(msg.data.Data())
}

func (msg *StdMsgsFloat64) GetDataAsString() string {
	//TODO: to implement the stringify opt
	myRetValue := ""
	return myRetValue
}

func (msg *StdMsgsFloat64) DestroyMessage() {
	cwrap.DestroyStdMsgsMsgFloat64(msg.data)
}

func GetMessageTypeFromStdMsgsFloat64() MessageTypeSupport {
	ret := cwrap.GetMessageTypeFromStdMsgsMsgFloat64()
	return MessageTypeSupport{ret}
}
