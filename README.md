# `flip` <em>the toolkit for open data</em>

`flip` is a command line tool that provides progamatic access to Flipside's on-chain data exchange.

## What can you do with this tool?

1. Access a fully featured DBT environment and curate your own datasets.
2. Gain programmatic access to Flipside's data exchange (based in snowflake) and utilize your favorite programming languages to interact with on-chain data (python, r, go, typescript, etc.)

## Getting Started

1. Download the latest release [here](https://github.com/FlipsideCrypto/flip/releases) for your machine (mac, linux, windows) and unpack it's contents.
2. Login with your Flipside Velocity account (`./flip auth login`).
3. Run `./flip --help` to see all the available commands.

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

<b>Coming soon...</b>

### Curate (coming very soon)

`flip curate init`
Generate a Flipside DBT project on your local machine.

`flip curate dbt-console`
Enter into a fully-featured dbt environment. No DBT install is required. The only requirement is that you have docker installed on your machine.

`flip curate dbt-docs`
Compile and access documentation for your Flipside DBT project.
