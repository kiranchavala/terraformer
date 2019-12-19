// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccessKeyLastUsedInput struct {
	_ struct{} `type:"structure"`

	// The identifier of an access key.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// AccessKeyId is a required field
	AccessKeyId *string `min:"16" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAccessKeyLastUsedInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccessKeyLastUsedInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccessKeyLastUsedInput"}

	if s.AccessKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessKeyId"))
	}
	if s.AccessKeyId != nil && len(*s.AccessKeyId) < 16 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessKeyId", 16))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetAccessKeyLastUsed request. It is
// also returned as a member of the AccessKeyMetaData structure returned by
// the ListAccessKeys action.
type GetAccessKeyLastUsedOutput struct {
	_ struct{} `type:"structure"`

	// Contains information about the last time the access key was used.
	AccessKeyLastUsed *AccessKeyLastUsed `type:"structure"`

	// The name of the AWS IAM user that owns this access key.
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetAccessKeyLastUsedOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAccessKeyLastUsed = "GetAccessKeyLastUsed"

// GetAccessKeyLastUsedRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves information about when the specified access key was last used.
// The information includes the date and time of last use, along with the AWS
// service and Region that were specified in the last request made with that
// key.
//
//    // Example sending a request using GetAccessKeyLastUsedRequest.
//    req := client.GetAccessKeyLastUsedRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetAccessKeyLastUsed
func (c *Client) GetAccessKeyLastUsedRequest(input *GetAccessKeyLastUsedInput) GetAccessKeyLastUsedRequest {
	op := &aws.Operation{
		Name:       opGetAccessKeyLastUsed,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAccessKeyLastUsedInput{}
	}

	req := c.newRequest(op, input, &GetAccessKeyLastUsedOutput{})
	return GetAccessKeyLastUsedRequest{Request: req, Input: input, Copy: c.GetAccessKeyLastUsedRequest}
}

// GetAccessKeyLastUsedRequest is the request type for the
// GetAccessKeyLastUsed API operation.
type GetAccessKeyLastUsedRequest struct {
	*aws.Request
	Input *GetAccessKeyLastUsedInput
	Copy  func(*GetAccessKeyLastUsedInput) GetAccessKeyLastUsedRequest
}

// Send marshals and sends the GetAccessKeyLastUsed API request.
func (r GetAccessKeyLastUsedRequest) Send(ctx context.Context) (*GetAccessKeyLastUsedResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAccessKeyLastUsedResponse{
		GetAccessKeyLastUsedOutput: r.Request.Data.(*GetAccessKeyLastUsedOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAccessKeyLastUsedResponse is the response type for the
// GetAccessKeyLastUsed API operation.
type GetAccessKeyLastUsedResponse struct {
	*GetAccessKeyLastUsedOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAccessKeyLastUsed request.
func (r *GetAccessKeyLastUsedResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
