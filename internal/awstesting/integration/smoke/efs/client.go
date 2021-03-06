// +build integration

//Package efs provides gucumber integration tests support.
package efs

import (
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	_ "github.com/aws/aws-sdk-go-v2/internal/awstesting/integration/smoke"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@efs", func() {
		// FIXME remove custom region
		cfg := integration.Config()
		cfg.Region = "us-west-2"

		gucumber.World["client"] = efs.New(cfg)
	})
}
