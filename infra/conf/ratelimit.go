package conf

import (
	"github.com/golang/protobuf/proto"
	"v2ray.com/core/app/ratelimit"
)

type LimitSettingConfig struct {
	UpRate       int64 `json:"upRate"`
	UpCapacity   int64 `json:"upCapacity"`
	DownRate     int64 `json:"downRate"`
	DownCapacity int64 `json:"downCapacity"`
}

func (c *LimitSettingConfig) Build() (*ratelimit.LimitSettings, error) {
	return &ratelimit.LimitSettings{
		UpRate:       c.UpRate,
		UpCapacity:   c.UpCapacity,
		DownRate:     c.DownRate,
		DownCapacity: c.DownCapacity,
	}, nil
}

type LimitRuleConfig struct {
	Users       []string            `json:"user"`
	InboundTags []string            `json:"inboundTag"`
	Settings    *LimitSettingConfig `json:"settings"`
}

func (c *LimitRuleConfig) Build() (*ratelimit.LimitRule, error) {
	ruleConfig := &ratelimit.LimitRule{}
	ruleConfig.UserEmail = c.Users
	ruleConfig.InboundTag = c.InboundTags
	settingsConfig, err := c.Settings.Build()
	if err != nil {
		return nil, err
	}
	ruleConfig.Settings = settingsConfig

	return ruleConfig, nil
}

type RateLimitConfig struct {
	Rule []*LimitRuleConfig `json:"rules"`
}

func (c *RateLimitConfig) Build() (proto.Message, error) {
	config := &ratelimit.Config{}
	for _, rule := range c.Rule {
		ruleConfig, err := rule.Build()
		if err != nil {
			return nil, err
		}
		config.Rule = append(config.Rule, ruleConfig)
	}

	return config, nil
}
