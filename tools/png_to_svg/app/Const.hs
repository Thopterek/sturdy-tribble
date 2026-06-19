module Const where

-- Bro what is this as for saving the path we need extra import hahah
path :: FilePath
path = "../../overview/pixel_art_sketches/"

-- The text for box that differes per the colour
black :: String
black = "outline"

-- Couple that to be used later (not in d6 / d20)
darkspot :: String
darkspot = "darkspot"
between :: String
between = "between"

-- Back to text used before
shadow :: String
shadow = "shadow"
slight :: String
slight = "small_light"
gbase :: String
gbase = "grey_base"
white :: String
white = "inner_white"

-- Making an rectangle
recStart :: String
recStart = "<rect class="
recEnd :: String
recEnd = "/>"
