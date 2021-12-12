package rstarhelper

type ErrorRsp struct {
	ErrMsg string `json:"ErrMsg"`
	Ret    int64  `json:"Ret"`
}

type InitWxidsReq struct {
	CurrentChatRoomContactSeq int64 `json:"CurrentChatRoomContactSeq"`
	CurrentWxcontactSeq       int64 `json:"CurrentWxcontactSeq"`
}

type InitWxidsRsp struct {
	ContinueFlag              int64    `json:"ContinueFlag"`
	CurrentChatRoomContactSeq int64    `json:"CurrentChatRoomContactSeq"`
	CurrentWxcontactSeq       int64    `json:"CurrentWxcontactSeq"`
	Wxid                      []string `json:"Wxid"`
}
