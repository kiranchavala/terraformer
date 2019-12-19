// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ReplaceRouteTableAssociationInput struct {
	_ struct{} `type:"structure"`

	// The association ID.
	//
	// AssociationId is a required field
	AssociationId *string `locationName:"associationId" type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the new route table to associate with the subnet.
	//
	// RouteTableId is a required field
	RouteTableId *string `locationName:"routeTableId" type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceRouteTableAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceRouteTableAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReplaceRouteTableAssociationInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}

	if s.RouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ReplaceRouteTableAssociationOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the new association.
	NewAssociationId *string `locationName:"newAssociationId" type:"string"`
}

// String returns the string representation
func (s ReplaceRouteTableAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

const opReplaceRouteTableAssociation = "ReplaceRouteTableAssociation"

// ReplaceRouteTableAssociationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Changes the route table associated with a given subnet in a VPC. After the
// operation completes, the subnet uses the routes in the new route table it's
// associated with. For more information about route tables, see Route Tables
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
//
// You can also use ReplaceRouteTableAssociation to change which table is the
// main route table in the VPC. You just specify the main route table's association
// ID and the route table to be the new main route table.
//
//    // Example sending a request using ReplaceRouteTableAssociationRequest.
//    req := client.ReplaceRouteTableAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReplaceRouteTableAssociation
func (c *Client) ReplaceRouteTableAssociationRequest(input *ReplaceRouteTableAssociationInput) ReplaceRouteTableAssociationRequest {
	op := &aws.Operation{
		Name:       opReplaceRouteTableAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReplaceRouteTableAssociationInput{}
	}

	req := c.newRequest(op, input, &ReplaceRouteTableAssociationOutput{})
	return ReplaceRouteTableAssociationRequest{Request: req, Input: input, Copy: c.ReplaceRouteTableAssociationRequest}
}

// ReplaceRouteTableAssociationRequest is the request type for the
// ReplaceRouteTableAssociation API operation.
type ReplaceRouteTableAssociationRequest struct {
	*aws.Request
	Input *ReplaceRouteTableAssociationInput
	Copy  func(*ReplaceRouteTableAssociationInput) ReplaceRouteTableAssociationRequest
}

// Send marshals and sends the ReplaceRouteTableAssociation API request.
func (r ReplaceRouteTableAssociationRequest) Send(ctx context.Context) (*ReplaceRouteTableAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReplaceRouteTableAssociationResponse{
		ReplaceRouteTableAssociationOutput: r.Request.Data.(*ReplaceRouteTableAssociationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReplaceRouteTableAssociationResponse is the response type for the
// ReplaceRouteTableAssociation API operation.
type ReplaceRouteTableAssociationResponse struct {
	*ReplaceRouteTableAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReplaceRouteTableAssociation request.
func (r *ReplaceRouteTableAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
