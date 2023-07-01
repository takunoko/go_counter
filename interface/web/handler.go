package web

import (
	"github.com/creasty/defaults"
)

type ApiHandlers struct{}

func Initialize(resVal any) error {
	// レスポンスパラメータのデフォルト値を設定
	if err := defaults.Set(resVal); err != nil {
		return err
	}
	return nil
}
