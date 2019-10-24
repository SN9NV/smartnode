package main

import (
    "os"

    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/rocketpool-api/node"
    "github.com/rocket-pool/smartnode/shared/utils/api"
    cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)


// Run
func main() {

    // Initialise application
    app := cli.NewApp()

    // Set application info
    app.Name = "rocketpool-api"
    app.Usage = "Rocket Pool node API"
    app.Version = "0.0.1"
    app.Authors = []cli.Author{
        cli.Author{
            Name:  "David Rugendyke",
            Email: "david@rocketpool.net",
        },
        cli.Author{
            Name:  "Jake Pospischil",
            Email: "jake@rocketpool.net",
        },
    }
    app.Copyright = "(c) 2019 Rocket Pool Pty Ltd"

    // Configure application
    cliutils.Configure(app)

    // Register commands
    node.RegisterCommands(app, "node", []string{"n"})

    // Run application
    if err := app.Run(os.Args); err != nil {
        api.PrintErrorResponse(nil, err)
    }

}

