package s3type

import "fmt"

type ParseEventNameError struct {
	message string
}

func (p *ParseEventNameError) String() string {
	return p.message
}

func (p ParseEventNameError) Error() string {
	return p.message
}

type EventName int

const (
	ObjectAccessedGet = iota
	ObjectAccessedGetRetention
	ObjectAccessedGetLegalHold
	ObjectAccessedHead
	ObjectAccessedAttributes
	ObjectCreatedCompleteMultipartUpload
	ObjectCreatedCopy
	ObjectCreatedPost
	ObjectCreatedPut
	ObjectCreatedPutRetention
	ObjectCreatedPutLegalHold
	ObjectTaggingPut
	ObjectTaggingDelete
	ObjectRemovedDelete
	ObjectRemovedDeleteMarkerCreated
	ObjectRemovedDeleteAllVersions
	ObjectRemovedNoOP
	BucketCreated
	BucketRemoved
	ObjectReplicationFailed
	ObjectReplicationComplete
	ObjectReplicationMissedThreshold
	ObjectReplicationReplicatedAfterThreshold
	ObjectReplicationNotTracked
	ObjectRestorePost
	ObjectRestoreCompleted
	ObjectTransitionFailed
	ObjectTransitionComplete
	ScannerManyVersions
	ScannerLargeVersions
	ScannerBigPrefix
	LifecycleDelMarkerExpirationDelete
	ObjectAclPut
	LifecycleExpirationDelete
	LifecycleExpirationDeleteMarkerCreated
	LifecycleTransition
	IntelligentTiering

	ObjectAccessedAll
	ObjectCreatedAll
	ObjectRemovedAll
	ObjectReplicationAll
	ObjectRestoreAll
	ObjectTaggingAll
	LifecycleExpirationAll
	ObjectTransitionAll
	ObjectScannerAll

	Everything
	ObjectRemovedAbortMultipartUpload
	ObjectCreatedCreateMultipartUpload
	ObjectRemovedDeleteObjects
)

func ParseEventName(name string) (EventName, error) {
	switch name {
	case "s3:BucketCreated:*":
		return BucketCreated, nil
	case "s3:BucketRemoved:*":
		return BucketRemoved, nil
	case "s3:ObjectAccessed:*":
		return ObjectAccessedAll, nil
	case "s3:ObjectAccessed:Get":
		return ObjectAccessedGet, nil
	case "s3:ObjectAccessed:GetRetention":
		return ObjectAccessedGetRetention, nil
	case "s3:ObjectAccessed:GetLegalHold":
		return ObjectAccessedGetLegalHold, nil
	case "s3:ObjectAccessed:Head":
		return ObjectAccessedHead, nil
	case "s3:ObjectAccessed:Attributes":
		return ObjectAccessedAttributes, nil
	case "s3:ObjectCreated:*":
		return ObjectCreatedAll, nil
	case "s3:ObjectCreated:CompleteMultipartUpload":
		return ObjectCreatedCompleteMultipartUpload, nil
	case "s3:ObjectCreated:Copy":
		return ObjectCreatedCopy, nil
	case "s3:ObjectCreated:Post":
		return ObjectCreatedPost, nil
	case "s3:ObjectCreated:Put":
		return ObjectCreatedPut, nil
	case "s3:ObjectCreated:PutRetention":
		return ObjectCreatedPutRetention, nil
	case "s3:ObjectCreated:PutLegalHold":
		return ObjectCreatedPutLegalHold, nil
	case "s3:ObjectCreated:PutTagging":
		return ObjectTaggingPut, nil
	case "s3:ObjectCreated:DeleteTagging":
		return ObjectTaggingDelete, nil

		// TODO: complete the events
	}

	return -1, ParseEventNameError{message: fmt.Sprintf("Invalid event name:%s", name)}
}
