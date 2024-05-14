# This makes the subsequent variables available to child shells
.EXPORT_ALL_VARIABLES:

# include ../.env

# Collect Last Target, convert to variable, and consume the target.
# Allows passing arguments to the target recipes from the make command line.
CMD_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
# Consume them to prevent interpretation as targets
$(eval $(CMD_ARGS):;@:)
# Service for command args
ARGUMENT  := $(word 1,${CMD_ARGS})

##
## Usage:
##  make [target] [ARGUMENT]
##

help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

######################
## Git
######################

commit:		## Short hand for git Add and Commit
	git add .; git commit -m ${ARGUMENT}; git push

######################
## Terraform
######################

init:		## Init for Terraform Modules # -upgrade?
	terraform init

validate:	## Validate Configuration Files
	terraform validate

plan:	## Plan for Terraform Modules
	terraform plan --var-file ../go-cfg.tfvars --var-file ../go-secrets.tfvars

apply:		## Apply for Terraform Modules
	terraform apply --var-file ../go-cfg.tfvars --var-file ../go-secrets.tfvars

destroy:		## Destroy for Terraform Modules
	terraform destroy --var-file ../go-cfg.tfvars

sa_pk:		## Service Account Private Key
	@terraform output -json | jq -r '.service_account_private_key.value' | base64 --decode
