# Yet Another NodeJS Package Manager (YPM)

WARNING: DO NOT USE!! This is still a work-in-progress.

`YPM` is not a normal package manager. It is a batteries-included, auto-detecting, modular package managing system. The primary goal is to make interworking applications within a single repo automatic and intuitive to integrate and work with.

What does this mean? Well, unlike the other package managers, you can use `YPM` for a single repo or a monorepo, out of the box. No configuration is needed to support either use case. The only thing required to detect another workspace, automatically, is the existence of a `package.json` file in a directory and a `name` property to identify one workspace from another by their names.

## Why Go?

Go has proven to not only be a very performant general purpose system programming language, but it builds a static binary, and it is easy to learn and makes building performant solutions quite fun. There are other benefits like being a function-oriented programming language, however, the biggest gains are that reduction is execution overhead and memory usage compared to the NodeJS runtime. Traditional JS/TS solutions use and (normally) require NodeJS to manage packages, which means they do not have the same level of control, performance, or memory efficiency as Go has. The fundamental idea was to be able to remove the package manager as a performance bottleneck, extend it beyond NodeJS, and leverage the performance gains we get from using Go. An example of a successful use case for this is [ESBuild](https://github.com/evanw/esbuild), which is currently one of the fastest growing JS/TS bundlers in the ecosystem, known for its performance, and written in Go. I hope to achieve the same goals here, but as a package manager.

## Commands

There are a few reserved keywords for `ypm`. Currently, those are:

| Keyword(s) | Description                                                                                                                                                                                                                                                                                    |
| :--------- | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `repo`     | This command is used to create the initial repo for YPM. This is important because YPM uses metadata to keep track of what workspace(s), if any, do exist and manages those automatically.                                                                                                     |
| `init`     | This is similar to the `repo` command, except it just generates the base `package.json` and auto-links the workspace during the creation. This is NOT required for `YPM` to detect your added workspace. If a `package.json` is created by hand, that will automatically be detected by `YPM`. |
| `add`      | Adds a desired package and contains flags for different types of packages (i.e devDependencies, etc).                                                                                                                                                                                          |
| `remove`   | Removes a given package from the configuration and uninstalls it from your project.                                                                                                                                                                                                            |
| `wksp`     | Invokes something under the strict context of a workspace (not require, but may be needed if a package name conflicts with a YPM keyword).                                                                                                                                                     |

---

This creates a `package.json` with a specific name, with the default version of `0.0.1`, and generates the `YPM` configuration directory `.ypm`.

```bash
ypm repo <name>
```

This does the same thing as the other command but assigns it to a specific version.

```bash
ypm repo <name>@<version>
```

---

This creates a `package.json` with the `name` that was passed in and a default version of `0.0.1`.

```bash
ypm init <name>
```

This creates a `package.json` file with both the `name` and `version` set to whatever is passed in.

```bash
ypm init <name>@<version>
```

---

## WIP

This is how you can invoke a script from a workspace by its name.

```bash
ypm @myproject/app dev
```

If there is a namespace conflict between a package name and a reserved keyword in `ypm` you can accomplish the same result as the example above doing this.

```bash
ypm wksp @myproject/app dev
```
