package controllers

import (
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

// GetAllTools uses the helm client to query a cluster for all Helm charts
func GetAllTools() ([]*release.Release, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)

	err := actionConfig.Init(settings.RESTClientGetter(), "", os.Getenv("HELM_DRIVER"), nil)
	if err != nil {
		return nil, err
	}

	client := action.NewList(actionConfig)
	client.Deployed = true
	results, err := client.Run()
	if err != nil {
		return nil, err
	}

	return results, nil
}
