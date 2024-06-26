// Copyright 2024 Michael Davis
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authentication

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/Michad/tilegroxy/internal/config"
	"github.com/google/uuid"
)

type StaticKeyConfig struct {
	Key string
}

type StaticKey struct {
	Config *StaticKeyConfig
}

func ConstructStaticKey(config *StaticKeyConfig, errorMessages *config.ErrorMessages) (*StaticKey, error) {
	if config.Key == "" {
		keyUuid, err := uuid.NewRandom()

		if err != nil {
			return nil, err
		}

		keyStr := strings.ReplaceAll(keyUuid.String(), "-", "")

		slog.Warn(fmt.Sprintf("Generated authentication key: %v\n", keyStr))
		config.Key = keyStr
	}

	return &StaticKey{config}, nil
}

func (c StaticKey) Preauth(req *http.Request) bool {
	h := req.Header["Authorization"]
	return len(h) > 0 && h[0] == "Bearer "+c.Config.Key
}
