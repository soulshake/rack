package aws_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	provider "github.com/convox/rack/provider/aws"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("RACK", "convox-test")
}

func TestGenerateStackInput(t *testing.T) {
	// Expected
	expectedParameters := []*cloudformation.Parameter{
		&cloudformation.Parameter{
			ParameterKey:   aws.String("ChangingA"),
			ParameterValue: aws.String("new-a"),
		},
		&cloudformation.Parameter{
			ParameterKey:   aws.String("ChangingB"),
			ParameterValue: aws.String("new-b"),
		},
		&cloudformation.Parameter{
			ParameterKey:     aws.String("NotChanging"),
			UsePreviousValue: aws.Bool(true),
		},
		&cloudformation.Parameter{
			ParameterKey:     aws.String("NotChangingNoEcho"),
			UsePreviousValue: aws.Bool(true),
		},
	}

	server := serverReturningFixture(t, "new-formation-template.json")
	defer server.Close()

	expectedStackInput := &cloudformation.UpdateStackInput{
		Capabilities: []*string{aws.String("CAPABILITY_IAM")},
		Parameters:   expectedParameters,
		StackName:    aws.String("test"),
		TemplateURL:  aws.String(server.URL),
	}

	// Actual
	stackName := "test"
	templateURL := server.URL
	currentParams := map[string]string{
		"ChangingA":         "current-a",
		"ChangingB":         "current-b",
		"NotChanging":       "not-changing",
		"NotChangingNoEcho": "****",
	}
	changes := map[string]string{
		"ChangingA": "new-a",
		"ChangingB": "new-b",
	}

	actualStackInput, _ := provider.GenerateStackInput(stackName, templateURL, currentParams, changes)

	assert.Equal(t, expectedStackInput, actualStackInput)
}

func serverReturningFixture(t *testing.T, fixtureName string) *httptest.Server {
	contents, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s", fixtureName))
	if err != nil {
		t.Error(err.Error())
	}

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(contents)
	}))
}
