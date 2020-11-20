# a git tag/branch/commit within anchore/anchore-engine repo
ANCHORE_ENGINE_REF = 68572ed6231501652dc3eaf7c2752919d0c0e367
OPENAPI_GENERATOR_VERSION = v4.1.3
PROJECT_ROOT = pkg
ENGINE_ROOT = $(PROJECT_ROOT)/external
ENGINE_SWAGGER_DOC = $(PROJECT_ROOT)/swagger-external-$(ANCHORE_ENGINE_REF).yaml

define generate_openapi_client
	# remove previous API clients
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo rm -rf ${3}; else rm -rf ${3}; fi

	# generate the API client
	docker run \
		--rm \
		-v $${PWD}:/local \
		--env GIT_USER_ID=anchore \
		--env GIT_REPO_ID=client-go \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
			generate \
				--additional-properties packageName=${2} \
				--additional-properties withGoCodegenComment=true \
				-i /local/${1} \
				-g go \
				-o /local/${3}

	# keep permissions the same as the user
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo chown -R $(shell id -u):$(shell id -g) ${3}; fi

	# remove unused files (go.mod is curated manually)
	rm ${3}/{.travis.yml,git_push.sh,go.*}
endef

.DEFAULT_GOAL :=
all: clean generate-engine-client

$(ENGINE_SWAGGER_DOC):
	mkdir -p $(PROJECT_ROOT)
	curl https://raw.githubusercontent.com/anchore/anchore-engine/$(ANCHORE_ENGINE_REF)/anchore_engine/services/apiext/swagger/swagger.yaml -o $(ENGINE_SWAGGER_DOC)

.PHONY :=
generate-engine-client: $(ENGINE_SWAGGER_DOC) ## Generate engine external API client from Anchore OpenAPI spec
	$(call generate_openapi_client,$(ENGINE_SWAGGER_DOC),client,$(ENGINE_ROOT))

.PHONY :=
clean:
	rm -rf $(PROJECT_ROOT)

.PHONY :=
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
