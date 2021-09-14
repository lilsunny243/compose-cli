/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package ecs

import (
	"context"

	"github.com/compose-spec/compose-go/types"

	"github.com/docker/compose/v2/pkg/api"
)

func (b *ecsAPIService) RunOneOffContainer(ctx context.Context, project *types.Project, opts api.RunOptions) (int, error) {
	return 0, api.ErrNotImplemented
}

func (b *ecsAPIService) Remove(ctx context.Context, project *types.Project, options api.RemoveOptions) error {
	return api.ErrNotImplemented
}