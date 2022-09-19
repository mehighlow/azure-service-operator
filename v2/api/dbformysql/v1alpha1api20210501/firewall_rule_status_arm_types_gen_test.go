// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

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

func Test_FirewallRule_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRule_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRule_STATUS_ARM, FirewallRule_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRule_STATUS_ARM runs a test to see if a specific instance of FirewallRule_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRule_STATUS_ARM(subject FirewallRule_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRule_STATUS_ARM
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

// Generator of FirewallRule_STATUS_ARM instances for property testing - lazily instantiated by
// FirewallRule_STATUS_ARMGenerator()
var firewallRule_STATUS_ARMGenerator gopter.Gen

// FirewallRule_STATUS_ARMGenerator returns a generator of FirewallRule_STATUS_ARM instances for property testing.
// We first initialize firewallRule_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FirewallRule_STATUS_ARMGenerator() gopter.Gen {
	if firewallRule_STATUS_ARMGenerator != nil {
		return firewallRule_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRule_STATUS_ARM(generators)
	firewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FirewallRule_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRule_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForFirewallRule_STATUS_ARM(generators)
	firewallRule_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FirewallRule_STATUS_ARM{}), generators)

	return firewallRule_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFirewallRule_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFirewallRule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FirewallRuleProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_FirewallRuleProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FirewallRuleProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFirewallRuleProperties_STATUS_ARM, FirewallRuleProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFirewallRuleProperties_STATUS_ARM runs a test to see if a specific instance of FirewallRuleProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFirewallRuleProperties_STATUS_ARM(subject FirewallRuleProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FirewallRuleProperties_STATUS_ARM
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

// Generator of FirewallRuleProperties_STATUS_ARM instances for property testing - lazily instantiated by
// FirewallRuleProperties_STATUS_ARMGenerator()
var firewallRuleProperties_STATUS_ARMGenerator gopter.Gen

// FirewallRuleProperties_STATUS_ARMGenerator returns a generator of FirewallRuleProperties_STATUS_ARM instances for property testing.
func FirewallRuleProperties_STATUS_ARMGenerator() gopter.Gen {
	if firewallRuleProperties_STATUS_ARMGenerator != nil {
		return firewallRuleProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUS_ARM(generators)
	firewallRuleProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(FirewallRuleProperties_STATUS_ARM{}), generators)

	return firewallRuleProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFirewallRuleProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EndIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["StartIpAddress"] = gen.PtrOf(gen.AlphaString())
}