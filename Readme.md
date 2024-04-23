# Frontend Masters Course:  Build Go Apps that Scale on AWS

This is a companion repo containing the code solutions for the [Build Go Apps that Scale on AWS course on Frontend Masters](https://frontendmasters.com/courses/go-aws)

## Setup Instructions

In order to complete this course, you should create an **AWS account**, install **Go**, the **AWS CLI** and the **AWS CDK**.

### Install Go (preferably version > = 1.18)
- Download and install Go: [Go Download Page](https://go.dev/dl/)

### Install the AWS CLI
- Create an [AWS Account](https://portal.aws.amazon.com/billing/signup) if you don't already have one.
- Install latest version of the [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
  - Verify the installation by running `aws --version`
- Configure an AWS user permissions 
  - Create an [IAM or IAM Identity Center administrative account](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-prereqs.html#getting-started-prereqs-iam) 
- Step 3 : Configure the AWS CLI
  - Configure [short or long term credentials ](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-quickstart.html)
  - Confirm your setup with: `aws s3 ls` and `aws sts get-caller-identity`

### CDK installed (use CDK command)
- [Download CDK](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html#getting_started_install)
- Make sure you bootstrap your environment with:
```bash
cdk bootstrap aws://ACCOUNT-NUMBER/REGION
```
Here's [more info](https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html#getting_started_bootstrap) about boostrapping the CDK. Your `ACCOUNT-NUMBER` and `REGION` can be found the AWS console.

### CloudWatch Logs

Some users received CloudWatch logging errors during the API Gateway lessons. Logging to CloudWatch is not required, but you can prevent these errors by [following these steps.](https://repost.aws/knowledge-center/api-gateway-cloudwatch-logs)

### Gopls Extension

Download the [Gopls Extension](https://github.com/golang/tools/blob/master/gopls/README.md) for your editor

## Code Solutions

All code for this course is written from scratch. The branches in this repo are code checkpoints for sections within the course:

* The `main` branch contains everything you need for the **Go Basics** section
* The `chapter_0` branch contains the solution to the **AWS & CDK** section
* branch `chapter_1` branch contains the solution to the **DynamoDB & Insertin Users** section
* branch `chapter_2` branch contains the solution to the **AAPI Gateway & Authentication Routes** section
* branch `chapter_3` branch contains the solution to the **JSON Web Tokens & Protected URLs** section

**Note:** The code in the above branches may not match what was coded during the course. See the `live_workshop` branch for the final course code. 

## cURL Commands

At the end of the course, the API routes are tested with cURL commands. You can copy/paste these commands instead of typing them out.

### Register

```bash
curl -X POST AWS_SERVER_URL/register -H "Content-Type: application/json" -d '{"username":"USERNAME", "password":"PASSWORD"}'
```

### Login

```bash
curl -X POST AWS_SERVER_URL/login -H "Content-Type: application/json" -d '{"username":"USERNAME", "password":"PASSWORD"}'
```

### Access Protected Route

```bash
curl -X GET AWS_SERVER_URL/protected -H "Content-Type: application/json" -H "Authorization: Bearer JWT_TOKEN"
```

## Remove infrastructure create with CDK

To remove the infrastructure created with the AWS CDK, run `cdk destory`. 
