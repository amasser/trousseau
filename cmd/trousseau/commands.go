package main

import (
	"github.com/codegangsta/cli"
	"github.com/oleiade/trousseau"
	"log"
	"fmt"
	"strings"
)

func CreateCommand() cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "create the trousseau data store",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Invalid number of arguments provided to create command")
			}

			var recipients []string = strings.Split(c.Args()[0], ",")
			trousseau.CreateAction(recipients)
		},
	}
}

func PushCommand() cli.Command {
	return cli.Command{
		Name:   "push",
		Usage:  "pushes the trousseau to remote storage",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Invalid number of arguments provided to push command")
			}

			var destination string = c.Args().First()
			trousseau.PushAction(destination, c.String("ssh-private-key"), c.Bool("ask-password"))
		},
		Flags: []cli.Flag{
			OverwriteFlag(),
			AskPassword(),
			VerboseFlag(),
			SshPrivateKeyPathFlag(),
		},
	}
}

func PullCommand() cli.Command {
	return cli.Command{
		Name:   "pull",
		Usage:  "pull the trousseau from remote storage",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Invalid number of arguments provided to pull command")
			}

			var source string = c.Args().First()
			trousseau.PullAction(source, c.String("ssh-private-key"), c.Bool("ask-password"))
		},
		Flags: []cli.Flag{
			OverwriteFlag(),
			AskPassword(),
			VerboseFlag(),
			SshPrivateKeyPathFlag(),
		},
	}
}

func ExportCommand() cli.Command {
	return cli.Command{
		Name:   "export",
		Usage:  "export the encrypted trousseau to local fs",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Invalid number of arguments provided to export command")
			}

			var to string = c.Args().First()
			trousseau.ExportAction(to, c.Bool("plain"))
		},
		Flags: []cli.Flag{
			OverwriteFlag(),
			PlainFlag(),
			VerboseFlag(),
		},
	}
}

func ImportCommand() cli.Command {
	return cli.Command{
		Name:   "import",
		Usage:  "import an encrypted trousseau from local fs",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Invalid number of arguments provided to import command")
			}

			var strategy trousseau.ImportStrategy
			var yours bool = c.Bool("yours")
			var theirs bool = c.Bool("theirs")
			var overwrite bool = c.Bool("overwrite")
			var activated uint = 0

			// Ensure two import strategies were not provided at
			// the same time. Otherwise, throw an error
			for _, flag := range []bool{yours, theirs, overwrite} {
				if flag {
					activated += 1
				}
				if activated >= 2 {
					log.Fatal("--yours, --theirs and --overwrite options are mutually exclusive")
				}
			}

			// Return proper ImportStrategy according to
			// provided flags
			if overwrite == true {
				strategy = trousseau.IMPORT_OVERWRITE
			} else if theirs == true {
				strategy = trousseau.IMPORT_THEIRS
			} else {
				strategy = trousseau.IMPORT_YOURS
			}

			var from string = c.Args().First()
			trousseau.ImportAction(from, strategy, c.Bool("plain"))
		},
		Flags: []cli.Flag{
			VerboseFlag(),
			PlainFlag(),
			OverwriteFlag(),
			TheirsFlag(),
			YoursFlag(),
		},
	}
}

func ListRecipientsCommand() cli.Command {
	return cli.Command{
		Name:   "list-recipients",
		Usage:  "lists trousseau data store recipients",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 0) {
				log.Fatal("Incorrect number of arguments provided to 'list-recipients' command")
			}

			trousseau.ListRecipientsAction()
		},
		Flags: []cli.Flag{
			VerboseFlag(),
		},
	}
}

func AddRecipientCommand() cli.Command {
	return cli.Command{
		Name:   "add-recipient",
		Usage:  "add a recipient to the encrypted trousseau",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Incorrect number of arguments to 'add-recipient' command")
			}

			trousseau.AddRecipientAction(c.Args().First())

			if c.Bool("verbose") == true {
				trousseau.Logger.Info(fmt.Sprintf("Recipient added to trousseau data store: %s", c.Args().First()))
			}
		},
		Flags: []cli.Flag{
			VerboseFlag(),
		},
	}
}

func RemoveRecipientCommand() cli.Command {
	return cli.Command{
		Name:   "remove-recipient",
		Usage:  "remove a recipient of the encrypted trousseau",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Incorrect number of arguments to 'remove-recipient' command")
			}

			trousseau.RemoveRecipientAction(c.Args().First())

			if c.Bool("verbose") == true {
				fmt.Printf("Recipient removed from trousseau data store: %s", c.Args().First())
			}

		},
		Flags: []cli.Flag{
			VerboseFlag(),
		},
	}
}

func SetCommand() cli.Command {
	return cli.Command{
		Name:   "set",
		Usage:  "sets a key value pair in the store",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 2) {
				log.Fatal("Invalid number of arguments provided to set command")
			}

			var key string = c.Args().First()
			var value string = c.Args()[1]
			var file string = c.String("file")

			trousseau.SetAction(key, value, file)

			if c.Bool("verbose") == true {
				trousseau.Logger.Info(fmt.Sprintf("%s:%s", key, value))
			}
		},
		Flags: []cli.Flag{
			FileFlag(),
			VerboseFlag(),
		},
	}
}

func GetCommand() cli.Command {
	return cli.Command{
		Name:   "get",
		Usage:  "get a value from the trousseau",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Invalid number of arguments provided to get command")
			}

			var key string = c.Args().First()
			var file string = c.String("file")
			trousseau.GetAction(key, file)
		},
		Flags: []cli.Flag{
			FileFlag(),
		},
	}
}

func RenameCommand() cli.Command {
	return cli.Command{
		Name:   "rename",
		Usage:  "rename an existing key",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 2) {
				log.Fatal("Invalid number of arguments provided to rename command")
			}

			var src string = c.Args().First()
			var dest string = c.Args()[1]

			trousseau.RenameAction(src, dest, c.Bool("overwrite"))

			if c.Bool("verbose") == true {
				trousseau.Logger.Info(fmt.Sprintf("renamed: %s to %s", src, dest))
			}
		},
		Flags: []cli.Flag{
			OverwriteFlag(),
			VerboseFlag(),
		},
	}
}

func DelCommand() cli.Command {
	return cli.Command{
		Name:   "del",
		Usage:  "delete the point key pair from the store",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 1) {
				log.Fatal("Invalid number of arguments provided to del command")
			}

			var key string = c.Args().First()

			trousseau.DelAction(key)

			if c.Bool("verbose") == true {
				trousseau.Logger.Info(fmt.Sprintf("deleted: %s", c.Args()[0]))
			}
		},
		Flags: []cli.Flag{
			VerboseFlag(),
		},
	}
}

func KeysCommand() cli.Command {
	return cli.Command{
		Name:   "keys",
		Usage:  "Lists the store keys",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 0) {
				log.Fatal("Invalid number of arguments provided to keys command")
			}

			trousseau.KeysAction()
		},
		Flags: []cli.Flag{
			VerboseFlag(),
		},
	}
}

func ShowCommand() cli.Command {
	return cli.Command{
		Name:   "show",
		Usage:  "shows trousseau content",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 0) {
				log.Fatal("Invalid number of arguments provided to show command")
			}

			trousseau.ShowAction()
		},
	}
}

func MetaCommand() cli.Command {
	return cli.Command{
		Name:   "meta",
		Usage:  "shows trousseau metadata",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 0) {
				log.Fatal("Invalid number of arguments provided to meta command")
			}

			trousseau.MetaAction()
		},
	}
}

func UpgradeCommand() cli.Command {
	return cli.Command{
		Name:   "upgrade",
		Usage:  "Upgrade your data store to a newer version",
		Action: func(c *cli.Context) {
			if !hasExpectedArgs(c.Args(), 0) {
				log.Fatal("Invalid number of arguments provided to upgrade command")
			}

			trousseau.UpgradeAction(c.Bool("yes"), c.Bool("no-backup"))
		},
		Flags: []cli.Flag{
			YesFlag(),
			NoBackupFlag(),
		},
	}
}

// hasExpectedArgs checks whether the number of args are as expected.
func hasExpectedArgs(args []string, expected int) bool {
	switch expected {
	case -1:
		if len(args) > 0 {
			return true
		} else {
			return false
		}
	default:
		if len(args) == expected {
			return true
		} else {
			return false
		}
	}
}

