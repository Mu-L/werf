{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
Generate werf environment variables for specified CI system.

Currently supported only GitLab (gitlab) and GitHub (github) CI systems

{{ header }} Syntax

```shell
werf ci-env CI_SYSTEM [options]
```

{{ header }} Examples

```shell
  # Load generated werf environment variables on GitLab job runner
  $ . $(werf ci-env gitlab --as-file)

  # Load generated werf environment variables on GitLab job runner using powershell
  $ Invoke-Expression -Command "werf ci-env gitlab --as-file --shell powershell" | Out-String -OutVariable WERF_CI_ENV_SCRIPT_PATH
  $ . $WERF_CI_ENV_SCRIPT_PATH.Trim()

  # Load generated werf environment variables on GitLab job runner using cmd.exe
  $ FOR /F "tokens=*" %g IN ('werf ci-env gitlab --as-file --shell cmdexe') do (SET WERF_CI_ENV_SCRIPT_PATH=%g)
  $ %WERF_CI_ENV_SCRIPT_PATH%
```

{{ header }} Options

```shell
      --as-env-file=false
            Create the .env file and print the path for sourcing (default $WERF_AS_ENV_FILE).
      --as-file=false
            Create the script and print the path for sourcing (default $WERF_AS_FILE).
      --config=''
            Use custom configuration file (default $WERF_CONFIG or werf.yaml in working directory)
      --config-templates-dir=''
            Custom configuration templates directory (default $WERF_CONFIG_TEMPLATES_DIR or .werf   
            in working directory)
      --dev=false
            Enable developer mode (default $WERF_DEV)
      --dir=''
            Use custom working directory (default $WERF_DIR or current directory)
      --docker-config=''
            Specify docker config directory path. Default $WERF_DOCKER_CONFIG or $DOCKER_CONFIG or  
            ~/.docker (in the order of priority)
            Command will copy specified or default (~/.docker) config to the temporary directory    
            and may perform additional login with new config.
      --env=''
            Use specified environment (default $WERF_ENV)
      --home-dir=''
            Use specified dir to store werf cache files and dirs (default $WERF_HOME or ~/.werf)
      --log-color-mode='auto'
            Set log color mode.
            Supported on, off and auto (based on the stdout’s file descriptor referring to a        
            terminal) modes.
            Default $WERF_LOG_COLOR_MODE or auto mode.
      --log-debug=false
            Enable debug (default $WERF_LOG_DEBUG).
      --log-pretty=true
            Enable emojis, auto line wrapping and log process border (default $WERF_LOG_PRETTY or   
            true).
      --log-quiet=false
            Disable explanatory output (default $WERF_LOG_QUIET).
      --log-terminal-width=-1
            Set log terminal width.
            Defaults to:
            * $WERF_LOG_TERMINAL_WIDTH
            * interactive terminal width or 140
      --log-verbose=false
            Enable verbose output (default $WERF_LOG_VERBOSE).
      --non-strict-giterminism-inspection=false
            Change some errors to warnings during giterminism inspection (more info                 
            https://werf.io/v1.2-alpha/documentation/advanced/configuration/giterminism.html,       
            default $WERF_NON_STRICT_GITERMINISM_INSPECTION)
  -o, --output-file-path=''
            Write to custom file (default $WERF_OUTPUT_FILE_PATH).
      --shell=''
            Set to cmdexe, powershell or use the default behaviour that is compatible with any unix 
            shell (default $WERF_SHELL).
      --tmp-dir=''
            Use specified dir to store tmp files and dirs (default $WERF_TMP_DIR or system tmp dir)
```

