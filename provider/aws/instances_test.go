package aws_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInstancesList(t *testing.T) {

	provider := StubAwsProvider()
	defer provider.Close()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	provider.EC2 = createEC2Mock(mockCtrl)
	provider.ECS = createECSContainerInstancesMock(mockCtrl)

	is, err := provider.InstanceList()

	assert.Nil(t, err)
	assert.EqualValues(
		t,
		ec2Instances,
		is,
	)
}
