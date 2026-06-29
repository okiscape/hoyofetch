# hoyofetch
> CLI tool for fetching Hoyoverse game account data.

> [!CAUTION]
> STILL IN DEVELOPMENT, DON'T DOWNLOAD IF YOU WON'T CONTRIBUTE!

**Supported Games**
| Abbr | Game |
|------|------|
| `zzz` | Zenless Zone Zero |
| `hsr` | Honkai: Star Rail |
| `gi` | Genshin Impact |
| `tot` | Tears of Themis |
| `hi3rd` | Honkai Impact 3rd |
| `pp` | Petit Planet |
| `hna` | Honkai: New Anima |

**Build**
```sh
make build
```
The binary will be placed at `./hoyofetch`.

**Usage**
```sh
hoyofetch get <game abbr>
```

**Config**
Config is stores at `~/.config/konawalls/config.json`.
**config.json schema**
```json
{
	"accounts": [
		{
			"name": "",
			"auth": {
				"ltuid_v2": "",
				"ltoken_v2": "",
				"ltmid_v2": ""
			},
			"games": ["zzz", "gi"]
		}
	]
}
```
