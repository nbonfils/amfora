package config

//go:generate ./default.sh
var defaultConf = []byte(`# This is the default config file.
# It also shows all the default values, if you don't create the file.

# All URL values may omit the scheme and/or port, as well as the beginning double slash
# Valid URL examples:
# gemini://example.com
# //example.com
# example.com
# example.com:123


[a-general]
# Press Ctrl-H to access it
home = "gemini://gemini.circumlunar.space"

# Follow up to 5 Gemini redirects without prompting.
# A prompt is always shown after the 5th redirect and for redirects to protocols other than Gemini.
# If set to false, a prompt will be shown before following redirects.
auto_redirect = false

# What command to run to open a HTTP(S) URL.
# Set to "default" to try to guess the browser, or set to "off" to not open HTTP(S) URLs.
# If a command is set, than the URL will be added (in quotes) to the end of the command.
# A space will be prepended to the URL.
#
# The best to define a command is using a string array.
# Examples:
# http = ['firefox']
# http = ['custom-browser', '--flag', '--option=2']
# http = ['/path/with spaces/in it/firefox']
#
# Note the use of single quotes, so that backslashes will not be escaped.
# Using just a string will also work, but it is deprecated, and will degrade if
# you use paths with spaces.

http = 'default'

# Any URL that will accept a query string can be put here
search = "gemini://gus.guru/search"

# Whether colors will be used in the terminal
color = true

# Whether ANSI color codes from the page content should be rendered
ansi = true

# Whether to replace list asterisks with unicode bullets
bullets = true

# A number from 0 to 1, indicating what percentage of the terminal width the left margin should take up.
left_margin = 0.15

# The max number of columns to wrap a page's text to. Preformatted blocks are not wrapped.
max_width = 100

# 'downloads' is the path to a downloads folder.
# An empty value means the code will find the default downloads folder for your system.
# If the path does not exist it will be created.
# Note the use of single quotes, so that backslashes will not be escaped.
downloads = ''

# Max size for displayable content in bytes - after that size a download window pops up
page_max_size = 2097152  # 2 MiB
# Max time it takes to load a page in seconds - after that a download window pops up
page_max_time = 10

# Whether to replace tab numbers with emoji favicons, which are cached.
emoji_favicons = false


[auth]
# Authentication settings
# Note the use of single quotes for values, so that backslashes will not be escaped.

[auth.certs]
# Client certificates
# Set domain name equal to path to client cert
# "example.com" = 'mycert.crt'

[auth.keys]
# Client certificate keys
# Set domain name equal to path to key for the client cert above
# "example.com" = 'mycert.key'


[keybindings]
# In the future there will be more settings here.

# Hold down shift and press the numbers on your keyboard (1,2,3,4,5,6,7,8,9,0) to set this up.
# It is default set to be accurate for US keyboards.
shift_numbers = "!@#$%^&*()"


[url-handlers]
# Allows setting the commands to run for various URL schemes.
# E.g. to open FTP URLs with FileZilla set the following key:
#   ftp = 'filezilla'
# You can set any scheme to "off" or "" to disable handling it, or
# just leave the key unset.
#
# DO NOT use this for setting the HTTP command.
# Use the http setting in the "a-general" section above.
#
# NOTE: These settings are overrided by the ones in the proxies section.
# Note the use of single quotes, so that backslashes will not be escaped.

# This is a special key that defines the handler for all URL schemes for which
# no handler is defined.
other = 'off'


[cache]
# Options for page cache - which is only for text/gemini pages
# Increase the cache size to speed up browsing at the expense of memory

# Zero values mean there is no limit
max_size = 0  # Size in bytes
max_pages = 30 # The maximum number of pages the cache will store


[proxies]
# Allows setting a Gemini proxy for different schemes.
# The settings are similar to the url-handlers section above.
# E.g. to open a gopher page by connecting to a Gemini proxy server:
#   gopher = "example.com:123"
#
# Port 1965 is assumed if no port is specified.
#
# NOTE: These settings override any external handlers specified in
# the url-handlers section.
#
# Note that HTTP and HTTPS are treated as separate protocols here.


[theme]
# This section is for changing the COLORS used in Amfora.
# These colors only apply if 'color' is enabled above.
# Colors can be set using a W3C color name, or a hex value such as "#ffffff".

# Note that not all colors will work on terminals that do not have truecolor support.
# If you want to stick to the standard 16 or 256 colors, you can get
# a list of those here: https://jonasjacek.github.io/colors/
# DO NOT use the names from that site, just the hex codes.

# Definitions:
#   bg = background
#   fg = foreground
#   dl = download
#   btn = button
#   hdg = heading
#   bkmk = bookmark
#   modal = a popup window/box in the middle of the screen

# EXAMPLES:
# hdg_1 = "green"
# hdg_2 = "#5f0000"

# Available keys to set:

# bg: background for pages, tab row, app in general
# tab_num: The number/highlight of the tabs at the top
# tab_divider: The color of the divider character between tab numbers: |
# bottombar_label: The color of the prompt that appears when you press space
# bottombar_text: The color of the text you type
# bottombar_bg

# hdg_1
# hdg_2
# hdg_3
# amfora_link: A link that Amfora supports viewing. For now this is only gemini://
# foreign_link: HTTP(S), Gopher, etc
# link_number: The silver number that appears to the left of a link
# regular_text: Normal gemini text, and plaintext documents
# quote_text
# preformatted_text
# list_text

# btn_bg: The bg color for all modal buttons
# btn_text: The text color for all modal buttons

# dl_choice_modal_bg
# dl_choice_modal_text
# dl_modal_bg
# dl_modal_text
# info_modal_bg
# info_modal_text
# error_modal_bg
# error_modal_text
# yesno_modal_bg
# yesno_modal_text
# tofu_modal_bg
# tofu_modal_text

# input_modal_bg
# input_modal_text
# input_modal_field_bg: The bg of the input field, where you type the text
# input_modal_field_text: The color of the text you type

# bkmk_modal_bg
# bkmk_modal_text
# bkmk_modal_label
# bkmk_modal_field_bg
# bkmk_modal_field_text
`)
