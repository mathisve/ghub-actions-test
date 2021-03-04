package main

import "github.com/sethvargo/go-githubactions"

func main() {
	name := githubactions.GetInput("name")
	if name == "" {
		githubactions.Fatalf("missing input 'name'")
	}

	githubactions.AddMask(name)
}