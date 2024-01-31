/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"html/template"
	"os"

	"github.com/spf13/cobra"
)

const deploymentTemplate string = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.Name}}
spec:
  replicas: {{.Replicas}}
  selector:
    matchLabels:
      app: {{.Name}}
  template:
    metadata:
      labels:
        app: {{.Name}}
    spec:
      containers:
      - name: {{.Name}}
        image: {{.Image}}
        ports:
        - containerPort: {{.Port}}
`

type Deployment struct {
	Name     string
	Replicas int
	Image    string
	Port     int
}

func YAMLIfy() *cobra.Command {

	var name, image string
	var replicas, port int
	var rootCmd = &cobra.Command{

		Use:   "YAMLIfy.git",
		Short: "A CLI tool that helps to create K8s deployment configuration yaml ",
		RunE: func(cmd *cobra.Command, args []string) error {
			deployment := Deployment{
				Name:     name,
				Replicas: replicas,
				Image:    image,
				Port:     port,
			}

			t := template.Must(template.New("deployment").Parse(deploymentTemplate))

			err := t.Execute(os.Stdout, deployment)
			if err != nil {
				return err
			}

			return nil
		},
	}
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "The name of the deployment")
	rootCmd.MarkFlagRequired("name")
	rootCmd.Flags().StringVarP(&image, "image", "i", "", "The Docker image to deploy")
	rootCmd.MarkFlagRequired("image")
	rootCmd.Flags().IntVarP(&replicas, "replicas", "r", 1, "The number of replicas to deploy")
	rootCmd.Flags().IntVarP(&port, "port", "p", 80, "The port the container should listen on")
	return rootCmd

}

// rootCmd represents the base command when called without any subcommands

func Execute() {
	err := YAMLIfy().Execute()
	if err != nil {
		os.Exit(1)
	}
}
