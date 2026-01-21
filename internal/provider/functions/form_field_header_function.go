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
	_ function.Function = FormFieldHeaderFunction{}
)

func NewFormFieldHeaderFunction() function.Function {
	return FormFieldHeaderFunction{}
}

type FormFieldHeaderFunction struct{}

func (r FormFieldHeaderFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "form_field_header"
}

func (r FormFieldHeaderFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary: "Field template for a header",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name: "field_id",
			},
			function.MapParameter{
				ElementType: types.StringType,
				Name:        "title",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"id":       types.StringType,
				"title":    types.MapType{ElemType: types.StringType},
				"type":     types.StringType,
				"optional": types.BoolType,
			},
		},
	}
}

func (r FormFieldHeaderFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var idData string
	var titleData map[string]string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &idData, &titleData))

	if resp.Error != nil {
		return
	}

	result := struct {
		Id       string            `tfsdk:"id"`
		Title    map[string]string `tfsdk:"title"`
		Type     string            `tfsdk:"type"`
		Optional bool              `tfsdk:"optional"`
	}{
		Id:    idData,
		Title: titleData,
		Type:  "header",
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, result))
}
