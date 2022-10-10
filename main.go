package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/ec2"
)

func NewMyStack(scope constructs.Construct, id string) (cdktf.TerraformStack, string) {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("us-east-2"),
	})

	tgw := ec2.NewEc2TransitGateway(stack, jsii.String("tgw"), &ec2.Ec2TransitGatewayConfig{
		Description: jsii.String("test-terraform"),
		Tags: &map[string]*string{
			"Name": jsii.String("CDKTF-Demo"),
		},
	})

	cdktf.NewTerraformOutput(stack, jsii.String("tgw_id"), &cdktf.TerraformOutputConfig{
		Value: tgw.Id(),
	})

	return stack, *tgw.Id()
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "learn-cdktf")

	app.Synth()
}
