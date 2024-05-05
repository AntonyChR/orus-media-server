#!/bin/bash

# This script is used to configure Git hooks for the Orus Media Server project.
# It sets the core.hooksPath configuration to point to the 'git-hooks' directory,
# which contains custom Git hooks specific to the project.

# Execute "git config --unset core.hooksPath" to remove the custom hooks path.
git config core.hooksPath git-hooks/

cd git-hooks
chmod +x pre-commit
