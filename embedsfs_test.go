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
	embedsfs *EmbedsFS
}

//go:embed embeds
var embeds embed.FS

func (suite *EmbedsFSSuite) BeforeTest(suiteName, testName string) {
	suite.embedsfs = NewEmbedsFS(embeds, "embeds")
}

func (suite *EmbedsFSSuite) Test_Filename() {
	actual := suite.embedsfs.Filename("readme.txt")
	assert.Equal(suite.T(), `embeds/readme.txt`, actual)
}

func (suite *EmbedsFSSuite) Test_FileContent() {
	actual, err := suite.embedsfs.Content("readme.txt")
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), actual)
	assert.Contains(suite.T(), fmt.Sprintf("%s", actual), "wu.shaohua@foxmail.com")
}
