# This makes the subsequent variables available to child shells
.EXPORT_ALL_VARIABLES:

# Bring the ephemeral env into the Makefile
include ../go-cfg.tfvars

ENV_NAME := $(shell echo $(env-name) | tr -d '"')

# login:  	## Who Am I for AWS
# 	@ AWS_ACCOUNT_ID=$$(aws sts get-caller-identity --output text --query 'Account'); \
# 	AWS_IAM_ARN=$$(aws sts get-caller-identity --output text --query 'Arn'); \
# 	echo "Running as $${AWS_IAM_ARN} in $${AWS_REGION} for $${AWS_ACCOUNT_ID}."

init:		## Terraform init with credentials, and ephemeral stuff
		terraform init 
		
validate:	## Terraform Validate
		terraform validate

plan:		## Terraform Plan
		terraform plan --var-file ../go-cfg.tfvars

apply:  	## Terraform Apply
		date; \
		terraform apply --var-file ../go-cfg.tfvars -auto-approve

destroy:	## Terraform Destroy
		terraform destroy --var-file ../go-cfg.tfvars
