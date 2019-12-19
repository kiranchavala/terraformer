// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteStackInstancesInput struct {
	_ struct{} `type:"structure"`

	// The names of the AWS accounts that you want to delete stack instances for.
	//
	// Accounts is a required field
	Accounts []string `type:"list" required:"true"`

	// The unique identifier for this stack set operation.
	//
	// If you don't specify an operation ID, the SDK generates one automatically.
	//
	// The operation ID also functions as an idempotency token, to ensure that AWS
	// CloudFormation performs the stack set operation only once, even if you retry
	// the request multiple times. You can retry stack set operation requests to
	// ensure that AWS CloudFormation successfully received them.
	//
	// Repeating this stack set operation with a new operation ID retries all stack
	// instances whose status is OUTDATED.
	OperationId *string `min:"1" type:"string" idempotencyToken:"true"`

	// Preferences for how AWS CloudFormation performs this stack set operation.
	OperationPreferences *StackSetOperationPreferences `type:"structure"`

	// The regions where you want to delete stack set instances.
	//
	// Regions is a required field
	Regions []string `type:"list" required:"true"`

	// Removes the stack instances from the specified stack set, but doesn't delete
	// the stacks. You can't reassociate a retained stack or add an existing, saved
	// stack to a new stack set.
	//
	// For more information, see Stack set operation options (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html#stackset-ops-options).
	//
	// RetainStacks is a required field
	RetainStacks *bool `type:"boolean" required:"true"`

	// The name or unique ID of the stack set that you want to delete stack instances
	// for.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteStackInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteStackInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteStackInstancesInput"}

	if s.Accounts == nil {
		invalidParams.Add(aws.NewErrParamRequired("Accounts"))
	}
	if s.OperationId != nil && len(*s.OperationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OperationId", 1))
	}

	if s.Regions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Regions"))
	}

	if s.RetainStacks == nil {
		invalidParams.Add(aws.NewErrParamRequired("RetainStacks"))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}
	if s.OperationPreferences != nil {
		if err := s.OperationPreferences.Validate(); err != nil {
			invalidParams.AddNested("OperationPreferences", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteStackInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for this stack set operation.
	OperationId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteStackInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteStackInstances = "DeleteStackInstances"

// DeleteStackInstancesRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Deletes stack instances for the specified accounts, in the specified regions.
//
//    // Example sending a request using DeleteStackInstancesRequest.
//    req := client.DeleteStackInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DeleteStackInstances
func (c *Client) DeleteStackInstancesRequest(input *DeleteStackInstancesInput) DeleteStackInstancesRequest {
	op := &aws.Operation{
		Name:       opDeleteStackInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteStackInstancesInput{}
	}

	req := c.newRequest(op, input, &DeleteStackInstancesOutput{})
	return DeleteStackInstancesRequest{Request: req, Input: input, Copy: c.DeleteStackInstancesRequest}
}

// DeleteStackInstancesRequest is the request type for the
// DeleteStackInstances API operation.
type DeleteStackInstancesRequest struct {
	*aws.Request
	Input *DeleteStackInstancesInput
	Copy  func(*DeleteStackInstancesInput) DeleteStackInstancesRequest
}

// Send marshals and sends the DeleteStackInstances API request.
func (r DeleteStackInstancesRequest) Send(ctx context.Context) (*DeleteStackInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteStackInstancesResponse{
		DeleteStackInstancesOutput: r.Request.Data.(*DeleteStackInstancesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteStackInstancesResponse is the response type for the
// DeleteStackInstances API operation.
type DeleteStackInstancesResponse struct {
	*DeleteStackInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteStackInstances request.
func (r *DeleteStackInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
