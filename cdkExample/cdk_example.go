package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type Take2GoCdkStackProps struct {
	awscdk.StackProps
}

func NewTake2GoCdkStack(scope constructs.Construct, id string, props *Take2GoCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	table := awsdynamodb.NewTable(stack, jsii.String("myUserTable2"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("userTable2"),
	})

	// define our lambda here
	myFunction := awslambda.NewFunction(stack, jsii.String("myLambdaFunction2"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("lambda/function.zip"), nil),
		Handler: jsii.String("main"),
	})

	// define our api gateway
	api := awsapigateway.NewRestApi(stack, jsii.String("myAPIGateway2"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
			AllowMethods: jsii.Strings("GET", "POST", "DELETE", "PUT", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
		DeployOptions: &awsapigateway.StageOptions{
			LoggingLevel: awsapigateway.MethodLoggingLevel_INFO,
		},
		EndpointConfiguration: &awsapigateway.EndpointConfiguration{
			Types: &[]awsapigateway.EndpointType{awsapigateway.EndpointType_REGIONAL},
		},
	})

	// integrate the routes + methods to this endpoint
	integration := awsapigateway.NewLambdaIntegration(myFunction, nil)

	registerRoute := api.Root().AddResource(jsii.String("register"), nil)
	registerRoute.AddMethod(jsii.String("POST"), integration, nil)

	loginRoute := api.Root().AddResource(jsii.String("login"), nil)
	loginRoute.AddMethod(jsii.String("POST"), integration, nil)

	protectedRoute := api.Root().AddResource(jsii.String("protected"), nil)
	protectedRoute.AddMethod(jsii.String("GET"), integration, nil)

	table.GrantReadWriteData(myFunction)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewTake2GoCdkStack(app, "Take2GoCdkStack", &Take2GoCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
