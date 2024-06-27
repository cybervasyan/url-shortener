package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
	"net/url"
	"url-shortener/shortener"
	"url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"longUrl" binding:"required"`
	UserId  string `json:"userId" binding:"required"`
}

func isValidUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func CreateShortUrl(context *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := context.ShouldBindJSON(&creationRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidUrl(creationRequest.LongUrl) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Введите корректную ссылку"})
		return
	}

	shortUrl := shortener.GenerateShortenedUrl(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:1488/"
	fullShortUrl := host + shortUrl

	qrCode, err := qrcode.Encode(fullShortUrl, qrcode.Medium, 256)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": fullShortUrl,
		"qr_code":   qrCode,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveOriginalUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
