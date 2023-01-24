package main

import (
	"github.com/meddler-io/watchdog/bootstrap"
)

func main() {

	private_key := `-----BEGIN RSA PRIVATE KEY-----
MIIG4wIBAAKCAYEAzdCLcwQgyVTr24vBU6ZFk5Ye1i9hZPxzKp6tlfTBFEQmufS+
d0lgnMrLL0Svod/Ankm/yuiy1v+rHW/v/krsPBPIdL4ItnrAPJrScLXXKy2wbQhd
8quewRMP/e81C1P0hIDhULgTz6oKojV7nOb1s70E3euKJ2yNOddne9Lsg6VDGdfm
1XVXXXK1QOdXk62nZm/5CQ/q+XGORhOF2iE90+nfLkP4QrdbpvWgyMCLVWI6m4QD
CjyhtcLJzlJ2Vrz0vm3Nfb5nM11Ug9FDVHG3X4bk9U1HncoR4axzgCFWBw5jIiux
SYtO33UWEqtpq3VGUzeP7brVyMYdt5FJFznqmwlsELpLT4z772XaJarDxfM1Y6GF
UTSMxzZEszOztd5IyaHgGnh+DsJbv8pkFO7R27gryzloPi2sMyY2Ox5K1/Ksn74S
DLbWexohWglsCvmI+kixXaZCMoqwqQSlq/B8Am6CR920nhITadhSxRHvl30OSX7W
U7x/GaFlsk3nDYxxAgMBAAECggGAQZDkO9ne0AS7to9JIWrg02mJnVYuwsshrTwr
tI/qSpiDdc4gN0QrAQghvoc1JG5xoHuCHBQRlSiRzKjyXaJkq6bpvzci828berMf
w0RKrbi6caGuvDYsNvnpqE3iKh+uz9926A4cUZWsGYzMOmnhcEcRG9mDg5ezTe/v
SaUiL7YgMEV1hh8BnAWE9I4+eg9iOuIFBfwKMTkZx/1EMP9SAQaKUOTaW/apvgCW
Hh4oxSSr4wslJMe4BH1vaMkgZfV/uA6M+guQhKKqyjvBV9ynj1j+cl3YBEhELHqi
G5CwWIizNQc7k2OoGQnFsLHlcE9jdsdAiTlKs7XQAVOA26bvzBpL7EN0Pn5K4nrE
InY5sBr5nsAMljn+UhyWzBS7UiPuC6L3fA1PmOWj43ZQvw+mQLsOAFq40a3qmDCQ
lPvgMN7kVnIR3v6HZ36tWamMue65vv5eUL8fRlP5VYb7iIaAo15ZRAuzoZxIO4ah
zxCT6ZKzb5hkkjblWDna/oJK7cIhAoHBAPH15R/4ylYig8Tl0aq+vgm7OjiefVFa
5tqbt8OW6cMZT0kJzQL+sIsQ0wVrK9AOQJF6pkpcY8oWECF3TRzHcG0pHbyyo8fj
AdSBcNsAcB7dcGIsLPmSWfqP5oycKQ1NizfWCaqonS7VNXcFFBMqW5finBE45j/R
nS5DoqY79nKI2PYb/Iil/LqrNpY52Zui/SpxIAhKcfD3qI87jxL1MO03nfhKjp8r
J/RHt4dEQirw0+2TXu2TKIgImtNo7SyyFQKBwQDZwbwq8+fz02g9jU3D6iikCwNy
M5q4uejths5yB17vKsMxqBOS2d+gg5HEJtVvDD1TJy5Y0eDBQA7fQsn8Q7F9xbjm
cT78u3H9l01YLF9lNP3BmntPML8SdPdp9qzh6QwMvog+XPECMedRsdzmz7O9k7UC
aSjMB01nxhhn7VOHVyyXud4tkRprbyCdOXWaCkH5rXgu6a5lY0OgDB9PNvgSmlOX
4evuxEUkzHOjyqYpyZkEDfG9LVdEIx03QtWIs+0CgcEArXiH3tY16m8wXZgPLAU4
pi1vXmxymEM9w3lk1Ht7+P5KU6kNLJqA5BMZ14awkKfHLwXp9uIqQRZ0Vin8RSfP
uNODmGfS+FoOU9bBnKHBwL8t9ZSKYhvFGtjUh/blj74Kk+uGXq4okw0VGdGwRNVX
eZtkHQzYNuUHdcXT9K6E0DTAqCIEecvXV9WseyPambTPIQCd4JPPAN3uIVKoDwVv
IIOBZCPAAOml6YMJV5degP7MUkYYBug8ZNsWdX+Cm0rJAoHAbi3gj7spue1V6Q7+
Mgrg1V/dNoN13dHzUXvMKVYhatIwQYfRn5Qt2oG314/unmmK/e/tO3O2TRGkdMLO
gJ2fyQuqKq9SN36AhtdvgxWwjvzpCHSpv7/ibexQ44S0pPnN5wTg2u0b5QhrCqL8
Kq0dPObnx3h98/4d5EwPPUaP5QLuxTZ+fH/W5lUsI2FXZ8GY8EQdc185poM7TFfV
37bofkEn2smtHvqgIaZxPIbk3G/sxGhp0FaIuSkrNzLR6IFhAoHAQVgTIEyIXBFP
+njvAHAWyZOOVazx1HSa/4jIXx48xf57L57x6HQLwQ3HrUEevq0IM0BhWI1s2ezp
yPwxwWVv4mL0gtYdDZ8yZPF7rCURuS2KDE63xM+89amQWaRwpR81u9SiPw5iebeE
zD0QpcaEt398iX2aj3FbieB53NB2pzwnC3SQUpsbKiDe/jJqZvvhzduzDflrCL1B
RTDunsCza2FxlBgEk6kDbT7opKUupp0aA51BQyIQ6sv2S5Q0Khhl
-----END RSA PRIVATE KEY-----
`

	// studiogangster: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
	print("Auth Token")

	bootstrap.Clone("https://github.com/meddler-io/secoflex.git", "/Users/apple/workspaces/watchdog/tmp/", "token", "", "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")

	print("Auth Password")

	bootstrap.Clone("https://github.com/meddler-io/secoflex.git", "/Users/apple/workspaces/watchdog/tmp/", "password", "studiogangster", "Prakhar@1993")

	print("SSH")

	bootstrap.Clone("git@github.com:meddler-io/secoflex.git", "/Users/apple/workspaces/watchdog/tmp/", "ssh", private_key, "")

	print("No Auth")

	bootstrap.Clone("https://github.com/meddler-io/watchdog.git", "/Users/apple/workspaces/watchdog/tmp/", "no_auth", "", "")

	// if got != "PopulateStr" {
	// t.Errorf("Abs(-1) = %d; want 1", got)
	// }
}
