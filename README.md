# Helper binary to install xkb keyboard layouts

This script notably helps to install a new ruleset within
the only file that actually gets parsed by the system to
list the available layouts.

## Usage

Once the binanry gets properly installed (from a binary on the
release page)
```bash
xkb-install -h
```

Meanwhile to quickly iterate and test it, here's an example
```bash
go run . -S fr -V optimot_ergo -d "French (Optimot, clavier Ergo)" --compose=./Optimot-Compose.txt ./Optimot-Ergo.xkb
# Alternatively
xkb-install -S fr -V optimot_ergo -d "French (Optimot, clavier Ergo)" --compose=./Optimot-Compose.txt ./Optimot-Ergo.xkb
```

The `Optimot-Ergo.xkb` file is expected to have symbol data about the variant like
```
xkb_symbols "optimot_ergo" {
    ...
};
```

## Features

- Idempotent
- Makes a layout available in system dropdowns

Extra details:
- [-] Install an extra `XkbSymbols` table in the `/usr/share/X11/xkb/symbols/XX` file
- [x] Add extra variants in the `evdev.lst` file under the `XX` symbol
- [x] Add extra variants in the `evdev.xml` file under the `XX` layout/symbol
  + [x] edit the XML file in tree
  + [x] backup the old file with a timestamp
  + [x] write the new file after getting sudo rights
- [x] Install an optional `.Xcompose` file in the home of `$USER`
