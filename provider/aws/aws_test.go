package aws_test

import (
	"bytes"
	"net/http/httptest"
	"time"

	awsgo "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/convox/logger"
	"github.com/convox/rack/api/awsutil"
	"github.com/convox/rack/api/structs"
	"github.com/convox/rack/provider/aws"
	"github.com/convox/rack/provider/aws/mocks"
	"github.com/golang/mock/gomock"
)

var sortableTime = "20060102.150405.000000000"

func init() {
	logger.Output = &bytes.Buffer{}
}

type AwsStub struct {
	*aws.AWSProvider
	server *httptest.Server
}

func (a *AwsStub) Close() {
	a.server.Close()
}

// StubAwsProvider creates an httptest server with canned Request / Response
// cycles, and sets CurrentProvider to a new AWS provider that uses
// the test server as the endpoint
func StubAwsProvider(cycles ...awsutil.Cycle) *AwsStub {
	handler := awsutil.NewHandler(cycles)
	s := httptest.NewServer(handler)

	p := &aws.AWSProvider{
		Region:           "us-test-1",
		Endpoint:         s.URL,
		Access:           "test-access",
		Secret:           "test-secret",
		Token:            "test-token",
		Cluster:          "cluster-test",
		Development:      true,
		DockerImageAPI:   "rack/web",
		DynamoBuilds:     "convox-builds",
		DynamoReleases:   "convox-releases",
		NotificationHost: "notifications.example.org",
		Password:         "password",
		Rack:             "convox",
		RegistryHost:     "registry.example.org",
		SettingsBucket:   "convox-settings",
		SkipCache:        true,
	}

	return &AwsStub{p, s}
}

var ec2Instance1 = structs.Instance{
	Agent:     true,
	Cpu:       0,
	Id:        "i-test-id-1",
	Memory:    0,
	PrivateIp: "127.0.0.1",
	Processes: 0,
	PublicIp:  "127.0.0.2",
	Status:    "active",
	Started:   time.Time{},
}

var ec2Instance2 = structs.Instance{
	Agent:     true,
	Cpu:       0,
	Id:        "i-test-id-2",
	Memory:    0,
	PrivateIp: "128.0.0.1",
	Processes: 0,
	PublicIp:  "128.0.0.2",
	Status:    "active",
	Started:   time.Time{},
}

var ec2Instance3 = structs.Instance{
	Agent:     false,
	Cpu:       0,
	Id:        "i-test-id-3",
	Memory:    0,
	PrivateIp: "129.0.0.1",
	Processes: 0,
	PublicIp:  "129.0.0.2",
	Status:    "active",
	Started:   time.Time{},
}

var ec2Instances = structs.Instances{
	ec2Instance1,
	ec2Instance2,
	ec2Instance3,
}

// createECSContainerInstancesMock creates a standard generic mock struct for ECS.
func createECSContainerInstancesMock(mockCtrl *gomock.Controller) *mocks.MockECSAPI {

	ecsMock := mocks.NewMockECSAPI(mockCtrl)
	ecsMock.EXPECT().ListContainerInstances(&ecs.ListContainerInstancesInput{
		Cluster:   awsgo.String("cluster-test"),
		NextToken: awsgo.String(""),
	}).Return(&ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []*string{
			awsgo.String("arn-test-instance-1"),
			awsgo.String("arn-test-instance-2"),
			awsgo.String("arn-test-instance-3"),
		},
	}, nil)

	ecsMock.EXPECT().DescribeContainerInstances(&ecs.DescribeContainerInstancesInput{
		Cluster: awsgo.String("cluster-test"),
		ContainerInstances: []*string{
			awsgo.String("arn-test-instance-1"),
			awsgo.String("arn-test-instance-2"),
			awsgo.String("arn-test-instance-3"),
		},
	}).Return(&ecs.DescribeContainerInstancesOutput{
		ContainerInstances: []*ecs.ContainerInstance{
			&ecs.ContainerInstance{
				AgentConnected: awsgo.Bool(ec2Instance1.Agent),
				Ec2InstanceId:  awsgo.String(ec2Instance1.Id),
				Status:         awsgo.String(ec2Instance1.Status),
				RegisteredResources: []*ecs.Resource{
					&ecs.Resource{
						Name:         awsgo.String("MEMORY"),
						IntegerValue: awsgo.Int64(512),
					},
					&ecs.Resource{
						Name:         awsgo.String("CPU"),
						IntegerValue: awsgo.Int64(1024),
					},
				},
				RemainingResources: []*ecs.Resource{},
			},
			&ecs.ContainerInstance{
				AgentConnected: awsgo.Bool(ec2Instance2.Agent),
				Ec2InstanceId:  awsgo.String(ec2Instance2.Id),
				Status:         awsgo.String(ec2Instance2.Status),
				RegisteredResources: []*ecs.Resource{
					&ecs.Resource{
						Name:         awsgo.String("MEMORY"),
						IntegerValue: awsgo.Int64(512),
					},
					&ecs.Resource{
						Name:         awsgo.String("CPU"),
						IntegerValue: awsgo.Int64(1024),
					},
				},
				RemainingResources: []*ecs.Resource{},
			},
			&ecs.ContainerInstance{
				AgentConnected: awsgo.Bool(ec2Instance3.Agent),
				Ec2InstanceId:  awsgo.String(ec2Instance3.Id),
				Status:         awsgo.String(ec2Instance3.Status),
				RegisteredResources: []*ecs.Resource{
					&ecs.Resource{
						Name:         awsgo.String("MEMORY"),
						IntegerValue: awsgo.Int64(512),
					},
					&ecs.Resource{
						Name:         awsgo.String("CPU"),
						IntegerValue: awsgo.Int64(1024),
					},
				},
				RemainingResources: []*ecs.Resource{},
			},
		},
	}, nil)

	return ecsMock
}

// createEC2Mock creates a standard generic mock struct for EC2.
// Expected EC2 calls may not be called and order is not enforced.
func createEC2Mock(mockCtrl *gomock.Controller) *mocks.MockEC2API {

	ec2Mock := mocks.NewMockEC2API(mockCtrl)
	ec2Mock.EXPECT().DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{Name: awsgo.String("instance-id"), Values: []*string{
				awsgo.String(ec2Instance1.Id),
				awsgo.String(ec2Instance2.Id),
				awsgo.String(ec2Instance3.Id),
			}},
		},
		MaxResults: awsgo.Int64(1000),
	}).Return(&ec2.DescribeInstancesOutput{
		NextToken: awsgo.String(""),
		Reservations: []*ec2.Reservation{
			&ec2.Reservation{
				Instances: []*ec2.Instance{
					&ec2.Instance{
						InstanceId:       awsgo.String(ec2Instance1.Id),
						PrivateIpAddress: awsgo.String(ec2Instance1.PrivateIp),
						PublicIpAddress:  awsgo.String(ec2Instance1.PublicIp),
						LaunchTime:       awsgo.Time(time.Time{}),
					},
					&ec2.Instance{
						InstanceId:       awsgo.String(ec2Instance2.Id),
						PrivateIpAddress: awsgo.String(ec2Instance2.PrivateIp),
						PublicIpAddress:  awsgo.String(ec2Instance2.PublicIp),
						LaunchTime:       awsgo.Time(time.Time{}),
					},
					&ec2.Instance{
						InstanceId:       awsgo.String(ec2Instance3.Id),
						PrivateIpAddress: awsgo.String(ec2Instance3.PrivateIp),
						PublicIpAddress:  awsgo.String(ec2Instance3.PublicIp),
						LaunchTime:       awsgo.Time(time.Time{}),
					},
				},
			},
		},
	}, nil)

	return ec2Mock
}
