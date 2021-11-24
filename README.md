# `flip` <em>the toolkit for open blockchain data</em>

`flip` is a command line tool that provides progamatic access to Flipside's on-chain data exchange.

## What can you do with this tool?

1. Access a fully featured DBT environment and curate your own datasets.
2. Gain programmatic access to Flipside's data exchange (based in snowflake) and utilize your favorite programming languages to interact with on-chain data (python, r, go, typescript, etc.)

## Getting Started

1. [PREREQUISITE] Download [Docker for Desktop](https://www.docker.com/products/docker-desktop).
2. Download the latest release [here](https://github.com/FlipsideCrypto/flip/releases/tag/0.3.0) for your machine (mac, linux, windows) and unpack its contents.
3. Login with your Flipside Velocity account (`./flip auth login`).
4. Run `./flip --help` to see all the available commands.

## Commands

### Auth

`flip auth login`
Login to the tool with your Flipside Velocity account. If you don't have a Flipside Velocity account sign-up [here](https://app.flipsidecrypto.com/auth/signup).

`flip auth logout`
Logout of your account.

`flip auth me`
View details on your account.

### DataX

`flip datax creds`
Return credentials to Flipside's Data Exchange. These creds will provide you programmatic access to Flipside's curated data views.

### Curate

The following commands allow you to interact with a fully-featured curation environment powered by DBT.

`flip curate init`
Generate a Flipside DBT project on your local machine.

! NOTE !
After generating a project `cd <your project directory>` to run the commands below.

`flip curate dbt-console`
Enter into a fully-featured dbt environment. No DBT install is required. The only requirement is that you have docker installed on your machine. This command must be run from the root of your DBT project.

`flip curate dbt-docs`
Compile and access documentation for your Flipside DBT project. This command must be run from the root of your dbt project.

`flip curate reset-env`
Remove the current docker image from your machine and re-download the latest version.

`flip curate flipside-docs`
Retrieve a link to Flipside's official [table documentation](https://docs.flipsidecrypto.com/our-data/tables).
