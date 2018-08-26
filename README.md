# pkgr

[![Build Status](https://travis-ci.org/mrtazz/pkgr.svg?branch=master)](https://travis-ci.org/mrtazz/pkgr)
[![Coverage Status](https://coveralls.io/repos/mrtazz/pkgr/badge.svg?branch=master&service=github)](https://coveralls.io/github/mrtazz/pkgr?branch=master)
[![Packagecloud](https://img.shields.io/badge/packagecloud-available-brightgreen.svg)](https://packagecloud.io/mrtazz/pkgr)
[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT)

## Overview
Simple tool to create FreeBSD packages from a directory. This basically
implements that one single use case from [fpm][]. If you need anything more
than this you should probably check out fpm.

## Usage

```
portal:pkgr[master]% cat MANIFEST
{
  "name": "pkgr",
  "version": "0.1.0-1",
  "comment": "create pkgng packages from directory",
  "desc": "create pkgng packages from directory",
  "maintainer": "Daniel Schauenberg <d@unwiredcouch.com>",
  "www": "https://github.com/mrtazz/pkgr"
}
portal:pkgr[master]% ./pkgr --manifest MANIFEST --path usr
```

## Why?
I dealt with a bunch of gem dependency issues on my FreeBSD build box with
fpm. And after tracking the issues down and opening issues on fpm I got
curious how it would look like to implement the one use case I needed in Go.

## Inspiration
This basically implements one simple use case of [fpm][].


[fpm]: https://github.com/jordansissel/fpm
