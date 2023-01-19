# go-tic-tac

Simple GUI game of Tic-Tac-Toe written in Golang and Fyne framework v2.3.0. For 2 local (human) players. Works for
Windows, MacOS and Linux. Can even be compiled to Android and iOS (macOS required) apps. 
- Learn more about fyne at https://developer.fyne.io/

## Quick Start

Jump to the releases page and download pre-built executable. No installation required. Just double click to start.

## Building locally

1. Make sure you have >=Go 1.18 and GCC installed on your machine, then add both to your PATH. Follow
   this [guide](https://developer.fyne.io/started/#prerequisites) for your specific OS.
2. To verify *gcc* is correctly installed, open terminal/CMD, enter `gcc --version`
3. Install the Fyne v2 CLI. `go install fyne.io/fyne/v2@latest`
4. Go to this project root directory, and open terminal and enter

```bash
# For more flags, use: fyne build --help
fyne build
```
5. First time build will take quite some time to complete. But subsequent builds will be much faster.
6. To build a smaller executable (50% smaller than one above), use:

```bash
# For more flags, use: fyne release --help
# --icon is optional. --id is required. See docs. 
fyne release --icon game_icon.png --id com.yourdomain.appName
```

## Credits
Icon by  [Vlad Marin, IconFinder](https://www.iconfinder.com/icons/190320/game_tac_tic_red_toe_icon)

## License

This project is [MIT](LICENSE) licensed.

## Pull Requests & Contributions
So much improvement can be done on the game. This is as simple as it can be. Pull requests and issues are much welcome.

