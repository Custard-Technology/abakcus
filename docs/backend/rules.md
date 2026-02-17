role: you are a senior engineer

Structure:
/cmd: Contains application-specific entry points, with a main.go file for each distinct application or service binary (e.g., cmd/api, cmd/worker).

/internal: Houses private application code and business logic that should not be imported by external projects. The Go compiler enforces this access restriction.

internal/handler: HTTP or RPC handlers that process requests and call the business logic.

internal/service: The core business logic layer, containing the application's main functions.

internal/repository: Handles interactions with data storage, such as databases or caches.

internal/domain or internal/models: Defines core domain entities and data structures.

/pkg: Contains public, reusable packages that can be safely used by other external projects (e.g., a client library or shared utilities).

/api: Holds API-related code, such as OpenAPI specifications, protocol buffer definitions, or generated client code.

/configs: Stores configuration files for different environments (development, production, etc.).

/scripts: Contains various build, installation, or deployment scripts (e.g., Makefiles).

/tests: Houses both unit and integration tests. 

rules: 
- Prioritize Readability and Simplicity: Write code that is easy to understand and maintain. Go favors short, clear naming conventions and minimal nesting.

- Use the Standard Library: Prefer the robust and well-optimized standard library over third-party packages when possible to minimize dependencies and potential security risks. 

- Handle Errors Explicitly: Never ignore returned errors. Check every error and handle it appropriately, logging details internally but avoiding exposure of sensitive information to users.

- Modularize and Document: Structure your project into well-defined, independent packages with clear boundaries. Document exported identifiers using godoc comments.

- Avoid Nesting: Handle errors early with a return statement to reduce cognitive load and avoid deep nesting.

- Validate and Sanitize Input: Never trust user input. Use strong validation rules and sanitize all inputs to prevent injection attacks (e.g., SQL injection, path traversal).

- Automate Code Analysis:
- Profile and Test

** NB:
when done with implementing task update always ./docs/notes.md
about what you did and why that was the best decision