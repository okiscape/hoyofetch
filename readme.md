# hoyofetch
> CLI tool for fetching Hoyoverse game account data.

## Supported Games
| Abbr | Game |
|------|------|
| `zzz` | Zenless Zone Zero |
| `hsr` | Honkai: Star Rail |
| `gi` | Genshin Impact |
| `tot` | Tears of Themis |
| `hi3rd` | Honkai Impact 3rd |
| `pp` | Petit Planet |
| `hna` | Honkai: New Anima |

## Build
```sh
make build
```
The binary will be placed at `./hoyofetch`.

## Usage
```sh
hoyofetch get <game abbr>
```

## Config
Config is stored at `~/.config/hoyofetch/config.json`.

### How to login?
To use this tool you have to get your hoyoverse account tokens.\
To get those, log in at https://www.hoyolab.com/ and read cookies from your browser.\
*For example: [cookie-editor](https://addons.mozilla.org/en-US/firefox/addon/cookie-editor/) for firefox based browsers or [MILK cookie manager](https://chromewebstore.google.com/detail/milk-—-cookie-manager/haipckejfdppjfblgondaakgckohcihp) for chromium-based.\*
In cookies after login you need to find `ltuid_v2`, `ltoken_v2`, and `ltmid_v2` and insert them into config.json.\
 
### config.json schema
```jsonc
{
	"accounts": [
		{
			"name": "",
			"auth": {
				"ltuid_v2": "",  // YOUR COOKIES
				"ltoken_v2": "", // YOUR COOKIES
				"ltmid_v2": ""   // YOUR COOKIES
			}
		}
	],
	"modules": [
		{
			"type": "game",
			"format": "hoyofetch - {}",
			"display": "lower"
		},
		{ 
			"type": "username",
			"format":   "  username   {}"
		},
		{
			"type": "userId",
			"format":   "  user id    {}"
		},
		{
			"format":   "  level      {}",
			"type": "level"
		},
		{
			"type": "server",
			"format":   "  server     {}"
		},
		{
			"type": "text",
			"format":   "  achievements"
		},
		{ 
			"type": "achievements",
			"format":   "    {name}: {value}"
		}
	]
}
```
### Formatting and types
Here's all available modules and their options:
<table>
  <tr>
    <td><p align="center"><code>game</code></p></td>
    <td><p align="left"><code>{}</code> - game name</p></td>
  </tr>
  <tr>
    <td><p align="center"><code>username</code></p></td>
    <td><p align="left"><code>{}</code> - account username</p></td>
  </tr>
  <tr>
    <td><p align="center"><code>userId</code></p></td>
    <td><p align="left"><code>{}</code> - account UID</p></td>
  </tr>
  <tr>
    <td><p align="center"><code>level</code></p></td>
    <td><p align="left"><code>{}</code> - account level</p></td>
  </tr>
  <tr>
    <td><p align="center"><code>server</code></p></td>
    <td><p align="left"><code>{}</code> - account server name</p></td>
  </tr>
  <tr>
    <td><p align="center"><code>text</code></p></td>
    <td><p align="left">don't replaces anything</p></td>
  </tr>
  <tr>
    <td><p align="center"><code>achievements</code></p></td>
    <td><p align="left"><code>{name}</code> - achievement name</p>
    <p align="left"><code>{value}</code> - achievement value</p></td>
  </tr>
</table>

 - `achievements` displays multiple times for each achievement
 - Display options available for each module: `lower`, `upper`, `capitalize` and `trim`
