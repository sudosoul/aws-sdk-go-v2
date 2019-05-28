// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeEntityRecognizerRequest
type DescribeEntityRecognizerInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that identifies the entity recognizer.
	//
	// EntityRecognizerArn is a required field
	EntityRecognizerArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEntityRecognizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEntityRecognizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEntityRecognizerInput"}

	if s.EntityRecognizerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityRecognizerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeEntityRecognizerResponse
type DescribeEntityRecognizerOutput struct {
	_ struct{} `type:"structure"`

	// Describes information associated with an entity recognizer.
	EntityRecognizerProperties *EntityRecognizerProperties `type:"structure"`
}

// String returns the string representation
func (s DescribeEntityRecognizerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEntityRecognizer = "DescribeEntityRecognizer"

// DescribeEntityRecognizerRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Provides details about an entity recognizer including status, S3 buckets
// containing training data, recognizer metadata, metrics, and so on.
//
//    // Example sending a request using DescribeEntityRecognizerRequest.
//    req := client.DescribeEntityRecognizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/DescribeEntityRecognizer
func (c *Client) DescribeEntityRecognizerRequest(input *DescribeEntityRecognizerInput) DescribeEntityRecognizerRequest {
	op := &aws.Operation{
		Name:       opDescribeEntityRecognizer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEntityRecognizerInput{}
	}

	req := c.newRequest(op, input, &DescribeEntityRecognizerOutput{})
	return DescribeEntityRecognizerRequest{Request: req, Input: input, Copy: c.DescribeEntityRecognizerRequest}
}

// DescribeEntityRecognizerRequest is the request type for the
// DescribeEntityRecognizer API operation.
type DescribeEntityRecognizerRequest struct {
	*aws.Request
	Input *DescribeEntityRecognizerInput
	Copy  func(*DescribeEntityRecognizerInput) DescribeEntityRecognizerRequest
}

// Send marshals and sends the DescribeEntityRecognizer API request.
func (r DescribeEntityRecognizerRequest) Send(ctx context.Context) (*DescribeEntityRecognizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEntityRecognizerResponse{
		DescribeEntityRecognizerOutput: r.Request.Data.(*DescribeEntityRecognizerOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEntityRecognizerResponse is the response type for the
// DescribeEntityRecognizer API operation.
type DescribeEntityRecognizerResponse struct {
	*DescribeEntityRecognizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEntityRecognizer request.
func (r *DescribeEntityRecognizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}