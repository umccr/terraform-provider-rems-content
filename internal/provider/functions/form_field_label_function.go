// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = FormFieldLabelFunction{}
)

func NewFormFieldLabelFunction() function.Function {
	return FormFieldLabelFunction{}
}

type FormFieldLabelFunction struct{}

func (r FormFieldLabelFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "form_field_label"
}

func (r FormFieldLabelFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Field template for a label",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name: "title",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"title":    types.StringType,
				"type":     types.StringType,
				"optional": types.BoolType,
			},
		},
	}
}

func (r FormFieldLabelFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var functionData string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &functionData))

	if resp.Error != nil {
		return
	}

	result := struct {
		Title    string `tfsdk:"title"`
		Type     string `tfsdk:"type"`
		Optional bool   `tfsdk:"optional"`
	}{
		Title: functionData,
		Type:  "label",
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
