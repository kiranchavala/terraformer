// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package firehose

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TagDeliveryStreamInput struct {
	_ struct{} `type:"structure"`

	// The name of the delivery stream to which you want to add the tags.
	//
	// DeliveryStreamName is a required field
	DeliveryStreamName *string `min:"1" type:"string" required:"true"`

	// A set of key-value pairs to use to create the tags.
	//
	// Tags is a required field
	Tags []Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s TagDeliveryStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagDeliveryStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagDeliveryStreamInput"}

	if s.DeliveryStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeliveryStreamName"))
	}
	if s.DeliveryStreamName != nil && len(*s.DeliveryStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeliveryStreamName", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagDeliveryStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagDeliveryStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagDeliveryStream = "TagDeliveryStream"

// TagDeliveryStreamRequest returns a request value for making API operation for
// Amazon Kinesis Firehose.
//
// Adds or updates tags for the specified delivery stream. A tag is a key-value
// pair that you can define and assign to AWS resources. If you specify a tag
// that already exists, the tag value is replaced with the value that you specify
// in the request. Tags are metadata. For example, you can add friendly names
// and descriptions or other types of information that can help you distinguish
// the delivery stream. For more information about tags, see Using Cost Allocation
// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide.
//
// Each delivery stream can have up to 50 tags.
//
// This operation has a limit of five transactions per second per account.
//
//    // Example sending a request using TagDeliveryStreamRequest.
//    req := client.TagDeliveryStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/firehose-2015-08-04/TagDeliveryStream
func (c *Client) TagDeliveryStreamRequest(input *TagDeliveryStreamInput) TagDeliveryStreamRequest {
	op := &aws.Operation{
		Name:       opTagDeliveryStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagDeliveryStreamInput{}
	}

	req := c.newRequest(op, input, &TagDeliveryStreamOutput{})
	return TagDeliveryStreamRequest{Request: req, Input: input, Copy: c.TagDeliveryStreamRequest}
}

// TagDeliveryStreamRequest is the request type for the
// TagDeliveryStream API operation.
type TagDeliveryStreamRequest struct {
	*aws.Request
	Input *TagDeliveryStreamInput
	Copy  func(*TagDeliveryStreamInput) TagDeliveryStreamRequest
}

// Send marshals and sends the TagDeliveryStream API request.
func (r TagDeliveryStreamRequest) Send(ctx context.Context) (*TagDeliveryStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagDeliveryStreamResponse{
		TagDeliveryStreamOutput: r.Request.Data.(*TagDeliveryStreamOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagDeliveryStreamResponse is the response type for the
// TagDeliveryStream API operation.
type TagDeliveryStreamResponse struct {
	*TagDeliveryStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagDeliveryStream request.
func (r *TagDeliveryStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
