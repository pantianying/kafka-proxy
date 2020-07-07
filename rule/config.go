package rule

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/nacos_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/common/http_agent"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/sirupsen/logrus"
)

var (
	ruleConfig  *RuleConfig
	nacosClient config_client.ConfigClient
	nacosDataId = "kafka-proxy-config"
	nacosGroup  = "DEFAULT_GROUP"
)

func init() {
	//ruleConfig = &RuleConfig{
	//	ClientTopicRuleMap: map[string]*ClientTopicRule{
	//		"test-topic-yangchun": {
	//			ClientTopic: "test-topic-yangchun",
	//			TestKey:     []string{"aabbaabaabb"},
	//			BrokerTopicRuleMap: map[string]*BrokerTopicRule{
	//				"test-topic-yangchun1": {TestKey: []string{"aabbaabaabb"}},
	//			},
	//		},
	//	},
	//}
	nacosClient = createConfigClientTest()
	content, _ := nacosClient.GetConfig(vo.ConfigParam{
		DataId: nacosDataId,
		Group:  nacosGroup,
	})
	updateConfig(content)

	nacosClient.ListenConfig(vo.ConfigParam{
		DataId: nacosDataId,
		Group:  nacosGroup,
		OnChange: func(namespace, group, dataId, data string) {
			logrus.Info("config changed group:" + group + ", dataId:" + dataId + ", data:" + data)
			updateConfig(data)
		},
	})
}

func updateConfig(content string) error {
	if ruleConfig == nil {
		ruleConfig = &RuleConfig{}
	}
	err := yaml.Unmarshal([]byte(content), ruleConfig)
	if err != nil {
		panic(err)
	}
	for k := range ruleConfig.ClientTopicRuleMap {
		ruleConfig.ClientTopicRuleMap[k].ClientTopic = k
		tmpKeys := []string{}
		for _, v2 := range ruleConfig.ClientTopicRuleMap[k].BrokerTopicRuleMap {
			tmpKeys = append(tmpKeys, v2.TestKey...)
		}
		ruleConfig.ClientTopicRuleMap[k].TestKey = tmpKeys
	}
	bytes, _ := json.Marshal(ruleConfig)
	logrus.Infof("get nacos config:%+v", string(bytes))
	return err
}

func createConfigClientTest() config_client.ConfigClient {
	nacosClientConfig := constant.ClientConfig{
		TimeoutMs:           10 * 1000,
		BeatInterval:        5 * 1000,
		ListenInterval:      300 * 1000,
		NotLoadCacheAtStart: true,
		Username:            "nacos",
		Password:            "nacos",
	}

	nacosServerConfig := constant.ServerConfig{
		IpAddr:      "nacos-dev.dian.so",
		Port:        80,
		ContextPath: "/nacos",
	}

	nc := nacos_client.NacosClient{}
	nc.SetServerConfig([]constant.ServerConfig{nacosServerConfig})
	nc.SetClientConfig(nacosClientConfig)
	nc.SetHttpAgent(&http_agent.HttpAgent{})
	client, _ := config_client.NewConfigClient(&nc)
	return client
}
