package cache

var topicConfig = &TopicConfig{}

type CfgCache interface {
	GetBrokerTopicByKey(key string) string
	GetClientTopicByKey(key string) string
}

type TopicConfig struct{}

func (t *TopicConfig) CheckIsReplaceTopic(key, value []byte) bool {
	return true
}

func GetTopicCfg() *TopicConfig {
	return topicConfig
}
