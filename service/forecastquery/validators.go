// Code generated by smithy-go-codegen DO NOT EDIT.

package forecastquery

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpQueryForecast struct {
}

func (*validateOpQueryForecast) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpQueryForecast) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*QueryForecastInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpQueryForecastInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpQueryWhatIfForecast struct {
}

func (*validateOpQueryWhatIfForecast) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpQueryWhatIfForecast) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*QueryWhatIfForecastInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpQueryWhatIfForecastInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpQueryForecastValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpQueryForecast{}, middleware.After)
}

func addOpQueryWhatIfForecastValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpQueryWhatIfForecast{}, middleware.After)
}

func validateOpQueryForecastInput(v *QueryForecastInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueryForecastInput"}
	if v.ForecastArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ForecastArn"))
	}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpQueryWhatIfForecastInput(v *QueryWhatIfForecastInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueryWhatIfForecastInput"}
	if v.WhatIfForecastArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WhatIfForecastArn"))
	}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
