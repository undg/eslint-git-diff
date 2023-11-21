## Motivation

Running eslint in large projects can be very time consuming. This solution can reduce time of linting from dozens of seconds to milliseconds. It will lint only files that are changed relatively to given branch (`dev` by default). It can also run in watch mode.

## Dependencies

Those apps are necessary to be presents in your operating system. Install them in whatever way you like.

* [eslint](https://eslint.org/)
* [eslint_d](https://www.npmjs.com/package/eslint_d)
* [git](https://git-scm.com/)

## Installation

Download compressed archive from [latest release](https://github.com/undg/eslint-git-diff/releases/latest) page. Unpack it to one of directories from your `$PATH`. Make it executable with command `chmod +x ...`

To list `$PATH` directories, you can run that command
```bash
echo $PATH | tr ':' '\n'
# or
sed 's/:/\n/g' <<< "$PATH"
```

## Compile from source

If your OS architecture is not present on [release page](https://github.com/undg/eslint-git-diff/releases) here are steps to compile it from source.

```bash
git clone https://github.com/undg/eslint-git-diff
cd eslint-git-diff
go build -o build/eslint-git-diff
cd build
chmod +x eslint-git-diff
```

## Known problems

* eslint_d heavily cashing: After hooping between branches it's cache can be out of sync and some nonsense can be produced in error log. In that case you can restart it witch command `esling_d --reset`. It's know limitation of any caching system.
* I need to pass flag XX or YY to underlying eslint: WIP, proxy for all eslint flags is not implemented yet, but will be in future.
