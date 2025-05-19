package qwen

import "github.com/cloudwego/eino/components/model"

// options is the specific options for the qwen
type options struct {
	// EnableThinking enables thinking mode
	// Optional. Default: base on the Model
	// https://help.aliyun.com/zh/model-studio/deep-thinking
	EnableThinking *bool `json:"enable_thinking,omitempty"`
}

// WithEnableThinking is the option to set the enable thinking for the model.
func WithEnableThinking(enableThinking bool) model.Option {
	return model.WrapImplSpecificOptFn(func(opt *options) {
		opt.EnableThinking = &enableThinking
	})
}
