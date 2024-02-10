# Yet Another NodeJS Package Manager

`YPM` is not a normal package manager. It is a batteries-included, auto-detecting, modular package managing system.

What does this mean? Well, unlike the other package managers, you can use `YPM` for a single repo or a monorepo, out of the box. No configuration is needed to support either use case. The only think required to detect another workspace, automatically, is the existence of a `package.json` file in a directory and a `name` property to identify one workspace from another by their names.

## Commands

| Command                                     | Description                                                                                |
| :------------------------------------------ | :----------------------------------------------------------------------------------------- |
| `ypm init <project_name>`                   | Creates a `package.json` file with the name specified and defaults the version to `0.0.1`. |
| `ypm init <project_name>@<project_version>` | Creates a `package.json` file with the name and version specified.                         |
