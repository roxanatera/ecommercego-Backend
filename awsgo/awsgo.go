package awsgo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var (
	Ctx context.Context
	Cfg aws.Config
	err error
)

func InitAWS() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("error to load SDK config, " + err.Error())
	}




}