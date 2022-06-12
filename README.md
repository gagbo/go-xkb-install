# Helper binary to install xkb keyboard layouts

This script notably helps to install a new ruleset within
the only file that actually gets parsed by the system to
list the available layouts.

## Usage

Once the binanry gets properly installed
```
xkb-install -h
```

Meanwhile to quickly iterate and test it, here's an example
```
go run . -S fr -V optimot_ergo -d "French (Optimot, clavier Ergo)" ./Optimot-Ergo.xkb
```

The `Optimot-Ergo.xkb` file is expected to have symbol data about the variant like
```
partial alphanumeric_keys
xkb_symbols "optimot_ergo" {
    ...
};
```

## Features

- Idempotent
- Makes a layout available in system dropdowns

Extra details:
- [ ] Install an extra `XkbSymbols` table in the `/usr/share/X11/xkb/symbols/XX` file
- [ ] Add extra variants in the `evdev.lst` file under the `XX` symbol
- [-] Add extra variants in the `evdev.xml` file under the `XX` layout/symbol
  + [x] edit the XML file in tree
  + [ ] backup the old file with a timestamp
  + [ ] write the new file after getting sudo rights
- [ ] Install an optional `.Xcompose` file in the home of `$USER`
