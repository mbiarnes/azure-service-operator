// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220131previewstorage

import (
	"encoding/json"
	v20220131ps "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20220131previewstorage"
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

func Test_FederatedIdentityCredential_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FederatedIdentityCredential to hub returns original",
		prop.ForAll(RunResourceConversionTestForFederatedIdentityCredential, FederatedIdentityCredentialGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFederatedIdentityCredential tests if a specific instance of FederatedIdentityCredential round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFederatedIdentityCredential(subject FederatedIdentityCredential) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220131ps.FederatedIdentityCredential
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FederatedIdentityCredential
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FederatedIdentityCredential_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FederatedIdentityCredential to FederatedIdentityCredential via AssignProperties_To_FederatedIdentityCredential & AssignProperties_From_FederatedIdentityCredential returns original",
		prop.ForAll(RunPropertyAssignmentTestForFederatedIdentityCredential, FederatedIdentityCredentialGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFederatedIdentityCredential tests if a specific instance of FederatedIdentityCredential can be assigned to v1api20220131previewstorage and back losslessly
func RunPropertyAssignmentTestForFederatedIdentityCredential(subject FederatedIdentityCredential) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220131ps.FederatedIdentityCredential
	err := copied.AssignProperties_To_FederatedIdentityCredential(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FederatedIdentityCredential
	err = actual.AssignProperties_From_FederatedIdentityCredential(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_FederatedIdentityCredential_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredential via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredential, FederatedIdentityCredentialGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredential runs a test to see if a specific instance of FederatedIdentityCredential round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredential(subject FederatedIdentityCredential) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredential
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

// Generator of FederatedIdentityCredential instances for property testing - lazily instantiated by
// FederatedIdentityCredentialGenerator()
var federatedIdentityCredentialGenerator gopter.Gen

// FederatedIdentityCredentialGenerator returns a generator of FederatedIdentityCredential instances for property testing.
func FederatedIdentityCredentialGenerator() gopter.Gen {
	if federatedIdentityCredentialGenerator != nil {
		return federatedIdentityCredentialGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFederatedIdentityCredential(generators)
	federatedIdentityCredentialGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential{}), generators)

	return federatedIdentityCredentialGenerator
}

// AddRelatedPropertyGeneratorsForFederatedIdentityCredential is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFederatedIdentityCredential(gens map[string]gopter.Gen) {
	gens["Spec"] = UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator()
	gens["Status"] = UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator()
}

func Test_UserAssignedIdentities_FederatedIdentityCredential_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentities_FederatedIdentityCredential_Spec to UserAssignedIdentities_FederatedIdentityCredential_Spec via AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec & AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentities_FederatedIdentityCredential_Spec, UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentities_FederatedIdentityCredential_Spec tests if a specific instance of UserAssignedIdentities_FederatedIdentityCredential_Spec can be assigned to v1api20220131previewstorage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentities_FederatedIdentityCredential_Spec(subject UserAssignedIdentities_FederatedIdentityCredential_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220131ps.UserAssignedIdentities_FederatedIdentityCredential_Spec
	err := copied.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentities_FederatedIdentityCredential_Spec
	err = actual.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_UserAssignedIdentities_FederatedIdentityCredential_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_FederatedIdentityCredential_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_Spec, UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_Spec runs a test to see if a specific instance of UserAssignedIdentities_FederatedIdentityCredential_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_Spec(subject UserAssignedIdentities_FederatedIdentityCredential_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_FederatedIdentityCredential_Spec
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

// Generator of UserAssignedIdentities_FederatedIdentityCredential_Spec instances for property testing - lazily
// instantiated by UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator()
var userAssignedIdentities_FederatedIdentityCredential_SpecGenerator gopter.Gen

// UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator returns a generator of UserAssignedIdentities_FederatedIdentityCredential_Spec instances for property testing.
func UserAssignedIdentities_FederatedIdentityCredential_SpecGenerator() gopter.Gen {
	if userAssignedIdentities_FederatedIdentityCredential_SpecGenerator != nil {
		return userAssignedIdentities_FederatedIdentityCredential_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_Spec(generators)
	userAssignedIdentities_FederatedIdentityCredential_SpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_FederatedIdentityCredential_Spec{}), generators)

	return userAssignedIdentities_FederatedIdentityCredential_SpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_Spec(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentities_FederatedIdentityCredential_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from UserAssignedIdentities_FederatedIdentityCredential_STATUS to UserAssignedIdentities_FederatedIdentityCredential_STATUS via AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS & AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS, UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS tests if a specific instance of UserAssignedIdentities_FederatedIdentityCredential_STATUS can be assigned to v1api20220131previewstorage and back losslessly
func RunPropertyAssignmentTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS(subject UserAssignedIdentities_FederatedIdentityCredential_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220131ps.UserAssignedIdentities_FederatedIdentityCredential_STATUS
	err := copied.AssignProperties_To_UserAssignedIdentities_FederatedIdentityCredential_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual UserAssignedIdentities_FederatedIdentityCredential_STATUS
	err = actual.AssignProperties_From_UserAssignedIdentities_FederatedIdentityCredential_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_UserAssignedIdentities_FederatedIdentityCredential_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_FederatedIdentityCredential_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS, UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS runs a test to see if a specific instance of UserAssignedIdentities_FederatedIdentityCredential_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentities_FederatedIdentityCredential_STATUS(subject UserAssignedIdentities_FederatedIdentityCredential_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_FederatedIdentityCredential_STATUS
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

// Generator of UserAssignedIdentities_FederatedIdentityCredential_STATUS instances for property testing - lazily
// instantiated by UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator()
var userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator gopter.Gen

// UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator returns a generator of UserAssignedIdentities_FederatedIdentityCredential_STATUS instances for property testing.
func UserAssignedIdentities_FederatedIdentityCredential_STATUSGenerator() gopter.Gen {
	if userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator != nil {
		return userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS(generators)
	userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_FederatedIdentityCredential_STATUS{}), generators)

	return userAssignedIdentities_FederatedIdentityCredential_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentities_FederatedIdentityCredential_STATUS(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}
