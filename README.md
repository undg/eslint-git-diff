## Dependencies

Those apps are necessary to be presents in your system. Install them in whatever way you like.

* (eslint)[https://eslint.org/]
* (eslint_d)[https://www.npmjs.com/package/eslint_d]
* (git)[https://git-scm.com/]

## Installation

Download compressed archive from (latest release)[https://github.com/undg/eslint-git-diff/releases/latest] page. Unpack it to one of directories from your `$PATH`.

To list `$PATH` directories, you can run that command
```bash
echo $PATH | tr ':' '\n'
# or
sed 's/:/\n/g' <<< "$PATH"
```

## Compile from source

If your OS architecture is not present on (release page)[https://github.com/undg/eslint-git-diff/releases] here are steps to compile it from source.

```bash
git clone https://github.com/undg/eslint-git-diff
cd eslint-git-diff
go build -o build/eslint-git-diff
cd build
chmod +x eslint-git-diff
```

