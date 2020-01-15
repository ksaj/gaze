/**
 * Gaze (https://github.com/wtetsu/gaze/)
 * Copyright 2020-present wtetsu
 * Licensed under MIT
 */

package config

// Default returns the default configuration
func Default() string {
	return `# Gaze configuration(priority: default < ~/.gaze.yml < ./.gaze.yaml < -f option)
commands:
- ext: .go
  run: go run "{{file}}"
- ext: .py
  run: python "{{file}}"
- ext: .rb
  run: ruby "{{file}}"
- ext: .js
  run: node "{{file}}"
- ext: .d
  run: dmd -run "{{file}}"
- ext: .groovy
  run: groovy "{{file}}"
- ext: .php
  run: php "{{file}}"
- ext: .pl
  run: perl "{{file}}"
- ext: .java
  run: java "{{file}}"
- ext: .kts
  run: kotlinc -script "{{file}}"
- re: ^Dockerfile$
  run: docker build -f "{{file}}" .
`
}