# Docker Flow Proxy

## Fork

This fork updates versions of HAProxy, Go and Alpine to the latest versions.
Some tests skipped that wouldn't run inside the Docker builder, and Docker Autotests disabled in order to get it automatically building in Docker Hub.

## Original

**This project needs adoption. I (@vfarcic) moved to Kubernetes and cannot dedicate time to this project anymore. Similarly, involvement from other contributors dropped as well. Please consider contributing yourself if you think this project is useful.**

[![GitHub release](https://img.shields.io/github/release/docker-flow/docker-flow-proxy.svg)]()
[![license](https://img.shields.io/github/license/docker-flow/docker-flow-proxy.svg)]()
[![Docker Pulls](https://img.shields.io/docker/pulls/vfarcic/docker-flow-proxy.svg)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/docker-flow/docker-flow-proxy)](https://goreportcard.com/report/github.com/docker-flow/docker-flow-proxy)

The goal of the *Docker Flow Proxy* project is to provide an easy way to reconfigure proxy every time a new service is deployed, or when a service is scaled. It does not try to "reinvent the wheel", but to leverage the existing leaders and combine them through an easy to use integration. It uses [HAProxy](http://www.haproxy.org/) as a proxy and adds custom logic that allows on-demand reconfiguration.

Supported archetectures are:

- linux-amd64
- linux-arm

Please visit the **[project documentation](http://proxy.dockerflow.com)** for more info or join the #df-proxy Slack channel in [DevOps20](http://slack.devops20toolkit.com/) if you have any questions, suggestions, or problems.

<a href='https://ko-fi.com/A655LRB' target='_blank'><img height='36' style='border:0px;height:36px;' src='https://az743702.vo.msecnd.net/cdn/kofi2.png?v=0' border='0' alt='Buy Me a Coffee at ko-fi.com' /></a>
