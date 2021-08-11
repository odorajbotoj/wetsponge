package wetsponge

import (
	"encoding/json"
	"fmt"
)

const (
	SUBSCRIBE      uint8 = 0
	UNSUBSCRIBE    uint8 = 1
	COMMANDREQUEST uint8 = 2
)

type CmdReqBodyOrigin struct {
	Type string `json:"type,omitempty"`
}
type CmdReqBody struct {
	EventName   string            `json:"eventName,omitempty"`
	Origin      *CmdReqBodyOrigin `json:"origin,omitempty"`
	CommandLine string            `json:"commandLine,omitempty"`
	Version     int8              `json:"version,omitempty"`
}
type CmdReqHeader struct {
	RequestId      string `json:"requestId"`
	MessagePurpose string `json:"messagePurpose"`
	Version        int8   `json:"version"`
	MessageType    string `json:"messageType"`
}
type CmdReq struct {
	Body   CmdReqBody   `json:"body"`
	Header CmdReqHeader `json:"header"`
}

func MakeCmdReq(mode uint8, reqId UUID, arg string) ([]byte, error) {
	var err error = nil
	var body = CmdReqBody{}
	var header = CmdReqHeader{RequestId: reqId.String(), Version: 1, MessageType: "commandRequest"}
	switch mode {
	case 0:
		body.EventName = arg
		header.MessagePurpose = "subscribe"
	case 1:
		body.EventName = arg
		header.MessagePurpose = "unsubscribe"
	case 2:
		var origin = CmdReqBodyOrigin{Type: "player"}
		body.Origin = &origin
		body.CommandLine = arg
		body.Version = 1
		header.MessagePurpose = "commandRequest"
	default:
		err = fmt.Errorf("WSC_Err:MakeCmdReq:Unknown mode %d .", mode)
		return nil, err
	}
	var request = CmdReq{Body: body, Header: header}
	b_out, m_err := json.Marshal(request)
	if m_err != nil {
		return nil, m_err
	}
	return b_out, err
}
