package audit

import (
	"strox/src/module/s3type"
	"time"
)

type ObjectVersion struct {
	ObjectName string
	VersionId  string
}

type ApiDetails struct {
	Name                string
	Bucket              string
	Object              string
	Objects             []ObjectVersion
	Status              string
	StatusCode          int32
	InputBytes          int64
	HeaderBytes         int64
	TimeToFirstByte     string
	TimeToFirstByteInNs string
	TimeToResponse      string
	TimeToResponseInNs  string
}

type AuditEntry struct {
	Version       string
	DeploymentId  string
	SiteName      string
	Time          time.Duration
	Event         s3type.EventName
	EntryType     string
	Trigger       string
	Api           ApiDetails
	RemoteHost    string
	RequestId     string
	UserAgent     string
	RequestPath   string
	RequestHost   string
	RequestNode   string
	RequestClaims map[string]any
	RequestQuery  map[string]string
	RequestHeader map[string]string
	Tags          map[string]any
	AccessKey     string
	ParentUser    string
	Error         string
}
