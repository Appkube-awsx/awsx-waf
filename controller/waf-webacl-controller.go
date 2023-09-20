package controller

import (
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-waf/cmd"
	"github.com/aws/aws-sdk-go/service/waf"
)

func GetWafByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) ([]*waf.ListWebACLsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetWAfByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetWafByUserCreds(region string, accesskey string, secretKey string, crossAccountRoleArn string, externalId string) ([]*waf.ListWebACLsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accesskey, secretKey, crossAccountRoleArn, externalId)
	return GetWafByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetWafByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) ([]*waf.ListWebACLsOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := cmd.GetInstanceList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetWafWebAcl(clientAuth *client.Auth) ([]*waf.ListWebACLsOutput, error) {
	response, err := cmd.GetInstanceList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
