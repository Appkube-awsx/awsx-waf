# WAF CLi's

## To list all the WAF webAclId ,run the following command:

```bash
awsx-waf --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId>
```

## To retrieve the configuration details of a specific WAF wafcmd, run the following command:

```bash
awsx-waf getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId>  --webAclId <webAclId>
```
