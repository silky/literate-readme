#!/usr/bin/env ./go

```haskell
{-# LANGUAGE OverloadedStrings #-}

import Turtle
```

# Literate README

[![Build status](https://github.com/silky/literate-readme/workflows/CI/badge.svg)](https://github.com/silky/literate-readme/actions?query=workflow%3ACI) <a href="https://high5.cool/high5/754d8358-fe0b-5e5c-8407-beec85b1a603"><img src="https://high5.cool/static/img/high5-me-green.png" />
</a>

The readme that builds itself!

You can setup and build this project by running this very readme. This is what
the CI job does!

Example: `./README.lhs --setup --test`


```haskell
parser :: Parser (Bool, Bool, Bool)
parser = (,,) <$> switch "setup" 's' "Set up the stack environment."
              <*> switch "test"  't' "Build the project and run the tests."
              <*> switch "build" 'b' "Just build, don't run tests."
```

```haskell
main = void $ do
    (setup, test, build) <- options "Literate README" parser
    let ops = doSetup setup .&&. doBuild build .&&. doTest test
    ops .||. die "Step failed."

nop = shell "true" empty

stackOrNop op True = shell ("stack " <> op) empty
stackOrNop _  _    = nop
```

## Setup

```haskell
-- | Call this with: ./README.lhs --setup
doSetup = stackOrNop "setup"
```


## Build

```haskell
-- | Call this with: ./README.lhs --build
doBuild = stackOrNop "build"
```


## Test

```haskell
-- | Call this with: ./README.lhs --test
doTest = stackOrNop "test"
``` 
