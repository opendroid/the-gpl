package bot

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	dfMock "github.com/opendroid/the-gpl/mocks/df"

	"github.com/golang/mock/gomock"

	"github.com/opendroid/the-gpl/clients/df"
	"github.com/stretchr/testify/assert"
)

// TestInitCli tests initialization
func TestInitCli(t *testing.T) {
	InitCli()
	assert.Equal(t, df.GCPProjectID, *gcpProjectName)
	assert.Equal(t, df.DefaultLanguage, *lang)
	assert.Equal(t, false, *chat)
	assert.Equal(t, string(df.Draft), *env)
}

var command = []string{"-chat=false", "-project=unit-test"}
var agentResponses = []string{"cochatbot.hi_hcihy"}

// TestCLI_ExecCmd tests the chatting with the bot.
func TestCLI_ExecCmd(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	oldBot := bot
	logger = log.New(&bytes.Buffer{}, "BOT ", log.LstdFlags)
	defer func() { bot = oldBot }() // Restore bot
	mockBot := dfMock.NewMockBot(mockCtrl)
	first := mockBot.EXPECT().Converse(gomock.Any(), "hello").Return(agentResponses, nil).Times(1)
	second := mockBot.EXPECT().Converse(gomock.Any(), "i like to cancel").Return(agentResponses, nil).Times(1)
	third := mockBot.EXPECT().Converse(gomock.Any(), "taking too long").Return(agentResponses, nil).Times(1)
	gomock.InOrder(first, second, third)
	bot = mockBot

	InitCli()
	cmd.ExecCmd(command)
}

// TestCLI_ExecCmd tests the chatting with the bot.
func TestCLI_ExecCmd_BotNil(t *testing.T) {
	bot = nil
	logger = log.New(&bytes.Buffer{}, "BOT ", log.LstdFlags)
	InitCli()
	cmd.ExecCmd(command)
}

// TestCLI_ExecCmd_ConverseError tests the chatting with the bot.
func TestCLI_ExecCmd_ConverseError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	oldBot := bot
	logger = log.New(&bytes.Buffer{}, "BOT ", log.LstdFlags)
	defer func() { bot = oldBot }() // Restore bot
	mockBot := dfMock.NewMockBot(mockCtrl)
	mockBot.EXPECT().Converse(gomock.Any(), "hello").Return(nil, fmt.Errorf("mock Error.")).Times(1)
	bot = mockBot

	InitCli()
	cmd.ExecCmd(command)
}
