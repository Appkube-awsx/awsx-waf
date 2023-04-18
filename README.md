# WAF CLi's

## To list all the WAF webAclId ,run the following command:

```bash
awsx-waf --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId> --env <env>
```

## To retrieve the configuration details of a specific WAF wafcmd, run the following command:

```bash
awsx-waf getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --webAclId <webAclId>
```

## To retrieve the cost and usage details of a specific WAF wafcmd run the following command:

```bash
awsx-waf getCostData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env>
```

## To retrieve the cost Spikes of a specific WAF wafcmd, run the following command:

```bash
awsx-waf GetCostSpike -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --granularity <granularity> --startDate <startDate> --endDate <endDate> 
```
