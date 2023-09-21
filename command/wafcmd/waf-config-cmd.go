package wafcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "Config data of waf",
	Long:  `Config data of waf`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			webAclId, _ := cmd.Flags().GetString("webAclId")

			if webAclId != "" {
				GetClusterDetails(webAclId, *clientAuth)
			} else {
				log.Fatalln("waf web acl id not provided. program exit")
			}
		}
	},
}

func GetClusterDetails(webAclId string, auth client.Auth) *waf.GetWebACLOutput {

	log.Println("Getting aws waf details data")

	listClusterClient := client.GetClient(auth, client.WAF_CLIENT).(*waf.WAF)

	input := &waf.GetWebACLInput{
		WebACLId: aws.String(webAclId),
	}

	clusterDetailsResponse, err := listClusterClient.GetWebACL(input)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	log.Println(clusterDetailsResponse.String())

	return clusterDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("webAclId", "w", "", "web acl id")

	if err := GetConfigDataCmd.MarkFlagRequired("webAclId"); err != nil {
		fmt.Println(err)
	}
}
