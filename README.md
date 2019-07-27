# Drone — Cloudflare Caching
> Drone plugin to purge cache via Cloudflare's API

![](https://img.shields.io/badge/License-MIT-lightgray.svg?style=for-the-badge)
![](https://img.shields.io/docker/stars/jetrails/drone-cloudflare-caching.svg?style=for-the-badge&colorB=9f9f9f)
![](https://img.shields.io/docker/pulls/jetrails/drone-cloudflare-caching.svg?style=for-the-badge&colorB=9f9f9f)

## About

> **WARNING**: Logic works if authenticated with email and global API token, not currently working with scoped API tokens. This is an issue with Cloudflare and should hopefully be resolved soon.

Our Drone plugin enables the ability for your pipeline to interface with Cloudflare's API to purge cache. This plugin is written in Go and it uses the [cloudflare-go](https://github.com/cloudflare/cloudflare-go) package to communicate with Cloudflare's API. For information on Cloudflare's API please refer to their [documentation](https://api.cloudflare.com/#zone-purge-all-files) page.

## Cloudflare Token

The API token that is used to authenticate with Cloudflare's API can be created in Cloudflare's dashboard. It is recommended to create an API token that includes only the zone resource you want to manipulate and give edit permissions only to the Cache Purge resource.

## Build

Develop locally by running the plugin with the following commands. Also please note that you should specify environmental variables for the plugin either via the inline method (`FOO=bar go run src/main`) or via the `export FOO=bar` method before the `go run` command.

```shell
$ dep ensure
$ go run src/main.go
```

## Docker

Drone plugins work off of docker images. The following commands will go over building, pushing, and running the docker image for this plugin.

###### Build Docker Image:

```shell
$ docker build -t jetrails/drone-cloudflare-caching .
```

###### Run Docker Container:

You can then replicate the command that Drone will use to launch the plugin by running:

```shell
$ docker run --rm \
	-e PLUGIN_API_TOKEN="u4C7ev06GMS8_vWBTpjqtVReT3I7FwGpW7MG44ZD" \
	-e PLUGIN_ZONE_IDENTIFIER="eJzrjE44s6Ki67x1tSDJzI8LdXxM3nj7" \
	-e PLUGIN_ACTION="purge_everything" \
	-v $(pwd):/drone/src \
	-w /drone/src \
	jetrails/drone-cloudflare-caching
```

###### Push Docker Image:

Finally, push this image to our Docker Hub [repository](https://hub.docker.com/r/jetrails/drone-cloudflare-caching) (assuming you have permission):

```shell
$ docker push jetrails/drone-cloudflare-caching
```

## Usage

This plugin supports purging all cache, purging hosts, purging files, and purging tags. Please refer to the table below with all possible settings that can be passed to the plugin:

|       Name      |           Required         | Default | Case-Sensitive |                            Type                           |
|:---------------:|:--------------------------:|:-------:|:--------------:|:---------------------------------------------------------:|
|    api_token    |             Yes            |    -    |       Yes      |                           STRING                          |
| zone_identifier |             Yes            |    -    |       Yes      |                           STRING                          |
|      action     |             Yes            |    -    |       No       | ENUM[purge_everything,purge_hosts,purge_files,purge_tags] |
|       list      | action != purge_everything |    -    |       Yes      |                       ARRAY\<STRING\>                       |

## Examples

```yaml
kind: pipeline
name: default

steps:
-   name: cloudflare
    image: jetrails/drone-cloudflare-caching
    settings:
        api_token:
            from_secret: cloudflare_token
        zone_identifier:
            from_secret: cloudflare_zone_identifier
        action: purge_everything
```

```yaml
kind: pipeline
name: default

steps:
-   name: cloudflare
    image: jetrails/drone-cloudflare-caching
    settings:
        api_token:
            from_secret: cloudflare_token
        zone_identifier:
            from_secret: cloudflare_zone_identifier
        action: purge_hosts
        list:
        -   example.com
        -   foo.example.com
        -   bar.example.com
```

```yaml
kind: pipeline
name: default

steps:
-   name: cloudflare
    image: jetrails/drone-cloudflare-caching
    settings:
        api_token:
            from_secret: cloudflare_token
        zone_identifier:
            from_secret: cloudflare_zone_identifier
        action: purge_files
        list:
        -   https://example.com/script.js
        -   https://example.com/logo.svg
```

```yaml
kind: pipeline
name: default

steps:
-   name: cloudflare
    image: jetrails/drone-cloudflare-caching
    settings:
        api_token:
            from_secret: cloudflare_token
        zone_identifier:
            from_secret: cloudflare_zone_identifier
        action: purge_tags
        list:
        -   foo
        -   bar
```

## Feature Requests / Issues

Feel free to open an issue for any feature requests and issues that you may come across. For furthur inquery, please contact [development@jetrails.com](mailto://development@jetrails.com).
