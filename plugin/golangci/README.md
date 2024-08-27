# Plugin for golangci-lint
[Interfaceguard](https://github.com/jkeys089/interfaceguard) plugin for golangci-lint

## Example Config
`.golangci.yml`
```yaml
linters-settings:
  custom:
    interfaceguard:
      type: module
      description: interfaceguard checks subtle interface-related issues, including improper comparisons and type mismatches.
      original-url: https://github.com/jkeys089/interfaceguard
      settings:
        # Disable check for direct interface-to-interface comparisons.
        # Default: false
        disable-interface-to-interface: true
        # Disable check for direct interface-to-nil comparisons.
        # Default: false
        disable-interface-to-nil: true

linters:
  disable-all: true
  enable:
    - interfaceguard
```

## How To Run
__Preferred Method__ - utilize `golangci-lint custom` command:
1. create a `.custom-gcl.yml` config like below (see [reference](https://golangci-lint.run/plugins/module-plugins/#reference) for all options):
   ```yaml
   version: v1.60.3
   plugins:
     - module: 'github.com/jkeys089/interfaceguard/plugin/golangci'
       import: 'github.com/jkeys089/interfaceguard/plugin/golangci/module'
       version: latest
   ```
2. build and run custom-gcl
   ```shell
   golangci-lint custom
   ./custom-gcl linters | grep interfaceguard
   ./custom-gcl run ./...
   ```

__Alternative Method__ - utilize the go [plugin](https://pkg.go.dev/plugin) package (requires golangci-lint be built with `CGO_ENABLED` and only supported on Linux, FreeBSD, and macOS):
1. build the plugin for your system:
   ```shell
   git clone https://github.com/jkeys089/interfaceguard
   cd interfaceguard/plugin/golangci
   go build -o interfaceguard.so -buildmode=plugin plugin.go
   ```
2. update `linters-settings.custom.interfaceguard` in `.golangci.yml` to use `path` param instead of `type: module`:
   ```yaml
   linters-settings:
     custom:
       interfaceguard:
         path: /path/to/interfaceguard.so
         description: interfaceguard checks subtle interface-related issues, including improper comparisons and type mismatches.
         original-url: https://github.com/jkeys089/interfaceguard
         settings:
           # Disable check for direct interface-to-interface comparisons.
           # Default: false
           disable-interface-to-interface: true
           # Disable check for direct interface-to-nil comparisons.
           # Default: false
           disable-interface-to-nil: true

   linters:
     disable-all: true
     enable:
       - interfaceguard
   ```
3.  run golangci-lint as normal
   ```shell
   golangci-lint linters | grep interfaceguard
   golangci-lint run ./...
   ```