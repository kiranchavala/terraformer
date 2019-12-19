// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTimeToLiveInput struct {
	_ struct{} `type:"structure"`

	// The name of the table to be described.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTimeToLiveInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTimeToLiveInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTimeToLiveInput"}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTimeToLiveOutput struct {
	_ struct{} `type:"structure"`

	// The description of the Time to Live (TTL) status on the specified table.
	TimeToLiveDescription *TimeToLiveDescription `type:"structure"`
}

// String returns the string representation
func (s DescribeTimeToLiveOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTimeToLive = "DescribeTimeToLive"

// DescribeTimeToLiveRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Gives a description of the Time to Live (TTL) status on the specified table.
//
//    // Example sending a request using DescribeTimeToLiveRequest.
//    req := client.DescribeTimeToLiveRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DescribeTimeToLive
func (c *Client) DescribeTimeToLiveRequest(input *DescribeTimeToLiveInput) DescribeTimeToLiveRequest {
	op := &aws.Operation{
		Name:       opDescribeTimeToLive,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTimeToLiveInput{}
	}

	req := c.newRequest(op, input, &DescribeTimeToLiveOutput{})
	return DescribeTimeToLiveRequest{Request: req, Input: input, Copy: c.DescribeTimeToLiveRequest}
}

// DescribeTimeToLiveRequest is the request type for the
// DescribeTimeToLive API operation.
type DescribeTimeToLiveRequest struct {
	*aws.Request
	Input *DescribeTimeToLiveInput
	Copy  func(*DescribeTimeToLiveInput) DescribeTimeToLiveRequest
}

// Send marshals and sends the DescribeTimeToLive API request.
func (r DescribeTimeToLiveRequest) Send(ctx context.Context) (*DescribeTimeToLiveResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTimeToLiveResponse{
		DescribeTimeToLiveOutput: r.Request.Data.(*DescribeTimeToLiveOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTimeToLiveResponse is the response type for the
// DescribeTimeToLive API operation.
type DescribeTimeToLiveResponse struct {
	*DescribeTimeToLiveOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTimeToLive request.
func (r *DescribeTimeToLiveResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
