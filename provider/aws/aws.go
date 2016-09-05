package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acm/acmiface"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/ecs/ecsiface"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/convox/logger"
)

var (
	customTopic       = os.Getenv("CUSTOM_TOPIC")
	notificationTopic = os.Getenv("NOTIFICATION_TOPIC")
	sortableTime      = "20060102.150405.000000000"
)

// Logger is a package-wide logger
var Logger = logger.New("ns=provider.aws")

type AWSProvider struct {
	Region   string
	Endpoint string
	Access   string
	Secret   string
	Token    string

	Cluster           string
	Development       bool
	DockerImageAPI    string
	DynamoBuilds      string
	DynamoReleases    string
	NotificationHost  string
	NotificationTopic string
	Password          string
	Rack              string
	RegistryHost      string
	SettingsBucket    string
	Subnets           string
	SubnetsPrivate    string
	Vpc               string
	VpcCidr           string

	SkipCache bool

	// AWS services
	ACM            acmiface.ACMAPI
	CloudFormation cloudformationiface.CloudFormationAPI
	CloudWatch     cloudwatchiface.CloudWatchAPI
	CloudWatchLogs cloudwatchlogsiface.CloudWatchLogsAPI
	DynamoDB       dynamodbiface.DynamoDBAPI
	EC2            ec2iface.EC2API
	ECR            ecriface.ECRAPI
	ECS            ecsiface.ECSAPI
	IAM            iamiface.IAMAPI
	S3             s3iface.S3API
	SNS            snsiface.SNSAPI
}

// NewProviderFromEnv returns a new AWS provider from env vars
func NewProviderFromEnv() *AWSProvider {
	p := &AWSProvider{
		Region:            os.Getenv("AWS_REGION"),
		Endpoint:          os.Getenv("AWS_ENDPOINT"),
		Access:            os.Getenv("AWS_ACCESS"),
		Secret:            os.Getenv("AWS_SECRET"),
		Token:             os.Getenv("AWS_TOKEN"),
		Cluster:           os.Getenv("CLUSTER"),
		Development:       os.Getenv("DEVELOPMENT") == "true",
		DockerImageAPI:    os.Getenv("DOCKER_IMAGE_API"),
		DynamoBuilds:      os.Getenv("DYNAMO_BUILDS"),
		DynamoReleases:    os.Getenv("DYNAMO_RELEASES"),
		NotificationHost:  os.Getenv("NOTIFICATION_HOST"),
		NotificationTopic: os.Getenv("NOTIFICATION_TOPIC"),
		Password:          os.Getenv("PASSWORD"),
		Rack:              os.Getenv("RACK"),
		RegistryHost:      os.Getenv("REGISTRY_HOST"),
		SettingsBucket:    os.Getenv("SETTINGS_BUCKET"),
		Subnets:           os.Getenv("SUBNETS"),
		SubnetsPrivate:    os.Getenv("SUBNETS_PRIVATE"),
		Vpc:               os.Getenv("VPC"),
		VpcCidr:           os.Getenv("VPCCIDR"),
	}

	s := session.New()

	p.ACM = acm.New(s, p.config())
	p.CloudFormation = cloudformation.New(s, p.config())
	p.CloudWatch = cloudwatch.New(s, p.config())
	p.CloudWatchLogs = cloudwatchlogs.New(s, p.config())
	p.DynamoDB = dynamodb.New(s, p.config())
	p.EC2 = ec2.New(s, p.config())
	p.ECR = ecr.New(s, p.config())
	p.ECS = ecs.New(s, p.config())
	p.IAM = iam.New(s, p.config())
	p.S3 = s3.New(s, p.config().WithS3ForcePathStyle(true)) // path style is easier to test
	p.SNS = sns.New(s, p.config())

	return p
}

/** services ****************************************************************************************/

func (p *AWSProvider) config() *aws.Config {
	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(p.Access, p.Secret, p.Token),
	}

	if p.Region != "" {
		config.Region = aws.String(p.Region)
	}

	if p.Endpoint != "" {
		config.Endpoint = aws.String(p.Endpoint)
	}

	if os.Getenv("DEBUG") != "" {
		config.WithLogLevel(aws.LogDebugWithHTTPBody)
	}

	return config
}

// IsTest returns true when we're in test mode
func (p *AWSProvider) IsTest() bool {
	return p.Region == "us-test-1"
}
