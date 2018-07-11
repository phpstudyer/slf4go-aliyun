package aliyun

import (
	"testing"

	"github.com/dynamicgo/go-config/source/file"

	config "github.com/dynamicgo/go-config"
)

func init() {
	config := config.NewConfig()

	config.Load(file.NewSource(file.WithPath("../config/aliyun.json")))
}

func TestAliyun(t *testing.T) {

}
