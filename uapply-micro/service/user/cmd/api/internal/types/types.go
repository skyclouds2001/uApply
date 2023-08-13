// Code generated by goctl. DO NOT EDIT.
package types

type IfExist struct {
	Exist  bool         `json:"exist"`
	Resume AddResumeReq `json:"resume,omitempty"`
}

type GetRegRsp struct {
	Rsp []ODInfo `json:"rsp"`
}

type ODInfo struct {
	DepID   int    `json:"dep_id"`
	DepName string `json:"dep_name"`
	OrgID   int    `json:"org_id"`
	OrgName string `json:"org_name"`
	First   int    `json:"first"`
	Second  int    `json:"second"`
	Final   int    `json:"final"`
}

type RegReq struct {
	OrgId int `json:"org_id"`
	DepId int `json:"dep_id"`
}

type RegRsp struct {
	Msg string `json:"msg"`
}

type GetDepsReq struct {
	Id string `path:"id"`
}

type GetDepsRsp struct {
	Deps []DepInfo `json:"deps"`
}

type DepInfo struct {
	DepID   int    `json:"dep_id"`
	DepName string `json:"dep_name"`
}

type AddResumeReq struct {
	Name       string `json:"name" binding:"required"`
	Sex        string `json:"sex,options=男|女"`
	StudentNum string `json:"student_num" binding:"required"`
	AddressNum string `json:"address_num,optional"`
	Major      string `json:"major,optional"`
	PhoneNum   string `json:"phone_num" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Intro      string `json:"intro" binding:"required"`
}

type AddResumeRsp struct {
	Msg string `json:"msg"`
}

type EnrollingRsp struct {
	Organizations []Enrolling `json:"organizations"`
}

type Enrolling struct {
	OrgId     int    `json:"org_id"`
	OrgName   string `json:"org_name"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type LoginReq struct {
	Code string `path:"code"`
}

type LoginRsp struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
	UID          int    `json:"uid"`
}

type ResumeOkResp struct {
	IsOk bool `json:"is_ok"`
}

type MiniCodeResp struct {
	Code int `json:"code"`
}

type XdLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
