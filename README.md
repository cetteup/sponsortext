# sponsortext

![](https://img.shields.io/badge/status-draft-white)
![](https://img.shields.io/badge/game-Battlefield_2-blue)
![](https://img.shields.io/badge/game-Battlefield_2142-blue)

An open standard for exposing Refractor 2 server variables via the otherwise unused "sponsor text" setting.

## Background

Neither Battlefield 2 not Battlefield 2142 offer a way of exposing custom variables for a server.
Thus, servers cannot (directly) provide information beyond what DICE deemed relevant some 20 years ago.
In contrast, servers for the equally old Call of Duty games do support adding custom variables.
The variables allow the games to nicely integrate with modern, web-based server browsers.
For example, [cod.pm](https://cod.pm/faq#6071) uses such custom variables to provide links to a server's Discord, website and/or TeamSpeak server.

## Sponsor text?

`sv.sponsorText` is a server setting available for both Battlefield 2 and Battlefield 2142.
A simple string ("text") variable with no hard character limit - that is _completely_ unused by the games themselves. 
In fact, the official server launchers don't even include `SponsorText` in the graphical server setting interface.

Overall, `sv.sponsorText` is perfect because:

* it is available on any Battlefield 2/Battlefield 2142 server
* it holds a string value, meaning it can hold any type of data
* it does not have a hard character limit
* it is not being used by the games
* it is included in the GameSpy-protocol query response, meaning it is easy to obtain externally
* it can be changed without restarting the server

## Syntax

```properties
sv.sponsorText "$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$"
```

### Keys and values

The proposed syntax provides a simple key-value structure.
Keys and values are considered to always be strings.
Key-value pairs are separated by semicolons.
Within a pair, key and value are separated by an equal sign.
Keys may _not_ contain whitespaces.
Values may contain any printable character, including whitespaces.

### Prefix and optional suffix

To allow adding other, non-variable values to `sponsorText`, the variables-section _must_ be prefixed with `$vars:`.
A `$`-suffix _can_ be used to end the variables-section (early).

### Escaping reserved characters

Any reserved character (`=`, `;`, `$`) can be escaped using a backslash, including backslashes themselves.
The escaped character will be considered part of the key/value, rather than part of the syntax.

### Parsing

Whitespaces before, in or after keys must be skipped.
When parsing a value, leading and trailing whitespaces must be omitted.
Consecutive whitespaces in values are to be parsed as a single whitespace unless the whitespaces are escaped.
Incomplete keys (keys not followed by an equal sign) must be skipped.
Keys without a value are considered present, but empty.

## Examples

The examples show the `sv.sponsorText` values without quotes.
They would be configured on the server as

```properties
sv.sponsorText "{example}"
```

### Variables only (with suffix)

```
$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$
```

**Equivalent JSON**

```json
{
  "discord": "https://discord.gg/vx4AKRfj",
  "provider": "bf2hub.com"
}
```

### Variables only (without suffix)

```
$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com
```

**Equivalent JSON**

```json
{
  "discord": "https://discord.gg/vx4AKRfj",
  "provider": "bf2hub.com"
}
```

### Non-variable prefix

```
Join our event this Sunday! $vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$
```

**Equivalent JSON**

```json
{
  "discord": "https://discord.gg/vx4AKRfj",
  "provider": "bf2hub.com"
}
```

### Non-variable suffix

```
$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$ Apply to become an admin today!
```

**Equivalent JSON**

```json
{
  "discord": "https://discord.gg/vx4AKRfj",
  "provider": "bf2hub.com"
}
```

### Escaped syntax characters

```
$vars:\$trange\=key=https://example.com?query\=\$start\;end$
```

**Equivalent JSON:**

```json
{
  "$trange=key": "https://example.com?query=$start;end"
}
```

### Escaped whitespaces

```
$vars:four-spaces=\ \ \ \ $
```

**Equivalent JSON:**

```json
{
  "four-spaces": "    "
}
```

### Omitted whitespaces

```
$vars: trimmed = one  two  three $
```

**Equivalent JSON:**

```json
{
  "trimmed": "one two three"
}
```

### Incomplete key

```
$vars:discord=https://discord.gg/vx4AKRfj;website$
```

**Equivalent JSON**

```json
{
  "discord": "https://discord.gg/vx4AKRfj"
}
```

### Key with no value

```
$vars:discord=https://discord.gg/vx4AKRfj;teamspeak=$
```

**Equivalent JSON**

```json
{
  "discord": "https://discord.gg/vx4AKRfj",
  "teamspeak": ""
}
```

## Got feedback?

Open an issue, leave a pull request or simply hop into [Discord](https://discord.gg/GsYyMMjEga).

