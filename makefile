# Detect the operating system
ifeq ($(OS),Windows_NT)
	# Windows
	OS_TARGET := windows
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		# Linux
		OS_TARGET := linux
	endif
endif

# Define the script directory
SCRIPT_DIR := ./scripts

.PHONY: run

run:
ifeq ($(OS_TARGET),windows)
	# Windows: Use batch script to run PowerShell
	@cmd /c run_powershell.bat
else
	# Linux: Use Bash
	@bash "$(SCRIPT_DIR)/docker-deploy.sh"
endif
