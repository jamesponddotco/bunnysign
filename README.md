# `bunnysign`

`bunnysign` is a command-line tool designed to generate signed URLs for
[Bunny.net's Token Authentication
feature](https://docs.bunny.net/docs/cdn-token-authentication). It
allows users to create secure, expiring URLs to control access to
content hosted on [Bunny.net](https://bunny.net/).

## Installation

### From source

First install the dependencies:

- Go 1.21 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```bash
make
sudo make install
```

## Usage

```bash
$ bunnysign --help
NAME:
   bunnysign - generate a signed URL for Bunny.net with token authentication

USAGE:
   bunnysign [global options] [arguments...]

VERSION:
   0.1.0

GLOBAL OPTIONS:
   --url value, -u value         the URL to sign
   --key value, -k value         the key to sign with
   --expiration value, -e value  the expiration time of the signature (default: 8760h0m0s)
   --help, -h                    show help
   --version, -v                 print the version
```

See _bunnysign(1)_ after installing for more information.

## Contributing

Anyone can help make `bunnysign` better. Send patches on the [mailing
list](https://lists.sr.ht/~jamesponddotco/bunnysign-devel) and report bugs
on the [issue tracker](https://todo.sr.ht/~jamesponddotco/bunnysign).

You must sign-off your work using `git commit --signoff`. Follow the
[Linux kernel developer's certificate of
origin](https://www.kernel.org/doc/html/latest/process/submitting-patches.html#sign-your-work-the-developer-s-certificate-of-origin)
for more details.

All contributions are made under [the GPL-2.0 license](LICENSE.md).

## Resources

The following resources are available:

- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/bunnysign-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/bunnysign-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/bunnysign).

---

Released under the [GPL-2.0 license](LICENSE.md).
