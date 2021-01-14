// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a fleet.
func (c *Client) GetDeviceFleetReport(ctx context.Context, params *GetDeviceFleetReportInput, optFns ...func(*Options)) (*GetDeviceFleetReportOutput, error) {
	if params == nil {
		params = &GetDeviceFleetReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeviceFleetReport", params, optFns, addOperationGetDeviceFleetReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeviceFleetReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeviceFleetReportInput struct {

	// The name of the fleet.
	//
	// This member is required.
	DeviceFleetName *string
}

type GetDeviceFleetReportOutput struct {

	// The Amazon Resource Name (ARN) of the device.
	//
	// This member is required.
	DeviceFleetArn *string

	// The name of the fleet.
	//
	// This member is required.
	DeviceFleetName *string

	// The versions of Edge Manager agent deployed on the fleet.
	AgentVersions []types.AgentVersion

	// Description of the fleet.
	Description *string

	// Status of devices.
	DeviceStats *types.DeviceStats

	// Status of model on device.
	ModelStats []types.EdgeModelStat

	// The output configuration for storing sample data collected by the fleet.
	OutputConfig *types.EdgeOutputConfig

	// Timestamp of when the report was generated.
	ReportGenerated *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDeviceFleetReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeviceFleetReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeviceFleetReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetDeviceFleetReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeviceFleetReport(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetDeviceFleetReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "GetDeviceFleetReport",
	}
}