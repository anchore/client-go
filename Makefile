# The version of this client (should be in line with the highest supported engine/enterprise version
CLIENT_VERSION = 0.1.0

# where all generated code will be located
PROJECT_ROOT = pkg

# OpenAPI generator version to use
OPENAPI_GENERATOR_VERSION = v4.3.1
# note: v5 introduces the new command pattern approach, splitting request and execute + generating interfaces per service.
#OPENAPI_GENERATOR_VERSION = v5.0.0-beta3

# --- anchore engine references
# a git tag/branch/commit within anchore/anchore-engine repo
ENGINE_REF = ad0d3d654243d7f134cddadbda655ac9333a5e59
EXTAPI_CLIENT_ROOT = $(PROJECT_ROOT)/external
EXTAPI_OPENAPI_DOC = $(PROJECT_ROOT)/swagger-external-$(ENGINE_REF).yaml

# --- anchore enterprise references
# ANCHORE_ENTERPRISE_REF = ...


define generate_openapi_client
	# remove previous API clients
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo rm -rf ${3}; else rm -rf ${3}; fi

	# generate the API client
	docker run \
		--rm \
		-v $${PWD}:/local \
		openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} \
			generate \
				-c /local/generator.yaml \
				--additional-properties packageName=${2} \
				--additional-properties packageVersion=$(CLIENT_VERSION) \
				--http-user-agent "anchore-client/$(CLIENT_VERSION)/go" \
				-i /local/${1} \
				-o /local/${3} \
				-g go

	# keep permissions the same as the user
	@ if [[ "$$OSTYPE" == "linux-gnu" ]]; then sudo chown -R $(shell id -u):$(shell id -g) ${3}; fi

	# remove unused files (go.mod is curated manually)
	rm ${3}/{.travis.yml,git_push.sh,go.*}
endef

.PHONY :=
.DEFAULT_GOAL :=
update: clean generate-external-client ## pull all swagger definitions and generate client code

.PHONY :=
generate: generate-external-client ## generate all client code from all swagger documents

$(EXTAPI_OPENAPI_DOC): ## pull the engine external API swagger document
	mkdir -p $(PROJECT_ROOT)
	# note: the existing upstream swagger document needs to be corrected, otherwise invalid code will be generated.
	# the tr/sed cmds are a workaround for now.
	curl https://raw.githubusercontent.com/anchore/anchore-engine/$(ENGINE_REF)/anchore_engine/services/apiext/swagger/swagger.yaml > $(EXTAPI_OPENAPI_DOC)

.PHONY :=
generate-external-client: $(EXTAPI_OPENAPI_DOC) ## generate client code for engine external API
	$(call generate_openapi_client,$(EXTAPI_OPENAPI_DOC),external,$(EXTAPI_CLIENT_ROOT))

.PHONY :=
clean: ## remove all swagger documents and generated client code
	rm -rf $(PROJECT_ROOT)

.PHONY :=
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
