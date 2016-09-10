package aws_test

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/convox/rack/api/structs"
	"github.com/convox/rack/provider/aws/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func init() {
	//os.Setenv("RACK", "convox")
	//os.Setenv("DYNAMO_BUILDS", "convox-builds")
	//os.Setenv("DYNAMO_RELEASES", "convox-releases")
	//models.PauseNotifications = true
}

var expectedBuild1 = &structs.Build{
	Id:       "BFAKEID",
	App:      "httpd",
	Logs:     "",
	Manifest: "version: \"2\"\nservices:\n  web:\n    build: {}\n    command: null\n    image: httpd\n    ports:\n    - 80:80\n",
	Release:  "RFAKEID",
	Status:   "complete",
	Started:  time.Unix(1459780456, 178278576).UTC(),
	Ended:    time.Unix(1459780542, 440881687).UTC(),
}

func TestBuildGet(t *testing.T) {
	provider := StubAwsProvider()
	defer provider.Close()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	provider.DynamoDB = createDynamoGetItemMock(mockCtrl, provider.DynamoBuilds)
	b, err := provider.BuildGet("httpd", "BFAKEID")

	assert.Nil(t, err)
	assert.EqualValues(t, expectedBuild1, b)
}

func TestBuildDelete(t *testing.T) {
	provider := StubAwsProvider()
	defer provider.Close()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	provider.DynamoDB = createGetDeleteBuildMock(mockCtrl)
	provider.CloudFormation = createAppStackMock(mockCtrl, provider.Rack)
	provider.ECR = createECRMock(mockCtrl)

	b, err := provider.BuildDelete("httpd", "BFAKEID")

	assert.Nil(t, err)
	assert.EqualValues(t, expectedBuild1, b)
}

func TestBuildList(t *testing.T) {
	provider := StubAwsProvider()
	defer provider.Close()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfMock := createAppStackMock(mockCtrl, provider.Rack)
	dynamoMock := mocks.NewMockDynamoDBAPI(mockCtrl)

	dynamoMock.EXPECT().Query(&dynamodb.QueryInput{
		KeyConditions: map[string]*dynamodb.Condition{
			"app": &dynamodb.Condition{
				AttributeValueList: []*dynamodb.AttributeValue{&dynamodb.AttributeValue{S: aws.String("httpd")}},
				ComparisonOperator: aws.String("EQ"),
			},
		},
		IndexName:        aws.String("app.created"),
		Limit:            aws.Int64(20),
		ScanIndexForward: aws.Bool(false),
		TableName:        aws.String(provider.DynamoBuilds),
	}).Return(&dynamodb.QueryOutput{
		Items: []map[string]*dynamodb.AttributeValue{
			map[string]*dynamodb.AttributeValue{
				"app":      &dynamodb.AttributeValue{S: aws.String("httpd")},
				"id":       &dynamodb.AttributeValue{S: aws.String("BHINCLZYYVN")},
				"manifest": &dynamodb.AttributeValue{S: aws.String("web:\n  image: httpd\n  ports:\n  - 80:80\n")},
				"release":  &dynamodb.AttributeValue{S: aws.String("RVFETUHHKKD")},
				"status":   &dynamodb.AttributeValue{S: aws.String("complete")},
				"created":  &dynamodb.AttributeValue{S: aws.String(time.Unix(1459780456, 178278576).UTC().Format(sortableTime))},
				"ended":    &dynamodb.AttributeValue{S: aws.String(time.Unix(1459780542, 440881687).UTC().Format(sortableTime))},
			},
			map[string]*dynamodb.AttributeValue{
				"app":      &dynamodb.AttributeValue{S: aws.String("httpd")},
				"id":       &dynamodb.AttributeValue{S: aws.String("BNOARQMVHUO")},
				"manifest": &dynamodb.AttributeValue{S: aws.String("web:\n  image: httpd\n  ports:\n  - 80:80\n")},
				"release":  &dynamodb.AttributeValue{S: aws.String("RFVZFLKVTYO")},
				"status":   &dynamodb.AttributeValue{S: aws.String("complete")},
				"created":  &dynamodb.AttributeValue{S: aws.String(time.Unix(1459709087, 472025215).UTC().Format(sortableTime))},
				"ended":    &dynamodb.AttributeValue{S: aws.String(time.Unix(1459709198, 984281955).UTC().Format(sortableTime))},
			},
		},
	}, nil)

	provider.CloudFormation = cfMock
	provider.DynamoDB = dynamoMock
	b, err := provider.BuildList("httpd", 20)

	assert.Nil(t, err)
	assert.EqualValues(t, structs.Builds{
		structs.Build{
			Id:       "BHINCLZYYVN",
			App:      "httpd",
			Logs:     "",
			Manifest: "web:\n  image: httpd\n  ports:\n  - 80:80\n",
			Release:  "RVFETUHHKKD",
			Status:   "complete",
			Started:  time.Unix(1459780456, 178278576).UTC(),
			Ended:    time.Unix(1459780542, 440881687).UTC(),
		},
		structs.Build{
			Id:       "BNOARQMVHUO",
			App:      "httpd",
			Logs:     "",
			Manifest: "web:\n  image: httpd\n  ports:\n  - 80:80\n",
			Release:  "RFVZFLKVTYO",
			Status:   "complete",
			Started:  time.Unix(1459709087, 472025215).UTC(),
			Ended:    time.Unix(1459709198, 984281955).UTC(),
		},
	}, b)
}

func TestBuildLogs(t *testing.T) {
	provider := StubAwsProvider()
	defer provider.Close()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	cfMock := createAppStackMock(mockCtrl, provider.Rack)

	body := ioutil.NopCloser(strings.NewReader("RUNNING: docker pull httpd"))
	s3Mock := mocks.NewMockS3API(mockCtrl)
	s3Mock.EXPECT().GetObject(&s3.GetObjectInput{
		Bucket: aws.String("convox-httpd-settings-139bidzalmbtu"),
		Key:    aws.String("builds/BHINCLZYYVN.log"),
	}).Return(&s3.GetObjectOutput{
		Body: body,
	}, nil)

	provider.CloudFormation = cfMock
	provider.S3 = s3Mock
	l, err := provider.BuildLogs("httpd", "BHINCLZYYVN")

	assert.Nil(t, err)
	assert.Equal(t, "RUNNING: docker pull httpd", l)
}

func createDynamoGetItemMock(mockCtrl *gomock.Controller, table string) *mocks.MockDynamoDBAPI {

	dynamoMock := mocks.NewMockDynamoDBAPI(mockCtrl)
	dynamoMock.EXPECT().GetItem(&dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		Key: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{S: aws.String("BFAKEID")},
		},
		TableName: aws.String(table),
	}).Return(&dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String("BFAKEID"),
			},
			"app": &dynamodb.AttributeValue{
				S: aws.String("httpd"),
			},
			"manifest": &dynamodb.AttributeValue{
				S: aws.String("version: \"2\"\nservices:\n  web:\n    build: {}\n    command: null\n    image: httpd\n    ports:\n    - 80:80\n"),
			},
			"release": &dynamodb.AttributeValue{
				S: aws.String("RFAKEID"),
			},
			"status": &dynamodb.AttributeValue{
				S: aws.String("complete"),
			},
			"created": &dynamodb.AttributeValue{
				S: aws.String(time.Unix(1459780456, 178278576).UTC().Format("20060102.150405.000000000")),
			},
			"ended": &dynamodb.AttributeValue{
				S: aws.String(time.Unix(1459780542, 440881687).UTC().Format("20060102.150405.000000000")),
			},
		},
	}, nil)

	return dynamoMock
}

func createGetDeleteBuildMock(mockCtrl *gomock.Controller) *mocks.MockDynamoDBAPI {

	dynamoMock := createDynamoGetItemMock(mockCtrl, "convox-builds")
	dynamoMock.EXPECT().DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{S: aws.String("BFAKEID")},
		},
		TableName: aws.String("convox-builds"),
	}).Return(nil, nil)

	return dynamoMock
}

// TODO: this should go somewhere else
func createECRMock(mockCtrl *gomock.Controller) *mocks.MockECRAPI {

	ecrMock := mocks.NewMockECRAPI(mockCtrl)
	ecrMock.EXPECT().BatchDeleteImage(&ecr.BatchDeleteImageInput{
		ImageIds: []*ecr.ImageIdentifier{
			&ecr.ImageIdentifier{
				ImageTag: aws.String("web.BFAKEID"),
			},
		},
		RegistryId:     aws.String("132866487567"),
		RepositoryName: aws.String("convox-httpd-hqvvfosgxt"),
	}).Return(nil, nil)

	return ecrMock
}
