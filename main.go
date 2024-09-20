package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "mycli",
        Short: "MyCLI is a simple CLI application",
        Long:  `MyCLI is a simple CLI application built with Cobra in Go.`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Welcome to MyCLI!")
        },
    }

    var echoCmd = &cobra.Command{
        Use:   "echo [message]",
        Short: "Echo prints the input message",
        Long:  `Echo prints the input message to the console.`,
        Args:  cobra.MinimumNArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println(args[0])
        },
    }

    rootCmd.AddCommand(echoCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}