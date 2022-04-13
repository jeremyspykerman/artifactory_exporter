# JFrog Artifactory Exporter 

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/peimanja/artifactory_exporter/Build)](https://github.com/peimanja/artifactory_exporter/actions) [![Docker Build](https://img.shields.io/docker/cloud/build/peimanja/artifactory_exporter)](https://hub.docker.com/r/peimanja/artifactory_exporter/builds) [![Go Report Card](https://goreportcard.com/badge/github.com/peimanja/artifactory_exporter)](https://goreportcard.com/report/github.com/peimanja/artifactory_exporter)

A [Prometheus](https://prometheus.io) exporter for [JFrog Artifactory](https://jfrog.com/artifactory) stats. 


## Note

This exporter is under development and more metrics will be added later. Tested on Artifactory Commercial, Enterprise and OSS version `6.x` and `7.x`.

## Authentication

The Artifactory exporter requires **admin** user and it supports multiple means of authentication. The following methods are supported:
  * Basic Auth
  * Bearer Token

### Basic Auth

Basic auth may be used by setting `ARTI_USERNAME` and `ARTI_PASSWORD` environment variables.

### Bearer Token

Artifactory access tokens may be used via the Authorization header by setting `ARTI_ACCESS_TOKEN` environment variable.

## Usage

### Binary

Download the binary for your operation system from [release](https://github.com/peimanja/artifactory_exporter/releases) page and run it:
```bash
$ ./artifactory_exporter <flags>
```

### Docker

Set the credentials in `env_file_name` and you can deploy this exporter using the [peimanja/artifactory_exporter](https://registry.hub.docker.com/r/peimanja/artifactory_exporter/) Docker image:
:

```bash
$ docker run --env-file=env_file_name -p 9531:9531 peimanja/artifactory_exporter:latest <flags>
```

## Install with Helm

[Helm](https://helm.sh) must be installed to use the charts.
Please refer to Helm's [documentation](https://helm.sh/docs/) to get started.

Source code for the exporter helm chart can be found [here peimanja/helm-charts](https://github.com/peimanja/helm-charts/tree/main/charts/prometheus-artifactory-exporter)

Once Helm is set up properly, add the repo as follows:

### Prerequisites

- Kubernetes 1.8+ with Beta APIs enabled

### Add Repo

```console
helm repo add peimanja https://peimanja.github.io/helm-charts
helm repo update
```

_See [helm repo](https://helm.sh/docs/helm/helm_repo/) for command documentation._

### Configuration

See [Customizing the Chart Before Installing](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing). To see all configurable options with detailed comments, visit the chart's [values.yaml](https://github.com/peimanja/helm-charts/blob/main/charts/prometheus-artifactory-exporter/values.yaml), or run these configuration commands:


```console
# Helm 3
helm show values peimanja/prometheus-artifactory-exporter
```

Set your values in `myvals.yaml`:
```yaml
artifactory:
  url: http://artifactory:8081/artifactory
  accessToken: "xxxxxxxxxxxxxxxxxxxx"
  existingSecret: false

options:
  logLevel: info
  logFormat: logfmt
  telemetryPath: /metrics
  verifySSL: false
  timeout: 5s
```

### Install Chart

```console
# Helm 3
helm install -f myvals.yaml [RELEASE_NAME] peimanja/prometheus-artifactory-exporter
```

_See [configuration](#configuration) below._

_See [helm install](https://helm.sh/docs/helm/helm_install/) for command documentation._

### Flags

```bash
$  docker run peimanja/artifactory_exporter:latest -h
usage: main --artifactory.user=ARTIFACTORY.USER [<flags>]

Flags:
  -h, --help                    Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":9531"
                                Address to listen on for web interface and telemetry.
      --web.telemetry-path="/metrics"
                                Path under which to expose metrics.
      --artifactory.scrape-uri="http://localhost:8081/artifactory"
                                URI on which to scrape JFrog Artifactory.
      --artifactory.repo-quota-enable  
                                Flag that enables repo quotas data collection (https://github.com/jfrog/artifactory-user-plugins/tree/master/storage/repoQuota)                                
      --artifactory.ssl-verify  Flag that enables SSL certificate verification for the scrape URI
      --artifactory.timeout=5s  Timeout for trying to get stats from JFrog Artifactory.
      --log.level=info          Only log messages with the given severity or above. One of: [debug, info, warn, error]
      --log.format=logfmt       Output format of log messages. One of: [logfmt, json]
```

| Flag / Environment Variable | Required | Default | Description |
| --------------------------- | -------- | ------- | ----------- |
| `web.listen-address`<br/>`WEB_LISTEN_ADDR` | No | `:9531`| Address to listen on for web interface and telemetry. |
| `web.telemetry-path`<br/>`WEB_TELEMETRY_PATH` | No | `/metrics` | Path under which to expose metrics. |
| `artifactory.scrape-uri`<br/>`ARTI_SCRAPE_URI` | No | `http://localhost:8081/artifactory` | URI on which to scrape JFrog Artifactory. |
| `artifactory.ssl-verify`<br/>`ARTI_SSL_VERIFY` | No | `true` | Flag that enables SSL certificate verification for the scrape URI. |
| `artifactory.timeout`<br/>`ARTI_TIMEOUT` | No | `5s` | Timeout for trying to get stats from JFrog Artifactory. |
| `artifactory.repo-quota-enable`<br/> `ARTI_REPO_QUOTA_ENABLE` | No  | `false`  | Enables the quota monitoring. |
| `log.level` | No | `info` | Only log messages with the given severity or above. One of: [debug, info, warn, error]. |
| `log.format` | No | `logfmt` | Output format of log messages. One of: [logfmt, json]. |
| `ARTI_USERNAME` | *No | | User to access Artifactory |
| `ARTI_PASSWORD` | *No | | Password of the user accessing the Artifactory |
| `ARTI_ACCESS_TOKEN` | *No | | Access token for accessing the Artifactory |

* Either `ARTI_USERNAME` and `ARTI_PASSWORD` or `ARTI_ACCESS_TOKEN` environment variables has to be set.

### Metrics
The quota metrics are only available if the feature to collect them is enabled.

Some metrics are not available with Artifactory OSS license. The exporter returns the following metrics:

| Metric | Description | Labels | OSS support |
| ------ | ----------- | ------ | ------ |
| artifactory_storage_repo_quota_used_percent | Percentage of quota used |`name`, `package_type`, `type` |  TODO: confirm it works with OSS  |
| artifactory_storage_repo_quota_bytes  | Artifactory quota in bytes |`name`, `package_type`, `type` |  TODO: confirm it works with OSS  |
| artifactory_up | Was the last scrape of Artifactory successful. |  | &#9989; |
| artifactory_exporter_build_info | Exporter build information. | `version`, `revision`, `branch`, `goversion` | &#9989; |
| artifactory_exporter_total_scrapes | Current total artifactory scrapes. |  | &#9989; |
| artifactory_exporter_total_api_errors | Current total Artifactory API errors when scraping for stats. |  | &#9989; |
| artifactory_exporter_json_parse_failures |Number of errors while parsing Json. |  | &#9989; |
| artifactory_replication_enabled | Replication status for an Artifactory repository (1 = enabled). | `name`, `type`, `cron_exp` | |
| artifactory_security_groups | Number of Artifactory groups. | | |
| artifactory_security_users | Number of Artifactory users for each realm. | `realm` | |
| artifactory_storage_artifacts | Total artifacts count stored in Artifactory. |  | &#9989; |
| artifactory_storage_artifacts_size_bytes | Total artifacts Size stored in Artifactory in bytes. |  | &#9989; |
| artifactory_storage_binaries | Total binaries count stored in Artifactory. |  | &#9989; |
| artifactory_storage_binaries_size_bytes | Total binaries Size stored in Artifactory in bytes. |  | &#9989; |
| artifactory_storage_filestore_bytes | Total space in the file store in bytes. | `storage_dir`, `storage_type` | &#9989; |
| artifactory_storage_filestore_used_bytes | Space used in the file store in bytes. | `storage_dir`, `storage_type` | &#9989; |
| artifactory_storage_filestore_free_bytes | Space free in the file store in bytes. | `storage_dir`, `storage_type` | &#9989; |
| artifactory_storage_repo_used_bytes | Space used by an Artifactory repository in bytes. | `name`, `package_type`, `type` | &#9989; |
| artifactory_storage_repo_folders | Number of folders in an Artifactory repository. | `name`, `package_type`, `type` | &#9989; |
| artifactory_storage_repo_files | Number files in an Artifactory repository. | `name`, `package_type`, `type` | &#9989; |
| artifactory_storage_repo_items | Number Items in an Artifactory repository. | `name`, `package_type`, `type` | &#9989; |
| artifactory_artifacts_created_1m | Number of artifacts created in the repo (last 1 minute). | `name`, `package_type`, `type` | &#9989; |
| artifactory_artifacts_created_5m | Number of artifacts created in the repo (last 5 minutes). | `name`, `package_type`, `type` | &#9989; |
| artifactory_artifacts_created_15m | Number of artifacts created in the repo (last 15 minutes). | `name`, `package_type`, `type` | &#9989; |
| artifactory_artifacts_downloaded_1m | Number of artifacts downloaded from the repository (last 1 minute). | `name`, `package_type`, `type` | &#9989; |
| artifactory_artifacts_downloaded_5m | Number of artifacts downloaded from the repository (last 5 minutes). | `name`, `package_type`, `type` | &#9989; |
| artifactory_artifacts_downloaded_15m | Number of artifacts downloaded from the repository (last 15 minute). | `name`, `package_type`, `type` | &#9989; |
| artifactory_system_healthy | Is Artifactory working properly (1 = healthy). | | &#9989; |
| artifactory_system_license | License type and expiry as labels, seconds to expiration as value | `type`, `licensed_to`, `expires` | &#9989; |
| artifactory_system_version | Version and revision of Artifactory as labels. | `version`, `revision` | &#9989; |


### Grafana Dashboard

Dashboard can be found [here](https://grafana.com/grafana/dashboards/12113).


![Grafana dDashboard](/grafana/dashboard-screenshot-1.png)
![Grafana dDashboard](/grafana/dashboard-screenshot-2.png)
