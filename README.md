# Linear Workflows Manager

This repository implements Linear APIs and Webhooks in order to handle specific workflows that we required at Livestorm. In order to run this project, we provide two main methods:

- **Local Development:** This project is written using standard golang libraries and can be run by simply installing the required libararies using `go mod download` and then running the following command: `go run cmd/webhook/main.go`

- **Docker:** Alternatively, a Dockerfile (along with a docker-compose) is also included in the project in order to easily spawn an instance of this service.

## Features

- Mandatory Label management
- Personalised Slack Notifications (In Development)
- Automatic ticket column selection based on user's Slack Group (In Development)

Have a feature request, feel free to create a Github Issue for this and we'll analyse if this could be included into our service.
 
### Configuration
This project allows for a team-level management of workflows from Linear. In order to set this up for a given team, you'll need to perform the actions provided below. The configuration of a default team named `Tech` with the code "LIV" is provided by default to be used as inspiration.

##### Env Variables
You'll find a `.env.sample` file present at the root of the repository. This contains all the relevant keys required to run this project effectively. Copy this and recreate a `.env` file filling in your Linear API Key and a default state for your manager.

##### Configuring your webhook
In: `internal/webhook/linear.go`, you would need to add a route and handler for your team matching a corresponding folder in `internal/workflows/linear`. We'e already provided a means of authenticating the source of requests and this should allow you to easily get started with your team.

You can alternatively modify the existing `Tech` team's webhook to adapt this for your use case.

##### Setting up your team
Team-specific configurations can be found in `core/config/linear`. In order to add a configuration for your team, simply create a file in this folder with the name of your team and add the relevant configurations following the `BoardConfig` type provided in the `types.go` file present in the same folder.

An example of this configuration can be seen in `tech.go` file present in this folder. Mandatory Labels are defined using the format:

```
    {
		"In Progress": { // represents the state from which this configuration is supposed to be applied
			"global": {"type", "scope", "squad", "chapter"}, // global represents all tickets, you can alternatively use a `type: ` label in your ticket to condition the presence of labels by type. 
		},
		"Ready for deploy": {
			"bug":        {"bug-reason"},
			"regression": {"bug-reason"},
		},
	}
```

Furthermore, you would need to retrieve the list of BoardStates manually from the Linear API and organise them based on your choice. This is kept manual for the moment in order to render the state's order flexible for developers using the mandatory labels approach. A good example of this would be to keep the Cancelled state in the beginning to avoid label restrictions to apply.