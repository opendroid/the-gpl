package lissajous

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"image/gif"
	"os"
	"testing"
)

var config = Config{
	Cycles:     2,
	Resolution: 0.000001,
	Size:       512,
	NFrames:    12,
	DelayMS:    10,
}

// cd lissajous
// go test -run TestLissajous -v
func TestLissajous(t *testing.T) {
	gifImageName := "../../content/media/lis.gif"
	gifImage, err := os.Create(gifImageName)
	require.Nil(t, err)
	require.NotNil(t, gifImage)

	Lissajous(config, gifImage)
	_ = gifImage.Close()
	checkImageSizeAndType(t, gifImageName, "gif", 1024, 1024)
}

// go test -run TestLissajous_Interface -v
func TestLissajous_Interface(t *testing.T) {

	// Call on config
	gifImageName := "../../content/media/lis2.gif"
	gifImage, err := os.Create(gifImageName)
	require.Nil(t, err)
	require.NotNil(t, gifImage)
	config.Lissajous(gifImage)
	_ = gifImage.Close()
	checkImageSizeAndType(t, gifImageName, "gif", 1024, 1024)
}

// checkImageSizeAndType helper function
func checkImageSizeAndType(t *testing.T, gifImageName string, imageType string, imageWidth int, imageLength int) {
	// Check if image was written for size
	gifImage, err := os.Open(gifImageName)
	require.Nil(t, err)
	require.NotNil(t, gifImage)
	exifInfo, err := gif.Decode(gifImage)
	require.Nil(t, err)
	require.NotNil(t, imageType)
	assert.Equal(t, "gif", imageType)
	require.NotNil(t, exifInfo)
	t.Logf("Size: %v", exifInfo.Bounds().Size())
	imageSize := exifInfo.Bounds().Size()
	assert.Equal(t, imageWidth, imageSize.X)
	assert.Equal(t, imageLength, imageSize.Y)
}
