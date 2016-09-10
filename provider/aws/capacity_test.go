package aws_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/convox/rack/api/structs"
	"github.com/convox/rack/provider/aws/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCapacityGet(t *testing.T) {
	provider := StubAwsProvider()
	defer provider.Close()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ecsMock := createClusterServciesMock(mockCtrl)

	ecsMock.EXPECT().DescribeTaskDefinition(&ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String("task-def-1"),
	}).Return(&ecs.DescribeTaskDefinitionOutput{
		TaskDefinition: &ecs.TaskDefinition{
			ContainerDefinitions: []*ecs.ContainerDefinition{
				&ecs.ContainerDefinition{
					Name: aws.String("container-1"),
					PortMappings: []*ecs.PortMapping{
						&ecs.PortMapping{
							ContainerPort: aws.Int64(8080),
							HostPort:      aws.Int64(9000),
						},
					},
					Cpu:    aws.Int64(1024),
					Memory: aws.Int64(2000),
				},
			},
		},
	}, nil).Times(2)

	ecsMock.EXPECT().DescribeTaskDefinition(&ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String("task-def-2"),
	}).Return(&ecs.DescribeTaskDefinitionOutput{
		TaskDefinition: &ecs.TaskDefinition{
			ContainerDefinitions: []*ecs.ContainerDefinition{
				&ecs.ContainerDefinition{
					Name: aws.String("container-2"),
					PortMappings: []*ecs.PortMapping{
						&ecs.PortMapping{
							ContainerPort: aws.Int64(8080),
							HostPort:      aws.Int64(9000),
						},
					},
					Cpu:    aws.Int64(1024),
					Memory: aws.Int64(2000),
				},
			},
		},
	}, nil).Times(2)

	provider.ECS = ecsMock
	r, err := provider.CapacityGet()

	assert.Nil(t, err)
	assert.EqualValues(t, &structs.Capacity{
		ClusterCPU:     3072,
		ClusterMemory:  1536,
		InstanceCPU:    1024,
		InstanceMemory: 512,
		ProcessCount:   9,
		ProcessCPU:     9216,
		ProcessMemory:  18000,
		ProcessWidth:   9,
	}, r)
}

func createClusterServciesMock(mockCtrl *gomock.Controller) *mocks.MockECSAPI {

	ecsMock := createECSContainerInstancesMock(mockCtrl)

	ecsMock.EXPECT().ListServices(&ecs.ListServicesInput{
		Cluster: aws.String("cluster-test"),
	}).Return(&ecs.ListServicesOutput{
		ServiceArns: []*string{
			aws.String("arn-service-1"),
		},
	}, nil)

	ecsMock.EXPECT().DescribeServices(&ecs.DescribeServicesInput{
		Cluster: aws.String("cluster-test"),
		Services: []*string{
			aws.String("arn-service-1"),
		},
	}).Return(&ecs.DescribeServicesOutput{
		Services: []*ecs.Service{
			&ecs.Service{
				TaskDefinition: aws.String("task-def-1"),
				DesiredCount:   aws.Int64(5),
				LoadBalancers: []*ecs.LoadBalancer{
					&ecs.LoadBalancer{
						LoadBalancerName: aws.String("elb-name"),
						ContainerName:    aws.String("container-1"),
						ContainerPort:    aws.Int64(8080),
					}},
				Deployments: []*ecs.Deployment{
					&ecs.Deployment{
						TaskDefinition: aws.String("task-def-1"),
						DesiredCount:   aws.Int64(5),
					}},
			},
			&ecs.Service{
				TaskDefinition: aws.String("task-def-2"),
				DesiredCount:   aws.Int64(4),
				LoadBalancers: []*ecs.LoadBalancer{
					&ecs.LoadBalancer{
						LoadBalancerName: aws.String("elb-name"),
						ContainerName:    aws.String("container-2"),
						ContainerPort:    aws.Int64(8080),
					}},
				Deployments: []*ecs.Deployment{
					&ecs.Deployment{
						TaskDefinition: aws.String("task-def-2"),
						DesiredCount:   aws.Int64(4),
					}},
			},
		},
	}, nil)

	return ecsMock
}
