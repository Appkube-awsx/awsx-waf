package cmd

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
       

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey,crossAccountRoleArn, externalId)

		if authFlag {
			WebACLId, _ := cmd.Flags().GetString("WebACLId")
			if(WebACLId != ""){
				getCostDetails(region, crossAccountRoleArn, acKey, secKey, externalId)
			}else{
				log.Fatalln("WebACLId not provided. Program exit")
			}
			
		 }
	}, 
} 

func  GetWebACLData(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string, WebACLId string) *waf.GetWebACLOutput {
	log.Println("Getting aws waf metaData Count summary")
	webAclClient := client.GetClient(region, crossAccountRoleArn,  accessKey, secretKey, externalId)
	input := &waf.GetWebACLInput{
		WebACLId: aws.String(WebACLId),
	}
	webAclResponse, err := webAclClient.GetWebACL(input)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println(webAclResponse)
	return webAclResponse
}



func init() {
	GetConfigDataCmd.Flags().StringP("WebACLId", "t", "", "Web acl id ")

	if err := GetConfigDataCmd.MarkFlagRequired("WebACLId"); err != nil {
		fmt.Println(err)
	}
}