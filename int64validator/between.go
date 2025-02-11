package int64validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
)

var _ tfsdk.AttributeValidator = betweenValidator{}

// betweenValidator validates that an integer Attribute's value is in a range.
type betweenValidator struct {
	min, max int64
}

// Description describes the validation in plain text formatting.
func (validator betweenValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be between %d and %d", validator.min, validator.max)
}

// MarkdownDescription describes the validation in Markdown formatting.
func (validator betweenValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// Validate performs the validation.
func (validator betweenValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	i, ok := validateInt(ctx, request, response)

	if !ok {
		return
	}

	if i < validator.min || i > validator.max {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.AttributePath,
			validator.Description(ctx),
			fmt.Sprintf("%d", i),
		))

		return
	}
}

// Between returns an AttributeValidator which ensures that any configured
// attribute value:
//
//   - Is a number, which can be represented by a 64-bit integer.
//   - Is greater than or equal to the given minimum and less than or equal to the given maximum.
//
// Null (unconfigured) and unknown (known after apply) values are skipped.
func Between(min, max int64) tfsdk.AttributeValidator {
	if min > max {
		return nil
	}

	return betweenValidator{
		min: min,
		max: max,
	}
}
