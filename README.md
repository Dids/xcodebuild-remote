[![Build Status](https://travis-ci.org/Dids/xcodebuild-remote.svg?branch=master)](https://travis-ci.org/Dids/xcodebuild-remote)

# xcodebuild-remote

`xcodebuild-remote` is a tool for extending `xcodebuild` for the purpose of building remote repositories.

**NOTICE:** _Work in progress._ 

### Requirements

- [macOS](https://www.apple.com/lae/macos/) (only tested on macOS High Sierra)
- [Xcode](https://developer.apple.com/xcode/) (available on the App Store)
- [Homebrew](https://brew.sh/)

### Installation

> brew tap Dids/brewery  
> brew install xcodebuild-remote  

### Usage

**NOTICE:** _SVN support is planned but not there yet._ 

> xcodebuild-remote --url <repository url> <optional xcodebuild arguments>  

Various formats are supported for the repository url:  
- HTTP(S) (`https://repo.url`)  
- Git (`git@url:user/repo.git` and `git+ssh@url:user/repo.git`)  
- GitHub (`user/repo`)  

### License

See [LICENSE](LICENSE).
