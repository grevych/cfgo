# cfgo
Configure Go apps, merge config files depending on the environment. This package mimics [node-config](https://github.com/lorenwest/node-config) module using Viper [Viper](github.com/spf13/viper)

### Requirements
 - Viper

### Usage
Expected default config file:
 - config/default.yaml

Suggested config files:
 - config/local.yaml
 - config/develop.yaml
 - config/test.yaml
 - config/staging.yaml
 - config/production.yaml

NOTE: Configuration files could be in any format supported by viper


For an example, see the directory `examples/`:
