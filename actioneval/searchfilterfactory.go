// Copyright © 2017 Aqua Security Software Ltd. <info@aquasec.com>
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

package actioneval

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/aquasecurity/bench-common/common"
	"gopkg.in/yaml.v2"
)

type SearchFilterResult struct {
	Output  bytes.Buffer
	Errmsgs string
	State   common.State
	Lines   int
}

type ISearchFilter interface {
	SearchFilterHandler(workspacePath string, count bool) (result SearchFilterResult)
}

func SearchFilterFactory(searchFilterType string, mapSlice yaml.MapSlice, tarHeaders []tar.Header) (ISearchFilter, error) {

	switch searchFilterType {

	case common.TEXTSEARCH:
		return NewTextSearchFilter(mapSlice), nil
	case common.FILESEARCH:
		search, err := NewFileSearchFilter(mapSlice)
		return search.WithTarHeaders(tarHeaders), err
	default:
		return nil, fmt.Errorf("unsupported search type %s", searchFilterType)
	}
}
