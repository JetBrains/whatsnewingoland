# What's New In 2021.1

This is a collection of new features in GoLand 2021.1.

Some of these are listed here, some have their own separate demo,
for the convenience of showing the feature itself.

- [Run Targets support](#run-targets-support)
- [Run Configurations](#run-configurations)
- [Go 1.16 support](#go-116-support)
- [JSON support](#json-support)
- [Error handling](#error-handling)
- [Quick-fixes improvements](#quick-fixes-improvements)
- [Refactorings](#refactorings)
- [Go Modules support](#go-modules-support)
- [Formatter](#formatter)
- [Completion](#completion)
- [UI improvements](#ui-improvements)
- [Plugins](#plugins)
    - [Makefile](#makefile)
    - [Docker](#docker)
    - [Kubernetes](#kubernetes)
    - [HTTP Client](#http-client)
- [Web Features](#web-features)
- [Version Control](#version-control)
- [Database support](#database-support)

If you want, you can watch a recording of the presentation based on
this content:

[![What's New in GoLand 2021.1](https://img.youtube.com/vi/3krpWjdisGo/0.jpg)](https://www.youtube.com/watch?v=3krpWjdisGo "What's New in GoLand 2021.1")

## Run Targets support

01. Run Targets support is here!

    You can now _run_ and _debug_ your application and its tests
    (including coverage reports, and profiling benchmarks)
    in a Docker Container, a WSL 2 environment, or via an SSH
    connection on a remote server/device.
    
    [Demo](01%20-%20run%20targets/main.go)

## Run Configurations

02. Run with elevated privileges service was introduced to better handle
    scenarios for applications that requires administrator level access.

    [Demo](02%20-%20run%20with%20sudo/main.go)    

## Go 1.16 support

03. `//go:embed` support
    [Demo](03%20-%20go-embed/main.go)
    
04. Incorrect usage of (t/b).Fatal* methods
    [Demo](04%20-%20incorrect%20tb-fatal/main_test.go)
    
05. Detect incorrect usages of asn1.Unmarshal
    [Demo](05%20-%20detect%20incorrect%20asn1%20unmarshal/main.go)
    
## JSON support

06. Generate json tags for struct types
    [Demo](06%20-%20json%20tags/main.go)
    
07. Copy JSON and paste Go
    [Demo](07%20-%20copy%20json%20paste%20go/main.go)
    
08. [JSON Lines](https://jsonlines.org) support
    [Demo](08%20-%20json%20lines/demo.jsonl)
    
## Error handling

09. Faster error handling thanks to new quick-fixes
    [Demo](09%20-%20faster%20error%20handling/main.go)
    
## Quick-fixes improvements

10. Improved create type quick-fix to now add the fields too
    [Demo](10%20-%20create%20type%20adds%20fields/main.go)
    
11. Improved type conversion quick-fix
    [Demo](11%20-%20better%20type%20conversion/main.go)
    
12. Mass generate getters/setters for structs
    [Demo](12%20-%20mass%20generate%20getters%20setters/main.go)

## Refactorings

13. Introduce Type refactoring converts an anonymous type to a named one.
    [Demo](13%20-%20introduce%20type/main.go)
    
## Go Modules support

14. Rename Go Modules with ease, from the `go.mod` file. 

    When to use it? Whenever you want to upgrade your module's major
    version to a newer one, or you want to change the location of your import.

    [Demo](14%20-%20rename%20go%20modules/go.mod)
    
## Formatter

15. The built-in formatter can now mirror the `goimports -local` behavior.
    
    This means you can stop relying on the File Watchers to run
    against your code and use the builtin tools to ensure consistent formatting across the code.
    
    [Demo](15%20-%20goimports%20local/cmd/demo/main.go)
    
## Completion

16. Machine Learning Assisted Completion is here.
    
    When enabled, code completion results are sorted based on rules learned from data we have gathered anonymously during our EAPs.

    If you want to see how this affects your results, enable [Settings/Preferences | Editor | General | Code Completion | Mark position changes](jetbrains://GoLand/settings?name=Editor--General--Code+Completion). This will display tiny up/down arrow icons in the completion list.

    **Note:** No source code was collected from our users for this feature. The collected information was about interactions with the code completion UI.
    
17. A new Postfix Completion `varCheckError` to handle errors faster.
    
    [Demo](17%20-%20varcheckerror%20postfix%20completion/main.go)
    
## UI improvements

18. Improved accessibility

    We added more labels to UI elements on the Welcome screen, the Project Structure view, and the VCS log.
    
    These elements and gutter icons are now read out correctly when a screen reader is enabled. 

19. Font weight (variations) support

    You can now select the font weight (variations) from [Settings/Preferences | Editor | Font](jetbrains://GoLand/settings?name=Editor--Font).

20. Windows 10 Recent list support

    You can now access your recent projects by right-clicking on the GoLand icon on the taskbar or on the Start menu in Windows.

    ![Windows 10 Recent Project List](Windows%20Recent%20Project%20List.png)

21. Maximize editor tabs in the split view

    Double-click the tab you're working with to maximize the editor window for it.
    To revert the tab to its original size, double-click it again.

## Plugins

### Makefile

22. Makefile Language plugin is now bundled with the IDE

    This brings native support for syntax highlighting, quick documentation, Find Usages for targets, and some navigation and code completion actions for Makefile.

    [Demo](22%20-%20Makefile/Makefile)

### Docker

23. Dockerfile editing improvements

    We've made a bunch of improvements, from folding of multi-stage containers
    to completion for image names. Macros can be used in Docker mounts.
    
    [Demo](23%20-%20Dockerfile%20improvements/Dockerfile)

24. BuildKit support is now supported for Docker run configurations. 
    This unlocks new build options, such as build mounts, network modes, and more.
    You can learn about these [here](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md).
    
    [Demo](24-BuildKit-support/Dockerfile)

### Kubernetes

25. Quickly delete resources from the Run icon

    [Demo](25%20-%20kubernetes%20-%20delete/db.yaml)
    
26. Quickly convert a multi-document YAML to a List

    [Demo](26%20-%20kubernetes%20-%20list/db.yaml)

27. Kustomize 3.7.0 components are now supported


### HTTP Client

28. The HTTP Client has a few new features
    - you can fold the response 
    - copy the response body to clipboard
    - a new inspection helps you convert your requests to use HTTPS
    
    [Demo](28%20-%20http%20client/main.go)

29. It is now possible to use custom SSL Client Authentication support
    by clicking `Add environment file` and selecting `Private`.


## Web Features

30. Svelte support improvements
    
    The Svelte plugin author is now part of our team,
    and we push forward with more changes for this based on your feedback.
    
    [Demo](30%20-%20svelte/App.svelte)

31. You can now preview HTML and other template files using the new
    built-in preview feature which uses Java Chromium Embedded Framework (JCEF)
    
    [Demo](31%20-%20builtin%20jcef%20preview/demo.html)

32. You can configure the ESLint scope via 
    [Settings/Preferences | Languages & Frameworks | JavaScript | Code Quality Tools | ESLint](jetbrains://GoLand/settings?name=Languages+%26+Frameworks--JavaScript--Code+Quality+Tools--ESLint).

33. JSDoc syntax highlighting has been improved.
    
    You can find new options for changing the color scheme for JSDoc elements,
    including JSDoc tag values and types, in
    [Settings/Preferences | Editor | Color Scheme | JavaScript/TypeScript](jetbrains://GoLand/settings?name=Editor--Color+Scheme--JavaScript).

34. You can now suppress inspections for JavaScript and TypeScript files on a per-file basis

    [Demo](34%20-%20supress%20js%20inspections/demo.mjs)

### Version Control

35. Support for Git commit message templates

36. Faster creation of Pull Requests for GitHub

37. Configure a profile for pre-commit inspections

## Database support

38. Improved databases support:
    
    - Support for the CockroachDB dialect.
    - Full support for Google Big Query.
    - Redshift driver 2.x improvements (you can now cancel running queries).
    - Support for Azure MFA.

39. The connection window has several new improvements in layout.

40. A new UI for grants is available for PostgreSQL, Redshift, Greenplum,
    MySQL, MariaDB, DB2, SQL Server, and Sybase.

41. Data in MongoDB collections can now be edited from the IDE.
