# minty

Experiment for minting OIDC tokens from GitHub Actions for use with OpenFaaS

Why would you want this?

* Enable third-parties to deploy to your openfaas organisation without sharing keys or passwords.
* Deploy to your own openfaas installation without saving or needing to rotate a password or key via a GitHub Actions Secret.

How does it work?

1) Install actions-federation plugin from OpenFaaS Pro to your local Kubernetes cluster or faasd installation
2) Configure the issuer for the plugin to: `https://vstoken.actions.githubusercontent.com`
3) The public key will be downloaded using the JWKS URL to validate any JWTS from GitHub Actions
4) Set a list of owners who can access the installation
5) Obtain an OIDC token from GitHub
6) Pick out the JWT from the response
7) Use the JWT via the `--token` flag with the `faas-cli`

Next, trigger a commit using this forked repository. If your name is in the owners list, this will work. If not, it will fail.

Test

Test 2