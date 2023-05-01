package wafcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-waf/authenticator"
	"github.com/Appkube-awsx/awsx-waf/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		print(authFlag)
		// authFlag := true
		if authFlag {
			webAclId, _ := cmd.Flags().GetString("webAclId")
			if webAclId != "" {
				getClusterDetails(region, crossAccountRoleArn, acKey, secKey, externalId, webAclId)
			} else {
				log.Fatalln("waf Acl Id not provided. Program exit")
			}
		}
	},
}

func getClusterDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string,  externalId string, webAclId string) *waf.GetWebACLOutput {
	log.Println("Getting aws waf details data")
	listClusterClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &waf.GetWebACLInput{
		WebACLId: aws.String(webAclId),
	}
	clusterDetailsResponse, err := listClusterClient.GetWebACL(input)
	log.Println(clusterDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return clusterDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("webAclId", "t", "", "web acl id")

	if err := GetConfigDataCmd.MarkFlagRequired("webAclId"); err != nil {
		fmt.Println(err)
	}
}
