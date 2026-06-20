module Const where

-- Bro what is this as for saving the path we need extra import hahah
path :: FilePath
path = "../../overview/pixel_art_sketches/"

-- The text for box that differes per the colour
black :: String
black = "\"outline\""

-- Couple that to be used later (not in d6 / d20)
darkspot :: String
darkspot = "\"darkspot\""

between :: String
between = "\"between\""

-- Back to text used before
shadow :: String
shadow = "\"shadow\""

slight :: String
slight = "\"small_light\""

gbase :: String
gbase = "\"grey_base\""

white :: String
white = "\"inner_white\""

-- Hardcoded data about making SVG
start :: String
start = "<svg width=\"64\" height=\"64\" viewBox=\"0 0 16 16\">"

end :: String
end = "</svg>"

-- Making an rectangle
recStart :: String
recStart = "<rect class="

recEnd :: String
recEnd = "/>"

-- left most corner on top x 0 y 0, width grows right, height grows down
data Rectangle = Rectangle {x :: Int, y :: Int, width :: Int, height :: Int, color :: String}
