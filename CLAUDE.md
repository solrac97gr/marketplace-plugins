# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a personal Claude Code marketplace repository for custom plugins and skills. The repository follows the Claude Code plugin marketplace structure, with a `marketplace.json` manifest that defines available plugins.

## Repository Structure

- **marketplace.json**: Root manifest defining the marketplace name, owner, and list of plugins
- **plugins/**: Directory containing individual plugin packages
  - Each plugin has its own `.claude-plugin/plugin.json` metadata file
  - Skills are defined in `skills/[skill-name]/SKILL.md` within each plugin

## Plugin Structure

Each plugin follows this structure:
```
plugins/[plugin-name]/
├── .claude-plugin/
│   └── plugin.json          # Plugin metadata (name, description, version)
└── skills/
    └── [skill-name]/
        └── SKILL.md         # Skill definition with frontmatter and instructions
```

### SKILL.md Format

Skill files use YAML frontmatter followed by the skill instructions:
```markdown
---
description: Brief description of what the skill does
disable-model-invocation: true  # Optional flag
---

[Skill instructions that Claude will follow when the skill is invoked]
```

## Adding New Plugins

1. Create a new directory under `plugins/[plugin-name]/`
2. Add `.claude-plugin/plugin.json` with metadata:
   ```json
   {
     "name": "plugin-name",
     "description": "What the plugin does",
     "version": "1.0.0"
   }
   ```
3. Create skills under `skills/[skill-name]/SKILL.md`
4. Register the plugin in root `marketplace.json` under the `plugins` array

## Adding New Skills to Existing Plugins

1. Create a new directory under `plugins/[plugin-name]/skills/[skill-name]/`
2. Add a `SKILL.md` file with frontmatter and instructions
3. Skills are automatically discovered from the plugin directory structure
