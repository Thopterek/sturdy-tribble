module Main where

import Codec.Picture
import Const
import System.FilePath ((</>))

-- Way to classify all of the colours used in Pixel Arts
classification :: PixelRGB8 -> String
classification (PixelRGB8 r g b)
    | r < 40 && g < 40 && b < 40 = black
    | r == 49 && g == 49 && b == 49 = darkspot
    | r == 84 && g == 85 && b == 85 = between
    | r == 121 && g == 120 && b == 120 = shadow
    | r == 161 && g == 160 && b == 160 = slight
    | r == 213 && g == 212 && b == 212 = gbase
    | r > 233 && g > 233 && b > 233 = white
    | otherwise = "\"ErrorBruh\""

-- Main caller? To make it like in C it would be
-- int writeRectangles(Image img, int w, int h)
-- where is for local variables and the recursion?!
-- writeRectangles :: Image PixelRGB8 -> Int -> Int
-- writeRectangles total = rec img w h
--  where
--    total = h * w
--    rec img  w h

-- Haskell, why are you using tabs instead of spaces :c
-- plus very random but the idea is: return value is the fn
-- int skrr(int x) {return (x * x);} and Haskell below
-- skrr :: Int -> Int
-- skrr x = x * x
main :: IO ()
main = do
    result <- readImage (path </> "bird.png")
    case result of
        Left err -> putStrLn $ "Error: " ++ err
        Right img -> do
            let w = dynamicMap imageWidth img
            let h = dynamicMap imageHeight img
            -- writeRectangles img w h
            putStrLn $ "Width: " ++ show w ++ ", Height: " ++ show h
