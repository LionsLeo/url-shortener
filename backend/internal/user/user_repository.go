package user

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"url-shortener/db"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Repository struct {
	Rdb *db.Db
}

func NewUserRepository(rdb *db.Db) *Repository {
	return &Repository{Rdb: rdb}
}

func (r *Repository) CreateOtpEntry(key string, jsonData []byte, ctx *gin.Context) error {
	err := r.Rdb.Redis.JSONSet(ctx, key, "$", jsonData).Err() // Setting with a TTL of 10 minutes
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Set TTL on the key (for example, 5 minutes)
	ttl := 5 * time.Minute
	err = r.Rdb.Redis.Expire(ctx, key, ttl).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *Repository) SendSmsOtp(to string, otp string) error {
	if os.Getenv("APP_MODE") == "dev" {
		return nil
	}
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromPhone := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(fromPhone)
	params.SetBody("Your OTP is: " + otp)

	_, err := client.Api.CreateMessage(params)
	return err
}

func (r *Repository) GetOtp(key string, ctx *gin.Context) (string, error) {
	jsonStr, err := r.Rdb.Redis.JSONGet(ctx, key).Result()
	if err != nil {
		// If key doesn't exist, you might get a redis.Nil error
		// or a generic error if there's a connection issue, etc.
		return "", fmt.Errorf("Could not find the key " + key)
	}

	// 4. Unmarshal the JSON into a map (or a struct if you prefer)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return "", fmt.Errorf("Error unmarshalling JSON from Redis " + err.Error())
	}

	// 5. Extract the "otp" field
	//    data["otp"] will be an interface{}; cast it to string if it's guaranteed to be a string.
	otp, ok := data["otp"].(string)
	if !ok {
		log.Fatal("OTP field not found or not a string in JSON")
		return "", fmt.Errorf("OTP field not found or not a string in JSON")
	}
	return otp, nil
}
