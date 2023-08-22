package main

import (
	"fmt"
	"github.com/eavesmy/yaml2go"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	rootCmd = &cobra.Command{
		Use:   "yaml2go-cli",
		Short: "yaml2go-cli is a cli-tool for yaml to go struct",
		Run:   cmdRun,
	}

	inputFile   *string
	outputFile  *string
	structName  *string
	packageName *string
)

func init() {
	inputFile = rootCmd.PersistentFlags().StringP("input", "i", "", "input yaml file path")
	outputFile = rootCmd.PersistentFlags().StringP("output", "o", "", "output go file path")
	structName = rootCmd.PersistentFlags().StringP("struct", "s", "Default", "struct name")
	packageName = rootCmd.PersistentFlags().StringP("package", "p", "main", "package name")

	if err := rootCmd.MarkPersistentFlagRequired("input"); err != nil {
		panic(err)
	}

	if err := rootCmd.MarkPersistentFlagRequired("output"); err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func cmdRun(cmd *cobra.Command, args []string) {

	baFile, err := os.ReadFile(*inputFile)

	if err != nil {
		panic(err)
	}

	dir, _ := path.Split(*outputFile)

	if dir != "" {
		// *packageName = dir
		*packageName = path.Base(dir)
	}

	y := yaml2go.NewStruct(*packageName, *structName, baFile)

	if y == nil {
		fmt.Println("nil ??", y == nil)
	}

	if err = y.DoYaml2Struct(); err != nil {
		// fmt.Println(err.Error())
		panic(err)
	}

	y.StructStr = "// Automatic generate. Do not modify this file.\n\n" + y.StructStr

	err = os.WriteFile(*outputFile, []byte(y.StructStr), 0644)

	// adapter for unix
	err = os.Chmod(*outputFile, 0643)
	if err != nil {
		panic(err)
	}
}
