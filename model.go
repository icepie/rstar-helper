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

type SearchContactReq struct {
	Scene    int64  `json:"Scene"`
	UserName string `json:"UserName"`
}

type SearchContactRsp struct {
	Alias           string `json:"Alias"`
	BigHeadImgURL   string `json:"BigHeadImgUrl"`
	ChatRoomOwner   string `json:"ChatRoomOwner"`
	ChatroomVersion int64  `json:"ChatroomVersion"`
	City            string `json:"City"`
	ContactType     int64  `json:"ContactType"`
	EncryptUsername string `json:"EncryptUsername"`
	ExtInfo         string `json:"ExtInfo"`
	ExtInfoExt      string `json:"ExtInfoExt"`
	LabelLists      string `json:"LabelLists"`
	MsgType         int64  `json:"MsgType"`
	NickName        string `json:"NickName"`
	Province        string `json:"Province"`
	Remark          string `json:"Remark"`
	Sex             int64  `json:"Sex"`
	Signature       string `json:"Signature"`
	SmallHeadImgURL string `json:"SmallHeadImgUrl"`
	Ticket          string `json:"Ticket"`
	UserName        string `json:"UserName"`
	VerifyFlag      int64  `json:"VerifyFlag"`
}
