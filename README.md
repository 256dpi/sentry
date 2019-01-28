# sentry

[![Build Status](https://travis-ci.org/256dpi/sentry.svg?branch=master)](https://travis-ci.org/256dpi/sentry)
[![Release](https://img.shields.io/github/release/256dpi/sentry.svg)](https://github.com/256dpi/sentry/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/256dpi/sentry)](http://goreportcard.com/report/256dpi/sentry)

A small tool to upload logged errors and crash reports of failing programs. The tool does not process Stderr line by line and instead uploads the written chunks in their original length in the hope of getting a complete dump of crashing programs.

## Usage

Provide `SENTRY_DSN` as an environment variable. Additionally, use `FILTER` to provide a semicolon separated list of Go regex patterns to filter out unwanted messages.

```bash
export SENTRY_DSN=THE_SENTRY_DSN
export FILTER="do not report this"
sentry my-app --some arguments
```
