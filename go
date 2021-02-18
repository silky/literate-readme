#!/bin/sh

# Drop the 1st arg which is just `./README.lhs`
# so this script makes sense.
shift

# Pass al the other args on.
stack runghc --resolver lts-8.24 --package turtle --package markdown-unlit -- "-pgmL markdown-unlit" README.lhs "$@"
