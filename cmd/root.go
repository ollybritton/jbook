package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/manifoldco/promptui"
	"github.com/ollybritton/jbook/pkg/jbook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/toqueteos/webbrowser"
)

var logger *logrus.Logger
var output string
var popopen bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jbook",
	Short: "jbook is a tool for generating an mdBook-formatted webpage for viewing jrnls.",
	Long:  `jbook is a tool for generating an mdBook-formatted webpage for viewing journals created by the jrnl tool.`,
	Args:  cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		var name string

		if len(args) == 0 {
			name = "default"
		} else {
			name = args[0]
		}

		logger.Infof("Using journal %q...", name)
		jrnl := jbook.Export(name)

		if len(jrnl.Entries) == 0 {
			logger.Fatalf("Journal %q is empty. Check your spelling and try again.", name)
		}

		dir := jbook.CreateSource(jrnl, "")
		book := jbook.Build(dir, output)

		webbrowser.Open("file://" + path.Join(book, "index.html"))

		fmt.Println("")
		logrus.Warnf("Do you want to delete the source and book?")

		choice := yesNo()

		if choice {
			logrus.Infof("Removing book and source...")
			os.RemoveAll(book)
			os.RemoveAll(dir)
		} else {
			logrus.Infof("Respecting your choice. Relevant directories:")
			logrus.Infof("* Rendered/Built Book: %v", book)
			logrus.Infof("* Source: %v", dir)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogging, checkTools)

	rootCmd.Flags().StringP("level", "l", "info", "log level (fatal, error, info, debug, trace)")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "where to output formatted journal (defaults to a temp directory)")
}

func initLogging() {
	logger = logrus.New()

	strLvl, err := rootCmd.Flags().GetString("level")
	if err != nil {
		logrus.Panicf("%v", err)
	}

	lvl, err := logrus.ParseLevel(strLvl)
	if err != nil {
		logrus.Fatalf("invalid level %q: try fatal, error, info, debug or trace", lvl)
	}

	logger.SetLevel(lvl)
	jbook.SetLogger(logger)
}

func checkTools() {
	ok := jbook.CheckTools()

	if !ok {
		logger.Fatalf("Missing either `jrnl` or `mdbook` from path. Are you sure you have them installed?")
	}
}

func yesNo() bool {
	fmt.Println("")

	prompt := promptui.Select{
		Label: "Select [Yes/No]",
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	fmt.Println("")
	return result == "Yes"
}
