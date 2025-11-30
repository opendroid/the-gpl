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

// TestNewBotCmd tests command creation and default flags
func TestNewBotCmd(t *testing.T) {
	cmd := NewBotCmd()
	assert.NotNil(t, cmd)
	assert.Equal(t, "bot", cmd.Use)

	// Check default flags
	project, _ := cmd.Flags().GetString("project")
	assert.Equal(t, df.GCPProjectID, project)

	lang, _ := cmd.Flags().GetString("lang")
	assert.Equal(t, df.DefaultLanguage, lang)

	chat, _ := cmd.Flags().GetBool("chat")
	assert.Equal(t, false, chat)

	env, _ := cmd.Flags().GetString("env")
	assert.Equal(t, string(df.Draft), env)
}

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

	cmd := NewBotCmd()
	cmd.SetArgs([]string{"--chat=false", "--project=unit-test"})
	err := cmd.Execute()
	assert.NoError(t, err)
}

// TestCLI_ExecCmd_BotNil tests the chatting with the bot when bot is nil.
func TestCLI_ExecCmd_BotNil(t *testing.T) {
	bot = nil
	logger = log.New(&bytes.Buffer{}, "BOT ", log.LstdFlags)

	cmd := NewBotCmd()
	cmd.SetArgs([]string{"--chat=false", "--project=unit-test"})
	err := cmd.Execute()
	assert.NoError(t, err)
}

// TestCLI_ExecCmd_ConverseError tests the chatting with the bot when converse returns error.
func TestCLI_ExecCmd_ConverseError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	oldBot := bot
	logger = log.New(&bytes.Buffer{}, "BOT ", log.LstdFlags)
	defer func() { bot = oldBot }() // Restore bot
	mockBot := dfMock.NewMockBot(mockCtrl)
	mockBot.EXPECT().Converse(gomock.Any(), "hello").Return(nil, fmt.Errorf("mock Error.")).Times(1)
	bot = mockBot

	cmd := NewBotCmd()
	cmd.SetArgs([]string{"--chat=false", "--project=unit-test"})
	err := cmd.Execute()
	assert.NoError(t, err)
}
