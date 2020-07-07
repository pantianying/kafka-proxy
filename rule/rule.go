package rule

import (
	"bytes"
	"fmt"
	"github.com/Shopify/sarama"
)

type RuleConfig struct {
	ClientTopicRuleMap map[string]*ClientTopicRule
}

// 一个clienttopic一个ClientTopicRule
type ClientTopicRule struct {
	ClientTopic string
	// 总的TestKey合集
	TestKey []string
	// brokerTopic--> testKey
	BrokerTopicRuleMap map[string]*BrokerTopicRule
}

type BrokerTopicRule struct {
	TestKey []string
}

func GetRuleCfg() *RuleConfig {
	return ruleConfig
}

func (r *RuleConfig) GetClientTopicRule(clientTopic string) sarama.ClientTopicRule {
	return r.ClientTopicRuleMap[clientTopic]
}

func (r *RuleConfig) CheckClientTopicIsRule(clientTopic string) bool {
	_, ok := r.ClientTopicRuleMap[clientTopic]
	return ok
}

func (r *RuleConfig) GetClientTopicByResponseTopic(brokerTopic string) string {
	for _, cr := range r.ClientTopicRuleMap {
		if _, ok := cr.BrokerTopicRuleMap[brokerTopic]; ok {
			return cr.ClientTopic
		}
	}
	return ""
}

func (c *ClientTopicRule) CheckIsReplaceTopic(key, value []byte) (bool, string) {
	for _, k := range c.TestKey {
		if k == string(key) || bytes.Contains(value, []byte(k)) {
			return true, k
		}
	}
	return false, ""
}

func (c *ClientTopicRule) BrokerTopic(configKey string) string {
	for brokerTopic, brokerRule := range c.BrokerTopicRuleMap {
		for _, k := range brokerRule.TestKey {
			if k == configKey {
				return brokerTopic
			}
		}
	}
	fmt.Println("unexpect BrokerTopic ", configKey, c)
	return ""
}
