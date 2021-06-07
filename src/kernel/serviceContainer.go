package kernel

import "github.com/ArtisanCloud/go-libs/object"

const CONFIG_BASE_INDEX = 0
const CONFIG_DEFAULT_INDEX = 1
const CONFIG_USER_INDEX = 2

type ServiceContainer struct {
	ID int

	DefaultConfig object.HashMap
	UserConfig    object.HashMap
	Config        []object.HashMap
}

func (container *ServiceContainer) getBaseConfig() object.HashMap {
	return object.HashMap{
		// http://docs.guzzlephp.org/en/stable/request-options.html
		"http": object.HashMap{
			"timeout":  30.0,
			"base_uri": "https://api.weixin.qq.com/",
		},
	}
}

func (container *ServiceContainer) GetConfig() []object.HashMap {
	// if container config has been setup
	if container.Config!=nil{
		return container.Config
	}

	// setup the container config firstly
	baseConfig := container.getBaseConfig()
	container.Config = []object.HashMap{
		baseConfig,
		container.DefaultConfig,
		container.UserConfig,
	}
	return container.Config
}