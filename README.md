# go-hcaptcha

https://godoc.org/github.com/kirari04/go-hcaptcha

## About

This package handles [hCaptcha](https://www.hcaptcha.com).

## Usage

Install the package in your environment:

```
go get github.com/kirari04/go-hcaptcha
```

To use it within your own code, import <tt>github.com/kirari04/go-hcaptcha</tt> and call:

```
recaptcha.Init (recaptchaPrivateKey)
```

once, to set the hCaptcha private key for your domain, then:

```
recaptcha.Confirm (clientIpAddress, recaptchaResponse)
```
