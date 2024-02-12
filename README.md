# Yet Another NodeJS Package Manager (YPM)

WARNING: DO NOT USE!! This is still a work-in-progress.

`YPM` is not a normal package manager. It is a batteries-included, auto-detecting, modular package managing system. The primary goal is to make interworking applications within a single repo automatic and intuitive to integrate and work with.

What does this mean? Well, unlike the other package managers, you can use `YPM` for a single repo or a monorepo, out of the box. No configuration is needed to support either use case. The only think required to detect another workspace, automatically, is the existence of a `package.json` file in a directory and a `name` property to identify one workspace from another by their names.

## Commands

There are a few reserved keywords for `ypm`. Currently, those are:

1. `repo`

- This command is used to create the initial repo for YPM. This is important because YPM uses metadata to keep track of what workspace(s), if any, do exist and manages those automatically.

2. `init`

- This is similar to the `ypm repo` command, except it just generates the base `package.json` and auto-links the workspace during the creation.

3. `add`

- Adds a desired package and contains flags for different types of packages (i.e devDependencies, etc).

4. `remove`

- Removes a given package from the configuration and uninstalls it from your project.

5. `wksp`

- Invokes something under the strict context of a workspace (not require, but may be needed if a package name conflicts with a YPM keyword)

These reserved words (if your project matches these works) will need to be invoked using the `wksp` namespace to invoke a workspace, by name, that conflicts with a reserved keyword.

---

```bash
ypm repo <name>
```

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
