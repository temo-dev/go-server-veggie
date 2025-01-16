package gin

import (
	"context"
	"net/http"
	commonError "server-veggie/common/error"
	"server-veggie/modules/aws/model"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func FindKeyS3Handler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bucketName = "veggie-bucket"
		var data *model.CreateLinkFileType
		//validate
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, commonError.ErrValidateInput(model.EntityName, err))
			return
		}
		validate := validator.New()
		if err := validate.Struct(data); err != nil {
			c.JSON(http.StatusBadRequest, commonError.ErrValidateInput(model.EntityName, err))
			return
		}
		// Load cấu hình AWS
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to load AWS config",
				"error":   err.Error(),
			})
			return
		}

		s3Client := s3.NewFromConfig(cfg)

		// Tạo Pre-signed URL
		presignClient := s3.NewPresignClient(s3Client)
		req, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
			Bucket:      &bucketName,
			Key:         &data.FileName,
			ContentType: &data.FileType,
			ACL:         types.ObjectCannedACLPublicRead, // Quyền truy cập file
		}, s3.WithPresignExpires(15*time.Minute))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to generate pre-signed URL",
				"error":   err.Error(),
			})
			return
		}

		// Trả Pre-signed URL cho client
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully generated pre-signed URL",
			"data":    req.URL,
		})
	}
}
