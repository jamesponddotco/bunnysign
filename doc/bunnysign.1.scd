bunnysign(1)

# NAME

bunnysign - command-line tool to generate signed URLs for Bunny.net

# SYNOPSIS

*bunnysign* [OPTIONS]

# DESCRIPTION

*bunnysign* is a command-line tool designed to generate signed URLs for
Bunny.net's Token Authentication feature. It allows users to create secure,
expiring URLs to control access to content hosted on Bunny.net.

# OPTIONS

*-u*, *--url*
	The URL to be signed.

*-k*, *--key*
	The key used to sign the URL.

*-e*, *--expiration*
	The expiration time of the signed URL in hours, minutes, or seconds.

*-h*, *--help*
	Show help message and quit.

*-v*, *--version*
	Show version number and quit.

## EXAMPLES

*Example 1. Sign a URL with a 1 hour expiration.*
	$ bunnysign -u 'https://test.b-cdn.net/my-partial/url/video.mp4' -k 'ed9b36bc-1901-43dd-b4d9-447e4451e90a' -e '1h'

*Example 2. Sign a URL with the default expiration of 1 year*
	$ bunnysign -u 'https://test.b-cdn.net/my-partial/url/video.mp4' -k 'ed9b36bc-1901-43dd-b4d9-447e4451e90a'

# AUTHORS

Maintained by James Pond <james@cipher.host>.
