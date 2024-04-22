#!/usr/bin/env bash

startdir=$(pwd)
tmpdir=$(mktemp -d)
mkdir -p .tilt/tmp
echo "$tmpdir" > .tilt/tmp/.tmp_git_dir
go run .tilt/init_gitea.go

cd $tmpdir
git clone http://test-user:password@localhost:3000/test-user/dev
rsync -a --exclude .git $startdir/ $tmpdir/dev/
cd $tmpdir/dev
git add -A && git commit -m "init commit" && git push
cd $startdir
