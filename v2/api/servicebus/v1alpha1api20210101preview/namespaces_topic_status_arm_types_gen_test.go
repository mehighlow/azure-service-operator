// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101preview

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

func Test_Namespaces_Topic_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Topic_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Topic_STATUS_ARM, Namespaces_Topic_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Topic_STATUS_ARM runs a test to see if a specific instance of Namespaces_Topic_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Topic_STATUS_ARM(subject Namespaces_Topic_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Topic_STATUS_ARM
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

// Generator of Namespaces_Topic_STATUS_ARM instances for property testing - lazily instantiated by
// Namespaces_Topic_STATUS_ARMGenerator()
var namespaces_Topic_STATUS_ARMGenerator gopter.Gen

// Namespaces_Topic_STATUS_ARMGenerator returns a generator of Namespaces_Topic_STATUS_ARM instances for property testing.
// We first initialize namespaces_Topic_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Topic_STATUS_ARMGenerator() gopter.Gen {
	if namespaces_Topic_STATUS_ARMGenerator != nil {
		return namespaces_Topic_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topic_STATUS_ARM(generators)
	namespaces_Topic_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topic_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Topic_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Topic_STATUS_ARM(generators)
	namespaces_Topic_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Topic_STATUS_ARM{}), generators)

	return namespaces_Topic_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Topic_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Topic_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_Topic_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Topic_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBTopicProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_SBTopicProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBTopicProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBTopicProperties_STATUS_ARM, SBTopicProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBTopicProperties_STATUS_ARM runs a test to see if a specific instance of SBTopicProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSBTopicProperties_STATUS_ARM(subject SBTopicProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBTopicProperties_STATUS_ARM
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

// Generator of SBTopicProperties_STATUS_ARM instances for property testing - lazily instantiated by
// SBTopicProperties_STATUS_ARMGenerator()
var sbTopicProperties_STATUS_ARMGenerator gopter.Gen

// SBTopicProperties_STATUS_ARMGenerator returns a generator of SBTopicProperties_STATUS_ARM instances for property testing.
// We first initialize sbTopicProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SBTopicProperties_STATUS_ARMGenerator() gopter.Gen {
	if sbTopicProperties_STATUS_ARMGenerator != nil {
		return sbTopicProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicProperties_STATUS_ARM(generators)
	sbTopicProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SBTopicProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBTopicProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForSBTopicProperties_STATUS_ARM(generators)
	sbTopicProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SBTopicProperties_STATUS_ARM{}), generators)

	return sbTopicProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSBTopicProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBTopicProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_STATUS_Active,
		EntityStatus_STATUS_Creating,
		EntityStatus_STATUS_Deleting,
		EntityStatus_STATUS_Disabled,
		EntityStatus_STATUS_ReceiveDisabled,
		EntityStatus_STATUS_Renaming,
		EntityStatus_STATUS_Restoring,
		EntityStatus_STATUS_SendDisabled,
		EntityStatus_STATUS_Unknown))
	gens["SubscriptionCount"] = gen.PtrOf(gen.Int())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSBTopicProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSBTopicProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUS_ARMGenerator())
}