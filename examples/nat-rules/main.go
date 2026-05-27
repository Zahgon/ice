// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

// Package main provides a simple CLI that prints the host and srflx addresses
// produced by different address rewrite (1:1) rule configurations. It is designed to be run
// inside the accompanying docker-compose topology so that each scenario can
// demonstrate how multi-homed hosts, srflx pools, CIDR scoping, and TCP muxing
// interact with the new rules.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/pion/ice/v4"
)

const (
	defaultTimeout = 8 * time.Second
)

type scenario struct {
	Key             string
	Title           string
	Description     string
	RewriteRules    []ice.AddressRewriteRule
	NetworkTypes    []ice.NetworkType
	CandidateTypes  []ice.CandidateType
	TimeoutOverride time.Duration
}

func (s scenario) timeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (s scenario) requiresTCPMux() bool { _ = "STUB: not implemented"; return false }

func main() {
	log.SetFlags(0)

	var (
		scenarioKey string
		listOnly    bool
		timeout     time.Duration
	)

	flag.StringVar(&scenarioKey, "scenario", "all", "Scenario key to run (use -list to see options)")
	flag.BoolVar(&listOnly, "list", false, "List available scenarios")
	flag.DurationVar(&timeout, "timeout", 0, "Override gather timeout for each scenario")
	flag.Parse()

	scenarios := buildScenarios(timeout)

	if listOnly {
		fmt.Println("Available scenarios:")
		for _, sc := range scenarios {
			fmt.Printf("  %s\t%s\n", sc.Key, sc.Title)
		}

		return
	}

	fmt.Println("Address rewrite rule demonstration client")
	printInterfaceSnapshot()

	ctx := context.Background()

	if scenarioKey == "all" {
		for _, sc := range scenarios {
			if err := runScenario(ctx, sc); err != nil {
				log.Fatalf("scenario %s failed: %v", sc.Key, err)
			}
		}

		return
	}

	sc, ok := findScenario(scenarios, scenarioKey)
	if !ok {
		log.Fatalf("unknown scenario %q. Use -list to see valid keys.", scenarioKey)
	}

	if err := runScenario(ctx, sc); err != nil {
		log.Fatalf("scenario %s failed: %v", sc.Key, err)
	}
}

func findScenario(scenarios []scenario, key string) (scenario, bool) {
	_ = "STUB: not implemented"
	return *new(scenario), false
}

func buildScenarios(timeout time.Duration) []scenario { _ = "STUB: not implemented"; return nil }

func buildMultiNetworkScenario() scenario { _ = "STUB: not implemented"; return *new(scenario) }

func buildIfaceScopedScenario() scenario { _ = "STUB: not implemented"; return *new(scenario) }

func buildSrflxScenario() scenario { _ = "STUB: not implemented"; return *new(scenario) }

func buildScopedCatchAllScenario() scenario { _ = "STUB: not implemented"; return *new(scenario) }

func runScenario(ctx context.Context, sc scenario) error {
	_ = "STUB: not implemented" //nolint:cyclop
	return nil
}

//nolint:contextcheck

//nolint:contextcheck

//nolint:err113

func formatCandidate(c ice.Candidate) string { _ = "STUB: not implemented"; return "" }

func printRules(rules []ice.AddressRewriteRule) { _ = "STUB: not implemented"; return }

func describeRuleScope(rule ice.AddressRewriteRule) string { _ = "STUB: not implemented"; return "" }

func formatNetworkList(networks []ice.NetworkType) string { _ = "STUB: not implemented"; return "" }

func envOrDefault(key, fallback string) string { _ = "STUB: not implemented"; return "" }

func printInterfaceSnapshot() { _ = "STUB: not implemented"; return }
