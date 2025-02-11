package mapvalidator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// validateMap ensures that the request contains a Map value.
func validateMap(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) (map[string]attr.Value, bool) {
	var m types.Map

	diags := tfsdk.ValueAs(ctx, request.AttributeConfig, &m)

	if diags.HasError() {
		response.Diagnostics.Append(diags...)

		return nil, false
	}

	if m.IsUnknown() || m.IsNull() {
		return nil, false
	}

	return m.Elements(), true
}
