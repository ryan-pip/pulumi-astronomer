---
description: Don't claim a root cause from circumstantial evidence — run an experiment that isolates the suspected variable before asserting "this is why X fails". The user has pushed back twice when I shipped "this is the bug" without verifying.
paths:
  - "**/*"
---

# Verify Before Claiming

When diagnosing a CI failure, behavior delta, or "X is broken because Y", verify the causal link before stating it. Inferring from circumstantial evidence ("there's a version skew → that must be the cause") is not the same as proving it ("ran with version A, ran with version B, observed the delta").

## Tells you're inferring, not verifying

- Hedged language: "this is likely…", "it must be that…", "consistent with the theory that…".
- Citing a fact that *could* be related, with no test that connects it to the failure.
- "Local matches what's checked in" — proves consistency between dev and repo; says nothing about CI's runtime.

## How to verify (concrete examples)

- **Suspected version drift** (e.g. `pulumi 3.232.0` vs `3.233.0` codegen): install both, run codegen with each, diff the output. Confirm the delta matches what CI sees.
- **Suspected platform drift** (Mac vs Linux): run the same binary in a Linux container (`podman run --platform linux/amd64 ...`) and compare against Mac output.
- **Suspected env / permissions**: replicate the env (token type, CI runner image, locale) and reproduce.

For any of these, the test is: "if I'm right, *this specific output* should appear / not appear."

## When you can't verify locally

Say so explicitly: "I haven't verified this is the cause; my best guess is X based on Y, but the test would be Z." Don't propose fixes predicated on unverified diagnosis.

## Don't

- Don't call a CI failure "transient" without evidence — a single retry passing can be unrelated to the bug.
- Don't conflate "matches checked-in state" with "matches CI's build" — different consistency claims.
- Don't reach for an admin-merge / bypass before you've actually proved the failing check is unrelated.
