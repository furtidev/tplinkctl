<p align="center">
  <b><i>tplinkctl</i></b>
</p>
<p align="center">
  (unofficial) experimental cli for tp-link routers
</p>

---

I created `tplinkctl` because I couldn't find a CLI for my router. Even though you can do a lot more with [OpenWRT](https://openwrt.org/), I did not want to bother with all of that. I wanted something simple to use. Plus my specific router is not supported by OpenWRT anymore. And thus, `tplinkctl` was born. 

Currently only **TL-WR840N** is supported as it is the only router I have available. You can write a driver for your router if it's not already a part of this project or imporve existing ones.

⚠️ This project is still in pre-alpha stage, so expect a lot of things to be missing or just straight up not working. 

## Roadmap
- TL-WR840N
  - [x] Ability to check connected clients.
  - [ ] Ability to check current status. (with a bunch of other useful information suhc as uptime, DNS etc. etc.)
  - [ ] Ability to reboot.
  - More will be coming soon!
- Write a driver for your router!

## Project Structure
I'm using `urfave/cli/v2` for the command line interface. That's the only dependency for now. It has nice [documentation](https://cli.urfave.org/v2/getting-started/). Each driver is it's own separate package with functions to interact with the CLI. I recommend checking the `main.go` file and the `wr840n` driver package to see how they interact. The `utils` package stores common functions.

This project is really new. It'd take some time before it all comes together I suppose. I'm also a beginner, so feedback on code and structure is appreciated (Use issues to give feedback).

## License
This project is licensed under the [MIT license](./LICENSE).