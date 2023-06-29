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
	rKey := "set"

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

func Test_cntRepo_CntUp(t *testing.T) {
	rCli := test.OpenTestRedis(t)
	key0 := "cnt_up_0"
	key1 := "cnt_up_1"
	keyMinus1 := "cnt_up_minus1"
	notExistKey := "cnt_up_not_ext"

	if err := rCli.MSet(context.Background(),
		map[string]interface{}{
			key0:      "0",
			key1:      "1",
			keyMinus1: "-1",
		},
	).Err(); err != nil {
		t.Error("初期データ作成エラー")
	}

	type fields struct {
		RCli *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "存在しないキーに対するカウントアップ",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: notExistKey,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "自然数に対するカウントアップ",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: key1,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "0に対するカウントアップ",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: key0,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "負の数に対するカウントアップ",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: keyMinus1,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := myRedis.NewDataRepository(tt.fields.RCli)
			got, err := dr.CntUp(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("CntUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CntUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cntRepo_CntDown(t *testing.T) {
	rCli := test.OpenTestRedis(t)
	key0 := "cnt_down_0"
	key1 := "cnt_down_1"
	keyMinus1 := "cnt_down_minus1"
	notExistKey := "cnt_down_not_ext"

	if err := rCli.MSet(context.Background(),
		map[string]interface{}{
			key0:      "0",
			key1:      "1",
			keyMinus1: "-1",
		},
	).Err(); err != nil {
		t.Error("初期データ作成エラー")
	}

	type fields struct {
		RCli *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "存在しないキーに対するカウントダウン",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: notExistKey,
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "自然数に対するカウントダウン",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: key1,
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "0に対するカウントダウン",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: key0,
			},
			want:    -1,
			wantErr: false,
		},
		{
			name: "負の数に対するカウントダウン",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: keyMinus1,
			},
			want:    -2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := myRedis.NewDataRepository(tt.fields.RCli)
			got, err := dr.CntDown(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("CntDown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CntDown() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cntRepo_Get(t *testing.T) {
	rCli := test.OpenTestRedis()
	key1 := "get_1"
	keyStr := "get_str"
	notExistKey := "get_not_ext"

	if err := rCli.MSet(context.Background(),
		map[string]interface{}{
			key1:   "1",
			keyStr: "aaa", // 文字列をセット
		},
	).Err(); err != nil {
		t.Error("初期データ作成エラー")
	}

	type fields struct {
		RCli *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "存在するキーに対する取得",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: key1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "存在しないキーに対する取得",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: notExistKey,
			},
			wantErr: true,
		},
		{
			name: "文字列に対する数値の取得",
			fields: fields{
				RCli: rCli,
			},
			args: args{
				ctx: context.Background(),
				key: keyStr,
			},
			wantErr: true, // エラーとなる
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dr := myRedis.NewDataRepository(tt.fields.RCli)
			got, err := dr.Get(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
