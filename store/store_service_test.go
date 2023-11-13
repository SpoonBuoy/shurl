package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()

}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	ol := "https://www.google.com"
	userUuid := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "short_url_1"

	SaveUrlMapping(shortUrl, ol, userUuid)
	retrievedUrl := RetrieveOriginalurl(shortUrl)

	assert.Equal(t, ol, retrievedUrl)
}
