package cli

import (
	"fmt"

	"github.com/markovidakovic/file-rover/internal/fileops"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filerover",
	Short: "A CLI tool for file management",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		// filename, _ := cmd.Flags().GetString("file")
		// dirname, _ := cmd.Flags().GetString("dir")

		if fileFlag != "" {
			err := fileops.CreateFile(fileFlag)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if dirFlag != "" {
			err := fileops.CreateDirectory(dirFlag)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Please provide either --file or --dir flag")
		}
	},
}

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if sourceFlag != "" && destinationFlag != "" {
			err := fileops.CopyFileOrDirectory(sourceFlag, destinationFlag)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Please provide both --source and --destination flags")
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if fileFlag != "" {
			err := fileops.DeleteFile(fileFlag)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else if dirFlag != "" {
			err := fileops.DeleteDirectory(dirFlag)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Please provide either --file or --dir flag")
		}
	},
}

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move a file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if sourceFlag != "" && destinationFlag != "" {
			err := fileops.MoveFileOrDirectory(sourceFlag, destinationFlag)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Please provider both --source and --destination flags")
		}
	},
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display file info",
	Run: func(cmd *cobra.Command, args []string) {
		if fileFlag != "" {
			err := fileops.DisplayFileInfo(fileFlag)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Please provide a file path using --file flag")
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List contents of a directory",
	Run: func(cmd *cobra.Command, args []string) {
		if dirFlag != "" {
			err := fileops.ListDirectoryContents(dirFlag)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Please provide a directory path using --dir flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(copyCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(moveCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(listCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
