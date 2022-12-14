// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies the extended metrics that you want a CloudWatch RUM app monitor to
// send to a destination. Valid destinations include CloudWatch and Evidently. By
// default, RUM app monitors send some metrics to CloudWatch. These default metrics
// are listed in CloudWatch metrics that you can collect with CloudWatch RUM
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-metrics.html).
// If you also send extended metrics, you can send metrics to Evidently as well as
// CloudWatch, and you can also optionally send the metrics with additional
// dimensions. The valid dimension names for the additional dimensions are
// BrowserName, CountryCode, DeviceType, FileType, OSName, and PageId. For more
// information, see  Extended metrics that you can send to CloudWatch and
// CloudWatch Evidently
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-vended-metrics.html).
// The maximum number of metric definitions that you can specify in one
// BatchCreateRumMetricDefinitions operation is 200. The maximum number of metric
// definitions that one destination can contain is 2000. Extended metrics sent are
// charged as CloudWatch custom metrics. Each combination of additional dimension
// name and dimension value counts as a custom metric. For more information, see
// Amazon CloudWatch Pricing (https://aws.amazon.com/cloudwatch/pricing/). You must
// have already created a destination for the metrics before you send them. For
// more information, see PutRumMetricsDestination
// (https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_PutRumMetricsDestination.html).
// If some metric definitions specified in a BatchCreateRumMetricDefinitions
// operations are not valid, those metric definitions fail and return errors, but
// all valid metric definitions in the same operation still succeed.
func (c *Client) BatchCreateRumMetricDefinitions(ctx context.Context, params *BatchCreateRumMetricDefinitionsInput, optFns ...func(*Options)) (*BatchCreateRumMetricDefinitionsOutput, error) {
	if params == nil {
		params = &BatchCreateRumMetricDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCreateRumMetricDefinitions", params, optFns, c.addOperationBatchCreateRumMetricDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCreateRumMetricDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateRumMetricDefinitionsInput struct {

	// The name of the CloudWatch RUM app monitor that is to send the metrics.
	//
	// This member is required.
	AppMonitorName *string

	// The destination to send the metrics to. Valid values are CloudWatch and
	// Evidently. If you specify Evidently, you must also specify the ARN of the
	// CloudWatchEvidently experiment that will receive the metrics and an IAM role
	// that has permission to write to the experiment.
	//
	// This member is required.
	Destination types.MetricDestination

	// An array of structures which define the metrics that you want to send.
	//
	// This member is required.
	MetricDefinitions []types.MetricDefinitionRequest

	// This parameter is required if Destination is Evidently. If Destination is
	// CloudWatch, do not use this parameter. This parameter specifies the ARN of the
	// Evidently experiment that is to receive the metrics. You must have already
	// defined this experiment as a valid destination. For more information, see
	// PutRumMetricsDestination
	// (https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_PutRumMetricsDestination.html).
	DestinationArn *string

	noSmithyDocumentSerde
}

type BatchCreateRumMetricDefinitionsOutput struct {

	// An array of error objects, if the operation caused any errors.
	//
	// This member is required.
	Errors []types.BatchCreateRumMetricDefinitionsError

	// An array of structures that define the extended metrics.
	MetricDefinitions []types.MetricDefinition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchCreateRumMetricDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchCreateRumMetricDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchCreateRumMetricDefinitions{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchCreateRumMetricDefinitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateRumMetricDefinitions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchCreateRumMetricDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rum",
		OperationName: "BatchCreateRumMetricDefinitions",
	}
}
