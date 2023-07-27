# renamer

```sh
go install git.sr.ht/~mendelmaleh/renamer@latest
```

Rename files with patterns.

For example, have a downloaded youtube playlist, the format is long, messy, and it sorts by title:
```sh
air ~/d/y/atfc $ ls
ATFC： Advancing To Black Steel And Train Track Making Line (Episode 33) [Wzcg0Ys8AMI].webm
ATFC： Advancing To Steel Age (Episode 26) [ywlr8JwaZPM].webm
ATFC： Automatic Cow Milking Machine (Episode 28) [WJeowXn9dbc].webm
ATFC： Big Rotating Mining Machine (Episode 32) [x_IcGbwqRLM].webm
ATFC： Finally Getting A Lava Source Close To Base And Big Windmill (Episode 30) [fLbssYZk5OM].webm
ATFC： Getting Into Create Mod Trains (Episode 31) [D-GNsO4EdoQ].webm
ATFC： Making A Fancy Barn Door (Episode 29) [nQSXLCc2bpM].webm
ATFC： Progressing To The Final Metal Tier (Episode 34) [3bHakTp5Ilk].webm
ATFC： Raw Stone Mining Tunnel Bore (Episode 27) [cbmHifveUlA].webm
ATFC： Testing Powder Kegs And Finding Diamonds (Episode 35) [Lzo2CNckyx4].webm
ATFC： Whiskey Cobble Generator  (Episode 36) [4FqhD2LWflA].webm
```

Set the sorting and naming patterns, and see what the transformation would
be. The sorting match group is put in the beginning of the filename, followed
by a dash, and the naming match in snake case:
```sh
air ~/d/y/atfc $ renamer -sort 'Episode (\d+)' -name 'ATFC： (.+)\(' | sort
26-advancing_to_steel_age.webm <- ATFC： Advancing To Steel Age (Episode 26) [ywlr8JwaZPM].webm
27-raw_stone_mining_tunnel_bore.webm <- ATFC： Raw Stone Mining Tunnel Bore (Episode 27) [cbmHifveUlA].webm
28-automatic_cow_milking_machine.webm <- ATFC： Automatic Cow Milking Machine (Episode 28) [WJeowXn9dbc].webm
29-making_a_fancy_barn_door.webm <- ATFC： Making A Fancy Barn Door (Episode 29) [nQSXLCc2bpM].webm
30-finally_getting_a_lava_source_close_to_base_and_big_windmill.webm <- ATFC： Finally Getting A Lava Source Close To Base And Big Windmill (Episode 30) [fLbssYZk5OM].webm
31-getting_into_create_mod_trains.webm <- ATFC： Getting Into Create Mod Trains (Episode 31) [D-GNsO4EdoQ].webm
32-big_rotating_mining_machine.webm <- ATFC： Big Rotating Mining Machine (Episode 32) [x_IcGbwqRLM].webm
33-advancing_to_black_steel_and_train_track_making_line.webm <- ATFC： Advancing To Black Steel And Train Track Making Line (Episode 33) [Wzcg0Ys8AMI].webm
34-progressing_to_the_final_metal_tier.webm <- ATFC： Progressing To The Final Metal Tier (Episode 34) [3bHakTp5Ilk].webm
35-testing_powder_kegs_and_finding_diamonds.webm <- ATFC： Testing Powder Kegs And Finding Diamonds (Episode 35) [Lzo2CNckyx4].webm
36-whiskey_cobble_generator.webm <- ATFC： Whiskey Cobble Generator  (Episode 36) [4FqhD2LWflA].webm
```

To rename the files, rerun the command with `-yes`:
```sh
air ~/d/y/atfc $ renamer -sort 'Episode (\d+)' -name 'ATFC： (.+)\(' -yes
air ~/d/y/atfc $ ls
26-advancing_to_steel_age.webm                                        32-big_rotating_mining_machine.webm
27-raw_stone_mining_tunnel_bore.webm                                  33-advancing_to_black_steel_and_train_track_making_line.webm
28-automatic_cow_milking_machine.webm                                 34-progressing_to_the_final_metal_tier.webm
29-making_a_fancy_barn_door.webm                                      35-testing_powder_kegs_and_finding_diamonds.webm
30-finally_getting_a_lava_source_close_to_base_and_big_windmill.webm  36-whiskey_cobble_generator.webm
31-getting_into_create_mod_trains.webm
```
