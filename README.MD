# BT_PROFILE

## Archiving, I no longer need this

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Donate](paypal.svg?raw=true)](https://www.paypal.me/ugjka)

Switch between the HSP/HFP and A2DP Bluetooth profiles from the system tray

## Usage

```text
$ bt_profile --help
Usage of bt_profile:
  -codec string
        custom codec e.g. aac, ldac
  -quit
        show the quit item
  -sink string
        headset's pulseaudio sink (default "1")
```

* custom codecs are supported by https://github.com/EHfive/pulseaudio-modules-bt

## Example

Find the sink first

```text
$ pactl list cards | grep Name
        Name: alsa_card.pci-0000_00_1b.0
        Name: bluez_card.F4_0E_11_7E_93_EB
```

Then use it

`$ bt_profile -sink "bluez_card.F4_0E_11_7E_93_EB"`

![screenshot](bt_profile.png?raw=true)

Now you can put this in your DE start up script

## Installation

Archlinux PKGBUILD in the archlinux folder

### Using make

You need make, go, go-tools, libappindicator-gtk3

Build with `make`

Install with `make install`

Uninstall with `make uninstall`

Clean up with `make clean`
