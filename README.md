# Sample Function: MySQL

## Introduction

This repository contains a sample function that lets you connect to a MySQL instance and demonstrates performing some queries on it like creating a table and doing some selects. You can deploy it on DigitalOcean's AppPlatform as a Serverless Function component. You can optionally add your app as a trusted source to a DO Managed DBaaS instance and have the App connect to it securely.

Documentation is available at https://docs.digitalocean.com/products/functions

### Requirements

* You need a DigitalOcean account. If you don't already have one, you can sign up at [https://cloud.digitalocean.com/registrations/new](https://cloud.digitalocean.com/registrations/new).
* You need to have a MySQL database that the function can connect to. You can [create a MySQL cluster](https://docs.digitalocean.com/products/databases/mysql/how-to/create/) on DO managed DBaaS
* You need to provide the database connection string of the form `mysql://{user}:{password}@{db_fqdn}:25060/{db_name}?ssl-mode=REQUIRED` in the `DATABASE_URL` variable in the `.env` file.
* To deploy from the command line, you will need the [DigitalOcean `doctl` CLI](https://github.com/digitalocean/doctl/releases).


## Deploying the Function

```bash
# clone this repo
git clone git@github.com:digitalocean/sample-functions-mysql.git
```

```
# deploy the project, using a remote build so that compiled executable matched runtime environment
> doctl serverless deploy sample-functions-mysql --remote-build --env sample-functions-mysql/packages/ts/hello/.env
Deploying 'sample-functions-mysql'
  to namespace 'fn-...'
  on host 'https://faas-...'
Submitted action 'ts/hello' for remote building and deployment in runtime go:default (id: 7f947f1f0050461a947f1f0050261a5f)
Processing of action 'ts/hello' is still running remotely ...
Deployment status recorded in 'sample-functions-mysql/.deployed'

Deployed functions ('doctl sbx fn get <funcName> --url' for URL):
  - ts/hello
```

### Learn More

You can learn more about Functions and App Platform integration in [the official App Platform Documentation](https://www.digitalocean.com/docs/app-platform/).
