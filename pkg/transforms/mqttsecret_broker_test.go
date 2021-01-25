// +build brokerRunning
//
// Copyright (c) 2020 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// This test will only be executed if the tag brokerRunning is added when running
// the tests with a command like:
// go test -tags brokerRunning

package transforms

import (
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/interfaces/mocks"
)

func TestMQTTSendWithData(t *testing.T) {
	sender := NewMQTTSecretSender(MQTTSecretConfig{}, true)
	sender.mqttConfig = MQTTSecretConfig{
		SecretPath: "/mqtt",
	}
	sender.client = MQTT.NewClient(sender.opts)
	mockSecretProvider := &mocks.SecretProvider{}
	mockSecretProvider.On("SecretsLastUpdated").Return(time.Now())
	context.SecretProvider = mockSecretProvider
	sender.MQTTSend(context, "sendme")
	// require.True(t, continuePipeline)
	// require.Error(t, result.(error))
}