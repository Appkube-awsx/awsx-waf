package command

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-waf/command/wafcmd"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/spf13/cobra"
)

var AwsxWafListCmd = &cobra.Command{
	Use:   "getWafListDetails",
	Short: "getWafListDetails command gets resource counts",
	Long:  `getWafListDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command getWafListDetails started")

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)

		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetWebAclList(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
}

func GetWebAclList(auth client.Auth) (*waf.ListWebACLsOutput, error) {
	log.Println("aws waf list details")
	wafclient := client.GetClient(auth, client.WAF_CLIENT).(*waf.WAF)
	wafRequest := &waf.ListWebACLsInput{}
	wafclusterResponse, err := wafclient.ListWebACLs(wafRequest)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println(wafclusterResponse)
	return wafclusterResponse, err
}

func Execute() {
	err := AwsxWafListCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxWafListCmd.AddCommand(wafcmd.GetConfigDataCmd)

	AwsxWafListCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxWafListCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxWafListCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxWafListCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxWafListCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxWafListCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxWafListCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn Required")
	AwsxWafListCmd.PersistentFlags().String("externalId", "", "aws externalId Required")

}
