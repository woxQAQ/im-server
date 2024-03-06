package errs

import "errors"

var (
	ErrSenderIdNotMatch = errors.New("SenderId not match the userId it comes from")
	ErrGroupIdNotFound  = errors.New("GroupId is not in the request")
	ErrRecvIdNotFound   = errors.New("ReceiverId is not in the request")

	ErrArgumentErr = errors.New("ArgumentError")
)
