.PHONY = unittests run shell help
PYTHON = python3
PYTEST = pytest
POETRY = poetry

help: ## Show usage instructions
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install Python packages to virtual environment
	${POETRY} install

run: shell ## Run Flask application
	${PYTHON} run.py

shell: install ## Activate Python virtual environment
	poetry shell

unittests: shell ## Run unit tests
	${PYTEST} --cov tracker --cov-branch --cov-report term-missing tests
