# go-tic-tac

Multiplayer tic-tac-toe game written 100% in Golang, using Fyne GUI toolkit and Websockets. Works for Windows, macOS and
Linux.


- Learn more about Fyne at https://developer.fyne.io/started/
- Official GitHub for Fyne: https://github.com/fyne-io/fyne

## Game Server

You will also need a running _gameserver_, [available here](https://github.com/Longwater1234/server-tic-tac), also
written in Golang. The server is very tiny (~5MB), lightweight, and can handle lots of concurrent players without
sweating your RAM or CPU. It's purely a console based program.

### Screenshot

![screenshot.PNG](screenshot.PNG)

## Requirements

- Go 1.19 or higher
- C compiler (eg. gcc or Clang), and should be added to your PATH. Follow
  this quick [guide](https://developer.fyne.io/started/#prerequisites)
- For Windows users, the easiest & fastest way to get GCC, is to install it from [here](https://jmeubank.github.io/tdm-gcc/download/).
- Other OS, you can use your package manager.
- Any graphics driver installed.

## Building locally

1. Verify your machine has the requirements listed above
2. Install the Fyne v2 CLI: `go install fyne.io/fyne/v2/cmd/fyne@latest`
3. Go to this project root directory, open terminal and enter:

   ```bash
   # For more flags, use: fyne package --help
   fyne package
   ```

4. Be patient, first time build will take much longer to complete than subsequent ones.
5. For an optimized, smaller package (50% smaller), use command below. Icon will be automatically attached.

    ```bash
    # Flag --id (appID) is required. See docs https://developer.fyne.io/started/distribution
    fyne package --release --id com.yourdomain.appName
    ```

## Credits

Free Icon by [Vlad Marin, IconFinder](https://www.iconfinder.com/icons/190320/game_tac_tic_red_toe_icon).

## License

&copy; 2023, Davis Tibbz. This project is [MIT](LICENSE) licensed.

## Pull Requests & Contributions

Pull requests and issues are much welcome.

