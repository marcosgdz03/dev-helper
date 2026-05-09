package cmd

import (
	"fmt"
	"strings"

	"github.com/marcosgdz03/dev-helper/core"
	"github.com/spf13/cobra"
)

var dockerizeCmd = &cobra.Command{
	Use:   "dockerize",
	Short: "Generate a Dockerfile for the project",
	Long: `Create a Dockerfile appropriate for the project language and framework.

Examples:
  dev-helper dockerize --lang go --framework gin --output ./my-api
  dev-helper dockerize --lang node --framework express`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDockerize()
	},
}

var dockerizeLang string
var dockerizeFramework string
var dockerizeOutput string

func init() {
	dockerizeCmd.Flags().StringVarP(&dockerizeLang, "lang", "l", "go", "project language")
	dockerizeCmd.Flags().StringVarP(&dockerizeFramework, "framework", "f", "gin", "project framework")
	dockerizeCmd.Flags().StringVarP(&dockerizeOutput, "output", "o", ".", "project directory")
	rootCmd.AddCommand(dockerizeCmd)
}

func runDockerize() error {
	if dockerizeLang == "" {
		return fmt.Errorf("--lang is required")
	}

	scaff := core.NewScaffolder(cfgFile)

	lang := strings.ToLower(dockerizeLang)
	framework := strings.ToLower(dockerizeFramework)

	if !scaff.Config.ValidateLanguage(lang) {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	return scaff.CreateDockerfile(lang, framework, dockerizeOutput)
}
