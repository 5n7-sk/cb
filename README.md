<p align="center">
  <a href="https://github.com/skmatz/cb">
    <img src="./assets/images/banner.png" width="1000" alt="banner" />
  </a>
</p>

<p align="center">
  <a href="https://github.com/skmatz/cb/actions?query=workflow%3Abuild">
    <img
      src="https://github.com/skmatz/cb/workflows/build/badge.svg"
      alt="build"
    />
  </a>
  <a href="https://github.com/skmatz/cb/actions?query=workflow%3Arelease">
    <img
      src="https://github.com/skmatz/cb/workflows/release/badge.svg"
      alt="release"
    />
  </a>
  <a href="./LICENSE">
    <img
      src="https://img.shields.io/github/license/skmatz/cb"
      alt="license"
    />
  </a>
  <a href="./go.mod">
    <img
      src="https://img.shields.io/github/go-mod/go-version/skmatz/cb"
      alt="go version"
    />
  </a>
  <a href="https://github.com/skmatz/cb/releases/latest">
    <img
      src="https://img.shields.io/github/v/release/skmatz/cb"
      alt="release"
    />
  </a>
</p>

<p align="center">
  <img src="./assets/images/demo.png" width="640" alt="demo" />
</p>

# cb

:bookmark: Chrome bookmarks from terminal

**cb** is an interactive CLI to open Chrome bookmarks from the terminal.  
Just run `cb` and you can search for Chrome bookmarks and open in your browser.

## Usage

Basically you just run the command.  
Note that Windows users cannot use arrow keys to move due to [this issue](https://github.com/manifoldco/promptui/issues/49), so type "?" to switch the search mode and use "j" or "k" keys like Vim to move up or down.

## Install

### Binary

Get binary from [releases](https://github.com/skmatz/cb/releases).

### Source

```sh
go get github.com/skmatz/cb
```
