package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity/cache"
)

func getMSCredentials() (azcore.TokenCredential, error) {
	record, err := readAuthRecord()
	if err != nil {
		log.Println("Error Credentials Auth: ", err)
	}
	c, _ := cache.New(nil)
	credentials, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		TenantID:             os.Getenv("TENANT_ID"),
		ClientID:             os.Getenv("APP_ID"),
		AuthenticationRecord: record,
		Cache:                c,
		UserPrompt: func(ctx context.Context, dcm azidentity.DeviceCodeMessage) error {

			fmt.Println(dcm.Message)
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	if record == (azidentity.AuthenticationRecord{}) {

		record, err = credentials.Authenticate(context.TODO(), &policy.TokenRequestOptions{
			Scopes: strings.Split(os.Getenv("SCOPES"), ","),
		})
		if err == nil {
			err = setAuthRecord(record)
		}
		log.Println("Error Credentials Auth: ", err)

	}

	return credentials, nil

}

func readAuthRecord() (azidentity.AuthenticationRecord, error) {

	var record = azidentity.AuthenticationRecord{}
	content, err := os.ReadFile("./auth.json")
	if err == nil {
		err = json.Unmarshal(content, &record)
	}

	return record, err

}

func setAuthRecord(record azidentity.AuthenticationRecord) error {
	jsn, err := json.Marshal(record)
	if err == nil {
		err = os.WriteFile("./auth.json", jsn, os.ModeAppend)
	}
	return err
}
