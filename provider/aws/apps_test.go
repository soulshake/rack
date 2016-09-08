package aws_test

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/convox/rack/api/structs"
	"github.com/convox/rack/provider/aws/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("RACK", "convox")
}

var expectedHttpApp = &structs.App{
	Name:    "httpd",
	Release: "RFAKEID",
	Status:  "running",
	Outputs: map[string]string{
		"BalancerWebHost":       "httpd-web-7E5UPCM-1241527783.us-east-1.elb.amazonaws.com",
		"LogGroup":              "convox-httpd-LogGroup-L4V203L35WRM",
		"RegistryId":            "132866487567",
		"RegistryRepository":    "convox-httpd-hqvvfosgxt",
		"Settings":              "convox-httpd-settings-139bidzalmbtu",
		"WebPort80Balancer":     "80",
		"WebPort80BalancerName": "httpd-web-7E5UPCM",
	},
	Parameters: map[string]string{
		"WebMemory":              "256",
		"WebCpu":                 "256",
		"Release":                "RFAKEID",
		"Subnets":                "subnet-13de3139,subnet-b5578fc3,subnet-21c13379",
		"Internal":               "No",
		"WebPort80ProxyProtocol": "No",
		"VPC":                  "vpc-f8006b9c",
		"Cluster":              "convox-Cluster-1E4XJ0PQWNAYS",
		"Key":                  "arn:aws:kms:us-east-1:132866487567:key/d9f38426-9017-4931-84f8-604ad1524920",
		"Repository":           "",
		"WebPort80Balancer":    "80",
		"Environment":          "https://convox-httpd-settings-139bidzalmbtu.s3.amazonaws.com/releases/RFAKEID/env",
		"WebPort80Certificate": "",
		"WebPort80Host":        "56694",
		"WebDesiredCount":      "1",
		"WebPort80Secure":      "No",
		"Version":              "20160330143438",
	},
	Tags: map[string]string{
		"Name":   "httpd",
		"Type":   "app",
		"System": "convox",
		"Rack":   "convox",
	},
}

func TestAppGet(t *testing.T) {
	provider := StubAwsProvider()
	defer provider.Close()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfMock := createHttpdStackMock(mockCtrl, provider.Rack)

	provider.CloudFormation = cfMock

	a, err := provider.AppGet("httpd")

	assert.Nil(t, err)
	assert.EqualValues(t, expectedHttpApp, a)
}

func createHttpdStackMock(mockCtrl *gomock.Controller, rack string) *mocks.MockCloudFormationAPI {

	cfMock := mocks.NewMockCloudFormationAPI(mockCtrl)
	cfMock.EXPECT().DescribeStacks(&cloudformation.DescribeStacksInput{
		StackName: aws.String(rack + "-" + "httpd"),
	}).Return(&cloudformation.DescribeStacksOutput{
		Stacks: []*cloudformation.Stack{
			&cloudformation.Stack{
				StackName:   aws.String("convox-httpd"),
				StackStatus: aws.String("UPDATE_COMPLETE"),
				Tags: []*cloudformation.Tag{
					&cloudformation.Tag{
						Key:   aws.String("Name"),
						Value: aws.String("httpd"),
					},
					&cloudformation.Tag{
						Key:   aws.String("Type"),
						Value: aws.String("app"),
					},
					&cloudformation.Tag{
						Key:   aws.String("System"),
						Value: aws.String("convox"),
					},
					&cloudformation.Tag{
						Key:   aws.String("Rack"),
						Value: aws.String("convox"),
					},
				},
				Parameters: []*cloudformation.Parameter{
					&cloudformation.Parameter{
						ParameterKey:   aws.String("Release"),
						ParameterValue: aws.String("RFAKEID"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("Internal"),
						ParameterValue: aws.String("No"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("WebMemory"),
						ParameterValue: aws.String("256"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("WebCpu"),
						ParameterValue: aws.String("256"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("Subnets"),
						ParameterValue: aws.String("subnet-13de3139,subnet-b5578fc3,subnet-21c13379"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("WebPort80ProxyProtocol"),
						ParameterValue: aws.String("No"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("VPC"),
						ParameterValue: aws.String("vpc-f8006b9c"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("Cluster"),
						ParameterValue: aws.String("convox-Cluster-1E4XJ0PQWNAYS"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("Key"),
						ParameterValue: aws.String("arn:aws:kms:us-east-1:132866487567:key/d9f38426-9017-4931-84f8-604ad1524920"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("Repository"),
						ParameterValue: aws.String(""),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("WebPort80Balancer"),
						ParameterValue: aws.String("80"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("Environment"),
						ParameterValue: aws.String("https://convox-httpd-settings-139bidzalmbtu.s3.amazonaws.com/releases/RFAKEID/env"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("WebPort80Certificate"),
						ParameterValue: aws.String(""),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("WebPort80Host"),
						ParameterValue: aws.String("56694"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("WebDesiredCount"),
						ParameterValue: aws.String("1"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("WebPort80Secure"),
						ParameterValue: aws.String("No"),
					},
					&cloudformation.Parameter{
						ParameterKey:   aws.String("Version"),
						ParameterValue: aws.String("20160330143438"),
					},
				},
				Outputs: []*cloudformation.Output{
					&cloudformation.Output{
						OutputKey:   aws.String("BalancerWebHost"),
						OutputValue: aws.String("httpd-web-7E5UPCM-1241527783.us-east-1.elb.amazonaws.com"),
					},
					&cloudformation.Output{
						OutputKey:   aws.String("LogGroup"),
						OutputValue: aws.String("convox-httpd-LogGroup-L4V203L35WRM"),
					},
					&cloudformation.Output{
						OutputKey:   aws.String("RegistryId"),
						OutputValue: aws.String("132866487567"),
					},
					&cloudformation.Output{
						OutputKey:   aws.String("RegistryRepository"),
						OutputValue: aws.String("convox-httpd-hqvvfosgxt"),
					},
					&cloudformation.Output{
						OutputKey:   aws.String("Settings"),
						OutputValue: aws.String("convox-httpd-settings-139bidzalmbtu"),
					},
					&cloudformation.Output{
						OutputKey:   aws.String("WebPort80Balancer"),
						OutputValue: aws.String("80"),
					},
					&cloudformation.Output{
						OutputKey:   aws.String("WebPort80BalancerName"),
						OutputValue: aws.String("httpd-web-7E5UPCM"),
					},
				},
			},
		},
	}, nil)

	return cfMock
}
