package main

import (
	"fmt"
	"os/exec"

	"github.com/drone/drone-go/drone"
)

// Target the API to execute against
func api(api API) *exec.Cmd {
	uri := api.URI
	require("api", uri)
	fmt.Printf("Target api %s\n", uri)
	return exec.Command("cf", "api", uri)
}

// Login to cloud foundry
func login(credentials Credentials) *exec.Cmd {
	user, pass := credentials.User, credentials.Password
	require("user", user)
	require("password", pass)

	fmt.Println("Logging in...")
	return exec.Command("cf", "auth", user, pass)
}

// cf target
func target(vargs Target) *exec.Cmd {
	org, space := vargs.Org, vargs.Space
	require("org", org)
	require("space", space)
	fmt.Printf("Targeting %s:%s...\n", org, space)
	return exec.Command("cf", "target", "-o", org, "-s", space)
}

// cf deploy
func deploy(workspace drone.Workspace, app App, route Route, flags Flags) *exec.Cmd {
	fmt.Println("Deploy")
	args := combine(
		[]string{"push"},
		app.args(),
		route.args(),
		flags.args(),
	)

	cmd := exec.Command("cf", args...)
	cmd.Dir = workspace.Path
	return cmd
}
