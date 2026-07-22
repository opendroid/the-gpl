package aitutor

import (
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	anthropicMock "github.com/opendroid/the-gpl/mocks/anthropic"
)

// TestTutor_Ask_Success tests a plain question with no chapter context.
func TestTutor_Ask_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := anthropicMock.NewMockClient(mockCtrl)
	mockClient.EXPECT().
		Ask(gomock.Any(), gomock.Any(), "What is a goroutine?").
		Return("A goroutine is a lightweight thread.", nil).
		Times(1)

	tutor := NewTutor(mockClient)
	answer, err := tutor.Ask("What is a goroutine?", "")
	assert.NoError(t, err)
	assert.Equal(t, "A goroutine is a lightweight thread.", answer)
}

// TestTutor_Ask_ChapterContext tests that chapter context is folded into the
// user content sent to the client.
func TestTutor_Ask_ChapterContext(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := anthropicMock.NewMockClient(mockCtrl)
	mockClient.EXPECT().
		Ask(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ interface{}, _ string, userContent string) (string, error) {
			assert.True(t, strings.Contains(userContent, "Chapter context:"))
			assert.True(t, strings.Contains(userContent, "package chapter5"))
			assert.True(t, strings.Contains(userContent, "How does HTML traversal work?"))
			return "It walks the DOM tree.", nil
		}).
		Times(1)

	tutor := NewTutor(mockClient)
	answer, err := tutor.Ask("How does HTML traversal work?", "package chapter5")
	assert.NoError(t, err)
	assert.Equal(t, "It walks the DOM tree.", answer)
}

// TestTutor_Ask_Error tests that client errors propagate to the caller.
func TestTutor_Ask_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := anthropicMock.NewMockClient(mockCtrl)
	mockClient.EXPECT().
		Ask(gomock.Any(), gomock.Any(), gomock.Any()).
		Return("", fmt.Errorf("claude API error: rate limited")).
		Times(1)

	tutor := NewTutor(mockClient)
	answer, err := tutor.Ask("What is a goroutine?", "")
	assert.Error(t, err)
	assert.Equal(t, "", answer)
}
