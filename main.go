package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	sw "github.com/igaskin/go-honeygain/go"
)

func main() {
	client := sw.NewAPIClient(sw.NewConfiguration())
	auth := sw.ApiUsersTokensPostRequest{}
	r := auth.Auth(sw.Auth{
		Email:    aws.String(os.Getenv("HONEYGAIN_EMAIL")),
		Password: aws.String(os.Getenv("HONEYGAIN_PASSWORD")),
	})
	resp, err := client.DefaultApi.UsersTokensPostExecute(r)
	if err.Error() != "" {
		fmt.Println("failed to get token")
	}
	fmt.Println(resp)
}
