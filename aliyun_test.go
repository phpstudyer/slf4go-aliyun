package aliyun

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dynamicgo/slf4go"

	"github.com/dynamicgo/go-config/source/file"

	config "github.com/dynamicgo/go-config"
)

var logconfig config.Config

func init() {
	logconfig = config.NewConfig()

	logconfig.Load(file.NewSource(file.WithPath("../config/aliyun.json")))
}

var message = `
hello hello worldhello worldhello worldhello worldhello worldhello 
worldhello worldhello worldhello worldworldhello worldhello worldhello
 worldhello worldhello worldhello worldhello worldhello worldhello worldhello
  worldhello worldhello worldhello worldhello worldhello worldhello worldhello 
  worldhello worldhello worldhello world
  hello hello worldhello worldhello worldhello worldhello worldhello 
worldhello worldhello worldhello worldworldhello worldhello worldhello
 worldhello worldhello worldhello worldhello worldhello worldhello worldhello
  worldhello worldhello worldhello worldhello worldhello worldhello worldhello 
  worldhello worldhello worldhello world
  hello hello worldhello worldhello worldhello worldhello worldhello 
worldhello worldhello worldhello worldworldhello worldhello worldhello
 worldhello worldhello worldhello worldhello worldhello worldhello worldhello
  worldhello worldhello worldhello worldhello worldhello worldhello worldhello 
  worldhello worldhello worldhello world
  hello hello worldhello worldhello worldhello worldhello worldhello 
worldhello worldhello worldhello worldworldhello worldhello worldhello
 worldhello worldhello worldhello worldhello worldhello worldhello worldhello
  worldhello worldhello worldhello worldhello worldhello worldhello worldhello 
  worldhello worldhello worldhello world
  hello hello worldhello worldhello worldhello worldhello worldhello 
worldhello worldhello worldhello worldworldhello worldhello worldhello
 worldhello worldhello worldhello worldhello worldhello worldhello worldhello
  worldhello worldhello worldhello worldhello worldhello worldhello worldhello 
  worldhello worldhello worldhello world
hello hello worldhello worldhello worldhello worldhello worldhello 
worldhello worldhello worldhello worldworldhello worldhello worldhello
 worldhello worldhello worldhello worldhello worldhello worldhello worldhello
  worldhello worldhello worldhello worldhello worldhello worldhello worldhello 
  worldhello worldhello worldhello world`

func TestAliyun(t *testing.T) {

	err := slf4go.Load(logconfig)

	require.NoError(t, err)

	logger := slf4go.Get("test")

	for {
		logger.DebugF("test")
	}

}
