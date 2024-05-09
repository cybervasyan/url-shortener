package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	originalLink1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink1 := generateShortenedUrl(originalLink1, UserId)

	originalLink2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink2 := generateShortenedUrl(originalLink2, UserId)

	originalLink3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink3 := generateShortenedUrl(originalLink3, UserId)

	assert.Equal(t, shortLink1, "jTa4L57P")
	assert.Equal(t, shortLink2, "d66yfx7N")
	assert.Equal(t, shortLink3, "dhZTayYQ")
}
