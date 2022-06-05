package helper

import (
	v1 "github.com/AkiOuma/demo01/api/goapi/demo01/service/v1"
	"github.com/AkiOuma/demo01/internal/domain/entity"
	"github.com/AkiOuma/demo01/internal/domain/valueobject"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func V1User2EntityUser(source ...*v1.User) []*entity.User {
	ans := make([]*entity.User, 0, len(source))
	for _, v := range source {
		ans = append(ans, entity.NewUser(int(v.GetId()), &valueobject.User{
			Name:      v.GetName(),
			Height:    v.GetHeight(),
			Group:     int(v.GetGroup()),
			CreatedAt: v.GetCreatedAt().AsTime(),
			UpdatedAt: v.GetUpdatedAt().AsTime(),
		}))
	}
	return ans
}

func EntityUser2V1User(source ...*entity.User) []*v1.User {
	ans := make([]*v1.User, 0, len(source))
	for _, v := range source {
		ans = append(ans, &v1.User{
			Id:        int32(v.GetId()),
			Name:      v.GetName(),
			Height:    v.GetHeight(),
			Group:     int32(v.GetGroup()),
			CreatedAt: timestamppb.New(v.GetCreatedAt()),
			UpdatedAt: timestamppb.New(v.GetUpdatedAt()),
		})
	}
	return ans
}

func V1UserQuerier2EntityUserQuerier(source *v1.UserQuerier) *valueobject.UserQuerier {
	return &valueobject.UserQuerier{
		Id:             bulkInt32ToInt(source.GetId()),
		IdLower:        int(source.GetIdLower()),
		IdUpper:        int(source.GetIdUpper()),
		Name:           source.GetName(),
		SearchName:     source.GetSearchName(),
		HeightLower:    source.GetHeightLower(),
		HeightUpper:    source.GetHeightUpper(),
		Group:          bulkInt32ToInt(source.GetGroup()),
		GroupLower:     int(source.GetGroupLower()),
		GroupUpper:     int(source.GetGroupUpper()),
		CreatedAtLower: source.GetCreatedAtLower().AsTime(),
		CreatedAtUpper: source.GetCreatedAtUpper().AsTime(),
		UpdatedAtLower: source.GetUpdatedAtLower().AsTime(),
		UpdatedAtUpper: source.GetUpdatedAtUpper().AsTime(),
	}
}

func bulkInt32ToInt(source []int32) []int {
	ans := make([]int, 0, len(source))
	for _, v := range source {
		ans = append(ans, int(v))
	}
	return ans
}
