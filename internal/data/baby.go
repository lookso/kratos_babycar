package data

import (
	babyBiz "babycare/internal/biz/baby"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BabyData struct {
	data   *Data
	logger *log.Helper
}

func NewBabyData(data *Data, logger log.Logger) babyBiz.IBabyRepo {
	return &BabyData{data: data, logger: log.NewHelper(log.With(logger, "module", "data/baby"))}
}

// 对user.IUser接口的实现
func (u *BabyData) GeUserId(ctx context.Context, id int64) (int64, error) {
	if id<0{
		return 0,nil
	}else{
		return 100,nil
	}
	return id, nil
}
