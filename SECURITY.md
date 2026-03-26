# Security Policy

## Supported Versions

The project is pre-1.0. Security fixes target the latest `main` branch state.

## Reporting A Vulnerability

Do not open public GitHub issues for suspected security vulnerabilities that could expose users or maintainers.

Until a dedicated private reporting channel is published, report security issues directly to the maintainers through a private channel you already have or by opening a minimal public issue that requests a private contact path without disclosing exploit details.

## Secret Handling

- Never commit credentials, API tokens, passwords, private keys, or internal URLs.
- Keep local runtime configuration in ignored files such as `.env` or tool-specific local config directories.
- If a secret is committed accidentally, rotate it first and then remove it from the repository history.
