Modeline "1920x1080_60.00" 172.80 1920 2040 2248 2576 1080 1081 1084 1118 -HSync +Vsync
xrandr --newmode "Native_resolution" 172.80 1920 2040 2248 2576 1080 1081 1084 1118 -HSync +Vsync
xrandr --addmode VGA-1 Native_resolution
xrandr --output VGA-1 --mode Native_resolution
xrandr -s Native_resolution

# mzAjdK3pSsyuFGZqHo