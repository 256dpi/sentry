# sentry

A small tool to upload crash reports of failing programs. The tool will collect the Stderr output and upload it if the program exists with an abnormal status.

## Usage

```bash
export SENTRY_DSN=THE_SENTRY_DSN
sentry crashing-program --some arguments
```
