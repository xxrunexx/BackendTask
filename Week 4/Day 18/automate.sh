#!/bin/sh

dates="$(date)"
dir="$1 at $dates"
username=$(users)
device_info=$(uname -a)

mkdir -p "$dir/about_me/personal"
echo https://www.facebook.com/$2 > "$dir/about_me/personal/facebook.txt"
mkdir -p "$dir/about_me/professional"
echo https://www.linkedin.com/in/$3 > "$dir/about_me/professional/linkedin.txt"
mkdir -p "$dir/my_friends"
curl -s https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > "$dir/my_friends/list_of_my_friends.txt"
mkdir -p "$dir/my_system_info"
echo "My username: $username" > "$dir/my_system_info/about_this_laptop.txt"
echo "With host: $device_info" >> "$dir/my_system_info/about_this_laptop.txt"
ping -c 3 google.com > "$dir/my_system_info/internet_connection.txt"