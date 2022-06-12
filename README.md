# Helper binary to install xkb keyboard layouts

This script notably helps to install a new ruleset within
the only file that actually gets parsed by the system to
list the available layouts.

## Usage

## Features

- Idempotent
- Makes a layout available in system dropdowns

Extra details:
- Install an extra `XkbSymbols` table in the `/usr/share/X11/xkb/symbols/XX` file
- Add extra variants in the `evdev.lst` file under the `XX` symbol
- Add extra variants in the `evdev.xml` file under the `XX` layout/symbol
- Install an optional `.Xcompose` file in the home of `$USER`
