# Demonstrating templ CSP issue

When I use a script in templ then, if I want to use a content security policy, I need to include a `script-src 'unsafe-hashes'` content security policy because any javascript included in an html attribute (e.g. onload) needs the unsafe-hashes CSP.

The suggestion from [the unsafe-hashes CSP page](https://content-security-policy.com/unsafe-hashes/) is:

> Whenever you see the prefix unsafe in a CSP keyword, that means that using this is not the most secure way to go. It is better to refactor your code to avoid using HTML event handler attributes (such as onload, onclick, onmouseover etc.)

...of course you might accept the risk of having an unsafe CSP depending on what your app does, but many organisations will enforce a "safe" CSP.

The example in this repo has a [page with a script](components/page.templ). When served, this generates a page that looks like:

```html
<!doctype html>
<html lang="en">
    <head></head>
    <script type="text/javascript">
        function __templ_App_1e4f() {
            console.log("Loaded!")
        }
    </script>
    <body onload="__templ_App_1e4f()">
        <h1>Check the console for a log!</h1>
    </body>
</html>
```

Since the `body` tag includes an `onload` attribute with javascript in it, we need to include `unsafe-hashes` in the CSP.

This repo includes some tests to validate a fix.

## Testing

In two separate processes run:

1. `make run` to run the example app.
2. `make test` to run tests against the example app. 

The tests make sure:

1. The website loads without console/inspector errors in the browser (i.e. CSP doesn't block required scripts)
2. A CSP header is included and that CSP header does not contain the word `unsafe`

## Notes

### Is this really an issue?

Interestingly neither of the following open source tools regard `unsafe-hashes` as an issue when checking CSP:

1. <https://github.com/mozilla/http-observatory>
2. <https://github.com/google/csp-evaluator>

This may be because the first hasn't been updated in a long time, and the second is more around spotting common issues with CSPs based on a large data set of sites.

### Running on linux

I use `host.docker.internal` [in the lighthouse (website loads without error) tests](test/lighthouse/cmd.sh) - this may not work on linux. Let me know if you need a docker-compose to fix this!