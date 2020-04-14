# go-teamhealthcheck

[![Build Status](https://travis-ci.org/FelixGeelhaar/go-teamhealthcheck.svg?branch=master)](https://travis-ci.org/github/FelixGeelhaar/go-teamhealthcheck)

## Content

1. Concept of the Team/Squad Health Check
2. Goal of this Module
3. Usage
4. Contribution

### Concept of the Team Health Check

The Team/Squad Health Check was formerly mentioned from Spotify as the Squad
Health Check. It is used to evaluate and identify on how is the team doing. The
Stakeholders of this Health Check are the teams themselves as also shared
Leaders, so that those get a sense on where to focus their improvement efforts,
spot systemic problems, and help teams become more self-aware so they can focus
their improvement efforts too.
[further reading](https://labs.spotify.com/2014/09/16/squad-health-check-model/ 'Spotifys Blog')

### Goal of this Module

The module should help you to provide a simple model of the Health Check to just
be able to implement a Service or API to use the Health Check. Mostly because
the Teams I'm working with requested to gather the data not manually anymore,
but through an App or website, I decided to start simple with a model.

### Next Steps

This model just reached version `v0.1.0`. So it's the very beginning. Once I'm
starting to use it more often, this model will get extentend when it makes
sense. There will be example implementations following.

Otherwise just check the Changelog.

### Usage

As this is a simple model, the usage is concidered the following path:

```
  // Beginning of your go file / function / ...
  metric1 := teamhealthcheck.New("Metric 1")
  metric1.SetTendency(2)
  metric1.SetRed(2)
  ...
```

### Contribution

Feel free to contribute, but make sure to follow along with the following
guideline for Commits:

_Structural elements_:

- _fix_: a commit of the type fix patches a bug in your codebase (this
  correlates with PATCH in semantic versioning)
- _feat_: a commit of the type feat introduces a new feature to the codebase
  (this correlates with MINOR in semantic versioning).
- _BREAKING CHANGE_: a commit that has the text BREAKING CHANGE: at the
  beginning of its optional body or footer section introduces a breaking API
  change (correlating with MAJOR in semantic versioning). A BREAKING CHANGE can
  be part of commits of any type.
- _chore_, _docs_, _style_, _refactor_, _perf_ or _test_ are fine too

[further reading](https://www.conventionalcommits.org/en/v1.0.0-beta.4/#summary)

_Commit messages_:

- Separate subject from body with a blank line
- Limit the subject line to 50 characters
- Capitalize the subject line
- Do not end the subject line with a period
- Use the imperative mood in the subject line
- Wrap the body at 72 characters
- Use the body to explain what and why vs. how

[further reading](https://chris.beams.io/posts/git-commit/#seven-rules)
