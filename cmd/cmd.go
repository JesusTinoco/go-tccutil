package cmd

import (
	"fmt"

	"github.com/JesusTinoco/go-tccutil/tccutil"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(addClient)
	RootCmd.AddCommand(removeClient)
	RootCmd.AddCommand(enableClient)
	RootCmd.AddCommand(disableClient)
	RootCmd.AddCommand(listClients)
}

var RootCmd = &cobra.Command{
	Use:   "tccutil2",
	Short: "tccutil2 is a tccutil vitaminized",
	Long:  `tccutil2 allows you modify OS X's Accessibility Database from the command line (https://github.com/JesusTinoco/go-tccutil).`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tccutil2",
	Long:  "Print the version number of tccutil2",
	Run: func(cmd *cobra.Command, clients []string) {
		fmt.Println("tccutil2 v2.0 -- HEAD")
	},
}

var addClient = &cobra.Command{
	Use:        "add",
	Short:      "Add new clients to the database",
	Long:       "Add the new clients given as arguments to the database",
	SuggestFor: []string{"insert"},
	Run: func(cmd *cobra.Command, clients []string) {
		for _, client := range clients {
			tccutil.InsertClient(client)
		}
	},
}

var removeClient = &cobra.Command{
	Use:        "remove",
	Short:      "Remove clients from the database",
	Long:       "Remove from the database the new clients given as arguments",
	SuggestFor: []string{"delete", "uninstall", "erase"},
	Run: func(cmd *cobra.Command, clients []string) {
		for _, client := range clients {
			tccutil.RemoveClient(client)
		}
	},
}

var enableClient = &cobra.Command{
	Use:        "enable",
	Short:      "Enable the given clients",
	Long:       "Enable the given clients",
	SuggestFor: []string{"allow"},
	Run: func(cmd *cobra.Command, clients []string) {
		for _, client := range clients {
			tccutil.EnableClient(client)
		}
	},
}

var disableClient = &cobra.Command{
	Use:   "disable",
	Short: "Disable the given clients",
	Long:  "Disable the given clients",
	Run: func(cmd *cobra.Command, clients []string) {
		for _, client := range clients {
			tccutil.DisableClient(client)
		}
	},
}

var listClients = &cobra.Command{
	Use:   "list",
	Short: "List all the current clients registered.",
	Long: `Will print a list of all the current clients that
            are registered in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		clients := tccutil.ListClients()
		for _, client := range clients {
			fmt.Println(client.Name)
		}
	},
}
