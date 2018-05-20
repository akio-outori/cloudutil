package main

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/endpoints"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudformation"
)


func constructSession() (*cloudformation.CloudFormation) {
  cloudformationNew := session.Must(session.NewSession(&aws.Config{
    Region: aws.String(endpoints.UsEast2RegionID),
  }))
  cloudformationSession := cloudformation.New(cloudformationNew)

  return cloudformationSession
}

func getCurrentStacks() (stackNames []string) {
  cloudformationSession := constructSession()
  input, err            := cloudformationSession.DescribeStacks(nil)

  if err != nil {
  }

  for _, stack := range input.Stacks {
    stackNames = append(stackNames, *stack.StackName)
  }

  return stackNames
}

func getStackStatus(stackName string) (result string) {
  cloudformationSession := constructSession()
  input, err            := cloudformationSession.DescribeStacks(nil)

  if err != nil {
  }

  for _, stack := range input.Stacks {
    if *stack.StackName == stackName {
      result = *stack.StackStatus
      break
    }
  }

  return result
}

