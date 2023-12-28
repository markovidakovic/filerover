package cli

var (
	fileFlag        string
	dirFlag         string
	sourceFlag      string
	destinationFlag string
)

func init() {
	createCmd.Flags().StringVarP(&fileFlag, "file", "f", "", "Create a file with a specified name")
	createCmd.Flags().StringVarP(&dirFlag, "dir", "d", "", "Create a directory with a specified name")

	copyCmd.Flags().StringVarP(&sourceFlag, "source", "s", "", "Source file/directory path")
	copyCmd.Flags().StringVarP(&destinationFlag, "destination", "d", "", "Destination path")

	deleteCmd.Flags().StringVarP(&fileFlag, "file", "f", "", "Delete the specified file")
	deleteCmd.Flags().StringVarP(&dirFlag, "dir", "d", "", "Delete the specified directory")

	moveCmd.Flags().StringVarP(&sourceFlag, "source", "s", "", "Source file/directory path")
	moveCmd.Flags().StringVarP(&destinationFlag, "destination", "d", "", "Destination path")

	infoCmd.Flags().StringVarP(&fileFlag, "file", "f", "", "File path")

	listCmd.Flags().StringVarP(&dirFlag, "dir", "d", "", "Directory path")
}
