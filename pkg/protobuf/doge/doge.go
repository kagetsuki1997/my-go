package doge

import (
	"my-go/internal/doge/model"
	common "my-go/pkg/protobuf/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromDogeTypeProto(dogeTypeProto *DogeType) (*model.DogeType, error) {
	var dogeType model.DogeType
	switch *dogeTypeProto {
	case DogeType_DOGE_TYPE_UNSPECIFIED:
		return nil, status.Error(codes.InvalidArgument, "invalid doge type")
	case DogeType_DOGE_TYPE_DOGE:
		dogeType = model.DogeTypeDoge
	case DogeType_DOGE_TYPE_DOGE_CRY:
		dogeType = model.DogeTypeDogeCry
	case DogeType_DOGE_TYPE_DOGE_BUFFED:
		dogeType = model.DogeTypeDogeBuffed
	case DogeType_DOGE_TYPE_DOGE_KACHITORITAI:
		dogeType = model.DogeTypeDogeKachitoritai
	}

	return &dogeType, nil
}

func ToDogeTypeProto(dogeType *model.DogeType) (dogeTypeProto DogeType) {
	switch *dogeType {
	case model.DogeTypeDoge:
		dogeTypeProto = DogeType_DOGE_TYPE_DOGE
	case model.DogeTypeDogeCry:
		dogeTypeProto = DogeType_DOGE_TYPE_DOGE_CRY
	case model.DogeTypeDogeBuffed:
		dogeTypeProto = DogeType_DOGE_TYPE_DOGE_BUFFED
	case model.DogeTypeDogeKachitoritai:
		dogeTypeProto = DogeType_DOGE_TYPE_DOGE_KACHITORITAI
	}

	return
}

func ToDogeProto(doge *model.Doge) *Doge {
	dogeProto := Doge{
		Id:              common.ToUUIDProto(doge.ID),
		Name:            doge.Name,
		MagicNumber:     int64(doge.MagicNumber),
		Type:            ToDogeTypeProto(&doge.Type),
		CreateTimestamp: timestamppb.New(doge.CreateTimestamp),
	}

	return &dogeProto
}
