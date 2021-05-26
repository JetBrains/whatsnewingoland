# What's New In 2021.2

This is a collection of new features in GoLand 2021.2.

Some of these are listed here, some have their own separate demo,
for the convenience of showing the feature itself.

- [Go 1.17 support](#go-117-support)
- [Run gofmt after the builtin formatter](#run-gofmt-after-the-builtin-formatter)
- [Go modules support](#go-modules-support)
  - [Load go.mod changes manually](#load-gomod-changes-manually)
  - [Better handling of Go versions from go.mod files](#better-handling-of-go-versions-from-gomod-files)
  - [Unused dependency in go.mod inspection](#unused-dependency-in-gomod-inspection)
- [Better view of attached projects](#better-view-of-attached-projects)
- [Web Features](#web-features)
  - [Browser pages reload on save](#browser-pages-reload-on-save)
  - [Completion for parameter types based on function calls](#completion-for-parameter-types-based-on-function-calls)
  - [New arrow function live template](#new-arrow-function-live-template)
  - [Code completion for private npm packages](#code-completion-for-private-npm-packages)
  - [Support for TypeScript types in JSDoc](#support-for-typescript-types-in-jsdoc)
  - [Support for web-types](#support-for-web-types)
  - [React support](#react-support)
    - [Rename refactoring for React useState hooks](#rename-refactoring-for-react-usestate-hooks)
    - [Support for classnames and clsx libraries](#support-for-classnames-and-clsx-libraries)
- [Version Control](#version-control)
    - [Sign commit with GPG](#sign-commit-with-gpg)
    - [Run custom Tests before commit](#run-custom-tests-before-commit)
    - [Run custom Inspections before commit](#run-custom-inspections-before-commit)
    - [Run custom Cleanup inspections before commit](#run-custom-cleanup-inspections-before-commit)
- [Docker](#docker)
  - [Quick documentation for keys in Dockerfiles and docker-compose.yaml files](#quick-documentation-for-keys-in-dockerfiles-and-docker-composeyaml-files)
  - [Restart/pause/unpause containers](#restartpauseunpause-containers)
  - [Docker Compose](#docker-compose)
    - [Inspection for ports mapping type in docker-compose](#inspection-for-ports-mapping-type-in-docker-compose)
    - [Support for operations on multiple services/containers](#support-for-operations-on-multiple-servicescontainers)
    - [Inspection for shm_size key values](#inspection-for-shm_size-key-values)
    - [Inspection for device_cgroup_rules values](#inspection-for-device_cgroup_rules-values)
- [Kubernetes](#kubernetes)
  - [Configure custom namespaces manually](#configure-custom-namespaces-manually)
- [Platform](#platform)
  - [Accessibility](#accessibility)
  - [Actions on save](#actions-on-save)
  - [Improvements for the Settings/Preferences dialog](#improvements-for-the-settingspreferences-dialog)
  - [Automatic cache and logs cleanup](#automatic-cache-and-logs-cleanup)
  - [Search for text in the Local History](#search-for-text-in-the-local-history)
  - [New Terminal options](#new-terminal-options)
  - [Select the shell of a new Terminal](#select-the-shell-of-a-new-terminal)
  - [Better project icon](#better-project-icon)
  - [Preview files when debugging](#preview-files-when-debugging)
  - [New Scratch File from selection](#new-scratch-file-from-selection)
  - [Remove Scratch File when closing it](#remove-scratch-file-when-closing-it) 

If you want, you can watch a recording of the presentation based on
this content:

[![What's New in GoLand 2021.2](https://img.youtube.com/vi/4JAmpWdPe-k/0.jpg)](https://www.youtube.com/watch?v=4JAmpWdPe-k "What's New in GoLand 2021.2")


## Go 1.17 support

  01. Convert a slice to an array pointer

  Go 1.17 brings a new language change:

  > Converting a slice to an array pointer yields a pointer to the underlying array of the slice. If the length of the slice is less than the length of the array, a run-time panic occurs.
  
  Source: [Go 1.17 release notes](https://tip.golang.org/ref/spec#Conversions_from_slice_to_array_pointer).

  [Demo](01%20-%20convert%20slice%20to%20array%20pointer/main.go)

## Run gofmt after the builtin formatter

  02. If you want to run gofmt, you can now do it in a fourth way: as a step after the builtin formatter.

  To use this feature, check the `On code reformat` option under `Settings/Preferences | Editor | Code Style | Go | Other`.

  Now you can use the default shortcut, `Ctrl + Alt + L` on Windows/Linux or `Cmd + Alt + L` on macOS, to trigger the builtin formatter.

  **Pro tip**: You can also search for `Run gofmt` in the IDE `Settings/Preferences`, and activate the option as
  described above.

  ![Run gofmt after builtin formatter.png](02%20-%20run%20gofmt%20after%20builtin%20formatter.png) 

## Go modules support

  ### Load go.mod changes manually

  03. Up until now, changes in the go.mod file where automatically detected. However, this could lead to situations
  where the update was triggered before a user could finish typing everything needed. Additionally, some users
  told us they prefer to launch the update process manually. We listened to this feedback and implemented a new way to manually trigger the updates after editing a go.mod file. To use it,
  navigate to `Settings/Preferences | Build, Execution, Deployment | Build Tools` and select `External changes` instead of `Any changes`.

  ![Load go.mod changes manually](03%20-%20load%20go%20mod%20changes%20manually/03%20-%20load%20go%20mod%20changes%20manually%20-%201.png)

  After this, you'll see a new option appearing on the top right of the editor after making any changes in the file.

  ![Trigger go.mod changes manually](03%20-%20load%20go%20mod%20changes%20manually/03%20-%20load%20go%20mod%20changes%20manually%20-%202.png)

  [Demo](03%20-%20load%20go%20mod%20changes%20manually/go.mod)

  ### Better handling of Go versions from go.mod files

  04. Until now, the version Go SDK used in the project was the one that GoLand took into account when performing inspections
  or other tasks. However, this is not always correct, as the `go` directive in the `go.mod` file might specify a different
  version. That's why, since this release, GoLand supports looking at the Go version from `go.mod` files rather than just the
  project SDK.

  [Demo](04%20-%20better%20handling%20of%20go%20version%20from%20go%20mod%20files/main.go)

## Unused dependency in go.mod inspection

  05. You can now determine the unused module dependencies in `go.mod` files thanks to a new inspection that marks them with a grey color.

  [Demo](05%20-%20unused%20dependendencies%20in%20gomod%20files/go.mod)

## Better view of attached projects

  06. Working with multiple projects at the same time is now cleaner thanks to the improved display of attached projects.
  To attach a project to an existing one, use the Open command and select the "Attach" option in the follow-up dialog.
  ![Attach project dialog](06%20-%20attach%20project%20dialog.png)

  After this, you can work with the new project alongside the existing one.
  ![Project view with attached project](06%20-%20attach%20project%20view.png)

  When you finish with the attached project, you can detach it using the `right-click | Remove from Project View` action.
  ![Detach an attached project](06%20-%20attach%20project%20detach.png)

## Web Features

### Browser pages reload on save

  07. The IDE can automatically reload the page opened in a browser after you make changes to an HTML file, or the linked CSS and JavaScript files. Reloading is triggered whenever changes are saved in the IDE, either automatically or using _Ctrl + S on Windows/Linux_ or _Cmd + S on macOS_, or when changes are made to a file externally.

  To use this, open the file from the IDE via the icon on the upper right-side of the editor or via `View | Open in Browser | <browser name>`.

  [Demo](07%20-%20automatic%20browser%20files%20reload/index.html)

### Completion for parameter types based on function calls

  08. The IDE will show you completion suggestions based on the type used in the function call for the parameters in the function body.

  [Demo](08%20-%20js%20completion%20for%20parameter%20types%20based%20on%20function%20calls/demo.js)

  **Pro tip:** Use the Smart Type Completion to narrow-down the list further.

### New arrow function Live Template

  09. Writing a new arrow function is now even faster thanks to the new `arf` Live Template.

  [Demo](09%20-%20new%20arrow%20function%20live%20template/demo.js)

### Code completion for private npm packages

  10. Code completion now works for private npm packages in `package.json` files. The IDE will let you browse information about the latest versions of the package, similarly to the public packages.

### Support for TypeScript types in JSDoc

  11. You can now use the TypeScript syntax within JSDoc comments in .js files.

  [Demo](11%20-%20typescript%20type%20references%20in%20jsdoc/ref.js)

### Support for web-types

  12. Web-types is an open source standard for documenting various web frameworks. It's framework-agnostic and provides IDEs and other tools with metadata information about contents of a web components library.
  Web-types are developed by JetBrains and work on this format started a few months prior to this release. You might already see the effects of this when using the Vue.js plugin.  
  To learn more about web-types, check out [this article](https://blog.jetbrains.com/webstorm/2021/01/web-types/).

### React support

#### Rename refactoring for React useState hooks

  13. When using the `Rename refactoring`, Shift + F6 on Windows/Linux or  ⇧ + F6 on macOS, you it will detect if you are using useStates from React and rename both values at once.

  [Demo](13%20-%20rename%20refactoring%20for%20react%20usestate%20hooks/demo.js)

#### Support for classnames and clsx libraries

  14. When using the `classnames` and `clsx` libraries, the IDE will show completion suggestions for your CSS classes and resolve all symbols in string literals and properties with literal names.

  [Demo](14%20-%20support%20for%20classnames%20and%20clsx%20libraries/demo.js)

## Version Control

### Sign commit with GPG

  15. Signing a commit with a GPG key is now possible.
  After configuring the git client with the correct key, head over to `Settings/Preferences | Version Control | Git`
  and ensure that the option `Configure GPG Key...` is turned on and points to the expected key.

  ![Sign commit with GPG key](15%20-%20sign%20commit%20with%20gpg%20key.png)

### Run custom Tests before commit

  16. You can now configure a test to run before committing your files. Invoke the Commit feature using Ctrl + K on Windows/Linux and Cmd + K on macOS,
  then select the `Commit options` and check the `Run Tests` option and select which run configuration you wish to run.

  **Note:** You must have the `Use non-modal commit interface` activated before you can use this feature. Check your settings under `Settings/Preferences | Version Control | Commit`.

  ![Run tests before commit](16%20-%20run%20tests%20before%20commit.png)

  **Pro-tip:** You can use the [Compound configuration type](https://www.jetbrains.com/help/go/run-debug-configuration.html#compound-configs) to run more than a single run configuration.

  [Demo](16%20-%20run%20tests%20before%20commit/main_test.go)

### Run custom Inspections before commit

  17. Committing error-free code reduces the time it takes for reviews to go happen and makes everyone happy. Since not all errors may be caught by tests, you can choose to run Inspections before a commit. This will improve your chances to have fewer bugs in code.

  To use this feature, invoke the `Commit` feature, then select the `Analyze code` feature. 

  ![Run inspections before commit.png](17%20-%20run%20inspections%20before%20commit.png)

### Run custom Cleanup inspections before commit

  18. Tests and Inspections are not the only features that you can run before a commit. You can also trigger an automatic code cleanup,
  using the `Cleanup` feature from the commit dialog.

  ![Code cleanup before commit](18%20-%20code%20cleanup%20before%20commit.png)

## Docker

### Quick documentation for keys in Dockerfiles and docker-compose.yaml files

  19. Dockerfile and docker-compose.yaml files now support displaying documentation references via the Quick Documentation, Ctrl + Q on Windows/Linux or F1 on macOS.

  [Demo](19%20-%20docker%20quick%20documentation/Dockerfile)

### Restart/pause/unpause containers

  20. Containers can now be paused/resumed or restarted from the Services view.

  [Demo](20%20-%20docker%20pause%20container)

### Docker Compose

#### Inspection for ports mapping type in docker-compose

  21. A new inspection checks if the ports mapping in docker-compose.yaml files is correct.

  [Demo](21%20-%20docker-compose%20ports%20mapping%20validation/docker-compose.yaml)

#### Support for operations on multiple services/containers

  22. You can now select multiple containers or Docker Compose Services from the Services view, Alt + 8 on Windows/Linux or Cmd + 8 on macOS,
  and perform common operations, such as start, stop, restart, pause, or delete on them.
  
  [Demo](22%20-%20docker-compose%20handle%20multiple%20containers/docker-compose.yaml)

#### Inspection for shm_size key values

  23. A new inspection for docker-compose files will tell you when the `shm_size` values are incorrect.

  [Demo](23%20-%20docker-shm-size-inspection/docker-compose.yaml)

#### Inspection for device_cgroup_rules values

  24. When using device_cgroup_rules, the value types are now correctly inspected, and the IDE will inform you if 
  you are not using a string, or a list of strings.

  [Demo](24%20-%20docker-device-cgroup-rules-inspection/docker-compose.yaml)

## Kubernetes

### Configure custom namespaces manually

  25. If the plugin cannot automatically infer the available Kubernetes namespaces, now you can configure them manually.
  This will solve the problem when the user does not have the required permissions to perform namespace listing.
  To use this feature, head over to `Settings/Preferences | Build, Execution, Deployment | Kubernetes` and configure the
  namespace in the corresponding section.
  
  ![Configure Kubernetes namespaces manually](25%20-%20configure%20kubernetes%20namespaces%20manually.png)

## Platform

### Accessibility
  
  26. We reduced the number of screen read help tooltips.

  A list of code completion options will be read out loud on macOS. The content of the selected combo box and combo box lists was added too.

  The results in Search Everywhere are now properly pronounced.

### Actions on save

  27. It's now possible to trigger various IDE actions on save. Before, you could use File Watchers to achieve a part of this functionality, but IDE actions such as Optimize Imports, or using the builtin code formatter were not possible.
  Starting with 2021.2, you can turn on such actions under `Settings/Preferences | Tools | Actions on Save`.

  ![Run actions on save.png](27%20-%20actions%20on%20save.png)

### Improvements for the Settings/Preferences dialog

  28. We added a new node called `Advanced Settings` in `Settings/Preferences` dialog. Most of the options there were transferred from the `Registry...` as we wanted to make those options easier to access.

  You'll also find some new configuration options there, such as the ability to set a left margin in `Distraction-free mode`.

  ![Advanced settings](28%20-%20advanced%20settings.png)

### Automatic cache and logs cleanup

  29. The IDE will automatically clean up any cache and log directories that were last updated more than 180 days ago. The system settings and plugins directories will stay intact. To trigger this process manually, go to `Help | Delete Leftover IDE Directories...`.

  ![Delete leftover IDE directories](29%20-%20delete%20leftover%20ide%20directories.png)

### Search for text in the Local History

  30. With the new Search field to the `Local History` dialog you quickly get to the necessary text in your revisions. To invoke it, press the `Find` shortcut, Ctrl + F on Windows/Linux or Cmd + F on macOS.

  ![Search text in local history](30%20-%20search%20text%20in%20local%20history.png)

### New Terminal options

  31. You can now select the cursor shape in the builtin Terminal.

  And, if you are using macOS, you can now use `Option as Meta key`, similar to the same-name option in the native Terminal on macOS. This allows the `⌥` key on the keyboard to act as a meta modifier that can be used in combination with other keys instead of just as an Escape key. For example, you can use the following shortcuts:

    ⌥ + F – go to the next word
    ⌥ + B – go back a word
    ⌥ + D – delete the next word

  You can tick the checkboxes for the new options in Settings/Preferences | Tools | Terminal.

  ![Terminal cursor shape](31%20-%20terminal%20cursor%20shape.png)

### Select the shell of a new Terminal

  32. Up until now, you could configure which shell you want to use only from the `Settings/Preferences | Tools | Terminal`. Since this was not flexible enough, we introduced a new way to create new Terminal tabs. Click on the `⋁` button next to the existing tabs and select which shell to use. You can also create a new SSH session from it, if you have any servers configured or configure a new one.

  ![Select the shell of a new Terminal](32%20-%20new%20terminal%20shell.png)

### Better project icon

  33. You can add a project icon to the project from the projects list on the Welcome screen. Upload custom icons by right-clicking on any project and select `Choose Project Icon...` from the context menu.

  ![Better project icon](33%20-%20better%20project%20icon.png)

### Preview files when debugging

  34. The `Preview tab` was extended to work with the files that open during debugging. To activate this feature, navigate to `Settings/Preferences | Editor | General | Editor Tabs`.

  When using this feature, you'll be able to preview files without worrying about creating a new tab/closing them after/losing your existing tabs.

### New Scratch File from selection

  35. You can now select some code, press Alt+Enter on Windows/Linux or ⌥ + ⏎ on macOS, and then choose `Create new scratch file from selection` to quickly create a new scratch containing the same code.

### Remove Scratch File when closing it

  36. When you close an empty scratch file, it will automatically be removed.