# go-tic-tac

Simple GUI game of Tic-Tac-Toe written in Golang and Fyne v2.3.0. For 2 local (human) players. Works for Windows, macOS
and Linux. Can even be compiled to Android and iOS (macOS required) apps.

- Learn more about fyne at https://developer.fyne.io/

## Building locally

1. Make sure you have Go>=1.18 and GCC installed on your machine, then add both to your PATH. Follow
   this [guide](https://developer.fyne.io/started/#prerequisites) for your specific OS.
2. To verify *gcc* is correctly installed, open terminal/CMD, enter `gcc --version`
3. Install the Fyne v2 CLI. `go install fyne.io/fyne/v2/cmd/fyne@latest`
4. Go to this project root directory, open terminal and enter:

```bash
# For more flags, use: fyne package --help
fyne package
```

5. First time build will take quite some time to complete. But subsequent builds will be much faster.
6. For an optimized, smaller executable (50% smaller than one above), use command below. Another option for compression
   is to use `upx` tool.

```bash
# Flag --id (appID) is required. See docs https://developer.fyne.io/started/distribution
fyne package --release --id com.yourdomain.appName
```

## Credits

Free Icon by [Vlad Marin, IconFinder](https://www.iconfinder.com/icons/190320/game_tac_tic_red_toe_icon).

## License

&copy; 2023, Davis Tibbz. This project is [MIT](LICENSE) licensed.

## Pull Requests & Contributions

So much improvement can be done on the game. This is as simple as it can be. Pull requests and issues are much welcome.

