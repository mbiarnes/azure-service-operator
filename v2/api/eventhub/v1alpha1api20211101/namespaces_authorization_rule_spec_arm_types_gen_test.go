// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Namespaces_AuthorizationRule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_AuthorizationRule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_AuthorizationRule_Spec_ARM, Namespaces_AuthorizationRule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_AuthorizationRule_Spec_ARM runs a test to see if a specific instance of Namespaces_AuthorizationRule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_AuthorizationRule_Spec_ARM(subject Namespaces_AuthorizationRule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_AuthorizationRule_Spec_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Namespaces_AuthorizationRule_Spec_ARM instances for property testing - lazily instantiated by
// Namespaces_AuthorizationRule_Spec_ARMGenerator()
var namespaces_AuthorizationRule_Spec_ARMGenerator gopter.Gen

// Namespaces_AuthorizationRule_Spec_ARMGenerator returns a generator of Namespaces_AuthorizationRule_Spec_ARM instances for property testing.
// We first initialize namespaces_AuthorizationRule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_AuthorizationRule_Spec_ARMGenerator() gopter.Gen {
	if namespaces_AuthorizationRule_Spec_ARMGenerator != nil {
		return namespaces_AuthorizationRule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Spec_ARM(generators)
	namespaces_AuthorizationRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_AuthorizationRule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_AuthorizationRule_Spec_ARM(generators)
	namespaces_AuthorizationRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_AuthorizationRule_Spec_ARM{}), generators)

	return namespaces_AuthorizationRule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_AuthorizationRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_AuthorizationRule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_AuthorizationRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(AuthorizationRuleProperties_ARMGenerator())
}

func Test_AuthorizationRuleProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationRuleProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationRuleProperties_ARM, AuthorizationRuleProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationRuleProperties_ARM runs a test to see if a specific instance of AuthorizationRuleProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationRuleProperties_ARM(subject AuthorizationRuleProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationRuleProperties_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AuthorizationRuleProperties_ARM instances for property testing - lazily instantiated by
// AuthorizationRuleProperties_ARMGenerator()
var authorizationRuleProperties_ARMGenerator gopter.Gen

// AuthorizationRuleProperties_ARMGenerator returns a generator of AuthorizationRuleProperties_ARM instances for property testing.
func AuthorizationRuleProperties_ARMGenerator() gopter.Gen {
	if authorizationRuleProperties_ARMGenerator != nil {
		return authorizationRuleProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationRuleProperties_ARM(generators)
	authorizationRuleProperties_ARMGenerator = gen.Struct(reflect.TypeOf(AuthorizationRuleProperties_ARM{}), generators)

	return authorizationRuleProperties_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationRuleProperties_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationRuleProperties_ARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(AuthorizationRuleProperties_Rights_Listen, AuthorizationRuleProperties_Rights_Manage, AuthorizationRuleProperties_Rights_Send))
}
