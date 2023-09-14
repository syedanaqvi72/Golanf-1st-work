/*Go modules are a dependency management system introduced in Go (also known as Golang) starting from version 1.11. They were designed to solve the problem of versioning and dependency management in Go projects. Before modules, Go used a simpler approach that required developers to organize their code in a specific directory structure within the GOPATH, which made it challenging to manage dependencies and ensure reproducibility.

Here's what Go modules are and how they are used:

Dependency Management: Go modules provide a way to manage the dependencies (external packages or libraries) used in a Go project. They allow you to specify the exact version of a dependency that your project needs, ensuring that your project will work consistently, even if the dependencies change over time.

Module Initialization: To start using Go modules, you initialize your project as a Go module using the go mod init command. This command creates a go.mod file in your project directory, which contains information about your project and its dependencies.

Dependency Resolution: When you import packages in your code, Go modules automatically download and manage the required dependencies based on the versions specified in your go.mod file. It maintains a cache of downloaded dependencies to ensure efficient builds.

Versioning: Go modules use semantic versioning (semver) to specify dependency versions. You can pin a specific version, use a version range, or even use a commit hash from a source code repository to specify a dependency version.

Reproducibility: Go modules improve the reproducibility of builds because they make sure your project always uses the same versions of dependencies. This is crucial for ensuring that your application works consistently across different environments.

Dependency Upgrades: You can easily update your project's dependencies to their latest compatible versions using the go get or go get -u command.

Vendor Directory: Go modules also provide a vendor directory where you can copy specific versions of your dependencies. This can be useful for projects that need to work offline or in environments with restricted internet access.

Compatibility: Go modules are designed to be compatible with the older GOPATH-based workflow, so you can gradually transition your existing projects to use modules.

It is possible to replace the code of a module with the code of another module with the directive “replace” :

replace (
   gitlab.com/loir402/bluesodium v2.0.1 => gitlab.com/loir402/bluesodium2 v1.0.0
   gitlab.com/loir402/corge => ./corgeforked
)
The replaced version is at the left of the arrow; the replacement is at the right.

The replacement module can be :

Stored on a code sharing website (ex: Github, GitLab .…)

Stored locally

Some important notes :

The replacement module should have the same module directive (the first line of the go.mod file).

Should the replacement specify a version? It depends on the location of the replacement :

Distant (Github, Gitlab) : required

Local: not required

A specific version of a module can be replaced OR all versions can be replaced.