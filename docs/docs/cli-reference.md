---
sidebar_position: 8
sidebar_label: CLI Reference
---

# CLI Reference

All Gpt4cli CLI commands and their options.

## Usage

```bash
gpt4cli [command] [flags]
g4c [command] [flags] # 'g4c' is an alias for 'gpt4cli'
```

## Help

Built-in help.

```bash
gpt4cli help
g4c h # alias
```

`--all/-a`: List all commands.

For help on a specific command, use:

```bash
gpt4cli [command] --help
```

## REPL

The easiest way to use Gpt4cli is through the REPL. Start it in your project directory with:

```bash
gpt4cli
```

or for short:

```bash
g4c
```

### Flags

The REPL has a few convenient flags you can use to start it with different modes, autonomy settings, and model packs. You can pass any of these to `gpt4cli` or `g4c` when starting the REPL.

```
  Mode
    --chat, -c     Start in chat mode (for conversation without making changes)
    --tell, -t     Start in tell mode (for implementation)

  Autonomy
    --no-auto      None → step-by-step, no automation
    --basic        Basic → auto-continue plans, no other automation
    --plus         Plus → auto-update context, smart context, auto-commit changes
    --semi         Semi-Auto → auto-load context
    --full         Full-Auto → auto-apply, auto-exec, auto-debug

  Models
    --daily        Daily driver pack (default models, balanced capability, cost, and speed)
    --reasoning    Similar to daily driver, but uses reasoning model for planning
    --strong       Strong pack (more capable models, higher cost and slower)
    --cheap        Cheap pack (less capable models, lower cost and faster)
    --oss          Open source pack (open source models)
    --gemini-exp   Gemini experimental pack (Gemini 2.5 Pro Experimental for planning and coding, default models for other roles)
```

All commands listed below can be run in the REPL by prefixing them with a backslash (`\`), e.g. `\new`.

## Plans

### new

Start a new plan.

```bash
gpt4cli new
gpt4cli new -n new-plan # with name
```

`--name/-n`: Name of the new plan. The name is generated automatically after first prompt if no name is specified on creation.

`--context-dir/-d`: Base directory to load context from when auto-loading context is enabled. Defaults to `.` (current directory). Set a different directoy if you don't want all files to be included in the project map.

`--no-auto`: Start the plan with auto-mode 'None' (step-by-step, no automation).

`--basic`: Start the plan with auto-mode 'Basic' (auto-continue plans, no other automation).

`--plus`: Start the plan with auto-mode 'Plus' (auto-update context, smart context, auto-commit changes).

`--semi`: Start the plan with auto-mode 'Semi-Auto' (auto-load context).

`--full`: Start the plan with auto-mode 'Full-Auto' (auto-apply, auto-exec, auto-debug).

`--daily`: Start the plan with the daily driver model pack.

`--reasoning`: Start the plan with the reasoning model pack.

`--strong`: Start the plan with the strong model pack.

`--cheap`: Start the plan with the cheap model pack.

`--oss`: Start the plan with the open source model pack.

`--gemini-exp`: Start the plan with the Gemini experimental model pack.

### plans

List plans. Output includes index, when each plan was last updated, the current branch of each plan, the number of tokens in context, and the number of tokens in the conversation (prior to summarization).

Includes full details on plans in current directory. Also includes names of plans in parent directories and child directories.

```bash
gpt4cli plans
gpt4cli plans --archived # list archived plans only

g4c pl # alias
```

`--archived/-a`: List archived plans only.

### current

Show current plan. Output includes when the plan was last updated and created, the current branch, the number of tokens in context, and the number of tokens in the conversation (prior to summarization).

```bash
gpt4cli current
g4c cu # alias
```

### cd

Set current plan by name or index.

```bash
gpt4cli cd # select from a list of plans
gpt4cli cd some-plan # by name
gpt4cli cd 4 # by index in `gpt4cli plans`
```

With no arguments, Gpt4cli prompts you with a list of plans to select from.

With one argument, Gpt4cli selects a plan by name or by index in the `gpt4cli plans` list.

### delete-plan

Delete a plan by name, index, range, pattern, or select from a list.

```bash
gpt4cli delete-plan # select from a list of plans
gpt4cli delete-plan some-plan # by name
gpt4cli delete-plan 4 # by index in `gpt4cli plans`
gpt4cli delete-plan 2-4 # by range of indices
gpt4cli delete-plan 'docs-*' # by pattern
gpt4cli delete-plan --all # delete all plans
g4c dp # alias
```

`--all/-a`: Delete all plans.

### rename

Rename the current plan.

```bash
gpt4cli rename # prompt for new name
gpt4cli rename new-name # set new name
```

### archive

Archive a plan.

```bash
gpt4cli archive # select from a list of plans
gpt4cli archive some-plan # by name
gpt4cli archive 4 # by index in `gpt4cli plans`

g4c arc # alias
```

### unarchive

Unarchive a plan.

```bash
gpt4cli unarchive # select from a list of archived plans
gpt4cli unarchive some-plan # by name
gpt4cli unarchive 4 # by index in `gpt4cli plans --archived`
g4c unarc # alias
```

## Context

### load

Load files, directories, directory layouts, URLs, notes, images, or piped data into context.

```bash
gpt4cli load component.ts # single file
gpt4cli load component.ts action.ts reducer.ts # multiple files
gpt4cli load lib -r # loads lib and all its subdirectories
gpt4cli load tests/**/*.ts # loads all .ts files in tests and its subdirectories
gpt4cli load . --tree # loads the layout of the current directory and its subdirectories (file names only)
gpt4cli load https://redux.js.org/usage/writing-tests # loads the text-only content of the url
npm test | gpt4cli load # loads the output of `npm test`
gpt4cli load -n 'add logging statements to all the code you generate.' # load a note into context
gpt4cli load ui-mockup.png # load an image into context

g4c l component.ts # alias
```

`--recursive/-r`: Load an entire directory and all its subdirectories.

`--tree`: Load directory tree layout with file names only.

`--map`: Load file map of the given directory (function/method/class signatures, variable names, types, etc.)

`--note/-n`: Load a note into context.

`--force/-f`: Load files even when ignored by .gitignore or .gpt4cliignore.

`--detail/-d`: Image detail level when loading an image (high or low)—default is high. See https://platform.openai.com/docs/guides/vision/low-or-high-fidelity-image-understanding for more info.

### ls

List everything in the current plan's context. Output includes index, name, type, token size, when the context added, and when the context was last updated.

```bash
gpt4cli ls

gpt4cli list-context # longer alias
```

### rm

Remove context by index, range, name, or glob.

```bash
gpt4cli rm some-file.ts # by name
gpt4cli rm app/**/*.ts # by glob pattern
gpt4cli rm 4 # by index in `gpt4cli ls`
gpt4cli rm 2-4 # by range of indices

gpt4cli remove # longer alias
gpt4cli unload # longer alias
```

### show

Output context by name or index.

```bash
gpt4cli show some-file.ts # by name
gpt4cli show 4 # by index in `gpt4cli ls`
```

### update

Update any outdated context.

```bash
gpt4cli update
g4c u # alias
```

### clear

Remove all context.

```bash
gpt4cli clear
```

## Control

### tell

Describe a task.

```bash
gpt4cli tell -f prompt.txt # from file
gpt4cli tell # open vim to write prompt
gpt4cli tell "add a cancel button to the left of the submit button" # inline

g4c t # alias
```

`--file/-f`: File path containing prompt.

`--stop/-s`: Stop after a single model response (don't auto-continue). Defaults to opposite of config value `auto-continue`.

`--no-build/-n`: Don't build proposed changes into pending file updates. Defaults to opposite of config value `auto-build`.

`--bg`: Run task in the background. Only allowed if `--auto-load-context` and `--apply/-a` are not enabled. Not allowed with the default [autonomy level](./core-concepts/autonomy.md) in Gpt4cli v2.

`--auto-update-context`: Automatically confirm context updates. Defaults to config value `auto-update-context`.

`--auto-load-context`: Automatically load context using project map. Defaults to config value `auto-load-context`.

`--smart-context`: Use smart context to only load the necessary file(s) for each step during implementation. Defaults to config value `smart-context`.

`--no-exec`: Don't execute commands after successful apply. Defaults to opposite of config value `can-exec`.

`--auto-exec`: Automatically execute commands after successful apply without confirmation. Defaults to config value `auto-exec`.

`--debug`: Automatically execute and debug failing commands (optionally specify number of tries—default is 5). Defaults to config values of `auto-debug` and `auto-debug-tries`.

`--apply/-a`: Automatically apply changes (and confirm context updates). Defaults to config value `auto-apply`.

`--commit/-c`: Commit changes to git when `--apply/-a` is passed. Defaults to config value `auto-commit`.

`--skip-commit`: Don't commit changes to git. Defaults to opposite of config value `auto-commit`.

### continue

Continue the plan.

```bash
gpt4cli continue

g4c c # alias
```

`--stop/-s`: Stop after a single model response (don't auto-continue). Defaults to opposite of config value `auto-continue`.

`--no-build/-n`: Don't build proposed changes into pending file updates. Defaults to opposite of config value `auto-build`.

`--bg`: Run task in the background. Only allowed if `--auto-load-context` and `--apply/-a` are not enabled. Not allowed with the default [autonomy level](./core-concepts/autonomy.md) in Gpt4cli v2.

`--auto-update-context`: Automatically confirm context updates. Defaults to config value `auto-update-context`.

`--auto-load-context`: Automatically load context using project map. Defaults to config value `auto-load-context`.

`--smart-context`: Use smart context to only load the necessary file(s) for each step during implementation. Defaults to config value `smart-context`.

`--no-exec`: Don't execute commands after successful apply. Defaults to opposite of config value `can-exec`.

`--auto-exec`: Automatically execute commands after successful apply without confirmation. Defaults to config value `auto-exec`.

`--debug`: Automatically execute and debug failing commands (optionally specify number of tries—default is 5). Defaults to config values of `auto-debug` and `auto-debug-tries`.

`--apply/-a`: Automatically apply changes (and confirm context updates). Defaults to config value `auto-apply`.

`--commit/-c`: Commit changes to git when `--apply/-a` is passed. Defaults to config value `auto-commit`.

`--skip-commit`: Don't commit changes to git. Defaults to opposite of config value `auto-commit`.

### build

Build any unbuilt pending changes from the plan conversation.

```bash
gpt4cli build
g4c b # alias
```

`--bg`: Build in the background. Not allowed if `--apply/-a` is enabled.

`--stop/-s`: Stop after a single model response (don't auto-continue). Defaults to opposite of config value `auto-continue`.

`--no-build/-n`: Don't build proposed changes into pending file updates. Defaults to opposite of config value `auto-build`.

`--auto-update-context`: Automatically confirm context updates. Defaults to config value `auto-update-context`.

`--no-exec`: Don't execute commands after successful apply. Defaults to opposite of config value `can-exec`.

`--auto-exec`: Automatically execute commands after successful apply without confirmation. Defaults to config value `auto-exec`.

`--debug`: Automatically execute and debug failing commands (optionally specify number of tries—default is 5). Defaults to config values of `auto-debug` and `auto-debug-tries`.

`--apply/-a`: Automatically apply changes (and confirm context updates). Defaults to config value `auto-apply`.

`--commit/-c`: Commit changes to git when `--apply/-a` is passed. Defaults to config value `auto-commit`.

`--skip-commit`: Don't commit changes to git. Defaults to opposite of config value `auto-commit`.

### chat

Ask a question or chat without making any changes.

```bash
gpt4cli chat "is it clear from the context how to add a new line chart?"
g4c ch # alias
```

`--file/-f`: File path containing prompt.

`--bg`: Run task in the background. Not allowed if `--auto-load-context` is enabled. Not allowed with the default [autonomy level](./core-concepts/autonomy.md) in Gpt4cli v2.

`--auto-update-context`: Automatically confirm context updates. Defaults to config value `auto-update-context`.

`--auto-load-context`: Automatically load context using project map. Defaults to config value `auto-load-context`.

### debug

Repeatedly run a command and automatically attempt fixes until it succeeds, rolling back changes on failure. Defaults to 5 tries before giving up.

```bash
gpt4cli debug 'npm test' # try 5 times or until it succeeds
gpt4cli debug 10 'npm test' # try 10 times or until it succeeds
g4c db 'npm test' # alias
```

`--commit/-c`: Commit changes to git when `--apply/-a` is passed. Defaults to config value `auto-commit`.

`--skip-commit`: Don't commit changes to git. Defaults to opposite of config value `auto-commit`.

## Changes

### diff

Review pending changes in 'git diff' format or in a local browser UI.

```bash
gpt4cli diff
gpt4cli diff --ui
```

`--plain/-p`: Output diffs in plain text with no ANSI codes.

`--ui/-u`: Review pending changes in a local browser UI.

`--side-by-side/-s`: Show diffs UI in side-by-side view

`--line-by-line/-l`: Show diffs UI in line-by-line view

### apply

Apply pending changes to project files.

```bash
gpt4cli apply
g4c ap # alias
```

`--auto-update-context`: Automatically confirm context updates. Defaults to config value `auto-update-context`.

`--no-exec`: Don't execute commands after successful apply. Defaults to opposite of config value `can-exec`.

`--auto-exec`: Automatically execute commands after successful apply without confirmation. Defaults to config value `auto-exec`.

`--debug`: Automatically execute and debug failing commands (optionally specify number of tries—default is 5). Defaults to config values of `auto-debug` and `auto-debug-tries`.

`--commit/-c`: Commit changes to git when `--apply/-a` is passed. Defaults to config value `auto-commit`.

`--skip-commit`: Don't commit changes to git. Defaults to opposite of config value `auto-commit`.

`--full`: Apply the plan and debug in full auto mode.

### reject

Reject pending changes to one or more project files.

```bash
gpt4cli reject # select from a list of pending files to reject
gpt4cli reject file.ts # one file
gpt4cli reject file.ts another-file.ts # multiple files
gpt4cli reject --all # all pending files

g4c rj file.ts # alias
```

`--all/-a`: Reject all pending files.

## History

### log

Show plan history.

```bash
gpt4cli log

gpt4cli history # alias
gpt4cli logs # alias
```

### rewind

Rewind to a previous state.

```bash
gpt4cli rewind # select from a list of previous states to rewind to
gpt4cli rewind 3 # rewind 3 steps
gpt4cli rewind a7c8d66 # rewind to a specific step from `gpt4cli log`
```

### convo

Show the current plan's conversation.

```bash
gpt4cli convo
gpt4cli convo 1 # show a specific message
gpt4cli convo 1-5 # show a range of messages
gpt4cli convo 3- # show all messages from 3 to the end
```

`--plain/-p`: Output conversation in plain text with no ANSI codes.

### summary

Show the latest summary of the current plan.

```bash
gpt4cli summary
```

`--plain/-p`: Output summary in plain text with no ANSI codes.

## Branches

### branches

List plan branches. Output includes index, name, when the branch was last updated, the number of tokens in context, and the number of tokens in the conversation (prior to summarization).

```bash
gpt4cli branches
g4c br # alias
```

### checkout

Checkout or create a branch.

```bash
gpt4cli checkout # select from a list of branches or prompt to create a new branch
gpt4cli checkout some-branch # checkout by name or create a new branch with that name

g4c co # alias
```

### delete-branch

Delete a branch by name or index.

```bash
gpt4cli delete-branch # select from a list of branches
gpt4cli delete-branch some-branch # by name
gpt4cli delete-branch 4 # by index in `gpt4cli branches`

g4c dlb # alias
```

## Background Tasks / Streams

### ps

List active and recently finished plan streams. Output includes stream ID, plan name, branch name, when the stream was started, and the stream's status (active, finished, stopped, errored, or waiting for a missing file to be selected).

```bash
gpt4cli ps
```

### connect

Connect to an active plan stream.

```bash
gpt4cli connect # select from a list of active streams
gpt4cli connect a4de # by stream ID in `gpt4cli ps`
gpt4cli connect some-plan main # by plan name and branch name
g4c conn # alias
```

### stop

Stop an active plan stream.

```bash
gpt4cli stop # select from a list of active streams
gpt4cli stop a4de # by stream ID in `gpt4cli ps`
gpt4cli stop some-plan main # by plan name and branch name
```

## Configuration

### config

Show current plan config. Output includes configuration settings for the plan, such as autonomy level, model settings, and other plan-specific options.

```bash
gpt4cli config
```

### config default

Show the default config used for new plans. Output includes the default configuration settings that will be applied to newly created plans.

```bash
gpt4cli config default
```

### set-config

Update configuration settings for the current plan.

```bash
gpt4cli set-config # select from a list of config options
gpt4cli set-config auto-context true # set a specific config option
```

With no arguments, Gpt4cli prompts you to select from a list of config options.

With arguments, allows you to directly set specific configuration options for the current plan.

### set-config default

Update the default configuration settings for new plans.

```bash
gpt4cli set-config default # select from a list of config options
gpt4cli set-config default auto-mode basic # set a specific default config option
```

Works exactly the same as set-config above, but sets the default configuration for all new plans instead of only the current plan.

### set-auto

Update the auto-mode (autonomy level) for the current plan.

```bash
gpt4cli set-auto # select from a list of auto-modes
gpt4cli set-auto full # set to full automation
gpt4cli set-auto semi # set to semi-auto mode
gpt4cli set-auto plus # set to plus mode
gpt4cli set-auto basic # set to basic mode
gpt4cli set-auto none # set to none (step-by-step, no automation)
```

With no arguments, Gpt4cli prompts you to select from a list of automation levels.

With one argument, Gpt4cli sets the automation level directly to the specified value.

### set-auto default

Set the default auto-mode for new plans.

```bash
gpt4cli set-auto default # select from a list of auto-modes
gpt4cli set-auto default basic # set default to basic mode
```

Works exactly the same as set-auto above, but sets the default automation level for all new plans instead of only the current plan.

## Models

### models

Show current plan models and model settings.

```bash
gpt4cli models
```

### models default

Show org-wide default models and model settings for new plans.

```bash
gpt4cli models default
```

### models available

Show available models.

```bash
gpt4cli models available # show all available models
gpt4cli models available --custom # show available custom models only
```

`--custom`: Show available custom models only.

### set-model

Update current plan models or model settings.

```bash
gpt4cli set-model # select from a list of models and settings
gpt4cli set-model planner openai/gpt-4 # set the model for a role
gpt4cli set-model gpt-4-turbo-latest # set the current plan's model pack by name (sets all model roles at once—see `model-packs` below)
gpt4cli set-model builder temperature 0.1 # set a model setting for a role
gpt4cli set-model max-tokens 4000 # set the planner model overall token limit to 4000
gpt4cli set-model max-convo-tokens 20000  # set how large the conversation can grow before Gpt4cli starts using summaries
```

With no arguments, Gpt4cli prompts you to select from a list of models and settings.

With arguments, can take one of the following forms:

- `gpt4cli set-model [role] [model]`: Set the model for a role.
- `gpt4cli set-model [model-pack]`: Set the current plan's model pack by name.
- `gpt4cli set-model [role] [setting] [value]`: Set a model setting for a role.
- `gpt4cli set-model [setting] [value]`: Set a model setting for the current plan.

Models are specified as `provider/model-name`, e.g. `openai/gpt-4`, `openrouter/anthropic/claude-opus-3`, `together/mistralai/Mixtral-8x22B-Instruct-v0.1`, etc.

See all the model roles [here](./models/roles.md).

Model role settings:

- `temperature`: Higher temperature means more randomness, which can produce more creativity but also more errors.
- `top-p`: Top-p sampling is a way to prevent the model from generating improbable text by only considering the most likely tokens.

Plan settings:

- `max-tokens`: The overall token limit for the planner model.
- `max-convo-tokens`: How large the conversation can grow before Gpt4cli starts using summaries.
- `reserved-output-tokens`: The number of tokens reserved for output from the model.

### set-model default

Update org-wide default model settings for new plans.

```bash
gpt4cli set-model default # select from a list of models and settings
gpt4cli set-model default planner openai/gpt-4 # set the model for a role
# etc. — same options as `set-model` above
```

Works exactly the same as `set-model` above, but sets the default model settings for all new plans instead of only the current plan.

### models add

Add a custom model.

```bash
gpt4cli models add
```

Gpt4cli will prompt you for all required information to add a custom model.

### models delete

Delete a custom model.

```bash
gpt4cli models delete # select from a list of custom models
gpt4cli models delete some-model # by name
gpt4cli models delete 4 # by index in `gpt4cli models available --custom`
```

### model-packs

Show all available model packs.

```bash
gpt4cli model-packs
```

`--custom`: Show available custom (user-created) model packs only.

### model-packs create

Create a new custom model pack.

```bash
gpt4cli model-packs create
```

Gpt4cli will prompt you for all required information to create a custom model pack.

### model-packs show

Show a built-in or custom model pack's settings.

```bash
gpt4cli model-packs show # select from a list of built-in and custom model packs
gpt4cli model-packs show some-model-pack # by name
```

### model-packs update

Update a custom model pack's settings.

```bash
gpt4cli model-packs update # select from a list of custom model packs
gpt4cli model-packs update some-model-pack # by name
```

### model-packs delete

Delete a custom model pack.

```bash
gpt4cli model-packs delete
gpt4cli model-packs delete some-model-pack # by name
gpt4cli model-packs delete 4 # by index in `gpt4cli model-packs --custom`
```

## Account Management

### sign-in

Sign in, accept an invite, or create an account.

```bash
gpt4cli sign-in
```

`--pin`: Sign in with a pin from the Gpt4cli Cloud web UI.

Unless you pass `--pin` (from the Gpt4cli Cloud web UI), Gpt4cli will prompt you for all required information to sign in, accept an invite, or create an account.

### invite

Invite a user to join your org.

```bash
gpt4cli invite # prompt for email, name, and role
gpt4cli invite name@domain.com 'Full Name' member # invite with email, name, and role
```

Users can be invited as `member`, `admin`, or `owner`.

### revoke

Revoke an invite or remove a user from your org.

```bash
gpt4cli revoke # select from a list of users and invites
gpt4cli revoke name@domain.com # by email
```

### users

List users and pending invites in your org.

```bash
gpt4cli users
```

## Gpt4cli Cloud

### billing

Show the billing settings page.

```bash
gpt4cli billing
```

### usage

Show Gpt4cli Cloud current balance and usage report. Includes recent spend, amount saved by input caching, a breakdown of spend by plan, category, and model, and a log of individual transactions with the `--log` flag.

Defaults to showing usage for the current session if you're using the REPL. Otherwise, defaults to showing usage for the day so far.

Requires **Integrated Models** mode.

```bash
gpt4cli usage
```

`--today`: Show usage for the day so far.

`--month`: Show usage for the current billing month.

`--plan`: Show usage for the current plan.

`--log`: Show a log of individual transactions. Defaults to showing the log for the current session if you're using the REPL. Otherwise, defaults to showing the log for the day so far. Works with `--today`, `--month`, and `--plan` flags.

Flags for `usage --log`:

`--debits`: Show only debits in the log.

`--purchases`: Show only purchases in the log.

`--page-size/-s`: Number of transactions to display per page.

`--page/-p`: Page number to display.






