package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-waf/authenticator"
	"github.com/Appkube-awsx/awsx-waf/client"
	cmd "github.com/Appkube-awsx/awsx-waf/cmd/wafcmd"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/spf13/cobra"
)

var getWafWebACLmetdataCmd = &cobra.Command{
	Use:   "getwafMetadata",
	Short: "getwafMetadata command gets resource counts",
	Long:  `getwafMetadata command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command waf details started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()
		

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey,crossAccountRoleArn, externalId)
		

		if authFlag {
			getListWebACLData(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
		
	},
}

func  getListWebACLData(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) *waf.ListWebACLsOutput  {
	log.Println("Getting aws waf metaData Count summary")
	webAclClient := client.GetClient(region, crossAccountRoleArn,  accessKey, secretKey, externalId)
	input := &waf.ListWebACLsInput{
		
	}
	webAclResponse, err := webAclClient.ListWebACLs(input)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println(webAclResponse)
	return webAclResponse
}

func Execute() {
	err := getWafWebACLmetdataCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	
	getWafWebACLmetdataCmd.AddCommand(cmd.GetConfigDataCmd)
	getWafWebACLmetdataCmd.AddCommand(cmd.GetCostDataCmd)

	getWafWebACLmetdataCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	getWafWebACLmetdataCmd.PersistentFlags().String("accountId", "", "aws account number")
	getWafWebACLmetdataCmd.PersistentFlags().String("zone", "", "aws region")
	getWafWebACLmetdataCmd.PersistentFlags().String("accessKey", "", "aws access key")
	getWafWebACLmetdataCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	getWafWebACLmetdataCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	getWafWebACLmetdataCmd.PersistentFlags().String("externalId", "", "aws externalId is required key")
	getWafWebACLmetdataCmd.PersistentFlags().String("WebACLId", "", "aws WebACLId is required key")
	
	

}
