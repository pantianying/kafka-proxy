package rule

import (
	"bytes"
	"fmt"
	"github.com/Shopify/sarama"
)

//
//type TopicRule interface {
//	GetClientTopicRule(clientTopic string) ClientTopicRule
//	// 判断该topic是否有测试配置
//	CheckClientTopicIsRule(clientTopic string) bool
//	// 通过测试topic找到真实的clientTopic
//	// 如果不是测试的则返回 ""
//	GetClientTopicByResponseTopic(brokerTopic string) string
//}
//
//type ClientTopicRule interface {
//	// 确认是不是需要替换topic
//	// 传入msg的key和value，返回是否匹配以及根据哪个配置ConfigKey匹配到
//	CheckIsReplaceTopic(key, value []byte) (bool, string)
//	// 通过配置值返回brokerTopic
//	BrokerTopic(ConfigKey string) string
//}

var ruleConfig *RuleConfig

func init() {
	ruleConfig = &RuleConfig{
		ClientTopicRuleMap: map[string]*ClientTopicRule{
			"test-topic-yangchun": {
				ClientTopic: "test-topic-yangchun",
				TestKey:     []string{"aabbaabaabb"},
				BrokerTopicTestKeyMap: map[string][]string{
					"test-topic-yangchun1": []string{"aabbaabaabb"},
				},
			},
		},
	}
}

type RuleConfig struct {
	ClientTopicRuleMap map[string]*ClientTopicRule
}

// 一个clienttopic一个ClientTopicRule
type ClientTopicRule struct {
	ClientTopic string
	// 总的TestKey合集
	TestKey []string
	// brokerTopic--> testKey
	BrokerTopicTestKeyMap map[string][]string
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
		if _, ok := cr.BrokerTopicTestKeyMap[brokerTopic]; ok {
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
	for brokerTopic, testKeys := range c.BrokerTopicTestKeyMap {
		for _, k := range testKeys {
			if k == configKey {
				return brokerTopic
			}
		}
	}
	fmt.Println("unexpect BrokerTopic ", configKey, c)
	return ""
}

//func (t *TopicConfig) BrokerTopic() string {
//	return "test-topic-yangchun1"
//}
//
//func (t *TopicConfig) ClientTopic() string {
//	return "test-topic-yangchun"
//}
