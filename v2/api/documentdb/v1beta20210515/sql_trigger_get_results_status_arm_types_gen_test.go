// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_SqlTriggerGetResults_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetResults_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetResults_STATUS_ARM, SqlTriggerGetResults_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetResults_STATUS_ARM runs a test to see if a specific instance of SqlTriggerGetResults_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetResults_STATUS_ARM(subject SqlTriggerGetResults_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetResults_STATUS_ARM
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

// Generator of SqlTriggerGetResults_STATUS_ARM instances for property testing - lazily instantiated by
// SqlTriggerGetResults_STATUS_ARMGenerator()
var sqlTriggerGetResults_STATUS_ARMGenerator gopter.Gen

// SqlTriggerGetResults_STATUS_ARMGenerator returns a generator of SqlTriggerGetResults_STATUS_ARM instances for property testing.
// We first initialize sqlTriggerGetResults_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlTriggerGetResults_STATUS_ARMGenerator() gopter.Gen {
	if sqlTriggerGetResults_STATUS_ARMGenerator != nil {
		return sqlTriggerGetResults_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResults_STATUS_ARM(generators)
	sqlTriggerGetResults_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResults_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSqlTriggerGetResults_STATUS_ARM(generators)
	sqlTriggerGetResults_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_STATUS_ARM{}), generators)

	return sqlTriggerGetResults_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetResults_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetResults_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlTriggerGetResults_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerGetResults_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlTriggerGetProperties_STATUS_ARMGenerator())
}

func Test_SqlTriggerGetProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetProperties_STATUS_ARM, SqlTriggerGetProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetProperties_STATUS_ARM runs a test to see if a specific instance of SqlTriggerGetProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetProperties_STATUS_ARM(subject SqlTriggerGetProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_STATUS_ARM
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

// Generator of SqlTriggerGetProperties_STATUS_ARM instances for property testing - lazily instantiated by
// SqlTriggerGetProperties_STATUS_ARMGenerator()
var sqlTriggerGetProperties_STATUS_ARMGenerator gopter.Gen

// SqlTriggerGetProperties_STATUS_ARMGenerator returns a generator of SqlTriggerGetProperties_STATUS_ARM instances for property testing.
func SqlTriggerGetProperties_STATUS_ARMGenerator() gopter.Gen {
	if sqlTriggerGetProperties_STATUS_ARMGenerator != nil {
		return sqlTriggerGetProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlTriggerGetProperties_STATUS_ARM(generators)
	sqlTriggerGetProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_STATUS_ARM{}), generators)

	return sqlTriggerGetProperties_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlTriggerGetProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerGetProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlTriggerGetProperties_Resource_STATUS_ARMGenerator())
}

func Test_SqlTriggerGetProperties_Resource_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_Resource_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetProperties_Resource_STATUS_ARM, SqlTriggerGetProperties_Resource_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetProperties_Resource_STATUS_ARM runs a test to see if a specific instance of SqlTriggerGetProperties_Resource_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetProperties_Resource_STATUS_ARM(subject SqlTriggerGetProperties_Resource_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_Resource_STATUS_ARM
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

// Generator of SqlTriggerGetProperties_Resource_STATUS_ARM instances for property testing - lazily instantiated by
// SqlTriggerGetProperties_Resource_STATUS_ARMGenerator()
var sqlTriggerGetProperties_Resource_STATUS_ARMGenerator gopter.Gen

// SqlTriggerGetProperties_Resource_STATUS_ARMGenerator returns a generator of SqlTriggerGetProperties_Resource_STATUS_ARM instances for property testing.
func SqlTriggerGetProperties_Resource_STATUS_ARMGenerator() gopter.Gen {
	if sqlTriggerGetProperties_Resource_STATUS_ARMGenerator != nil {
		return sqlTriggerGetProperties_Resource_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_Resource_STATUS_ARM(generators)
	sqlTriggerGetProperties_Resource_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_Resource_STATUS_ARM{}), generators)

	return sqlTriggerGetProperties_Resource_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_Resource_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetProperties_Resource_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.OneConstOf(
		SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_All,
		SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Create,
		SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Delete,
		SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Replace,
		SqlTriggerGetProperties_Resource_TriggerOperation_STATUS_Update))
	gens["TriggerType"] = gen.PtrOf(gen.OneConstOf(SqlTriggerGetProperties_Resource_TriggerType_STATUS_Post, SqlTriggerGetProperties_Resource_TriggerType_STATUS_Pre))
	gens["Ts"] = gen.PtrOf(gen.Float64())
}
