# BT_PROFILE

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Donate](paypal.svg?raw=true)](https://www.paypal.me/ugjka)

Switch between the HSP/HFP and A2DP Bluetooth profiles from the system tray

## Usage

```text
$ bt_profile --help
Usage of bt_profile:
  -quit
        show the quit item
  -sink string
        headset's pulseaudio sink (default "1")
```

## Example

`$ bt_profile -sink "bluez_card.F4_0E_11_7E_93_EB"`

![screenshot](bt_profile.png?raw=true)

## Installation

Archlinux PKGBUILD in the archlinux folder

### Using make

You need make, go, go-tools

Build with `make`

Install with `make install`

Uninstall with `make uninstall`

Clean up with `make clean`
