# Contributing to date-mcp

Thank you for your interest in contributing to date-mcp! This document provides guidelines and information for contributors.

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct. Please treat all community members with respect.

## Getting Started

### Prerequisites

- Go 1.24 or later
- Git
- Make (for using the Makefile)

### Setting Up Your Development Environment

1. Fork the repository on GitHub
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/date-mcp.git
   cd date-mcp
   ```
3. Add the upstream remote:
   ```bash
   git remote add upstream https://github.com/emj-io/date-mcp.git
   ```
4. Install dependencies:
   ```bash
   make deps
   ```
5. Run tests to ensure everything works:
   ```bash
   make test
   ```

## Development Workflow

### Making Changes

1. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
2. Make your changes
3. Add tests for your changes
4. Ensure all tests pass:
   ```bash
   make test
   ```
5. Run linting:
   ```bash
   make lint
   ```
6. Build the project:
   ```bash
   make build
   ```

### Commit Guidelines

- Use clear, descriptive commit messages
- Keep commits focused on a single change
- Reference issues in commit messages when applicable

Example commit messages:
- `feat: add timezone support to get_current_date tool`
- `fix: handle invalid date format parameters correctly`
- `docs: update README with new installation instructions`
- `test: add unit tests for date formatting edge cases`

### Pull Request Process

1. Update your branch with the latest upstream changes:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```
2. Push your branch to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
3. Open a Pull Request on GitHub
4. Fill out the PR template completely
5. Ensure all CI checks pass
6. Address any review feedback

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/tools/
```

### Adding Tests

- Add unit tests for all new functions
- Add integration tests for new MCP tools
- Ensure test coverage remains high
- Test both happy path and error cases

### Testing MCP Integration

When adding new MCP tools:

1. Test with a real MCP client (Cursor IDE recommended)
2. Verify tool descriptions are clear and helpful
3. Test all parameter combinations
4. Ensure error handling works correctly

## Code Style

### Go Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting (automated in CI)
- Follow the linting rules in `.golangci.yml`
- Write clear, self-documenting code
- Add comments for public APIs

### MCP Tool Guidelines

When implementing MCP tools:

- Use clear, descriptive tool names
- Provide comprehensive tool descriptions
- Include examples in descriptions when helpful
- Support appropriate parameter options
- Return well-structured responses
- Handle errors gracefully

## Project Structure

```
.
├── cmd/server/          # Main server entry point
├── internal/tools/      # MCP tool implementations
├── .github/            # GitHub workflows and templates
├── bin/               # Build output (gitignored)
├── Makefile           # Build automation
├── .golangci.yml      # Linting configuration
└── .goreleaser.yml    # Release configuration
```

## Release Process

Releases are automated via GitHub Actions:

1. Ensure your changes are merged to `main`
2. Create and push a version tag:
   ```bash
   git tag v1.0.0
   git push upstream v1.0.0
   ```
3. GitHub Actions will automatically build and release

## Getting Help

- Open an issue for bugs or feature requests
- Use GitHub Discussions for questions
- Check existing issues before creating new ones

## Security

If you discover a security vulnerability, please:

1. **Do not open a public issue**
2. Email the maintainers directly
3. Provide detailed information about the vulnerability
4. Allow time for the issue to be addressed before public disclosure

Thank you for contributing to date-mcp!