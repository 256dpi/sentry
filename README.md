# sentry

[![Build Status](https://travis-ci.org/256dpi/sentry.svg?branch=master)](https://travis-ci.org/256dpi/sentry)
[![Release](https://img.shields.io/github/release/256dpi/sentry.svg)](https://github.com/256dpi/sentry/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/256dpi/sentry)](http://goreportcard.com/report/256dpi/sentry)

A small tool to upload crash reports of failing programs. The tool will collect the Stderr output and upload it if the program exists with an abnormal status.

## Usage

```bash
export SENTRY_DSN=THE_SENTRY_DSN
sentry crashing-program --some arguments
```
