.PHONY: build

build:
	sam build

deploy-infra:
	sam build && aws-vault exec lou --no-session -- sam deploy

deploy-site:
	aws-vault exec lou --no-session -- aws s3 sync ./resume-site s3://louisnguyen-cloud-resume-2026