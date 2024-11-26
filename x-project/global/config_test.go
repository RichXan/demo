package global

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name        string
		configPath  string
		shouldPanic bool
	}{
		{
			name:        "Valid config file",
			configPath:  "../config/config.yml",
			shouldPanic: false,
		},
		{
			name:        "Invalid config path",
			configPath:  "non_existent_config.yml",
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				assert.Panics(t, func() {
					InitConfig(tt.configPath)
				})
			} else {
				assert.NotPanics(t, func() {
					InitConfig(tt.configPath)
				})
				// 验证配置是否正确加载
				assert.NotEmpty(t, AppConfig.System.Env)
			}
		})
	}
}

func TestConfigValues(t *testing.T) {
	// 首先初始化配置
	InitConfig("../config/config.yml")

	// 测试系统配置
	t.Run("System Config", func(t *testing.T) {
		assert.NotEmpty(t, AppConfig.System.Env)
		assert.NotEmpty(t, AppConfig.System.Version)
	})

	// 可以根据实际配置结构添加更多测试
}

// 基准测试
func BenchmarkInitConfig(b *testing.B) {
	configPath := "../config/config.yml"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InitConfig(configPath)
	}
}
