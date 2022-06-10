package cloudbucket

import (
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

func HandleFileUpload(c *gin.Context) {
	bucket := "medlit-bucket"

	var err error

	ctx := appengine.NewContext(c.Request)

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": "true",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"error": "false",
		"message": "File uploaded successfully",
		"pathname": "https://storage.googleapis.com" + u.EscapedPath(),
	})
}