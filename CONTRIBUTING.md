# 🤝 Contributing to dev-helper

We welcome contributions to dev-helper! Whether it's reporting a bug, suggesting a feature, or improving the codebase, your help is greatly appreciated.

## 🚀 How to Get Started

1. **Fork** the repository.
2. **Create** a new branch (`git checkout -b feature/AmazingFeature`).
3. **Commit** your changes (`git add . && git commit -m 'feat: Add AmazingFeature'`).
4. **Push** to the branch (`git push origin feature/AmazingFeature`).
5. **Open a Pull Request** against the `main` branch.

## 🎨 Code Vision & Standards

We aim for high code quality and consistency across all components. Please follow these guidelines:

*   **Idiomatic Go/Language:** Follow best practices for the language you are contributing to (e.g., Go conventions, Python PEP 8).
*   **Testing is Mandatory:** Every new feature or significant modification **must** include corresponding unit tests. The CI pipeline checks for test coverage, and we highly encourage implementing more robust test cases, especially for core packages like `core/config.go` and `core/scaffolder.go`.
*   **Documentation:** Ensure all exported functions, classes, and public interfaces (especially in `plugins/interface.go`) include clear godoc/docstring comments explaining their purpose, arguments, and return values.

## 🧪 Testing & Validation

We utilize a strict testing pipeline (see CI workflow) and encourage developers to:
1.  **Implement Unit Tests:** Cover logic within individual files.
2.  **Implement Integration Tests:** Test how components interact, particularly for the TUI Wizard flow (e.g., testing the selection sequence in `tui/wizard.go`).

## 🛠️ Developing New Functionality

When adding a major feature:
1.  Use the task system to plan out multiple dependent steps.
2.  Break down the task into logical components.
3.  Use the `code-agent` to work on code, assisted by the agent workflow.
4.  After implementation, run `make test` and `make lint` to ensure quality.
5.  If successful, create a clean commit message explaining the "why."
