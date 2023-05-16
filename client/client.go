package client

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-waf/awssession"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/waf"
)

func GetClient(region string, crossAccountRoleArn string, accessKey string,  secretKey string, externalId string) *waf.WAF{
	
	sessionName := "assume_role_session_name"
	return assumeRole(crossAccountRoleArn, sessionName, accessKey, secretKey, region, externalId)
}

func assumeRole(roleArn string, sessionName string, accesskey string, secretKey string, region string, externalId string) *waf.WAF {
	

	sess, err := awssession.GetSessionByCreds(region, accesskey, secretKey, "")

	if err != nil {
		fmt.Printf("failed to create aws session, %v\n", err)
		log.Fatal(err)
	}

	svc := sts.New(sess)
	assumeRoleInput := sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: aws.Int64(60 * 60 * 1),
		ExternalId:  aws.String(externalId),
	}
	if externalId != "nil" {
		fmt.Println("Trying to fetch env to assume new role")
		assumeRoleInput.ExternalId = aws.String(externalId)
	}
	result, err := svc.AssumeRole(&assumeRoleInput)

	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}
	fmt.Println("Assume role output: ", result)
	awsSession, err := awssession.GetSessionByCreds(region, *result.Credentials.AccessKeyId, *result.Credentials.SecretAccessKey, *result.Credentials.SessionToken)
	if err != nil {
		fmt.Printf("failed to assume role, %v\n", err)
		log.Fatal(err)
	}
	dbclusterClient := waf.New(awsSession)
	return dbclusterClient
}

