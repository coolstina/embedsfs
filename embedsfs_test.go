// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
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

package embedsfs

import (
	"embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestEmbedsFSSuite(t *testing.T) {
	suite.Run(t, &EmbedsFSSuite{})
}

type EmbedsFSSuite struct {
	suite.Suite
	embedsfsSpecifyPath *EmbedsFS
	embedsfs *EmbedsFS
}

//go:embed test/data/embeds
var embeds embed.FS

func (suite *EmbedsFSSuite) BeforeTest(suiteName, testName string) {
	suite.embedsfsSpecifyPath = NewEmbedsFS(embeds, WithPath("test/data/embeds"))
	suite.embedsfs = NewEmbedsFS(embeds)
}

func (suite *EmbedsFSSuite) Test_EmbedsFSSpecifyPath_Filename() {
	actual := suite.embedsfsSpecifyPath.Filename("readme.txt")
	assert.Equal(suite.T(), `test/data/embeds/readme.txt`, actual)

	actual = suite.embedsfsSpecifyPath.Filename("test/data/embeds/readme.txt")
	assert.Equal(suite.T(), `test/data/embeds/readme.txt`, actual)

	actual = suite.embedsfsSpecifyPath.Filename("api/swagger/spec.json")
	assert.Equal(suite.T(), `test/data/embeds/api/swagger/spec.json`, actual)
}

func (suite *EmbedsFSSuite) Test_EmbedsFSSpecifyPath_FileContent() {
	actual, err := suite.embedsfsSpecifyPath.Content("readme.txt")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), actual)
	assert.Contains(suite.T(), fmt.Sprintf("%s", actual), "wu.shaohua@foxmail.com")
}

func (suite *EmbedsFSSuite) Test_EmbedsFS_Filename() {
	actual := suite.embedsfs.Filename(
		"test/data/embeds/readme.txt")
	assert.Equal(suite.T(), `test/data/embeds/readme.txt`, actual)

	actual = suite.embedsfsSpecifyPath.Filename(
		"test/data/embeds/api/swagger/spec.json")
	assert.Equal(suite.T(), `test/data/embeds/api/swagger/spec.json`, actual)
}

func (suite *EmbedsFSSuite) Test_EmbedsFS_FileContent() {
	actual, err := suite.embedsfs.Content("test/data/embeds/readme.txt")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), actual)
	assert.Contains(suite.T(), fmt.Sprintf("%s", actual), "wu.shaohua@foxmail.com")
}

