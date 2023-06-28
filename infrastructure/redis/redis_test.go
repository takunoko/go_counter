package myRedis_test

import (
	"context"
	"github.com/redis/go-redis/v9"
	myRedis "go_web_counter/infrastructure/redis"
	"go_web_counter/pkg/test"
	"testing"
)

func Test_cntRepo_Set(t *testing.T) {
	rCli := test.OpenTestRedis(t)
	rKey := "k"

	type fields struct {
		RCli *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
		val int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: rKey,
				val: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := myRedis.NewDataRepository(tt.fields.RCli)
			if err := dr.Set(tt.args.ctx, tt.args.key, tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
