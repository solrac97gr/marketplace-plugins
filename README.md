# Marketplace Plugins

A personal collection of custom plugins and skills for [Claude Code](https://claude.ai/code).

## What is this?

This repository contains custom plugins that extend Claude Code's functionality through user-defined skills. Each plugin can provide one or more skills that can be invoked during Claude Code sessions using the `/skill-name` command.

## Available Plugins

### review-plugin

Adds a `/review` skill for quick code reviews.

**Usage**: `/review`

The skill analyzes selected code or recent changes for:
- Potential bugs or edge cases
- Security concerns
- Performance issues
- Readability improvements

## Installation

To use these plugins with Claude Code:

1. Clone this repository
2. Configure your Claude Code marketplace settings to point to this repository
3. The plugins will be available in your Claude Code sessions

## Repository Structure

```
marketplace-plugins/
├── .claude-plugin/
│   └── marketplace.json      # Marketplace manifest
├── plugins/
│   └── [plugin-name]/
│       ├── .claude-plugin/
│       │   └── plugin.json   # Plugin metadata
│       └── skills/
│           └── [skill-name]/
│               └── SKILL.md  # Skill definition
└── CLAUDE.md                 # Development guide
```

## Creating a New Plugin

1. Create a new directory under `plugins/`:
   ```bash
   mkdir -p plugins/my-plugin/.claude-plugin
   mkdir -p plugins/my-plugin/skills/my-skill
   ```

2. Add plugin metadata in `.claude-plugin/plugin.json`:
   ```json
   {
     "name": "my-plugin",
     "description": "Description of what your plugin does",
     "version": "1.0.0"
   }
   ```

3. Create a skill in `skills/my-skill/SKILL.md`:
   ```markdown
   ---
   description: Brief description
   disable-model-invocation: true
   ---

   Instructions for Claude when this skill is invoked.
   ```

4. Register the plugin in `.claude-plugin/marketplace.json`:
   ```json
   {
     "plugins": [
       {
         "name": "my-plugin",
         "source": "./plugins/my-plugin",
         "description": "Description of what your plugin does"
       }
     ]
   }
   ```

## Contributing

This is a personal marketplace, but feel free to fork and create your own custom plugins!

## License

MIT
