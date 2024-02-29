package convert

import (
	"github.com/woxQAQ/im-service/internal/rpc/imrpc_user/pb"
	"github.com/woxQAQ/im-service/pkg/common/sql/user"
)

func StrToGender(gender string) (pb.Gender, error) {
	switch gender {
	case "unknown":
		return pb.Gender_GENDER_UNKNOWN, nil
	case "male":
		return pb.Gender_GENDER_MALE, nil
	case "female":
		return pb.Gender_GENDER_FEMALE, nil
	}
	return -1, user.ErrGenderInvalid
}
