#!/usr/bin/env bash

startdir=$(pwd)
tmpdir=$(cat .tilt/tmp/.tmp_git_dir)

cd $tmpdir/dev
rsync -a --exclude .git $startdir/ $tmpdir/dev/
git add -A && git commit -m "update commit" && git push
cd $startdir
